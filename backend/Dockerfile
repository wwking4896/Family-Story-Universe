FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /fairy-castle-api ./cmd/api

FROM alpine:3.20
RUN addgroup -S app && adduser -S app -G app
USER app
COPY --from=build /fairy-castle-api /fairy-castle-api
EXPOSE 8080
ENTRYPOINT ["/fairy-castle-api"]
