package model

type TodoList struct {
	Id          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int    `json:"id"`
	UserId string `json:"UserId"`
	ListId string `json:"ListId"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int    `json:"id"`
	ListId string `json:"ListId"`
	ItemId string `json:"ItemId"`
}
