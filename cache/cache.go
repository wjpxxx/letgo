package cache

import (
	"core/cache/filecache"
	"core/cache/icache"
	"core/cache/memcache"
	"core/cache/redis"
)

var cacheList map[string]icache.ICacher
//Register 注册缓存对象
func Register(name string, cacher icache.ICacher) {
	if cacheList==nil{
		cacheList=make(map[string]icache.ICacher)
	}
	cacheList[name]=cacher
}
//NewCache 新建一个缓存对象
func NewCache(name string) icache.ICacher{
	return cacheList[name]
}
//NewRedis 新建一个redis对象
func NewRedis()redis.Rediser{
	return redis.NewRedis()
}
//init 注册系统的
func init(){
	Register("redis", redis.NewRedis().Master())
	Register("file", filecache.NewFileCache())
	Register("memcache", memcache.NewMemcache())
}