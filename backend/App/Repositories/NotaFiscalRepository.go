package Repositories

import (
	"fmt"
	"historico_consumo/App/Models"
	"historico_consumo/App/Requests"
	"historico_consumo/config"
	"strings"

	"github.com/jinzhu/copier"
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
func (repository *NotaFiscalRepository) Create(body Requests.NotaFiscalStoreRequest) {

	//obter ou criar estabelecimento
	var NotaFiscal Models.NotaFiscal
	estabelecimento := repository.insertOrCreateEstabelecimento(body.Estabelecimento)
	NotaFiscal.EstabelecimentoId = estabelecimento.Id
	//obter ou criar consumidor
	consumidor := repository.insertOrCreateConsumidor(body.Consumidor)
	NotaFiscal.ConsumidorId = consumidor.Id

	//criar nota fiscal
	nota_fiscal_model := repository.insertNotaFiscal(body, estabelecimento.Id, consumidor.Id)

	//obter ou criar forma de pagamento

	repository.insertOrCreateFormaPagamentoNotaFiscal(body.FormasPagamento, nota_fiscal_model.Id)

	//criar itens
	repository.insertItemNotaFiscal(nota_fiscal_model.Id, body.ItemNotaFiscal)
}

func (repository *NotaFiscalRepository) insertOrCreateEstabelecimento(estabelecimento_request Requests.EstabelecimentoStoreRequest) Models.Estabelecimento {
	var estabelecimento Models.Estabelecimento

	repository.db.GetConnection().
		Where("cnpj = ?", estabelecimento_request.Cnpj).
		Attrs(Models.Estabelecimento{
			Descricao: estabelecimento_request.Descricao,
			Endereco:  estabelecimento_request.Endereco,
			Cnpj:      estabelecimento_request.Cnpj,
		}).
		FirstOrCreate(&estabelecimento)
	return estabelecimento
	/* repository.db.GetConnection().
		Where("cnpj = ?", estabelecimento_request.Cnpj).
		First(&estabelecimento)

	if estabelecimento.Id == 0 {
		estabelecimento = Models.Estabelecimento{
			Cnpj:      estabelecimento_request.Cnpj,
			Descricao: estabelecimento_request.Descricao,
			Endereco:  estabelecimento_request.Endereco,
		}
		repository.db.GetConnection().Create(&estabelecimento)
	} */
	/* jsonBytes, err := json.MarshalIndent(estabelecimento_request, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar JSON:", err)
	} else {
		fmt.Printf("📦 Dados recebidos do front:\n%s\n", string(jsonBytes))
	} */
}
func (repository *NotaFiscalRepository) insertOrCreateConsumidor(consumidor_request Requests.ConsumidorStoreRequest) Models.Consumidor {
	var consumidor Models.Consumidor

	repository.db.GetConnection().
		FirstOrCreate(&consumidor, Models.Consumidor{
			Cpf: consumidor_request.Cpf,
		})
	return consumidor
}

func (repository *NotaFiscalRepository) insertNotaFiscal(
	nota_fiscal_request Requests.NotaFiscalStoreRequest,
	estabelecimento_id int,
	consumidor_id int,
) Models.NotaFiscal {
	var nota_fiscal_model Models.NotaFiscal
	var nota_fiscal_dados Models.NotaFiscal

	copier.Copy(&nota_fiscal_dados, &nota_fiscal_request)
	nota_fiscal_dados.EstabelecimentoId = estabelecimento_id
	nota_fiscal_dados.ConsumidorId = consumidor_id
	nota_fiscal_dados.ItemNotaFiscal = []Models.ItemNotaFiscal{}
	repository.db.GetConnection().
		Where("numero = ? AND serie = ?", nota_fiscal_request.Numero, nota_fiscal_request.Serie).
		Attrs(nota_fiscal_dados).
		FirstOrCreate(&nota_fiscal_model)
	return nota_fiscal_model
}

func (repository *NotaFiscalRepository) insertOrCreateFormaPagamentoNotaFiscal(
	formas_pagamento_nota_fiscalrequest []Requests.FormasPagamentoNotaFiscalStoreRequest,
	nota_fiscal_id int,
) []Models.FormaPagamentoNotaFiscal {
	var forma_pagamento_nota_fiscal_model []Models.FormaPagamentoNotaFiscal

	for _, forma_pagamento_nota_fiscal := range formas_pagamento_nota_fiscalrequest {
		var fp Models.FormaPagamentoNotaFiscal
		var fp_dados Models.FormaPagamentoNotaFiscal
		tipo_pagamento_model := repository.insertOrCreateTipoPagamento(forma_pagamento_nota_fiscal.Descricao)

		copier.Copy(&fp_dados, &forma_pagamento_nota_fiscal)
		fp_dados.NotaFiscalId = nota_fiscal_id
		fp_dados.TipoPagamentoId = tipo_pagamento_model.Id

		repository.db.GetConnection().
			Where("tipo_pagamento_id = ?", tipo_pagamento_model.Id).
			Where("nota_fiscal_id = ?", nota_fiscal_id).
			Attrs(fp_dados).
			FirstOrCreate(&fp)

		forma_pagamento_nota_fiscal_model = append(forma_pagamento_nota_fiscal_model, fp)
	}

	return forma_pagamento_nota_fiscal_model
}

func (repository *NotaFiscalRepository) insertOrCreateTipoPagamento(descricao_tipo_pagamento string) Models.TipoPagamento {
	var tipo_pagamento Models.TipoPagamento

	repository.db.GetConnection().
		FirstOrCreate(&tipo_pagamento, Models.TipoPagamento{
			Descricao: descricao_tipo_pagamento,
		})
	return tipo_pagamento
}

func (repository *NotaFiscalRepository) insertItemNotaFiscal(
	nota_fiscal_id int,
	item_nota_fiscal_request []Requests.ItemNotaFiscalStoreRequest,
) {

	for _, item := range item_nota_fiscal_request {
		var item_nota_fiscal_model = Models.ItemNotaFiscal{
			Quantidade:    item.Quantidade,
			PrecoUnitario: item.PrecoUnitario,
			ValorTotal:    item.ValorTotal,
		}
		fmt.Println("========================================================")
		fmt.Printf("inf encontrada/inserida: %+v\n", item_nota_fiscal_model)

		item := repository.insertOrCreateItem(
			item.Item.Codigo,
			item.Item.Descricao,
			item.Item.UnidadeMedida.Descricao,
		)

		item_nota_fiscal_model.NotaFiscalId = nota_fiscal_id
		item_nota_fiscal_model.ItemId = item.Id
		fmt.Printf("---item encontrada/inserida: %+v\n", item_nota_fiscal_model)
		repository.db.GetConnection().Debug().
			Omit("Item").
			Create(&item_nota_fiscal_model)
	}
}

func (repository *NotaFiscalRepository) insertOrCreateItem(
	codigo string,
	descricao string,
	unidade_medida_descricao string,
) Models.Item {
	//var item Models.Item
	//fmt.Printf(codigo, descricao, unidade_medida_descricao)
	unidade_medida := repository.insertOrCreateUnidadeMedida(unidade_medida_descricao)
	item_insert := Models.Item{
		Codigo:          codigo,
		Descricao:       descricao,
		UnidadeMedidaId: unidade_medida.Id,
	}
	fmt.Printf("unidadeMedidaID encontrada/inserida: %+v\n", item_insert.UnidadeMedidaId)
	fmt.Printf("item encontrada/inserida: %+v\n", item_insert)

	repository.db.GetConnection().
		Omit("UnidadeMedida").
		Select("codigo", "descricao", "unidade_medida_id").
		Create(&item_insert)
	return item_insert
}

func (repository *NotaFiscalRepository) insertOrCreateUnidadeMedida(
	descricao string,
) Models.UnidadeMedida {
	descricao = strings.ToLower(descricao)
	unidade_medida := Models.UnidadeMedida{}
	repository.db.GetConnection().Debug().
		Where("descricao = ?", descricao).
		Attrs(Models.UnidadeMedida{Descricao: descricao}).
		FirstOrCreate(&unidade_medida)

	fmt.Printf("unidade encontrada/inserida: %+v\n", unidade_medida)

	/* Debug().
	FirstOrCreate(&unidade_medida) */
	//fmt.Printf("Unidade encontrada/inserida: %+v\n", unidade_medida)

	return unidade_medida
}
