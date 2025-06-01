package main

import (
	"fmt"
	"net/http"

	"github.com/skozlovtsev/citatnik/pkg/routes"
)

const (
	port = ":8080"
)

func main() {
	mux := http.NewServeMux()
	routes.Register(mux)
	fmt.Println("Server listened on http://localhost" + port)
	fmt.Println(http.ListenAndServe(":8080", mux))
}
