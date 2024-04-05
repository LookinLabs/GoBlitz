# Go Web App Skeleton

> :warning: **NOTE:** This project is currently under development and does not have a stable version yet. Please use with caution.

The Go Web App Skeleton is a boilerplate for building web applications in Golang, designed to mainly run on containers but can also be set up in a VM. The project is structured in a way that makes it easy to add new business logic features. Its architecture is similar to layered architecture with some MVC components meshed in.

Web application built on top of the Gin Gonic framework are very fast and efficient. You can look for benchmarks [here](#benchmarks).

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

- Simple Ping API at /api/v1/ping
- Status page at /status
- PostgreSQL Connection
- Security Headers
- SSRF protection via Host Header Validation
- Static site serving at root path (/)
- HTML Templates
- Error Handling
- Container Compatible
- Comprehensive and quite strict quality scan
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
- `repository`- Repository is a layer that connects the application to external services like databases, cache servers, etc.
- `tests/` - GO Unit Tests for the web application

## Architecture

The architecture of the web application is layered, with some MVC components integrated.

The middleware layer handles routing and serves the static site from the `public` folder. It interacts with the handler layer to process requests from clients and generate appropriate responses.

The repository layer manages connections to external services such as databases and cache servers. It utilizes environment variables from the model component to establish connections with these external services.

The `handler` component plays a crucial role in processing requests received from the middleware layer. It enriches these requests with necessary data for generating responses, using the model component in conjunction with the repository layer. Additionally, it manages errors and sends appropriate responses back to the middleware layer.

The `model` component contains data structures for handlers and environment variable values used by the repository, middleware, and handler components.

The `config` component facilitates configuring the application through code. It stores configuration settings for the repository, middleware, and handler components. The config component is combined with the model to establish data structures that promote code reusability across all layers and components.


![Architecture](./docs/assets/architecture.png)

## Benchmarks

Benchmarks are done via ApacheBench. We're comparing here the performance of default web application written via Go and Laravel against /api/v1/ping endpoint which contains only JSON message: `{"message": "pong"}`.

### Go Web App Skeleton (Built on Gin-Gonic)

1. 100 requests and 10 concurrenct connections

```
ab -n 100 -c 10 -k http://localhost:8000/api/v1/ping
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/v1/ping
Document Length:        571 bytes

Concurrency Level:      10
Time taken for tests:   0.012 seconds
Complete requests:      100
Failed requests:        0
Keep-Alive requests:    100
Total transferred:      123300 bytes
HTML transferred:       57100 bytes
Requests per second:    8602.15 [#/sec] (mean)
Time per request:       1.162 [ms] (mean)
Time per request:       0.116 [ms] (mean, across all concurrent requests)
Transfer rate:          10357.86 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:     1    1   0.3      1       3
Waiting:        1    1   0.3      1       2
Total:          1    1   0.4      1       3

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      2
  99%      3
 100%      3 (longest request)
```

2. 100 requests and 100 concurrenct connections

```
ab -n 100 -c 100 -k http://localhost:8000/api/v1/ping
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/v1/ping
Document Length:        571 bytes

Concurrency Level:      100
Time taken for tests:   0.013 seconds
Complete requests:      100
Failed requests:        0
Keep-Alive requests:    100
Total transferred:      123300 bytes
HTML transferred:       57100 bytes
Requests per second:    7521.06 [#/sec] (mean)
Time per request:       13.296 [ms] (mean)
Time per request:       0.133 [ms] (mean, across all concurrent requests)
Transfer rate:          9056.12 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   0.5      3       3
Processing:     4    5   0.4      5       6
Waiting:        2    5   0.5      5       6
Total:          4    8   0.6      8       9

Percentage of the requests served within a certain time (ms)
  50%      8
  66%      8
  75%      8
  80%      8
  90%      8
  95%      8
  98%      9
  99%      9
 100%      9 (longest request)
```

3. 1000 requests and 100 concurrenct connections

```
ab -n 1000 -c 100 -k http://localhost:8000/api/v1/ping
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/v1/ping
Document Length:        571 bytes

Concurrency Level:      100
Time taken for tests:   0.057 seconds
Complete requests:      1000
Failed requests:        0
Keep-Alive requests:    1000
Total transferred:      1233000 bytes
HTML transferred:       571000 bytes
Requests per second:    17653.19 [#/sec] (mean)
Time per request:       5.665 [ms] (mean)
Time per request:       0.057 [ms] (mean, across all concurrent requests)
Transfer rate:          21256.23 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.9      0       3
Processing:     1    5   1.5      5      12
Waiting:        1    5   1.5      5      12
Total:          1    5   1.9      5      12

Percentage of the requests served within a certain time (ms)
  50%      5
  66%      6
  75%      6
  80%      6
  90%      8
  95%      9
  98%      9
  99%     10
 100%     12 (longest request)
```

### Laravel (default)

1. 100 requests and 10 concurrenct connections

```
ab -n 100 -c 10 -k http://localhost:80/api/v1/ping
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        nginx
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/ping
Document Length:        33822 bytes

Concurrency Level:      10
Time taken for tests:   6.029 seconds
Complete requests:      100
Failed requests:        0
Keep-Alive requests:    0
Total transferred:      3494800 bytes
HTML transferred:       3382200 bytes
Requests per second:    16.59 [#/sec] (mean)
Time per request:       602.927 [ms] (mean)
Time per request:       60.293 [ms] (mean, across all concurrent requests)
Transfer rate:          566.05 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:   219  557 100.1    546     836
Waiting:      218  556 100.2    546     835
Total:        219  557 100.1    547     836

Percentage of the requests served within a certain time (ms)
  50%    547
  66%    572
  75%    580
  80%    617
  90%    674
  95%    750
  98%    807
  99%    836
 100%    836 (longest request)
```
2. 100 requests and 100 concurrenct connections

```
ab -n 100 -c 100 -k http://localhost:80/api/v1/ping
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        nginx
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/ping
Document Length:        33822 bytes

Concurrency Level:      100
Time taken for tests:   5.932 seconds
Complete requests:      100
Failed requests:        0
Keep-Alive requests:    0
Total transferred:      3494800 bytes
HTML transferred:       3382200 bytes
Requests per second:    16.86 [#/sec] (mean)
Time per request:       5932.158 [ms] (mean)
Time per request:       59.322 [ms] (mean, across all concurrent requests)
Transfer rate:          575.32 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    3   0.6      4       4
Processing:   229 3007 1586.7   2938    5651
Waiting:      229 3006 1586.7   2938    5650
Total:        232 3010 1587.2   2942    5655
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%   2942
  66%   3986
  75%   4279
  80%   4534
  90%   5061
  95%   5422
  98%   5654
  99%   5655
 100%   5655 (longest request)
```

3. 1000 requests and 100 concurrenct connections

```
ab -n 1000 -c 100 -k http://localhost:80/api/v1/ping
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        nginx
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/v1/ping
Document Length:        33822 bytes

Concurrency Level:      100
Time taken for tests:   53.601 seconds
Complete requests:      1000
Failed requests:        0
Keep-Alive requests:    0
Total transferred:      34948000 bytes
HTML transferred:       33822000 bytes
Requests per second:    18.66 [#/sec] (mean)
Time per request:       5360.134 [ms] (mean)
Time per request:       53.601 [ms] (mean, across all concurrent requests)
Transfer rate:          636.72 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.8      0       4
Processing:   218 5083 877.5   5315    5709
Waiting:      217 5080 877.2   5314    5709
Total:        222 5083 876.8   5315    5711

Percentage of the requests served within a certain time (ms)
  50%   5315
  66%   5343
  75%   5383
  80%   5412
  90%   5451
  95%   5459
  98%   5474
  99%   5512
 100%   5711 (longest request)
```

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