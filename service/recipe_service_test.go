package service

// import (
//     "errors"
//     "testing"

//     "github.com/google/uuid"
//     m "github.com/oddball707/recipes/model"
// )

// func TestGetRecipe_Success(t *testing.T) {
//     id := uuid.New()
//     expectedRecipe := &m.Recipe{
//         ID:   id,
//         Name: "Pasta",
//     }

//     mockDAO := &MockRecipeDAO{
//         getRecipeFunc: func(uid uuid.UUID) (*m.Recipe, error) {
//             return expectedRecipe, nil
//         },
//     }

//     service := &RecipeService{recipeDAO: mockDAO}
//     recipe, err := service.GetRecipe(id)

//     if err != nil {
//         t.Errorf("Expected no error, got %v", err)
//     }
//     if recipe.Name != "Pasta" {
//         t.Errorf("Expected recipe name 'Pasta', got %s", recipe.Name)
//     }
// }

// func TestGetRecipe_Error(t *testing.T) {
//     id := uuid.New()
//     mockDAO := &MockRecipeDAO{
//         getRecipeFunc: func(uid uuid.UUID) (*m.Recipe, error) {
//             return nil, errors.New("recipe not found")
//         },
//     }

//     service := &RecipeService{recipeDAO: mockDAO}
//     _, err := service.GetRecipe(id)

//     if err == nil {
//         t.Errorf("Expected error, got nil")
//     }
// }

// func TestCreateRecipe_Success(t *testing.T) {
//     recipe := &m.Recipe{
//         ID:   uuid.New(),
//         Name: "Pizza",
//     }

//     mockDAO := &MockRecipeDAO{
//         createRecipeFunc: func(r *m.Recipe) error {
//             return nil
//         },
//     }

//     service := &RecipeService{recipeDAO: mockDAO}
//     result, err := service.CreateRecipe(recipe)

//     if err != nil {
//         t.Errorf("Expected no error, got %v", err)
//     }
//     if result.Name != "Pizza" {
//         t.Errorf("Expected recipe name 'Pizza', got %s", result.Name)
//     }
// }

// func TestCreateRecipe_Error(t *testing.T) {
//     recipe := &m.Recipe{
//         ID:   uuid.New(),
//         Name: "Burger",
//     }

//     mockDAO := &MockRecipeDAO{
//         createRecipeFunc: func(r *m.Recipe) error {
//             return errors.New("database error")
//         },
//     }

//     service := &RecipeService{recipeDAO: mockDAO}
//     _, err := service.CreateRecipe(recipe)

//     if err == nil {
//         t.Errorf("Expected error, got nil")
//     }
// }
