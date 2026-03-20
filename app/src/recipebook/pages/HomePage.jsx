import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
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

const FeaturedSection = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(4),
  background: 'linear-gradient(135deg, #C41E3A 0%, #8B1428 100%)',
  color: '#fff',
  borderRadius: '4px',
  marginBottom: theme.spacing(4),
  textAlign: 'center',
}));

const HomePage = () => {
  const navigate = useNavigate();
  const [recipes, setRecipes] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

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
    navigate(`/recipes/${recipe.id}`);
  };

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
            <Grid item xs={12} sm={6} md={4} key={recipe.id || recipe.name}>
              <RecipeCard
                recipe={recipe}
                onClick={() => handleRecipeSelect(recipe)}
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
