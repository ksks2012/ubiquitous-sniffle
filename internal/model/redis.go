package model

import (
	"context"
	"errors"

	"github.com/go-redis/redis"
	"github.com/ksks2012/ubiquitous-sniffle/global"
	"github.com/ksks2012/ubiquitous-sniffle/pkg/setting"
)

func NewCacheEngine(cacheSetting *setting.CacheSettingS) (*redis.Client, error) {
	if cacheSetting == nil {
		return nil, errors.New("cacheSetting is nil")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cacheSetting.Host + ":" + cacheSetting.Port,
		Password: cacheSetting.Password, // no password set
		DB:       cacheSetting.Number,   // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	global.Logger.Info(context.TODO(), pong)
	return client, nil
}

func GetKey(key string) (string, error) {
	val, err := global.CacheEngine.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func SetKey(key string, value string) error {
	err := global.CacheEngine.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
