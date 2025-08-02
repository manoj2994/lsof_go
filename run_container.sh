#!/bin/bash

docker run -it --rm \
  --name lsof_go \
  --pid=host \
  -v "$PWD":/app \
  -w /app \
  lsof_go