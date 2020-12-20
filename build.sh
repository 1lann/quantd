#!/bin/bash
go get github.com/tinylib/msgp
go get github.com/markbates/pkger/cmd/pkger
go generate ./...
go build -ldflags '-s -w -extldflags "-fno-PIC -static"' -tags 'osusergo netgo static_build' -o "assets/quantd_$(go env GOOS)_$(go env GOARCH)" ./cmd/quantd
pkger -include /assets -o assets
