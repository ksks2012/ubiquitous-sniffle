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