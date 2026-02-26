module recipes/lambda

go 1.25.0

require (
	github.com/aws/aws-lambda-go v1.49.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.16.2
	github.com/oddball707/recipes v0.1.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.8.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/oddball707/recipes/handler => ../handler
	github.com/oddball707/recipes/server => ../server
)
