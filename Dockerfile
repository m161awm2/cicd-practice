FROM golang:1.26.5-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test ./... && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /app


FROM alpine:3.23

RUN addgroup -S app && adduser -S app -G app
COPY --from=build /app /usr/local/bin/app

USER app
ENV GIN_MODE=release
EXPOSE 3000

CMD ["/usr/local/bin/app"]
