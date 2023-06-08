package main

import (
	"GO-WEB-II-MANHA/cmd/server/handler"
	"GO-WEB-II-MANHA/internal/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("erro ao carregar o arquivo env")
	}

	repo := db.NewRepository()
	service := db.NewService(repo)
	l := handler.NewLogin(service)

	r := gin.Default()
	pr := r.Group("/login")

	pr.POST("/salvar", l.Store())
	pr.GET("/", l.GetAll())
	r.Run()
}
