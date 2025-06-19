package Models

import "time"

type Estabelecimento struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Descricao *string   `gorm:"column:descricao;type:varchar(50)" json:"descricao"`
	Endereco  *string   `gorm:"column:endereco;type:varchar(50)" json:"endereco"`
	Cnpj      *string   `gorm:"column:cnpj;type:varchar(50)" json:"cnpj"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (Estabelecimento) TableName() string {
	return "estabelecimento"
}
