package main

import (
	"fmt"
	"go-web-api/src/config"
	"go-web-api/src/infra/persistence/database"
	"go-web-api/src/infra/persistence/migration"
)

func main() {
	fmt.Println("!... Hello World ...!")

	cfg := config.GetConfig()

	err := database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		fmt.Println(err.Error())
	}
	migration.Up1()

}
