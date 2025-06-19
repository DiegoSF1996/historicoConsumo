package Requests

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const layout = "02/01/2006 15:04:05-07:00"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type NotaFiscalStoreRequest struct {
	Serie                *string  `json:"serie"`
	Numero               *string  `json:"numero"`
	Emissao              *string  `json:"emissao"`
	ProtocoloAutorizacao *string  `json:"protocolo_autorizacao"`
	DataAutorizacao      *string  `json:"data_autorizacao"`
	ChaveDeAcesso        *string  `json:"chave_de_acesso"`
	QtdTotalItens        *int     `json:"qtd_total_itens"`
	ValorTotal           *float64 `json:"valor_total"`
	Desconto             *float64 `json:"desconto"`
	ValorAPagar          *float64 `json:"valor_a_pagar"`
	ValorTributos        *float64 `json:"valor_tributos"`
	EstabelecimentoId    int      `json:"estabelecimento_id"`
}
