FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /itk ./cmd/itk

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /itk .

EXPOSE 8080

CMD ["./itk"]