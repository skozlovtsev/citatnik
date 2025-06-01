package storage

import (
	. "github.com/skozlovtsev/citatnik/pkg/model"
)

type InMemoryStorage struct {
	LastId      int
	NextId      int
	Storage     []*Quote
	AuthorIndex map[string][]int
}

func NewInMemoryStorage() *InMemoryStorage {
	ims := &InMemoryStorage{
		LastId: 0,
		NextId: 0,
	}
	ims.Storage = make([]*Quote, 0)
	ims.AuthorIndex = make(map[string][]int)
	return ims
}

func (ims *InMemoryStorage) Add(value Quote) int {
	quote := value
	quote.Id = ims.NextId
	ims.Storage = append(ims.Storage, &quote)
	ims.AddAuthorIndex(value.Author, ims.NextId)

	ims.LastId = ims.NextId
	ims.NextId += 1
	return ims.LastId
}

func (ims *InMemoryStorage) GetById(id int) Quote {
	if id > ims.LastId {
		return Quote{}
	}

	value := ims.Storage[id]
	if value == nil {
		return Quote{}
	}
	return *value
}

func (ims *InMemoryStorage) GetAll() []Quote {
	var result []Quote
	for _, v := range ims.Storage {
		if v == nil {
			continue
		}
		result = append(result, *v)
	}
	return result
}

func (ims *InMemoryStorage) GetByAuthor(author string) []Quote {
	value, ok := ims.AuthorIndex[author]
	if !ok {
		return []Quote{}
	}
	var quotes []Quote
	for _, v := range value {
		quotes = append(quotes, ims.GetById(v))
	}
	return quotes
}

func (ims *InMemoryStorage) DeleteById(id int) Quote {
	if id > ims.LastId {
		return Quote{}
	}
	quote := ims.Storage[id]
	if quote == nil {
		return Quote{}
	}
	ims.DeleteAuthorIndex(quote.Author, id)
	ims.Storage[id] = nil
	return *quote
}

func (ims *InMemoryStorage) SearchAuthorIndex(author string) []int {
	value, ok := ims.AuthorIndex[author]
	if !ok {
		return []int{}
	}
	return value
}

func (ims *InMemoryStorage) AddAuthorIndex(author string, id int) {
	_, ok := ims.AuthorIndex[author]
	if ok {
		ims.AuthorIndex[author] = append(ims.AuthorIndex[author], id)
		return
	}
	ims.AuthorIndex[author] = []int{id}
}

func (ims *InMemoryStorage) DeleteAuthorIndex(author string, id int) int {
	value, ok := ims.AuthorIndex[author]
	if !ok {
		return -1
	}
	var index int
	for i, v := range value {
		if v == id {
			ims.AuthorIndex[author] = append(ims.AuthorIndex[author][:i], ims.AuthorIndex[author][i+1:]...)
			index = i
			break
		}
	}
	if len(ims.AuthorIndex[author]) == 0 {
		delete(ims.AuthorIndex, author)
	}
	return index
}

// Экземпляр InMemoryStorage
var Storage *InMemoryStorage

func init() {
	Storage = NewInMemoryStorage()
}
