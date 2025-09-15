package Models

import "time"

type Consumidor struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Descricao *string   `gorm:"column:descricao;type:varchar(150)" json:"descricao"`
	Cpf       *string   `gorm:"column:cpf;type:varchar(50)" json:"cpf"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (Consumidor) TableName() string {
	return "consumidor"
}
