#!/bin/bash

function start_dev() {
  cp dev.template.env dev.env
  docker-compose -f dev.docker-compose.yaml up --env-file dev.env -d
}

function stop_dev() {
  docker-compose -f dev.docker-compose.yaml down
}

"$@"