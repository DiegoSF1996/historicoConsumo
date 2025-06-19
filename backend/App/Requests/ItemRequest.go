package Requests

type ItemRequest struct {
	Codigo          *string `json:"codigo"`
	Descricao       string  `json:"descricao" binding:"required"`
	UnidadeMedidaId int     `json:"unidade_medida_id" binding:"required"`
	CategoriaItemId *int    `json:"categoria_item_id"`
}
