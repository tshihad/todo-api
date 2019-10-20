FROM golang:latest

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go test ./... && go build

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/todo .
EXPOSE 8080
CMD ./todo 