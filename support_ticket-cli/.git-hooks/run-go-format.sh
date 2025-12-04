#!/bin/bash
set -e

cd ticket-cli

echo "Running gofumpt..."
gofumpt -w .

echo "Running goimports..."
goimports -w .
