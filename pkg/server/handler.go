package server

import "github.com/redbow26/en-verre-et-contre-toux-api/pkg/activity"

func (server *Server) RegisterHandler() {
	router := server.app.Group("v1")

	activity.RegisterRouter(router)
}
