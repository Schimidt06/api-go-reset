package models

type Personalidade struct {
    Nome      string `json:"nome"`
    Descricao string `json:"descricao"`
}

var Personalidades []Personalidade
