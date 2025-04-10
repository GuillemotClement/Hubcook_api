package services

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/GuillemotClement/Hubcook_api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("mimix")

// register
func CreateUser(c *gin.Context) {
	// variable recupere les data JSON provenant de la requete
	var input models.User
	input.RoleId = 2
	// verification de la conversion du JSON et on le mets dans la variable input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Format invalide"})
		return
	}

	// recuperation du user pour verifier si il existe deja
	var existingUser models.User
	err := db.QueryRow(`SELECT username FROM users WHERE username = $1`, input.Username).Scan(&existingUser.Username)

	if err != sql.ErrNoRows && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur rencontrer dans le traitement de l'insription pour username"})
		return
	}

	// verification si username vide
	if existingUser.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Username deja utiliser"})
		return
	}

	err = db.QueryRow(`SELECT email FROM users WHERE email = $1`, input.Email).Scan(&existingUser.Email)

	if err != sql.ErrNoRows && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur rencontrer dans le traitement de l'insription pour l'email"})
		return
	}

	// verification si email vide alors email non utiliser
	if existingUser.Email != "" || existingUser.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Email deja utiliser"})
		return
	}

	// hash du mot de passe
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Echec du hash du mot de passe"})
		return
	}

	input.Password = string(hashed)

	// enregistrement de l'user en db
	// on prepare la variable pour recuperer l'id
	var userId int
	err = db.QueryRow(`INSERT INTO users(username, email, password, image, role_id) VALUES($1, $2, $3, $4, $5) RETURNING id`, input.Username, input.Email, input.Password, input.Image, input.RoleId).Scan(&userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Echec de l'enregistrement de l'utilisateur"})
		return
	}

	// retourne le token
	c.JSON(http.StatusAccepted, gin.H{"message": "user created successfully"})
}

// login
func Login(c *gin.Context) {
	w := c.Writer
	// on prepare la struct pour recuperer les donnees du json
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Format invalide"})
		return
	}

	var hashPassword string
	var roleId uint
	var userId uint
	err := db.QueryRow(`SELECT id, password, role_id FROM users WHERE username = $1`, input.Username).Scan(&userId, &hashPassword, &roleId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur lors du traitement du login"})
		return
	}

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"Erreur": "Echec de l'authentification"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Erreur": "Echec de compare mdp"})
		return
	}

	token, err := GenerateToken(userId, roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Echec du generation du token"})
		return
	}

	var userInfo struct {
		Id       uint   `json:"id"`
		Username string `json:"username"`
		Image    string `json:"image"`
		Role     string `json:"role"`
		Email    string `json:"email"`
	}

	if err := db.QueryRow(`SELECT id, username, email, image, role_id FROM users WHERE username = $1`, input.Username).Scan(&userInfo.Id, &userInfo.Username, &userInfo.Email, &userInfo.Image, &userInfo.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Echec de recuperation des infos de l'utilisateur"})
		return
	}

	// generation du cookie pour fournir le token
	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Secure = false
	cookie.SameSite = http.SameSiteDefaultMode
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(w, &cookie)

	c.JSON(http.StatusOK, gin.H{"message": "authentification reussie", "userInfo": userInfo})
}

func GenerateToken(id uint, role uint) (string, error) {
	// definition du temps d'expiration du token
	expirationTime := time.Now().Add((24 * 7) * time.Hour)

	claims := &models.Claims{
		Id:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

// middleware authentification
func AuthMiddleware(c *gin.Context) {

	tokenString, err := c.Request.Cookie("token")
	// recuperation du token depuis le header

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Erreur token"})
		return
	}

	if tokenString.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Token manquant"})
		return
	}

	claims := &models.Claims{}

	// recuperation du token
	token, err := jwt.ParseWithClaims(tokenString.Value, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Token invalide"})
		c.Abort()
		return
	}

	c.Set("userId", claims.Id)
	c.Set("roleId", claims.Role)
	c.Next()
}

func Logout(c *gin.Context) {
	// recuperation du writter depuis le context Gin
	w := c.Writer

	// on tente de recuperer le cookie token
	cookieResult, err := c.Request.Cookie("token")

	log.Printf("Cookie recu : %v", cookieResult)

	if err != nil {
		if err == http.ErrNoCookie {
			log.Printf("Erreur cookie non present :%v", err)
			// pas de cookie
			c.JSON(http.StatusBadRequest, gin.H{"message": "Aucun coookie trouver"})
			return
		}
		log.Printf("Seconde erreur : %v", err)
		// autre cas d'erreur
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur lors de la recuperation de la requete"})
		return
	}

	// on ecrase le cookie pour le supprimer.
	cookie := &http.Cookie{
		Name:     "token", // nom du cookie
		Value:    "",
		MaxAge:   -1,  // supression immediate
		Path:     "/", // meme path que lors de la creation
		HttpOnly: true,
	}

	http.SetCookie(w, cookie) // ecrasemeent de l'ancien cookie

	c.JSON(http.StatusOK, gin.H{"message": "Logout"})
}
