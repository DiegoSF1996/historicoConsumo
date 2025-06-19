package Requests

type ItemNotaFiscalRequest struct {
	NotaFiscalId  int     `json:"nota_fiscal_id" binding:"required"`
	ItemId        int     `json:"item_id" binding:"required"`
	Quantidade    float64 `json:"quantidade" binding:"required"`
	PrecoUnitario float64 `json:"preco_unitario" binding:"required"`
	ValorTotal    float64 `json:"valor_total" binding:"required"`
}
