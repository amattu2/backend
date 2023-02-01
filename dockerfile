FROM golang:latest

WORKDIR /backend

ADD . ./

RUN go install

RUN go build -a -o ./server

CMD ["./server"]

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl --fail http://localhost:8080/status || exit 1
