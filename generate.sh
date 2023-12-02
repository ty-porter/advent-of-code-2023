#!/bin/bash
set -eu pipefail

PREVDAY=$(ls | grep day | sort -V | tail -n 1 | tr -d '\n' | tail -c 1)
NEXTDAY=$((PREVDAY + 1))
NEXTDAY_DIR="day$NEXTDAY"

echo "Created directory $NEXTDAY_DIR"

mkdir $NEXTDAY_DIR

pushd $NEXTDAY_DIR

# Copy go template
cp ../main.go.template main.go

# Init the mod
go mod init

# Copy blank inputs for test data and actual input
touch part1.txt
touch part2.txt
touch input.txt

popd

echo "AoC 2023 Day $NEXTDAY scaffold created! Good luck!"
