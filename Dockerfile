FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /unknown-app

# RUNNER
FROM alpine

WORKDIR /

COPY --from=build /unknown-app /unknown-app

EXPOSE 8080

ENTRYPOINT ["/unknown-app"]