#!/usr/bin/env bash
# Initialize a new challenge

set -euo pipefail

DAY_NO=$1

git checkout -b "day-$DAY_NO"
mkdir -p "$DAY_NO"
pushd "$DAY_NO"
    cat <<EOF > README.md
# Problem #$DAY_NO

https://adventofcode.com/2023/day/$DAY_NO

## Problem statement

...

## Input

https://adventofcode.com/2023/day/$DAY_NO/input

## Requirements

...

## Run

Build and run: \`make all\` or:

1. Build: \`make\`
2. Run:   \`make run\`
EOF
    cat <<EOF > Makefile
SHELL:=/bin/bash

all: build run

build:
    echo "build command"

run:
    echo "run command"
EOF
popd
