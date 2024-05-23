#!/bin/bash
go test $(go list ./{controllers,repos,services}/... | grep -v mock) -tags integration -coverprofile /tmp/cover.out
go tool cover -html=/tmp/cover.out
rm /tmp/cover.out
