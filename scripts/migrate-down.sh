#!/bin/bash

migrate -path database/migration -database "postgres://user:12345678@localhost:5432/db?sslmode=disable" -verbose down