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
save_res_options="-s -o $res_file"

command="${1}"

show-response-body () {
  echo $(cat $res_file)
}

value-or-get () {
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

if [ $command = "createOrder" ]; then
  restaurant_id=$(value-or-get restaurant_id $2)
  curl $post_options $debug_options $save_res_options \
    -d $(printf '%s' $(cat <<- EOF
      {
        "restaurant_id": "$restaurant_id",
        "payment_information": { "token": "daiji na token" },
        "delivery_information": {
          "address": { "zip_code": "300-9999" }
        },
        "order_line_items": {
          "line_items": [
            { "menu_item_id": "kore", "quantity": 5 },
            { "menu_item_id": "soukai", "quantity": 1 }
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

elif [ $command = "createRestaurant" ]; then
  curl $post_options $debug_options $save_res_options \
    -d '
      {
        "name": "macdonalds",
        "restaurant_menu": {
          "menu_items": [
            { "name": "fries large", "price": 300 },
            { "name": "cheese burger", "price": 200 }
          ]
        }
      }
    ' \
    ftgo-restaurant:3002/restaurants
  show-response-body
  store-id-from-file restaurant_id

elif [ $command = "getRestaurant" ]; then
  id=$(value-or-get restaurant_id $2)
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
  id=$(value-or-get consumer_id $2)
  curl $debug_options \
    ftgo-consumer:3003/consumers/$id
fi
