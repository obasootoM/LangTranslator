FROM golang:1-alpine3.15
WORKDIR /app
COPY . .
RUN go build -o main main.go
EXPOSE 9000
CMD [ "/main/app" ]