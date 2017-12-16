#!/bin/sh
goimports -w .
go fmt
#go build -race
go test -race
#./pipelines arg0
