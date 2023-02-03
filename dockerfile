FROM golang:latest

ARG ADDR
ENV ADDR $ADDR
ARG PORT
ENV PORT $PORT
ARG APP_URL
ENV APP_URL $APP_URL
ARG REQUESTMAX
ENV REQUESTMAX $REQUESTMAX
ARG APP_ENV=prod
ENV APP_ENV $APP_ENV

WORKDIR /backend

ADD . ./

RUN go install

RUN go build -a -o ./server

CMD ["./server"]

EXPOSE ${PORT}

HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl --fail http://${ADDR}:${PORT}/status || exit 1
