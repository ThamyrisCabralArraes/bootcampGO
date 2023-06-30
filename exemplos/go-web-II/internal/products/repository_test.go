package products

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetAll(t *testing.T) {
	input := []Product{
		{
			ID:         1,
			Nome:       "CellPhone",
			Tipo:       "Tech",
			Quantidade: 3,
			Preco:      250,
		}, {
			ID:         2,
			Nome:       "Notebook",
			Tipo:       "Tech",
			Quantidade: 10,
			Preco:      1750.5,
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := Mock{
		Data: dataJson,
	}
	storeStub := FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	fmt.Println(storeStub)
	myRepo := NewRepository()
	resp, _ := myRepo.GetAll()
	assert.Equal(t, input, resp)
}
