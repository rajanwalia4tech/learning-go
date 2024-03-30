const express = require('express');
const db = require('./db');

const app = express();

// Setup routes
app.get('/books', async (req, res) => {
  try {
    const { rows } = await db.query('SELECT * FROM books');
    res.json(rows);
  } catch (error) {
    console.error('Error retrieving books:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

// Add more routes for other CRUD operations (e.g., POST, PUT, DELETE)

module.exports = app;
