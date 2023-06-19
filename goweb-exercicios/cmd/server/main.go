package main

import (
	"GOWEB-EXERCICIOS/cmd/server/handler"
	"GOWEB-EXERCICIOS/docs"
	"GOWEB-EXERCICIOS/internal/repository"
	"GOWEB-EXERCICIOS/internal/service"
	"GOWEB-EXERCICIOS/pkg/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
Instanciamos cada camada do domínio Products e usaremos os métodos do controlador para cada endpoint.
*/

// @title MELI Bootcamp API - Produtos
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error ao carregar o arquivo .env")
	}

	store := store.Factory("arquivo", "produtos.json")
	if store == nil {
		log.Fatal("Não foi possivel criar a store")
	}

	repo := repository.NewRepository(store) // Criação da instância Repository
	service := service.NewService(repo)     // Criação da instância Service
	p := handler.NewProduct(service)        // Criação do Controller

	r := gin.Default()
	pr := r.Group("/produtos")
	{
		pr.Use(TokenAuthMiddleware())

		pr.POST("/post", p.Salvar())
		pr.GET("/getAll", p.GetAll())
		pr.PUT("/:id", p.Update())
		pr.DELETE("/:id", p.Delete())
		pr.PATCH("/updateNome/:id", p.UpdateNome())
		pr.PATCH("/updatePreco/:id", p.UpdatePreco())
	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}

//Func main com outros exercícios sem estrutura
// func main() {
// 	r := gin.Default()

// 	pr := r.Group("/produtos")
// 	{
// 		pr.POST("/salvar", handler.)
// 		pr.GET("/getAll", handler.GetAll())
// 		// pr.GET("/get", controller.Buscar())
// 		// pr.GET("/getAll", controller.GetAllWithFilters)
// 		// pr.GET("/getProductById/:id", controller.GetProductById())
// 	}

// 	r.Run()

// }
