package activity

import "github.com/gofiber/fiber/v2"

func RegisterRouter(router fiber.Router) {
	r := router.Group("activities")

	r.Get("/", getActivities)
	r.Get("/:id", getActivity)
	r.Post("/", createActivity)
	r.Put("/:id", updateActivity)
	r.Delete("/:id", deleteActivity)
}
