package Models

import "time"

type UnidadeMedida struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Descricao string    `gorm:"column:descricao;type:varchar(90)" json:"descricao"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (UnidadeMedida) TableName() string {
	return "unidade_medida"
}
