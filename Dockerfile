FROM golang:1.26.5-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /server .

FROM scratch
COPY --from=build /server /server
USER 65532:65532
ENV GIN_MODE=release PORT=3000
EXPOSE 3000
ENTRYPOINT ["/server"]
