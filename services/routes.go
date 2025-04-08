package services

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	// initialisation du router
	r := gin.Default()

	// initialisation du middleware et du CORS

	// Route publique
	r.POST("/register")
	r.POST("login")

	return r
}
