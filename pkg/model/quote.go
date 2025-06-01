package model

// Quote - Основная структура данных с которой работаем
type Quote struct {
	Id     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}
