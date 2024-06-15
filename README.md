# **Intro**

A URL shortener

# **Command**

## **Server**

```
go run cmd/main.go

```

## **pkg testing**

- hash
    
    ```
    go test ./pkg/hash/
    ```

- test post method

    ```
    wget --post-data 'url=www.google.com&method=md5' localhost:18080/api/v1/shorten -O ./log/md5.json
    ```

    ```
    wget --post-data 'url=www.google.com&method=sha1' localhost:18080/api/v1/shorten -O ./log/sha1.json
    ```

    ```
    wget --post-data 'url=www.google.com&method=crc32' localhost:18080/api/v1/shorten -O ./log/crc32.json
    ```

# Planning

## Feature

- [ ]  API for create short URL
- [ ]  API for redirected to long URL
- [ ]  Cache

## DB table

- UUID by snowflake
- short_url
- long_url

# plugin

- storage
- shorten method
    - hashing method
        - collision resolution
        - decrease hash value
        - bloom filter
    - base62