package main

import (
	"log"
	"net/http"

	"github.com/arcadefx/schnell/routes"
)

func main() {
	routes.UseRoutes()

	log.Fatal(http.ListenAndServe(":3000", nil))
}
