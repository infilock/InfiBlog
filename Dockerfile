## Build the application from source
FROM golang:1.21.3 AS build-infiblog

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=1 GOOS=linux go build -o /infiblog

## Run the tests in the container
FROM build-infiblog AS test-infiblog

RUN go test -v ./...

## Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-infiblog /infiblog /

EXPOSE 8080

ENTRYPOINT ["/infiblog"]