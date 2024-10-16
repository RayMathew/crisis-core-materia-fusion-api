# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)" 
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## install golangci-lint first. go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
lint:
	golangci-lint run

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the api application
.PHONY: build
build:
	go build -o=/tmp/bin/api ./api
	
## run: run the api application
.PHONY: run
run: build
	/tmp/bin/api

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/api" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

# ==================================================================================== #
# DEBUG
# ==================================================================================== #

# The Dockerfile is currently written for building on the Google Cloud. 
# If needed in local, add this line COPY .env /app/.env, as the env values in Cloud are fetched using the Cloud's env variables manager.
# Then use below commands.

## build docker image
dbuild:
	docker build -t materia-fusion-api .

## create and run container
drun:
	docker run --name my-materia-api -p 4444:4444 materia-fusion-api

## stop container
dstop:
	docker stop my-materia-api

## start container
dstart:
	docker start my-materia-api

