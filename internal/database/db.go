package database

import (
    "database/sql"
    "fmt"
)

func ConectaDB() *sql.DB {
    // Aqui no futuro colocaremos a conexão real com o PostgreSQL
    fmt.Println("Conexão com o banco preparada!")
    return nil
}
