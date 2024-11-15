FROM golang:1.23.2
WORKDIR /app
COPY . .
RUN go build .
EXPOSE 80
CMD ["./pokeneas"]