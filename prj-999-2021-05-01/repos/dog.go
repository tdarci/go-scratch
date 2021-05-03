package repos

import (
	"context"

	"github.com/tdarci/prj-999/config"
	"github.com/tdarci/prj-999/models"
)

type Dog struct {
	*config.Config
}

func NewDog(cfg *config.Config) *Dog {
	return &Dog{
		Config: cfg,
	}
}

func (d *Dog) Get(ctx context.Context, id int) (*models.Dog, error) {

	dog := &models.Dog{}

	err := d.DB().QueryRowContext(ctx, "select id, name from dog where id = ?", id).Scan(&dog.ID, &dog.Name)

	if err != nil {
		return nil, err
	}
	return dog, nil
}
