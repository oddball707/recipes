package service

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	d "github.com/oddball707/recipes/dao"
	m "github.com/oddball707/recipes/model"
	"github.com/stretchr/testify/assert"
)

func TestGetRecipe(t *testing.T) {
	id := uuid.New()
	getRecipeSuccess := &m.Recipe{
		ID:   id,
		Name: "Pasta",
	}

	testCases := []struct {
		name           string
		id             uuid.UUID
		setupMock      func(*d.MockRecipeDAO)
		expectedRecipe *m.Recipe
		expectError    bool
	}{
		{
			name: "Success",
			id:   id,
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().GetRecipe(id).Return(getRecipeSuccess, nil)
			},
			expectedRecipe: getRecipeSuccess,
			expectError:    false,
		},
		{
			name: "Error",
			id:   id,
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().GetRecipe(id).Return(nil, errors.New("recipe not found"))
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDAO := d.NewMockRecipeDAO(t)
			tc.setupMock(mockDAO)

			service := &RecipeService{recipeDAO: mockDAO}
			recipe, err := service.GetRecipe(tc.id)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRecipe, recipe)
			}
		})
	}
}

func TestCreateRecipe(t *testing.T) {
	createRecipeSuccess := &m.Recipe{
		ID:   uuid.New(),
		Name: "Pizza",
	}
	createRecipeError := &m.Recipe{
		ID:   uuid.New(),
		Name: "Burger",
	}

	testCases := []struct {
		name           string
		recipe         *m.Recipe
		setupMock      func(*d.MockRecipeDAO)
		expectedRecipe *m.Recipe
		expectError    bool
	}{
		{
			name:   "Success",
			recipe: createRecipeSuccess,
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().CreateRecipe(createRecipeSuccess).Return(nil)
			},
			expectedRecipe: createRecipeSuccess,
			expectError:    false,
		},
		{
			name:   "Error",
			recipe: createRecipeError,
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().CreateRecipe(createRecipeError).Return(errors.New("database error"))
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDAO := d.NewMockRecipeDAO(t)
			tc.setupMock(mockDAO)

			service := &RecipeService{recipeDAO: mockDAO}
			recipe, err := service.CreateRecipe(tc.recipe)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRecipe, recipe)
			}
		})
	}
}

func TestListRecipes(t *testing.T) {
	listRecipeSuccess := []*m.Recipe{
		{ID: uuid.New(), Name: "Soup"},
		{ID: uuid.New(), Name: "Salad"},
	}

	testCases := []struct {
		name           string
		setupMock      func(*d.MockRecipeDAO)
		expectedRecipe []*m.Recipe
		expectError    bool
	}{
		{
			name: "Success",
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().GetAllRecipes().Return(listRecipeSuccess, nil)
			},
			expectedRecipe: listRecipeSuccess,
			expectError:    false,
		},
		{
			name: "Error",
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().GetAllRecipes().Return(nil, errors.New("database error"))
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDAO := d.NewMockRecipeDAO(t)
			tc.setupMock(mockDAO)

			service := &RecipeService{recipeDAO: mockDAO}
			recipes, err := service.ListRecipes()

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRecipe, recipes)
			}
		})
	}
}

func TestUpdateRecipe(t *testing.T) {
	updateRecipeSuccess := &m.Recipe{
		ID:   uuid.New(),
		Name: "Taco",
	}
	updateRecipeError := &m.Recipe{
		ID:   uuid.New(),
		Name: "Burrito",
	}

	testCases := []struct {
		name           string
		recipe         *m.Recipe
		setupMock      func(*d.MockRecipeDAO)
		expectedRecipe *m.Recipe
		expectError    bool
	}{
		{
			name:   "Success",
			recipe: updateRecipeSuccess,
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().UpdateRecipe(updateRecipeSuccess).Return(nil)
			},
			expectedRecipe: updateRecipeSuccess,
			expectError:    false,
		},
		{
			name:   "Error",
			recipe: updateRecipeError,
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().UpdateRecipe(updateRecipeError).Return(errors.New("database error"))
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDAO := d.NewMockRecipeDAO(t)
			tc.setupMock(mockDAO)

			service := &RecipeService{recipeDAO: mockDAO}
			recipe, err := service.UpdateRecipe(tc.recipe)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRecipe, recipe)
			}
		})
	}
}

func TestDeleteRecipe(t *testing.T) {
	id := uuid.New()

	testCases := []struct {
		name        string
		id          uuid.UUID
		setupMock   func(*d.MockRecipeDAO)
		expectError bool
	}{
		{
			name: "Success",
			id:   id,
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().DeleteRecipe(id).Return(nil)
			},
			expectError: false,
		},
		{
			name: "Error",
			id:   id,
			setupMock: func(mockDAO *d.MockRecipeDAO) {
				mockDAO.EXPECT().DeleteRecipe(id).Return(errors.New("database error"))
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDAO := d.NewMockRecipeDAO(t)
			tc.setupMock(mockDAO)

			service := &RecipeService{recipeDAO: mockDAO}
			err := service.DeleteRecipe(tc.id)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
