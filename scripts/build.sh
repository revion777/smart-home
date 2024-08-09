#!/bin/bash
set -e

echo "Building project..."
go build -o bin/smart-home ./...
