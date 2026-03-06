# Recipe App Frontend

A sophisticated, elegant recipe application inspired by **spiceandstiletto.com**, built with React and Material-UI.

## Features

✨ **Modern Design**
- Elegant earth-tone color palette (browns, tans, creams)
- Responsive grid layout
- Smooth animations and transitions
- Professional typography with Lato font

📖 **Recipe Management**
- Browse all recipes in a beautiful grid
- View detailed recipe information
- Search recipes by name or description
- Ingredient lists with quantities and units
- Step-by-step cooking instructions

🔄 **API Integration**
- Connected to local backend API (`http://localhost:8080/api`)
- Real-time recipe fetching and filtering
- CRUD operations support

## Project Structure

```
src/recipebook/
├── App.jsx                 # Main app component
├── theme.js               # Material-UI theme (colors, typography)
├── index.js               # Exports
├── components/
│   ├── Navigation.jsx     # Header with hero section & search
│   ├── RecipeCard.jsx     # Recipe card component
│   ├── RecipeDetail.jsx   # Full recipe view
│   └── Footer.jsx         # Footer with links
├── pages/
│   └── HomePage.jsx       # Main recipe listing page
└── services/
    └── recipeApi.js       # API client using Axios
```

## Getting Started

### Prerequisites
- Node.js 16+
- Backend API running on `http://localhost:8080`

### Installation

```bash
cd app
npm install
```

### Development

```bash
npm run dev
```

The app will start on `http://localhost:5173` (Vite default).

### Building

```bash
npm run build
```

Builds the app for production to the `dist` folder.

## API Endpoints

The frontend expects the following REST API:

- `POST /api/list` - Get all recipes
- `POST /api/get` - Get recipe by ID
- `POST /api/create` - Create new recipe
- `POST /api/update` - Update recipe
- `POST /api/delete` - Delete recipe

## Recipe Data Structure

```json
{
  "ID": "uuid-string",
  "Name": "Recipe Name",
  "Description": "Brief description",
  "Ingredients": [
    {
      "Name": "Ingredient name",
      "Quantity": 1.5,
      "Unit": {
        "Name": "cup",
        "Abbreviation": "c"
      }
    }
  ],
  "Instructions": [
    {
      "Instruction": "Step description"
    }
  ]
}
```

## Design Philosophy

The app follows the elegant, sophisticated aesthetic of **spiceandstiletto.com**:

- **Color Palette**: Saddle brown (#8B4513), soft tan (#D4A574), cream backgrounds (#FAFAF8)
- **Typography**: Clean Lato font with strong hierarchy
- **Spacing**: Generous whitespace for a magazine-like feel
- **Interactions**: Subtle shadows and smooth transitions on hover
- **Layout**: Responsive grid that adapts to mobile, tablet, and desktop

## Dependencies

- **@mui/material** - Component library
- **@emotion/react & @emotion/styled** - Styling
- **axios** - HTTP client
- **react** - UI framework
- **react-dom** - React rendering

## Customization

### Changing Colors

Edit `src/recipebook/theme.js`:

```javascript
palette: {
  primary: {
    main: '#8B4513',      // Primary brown
  },
  secondary: {
    main: '#D4A574',      // Accent tan
  },
}
```

### Changing API URL

Edit `src/recipebook/services/recipeApi.js`:

```javascript
const API_BASE_URL = 'http://localhost:8080/api';
```

## Future Enhancements

- [ ] Recipe creation/editing forms
- [ ] Image upload for recipes
- [ ] Category and tag filtering
- [ ] User ratings and favorites
- [ ] Recipe sharing and printing
- [ ] Mobile app version
- [ ] User authentication

## License

MIT - Feel free to use this as a template for your own recipe app!
