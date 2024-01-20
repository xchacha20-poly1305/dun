#!/bin/bash

set -e

git submodule update --remote

go mod tidy
