#builder stage
FROM golang:1-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

#run stage
FROM alpine
WORKDIR /app
COPY  --from=builder /app/main .
COPY  --from=builder /app/migrate ./migrate
COPY  app.env .
COPY start.sh .
COPY db/migration ./migration
COPY wait-for.sh .
EXPOSE 8000
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
