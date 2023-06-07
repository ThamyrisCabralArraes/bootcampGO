package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thamyriscabralarraes/bootcampGO/exemplos/go-web-II/cmd/server/handler"
	"github.com/thamyriscabralarraes/bootcampGO/exemplos/go-web-II/internal/products"
)

type request struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

var produtos []request
var lastID int

func Salvar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token invalido",
			})
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastID++
		req.ID = lastID

		produtos = append(produtos, req)
		ctx.JSON(200, req)
	}
}

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	r.Run()
}
