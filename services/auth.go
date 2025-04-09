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

	log.Println(err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Echec de l'enregistrement de l'utilisateur"})
		return
	}

	// retourne le token
	c.JSON(http.StatusAccepted, gin.H{"message": "user created successfully"})
}

// login
func Login(c *gin.Context) {
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
	err := db.QueryRow(`SELECT password, role_id FROM users WHERE username = $1`, input.Username).Scan(&hashPassword, &roleId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur lors du traitement du login"})
		return
	}

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"Erreur": "Echec de l'authentification"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Erreur": "Echec de l'authentification"})
		return
	}

	token, err := GenerateToken(input.Username, roleId)

	log.Printf("Token : %v", token)
	log.Printf("Err token : %v", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Echec du generation du token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "connexion reussis", "token": token})

}

func GenerateToken(username string, role uint) (string, error) {
	// definition du temps d'expiration du token
	expirationTime := time.Now().Add((24 * 7) * time.Hour)

	log.Print(expirationTime)

	claims := &models.Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	log.Print(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Print(token)
	return token.SignedString(jwtKey)
}

// middleware authentification
func AuthMiddleware(c *gin.Context) {
	// recuperation du token depuis le header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Token manquant"})
		return
	}

	claims := &models.Claims{}

	// recuperation du token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Token invalide"})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Set("role_id", claims.Role)
	c.Next()
}
