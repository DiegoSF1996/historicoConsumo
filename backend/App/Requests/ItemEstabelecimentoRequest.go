package Requests

type ItemEstabelecimentoRequest struct {
	ItemId            int  `json:"item_id" binding:"required"`
	ItemReferenciaId  *int `json:"item_referencia_id"`
	EstabelecimentoId *int `json:"estabelecimento_id"`
}
