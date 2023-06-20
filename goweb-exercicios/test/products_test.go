package test_test

import (
	"GOWEB-EXERCICIOS/cmd/server/handler"
	"GOWEB-EXERCICIOS/internal/repository"
	"GOWEB-EXERCICIOS/internal/service"
	"GOWEB-EXERCICIOS/pkg/store"
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// iniciar o servidor e definir rotas para teste
func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.Factory(store.FileType, "produtos_test.json")
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/produtos")
	pr.PUT("/:id", p.Update())
	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")
	return req, httptest.NewRecorder()
}

func Test_ProductsUpdate(t *testing.T) {
	r := createServer()
	req, response := createRequestTest(http.MethodPut, "/produtos/1", `{"nome": "Aspirador", "cor": "Branco", "preco": 4.99, "estoque": 4, "codigo": "AF333", "publicacao": true, "dataCriacao": "20231010"}`)

	r.ServeHTTP(response, req)

	// printar a rota
	assert.Equal(t, 200, response.Code, response.Body.String())

	if response.Code != 200 {
		t.Errorf("Falha na solicitação: %d - %s", response.Code, response.Body.String())
	}
}
