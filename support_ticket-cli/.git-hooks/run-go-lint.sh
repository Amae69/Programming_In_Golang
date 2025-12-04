#!/bin/bash
set -e

cd ticket-cli

echo "Running golangci-lint..."
golangci-lint run
