FROM golang:1.20.3

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build
EXPOSE 8080
CMD ["./api"]