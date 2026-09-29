package models

type Aluno struct {
    ID    int    `json:"id"`
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
    Email string `json:"email"`
}
