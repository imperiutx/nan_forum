# base go image
FROM golang:1.18-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o main .

RUN chmod +x /app/main

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/main /app
COPY app.env .
COPY start.sh /app/start.sh
COPY wait-for.sh /app/wait-for.sh
COPY db/migration ./db/migration

EXPOSE 8000

CMD [ "/app/main" ]

ENTRYPOINT [ "/app/start.sh" ]