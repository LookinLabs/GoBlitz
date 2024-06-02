# GoBlitz

> :warning: **NOTE:** This project is currently under development and does not have a stable version yet. Please use with caution.

GoBlitz is a powerful framework built on top of Gin Gonic, designed to help develop production ready web applications in Golang. GoBlitz creates an abstraction layer for the Gin-Gonic framework and provides a structured way to build web applications. It is designed to be fast, secure, and easy to use. The framework is built with security in mind and provides a set of security features to protect web applications from common web vulnerabilities.

Feel free to fork this repository as website boilerplate for your next project.

## Technologies used

- [Dependabot](https://github.com/dependabot) - Dependency Management
- [Gin Gonic](https://github.com/gin-gonic/gin) - Golang Web Framework
- [Goose](https://github.com/pressly/goose) - Golang Migrations framework
- [GoDotEnv](https://github.com/joho/godotenv) - Golang .env file parser
- [Golang CI Lint](https://github.com/golangci/golangci-lint) - Golang Linter
- [GoSec](https://github.com/securego/gosec) - Golang Security Scan
- [TailwindCSS](https://tailwindcss.com/docs/installation/play-cdn) - CSS Framework from the CDN
- [Air](https://github.com/cosmtrek/air) - Live reload for Go applications

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

## Getting Started

**Note!** Before getting started, make sure you have `Goose`, `Air`, `Golang CI Lint`, `Gosec` and `Go` installed on your host

1. Clone the repository

```bash
git clone git@github.com:KostLinux/GoBlitz.git my-web-application
```

2. Change the directory

```bash
cd my-web-application
```

3. Configure .env

```bash
cp .env.example .env
```

4. Setup the database

```bash
docker-compose up -d db
```

5. Run the migrations

```bash
make migrate-up
```

6. Generate `STATUSPAGE_API_KEY`

```
go run bin/api/generate_api_key.go
```

7. Paste the generated key into the `.env` file into the `STATUSPAGE_API_KEY` variable

8. Run the application

```bash
make run
```

9. Visit the application in your browser

Feel free to visit the application at `localhost:8000` and move around available paths

## Running the application with Live Reload (Air)

1. Install Air

```bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
mv ./bin/air /usr/local/bin/air
```

2. Run the application with Air

```bash
make
```

## Docs

You can find the documentation page [here](https://kostlinux.github.io/GoBlitz-Docs/).

## Contribution

Feel free to contribute to the project by creating a pull request.

Make sure to follow the [Contribution Guidelines](https://docs.github.com/en/get-started/exploring-projects-on-github/contributing-to-a-project).