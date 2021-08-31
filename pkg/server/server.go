package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func New(config ...fiber.Config) *Server {
	return &Server{
		app: fiber.New(config...),
	}
}
