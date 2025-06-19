package Repositories

import (
	"fmt"
	"historico_consumo/App/Models"
	"historico_consumo/App/Requests"
	"historico_consumo/config"

	"github.com/gin-gonic/gin"
)

type NotaFiscalRepository struct {
	db *config.DBConfig
}

func NewNotaFiscalRepository() *NotaFiscalRepository {
	db := config.NewDBConfig()
	db.GetConnection()

	return &NotaFiscalRepository{
		db: db,
	}
}
func (repository *NotaFiscalRepository) GetAll() (*[]Models.NotaFiscal, error) {
	var nota_fiscal *[]Models.NotaFiscal
	repository.db.GetConnection().Preload("Estabelecimento").Preload("Consumidor").Preload("ItemNotaFiscal").Find(&nota_fiscal)
	return nota_fiscal, nil
}
func (repository *NotaFiscalRepository) Create(c *gin.Context) {

	//obter ou criar estabelecimento

	//repository.insertOrCreateEstabelecimento()
	//obter ou criar consumidor
	repository.insertOrCreateConsumidor()
	//obter ou criar forma de pagamento
	repository.insertOrCreateFormaPagamento()
	//criar nota fiscal
	repository.insertNotaFiscal()
	//criar itens
	repository.insertItem()
}

func (repository *NotaFiscalRepository) insertOrCreateEstabelecimento(estabelecimento_request *Requests.EstabelecimentoRequest) {
	//var estabelecimento *Models.Estabelecimento
	fmt.Print(estabelecimento_request)
	//repository.db.GetConnection().FirstOrCreate(&estabelecimento, Models.Estabelecimento{Cnpj: "Teste"})
}
func (repository *NotaFiscalRepository) insertOrCreateConsumidor() {

}

func (repository *NotaFiscalRepository) insertOrCreateFormaPagamento() {

}

func (repository *NotaFiscalRepository) insertNotaFiscal() {

}

func (repository *NotaFiscalRepository) insertItem() {

}
