package infrastructure

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"rhSystem_server/app/infrastructure/routes"

	"github.com/rs/cors"
)

func InitServer() {
	godotenv.Load(".env")

	//portString := os.Getenv("PORT")

	// if portString == "" {
	// 	log.Fatal("Port is not found in the environment")
	// }

	handler := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedOrigins: []string{"*"},
	}).Handler(routes.Router())

	//fmt.Println("port: ", portString)

	//log.Fatal(http.ListenAndServe(": "+"8080", handler))
	log.Fatal(http.ListenAndServe(":8080", handler))
}
