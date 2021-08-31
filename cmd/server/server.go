package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/oklog/pkg/group"
	"github.com/redbow26/en-verre-et-contre-toux-api/pkg/common/handler"
	"github.com/redbow26/en-verre-et-contre-toux-api/pkg/database"
	"github.com/redbow26/en-verre-et-contre-toux-api/pkg/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os/user"
	"time"
)

func main() {
	// Database
	dsn := "host=localhost user=test password=test dbname=test port=5432"
	err := database.Init(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = database.DB.AutoMigrate(&user.User{}, &group.Group{})
	if err != nil {
		log.Fatal(err)
	}

	// Fiber app
	app := server.New(fiber.Config{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		ErrorHandler: handler.ErrorHandler,
	})

	app.RegisterHandler()

	if err := app.ListenWithGraceful(":8080"); err != nil {
		log.Fatal(err)
	}

	app.CleanupTask()
}
