import React, { useState } from 'react';
import {
  Box,
  Menu,
  MenuItem,
} from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import { styled } from '@mui/material/styles';

const NavMenu = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  gap: '2rem',
}));

const NavLink = styled(Box)(({ theme, isFooter }) => ({
  fontSize: isFooter ? '12px' : '10px',
  fontWeight: 500,
  color: isFooter ? 'rgba(255, 255, 255, 0.85)' : theme.palette.text.primary,
  cursor: 'pointer',
  position: 'relative',
  display: 'flex',
  alignItems: 'center',
  gap: '0.25rem',
  transition: 'color 0.3s ease',
  textDecoration: 'none',
  '&:hover': {
    color: isFooter ? '#FFFFFF' : theme.palette.primary.main,
  },
  '& a': {
    color: 'inherit',
    textDecoration: 'none',
  },
}));

const categories = [
  { label: 'Appetizers & Sides', href: '#appetizers' },
  { label: 'Mains', href: '#mains' },
  { label: 'Salads', href: '#salads' },
  { label: 'Desserts', href: '#desserts' },
  { label: 'Soups', href: '#soups' },
  { label: 'Cocktails', href: '#cocktails' },
];

export const RECIPE_CATEGORIES = categories;

const MenuLink = ({ href, label, isButton = false, isFooter = false }) => (
  <NavLink isFooter={isFooter}>
    <a href={href}>{label}</a>
  </NavLink>
);

const RecipeMenu = ({ variant = 'horizontal', isFooter = false }) => {
  const [recipesMenuAnchor, setRecipesMenuAnchor] = useState(null);

  const handleRecipesMouseEnter = (event) => {
    setRecipesMenuAnchor(event.currentTarget);
  };

  const handleRecipesMouseLeave = () => {
    setRecipesMenuAnchor(null);
  };

  return (
    <NavMenu>
      <MenuLink href="#meet-stacy" label="Meet Stacy" isFooter={isFooter} />

      <NavLink
        isFooter={isFooter}
        onMouseEnter={handleRecipesMouseEnter}
        onMouseLeave={handleRecipesMouseLeave}
      >
        <span>Recipes</span>
        <ExpandMoreIcon sx={{ fontSize: '1.2rem' }} />

        {/* Dropdown Menu */}
        <Menu
          anchorEl={recipesMenuAnchor}
          open={Boolean(recipesMenuAnchor)}
          onClose={handleRecipesMouseLeave}
          onMouseLeave={handleRecipesMouseLeave}
          MenuListProps={{
            onMouseLeave: handleRecipesMouseLeave,
          }}
          sx={{
            '& .MuiMenu-paper': {
              marginTop: '0.5rem',
              boxShadow: '0 4px 12px rgba(0, 0, 0, 0.15)',
            },
          }}
        >
          {categories.map((category) => (
            <MenuItem
              key={category.label}
              onClick={handleRecipesMouseLeave}
              component="a"
              href={category.href}
              sx={{
                color: isFooter ? 'rgba(255, 255, 255, 0.85)' : 'text.primary',
                backgroundColor: isFooter ? '#C41E3A' : 'inherit',
                '&:hover': {
                  backgroundColor: isFooter ? 'rgba(0, 0, 0, 0.2)' : 'rgba(196, 30, 58, 0.08)',
                  color: isFooter ? '#FFFFFF' : 'primary.main',
                },
              }}
            >
              {category.label}
            </MenuItem>
          ))}
        </Menu>
      </NavLink>

      <MenuLink href="#contact" label="Contact" isFooter={isFooter} />
    </NavMenu>
  );
};

export default RecipeMenu;
