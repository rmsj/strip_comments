SHELL := /bin/bash

# ==============================================================================
# test


test:
	go test -v

run:
	go build strip_comments.go && ./strip_comments