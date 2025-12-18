package api

import (
	"github.com/gofiber/fiber/v2"
	"job4j.ru/go-lang-base/internal/repository"
)

type Server struct {
	Repository *repository.RepoPg
}

func NewServer(repo *repository.RepoPg) *Server {
	return &Server{Repository: repo}
}

func (s *Server) Route(route fiber.Router) {
	route.Post("/item/", s.CreateItem)
	route.Get("/items/", s.GetItems)
}
