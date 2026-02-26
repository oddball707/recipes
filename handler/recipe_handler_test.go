package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	m "github.com/oddball707/recipes/model"
	s "github.com/oddball707/recipes/service"
	"github.com/stretchr/testify/assert"
)

func TestGetRecipe(t *testing.T) {
	id := uuid.New()
	recipe := &m.Recipe{ID: id, Name: "Test Recipe"}

	testCases := []struct {
		name           string
		id             string
		setupMock      func(*s.MockRecipeClient)
		expectedStatus int
		expectedBody   *m.Recipe
	}{
		{
			name: "Success",
			id:   id.String(),
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().GetRecipe(id).Return(recipe, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   recipe,
		},
		{
			name: "Error - Not Found",
			id:   id.String(),
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().GetRecipe(id).Return(nil, errors.New("not found"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Error - Invalid ID",
			id:             "invalid-id",
			setupMock:      func(mock *s.MockRecipeClient) {},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := s.NewMockRecipeClient(t)
			tc.setupMock(mockService)
			handler := NewRecipeHandler(mockService)

			req := httptest.NewRequest("GET", "/recipe?id="+tc.id, nil)
			rr := httptest.NewRecorder()

			handler.GetRecipe(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
			if tc.expectedBody != nil {
				var actual m.Recipe
				json.Unmarshal(rr.Body.Bytes(), &actual)
				assert.Equal(t, *tc.expectedBody, actual)
			}
		})
	}
}

func TestListRecipes(t *testing.T) {
	recipes := []*m.Recipe{{ID: uuid.New(), Name: "Recipe 1"}}

	testCases := []struct {
		name           string
		setupMock      func(*s.MockRecipeClient)
		expectedStatus int
		expectedBody   []*m.Recipe
	}{
		{
			name: "Success",
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().ListRecipes().Return(recipes, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   recipes,
		},
		{
			name: "Error",
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().ListRecipes().Return(nil, errors.New("db error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := s.NewMockRecipeClient(t)
			tc.setupMock(mockService)
			handler := NewRecipeHandler(mockService)

			req := httptest.NewRequest("GET", "/recipes", nil)
			rr := httptest.NewRecorder()

			handler.ListRecipes(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
			if tc.expectedBody != nil {
				var actual []*m.Recipe
				json.Unmarshal(rr.Body.Bytes(), &actual)
				assert.Equal(t, tc.expectedBody, actual)
			}
		})
	}
}

func TestCreateRecipe(t *testing.T) {
	recipe := &m.Recipe{Name: "New Recipe"}
	recipeWithID := &m.Recipe{ID: uuid.New(), Name: "New Recipe"}

	testCases := []struct {
		name           string
		body           *m.Recipe
		setupMock      func(*s.MockRecipeClient)
		expectedStatus int
		expectedBody   *m.Recipe
	}{
		{
			name: "Success",
			body: recipe,
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().CreateRecipe(recipe).Return(recipeWithID, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   recipeWithID,
		},
		{
			name: "Error",
			body: recipe,
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().CreateRecipe(recipe).Return(nil, errors.New("create error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := s.NewMockRecipeClient(t)
			tc.setupMock(mockService)
			handler := NewRecipeHandler(mockService)

			bodyBytes, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/recipe", bytes.NewReader(bodyBytes))
			rr := httptest.NewRecorder()

			handler.CreateRecipe(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
			if tc.expectedBody != nil {
				var actual m.Recipe
				json.Unmarshal(rr.Body.Bytes(), &actual)
				assert.Equal(t, *tc.expectedBody, actual)
			}
		})
	}
}

func TestUpdateRecipe(t *testing.T) {
	recipe := &m.Recipe{ID: uuid.New(), Name: "Updated Recipe"}

	testCases := []struct {
		name           string
		body           *m.Recipe
		setupMock      func(*s.MockRecipeClient)
		expectedStatus int
		expectedBody   *m.Recipe
	}{
		{
			name: "Success",
			body: recipe,
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().UpdateRecipe(recipe).Return(recipe, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   recipe,
		},
		{
			name: "Error",
			body: recipe,
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().UpdateRecipe(recipe).Return(nil, errors.New("update error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := s.NewMockRecipeClient(t)
			tc.setupMock(mockService)
			handler := NewRecipeHandler(mockService)

			bodyBytes, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("PUT", "/recipe", bytes.NewReader(bodyBytes))
			rr := httptest.NewRecorder()

			handler.UpdateRecipe(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
			if tc.expectedBody != nil {
				var actual m.Recipe
				json.Unmarshal(rr.Body.Bytes(), &actual)
				assert.Equal(t, *tc.expectedBody, actual)
			}
		})
	}
}

func TestDeleteRecipe(t *testing.T) {
	id := uuid.New()

	testCases := []struct {
		name           string
		id             string
		setupMock      func(*s.MockRecipeClient)
		expectedStatus int
	}{
		{
			name: "Success",
			id:   id.String(),
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().DeleteRecipe(id).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Error",
			id:   id.String(),
			setupMock: func(mock *s.MockRecipeClient) {
				mock.EXPECT().DeleteRecipe(id).Return(errors.New("delete error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Error - Invalid ID",
			id:             "invalid-id",
			setupMock:      func(mock *s.MockRecipeClient) {},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := s.NewMockRecipeClient(t)
			tc.setupMock(mockService)
			handler := NewRecipeHandler(mockService)

			req := httptest.NewRequest("DELETE", "/recipe?id="+tc.id, nil)
			rr := httptest.NewRecorder()

			handler.DeleteRecipe(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}

func TestParseGetReq(t *testing.T) {
	id := uuid.New()
	testCases := []struct {
		name          string
		url           string
		expectedID    uuid.UUID
		expectError   bool
		expectedError string
	}{
		{
			name:        "Valid UUID",
			url:         "/recipe?id=" + id.String(),
			expectedID:  id,
			expectError: false,
		},
		{
			name:          "No ID",
			url:           "/recipe",
			expectError:   true,
			expectedError: "id is required",
		},
		{
			name:        "Invalid UUID",
			url:         "/recipe?id=invalid",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.url, nil)
			parsedID, err := parseGetReq(req)

			if tc.expectError {
				assert.Error(t, err)
				if tc.expectedError != "" {
					assert.Equal(t, tc.expectedError, err.Error())
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedID, parsedID)
			}
		})
	}
}

func TestParseCreateReq(t *testing.T) {
	recipe := &m.Recipe{Name: "Test"}
	recipeBody, _ := json.Marshal(recipe)

	testCases := []struct {
		name           string
		body           []byte
		expectedRecipe *m.Recipe
		expectError    bool
	}{
		{
			name:           "Valid Body",
			body:           recipeBody,
			expectedRecipe: recipe,
			expectError:    false,
		},
		{
			name:        "Invalid JSON",
			body:        []byte("{invalid"),
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/recipe", bytes.NewReader(tc.body))
			parsedRecipe, err := parseCreateReq(req)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRecipe.Name, parsedRecipe.Name)
			}
		})
	}
}
