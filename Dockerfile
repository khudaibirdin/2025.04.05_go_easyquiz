FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o web_app cmd/app/main.go

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/service /app/
COPY --from=builder /app/configs /app/configs

EXPOSE 9000

CMD ["./web_app"]
