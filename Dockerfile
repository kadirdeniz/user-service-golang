FROM golang:1.21-alpine

RUN apk update && apk add make --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 3000

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go get -u -d github.com/golang-migrate/migrate

RUN [ "make" , "migrate-up" ]

ENTRYPOINT [ "make" , "run" ]