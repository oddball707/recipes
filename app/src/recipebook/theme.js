import { createTheme } from '@mui/material/styles';

// Elegant theme matching spiceandstiletto.com
// Red, white, and black color scheme
const theme = createTheme({
  palette: {
    primary: {
      main: '#C41E3A',      // Deep red/crimson
      light: '#D63A54',     // Lighter red
      dark: '#8B1428',      // Dark red
      contrastText: '#fff',
    },
    secondary: {
      main: '#000000',      // Black
      light: '#333333',     // Dark gray
      dark: '#000000',
      contrastText: '#fff',
    },
    background: {
      default: '#FFFFFF',   // White
      paper: '#FFFFFF',
    },
    text: {
      primary: '#000000',   // Black
      secondary: '#333333', // Dark gray
    },
    divider: '#E0E0E0',
  },
  typography: {
    fontFamily: '"Lato", "Segoe UI", "Roboto", "Helvetica Neue", sans-serif',
    h1: {
      fontSize: '3rem',
      fontWeight: 700,
      letterSpacing: '-1px',
      marginBottom: '1rem',
      color: '#000000',
    },
    h2: {
      fontSize: '2rem',
      fontWeight: 600,
      marginBottom: '1.5rem',
      color: '#000000',
    },
    h3: {
      fontSize: '1.5rem',
      fontWeight: 600,
      color: '#000000',
    },
    h4: {
      fontSize: '1.25rem',
      fontWeight: 600,
      color: '#000000',
    },
    h6: {
      fontSize: '1rem',
      fontWeight: 600,
      color: '#000000',
    },
    subtitle1: {
      fontSize: '1.1rem',
      fontWeight: 300,
      lineHeight: 1.6,
      color: '#333333',
    },
    body1: {
      fontSize: '1rem',
      lineHeight: 1.7,
      color: '#000000',
    },
    body2: {
      fontSize: '0.95rem',
      lineHeight: 1.6,
      color: '#333333',
    },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: 'uppercase',
          fontWeight: 600,
          letterSpacing: '0.5px',
          borderRadius: '2px',
        },
        contained: {
          boxShadow: 'none',
          backgroundColor: '#C41E3A',
          '&:hover': {
            backgroundColor: '#8B1428',
            boxShadow: '0 2px 8px rgba(196, 30, 58, 0.2)',
          },
        },
        outlined: {
          borderColor: '#C41E3A',
          color: '#C41E3A',
          '&:hover': {
            backgroundColor: 'rgba(196, 30, 58, 0.04)',
            borderColor: '#8B1428',
            color: '#8B1428',
          },
        },
      },
    },
    MuiMenu: {
      styleOverrides: {
        paper: {
          fontFamily: '"Montserrat", "Segoe UI", "Roboto", "Helvetica Neue", sans-serif',
          fontSize: '10px',
          letterSpacing: '1.2px',
        },
      },
    },
    MuiMenuItem: {
      styleOverrides: {
        root: {
          fontFamily: '"Montserrat", "Segoe UI", "Roboto", "Helvetica Neue", sans-serif',
          fontSize: '10px !important',
          letterSpacing: '1.2px',
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          boxShadow: '0 2px 8px rgba(0, 0, 0, 0.08)',
          border: '1px solid #E0E0E0',
          '&:hover': {
            boxShadow: '0 8px 16px rgba(196, 30, 58, 0.1)',
          },
          transition: 'all 0.3s ease',
        },
      },
    },
    MuiAppBar: {
      styleOverrides: {
        root: {
          backgroundColor: '#FFFFFF',
          color: '#000000',
          boxShadow: '0 2px 4px rgba(0, 0, 0, 0.08)',
          borderBottom: '1px solid #E0E0E0',
          textTransform: 'uppercase',
          letterSpacing: '1.2px',
        },
      },
    },
  },
});

export default theme;
