#!/bin/bash

HOST='localhost:18080'
API='api/v1/shorten'
URL='www.google.com'

methods=("md5" "sha1" "crc32")

for method in "${methods[@]}"
do
	wget --post-data "url=$URL&method=$method" "$HOST/$API" -O "./log/$method.json"
done