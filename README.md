# Basic-Auth-JWT-Server with Golang + Docker

Basic-auth-server with Golang + Docker. This server can return JWT.

## Quick Start

```bash
$ make up
```

try basic auth

```bash
$ curl --user test1@mail.com:pass1 http://localhost:8081/api/v1/auth/gettoken
```

## make command

```bash
$ make help
all                  go test & go build
build                build go binary
test                 go test
clean                remove go bainary
run                  run go bainary
initdb               create database and insert first data
up                   docker-compose up
down                 docker-compose down
execapi              exec api-server container
execdb               exec db-server container
help                 Display this help screen
```