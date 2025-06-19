package Requests

type EstabelecimentoRequest struct {
	Descricao *string `json:"descricao"`
	Endereco  *string `json:"endereco"`
	Cnpj      *string `json:"cnpj"`
}
