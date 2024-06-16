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
- test get url

    ```
    wget localhost:18080/api/v1/shortUrl/0a137b375cc3881a70e186ce2172c8d1 -O ./log/md5_long.json
    ```

    ```
    wget localhost:18080/api/v1/shortUrl/d8b99f68b208b5453b391cb0c6c3d6a9824f3c3a -O ./log/sha1_long.json
    ```

    ```
    wget localhost:18080/api/v1/shortUrl/526628817 -O ./log/crc32_long.json
    ```



# Planning

## Feature

- [X]  API for create short URL
- [X]  API for redirected to long URL
- [X]  Cache
- [ ]  Salt

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