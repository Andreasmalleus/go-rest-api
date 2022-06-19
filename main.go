package main

import (
	"github.com/Andreasmalleus/go-rest-api/routes"
)

func main() {
	router := routes.SetRouter()
	router.Run(":3000")
}
