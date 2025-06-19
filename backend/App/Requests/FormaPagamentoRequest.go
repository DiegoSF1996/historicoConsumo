package Requests

type FormaPagamentoRequest struct {
	Codigo          *string `json:"codigo"`
	Descricao       string  `json:"descricao" binding:"required"`
	TipoPagamentoId int     `json:"tipo_pagamento_id" binding:"required"`
	NotaFiscalId    int     `json:"nota_fiscal_id" binding:"required"`
	ValorPago       float64 `json:"valor_pago" binding:"required"`
}
