import React from 'react';
import {
  AppBar,
  Toolbar,
  Box,
  Container,
  Typography,
} from '@mui/material';
import { styled } from '@mui/material/styles';
import RecipeMenu from './RecipeMenu';

const StyledAppBar = styled(AppBar)(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
  color: theme.palette.text.primary,
  boxShadow: '0 2px 4px rgba(0, 0, 0, 0.08)',
  borderBottom: `1px solid ${theme.palette.divider}`,
}));

const Navigation = () => {

  return (
    <>
      <StyledAppBar position="sticky">
        <Container maxWidth="lg">
          <Toolbar disableGutters sx={{ py: 1 }}>
            {/* Navigation Menu */}
            <RecipeMenu />
          </Toolbar>
        </Container>
      </StyledAppBar>

      {/* Hero Section */}
      <Box
        sx={{
          background: 'white ',
          py: { xs: 6, md: 8 },
          textAlign: 'center',
          mb: 4,
        }}
      >
        <Container maxWidth="md">
          <img
            src="/spiceandstilleto.png"
            alt="Spice and Stilleto Logo"
            style={{
              width: 'auto',
              objectFit: 'contain',
              marginBottom: '1rem',
            }}
          />
          <Typography variant="subtitle1">
            A fabulous recipe collection for the hungry home chef
          </Typography>
        </Container>
      </Box>
    </>
  );
};

export default Navigation;
