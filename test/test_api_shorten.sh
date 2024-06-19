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

short_urls=("480a829c2f42c8c53626d162cf3c934b" "8e0d2e72cd72bc5dd83b5928aeb44bcc9c73e953" "1277064271")
long_names=("md5_long" "sha1_long" "crc32_long")

for i in "${!short_urls[@]}"
do
    wget "$HOST/$API/shortUrl/${short_urls[$i]}" -O "$log_dir/${long_names[$i]}.json"
done