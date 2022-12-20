#!/usr/bin/env bash

#go install github.com/swaggo/swag/cmd/swag@latest

mkdir -p assets

swag init -o assets/docs
