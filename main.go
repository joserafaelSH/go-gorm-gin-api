package main

import (
	"github.com/joserafaelSH/go-gorm-gin-api/database"
	"github.com/joserafaelSH/go-gorm-gin-api/routes"
)

func main() {
	database.DataBaseConnect()
	
	routes.HandleRequest()
}