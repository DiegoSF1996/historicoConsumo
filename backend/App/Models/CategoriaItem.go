package Models

import "time"

type CategoriaItem struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Descricao string    `gorm:"column:descricao;type:varchar(90)" json:"descricao"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (CategoriaItem) TableName() string {
	return "categoria_item"
}
