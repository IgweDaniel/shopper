package internal

import (
	"github.com/IgweDaniel/shopper/internal/config"
	"github.com/IgweDaniel/shopper/internal/contracts"
)

type Application struct {
	Config       *config.Config
	Repositories contracts.Repositories
}
