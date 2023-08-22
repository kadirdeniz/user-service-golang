#!/bin/bash

export PATH=$PATH:$(go env GOPATH)/bin

mockgen -package=mock -source=./internal/repository/user.go -destination=./test/mock/user_repository.go 