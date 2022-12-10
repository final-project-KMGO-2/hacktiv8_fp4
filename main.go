package main

import (
	"context"
	"hacktiv8_fp_2/config"
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/middleware"
	"hacktiv8_fp_2/repository"
	"hacktiv8_fp_2/routes"
	"hacktiv8_fp_2/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	db := config.SetupDatabaseConnection()
	userRepository := repository.NewUserRepository(db)
	transactionHistoryRepository := repository.NewTransactionHistoryRepository(db)
	productRepository := repository.NewProductRepo(db)
	categoryRepository := repository.NewCategoryRepository(db)

	jwtService := middleware.NewJWTService()
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)
	transactionHistoryService := service.NewTransactionHistoryService(transactionHistoryRepository)
	productService := service.NewProductService(productRepository)
	categoryService := service.NewCategoryService(categoryRepository)

	authController := controller.NewUserController(userService, authService, jwtService)
	transactionHistoryController := controller.NewTransactionHistoryController(transactionHistoryService, jwtService)
	productController := controller.NewProductController(jwtService, productService)
	categoryController := controller.NewCategoryController(categoryService, jwtService)

	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	routes.UserRoutes(server, authController, jwtService)
	routes.TransactionHistoryRoutes(server, transactionHistoryController, jwtService)
	routes.GenerateProductRoutes(server, productController, jwtService)
	routes.CategoryRoutes(server, categoryController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("error serving :", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("[%v] - Shutting down server\n", time.Now())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("error shutting down :", err)
	}

	<-ctx.Done()
	log.Println("timeout, exiting")

}
