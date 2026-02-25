package service

import (
	"github.com/google/uuid"
	"github.com/oddball707/recipes/dao"
	m "github.com/oddball707/recipes/model"
)

type RecipeService struct {
	recipeDAO dao.RecipeDAO
}

type RecipeClient interface {
	GetRecipe(id uuid.UUID) (*m.Recipe, error)
	CreateRecipe(recipe *m.Recipe) (*m.Recipe, error)
	UpdateRecipe(recipe *m.Recipe) (*m.Recipe, error)
	DeleteRecipe(id uuid.UUID) error
}

func NewRecipeService(recipeDAO dao.RecipeDAO) RecipeClient {
	return &RecipeService{
		recipeDAO: recipeDAO,
	}
}

func (s *RecipeService) GetRecipe(id uuid.UUID) (*m.Recipe, error) {
	return s.recipeDAO.GetRecipe(id)
}

func (s *RecipeService) CreateRecipe(recipe *m.Recipe) (*m.Recipe, error) {
	err := s.recipeDAO.CreateRecipe(recipe)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (s *RecipeService) UpdateRecipe(recipe *m.Recipe) (*m.Recipe, error) {
	err := s.recipeDAO.UpdateRecipe(recipe)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (s *RecipeService) DeleteRecipe(id uuid.UUID) error {
	return s.recipeDAO.DeleteRecipe(id)
}
