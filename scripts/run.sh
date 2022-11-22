#!/bin/bash

# always rebuild image
./scripts/build.sh

# run container
docker run -it \
    -m 1GiB \
    -e "GOMEMLIMIT=30MiB" \
    -e "GOGC=100" \
    -e "GODEBUG=gctrace=1" \
    -p "8080:8080" \
    gogc-demo
