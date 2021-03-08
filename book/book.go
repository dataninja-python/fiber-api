package book

import "github.com/gofiber/fiber"

func GetBooks(c *fiber.Ctx) {
	c.Send("All the books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("A single book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Add a new book")
}

func UpdateBook(c *fiber.Ctx) {
	c.Send("Update a book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete a book")
}
