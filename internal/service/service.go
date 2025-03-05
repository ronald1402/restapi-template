package service

import (
	"restapi/internal/repository"
)

type Service interface {
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
