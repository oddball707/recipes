curl -X POST http://localhost:8080/api/create \
-H "Content-Type: application/json" \
-d '{
  "name": "Spaghetti Carbonara",
  "ingredients": [
    { "name": "Spaghetti", "quantity": 200, "unit": "g" },
    { "name": "Guanciale", "quantity": 100, "unit": "g" },
    { "name": "Eggs", "quantity": 2},
    { "name": "Pecorino Romano", "quantity": 50, "unit": "g" },
    { "name": "Black Pepper", "quantity": 1, "unit": "tsp" }
  ],
  "instructions": [
    { "step": 1, "description": "Boil water for pasta." },
    { "step": 2, "description": "Cook guanciale until crispy." },
    { "step": 3, "description": "Mix eggs and cheese." },
    { "step": 4, "description": "Cook spaghetti." },
    { "step": 5, "description": "Combine everything and serve." }
  ]
}'

curl -X POST http://localhost:8080/api/list \
-H "Content-Type: application/json" \

