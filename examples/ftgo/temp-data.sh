#!/bin/bash

data_file="temp-data.txt"

cmd=${1:-init}

init () {
  cat > $data_file <<-EOF
order_id=
restaurant_id=
consumer_id=
EOF
}

write-by-key () {
  key=$1
  value=$2

  sed -i -e "s/^${key}=.*$/${key}=${value}/" $data_file
}

write-by-key-from-file () {
  key=$1
  file=$2

  value=$(cat $file)
  write-by-key $key $value
}

get-by-key () {
  key=$1

  value=$(sed -n -e "/^${key}=.*$/p" $data_file \
    | awk -F "=" '{print $2}')
  echo $value
}

if [ $cmd = "init" ]; then
  init

elif [ $cmd = "write-by-key" ]; then
  write-by-key $2 $3

elif [ $cmd = "write-by-key-from-file" ]; then
  write-by-key-from-file $2 $3

elif [ $cmd = "get-by-key" ]; then
  get-by-key $2 $3

fi
