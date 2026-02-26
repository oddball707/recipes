package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	m "github.com/oddball707/recipes/model"
	s "github.com/oddball707/recipes/service"
)

type Handler struct {
	service s.RecipeClient
}

type RecipeHandler interface {
	HealthHandler(w http.ResponseWriter, r *http.Request)
	ReadinessHandler(w http.ResponseWriter, r *http.Request)
	GetRecipe(w http.ResponseWriter, r *http.Request)
	CreateRecipe(w http.ResponseWriter, r *http.Request)
	UpdateRecipe(w http.ResponseWriter, r *http.Request)
	DeleteRecipe(w http.ResponseWriter, r *http.Request)
}

type CreateRequest struct {
	RecipeName  string `json:"recipeName"`
	Description string `json:"description"`

	Ingredients  []m.Ingredient  `json:"ingredients"`
	Instructions []m.Instruction `json:"instructions"`
}

func NewRecipeHandler(service s.RecipeClient) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Healthy")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Ready")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := parseGetReq(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	recipe, err := h.service.GetRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func (h *Handler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.service.ListRecipes()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func (h *Handler) CreateRecipe(w http.ResponseWriter, r *http.Request) {

	recipe, err := parseCreateReq(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	Recipe, err := h.service.CreateRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Recipe)
}

func (h *Handler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	recipe, err := parseCreateReq(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	Recipe, err := h.service.UpdateRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Recipe)
}

func (h *Handler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := parseGetReq(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = h.service.DeleteRecipe(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func parseGetReq(r *http.Request) (uuid.UUID, error) {
	id_string := r.URL.Query().Get("id")
	if id_string == "" {
		log.Println("Id required for get")
		return uuid.Nil, errors.New("id is required")
	}
	id, err := uuid.Parse(id_string)
	if err != nil {
		log.Print("Error parsing id - ", err)
		return uuid.Nil, err
	}
	return id, nil
}

func parseCreateReq(r *http.Request) (*m.Recipe, error) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Print("Error reading request body - ", err)
		return nil, err
	}

	var msg m.Recipe
	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Print("Error Unmarshalling request - ", err)
		return nil, err
	}

	return &msg, nil
}
