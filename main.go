package main

import (
	"fmt"
	"rhSystem_server/app/infrastructure"
	"rhSystem_server/app/infrastructure/database"
	//_ "rhSystem/docs"
)

func main() {
	fmt.Println("rhSystem server!")

	database.Init()

	infrastructure.InitServer()
}
