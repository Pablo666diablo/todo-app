package service

import (
	"github.com/Pablo666diablo/todo-app/pkg/repository"
)

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(*repository.Repository) *Service {
	return &Service{}
}
