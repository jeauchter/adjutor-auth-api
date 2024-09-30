#!/bin/sh

# Download Go module dependencies
go mod download

# Start Air for hot-reloading
air -c .air.toml