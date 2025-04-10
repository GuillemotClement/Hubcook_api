package main

import (
	"github.com/GuillemotClement/Hubcook_api/services"
)

func main() {

	port := ":8086"

	services.InitDB()

	// generation du seed des data en BDD
	services.GenerateFileSeedingDB()

	// initialisation du router
	r := services.SetupRouter()

	// autorisation du proxy
	r.SetTrustedProxies(nil) // retire l'erreur

	// lancement du serveur
	r.Run(port)
}
