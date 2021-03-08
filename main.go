package main

import (
	"fmt"

	"github.com/dataninja-python/fiber-api/book"
	"github.com/dataninja-python/fiber-api/database"

	"github.com/gofiber/fiber"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// func helloWorld(c *fiber.Ctx) {
// 	c.Send("Hi")
// }

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successful")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	// initialize the database
	initDatabase()
	// ensure database connection closes automatically when not needed
	defer database.DBConn.Close()

	// app.Use(cors.New())

	setupRoutes(app)

	app.Listen(3000)
}
