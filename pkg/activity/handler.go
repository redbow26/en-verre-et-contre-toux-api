package activity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redbow26/en-verre-et-contre-toux-api/pkg/activity/dto"
	"github.com/redbow26/en-verre-et-contre-toux-api/pkg/common/validation"
)

func getActivities(ctx *fiber.Ctx) error {
	activities := FindAll()
	return ctx.JSON(activities)
}

func getActivity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	activity := FindOne(id)

	if activity.ID == uuid.Nil {
		return fiber.NewError(fiber.StatusNotFound, "Activity "+id+" was not found")
	}
	return ctx.JSON(activity)
}

func createActivity(ctx *fiber.Ctx) error {
	data := new(dto.CreateActivityDto)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	errors := validation.Validate(*data)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	activity := Create(&Activity{
		Name:      data.Name,
		MaxPlayer: data.MaxPlayer,
	})

	if activity.ID == uuid.Nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Cannot create Activity")
	}
	return ctx.JSON(activity)
}

func updateActivity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data := new(dto.UpdateActivityDto)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	errors := validation.Validate(*data)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	activity := FindOne(id)

	activity.Update(&Activity{
		Name:      data.Name,
		MaxPlayer: data.MaxPlayer,
	})

	return ctx.JSON(activity)
}

func deleteActivity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	Delete(id)

	ctx.Status(fiber.StatusNoContent)
	return nil
}
