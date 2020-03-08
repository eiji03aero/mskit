#!/bin/bash

cmd=${1:-up}

function execute-docker-compose () {
  docker-compose \
    -f 'docker-compose.yml' \
    $@
}

function stop-docker-compose () {
  execute-docker-compose stop
}

if [ $cmd = 'up' ] && [ $# -le 1 ]; then
  execute-docker-compose up
  stop-docker-compose

elif [ $cmd = 'bash-o' ]; then
  execute-docker-compose exec ftgo-order /bin/bash
elif [ $cmd = 'bash-o-p' ]; then
  execute-docker-compose exec ftgo-order-postgres /bin/bash

elif [ $cmd = 'bash-k' ]; then
  execute-docker-compose exec ftgo-kitchen /bin/bash
elif [ $cmd = 'bash-k-m' ]; then
  execute-docker-compose exec ftgo-kitchen-mongo /bin/bash

elif [ $cmd = 'bash-r' ]; then
  execute-docker-compose exec ftgo-restaurant /bin/bash
elif [ $cmd = 'bash-r-m' ]; then
  execute-docker-compose exec ftgo-restaurant-mongo /bin/bash

elif [ $cmd = 'bash-rm' ]; then
  execute-docker-compose exec ftgo-rabbitmq /bin/bash

elif [ $cmd = 'setup-db' ]; then
  execute-docker-compose exec ftgo-order make setup
  execute-docker-compose exec ftgo-kitchen-mongo mongo --eval \
    'db.getSiblingDB("mskit").events.remove({});'
  execute-docker-compose exec ftgo-restaurant-mongo mongo --eval \
    'db.getSiblingDB("mskit").events.remove({});'

elif [ $cmd = 'reset-mq' ]; then
  execute-docker-compose exec ftgo-rabbitmq rabbitmqctl stop_app
  execute-docker-compose exec ftgo-rabbitmq rabbitmqctl reset
  execute-docker-compose exec ftgo-rabbitmq rabbitmqctl start_app

else
  execute-docker-compose $@
fi
