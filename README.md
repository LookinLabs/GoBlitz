# Example Go Web service

Example Go Web Service that serves static files and exposes a simple API path.

## Features

- Status page at /status
- PostgreSQL Connection
- Security Headers + Host Header Injection Fix
- Static files serving via public folder
- HTML Templates

## Running tests

Running tests is simple, just use your .env file you have in your project and run the following command:

```bash
env $(cat .env | xargs) go test -v ./tests/ -p 32
```