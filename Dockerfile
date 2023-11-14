# syntax=docker/dockerfile:1
FROM golang:1.21.1-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash make gcc gettext musl-dev

# dependencies
COPY go.mod go.sum ./
RUN go mod download

# build
COPY ./ ./
RUN go build -o ./bin/app cmd/app/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app ./
COPY --from=builder /usr/local/src/.env ./
COPY --from=builder /usr/local/src/migrations ./migrations

CMD ["./app"]