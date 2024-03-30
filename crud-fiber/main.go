package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://go_book_user:123456@localhost/go_book_db?sslmode=disable") // Replace with your PostgreSQL connection details
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/books", getAllBooks)
	app.Get("/books/:id", getBookByID)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	log.Fatal(app.Listen(":8000"))
}

func getAllBooks(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, title, author, price FROM books")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		books = append(books, book)
	}

	return c.JSON(books)
}

func getBookByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Book
	err := db.QueryRow("SELECT id, title, author, price FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author, &book.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(book)
}

func createBook(c *fiber.Ctx) error {
	var book Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, err := db.Exec("INSERT INTO books (title, author, price) VALUES ($1, $2, $3)", book.Title, book.Author, book.Price)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, _ := result.LastInsertId()
	book.ID = int(id)

	return c.Status(fiber.StatusCreated).JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := db.Exec("UPDATE books SET title = $1, author = $2, price = $3 WHERE id = $4", book.Title, book.Author, book.Price, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Book updated successfully",
	})
}

func deleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
