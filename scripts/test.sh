#!/bin/bash

go test \
    -v \
    -bench=. \
    -benchmem \
    -benchtime=10s \
    -memprofile=mem.out \
    -cpuprofile=cpu.out \
    -gcflags="-m -m" ./...
