package reader

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewUserBrewRecipeReader(db *ent.Client) *UserBrewRecipeReader {
	return &UserBrewRecipeReader{
		db: db,
	}
}

//go:generate mockgen -source=./user_brew_recipe.go -destination=./mock/user_brew_recipe_mock.go -package=reader
type UserBrewRecipe interface {
	GetByID(ctx context.Context, userBrewRecipeID int) (model.UserBrewRecipe, error)
}

type UserBrewRecipeReader struct {
	db *ent.Client
}

var _ UserBrewRecipe = (*UserBrewRecipeReader)(nil)

func (impl *UserBrewRecipeReader) GetByID(ctx context.Context, userBrewRecipeID int) (model.UserBrewRecipe, error) {
	return model.UserBrewRecipe{}, nil
}
