package book

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Book struct {
	Id       int    `json:"id"`
	Author   string `json:"author"`
	BookName string `json:"book_name"`
}

var books = []*Book{
	{
		Id:       1,
		Author:   "Kadri Küheylan",
		BookName: "Hasan Alinin Kırmızı Kurdelesi",
	},
	{
		Id:       2,
		Author:   "Bedri Kaynamaz",
		BookName: "Hasan Alinin Kırmızı Bandanası",
	},
}

func GetBooks(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"books": books})

}

func GetBook(ctx *fiber.Ctx) error {

	paramsId := ctx.Params("id")

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, book := range books {
		if book.Id == id {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"book": book})
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Not found"})
}

func NewBook(ctx *fiber.Ctx) error {
	type Request struct {
		Author   string `json:"author"`
		BookName string `json:"book_name"`
	}

	var body Request

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Cannot parse JSON",
		})
	}
	book := &Book{
		Id:       len(books) + 1,
		Author:   body.Author,
		BookName: body.BookName,
	}

	books = append(books, book)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"books": books, "message":"success"})
}

func UpdateBook(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}
	type Request struct {
		Author   *string `json:"author"`
		BookName *string `json:"book_name"`
	}
	var body Request
	err = ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	var updated *Book

	for _, update := range books {
		if update.Id == id {
			updated = update
			break
		}
	}



	if body.Author != nil {
		updated.Author = *body.Author
	}

	if body.BookName != nil {
		updated.BookName = *body.BookName
	}

	return ctx.Status(fiber.StatusOK).JSON(books)

}
func DeleteBook(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")

	id, err := strconv.Atoi(paramsId)

	if err!= nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Cannot parse id",
		})
	}

	for i,book := range books{
		if book.Id == id {
			books = append(books[:i], books[i+1:]...)

			return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"message":"Deleted Successfuly",
			})
		}
	}
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message":"Book not found",
	})
}
