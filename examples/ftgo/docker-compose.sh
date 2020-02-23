#!/bin/bash

COMMAND=${1:-up}

function execute-docker-compose () {
  docker-compose \
    -f 'docker-compose.yml' \
    $@
}

function stop-docker-compose () {
  execute-docker-compose stop
}

if [ $COMMAND = 'up' ] && [ $# -le 1 ]; then
  execute-docker-compose up
  stop-docker-compose

elif [ $COMMAND = 'bash-o' ]; then
  execute-docker-compose exec ftgo-order /bin/bash
elif [ $COMMAND = 'bash-o-p' ]; then
  execute-docker-compose exec ftgo-order-postgres /bin/bash

else
  execute-docker-compose $@
fi
