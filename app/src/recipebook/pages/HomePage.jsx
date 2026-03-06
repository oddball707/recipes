import React, { useState, useEffect } from 'react';
import {
  Container,
  Grid,
  Box,
  Typography,
  Button,
  CircularProgress,
  Alert,
  Paper,
} from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import { styled } from '@mui/material/styles';
import recipeApi from '../services/recipeApi';
import RecipeCard from '../components/RecipeCard.jsx';
import RecipeDetail from '../components/RecipeDetail.jsx';

const FeaturedSection = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(4),
  background: 'linear-gradient(135deg, #C41E3A 0%, #8B1428 100%)',
  color: '#fff',
  borderRadius: '4px',
  marginBottom: theme.spacing(4),
  textAlign: 'center',
}));

const HomePage = () => {
  const [recipes, setRecipes] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [selectedRecipe, setSelectedRecipe] = useState(null);

  // Fetch recipes on component mount
  useEffect(() => {
    fetchRecipes();
  }, []);

  const fetchRecipes = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await recipeApi.listRecipes();
      setRecipes(data || []);
    } catch (err) {
      console.error('Failed to fetch recipes:', err);
      setError('Failed to load recipes. Please try again later.');
    } finally {
      setLoading(false);
    }
  };

  const handleRecipeSelect = (recipe) => {
    setSelectedRecipe(recipe);
  };

  const handleBackToList = () => {
    setSelectedRecipe(null);
  };

  const handleDeleteRecipe = async (recipeId) => {
    try {
      await recipeApi.deleteRecipe(recipeId);
      fetchRecipes();
      setSelectedRecipe(null);
    } catch (err) {
      console.error('Failed to delete recipe:', err);
      setError('Failed to delete recipe. Please try again.');
    }
  };

  const handleEditRecipe = (recipe) => {
    // TODO: Implement edit modal/form
    console.log('Edit recipe:', recipe);
  };

  // Show recipe detail view
  if (selectedRecipe) {
    return (
      <RecipeDetail
        recipe={selectedRecipe}
        onBack={handleBackToList}
        onEdit={handleEditRecipe}
        onDelete={handleDeleteRecipe}
      />
    );
  }

  // Show loading state
  if (loading) {
    return (
      <Container maxWidth="lg" sx={{ py: 8, textAlign: 'center' }}>
        <CircularProgress />
        <Typography sx={{ mt: 2 }}>Loading recipes...</Typography>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ py: 4 }}>
      {/* Featured Section */}
      {recipes.length > 0 && (
        <FeaturedSection elevation={0}>
          <Typography variant="body1" sx={{ mb: 2, opacity: 0.95 }}>
            Discover fabulous recipes for the hungry home chef
          </Typography>
          <Typography variant="h2" sx={{ fontWeight: 700, opacity: 0.9 }}>
            {recipes.length} Recipes
          </Typography>
        </FeaturedSection>
      )}

      {/* Error Alert */}
      {error && <Alert severity="error" sx={{ mb: 3 }}>{error}</Alert>}

      {/* Recipes Grid */}
      {recipes.length > 0 ? (
        <Grid container spacing={3}>
          {recipes.map((recipe) => (
            <Grid item xs={12} sm={6} md={4} key={recipe.ID || recipe.Name}>
              <RecipeCard
                recipe={recipe}
                onClick={handleRecipeSelect}
              />
            </Grid>
          ))}
        </Grid>
      ) : (
        <Box sx={{ textAlign: 'center', py: 8 }}>
          <Typography variant="h6" color="text.secondary" sx={{ mb: 2 }}>
            No recipes yet. Add the first!
          </Typography>
          <Button variant="contained" color="primary" startIcon={<AddIcon />}>
            Add Recipe
          </Button>
        </Box>
      )}
    </Container>
  );
};

export default HomePage;
