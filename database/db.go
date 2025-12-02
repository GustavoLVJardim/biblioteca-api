package database

import (
	"log"
	"github.com/gustavolvjardim/biblioteca-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// variável global que armazena a instância do banco de dados.
var DB *gorm.DB

// conecta ao banco de dados e realiza a migração automática.
func InitDB() {
	var err error
	
	// Tenta abrir a conexão com o banco de dados SQLite
	DB, err = gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		// É melhor usar log.Fatalf ou panic, pois sem o DB o programa não pode continuar.
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso.")

	// Realiza a migração automática do(s) modelo(s)
	err = DB.AutoMigrate(&models.Book{}) // 3. Se Book é o nome do seu struct, use Book.Book{}
	if err != nil {
		log.Fatalf("Falha na migração automática: %v", err)
	}

	log.Println("Migração automática concluída com sucesso.")
}