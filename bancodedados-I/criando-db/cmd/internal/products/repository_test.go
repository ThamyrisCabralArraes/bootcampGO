package products_test

import (
	"criando-db/cmd/internal/products"
	"criando-db/cmd/internal/products/models"
	"criando-db/cmd/internal/products/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sqlRepository_GetOne(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := products.NewRepo(db)

	productEmpty := models.Product{}
	product := models.Product{
		Name:  "Mac",
		Type:  "eletronico",
		Count: 1,
		Price: 10,
	}
	product, err = repository.Store(product)
	assert.NoError(t, err)
	getResult, err := repository.GetOne(3333333)
	assert.NoError(t, err)
	assert.Equal(t, productEmpty, getResult)
	getResult, err = repository.GetOne(product.ID)
	assert.NoError(t, err)
	assert.NotNil(t, getResult)
	assert.Equal(t, product.Name, getResult.Name)

	defer db.Close()

}

func Test_sqlRepository_Update(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := products.NewRepo(db)

	product := models.Product{
		Name:  "Mac",
		Type:  "eletronico",
		Count: 1,
		Price: 10,
	}

	productUpdate := models.Product{
		Name:  "Mac",
		Type:  "eletronico",
		Count: 1,
		Price: 10000,
	}

	_, err = repository.Store(product)
	assert.NoError(t, err)
	getResult, err := repository.Update(productUpdate)
	assert.Equal(t, productUpdate, getResult)
	assert.NoError(t, err)

	defer db.Close()
}

func Test_sqlRepository_GetAll(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := products.NewRepo(db)

	expectedProduct := []models.Product{
		{ID: 4, Name: "Produto 1",
			Type:  "Tipo 1",
			Count: 10,
			Price: 10.5,
		},
		{Name: "Robo",
			Type:  "eletronico",
			Count: 1,
			Price: 5000,
		},
	}

	getResults, err := repository.GetAll()
	assert.Equal(t, expectedProduct[0], getResults[0])
	assert.NoError(t, err)

	defer db.Close()
}

func Test_sqlRepository_Delete(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := products.NewRepo(db)

	validateId := 1

	err = repository.Delete(validateId)
	assert.NoError(t, err)

	defer db.Close()

}
