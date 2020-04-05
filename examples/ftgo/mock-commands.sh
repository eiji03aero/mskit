#!/bin/bash

if [ $# -lt 1 ]; then
  cat <<- EOF
demo.sh: error

You need to pass one arguments to select what to demo
  - createOrder: creating order
EOF
  exit 1;
fi

res_file="./response.txt"
debug_options="--dump-header -"
post_options="-X POST"
patch_options="-X PATCH"
save_res_options="-s -o $res_file"

command="${1}"

show-response-body () {
  echo $(cat $res_file)
}

get-by-key-or-value () {
  key=$1
  value=$2

  if [ -n "$value" ]; then
    echo $value
  else
    echo $(./temp-data.sh get-by-key $key)
  fi
}

store-id-from-file () {
  key=$1
  ./temp-data.sh write-by-key-from-file $key $res_file
  rm $res_file
}

if [ $command = "seed" ]; then
  ./mock-commands.sh createRestaurant
  sleep 1s
  ./mock-commands.sh createConsumer
  sleep 1s
  ./mock-commands.sh createOrder

elif [ $command = "createOrder" ]; then
  restaurant_id=$(get-by-key-or-value restaurant_id $2)
  consumer_id=$(get-by-key-or-value consumer_id $3)
  curl $post_options $debug_options $save_res_options \
    -d $(printf '%s' $(cat <<- EOF
      {
        "restaurant_id": "$restaurant_id",
        "consumer_id": "$consumer_id",
        "payment_information": { "token": "daiji na token" },
        "delivery_information": {
          "address": { "zip_code": "300-9999" }
        },
        "order_line_items": {
          "line_items": [
            { "menu_item_id": "awesome-papas", "quantity": 5 }
          ]
        }
      }
EOF
    )) \
    ftgo-order:3000/orders
  show-response-body
  store-id-from-file order_id

elif [ $command = "getOrder" ]; then
  curl $debug_options \
    ftgo-order:3000/orders/$2

elif [ $command = "reviseOrder" ]; then
  order_id=$(get-by-key-or-value order_id $3)
  curl $patch_options $debug_options \
    -d $(printf '%s' $(cat <<- EOF
      {
        "order_line_items": {
          "line_items": [
            { "menu_item_id": "awesome-papas", "quantity": 8 }
          ]
        }
      }
EOF
    )) \
    ftgo-order:3000/orders/$order_id

elif [ $command = "createRestaurant" ]; then
  curl $post_options $debug_options $save_res_options \
    -d '
      {
        "name": "macdonalds",
        "restaurant_menu": {
          "menu_items": [
            { "id": "awesome-papas", "name": "fries large", "price": 300 },
            { "id": "second-yum", "name": "cheese burger", "price": 200 }
          ]
        }
      }
    ' \
    ftgo-restaurant:3002/restaurants
  show-response-body
  store-id-from-file restaurant_id

elif [ $command = "getRestaurant" ]; then
  id=$(get-by-key-or-value restaurant_id $2)
  curl $debug_options \
    ftgo-order:3000/restaurants/$id

elif [ $command = "createConsumer" ]; then
  curl $post_options $debug_options $save_res_options \
    -d '
      {
        "name": "pink guy"
      }
    ' \
    ftgo-consumer:3003/consumers
  show-response-body
  store-id-from-file consumer_id

elif [ $command = "getConsumer" ]; then
  id=$(get-by-key-or-value consumer_id $2)
  curl $debug_options \
    ftgo-consumer:3003/consumers/$id
fi
