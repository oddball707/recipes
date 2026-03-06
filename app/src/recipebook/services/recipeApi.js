import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const recipeApi = {
  /**
   * Get all recipes
   */
  async listRecipes() {
    try {
      const response = await api.post('/list', {});
      return response.data || [];
    } catch (error) {
      console.error('Error fetching recipes:', error);
      throw error;
    }
  },

  /**
   * Get a single recipe by ID
   */
  async getRecipe(id) {
    try {
      const response = await api.post('/get', { id });
      return response.data;
    } catch (error) {
      console.error('Error fetching recipe:', error);
      throw error;
    }
  },

  /**
   * Create a new recipe
   */
  async createRecipe(recipe) {
    try {
      const response = await api.post('/create', recipe);
      return response.data;
    } catch (error) {
      console.error('Error creating recipe:', error);
      throw error;
    }
  },

  /**
   * Update an existing recipe
   */
  async updateRecipe(recipe) {
    try {
      const response = await api.post('/update', recipe);
      return response.data;
    } catch (error) {
      console.error('Error updating recipe:', error);
      throw error;
    }
  },

  /**
   * Delete a recipe
   */
  async deleteRecipe(id) {
    try {
      const response = await api.post('/delete', { id });
      return response.data;
    } catch (error) {
      console.error('Error deleting recipe:', error);
      throw error;
    }
  },

  /**
   * Check API health
   */
  async health() {
    try {
      const response = await api.get('/health');
      return response.data;
    } catch (error) {
      console.error('API health check failed:', error);
      throw error;
    }
  },
};

export default recipeApi;
