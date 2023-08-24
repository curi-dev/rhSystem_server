package infrastructure

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"rhSystem_server/app/infrastructure/routes"
)

func InitServer() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}

	log.Fatal(http.ListenAndServe(": "+portString, routes.Router()))
}
