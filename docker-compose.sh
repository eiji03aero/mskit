#!/bin/bash

COMMAND=${1:-up}
container_name="mskit"

execute-docker-compose () {
  docker-compose \
    -f 'docker-compose.yml' \
    $@
}

stop-docker-compose () {
  execute-docker-compose stop
}

if [ $COMMAND = 'up' ] && [ $# -le 1 ]; then
  execute-docker-compose up -d
  execute-docker-compose exec $container_name /bin/bash
  stop-docker-compose
elif [ $COMMAND = 'bash' ]; then
  execute-docker-compose exec $container_name /bin/bash

elif [ $COMMAND = 'bash-ftgo-order' ]; then
  execute-docker-compose exec -w /app/examples/ftgo/pkg/order $container_name /bin/bash

else
  execute-docker-compose $@
fi
