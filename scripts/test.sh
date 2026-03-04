curl -X POST http://localhost:8080/api/create \
-H "Content-Type: application/json" \
-d '{
  "name": "Spaghetti Carbonara",
  "description": "A classic Italian pasta dish made with eggs, cheese, guanciale, and pepper.",
  "ingredients": [
    { "name": "Spaghetti", "quantity": 200, "unit": "g" },
    { "name": "Guanciale", "quantity": 100, "unit": "g" },
    { "name": "Eggs", "quantity": 2},
    { "name": "Pecorino Romano", "quantity": 50, "unit": "g" },
    { "name": "Black Pepper", "quantity": 1, "unit": "tsp" }
  ],
  "instructions": [
    { "step": 1, "text": "Boil water for pasta." },
    { "step": 2, "text": "Cook guanciale until crispy." },
    { "step": 3, "text": "Mix eggs and cheese." },
    { "step": 4, "text": "Cook spaghetti." },
    { "step": 5, "text": "Combine everything and serve." }
  ]
}'

curl -X POST http://localhost:8080/api/list \
-H "Content-Type: application/json" \

curl -X POST http://localhost:8080/api/delete \
-H "Content-Type: application/json" \
-d '{
  "id": "06d3bcf5-81cc-4341-a3da-461502035175"
}'
