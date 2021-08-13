package routes

import (
	"erbeyinn/book"
	"github.com/gofiber/fiber/v2"
)

func Route(route fiber.Router) {
	route.Get("", book.GetBooks)
	route.Post("", book.NewBook)
	route.Put("/:id",book.UpdateBook)
	route.Delete("/:id", book.DeleteBook)
	route.Get("/:id", book.GetBook)
}
