#!/bin/bash

mockgen -package=mock -source=./internal/user/repository.go -destination=./test/mock/user_repository.go 