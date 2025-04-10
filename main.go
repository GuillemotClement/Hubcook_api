package main

import (
	"github.com/GuillemotClement/Hubcook_api/services"
)

func main() {

	// port := ":8086"

	services.InitDB()
	services.GenerateFileSeedingDB()
	// services.GenerateUser(5)

	// initialisation du router
	// r := services.SetupRouter()

	// autorisation du proxy
	// r.SetTrustedProxies(nil) // retire l'erreur

	// lancement du serveur
	// r.Run(port)
}
