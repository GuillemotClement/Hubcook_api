package services

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// initialisation du router
	r := gin.Default()

	// initialisation du middleware et du CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Route publique
	r.POST("/register", CreateUser)
	r.POST("/login", Login)
	r.GET("/logout", Logout)

	// Route proteger
	protected := r.Group("/", AuthMiddleware)
	{
		protected.GET("/profil")
	}

	return r
}
