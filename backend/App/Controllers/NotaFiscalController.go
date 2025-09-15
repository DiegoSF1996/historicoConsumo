package Controllers

import (
	"historico_consumo/App/Repositories"
	"historico_consumo/App/Requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotaFiscalController struct {
	nota_fiscal_repository *Repositories.NotaFiscalRepository
}

func NewNotaFiscalController() *NotaFiscalController {
	repository := Repositories.NewNotaFiscalRepository()
	return &NotaFiscalController{
		nota_fiscal_repository: repository,
	}

	//nota_fiscal_repository: Repositories.NewNotaFiscalRepository(&Models.NotaFiscal{}),

}

func (controller *NotaFiscalController) Index(c *gin.Context) {

	nota_fiscal, err := controller.nota_fiscal_repository.GetAll()
	if err != nil {
		panic(err)
	}
	/* db := config.NewDBConfig()
	var nota_fiscal Models.NotaFiscal
	db.GetConnection().Preload("Estabelecimento").Preload("Consumidor").Find(&nota_fiscal) */
	c.JSON(200, gin.H{
		"message": "success",
		"data":    nota_fiscal,
	})
}

func (controller *NotaFiscalController) Create(c *gin.Context) {
	var body Requests.NotaFiscalStoreRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido", "detalhes": err.Error()})
		return
	}
	controller.nota_fiscal_repository.Create(body)
	// Agora vamos exibir formatado
	/* jsonBytes, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar JSON:", err)
	} else {
		fmt.Printf("📦 Dados recebidos do front:\n%s\n", string(jsonBytes))
	} */

}

func (controller *NotaFiscalController) Show() {
}

func (controller *NotaFiscalController) Update() {
}

func (controller *NotaFiscalController) Delete() {
}
