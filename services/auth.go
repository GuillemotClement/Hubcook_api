package services

import (
	"net/http"

	"github.com/GuillemotClement/Hubcook_api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// register
func CreateUser(c *gin.Context) {
	// variable recupere les data JSON provenant de la requete
	var input models.User

	// verification de la conversion du JSON et on le mets dans la variable input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Format invalide"})
	}

	// recuperation du user pour verifier si il existe deja
	var existingUser models.User
	err := db.QueryRow(`SELECT username, email FROM users WHERE username = $1`, input.Username).Scan(&existingUser.Username, &existingUser.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Erreur lors du traintement de l'inscription"})
		return
	}

	// verification si email vide alors email non utiliser
	if existingUser.Email != "" || existingUser.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Utilisateur existant"})
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
	if err := db.QueryRow(`INSERT INTO users(username, email, password) VALUES($1, $2, $3) RETURNING id`, input.Username, input.Email, input.Password).Scan(&userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erreur": "Echec de l'enregistrement de l'utilisateur"})
		return
	}
	// generation du token

	// retourne le token
	c.JSON(http.StatusAccepted, gin.H{"message": "user created successfully"})
}

// login
