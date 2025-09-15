package Requests

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const inputLayout = "02/01/2006 15:04:05-07:00"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}
	t, err := time.Parse(inputLayout, s)
	if err != nil {
		return err
	}
	ct.Time = t.Local()
	return nil
}

type NotaFiscalStoreRequest struct {
	Serie                *string                                 `json:"serie"`
	Numero               *string                                 `json:"numero"`
	Emissao              *string                                 `json:"emissao"`
	ProtocoloAutorizacao *string                                 `json:"protocolo_autorizacao"`
	DataAutorizacao      string                                  `json:"data_autorizacao"`
	ChaveDeAcesso        *string                                 `json:"chave_de_acesso"`
	QtdTotalItens        *int                                    `json:"qtd_total_itens"`
	ValorTotal           *float64                                `json:"valor_total"`
	Desconto             *float64                                `json:"desconto"`
	ValorAPagar          *float64                                `json:"valor_a_pagar"`
	ValorTributos        *float64                                `json:"valor_tributos"`
	Estabelecimento      EstabelecimentoStoreRequest             `json:"estabelecimento"`
	Consumidor           ConsumidorStoreRequest                  `json:"consumidor"`
	FormasPagamento      []FormasPagamentoNotaFiscalStoreRequest `json:"formas_pagamento"`
	ItemNotaFiscal       []ItemNotaFiscalStoreRequest            `json:"item_nota_fiscal"`
}
type EstabelecimentoStoreRequest struct {
	Descricao *string `json:"descricao"`
	Endereco  *string `json:"endereco"`
	Cnpj      *string `json:"cnpj"`
}

type ConsumidorStoreRequest struct {
	Cpf *string `json:"cpf"`
}
type FormasPagamentoNotaFiscalStoreRequest struct {
	Descricao string  `json:"descricao"`
	ValorPago float64 `json:"valor_pago"`
}

type ItemNotaFiscalStoreRequest struct {
	Item          ItemStoreRequest `json:"item"`
	Quantidade    float64          `json:"quantidade"`
	PrecoUnitario float64          `json:"preco_unitario"`
	ValorTotal    float64          `json:"valor_total"`
}

type ItemStoreRequest struct {
	Codigo        string               `json:"codigo"`
	Descricao     string               `json:"descricao" binding:"required"`
	UnidadeMedida UnidadeMedidaRequest `json:"unidade" binding:"required"`
}

type UnidadeMedidaStoreRequest struct {
	Descricao string `json:"descricao" binding:"required"`
}
