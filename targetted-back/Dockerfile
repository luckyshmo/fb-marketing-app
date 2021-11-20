FROM golang:1.16.2-alpine as builder

ENV GOPATH=/

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

# build go app
RUN go build -o targetted-back ./cmd/main.go

#Build destination container
FROM alpine:latest

ENV GOPATH=/go

# copy bin
COPY --from=builder $GOPATH/targetted-back ./

# copy PG migrations
COPY --from=builder $GOPATH/internal/repository/pg/migrations/*.sql ./migrations/

EXPOSE 8080
