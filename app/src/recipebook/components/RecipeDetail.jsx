import React, { useState, useEffect } from 'react';
import {
  Container,
  Box,
  Typography,
  Paper,
  Grid,
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableRow,
  Divider,
  CircularProgress,
  Alert,
  IconButton,
} from '@mui/material';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import { styled } from '@mui/material/styles';
import recipeApi from '../services/recipeApi';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  backgroundColor: theme.palette.background.paper,
  borderRadius: '4px',
}));

const SectionTitle = styled(Typography)(({ theme }) => ({
  fontSize: '1.25rem',
  fontWeight: 600,
  color: theme.palette.primary.main,
  marginTop: theme.spacing(3),
  marginBottom: theme.spacing(1.5),
  paddingBottom: theme.spacing(1),
  borderBottom: `2px solid ${theme.palette.secondary.main}`,
}));

const RecipeDetail = ({ recipe, onBack, onEdit, onDelete }) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  if (loading) {
    return (
      <Container maxWidth="md" sx={{ py: 4, textAlign: 'center' }}>
        <CircularProgress />
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="md" sx={{ py: 4 }}>
        <Alert severity="error">{error}</Alert>
        <Button onClick={onBack} sx={{ mt: 2 }}>
          Back to Recipes
        </Button>
      </Container>
    );
  }

  if (!recipe) {
    return (
      <Container maxWidth="md" sx={{ py: 4 }}>
        <Alert severity="info">Recipe not found</Alert>
        <Button onClick={onBack} sx={{ mt: 2 }}>
          Back to Recipes
        </Button>
      </Container>
    );
  }

  return (
    <Container maxWidth="md" sx={{ py: 4 }}>
      {/* Header */}
      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, mb: 3 }}>
        <IconButton onClick={onBack} size="small">
          <ArrowBackIcon />
        </IconButton>
        <Typography variant="h4" sx={{ flex: 1 }}>
          {recipe.Name}
        </Typography>
        <Button
          startIcon={<EditIcon />}
          variant="outlined"
          color="primary"
          onClick={() => onEdit && onEdit(recipe)}
          size="small"
        >
          Edit
        </Button>
        <Button
          startIcon={<DeleteIcon />}
          variant="outlined"
          color="error"
          onClick={() => {
            if (window.confirm('Are you sure you want to delete this recipe?')) {
              onDelete && onDelete(recipe.ID);
            }
          }}
          size="small"
        >
          Delete
        </Button>
      </Box>

      <StyledPaper>
        {/* Description */}
        <Typography variant="body1" paragraph sx={{ color: 'text.secondary', fontSize: '1.1rem' }}>
          {recipe.Description}
        </Typography>

        <Divider sx={{ my: 3 }} />

        {/* Ingredients */}
        {recipe.Ingredients && recipe.Ingredients.length > 0 && (
          <>
            <SectionTitle>Ingredients</SectionTitle>
            <TableContainer>
              <Table>
                <TableBody>
                  {recipe.Ingredients.map((ingredient, idx) => (
                    <TableRow key={idx}>
                      <TableCell sx={{ border: 'none', py: 1 }}>
                        <Typography variant="body2">
                          <strong>{ingredient.Quantity}</strong>{' '}
                          {ingredient.Unit?.Abbreviation || ingredient.Unit?.Name || ''}
                        </Typography>
                      </TableCell>
                      <TableCell sx={{ border: 'none', py: 1 }}>
                        <Typography variant="body2">{ingredient.Name}</Typography>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </>
        )}

        <Divider sx={{ my: 3 }} />

        {/* Instructions */}
        {recipe.Instructions && recipe.Instructions.length > 0 && (
          <>
            <SectionTitle>Instructions</SectionTitle>
            <Box>
              {recipe.Instructions.map((instruction, idx) => (
                <Box key={idx} sx={{ mb: 2 }}>
                  <Typography
                    variant="body2"
                    sx={{
                      display: 'flex',
                      gap: 2,
                      mb: 1,
                    }}
                  >
                    <Box
                      sx={{
                        minWidth: '24px',
                        height: '24px',
                        borderRadius: '50%',
                        backgroundColor: 'primary.main',
                        color: '#fff',
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        fontWeight: 600,
                        fontSize: '0.85rem',
                        flexShrink: 0,
                      }}
                    >
                      {idx + 1}
                    </Box>
                    <Typography variant="body2" sx={{ pt: 0.25 }}>
                      {instruction.Instruction}
                    </Typography>
                  </Typography>
                </Box>
              ))}
            </Box>
          </>
        )}
      </StyledPaper>
    </Container>
  );
};

export default RecipeDetail;
