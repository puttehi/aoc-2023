SHELL:=/bin/bash

all: build run

setup:
	poetry install

build: run # No-op, uses notebook

run: setup
	poetry run jupyter notebook
