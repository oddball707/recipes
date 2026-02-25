# Recipes PostgreSQL DAO

This package provides a Data Access Object (DAO) layer for interfacing with a PostgreSQL database to manage recipes, ingredients, and instructions. It uses `pgx`, a modern PostgreSQL driver for Go with connection pooling support.

## Setup

### 1. Install pgx Driver
The `github.com/jackc/pgx/v5` driver has been added to `go.mod`. Run:

```bash
go mod download
```

### 2. Create the Database
Connect to PostgreSQL and create the database:

```sql
CREATE DATABASE recipes;
```

### 3. Run Schema
Execute the schema file to create tables:

```bash
psql -U <username> -d recipes -f dao/schema.sql
```

Or execute the SQL queries directly:

```sql
-- Run the contents of dao/schema.sql in your PostgreSQL client
```

## Usage

### Initialize the Database Connection

```go
import "github.com/oddball707/recipes/dao"

// Create database connection with connection pooling
database, err := dao.NewDatabase(
    "localhost",  // host
    "5432",       // port
    "username",   // user
    "password",   // password
    "recipes",    // database name
)
if err != nil {
    log.Fatal(err)
}
defer database.Close()

// Create recipe DAO
recipeDAO := dao.NewRecipeDAO(database)
```

### CRUD Operations

#### Create a Recipe

```go
recipe := &model.Recipe{
    ID:          uuid.New(),
    Name:        "Chocolate Cake",
    Description: "A delicious chocolate cake",
    Ingredients: []model.Ingredient{
        {
            ID:       uuid.New(),
            Name_:    "Chocolate",
            Quantity: 200,
            Unit_:    model.UnitGram,
        },
    },
    Instructions: []model.Instruction{
        {
            ID:      uuid.New(),
            StepNum: 1,
            Text_:   "Preheat oven to 350F",
        },
    },
}

err := recipeDAO.CreateRecipe(recipe)
```

#### Get a Recipe

```go
recipe, err := recipeDAO.GetRecipe(recipeID)
if err != nil {
    log.Fatal(err)
}
```

#### Get All Recipes

```go
recipes, err := recipeDAO.GetAllRecipes()
if err != nil {
    log.Fatal(err)
}
```

#### Update a Recipe

```go
recipe.Name = "Updated Cake Name"
err := recipeDAO.UpdateRecipe(recipe)
```

#### Delete a Recipe

```go
err := recipeDAO.DeleteRecipe(recipeID)
```

## Database Schema

### Recipes Table
- `id` (UUID): Primary key
- `name` (VARCHAR): Recipe name
- `description` (TEXT): Recipe description
- `created_at` (TIMESTAMP): Creation timestamp
- `updated_at` (TIMESTAMP): Last update timestamp

### Ingredients Table
- `id` (UUID): Primary key
- `recipe_id` (UUID): Foreign key to recipes table
- `name` (VARCHAR): Ingredient name
- `quantity` (FLOAT): Quantity amount
- `unit` (VARCHAR): Unit of measurement
- `created_at` (TIMESTAMP): Creation timestamp

### Instructions Table
- `id` (UUID): Primary key
- `recipe_id` (UUID): Foreign key to recipes table
- `step_number` (INT): Order of instruction
- `text` (TEXT): Instruction text
- `created_at` (TIMESTAMP): Creation timestamp

## Features

- **Connection Pooling**: Efficient connection pooling with pgx for better performance
- **Context Support**: All operations are context-aware for proper cancellation and timeout handling
- **Transaction Support**: Create, Update, and Delete operations use database transactions to ensure data consistency
- **Cascading Deletes**: Deleting a recipe automatically deletes associated ingredients and instructions
- **Error Handling**: Comprehensive error messages for debugging
- **Connection Management**: Easy database connection creation and cleanup

## Environment Variables

Consider storing database credentials in environment variables for security:

```go
import "os"

host := os.Getenv("DB_HOST")
port := os.Getenv("DB_PORT")
user := os.Getenv("DB_USER")
password := os.Getenv("DB_PASSWORD")
dbname := os.Getenv("DB_NAME")

database, err := dao.NewDatabase(host, port, user, password, dbname)
```
