FROM golang:1.26.1

WORKDIR /app

# Copia apenas go.mod e go.sum primeiro (melhor cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do projeto
COPY . .

# Build
RUN go build -o main ./cmd/main.go

EXPOSE 8000

CMD ["./main"]