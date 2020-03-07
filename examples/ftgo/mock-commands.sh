#!/bin/bash

if [ $# -lt 1 ]; then
  cat <<- EOF
demo.sh: error

You need to pass one arguments to select what to demo
  - createOrder: creating order
EOF
  exit 1;
fi

command="${1}"

if [ $command = "createOrder" ]; then
  curl -X POST \
    --dump-header - \
    -d '
    {
      "payment_information": { "token": "daiji na token" },
      "delivery_information": {
        "address": { "zip_code": "359-0034" }
      },
      "order_line_items": {
        "line_items": [
          { "menu_item_id": "kore", "quantity": 5 },
          { "menu_item_id": "soukai", "quantity": 1 }
        ]
      }
    }' \
    localhost:3000/orders

elif [ $command = "createOrder-not-enough-items" ]; then
  curl -X POST \
    --dump-header - \
    -d '
    {
      "payment_information": { "token": "daiji na token" },
      "delivery_information": {
        "address": { "zip_code": "359-0034" }
      },
      "order_line_items": { "line_items": [] }
    }' \
    localhost:3000/order

elif [ $command = "getOrder" ]; then
  curl \
    --dump-header - \
    localhost:3000/orders/$2

elif [ $command = "createRestaurant" ]; then
  curl -X POST \
    --dump-header - \
    -d '
    {
      "name": "macdonalds",
      "restaurant_menu": {
        "menu_items": [
          { "name": "fries large", "price": 300 },
          { "name": "cheese burger", "price": 200 }
        ]
      }
    }' \
    localhost:3002/restaurants
elif [ $command = "getRestaurant" ]; then
  curl \
    --dump-header - \
    localhost:3000/restaurants/$2
fi
