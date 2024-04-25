# Use the official Golang image as the base image
FROM golang:1.20

WORKDIR /app

COPY . ./

RUN go build -o univerte ./cmd/univerte

EXPOSE 8080

CMD ["./univerte"]
