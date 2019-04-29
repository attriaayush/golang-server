FROM golang:alpine

ADD ./src go/src/app
WORKDIR /go/src/app

COPY ./main.go /usr/src/app/

ENV PORT=3001

CMD ["go", "run", "main.go"]