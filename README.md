# Go Web App Skeleton

Go Web App Skeleton is a boilerplate for building web applications in Golang. It is built using the Gin Gonic framework and PostgreSQL as the database. It also includes a Dockerfile for containerization. The project is structured in a way that it is easy to add new features and scale the application.

Feel free to fork the project and use it as a starting point for your next web application.

## Technologies used

- [Dependabot](https://github.com/dependabot) - Dependency Management
- [Gin Gonic](https://github.com/gin-gonic/gin) - Golang Web Framework
- [Goose](https://github.com/pressly/goose) - Golang Migrations framework
- [GoDotEnv](https://github.com/joho/godotenv) - Golang .env file parser
- [Golang CI Lint](https://github.com/golangci/golangci-lint) - Golang Linter
- [GoSec](https://github.com/securego/gosec) - Golang Security Scan
- [TailwindCSS](https://tailwindcss.com/docs/installation/play-cdn) - CSS Framework from the CDN

## Features

- Status page at /status
- Simple Ping API at /api/v1/ping
- PostgreSQL Connection
- Security Headers + Host Header Injection Fix
- Static files serving via public folder
- HTML Templates
- Container Compatible
- Error Handling
- Comprehensive Go Linting
- Quick Code Verification Workflow via Github Actions

## Folder structure

- `config` - Configuration as code (e.g. Environment Variables, Gin Gonic Configuration)
- `docs` - Web Application documentation
- `handlers` - Contain the logic for the web application (e.g. API, Error Handling, etc.)
- `middleware` - Middleware is as HTTP Web Server for API routes, static site serving, etc. 
- `migrations/` - Database migrations for the web application
- `models` - Models contain data structures
- `public/` - Static files for the UI
- `public/errors` - Error pages served by the web application
- `public/views` - HTML Templates, mostly used for rendering the UI via API
- `tests/` - GO Unit Tests for the web application

## Architectural Logic

![Architecture](./docs/assets/architecture.png)

## Available Make Commands

- `make build` - Build the application
- `make run` - Run the application
- `make test` - Run the tests
- `make fumpt` - Run the go fmt
- `make linter` - Run the comprehensive [GolangCILint](.golangci.yml) to check the code quality
- `make gosec` - Run the GoSec to check code for vulnerabilities
- `make mod-vendor` - Vendor the dependencies
- `make validate` - Runs `make linter`, `make test` and `make gosec` to validate the code
- `make migrate-create MIGRATION_NAME` - Create a new migration
- `make migrate-up` - Run the migrations
- `make migrate-down` - Rollback the migrations

## Available Paths

- `/` - Home page
- `/status` - Status page
- `API_PATH/ping` - API check endpoint
- `API_PATH/users` - Sample Users API endpoint

**Note!** Replace `API_PATH` with the actual path of the application. By default it's `/api/v1/`

## Managing environment variables

Environment variables are managed via struct in [envConfig Model](./src/model/envConfig.go). 

You can add new environment variables in the struct and configure the defaults in [envConfig Configuration File](./src/config/envConfig.go).

## Getting Started

1. Clone the repository

```bash
git clone git@github.com:KostLinux/example-go-web-app.git
```

2. Change the directory

```bash
cd example-go-web-app
```

3. Configure .env (optional)

```bash
cp .env.example .env && nano .env
```

4. Run the application

```bash
make run || go run main.go
```

5. Visit the application in your browser

Feel free to visit the application at localhost:8000 and move around available paths

## Contribution

Feel free to contribute to the project by creating a pull request. Make sure to follow the [Contribution Guidelines](https://docs.github.com/en/get-started/exploring-projects-on-github/contributing-to-a-project).