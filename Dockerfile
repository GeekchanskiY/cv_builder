FROM golang:1.22-alpine

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /cv_builder /app/cmd/main.go 

EXPOSE 8080

CMD ["/cv_builder"]