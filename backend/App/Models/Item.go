package Models

import "time"

type Item struct {
	ID              int           `gorm:"primaryKey" json:"id"`
	Codigo          string        `gorm:"column:codigo;type:varchar(60)" json:"codigo"`
	Descricao       string        `gorm:"column:descricao;type:varchar(90)" json:"descricao"`
	UnidadeMedidaId int           `gorm:"column:unidade_medida_id;type:int" json:"unidade_medida_id"`
	UnidadeMedida   UnidadeMedida `gorm:"belongsTo:unidade_medida"`
	CategoriaItemId int           `gorm:"column:categoria_item_id;type:int" json:"categoria_item_id"`
	CategoriaItem   CategoriaItem `gorm:"belongsTo:categoria_item"`
	CreatedAt       time.Time     `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (Item) TableName() string {
	return "item"
}
