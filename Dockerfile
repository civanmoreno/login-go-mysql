FROM golang:1.23

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod ./
COPY .air.toml ./
COPY . .

RUN go mod download

# Crea el directorio tmp para Air
RUN mkdir -p /app/tmp

CMD ["air", "-c", ".air.toml"]