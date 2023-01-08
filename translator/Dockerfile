FROM golang:1-alpine3.15
WORKDIR /app
COPY . .
RUN go build -o main main.go
EXPOSE 8090
CMD [ "/main/app" ]