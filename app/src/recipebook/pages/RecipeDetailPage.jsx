import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Container, CircularProgress, Alert, Button } from '@mui/material';
import recipeApi from '../services/recipeApi';
import RecipeDetail from '../components/RecipeDetail';

const RecipeDetailPage = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [recipe, setRecipe] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchRecipe = async () => {
      try {
        setLoading(true);
        const data = await recipeApi.getRecipe(id);
        setRecipe(data);
      } catch (err) {
        console.error('Failed to fetch recipe:', err);
        setError('Failed to load recipe. Please try again.');
      } finally {
        setLoading(false);
      }
    };

    fetchRecipe();
  }, [id]);

  const handleBack = () => {
    navigate('/');
  };

  const handleDelete = async (recipeId) => {
    try {
      await recipeApi.deleteRecipe(recipeId);
      navigate('/');
    } catch (err) {
      console.error('Failed to delete recipe:', err);
    }
  };

  const handleEdit = (recipe) => {
    // TODO: Implement edit
    console.log('Edit recipe:', recipe);
  };

  if (loading) {
    return (
      <Container maxWidth="md" sx={{ py: 8, textAlign: 'center' }}>
        <CircularProgress />
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="md" sx={{ py: 4 }}>
        <Alert severity="error">{error}</Alert>
        <Button onClick={handleBack} sx={{ mt: 2 }}>Back to Recipes</Button>
      </Container>
    );
  }

  return (
    <RecipeDetail
      recipe={recipe}
      onBack={handleBack}
      onEdit={handleEdit}
      onDelete={handleDelete}
    />
  );
};

export default RecipeDetailPage;
