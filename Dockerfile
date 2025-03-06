FROM golang:1.24-alpine3.21

WORKDIR /app
COPY src/* ./
RUN go mod download &&\
    go build -o main *.go
USER 1001
CMD ["/app/main"]

