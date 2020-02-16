#!/bin/bash

COMMAND=${1:-up}
container_name="mskit"

function execute-docker-compose () {
  docker-compose \
    -f 'docker-compose.yml' \
    $@
}

function stop-docker-compose () {
  execute-docker-compose stop
}

if [ $COMMAND = 'up' ] && [ $# -le 1 ]; then
  execute-docker-compose up -d
  execute-docker-compose exec $container_name /bin/bash
  stop-docker-compose
elif [ $COMMAND = 'bash' ]; then
  execute-docker-compose exec $container_name /bin/bash
else
  execute-docker-compose $@
fi
