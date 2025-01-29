document.addEventListener('DOMContentLoaded', () => {
    const bookForm = document.getElementById('book-form');
    const booksList = document.getElementById('books-list');

    // Handle form submission
    bookForm.addEventListener('submit', async (event) => {
        event.preventDefault();

        const book = {
            title: document.getElementById('title').value,
            author: document.getElementById('author').value,
            published_year: document.getElementById('published_year').value,
            isbn: document.getElementById('isbn').value,
            description: document.getElementById('description').value
        };

        try {
            // POST request to create a new book
            const response = await fetch('http://localhost:8080/books', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(book)
            });

            const data = await response.json();

            if (response.ok) {
                // Clear form
                bookForm.reset();
                // Reload books list
                loadBooks();
            } else {
                alert('Error adding book');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    });

    // Load all books
    async function loadBooks() {
        try {
            const response = await fetch('http://localhost:8080/books');
            const books = await response.json();

            // Clear the current list
            booksList.innerHTML = '';

            // Append each book to the list
            books.forEach(book => {
                const li = document.createElement('li');
                li.classList.add('list-group-item');
                li.innerHTML = `
                    <strong>${book.title}</strong><br>
                    Author: ${book.author}<br>
                    Year: ${book.published_year}<br>
                    ISBN: ${book.isbn}<br>
                    <button class="btn btn-warning btn-sm mt-2" onclick="updateBook('${book.id}')">Edit</button>
                    <button class="btn btn-danger btn-sm mt-2" onclick="deleteBook('${book.id}')">Delete</button>
                `;
                booksList.appendChild(li);
            });
        } catch (error) {
            console.error('Error:', error);
        }
    }

    async function deleteBook(id) {
        if (!confirm("Are you sure you want to delete this book?")) return;
    
        try {
            const response = await fetch(`http://localhost:8080/books/${id}`, {
                method: 'DELETE'
            });
    
            if (response.ok) {
                alert("Book deleted successfully");
                loadBooks(); // Refresh book list after deletion
            } else {
                const errorData = await response.json();
                alert("Error: " + errorData.error);
            }
        } catch (error) {
            console.error("Error:", error);
        }
    }
    
    async function updateBook(id) {
        const newTitle = prompt("Enter new title:");
        const newAuthor = prompt("Enter new author:");
        const newYear = prompt("Enter new published year:");
        const newISBN = prompt("Enter new ISBN:");
    
        if (!newTitle || !newAuthor || !newYear || !newISBN) {
            alert("All fields are required!");
            return;
        }
    
        const updatedBook = {
            title: newTitle,
            author: newAuthor,
            published_year: parseInt(newYear),
            isbn: newISBN
        };
    
        try {
            const response = await fetch(`http://localhost:8080/books/${id}`, {
                method: 'PUT',
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(updatedBook)
            });
    
            if (response.ok) {
                alert("Book updated successfully");
                loadBooks(); // Refresh book list after update
            } else {
                const errorData = await response.json();
                alert("Error: " + errorData.error);
            }
        } catch (error) {
            console.error("Error:", error);
        }
    }
    

    // Load books when the page loads
    loadBooks();
});
