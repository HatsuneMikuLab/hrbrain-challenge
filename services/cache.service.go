package services

import (
	"time"
)

type ICacheSerive interface {
	SetValue(key string, value string) 
	GetValue(key string) (string, bool)
}

type cacheItem struct {
	ExpireAt time.Time
	Value string
}

type cacheService struct {
	store map[string]*cacheItem
}

func NewCacheService() *cacheService {
	return &cacheService{ store: make(map[string]*cacheItem, 200)}
}

func (cs *cacheService) SetValue(key string, value string) {
	// add record to cache and set expire time as 10 minutes
	cs.store[key] = &cacheItem{ 
		ExpireAt: time.Now().Add(1 * time.Minute),
		Value: value,
	}
}

func (cs *cacheService) GetValue(key string) (string, bool) {
	if item, ok := cs.store[key]; ok {
		if time.Now().UnixNano() > item.ExpireAt.UnixNano() {
			delete(cs.store, key)
			return "", false
		}
		return item.Value, true
	}
	return "", false
}