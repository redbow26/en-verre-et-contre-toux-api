package server

import "github.com/gofiber/fiber/v2"

func (server *Server) Use(handler fiber.Handler) {
	server.app.Use(handler)
}
