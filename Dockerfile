FROM golang:1.21-alpine as base

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o http-server cmd/http/main.go

RUN apk -o no-cache add ca-certificates

FROM scratch

COPY --from=base ["build/http-server", "/http-server"]

COPY --from=base ["/etc/ssl/certs/ca-certificates.crt", "/etc/ssl/certs/"]

ARG GO_ENV="deploy"
ARG PGHOST
ARG PGPORT
ARG PGUSER
ARG PGPASSWORD
ARG PGDATABASE
ARG DATABASE_URL
ARG PORT
ARG JWT_SECRET_KEY

ENV GO_ENV=${GO_ENV}
ENV PGHOST=${PGHOST}
ENV PGPORT=${PGPORT}
ENV PGUSER=${PGUSER}
ENV PGPASSWORD=${PGPASSWORD}
ENV PGDATABASE=${PGDATABASE}
ENV DATABASE_URL=${DATABASE_URL}
ENV PORT=${PORT}
ENV JWT_SECRET_KEY=${JWT_SECRET_KEY}

CMD ["/http-server"]

