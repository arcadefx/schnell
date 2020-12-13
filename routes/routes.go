package routes

import (
	"net/http"

	"github.com/arcadefx/schnell/controllers"
)

// UseRoutes defines http paths to handle
func UseRoutes() {
	http.HandleFunc("/api/auth", func(w http.ResponseWriter, r *http.Request) {
		controllers.Authenticate(w, r)
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))
}
