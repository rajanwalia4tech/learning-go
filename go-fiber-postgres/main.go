package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"goorm.io/gorm"
)

type Repository struct {
	// db *sql.DB
	DB *gorm.DB
}

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func (r *Repository) SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Post("/book", r.CreateBook)
	api.Get("/books", r.GetBooks)
	api.Get("/book/:id", r.GetBook)
	api.Put("/book/:id", r.UpdateBook)
	api.Delete("/book/:id", r.DeleteBook)

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatalf("Error connecting to database")
	}

	r := Repository{
		DB: db,
	}

	// create a fiber server
	app := fiber.New()
	r.SetupRoutes(app)
	// listen on port 3000
	app.Listen(":8000")

}
