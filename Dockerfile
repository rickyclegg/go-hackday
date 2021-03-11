FROM golang:1.16.1
WORKDIR /usr/src/app
COPY . .
RUN go build .
EXPOSE 8080
CMD ["./main"]
