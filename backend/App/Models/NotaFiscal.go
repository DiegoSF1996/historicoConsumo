package Models

import (
	"time"
)

type NotaFiscal struct {
	ID                   int              `gorm:"primaryKey" json:"id"`
	DataEmissao          *time.Time       `gorm:"column:data_emissao;type:datetime" json:"data_emissao"`
	QtdTotalItens        *int             `gorm:"column:qtd_total_itens;type:int" json:"qtd_total_itens"`
	ValorAPagar          *float64         `gorm:"column:valor_a_pagar;type:float" json:"valor_a_pagar"`
	ValorTributos        *float64         `gorm:"column:valor_tributos;type:float" json:"valor_tributos"`
	Numero               *string          `gorm:"column:numero;type:varchar(50)" json:"numero"`
	Serie                *string          `gorm:"column:serie;type:varchar(50)" json:"serie"`
	Emissao              *string          `gorm:"column:emissao;type:varchar(50)" json:"emissao"`
	ChaveDeAcesso        *string          `gorm:"column:chave_de_acesso;type:varchar(255)" json:"chave_de_acesso"`
	ProtocoloAutorizacao *string          `gorm:"column:protocolo_autorizacao;type:varchar(255)" json:"protocolo_autorizacao"`
	EstabelecimentoId    int              `gorm:"column:estabelecimento_id;type:int" json:"estabelecimento_id"`
	Estabelecimento      Estabelecimento  `gorm:"hasOne:estabelecimento" json:"estabelecimento"`
	ConsumidorID         *int             `gorm:"column:consumidor_id;type:int" json:"consumidor_id"`
	Consumidor           Consumidor       `gorm:"hasOne:consumidor" json:"consumidor"`
	ItemNotaFiscal       []ItemNotaFiscal `gorm:"hasMany:item_nota_fiscal" json:"item_nota_fiscal"`
	CreatedAt            time.Time        `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt            time.Time        `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (NotaFiscal) TableName() string {
	return "nota_fiscal"
}
