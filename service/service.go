package service

import (
	"github.com/4halavandala/go_restAPI/model"
	"github.com/4halavandala/go_restAPI/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, pasword string) (string, error)
	PasrseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId, listId int) (model.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
