package glredis

import "github.com/redis/go-redis/v9"

var setHashKey = redis.NewScript(`
	local hashKey = KEYS[1]
	local field = ARGV[1]
	local value = ARGV[2]
	redis.call('HSET', hashKey, field, value)
`)

var delHashField = redis.NewScript(`
	local hashKey = KEYS[1]
	local field = ARGV[1]
	redis.call('HDEL', hashKey, field)
`)

var getHashFiled = redis.NewScript(`
	local hashKey = KEYS[1]
	local field = ARGV[1]
	return redis.call('HGET', hashKey, field)
`)
