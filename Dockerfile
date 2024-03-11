FROM golang:1.21.7-alpine
WORKDIR /app
EXPOSE 8080
COPY go.mod go.sum ./
RUN go get .
COPY . .
