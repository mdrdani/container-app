FROM golang:alpine3.16 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /aplikasi

FROM alpine:latest
WORKDIR /
COPY --from=build /aplikasi /aplikasi
EXPOSE 8144
ENTRYPOINT ["/aplikasi"]