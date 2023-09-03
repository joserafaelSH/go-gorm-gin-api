package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Name string `json:"name"`
	Cpf string `json:"cpf"`
	Rg string `json:"rg"`
}

var Alunos []Aluno 