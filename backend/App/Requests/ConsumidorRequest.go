package Requests

type ConsumidorRequest struct {
	Descricao *string `json:"descricao"`
	Cpf       *string `json:"cpf"`
}
