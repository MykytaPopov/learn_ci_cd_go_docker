FROM golang:1.17-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go clean --modecache
RUN go build -o main .

CMD [ "/app/main" ]
