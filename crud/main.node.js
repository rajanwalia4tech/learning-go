const express = require('express');
const app = express();
const port = 7000;

const { Pool } = require('pg');
const pool = new Pool({
  user: 'go_book_user',
  host: 'localhost',
  database: 'go_book_db',
  password: '123456',
  port: 5432,
});

app.use(express.json());

app.get('/books', async (req, res) => {
  try {
    const { rows } = await pool.query('SELECT id, title, author, price FROM books');
    res.json(rows);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

app.get('/books/:id', async (req, res) => {
  const { id } = req.params;
  try {
    const { rows } = await pool.query('SELECT id, title, author, price FROM books WHERE id = $1', [id]);
    if (rows.length === 0) {
      res.status(404).json({ error: 'Book not found' });
    } else {
      res.json(rows[0]);
    }
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

app.post('/books', async (req, res) => {
  const { title, author, price } = req.body;
  try {
    const { rows } = await pool.query('INSERT INTO books (title, author, price) VALUES ($1, $2, $3) RETURNING id', [title, author, price]);
    const id = rows[0].id;
    res.status(201).json({ id, title, author, price });
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

app.put('/books/:id', async (req, res) => {
  const { id } = req.params;
  const { title, author, price } = req.body;
  try {
    const result = await pool.query('UPDATE books SET title = $1, author = $2, price = $3 WHERE id = $4', [title, author, price, id]);
    if (result.rowCount === 0) {
      res.status(404).json({ error: 'Book not found' });
    } else {
      res.json({ message: 'Book updated successfully' });
    }
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

app.delete('/books/:id', async (req, res) => {
  const { id } = req.params;
  try {
    const result = await pool.query('DELETE FROM books WHERE id = $1', [id]);
    if (result.rowCount === 0) {
      res.status(404).json({ error: 'Book not found' });
    } else {
      res.json({ message: 'Book deleted successfully' });
    }
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
