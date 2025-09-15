package Models

import "time"

type ItemNotaFiscal struct {
	ID            int        `gorm:"primaryKey" json:"id"`
	NotaFiscalId  int        `gorm:"column:nota_fiscal_id" json:"nota_fiscal_id"`
	NotaFiscal    NotaFiscal `gorm:"belongsTo:nota_fiscal json:nota_fiscal"`
	ItemId        int        `gorm:"column:item_id;type:int" json:"item_id"`
	Item          Item       `gorm:"belongsTo:item"`
	Quantidade    float64    `gorm:"column:quantidade;type:float" json:"quantidade"`
	PrecoUnitario float64    `gorm:"column:preco_unitario;type:float" json:"preco_unitario"`
	ValorTotal    float64    `gorm:"column:valor_total;type:float" json:"valor_total"`
	CreatedAt     time.Time  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (ItemNotaFiscal) TableName() string {
	return "item_nota_fiscal"
}
