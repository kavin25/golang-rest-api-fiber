package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kavin25/go-fiber/book"
	"github.com/kavin25/go-fiber/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/books", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("Connection to DB is open")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDB()

	setupRoutes(app)
	app.Listen(":3000")

	// defer database.DBConn.Close()
}
