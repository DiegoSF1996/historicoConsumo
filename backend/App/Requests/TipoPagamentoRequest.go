package Requests

type TipoPagamentoRequest struct {
	Descricao string `json:"descricao" binding:"required"`
}
