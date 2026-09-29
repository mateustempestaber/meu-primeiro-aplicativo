package main

import (
    "fmt"
    "net/http"
    "cadastro-alunos/internal/handlers"
)

func main() {
    // Diz ao servidor qual função chamar quando acessarem a raiz "/"
    http.HandleFunc("/", handlers.HomeHandler)

    fmt.Println("Servidor Go rodando na porta 8080...")
    
    // Inicia o servidor web na porta 8080
    http.ListenAndServe(":8080", nil)
}
