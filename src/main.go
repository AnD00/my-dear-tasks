package main

import (
	db "github.com/AnD00/my-dear-tasks/src/databases"
	router "github.com/AnD00/my-dear-tasks/src/routes"
)

func main() {
	db.Initialize()
	defer db.Close()

	router.Router()
}
