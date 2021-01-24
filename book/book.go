package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kavin25/go-fiber/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks - Gets all the books
func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

// GetBook - Gets a single Book
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
	// return c.SendString("Single Books")
}

// NewBook - Creates a new book
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	result := db.Create(&book)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(book)
}

// DeleteBook - Deletes a book
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Book Found with ID")
	}
	db.Delete(&book, id)
	return c.JSON(fiber.Map{"message": "Success"})
}
