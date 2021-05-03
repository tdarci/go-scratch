package engine

import (
	"context"

	"github.com/tdarci/prj-999/config"
	"github.com/tdarci/prj-999/models"
	"github.com/tdarci/prj-999/repos"
)

type Engine struct {
	*config.Config
	dogRepo *repos.Dog
}

func NewEngine(cfg *config.Config) *Engine {
	return &Engine{
		Config:  cfg,
		dogRepo: repos.NewDog(cfg),
	}
}

func (e *Engine) Add(a, b int) int {
	return a + b
}

func (e *Engine) GetDog(ctx context.Context, id int) (*models.Dog, error) {
	return e.dogRepo.Get(ctx, id)
}
