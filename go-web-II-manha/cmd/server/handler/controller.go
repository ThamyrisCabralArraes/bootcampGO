package handler

import (
	"GO-WEB-II-MANHA/internal/db"
	"os"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Senha int    `json:"idade"`
}

var login []request

type Logins struct {
	service db.Servide
}

func NewLogin(db db.Servide) *Logins {
	return &Logins{
		service: db,
	}
}

func (c *Logins) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token invalido"})
			return
		}
		l, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, l)
	}

}

func (c *Logins) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token invalido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": "erro na request"})
			return
		}

		if req.Nome == "" {
			ctx.JSON(400, gin.H{"error": "nome é obrigatorio"})
			return
		}

		if req.Senha == 0 {
			ctx.JSON(400, gin.H{"error": "idade é obrigatoria"})
			return
		}

		l, err := c.service.Store(req.Nome, req.Senha)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// lastID := len(login)
		// if lastID == 0 {
		// 	req.ID = 1
		// } else {
		// 	req.ID = lastID + 1
		// }

		// login = append(login, req)
		ctx.JSON(200, l)
	}
}
