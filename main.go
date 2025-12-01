package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct que vai representar um livro (por enquanto em memória)
type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// Banco de dados fake (em memória)
var books = []Book{
	{ID: "1", Title: "Dom Casmurro", Author: "Machado de Assis", Quantity: 5},
	{ID: "2", Title: "1984", Author: "George Orwell", Quantity: 3},
}

// Rota de saúde (pra saber se a API tá viva)
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "API rodando perfeitamente"})
}

// Lista todos os livros
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Busca livro por ID
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "livro não encontrado"})
}

// Cria um novo livro
func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Gera um ID simples (depois vamos melhorar)
	newBook.ID = "999" // só pra hoje

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	// ROTAS DO DIA 1
	router.GET("/health", healthCheck)      // verifica se tá vivo
	router.GET("/books", getBooks)          // lista todos
	router.GET("/books/:id", getBookByID)   // busca por ID
	router.POST("/books", createBook)       // cria novo

	router.Run("localhost:8080")
}