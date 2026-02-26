package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oddball707/recipes/model"
)

type PostgresRecipeDAO struct {
	pool *pgxpool.Pool
}

type RecipeDAO interface {
	CreateRecipe(recipe *model.Recipe) error
	GetRecipe(id uuid.UUID) (*model.Recipe, error)
	GetAllRecipes() ([]*model.Recipe, error)
	UpdateRecipe(recipe *model.Recipe) error
	DeleteRecipe(id uuid.UUID) error
}

// NewRecipeDAO creates a new recipe DAO
func NewRecipeDAO(database *Database) RecipeDAO {
	return &PostgresRecipeDAO{pool: database.Pool}
}

// CreateRecipe inserts a new recipe with its ingredients and instructions
func (r *PostgresRecipeDAO) CreateRecipe(recipe *model.Recipe) error {
	ctx := context.Background()
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Insert recipe
	recipeQuery := `INSERT INTO recipes (name, description) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(ctx, recipeQuery, recipe.Name, recipe.Description).Scan(&recipe.ID)
	if err != nil {
		return fmt.Errorf("failed to insert recipe: %w", err)
	}

	// Insert ingredients
	ingredientQuery := `INSERT INTO ingredients (recipe_id, name, quantity, unit_id) VALUES ($1, $2, $3, $4)`
	for _, ingredient := range recipe.Ingredients {
		unitID, err := r.getOrCreateUnit(ctx, tx, ingredient.Unit)
		if err != nil {
			log.Printf("failed to get or create unit: %v", err)
			return err
		}
		_, err = tx.Exec(ctx, ingredientQuery, recipe.ID, ingredient.Name, ingredient.Quantity, unitID)
		if err != nil {
			log.Printf("failed to insert ingredient: %v", err)
			return fmt.Errorf("failed to insert ingredient: %w", err)
		}
	}

	// Insert instructions
	instructionQuery := `INSERT INTO instructions (recipe_id, step_number, text) VALUES ($1, $2, $3)`
	for _, instruction := range recipe.Instructions {
		_, err = tx.Exec(ctx, instructionQuery, recipe.ID, instruction.StepNumber, instruction.Text)
		if err != nil {
			log.Printf("failed to insert instruction: %v", err)
			return fmt.Errorf("failed to insert instruction: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// GetRecipe retrieves a recipe by ID with its ingredients and instructions
func (r *PostgresRecipeDAO) GetRecipe(id uuid.UUID) (*model.Recipe, error) {
	ctx := context.Background()
	recipe := &model.Recipe{}

	// Get recipe
	recipeQuery := `SELECT id, name, description FROM recipes WHERE id = $1`
	err := r.pool.QueryRow(ctx, recipeQuery, id).Scan(&recipe.ID, &recipe.Name, &recipe.Description)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("recipe not found")
		}
		return nil, fmt.Errorf("failed to get recipe: %w", err)
	}

	// Get ingredients
	ingredients, err := r.getIngredients(id)
	if err != nil {
		return nil, err
	}
	recipe.Ingredients = ingredients

	// Get instructions
	instructions, err := r.getInstructions(id)
	if err != nil {
		return nil, err
	}
	recipe.Instructions = instructions

	return recipe, nil
}

// GetAllRecipes retrieves all recipes
func (r *PostgresRecipeDAO) GetAllRecipes() ([]*model.Recipe, error) {
	ctx := context.Background()
	query := `SELECT id, name, description FROM recipes ORDER BY name`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get recipes: %w", err)
	}
	defer rows.Close()

	var recipes []*model.Recipe
	for rows.Next() {
		recipe := &model.Recipe{}
		err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan recipe: %w", err)
		}

		// Get ingredients and instructions for each recipe
		ingredients, err := r.getIngredients(recipe.ID)
		if err != nil {
			return nil, err
		}
		recipe.Ingredients = ingredients

		instructions, err := r.getInstructions(recipe.ID)
		if err != nil {
			return nil, err
		}
		recipe.Instructions = instructions

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

// UpdateRecipe updates an existing recipe
func (r *PostgresRecipeDAO) UpdateRecipe(recipe *model.Recipe) error {
	ctx := context.Background()
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Update recipe
	updateQuery := `UPDATE recipes SET name = $1, description = $2 WHERE id = $3`
	_, err = tx.Exec(ctx, updateQuery, recipe.Name, recipe.Description, recipe.ID)
	if err != nil {
		return fmt.Errorf("failed to update recipe: %w", err)
	}

	// Delete old ingredients
	_, err = tx.Exec(ctx, `DELETE FROM ingredients WHERE recipe_id = $1`, recipe.ID)
	if err != nil {
		return fmt.Errorf("failed to delete ingredients: %w", err)
	}

	// Insert new ingredients
	ingredientQuery := `INSERT INTO ingredients (recipe_id, name, quantity, unit_id) VALUES ($1, $2, $3, $4)`
	for _, ingredient := range recipe.Ingredients {
		unitID, err := r.getOrCreateUnit(ctx, tx, ingredient.Unit)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, ingredientQuery, recipe.ID, ingredient.Name, ingredient.Quantity, unitID)
		if err != nil {
			return fmt.Errorf("failed to insert ingredient: %w", err)
		}
	}

	// Delete old instructions
	_, err = tx.Exec(ctx, `DELETE FROM instructions WHERE recipe_id = $1`, recipe.ID)
	if err != nil {
		return fmt.Errorf("failed to delete instructions: %w", err)
	}

	// Insert new instructions
	instructionQuery := `INSERT INTO instructions (recipe_id, step_number, text) VALUES ($1, $2, $3)`
	for _, instruction := range recipe.Instructions {
		_, err = tx.Exec(ctx, instructionQuery, recipe.ID, instruction.StepNumber, instruction.Text)
		if err != nil {
			return fmt.Errorf("failed to insert instruction: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// DeleteRecipe deletes a recipe and its ingredients and instructions
func (r *PostgresRecipeDAO) DeleteRecipe(id uuid.UUID) error {
	ctx := context.Background()
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Delete ingredients
	_, err = tx.Exec(ctx, `DELETE FROM ingredients WHERE recipe_id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete ingredients: %w", err)
	}

	// Delete instructions
	_, err = tx.Exec(ctx, `DELETE FROM instructions WHERE recipe_id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete instructions: %w", err)
	}

	// Delete recipe
	result, err := tx.Exec(ctx, `DELETE FROM recipes WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete recipe: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("recipe not found")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Helper method to get ingredients for a recipe
func (r *PostgresRecipeDAO) getIngredients(recipeID uuid.UUID) ([]model.Ingredient, error) {
	ctx := context.Background()
	query := `
		SELECT i.name, i.quantity, u.name, u.abbreviation
		FROM ingredients i
		JOIN units u ON i.unit_id = u.id
		WHERE i.recipe_id = $1
		ORDER BY i.id`
	rows, err := r.pool.Query(ctx, query, recipeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get ingredients: %w", err)
	}
	defer rows.Close()

	var ingredients []model.Ingredient
	for rows.Next() {
		var ingredient model.Ingredient
		var unitName, unitAbbreviation string
		err := rows.Scan(&ingredient.Name, &ingredient.Quantity, &unitName, &unitAbbreviation)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ingredient: %w", err)
		}
		ingredient.Unit = &model.Unit{
			Name:         unitName,
			Abbreviation: unitAbbreviation,
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

// Helper method to get instructions for a recipe
func (r *PostgresRecipeDAO) getInstructions(recipeID uuid.UUID) ([]model.Instruction, error) {
	ctx := context.Background()
	query := `SELECT step_number, text FROM instructions WHERE recipe_id = $1 ORDER BY step_number`
	rows, err := r.pool.Query(ctx, query, recipeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get instructions: %w", err)
	}
	defer rows.Close()

	var instructions []model.Instruction
	for rows.Next() {
		var instruction model.Instruction
		err := rows.Scan(&instruction.StepNumber, &instruction.Text)
		if err != nil {
			return nil, fmt.Errorf("failed to scan instruction: %w", err)
		}
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

// getOrCreateUnit gets or creates a unit and returns its ID
func (r *PostgresRecipeDAO) getOrCreateUnit(ctx context.Context, tx pgx.Tx, unit *model.Unit) (uuid.UUID, error) {
	if unit == nil {
		unit = &model.Unit{Name: "", Abbreviation: ""}
	}
	var unitID uuid.UUID
	findUnitQuery := `SELECT id FROM units WHERE abbreviation = $1`
	err := tx.QueryRow(ctx, findUnitQuery, unit.Abbreviation).Scan(&unitID)
	if err == nil {
		return unitID, nil // Found existing unit
	}

	if err != pgx.ErrNoRows {
		log.Printf("failed to query for unit: %v", err)
		return uuid.Nil, fmt.Errorf("failed to query for unit: %w", err)
	}

	log.Printf("Unit (%s) not found, creating...", unit.Abbreviation)

	// Unit not found, create it
	unitID = uuid.New()
	insertUnitQuery := `INSERT INTO units (id, name, abbreviation) VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, insertUnitQuery, unitID, unit.Name, unit.Abbreviation)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert new unit: %w", err)
	}

	return unitID, nil
}
