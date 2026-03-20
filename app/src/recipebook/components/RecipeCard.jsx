import React from 'react';
import {
  Card,
  CardMedia,
  CardContent,
  Typography,
  Box,
  Chip,
  Stack,
} from '@mui/material';
import { styled } from '@mui/material/styles';

const StyledCard = styled(Card)(({ theme }) => ({
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  cursor: 'pointer',
  transition: 'all 0.3s ease',
  '&:hover': {
    transform: 'translateY(-4px)',
    boxShadow: '0 12px 24px rgba(196, 30, 58, 0.2)',
  },
}));

const StyledCardMedia = styled(CardMedia)({
  height: '250px',
  backgroundSize: 'cover',
  backgroundPosition: 'center',
  backgroundColor: '#E8E8E8',
});

const RecipeCard = ({ recipe, onClick }) => {
  // Fallback gradient colors - red/white scheme
  const getCardColor = (name) => {
    const colors = [
      'linear-gradient(135deg, #C41E3A 0%, #FF6B6B 100%)',
      'linear-gradient(135deg, #8B1428 0%, #C41E3A 100%)',
      'linear-gradient(135deg, #E63946 0%, #A4161A 100%)',
      'linear-gradient(135deg, #D62828 0%, #F77F00 100%)',
    ];
    const hash = name.charCodeAt(0);
    return colors[hash % colors.length];
  };

  return (
    <StyledCard onClick={() => onClick && onClick(recipe)}>
      <StyledCardMedia
        sx={{
          background: getCardColor(recipe.name),
        }}
      >
        {/* Image placeholder or actual image if available */}
      </StyledCardMedia>
      <CardContent sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column' }}>
        <Typography
          variant="h6"
          sx={{
            fontWeight: 600,
            color: 'primary.main',
            mb: 1,
            lineHeight: 1.4,
            minHeight: '2.8em',
          }}
        >
          {recipe.name}
        </Typography>

        <Typography
          variant="body2"
          color="text.secondary"
          sx={{
            mb: 2,
            display: '-webkit-box',
            WebkitLineClamp: 2,
            WebkitBoxOrient: 'vertical',
            overflow: 'hidden',
            lineHeight: 1.5,
          }}
        >
          {recipe.description}
        </Typography>

        <Box sx={{ mt: 'auto' }}>
          {recipe.ingredients && recipe.ingredients.length > 0 && (
            <Stack direction="row" spacing={0.5} sx={{ flexWrap: 'wrap', gap: 0.5 }}>
              <Typography variant="caption" color="text.secondary" sx={{ width: '100%' }}>
                {recipe.ingredients.length} ingredient{recipe.ingredients.length !== 1 ? 's' : ''}
              </Typography>
            </Stack>
          )}

          {recipe.instructions && recipe.instructions.length > 0 && (
            <Typography variant="caption" color="text.secondary" sx={{ display: 'block', mt: 0.5 }}>
              {recipe.instructions.length} step{recipe.instructions.length !== 1 ? 's' : ''}
            </Typography>
          )}
        </Box>
      </CardContent>
    </StyledCard>
  );
};

export default RecipeCard;
