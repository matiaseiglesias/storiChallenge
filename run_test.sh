#!/bin/sh
echo Running tests
echo Running unit tests
go test -v ./internal/services_test
echo Running integration tests
go test -cover