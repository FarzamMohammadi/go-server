FROM golang:1.16.3-alpine
RUN mkdir -p /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
WORKDIR /app
CMD ["go","run","/app/main.go"]
