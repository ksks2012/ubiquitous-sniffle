package global

import (
	"database/sql"

	"github.com/go-redis/redis"
)

var DBEngine *sql.DB

var CacheEngine *redis.Client
