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
    { "stepNumber": 1, "text": "Boil water for pasta." },
    { "stepNumber": 2, "text": "Cook guanciale until crispy." },
    { "stepNumber": 3, "text": "Mix eggs and cheese." },
    { "stepNumber": 4, "text": "Cook spaghetti." },
    { "stepNumber": 5, "text": "Combine everything and serve." }
  ]
}'

curl -X POST http://localhost:8080/api/list \
-H "Content-Type: application/json"

curl 'http://localhost:8080/api/get' \
  -H 'Content-Type: application/json' \
  --data-raw '{"id":"7f84f163-e1eb-4c99-a626-7659678b44e4"}'

curl 'http://localhost:8080/api/delete' \
  -H 'Content-Type: application/json' \
  --data-raw '{"id":"7f84f163-e1eb-4c99-a626-7659678b44e4"}'
