package main

import (
	"github.com/daniel/api-go-gin/database"
	"github.com/daniel/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
