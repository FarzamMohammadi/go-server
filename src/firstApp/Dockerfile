FROM golang:1.16.3-alpine
RUN mkdir -p /app
COPY . /app
WORKDIR /app
COPY go.mod ./
RUN go mod download
CMD ["go","run","/app/main.go"]