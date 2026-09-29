package handlers

import (
    "fmt"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Bem-vindo à página de cadastro de alunos em Go! 🚀")
}
