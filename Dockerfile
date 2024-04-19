# syntax=docker/dockerfile:1

FROM golang:1.22-alpine
ENV APP_HOME /go/src/fizzy
WORKDIR "$APP_HOME"

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod verify

COPY ./ .

RUN ls -lart /usr/local/go/src/

RUN go build -o fizzy ./src

EXPOSE 8080

CMD [ "./fizzy" ]