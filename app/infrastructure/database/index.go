package database

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"

	"sync"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

func GetDB() *sql.DB {
	return db
}

func Init() {
	once.Do(func() {

		godotenv.Load(".env")

		// user := os.Getenv("DATABASE_USER")
		// password := os.Getenv("DATABASE_PASSWORD")
		// host := os.Getenv("DATABASE_HOST")
		// database := os.Getenv("DATABASE_NAME")
		// sslMode := os.Getenv("SSL_MODE") // use "verify-full" in prod

		// dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", user, password, host, database, sslMode)
		dsn := os.Getenv("DATABASE_URL")

		fmt.Println("dsn: ", dsn)

		driver := "postgres"

		db, err = sql.Open(driver, dsn)

		if err != nil {
			log.Fatal(err)
		}

		err := db.Ping() // what do this function really does

		if err != nil {
			fmt.Println("ping err: ", err)

			panic("[database]: Ocorreu um erro no servidor")
		}
	})
}
