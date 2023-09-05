package infrastructure

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"rhSystem_server/app/infrastructure/routes"

	"github.com/rs/cors"
)

func InitServer() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}

	handler := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}).Handler(routes.Router())

	fmt.Println("handler: ", handler)

	log.Fatal(http.ListenAndServe(": "+portString, handler))
}
