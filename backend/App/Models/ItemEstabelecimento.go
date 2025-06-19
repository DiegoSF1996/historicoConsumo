package Models

import (
	"time"
)

type ItemEstabelecimento struct {
	Id                int       `gorm:"column:id"`
	ItemId            int       `gorm:"column:item_id;type:int" json:"item_id"`
	ItemReferenciaId  int       `gorm:"column:item_referencia_id;type:int" json:"item_referencia_id"`
	EstabelecimentoId int       `gorm:"column:estabelecimento_id;type:int" json:"estabelecimento_id"`
	CreatedAt         time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (ItemEstabelecimento) TableName() string {
	return "item_estabelecimento"
}
