package controller

import (
	rand "math/rand/v2"

	"github.com/skozlovtsev/citatnik/pkg/model"
	"github.com/skozlovtsev/citatnik/pkg/storage"
)

func AddQuote(quote model.Quote) model.Quote {
	quote.Id = storage.Storage.Add(quote)
	return quote
}

func GetQuotes() []model.Quote {
	return storage.Storage.GetAll()
}

func GetRandomQuote() model.Quote {
	quotes := storage.Storage.GetAll()
	return quotes[rand.IntN(len(quotes))]
}

func GetQuoteByAuthor(author string) []model.Quote {
	return storage.Storage.GetByAuthor(author)
}

func DeleteQuoteById(id int) model.Quote {
	return storage.Storage.DeleteById(id)
}
