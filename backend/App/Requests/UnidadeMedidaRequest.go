package Requests

type UnidadeMedidaRequest struct {
	Descricao string `json:"descricao" binding:"required"`
}
