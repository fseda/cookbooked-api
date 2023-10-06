FROM golang:1.21-alpine as base

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go install github.co/swaggo/swag/cmd/swag@latest

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o http-server cmd/http/main.go

FROM scratch

COPY --from=base ["build/http-server", "/http-server"]

ENV GO_ENV=production

ARG PGHOST
ARG PGPORT
ARG PGUSER
ARG PGPASSWORD
ARG PGDATABASE
ARG DATABASE_URL
ARG SERVER_PORT
ARG JWT_SECRET_KEY

CMD ["/http-server"]

