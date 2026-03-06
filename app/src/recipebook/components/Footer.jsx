import React from 'react';
import {
  Box,
  Container,
  Typography,
  Stack,
  Link,
  Divider,
} from '@mui/material';
import { styled } from '@mui/material/styles';
import InstagramIcon from '@mui/icons-material/Instagram';
import PinterestIcon from '@mui/icons-material/Pinterest';
import EmailIcon from '@mui/icons-material/Email';
import { RECIPE_CATEGORIES } from './RecipeMenu';
import RecipeMenu from './RecipeMenu';

const FooterContainer = styled(Box)(({ theme }) => ({
  backgroundColor: '#C41E3A',
  color: '#fff',
  marginTop: theme.spacing(8),
  paddingTop: theme.spacing(4),
  paddingBottom: theme.spacing(3),
  fontSize: '12px',
}));

const FooterLink = styled(Link)(({ theme }) => ({
  color: 'rgba(255, 255, 255, 0.85)',
  textDecoration: 'none',
  fontSize: '12px',
  transition: 'color 0.3s ease',
  '&:hover': {
    color: '#FFFFFF',
  },
}));

const Footer = () => {
  return (
    <FooterContainer>
      <Container maxWidth="lg">
        {/* Main Footer Content */}
        <Stack spacing={3} sx={{ mb: 3, alignItems: 'center', textAlign: 'center' }}>
          {/* Navigation Menu */}
          <Box>
            <RecipeMenu isFooter={true} />
          </Box>

          {/* Follow Us */}
          <Box>
            <Stack direction="row" spacing={1} sx={{ justifyContent: 'center' }}>
              <Link
                href="https://www.instagram.com/spiceandstiletto/"
                sx={{
                  color: 'rgba(255, 255, 255, 0.85)',
                  display: 'flex',
                  alignItems: 'center',
                  transition: 'color 0.3s ease',
                  '&:hover': {
                    color: '#FFFFFF',
                  },
                }}
              >
                <InstagramIcon fontSize="small" />
              </Link>
              <Link
                href="https://www.pinterest.com/spiceandstiletto/"
                sx={{
                  color: 'rgba(255, 255, 255, 0.85)',
                  display: 'flex',
                  alignItems: 'center',
                  transition: 'color 0.3s ease',
                  '&:hover': {
                    color: '#FFFFFF',
                  },
                }}
              >
                <PinterestIcon fontSize="small" />
              </Link>
              <Link
                href="mailto:hello@example.com"
                sx={{
                  color: 'rgba(255, 255, 255, 0.85)',
                  display: 'flex',
                  alignItems: 'center',
                  transition: 'color 0.3s ease',
                  '&:hover': {
                    color: '#FFFFFF',
                  },
                }}
              >
                <EmailIcon fontSize="small" />
              </Link>
            </Stack>
          </Box>
        </Stack>

        <Divider sx={{ backgroundColor: 'rgba(255, 255, 255, 0.2)', my: 2 }} />

        {/* Copyright */}
        <Typography variant="caption" sx={{ opacity: 0.8 }}>
          © {new Date().getFullYear()} Spice & Stilleto. A fabulous recipe collection.
        </Typography>
      </Container>
    </FooterContainer>
  );
};

export default Footer;
