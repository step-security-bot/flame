#!/usr/bin/env bash

pushd ..
DOCKER_BUILDKIT=1 docker build -f build/Dockerfile.dev --tag fledge .
popd
