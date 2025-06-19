package Requests

type CategoriaItemRequest struct {
	Descricao string `json:"descricao" binding:"required"`
}
