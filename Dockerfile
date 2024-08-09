FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY * ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /cv_builder

EXPOSE 8080

CMD ["/cv_builder"]