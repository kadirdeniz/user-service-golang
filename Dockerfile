FROM golang:1.18-alpine

RUN apk update && apk add make --no-cache git
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY . .

COPY ./go.* .

RUN go mod download

COPY . .

EXPOSE 3000

ENTRYPOINT [ "make","run" ]