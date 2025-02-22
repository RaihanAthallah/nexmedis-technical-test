package main

import (
	"fmt"
	"log"
	"net/http"
	"nexmedis-technical-test/app/config"
	"nexmedis-technical-test/app/handler"
	"nexmedis-technical-test/app/middleware"
	"nexmedis-technical-test/app/repository"
	"nexmedis-technical-test/app/routes"
	"nexmedis-technical-test/app/usecase"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Membuat router Gin
	db := config.ConnectDB()
	router := gin.Default()
	// Configure CORS
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowHeaders:     []string{"Content-Type", "Authorization", "token"}, // Add the "token" header here
				AllowCredentials: true,
			},
		),
	)

	//dependency injection
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	cartRepository := repository.NewCartRepository(db)
	bankRepository := repository.NewBankRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	productUsecase := usecase.NewProductUsecase(productRepository)
	cartUsecase := usecase.NewCartUsecase(cartRepository)
	bankUsecase := usecase.NewBankUsecase(bankRepository)

	userHandler := handler.NewAuthHandler(userUsecase)
	productHandler := handler.NewProductHandler(productUsecase)
	cartHandler := handler.NewCartHandler(cartUsecase)
	bankHandler := handler.NewBankHandler(bankUsecase)

	routes.SetupUserRoutes(router.Group("/api/v1"), userHandler)

	authGroup := router.Group("/api/v1/")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
		})
		routes.SetupProductRoutes(authGroup, productHandler)
		routes.SetupCartRoutes(authGroup, cartHandler)
		routes.SetupBankRoutes(authGroup, bankHandler)
	}

	port := os.Getenv("APP_PORT")
	// Menjalankan server
	fmt.Printf("Server is running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
