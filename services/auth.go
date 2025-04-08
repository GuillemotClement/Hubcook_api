package services

import (
	"net/http"

	"github.com/GuillemotClement/Hubcook_api/models"
	"github.com/gin-gonic/gin"
)

// register
func CreateUser(c *gin.Context) {
	// variable recupere les data JSON provenant de la requete
	var input models.User

	// verification de la conversion du JSON et on le mets dans la variable input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erreur": "Format invalide"})
	}

	// verification si l'user existe deja

	// hash du mot de passe

	// enregistrement de l'user en db

	// recuperation de l'user

	// generation du token

	// retourne le token
}

// login
