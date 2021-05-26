FROM golang:1.16.2-alpine as builder

ENV GOPATH=/

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

# RUN go test -v go/pkg/...

# build go app
RUN go build -o targetted-back ./cmd/main.go

#Build destination container
FROM alpine:latest

ENV GOPATH=/go

# copy bin
COPY --from=builder $GOPATH/targetted-back $GOPATH/server/server.crt $GOPATH/server/server.key ./

# copy PG migrations
COPY --from=builder $GOPATH/pkg/repository/pg/migrations/*.sql ./migrations/

EXPOSE 8080
