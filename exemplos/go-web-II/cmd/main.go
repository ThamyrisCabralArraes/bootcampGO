package main

import (
	"GO-WEB-II/cmd/server/handler"
	"GO-WEB-II/internal/products"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// 1 maneira de fazer o POST
// type request struct {
// 	ID         int     `json:"id"`
// 	Nome       string  `json:"nome"`
// 	Tipo       string  `json:"tipo"`
// 	Quantidade int     `json:"quantidade"`
// 	Preco      float64 `json:"preco"`
// }

// var produtos []request
// var lastID int

// func Salvar() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		token := ctx.GetHeader("token")
// 		if token != "123456" {
// 			ctx.JSON(401, gin.H{
// 				"error": "token invalido",
// 			})
// 		}
// 		var req request
// 		if err := ctx.ShouldBindJSON(&req); err != nil {
// 			ctx.JSON(400, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		lastID++
// 		req.ID = lastID

// 		produtos = append(produtos, req)
// 		ctx.JSON(200, req)
// 	}
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error ao carregar o arquivo env")
	}
	usuario := os.Getenv("MY_USER")
	password := os.Getenv("MY_PASS")

	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}
