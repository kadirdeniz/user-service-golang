FROM golang:1.18-alpine

RUN apk update && apk add make --no-cache git
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY . .

COPY ./go.* .

RUN go mod download

COPY . .

EXPOSE 3000

RUN go get -u -d github.com/golang-migrate/migrate

RUN ["make", "migrate-up"]

ENTRYPOINT [ "make","run" ]