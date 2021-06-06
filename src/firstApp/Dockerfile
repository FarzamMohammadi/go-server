FROM golang:1.16.3
RUN mkdir -p /app
COPY . /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
CMD ["go","run","/app/main.go"]
