package main

import (
	"bit-ly/database"
	"bit-ly/routes"
)

func main() {
	database.ConnectDatabase()
	routes.ConfigureRoutes()
}
