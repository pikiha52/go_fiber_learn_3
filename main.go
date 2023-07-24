package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"clean_architecture/api/routes"
	"clean_architecture/config"
	"clean_architecture/pkg/user"
)

func main() {
	app := fiber.New()

	database := ConnectDB()
	userRepo := user.NewRepo(database)
	userService := user.NewService(userRepo)

	api := app.Group("/api", logger.New())
	routes.UserRoutes(api, userService)

	app.Listen(":3000")
}

func ConnectDB() *gorm.DB {
	var DB *gorm.DB

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error parsing DB port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to open database!")
	}

	fmt.Println("Connection Opened to Database: ")

	return DB

}
