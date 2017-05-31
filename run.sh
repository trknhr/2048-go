#!/bin/sh

find ./ -name '*.go' -not -name '*_test.go' | xargs go run
