package glredis

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

type GoLuaRedisHash interface {
	SetHashKey(rdb *redis.Client, key []string, field, value interface{}) error
	DelHashField(rdb *redis.Client, key []string, filed string) error
	GetHashFiled(rdb *redis.Client, key []string, filed string) (string, error)
}

type GLRHash struct{}

func NewGLRHash() *GLRHash {
	return &GLRHash{}
}

// SetHashKey set hash filed value
func (glh *GLRHash) SetHashKey(rdb *redis.Client, key []string, field string, value interface{}) error {
	doCheckKey(key)
	doCheckHashField(field)
	err := setHashKey.Run(ctx, rdb, key, field, value).Err()
	return errIs(err)
}

// DelHashField delete hash filed
func (glh *GLRHash) DelHashField(rdb *redis.Client, key []string, field string) error {
	doCheckKey(key)
	doCheckHashField(field)
	err := delHashField.Run(ctx, rdb, key, field).Err()
	return errIs(err)
}

// GetHashFiled get hash filed
func (glh *GLRHash) GetHashFiled(rdb *redis.Client, key []string, field string) (interface{}, error) {
	doCheckKey(key)
	doCheckHashField(field)
	res, err := getHashFiled.Run(ctx, rdb, key, field).Result()
	if err = errIs(err); err != nil {
		return "", err
	}
	return res, nil
}

func errIs(err error) error {
	if errors.Is(err, redis.Nil) || errors.Is(err, nil) {
		return nil
	} else {
		return err
	}
}
