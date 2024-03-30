package model

type Book struct {
	Title     string
	Author    string
	Publisher string
}

func (b *Book) Create() error {
	_, err := db.Exec("INSERT INTO books (title, author, publisher) VALUES ($1, $2, $3)", b.Title, b.Author, b.Publisher)
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) Update() error {
	_, err := db.Exec("UPDATE books SET author=$1, publisher=$2 WHERE title=$3", b.Author, b.Publisher, b.Title)
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) Delete() error {
	_, err := db.Exec("DELETE FROM books WHERE title=$1", b.Title)
	if err != nil {
		return err
	}
	return nil
}

func GetAllBooks() ([]Book, error) {
	rows, err := db.Query("SELECT title, author, publisher FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Title, &book.Author, &book.Publisher)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
