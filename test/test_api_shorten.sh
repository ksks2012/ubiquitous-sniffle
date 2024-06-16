#!/bin/bash

HOST='localhost:18080'
API='api/v1'
URL='www.google.com'

methods=("md5" "sha1" "crc32")
log_dir="./log"

for method in "${methods[@]}"
do
    wget --post-data "url=$URL&method=$method" "$HOST/$API/shorten" -O "$log_dir/$method.json"
done

short_urls=("0a137b375cc3881a70e186ce2172c8d1" "d8b99f68b208b5453b391cb0c6c3d6a9824f3c3a" "526628817")
long_names=("md5_long" "sha1_long" "crc32_long")

for i in "${!short_urls[@]}"
do
    wget "$HOST/$API/shortUrl/${short_urls[$i]}" -O "$log_dir/${long_names[$i]}.json"
done