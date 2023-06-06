package main

import "github.com/gin-gonic/gin"

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
	r := gin.Default()
	pr := r.Group("/produtos")
	pr.POST("/salvar", Salvar())
	r.Run()
}
