package Models

import (
	"time"
)

type FormaPagamento struct {
	Id              int       `gorm:"column:id" json:"id"`
	Codigo          string    `gorm:"column:codigo;type:varchar(60)" json:"codigo"`
	Descricao       string    `gorm:"column:descricao;type:varchar(90)" json:"descricao"`
	TipoPagamentoId int       `gorm:"column:tipo_pagamento_id;type:int" json:"tipo_pagamento_id"`
	NotaFiscalId    int       `gorm:"column:nota_fiscal_id;type:int" json:"nota_fiscal_id"`
	ValorPago       float64   `gorm:"column:valor_pago;type:float" json:"valor_pago"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (FormaPagamento) TableName() string {
	return "forma_pagamento"
}
