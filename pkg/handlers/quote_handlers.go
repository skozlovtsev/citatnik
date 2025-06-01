package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/skozlovtsev/citatnik/pkg/controller"
	"github.com/skozlovtsev/citatnik/pkg/model"
)

func HandleAddQuote(w http.ResponseWriter, r *http.Request) {
	var quote model.Quote
	var body []byte
	var err error

	if body, err = io.ReadAll(r.Body); err != nil {
		return
	}
	if err := json.Unmarshal(body, &quote); err != nil {
		return
	}
	quote = controller.AddQuote(quote)
	result, _ := json.Marshal(quote)

	w.Write([]byte(result))
}

func HandleGetQuote(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author != "" {
		body, _ := json.Marshal(controller.GetQuoteByAuthor(author))

		w.Write([]byte(body))
		return
	}

	body, _ := json.Marshal(controller.GetQuotes())

	w.Write([]byte(body))
}

func HandleGetRandomQuote(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal(controller.GetRandomQuote())

	w.Write([]byte(body))
}

func HandleDeleteQuoteById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		w.Write([]byte(""))
		return
	}

	body, _ := json.Marshal(controller.DeleteQuoteById(id))
	w.Write([]byte(body))
}
