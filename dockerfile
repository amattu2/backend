FROM golang:latest

ARG ADDR
ENV ADDR $ADDR
ARG PORT
ENV PORT $PORT
ARG REQUESTMAX
ENV REQUESTMAX $REQUESTMAX

WORKDIR /backend

ADD . ./

RUN go install

RUN go build -a -o ./server

CMD ["./server"]

EXPOSE ${PORT}

HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl --fail http://${ADDR}:${PORT}/status || exit 1
