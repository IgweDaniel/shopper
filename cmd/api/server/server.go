package server

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/IgweDaniel/shopper/cmd/api/routes"
	"github.com/IgweDaniel/shopper/internal"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/database"
	"github.com/IgweDaniel/shopper/internal/repository"
	"github.com/IgweDaniel/shopper/internal/services"
)

func NewServer() *http.Server {

	cfg, err := internal.LoadConfig()
	db := database.New(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	repos := contracts.Repositories{
		User:    repository.NewPostgresUserRepository(db.DB()),
		Order:   repository.NewPostgresOrderRepository(db.DB()),
		Product: repository.NewPostgresProductRepository(db.DB()),
	}

	app := internal.Application{
		Config:       cfg,
		Repositories: repos,
	}
	services := contracts.Services{
		User:    services.NewUserService(&app),
		Order:   services.NewOrderService(&app),
		Product: services.NewProductService(&app),
	}
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      routes.RegisterRoutes(db, &services),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
