# Recipes
Free Recipe Book Service

## Development
See Makefile for examples
Must have Go and Docker installed locally

`make build` to build

`make start` to start local dev server

## Configuration

The application uses Viper to handle configuration management and automatically detects the runtime environment. All configuration is validated on startup - no unsafe defaults.

### Local Development
- **Source**: Reads from `config.local.json` file first

### Docker Container
- **Source**: Reads exclusively from environment variables set in `docker-compose.yml`

### AWS Lambda
- **Source**: Reads exclusively from environment variables set via Lambda configuration
