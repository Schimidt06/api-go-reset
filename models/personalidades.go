package models

type Personalidade struct {
	Id        int64  `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}

var Personalidades []Personalidade
