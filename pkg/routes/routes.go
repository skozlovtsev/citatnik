package routes

import (
	"net/http"

	"github.com/skozlovtsev/citatnik/pkg/handlers"
)

// Добавление маршрутов
func Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /quotes", handlers.HandleAddQuote)
	mux.HandleFunc("GET /quotes", handlers.HandleGetQuote)
	mux.HandleFunc("GET /quotes/random", handlers.HandleGetRandomQuote)
	mux.HandleFunc("DELETE /quotes/{id}", handlers.HandleDeleteQuoteById)
}
