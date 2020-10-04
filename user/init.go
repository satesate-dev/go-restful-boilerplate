package user

import (
	"database/sql"

	"github.com/go-redis/redis/v7"
)

var (
	DBPool    *sql.DB
	CachePool *redis.Client
)

func Init(db *sql.DB, cache *redis.Client) {
	DBPool = db
	CachePool = cache
}
