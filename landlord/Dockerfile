FROM golang:1.21-alpine
WORKDIR /app
COPY ../application ./
RUN go mod tidy
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
