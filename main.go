package main

import (
	"log"
	"net/http"
	"github.com/gustavolvjardim/biblioteca-api/database"
	"github.com/gustavolvjardim/biblioteca-api/models"
	"github.com/gin-gonic/gin"
)


// Rota de saúde (pra saber se a API tá viva)
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "API rodando perfeitamente"})
}

// Lista todos os livros
// Lista todos os livros
func getBooks(c *gin.Context) {
    var books []models.Book // Declara uma slice da struct do pacote models

    // Usa o GORM para buscar todos os registros na tabela 'books'
    // e preencher a slice 'books'
    database.DB.Find(&books) 
    
    // Verifica se a busca foi bem-sucedida
    if len(books) == 0 {
         // Se não encontrar livros, retorna 204 No Content (ou 200 com array vazio)
        c.IndentedJSON(http.StatusOK, []models.Book{})
        return
    }

    c.IndentedJSON(http.StatusOK, books)
}


// Busca livro por ID
func getBookByID(c *gin.Context) {
    id := c.Param("id")

    var book models.Book // Struct para onde o resultado será carregado

    // Usa o GORM para buscar o primeiro registro que satisfaça a condição (ID = id)
    result := database.DB.First(&book, id) 

    if result.Error != nil {
        // Se o erro for "record not found" (registro não encontrado)
        if result.Error.Error() == "record not found" {
            c.IndentedJSON(http.StatusNotFound, gin.H{"error": "livro não encontrado"})
            return
        }
        // Outro erro de DB
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erro no banco de dados"})
        return
    }

    c.IndentedJSON(http.StatusOK, book)
}

// Cria um novo livro
func createBook(c *gin.Context) {
    var newBook models.Book // Usa a struct do pacote models

    // 1. Faz o bind (deserializa o JSON)
    if err := c.BindJSON(&newBook); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
        return
    }

    // 2. Salva o novo livro no banco de dados
    // GORM irá popular automaticamente o campo ID (se definido com primaryKey)
    result := database.DB.Create(&newBook)
    
    if result.Error != nil {
        log.Printf("Erro ao criar livro: %v", result.Error)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Falha ao salvar o livro no DB"})
        return
    }
    
    // 3. Retorna o livro criado (agora com o ID gerado pelo DB)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {

	log.Println("Iniciando a API da Biblioteca...")

	database.InitDB()
	router := gin.Default()

	// ROTAS DO DIA 1
	router.GET("/health", healthCheck)      // verifica se tá vivo
	router.GET("/books", getBooks)          // lista todos
	router.GET("/books/:id", getBookByID)   // busca por ID
	router.POST("/books", createBook)       // cria novo

	router.Run("localhost:8080")
}