#!/bin/bash

# always rebuild image
./scripts/build.sh

# run container
docker run -it \
    -m 4GiB \
    --env-file .env \
    -p "8080:8080" \
    gogc-demo
