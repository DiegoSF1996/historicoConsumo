package Requests

import "time"

type NotaFiscalRequest struct {
	DataEmissao          *time.Time `json:"data_emissao"`
	QtdTotalItens        *int       `json:"qtd_total_itens"`
	ValorAPagar          *float64   `json:"valor_a_pagar"`
	ValorTributos        *float64   `json:"valor_tributos"`
	Numero               *string    `json:"numero"`
	Serie                *string    `json:"serie"`
	Emissao              *string    `json:"emissao"`
	ChaveDeAcesso        *string    `json:"chave_de_acesso"`
	ProtocoloAutorizacao *string    `json:"protocolo_autorizacao"`
	EstabelecimentoId    int        `json:"estabelecimento_id" binding:"required"`
	ConsumidorId         *int       `json:"consumidor_id"`
}
