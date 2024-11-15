FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o pokeneas
EXPOSE 8080
CMD ["./pokeneas"]