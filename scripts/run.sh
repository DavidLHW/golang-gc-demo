#!/bin/bash

# always rebuild image
./scripts/build.sh

# run container
docker run -it \
    -m 1GiB \
    --env-file .env \
    -p "8080:8080" \
    gogc-demo
