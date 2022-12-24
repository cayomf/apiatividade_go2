package models

type Entrega struct {
	Id int `json:"id"`
	AlunoId int `json:"aluno_id"`
	AtividadeId int `json:"atividade_id"`
	Nota string `json:"nota"`
	Disciplina string `json:"disciplina"`
}