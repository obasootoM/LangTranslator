#builder stage
FROM golang:1-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#run stage
FROM alpine:3.15
WORKDIR /app
COPY  --from=builder /app/main .
COPY  app.env .
EXPOSE 8000
CMD [ "/main/app" ]