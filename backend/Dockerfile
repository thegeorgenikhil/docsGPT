FROM golang:latest 

WORKDIR /app

COPY . . 

RUN go build .

EXPOSE 9000 

CMD ["./docsGPT"]
