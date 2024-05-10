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
nano .env
```

4. Run the application

```bash
make run || go run main.go
```

5. Visit the application in your browser

Feel free to visit the application at `localhost:8000` and move around available paths

## Advanced

### Running the application with Live Reload (Air)

1. Install Air

```bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
mv ./bin/air /usr/local/bin/air
```

2. Run the application with Air

```bash
make
```

### Running the application via Cognito

#### Prerequisites

- [AWS Account](https://aws.amazon.com/)
- [Configured AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)

#### Setting Up Cognito

Cognito can be set up via terraform code under aws/terraform directory. This code sets up the Cognito User Pool and Identity Pool.

1. Change the directory

```bash
cd aws/terraform
```

2. Initialize the terraform

```bash
terraform init
```

3. Plan the terraform

```bash
terraform plan
```

4. Apply the terraform

```bash
terraform apply
```

#### Configure Cognito User

1. Configure the Cognito User

You can sign up for a new user via Cognito Hosted UI.

Go to the `Cognito` tab in the AWS Console and click on the User Pools. Click on the User Pool you created and go to the `App Integration`.

By default there will be `GoBlitzClient` application client. Click on that client and scroll down till you see "Hosted UI" section. Press `View Hosted UI` and sign up for a new user.

2. Configure the .env

Configure Cognito Environment variables. The `AWS_COGNITO_USER_POOL_ID` and `AWS_COGNITO_CLIENT_ID` can be found in terminal after running `terraform apply` under `Changes on Outputs:`.

```
Changes to Outputs:
  client_id    = "ah37r9t409rjotnm8h575edio"
  user_pool_id = "us-east-1_5laVcI9aL"
```

```bash
AWS_COGNITO_USER_POOL_ID=user_pool_id
AWS_COGNITO_CLIENT_ID=client_id
```

`AWS_COGNITO_TOKEN_URL` is the URL to get the token from Cognito. It's in the format of `https://<user-pool-id>.auth.<region>.amazoncognito.com/oauth2/token`.

By default it's `https://goblitz.auth.us-east-1.amazoncognito.com/oauth2/token`

The `AWS_COGNITO_API_USER_EMAIL` and `AWS_COGNITO_API_USER_PASSWORD` are the email and password of the user you signed up with.

You can leave AWS_COGNITO_JWT_TOKEN empty. It will be filled after third step.

3. Get the JWT Token

You can get the JWT Token by running the following command:

```bash
python3 bin/fetch_jwt.py
```

4. Configure AWS_COGNITO_JWT_TOKEN

Configure the JWT Token in the .env file.

5. Run the application

```bash
make
```

6. Visit some API Path in your browser

## Docs

You can find the documentation page [here](https://kostlinux.github.io/GoBlitz-Docs/).

## Contribution

Feel free to contribute to the project by creating a pull request.

Make sure to follow the [Contribution Guidelines](https://docs.github.com/en/get-started/exploring-projects-on-github/contributing-to-a-project).