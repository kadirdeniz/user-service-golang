#!/bin/bash

mockgen -package=mock -source=./internal/repository/user.go -destination=./test/mock/user_repository.go 