package redis

import (
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/lib"
	"fmt"
	"sync"

	"github.com/garyburd/redigo/redis"
)

//全局实现者
var pool RedisPooler
var poolLock sync.Mutex

//NewPool 初始化数据库连接
func NewPool(config RedisConnect) RedisPooler{
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool==nil{
		pool=&RedisPool{}
		pool.Init(config)
	}
	return pool;
}
//Rediser 操作接口 
type Rediser interface {
	Master()Master
	Slave()Master
	SlaveByName(name string)Master
}
//Master 从接口
type Master interface{
	Set(key string, value interface{}, overtime int64) bool
	SetNx(key string, value interface{}) bool
	Get(key string, value interface{}) bool
	Del(key string) bool
	Ttl(key string) int64
	Expire(key string, overtime int64) bool
	Len(key string) int64
	FlushDB() bool
	Exists(key string) bool
	Keys(key string) []string
	RPush(key string, value ...interface{}) int64
	LPush(key string, value ...interface{}) int64
	LPop(key string, value interface{})bool
	RPop(key string, value interface{}) bool
	Type(key string) (string, bool)
	Ping() bool
	GetRequirepass() bool
	SetRequirepass(password string) bool
	Select(index int) bool
	HMset(key string, value lib.InRow) bool
	HDel(key string, field ...string) int
	HExists(key string, field string) int
	HGet(key string, field string, value interface{}) bool
	HGetAll(key string)lib.Row
	HLen(key string) int
	HKeys(key string) []string
	HSet(key string, field string, value interface{}) int
	HSetNx(key string, field string, value interface{}) bool
	SAdd(key string, values ...interface{}) int
	SCard(key string) int
	SDiff(keys ...string)[][]byte
	SDiffStore(destination string, keys ...string) int
	SInter(keys ...string)[][]byte
	SInterStore(destination string, keys ...string) int
	SIsMember(key string, value interface{}) bool
	SMembers(key string) [][]byte
	SMove(source, destination string, member interface{}) int
	SPop(key string,value interface{}) bool
	SRandMember(key string, count int) [][]byte
	SRem(key string, members ...interface{}) int
	Push(key string, value interface{}) bool
	Pop(key string, value interface{}) bool
}
//Redis redis对象
type Redis struct {
	pool RedisPooler
	isMaster bool
	slaveName string

}
//Master 主redis
func (r *Redis)Master()Master{
	r.isMaster=true
	return r
}
//Slave 从redis
func (r *Redis)Slave()Master{
	r.isMaster=false
	r.slaveName=""
	return r
}
//SlaveByName 从redis 通过名称
func (r *Redis)SlaveByName(name string)Master{
	r.isMaster=false
	r.slaveName=name
	return r
}
//getRedis 获得redis
func (r *Redis)getRedis() *redis.Pool{
	if r.isMaster{
		return r.pool.GetMaster()
	}else{
		if r.slaveName=="" {
			return r.pool.GetSlave()
		}else{
			return r.pool.GetSlaveByName(r.slaveName)
		}
		
	}
}
//SetPool 设置池
func (r *Redis)SetPool(pooler RedisPooler)Rediser {
	r.pool=pooler
	return r
}
//Set set操作
func (r *Redis)Set(key string, value interface{}, overtime int64) bool{
	if overtime>-1{
		_, err :=r.getRedis().Get().Do("SET",key,lib.Serialize(value),"EX",overtime)
		if err != nil {
			return false
		}
	}else{
		_, err :=r.getRedis().Get().Do("SET",key,lib.Serialize(value))
		if err != nil {
			return false
		}
	}
	return true
}
//SetNx SetNx 操作
func (r *Redis)SetNx(key string, value interface{}) bool{
	v, err :=redis.Int64(r.getRedis().Get().Do("SETNX", key, value))
	if err!=nil{
		return false
	}
	if v == 0 {
		return false
	}
	return true
}
//Get Get操作
func (r *Redis)Get(key string, value interface{}) bool{
	v, err :=redis.Bytes(r.getRedis().Get().Do("GET", key))
	if err!=nil{
		return false
	}
	lib.UnSerialize(v, value)
	return true
}
//Del Del操作
func (r *Redis)Del(key string) bool{
	_, err := r.getRedis().Get().Do("DEL", key)
	if err!=nil{
		return false
	}
	return true
}
//Ttl Ttl操作
func (r *Redis)Ttl(key string) int64{
	v, err :=redis.Int64(r.getRedis().Get().Do("TTL", key))
	if err!=nil{
		return -1
	}
	return v
}
//Expire Expire操作
func (r *Redis)Expire(key string, overtime int64) bool{
	_, err :=r.getRedis().Get().Do("EXPIRE", key, overtime)
	if err!=nil{
		return false
	}
	return true
}
//Len Len操作
func (r *Redis)Len(key string) int64{
	v, err :=redis.Int64(r.getRedis().Get().Do("LLEN", key))
	if err!=nil{
		return -1
	}
	return v
}
//FlushDB FlushDB操作
func (r *Redis)FlushDB() bool{
	_, err :=r.getRedis().Get().Do("FLUSHDB")
	if err != nil {
		return false
	}
	return true
}
//Exists Exists操作
func (r *Redis)Exists(key string) bool{
	v, err :=redis.Int(r.getRedis().Get().Do("EXISTS",key))
	if err != nil {
		return false
	}
	if v == 1 {
		return true
	}else{
		return false
	}
}
//Keys Keys操作
func (r *Redis)Keys(key string) []string{
	v, err := redis.Strings(r.getRedis().Get().Do("KEYS", key))
	if err != nil {
		return nil
	}
	return v
}
//RPush RPush操作
func (r *Redis)RPush(key string, value ...interface{}) int64{
	var arg []interface{}
	arg = append(arg, key)
	for _, d := range value {
		arg = append(arg, lib.Serialize(d))
	}
	v, err :=redis.Int64(r.getRedis().Get().Do("Rpush", arg...))
	if err != nil {
		return -1
	}
	return v
}
//LPush LPush操作
func (r *Redis)LPush(key string, value ...interface{}) int64{
	var arg []interface{}
	arg = append(arg, key)
	for _, d := range value {
		arg = append(arg, lib.Serialize(d))
	}
	v, err :=redis.Int64(r.getRedis().Get().Do("Lpush", arg...))
	if err != nil {
		return -1
	}
	return v
}
//LPop LPop操作
func (r *Redis)LPop(key string, value interface{})bool{
	v, err := redis.Bytes(r.getRedis().Get().Do("Lpop", key))
	if err != nil {
		return false
	}
	lib.UnSerialize(v,value)
	return true
}
//RPop RPop操作
func (r *Redis)RPop(key string, value interface{}) bool{
	v, err := redis.Bytes(r.getRedis().Get().Do("Rpop", key))
	if err != nil {
		return false
	}
	lib.UnSerialize(v,value)
	return true
}
//Push Push操作
func (r *Redis)Push(key string, value interface{}) bool{
	i:=r.LPush(key,value)
	if i>-1{
		return true
	}
	return false
}
//Push Push操作
func (r *Redis)Pop(key string, value interface{}) bool{
	return r.RPop(key,value)
}
//Type Type操作
func (r *Redis)Type(key string) (string, bool){
	v, err := redis.String(r.getRedis().Get().Do("TYPE", key))
	if err!=nil{
		return "",false
	}
	return v,true
}
//Ping Ping操作
func (r *Redis)Ping() bool{
	v, err := redis.String(r.getRedis().Get().Do("PING"))
	if err!=nil{
		return false
	}
	if v == "PONG" {
		return true
	} else {
		return false
	}
}
//GetRequirepass GetRequirepass操作
func (r *Redis)GetRequirepass() bool{
	v, err := redis.Strings(r.getRedis().Get().Do("CONFIG", "get", "requirepass"))
	if err != nil {
		return false
	}
	//fmt.Println(v)
	if v[1] == "" {
		return false
	} else {
		return true
	}
}
//SetRequirepass SetRequirepass操作
func (r *Redis)SetRequirepass(password string) bool{
	v, err := redis.String(r.getRedis().Get().Do("CONFIG", "set", "requirepass", password))
	if err != nil {
		return false
	}
	if v == "OK" {
		return true
	} else {
		return false
	}
}
//Select 选择数据库
func (r *Redis)Select(index int) bool{
	v, err := redis.String(r.getRedis().Get().Do("SELECT", index))
	if err != nil {
		return false
	}
	if v == "OK" {
		return true
	} else {
		return false
	}
}

//HMset HMset操作
func (r *Redis)HMset(key string, value lib.InRow) bool{
	var arg []interface{}
	arg = append(arg, key)
	for k,iv:=range value{
		arg = append(arg, k)
		arg = append(arg, lib.Serialize(iv))
	}
	v, err := redis.String(r.getRedis().Get().Do("HMSET", arg...))
	if err != nil {
		return false
	}
	if v == "OK" {
		return true
	} else {
		return false
	}
}

//HDel HDel操作
func (r *Redis)HDel(key string, field ...string) int{
	var args []interface{}
	args = append(args, key)
	for _, v := range field {
		args = append(args, v)
	}
	v, err := redis.Int(r.getRedis().Get().Do("HDEL", args...))
	if err != nil {
		return 0
	}
	return v
}

//HExists HExists操作
func (r *Redis)HExists(key string, field string) int{
	v, err := redis.Int(r.getRedis().Get().Do("HEXISTS", key, field))
	if err != nil {
		return -1
	}
	return v
}

//HGet HGet操作
func (r *Redis)HGet(key string, field string, value interface{}) bool{
	v, err := redis.Bytes(r.getRedis().Get().Do("HGET", key, field))
	if err != nil {
		return false
	}
	lib.UnSerialize(v,value)
	return true
}

//HGetAll HGetAll操作
func (r *Redis)HGetAll(key string)lib.Row{
	v, err :=redis.ByteSlices(r.getRedis().Get().Do("HGETALL", key))
	if err != nil {
		return nil
	}
	value := make(lib.Row)
	for i := 0; i < len(v); i = i + 2 {
		vl:=&lib.Data{}
		value[string(v[i])]=vl.Set(v[i+1])
	}
	return value
}


//HLen HLen操作
func (r *Redis)HLen(key string) int{
	v, err := redis.Int(r.getRedis().Get().Do("HLEN", key))
	if err != nil {
		return -1
	}
	return v
}

//HKeys HKeys操作
func (r *Redis)HKeys(key string) []string{
	v, err := redis.Strings(r.getRedis().Get().Do("HKEYS", key))
	if err != nil {
		return nil
	}
	return v
}

//HSet HSet操作
func (r *Redis)HSet(key string, field string, value interface{}) int{
	v, err := redis.Int(r.getRedis().Get().Do("HSET", key, field, lib.Serialize(value)))
	if err != nil {
		return -1
	}
	return v
}

//HSetNx HSetNx操作
func (r *Redis)HSetNx(key string, field string, value interface{}) bool{
	v, err := redis.Int(r.getRedis().Get().Do("HSETNX", key, field, lib.Serialize(value)))
	if err != nil {
		return false
	}
	if v==1{
		return true
	}
	return false
}
//SAdd SAdd 操作
func (r *Redis) SAdd(key string, values ...interface{}) int {
	var args []interface{}
	args = append(args, key)
	for _, v := range values {
		args = append(args, lib.Serialize(v))
	}
	v, err := redis.Int(r.getRedis().Get().Do("SADD", args...))
	if err != nil {
		return -1
	}
	return v
}
//SCard SCard 操作
func (r *Redis)SCard(key string) int{
	v, err := redis.Int(r.getRedis().Get().Do("SCARD", key))
	if err != nil {
		return -1
	}
	return v
}
//SDiff SDiff 操作命令返回第一个集合与其他集合之间的差异
func (r *Redis)SDiff(keys ...string)[][]byte{
	var args []interface{}
	for _, k := range keys {
		args = append(args, k)
	}
	v, err := redis.ByteSlices(r.getRedis().Get().Do("SDIFF", args...))
	if err != nil {
		return nil
	}
	return v
}
//SDiffStore SDiffStore 命令将给定集合之间的差集存储在指定的集合中
func (r *Redis)SDiffStore(destination string, keys ...string) int{
	var args []interface{}
	args = append(args, destination)
	for _, k := range keys {
		args = append(args, k)
	}
	v, err := redis.Int(r.getRedis().Get().Do("SDIFFSTORE", args...))
	if err != nil {
		return -1
	}
	return v
}
//SInter SInter 操作
func (r *Redis)SInter(keys ...string)[][]byte{
	var args []interface{}
	for _, k := range keys {
		args = append(args, k)
	}
	v, err := redis.ByteSlices(r.getRedis().Get().Do("SINTER", args...))
	if err != nil {
		return nil
	}
	return v
}
//SInterStore SInterStore 操作
func (r *Redis)SInterStore(destination string, keys ...string) int{
	var args []interface{}
	args = append(args, destination)
	for _, k := range keys {
		args = append(args, k)
	}
	v, err := redis.Int(r.getRedis().Get().Do("SINTERSTORE", args...))
	if err != nil {
		return -1
	}
	return v
}
//SIsMember SIsMember 命令判断成员元素是否是集合的成员。
func (r *Redis)SIsMember(key string, value interface{}) bool{
	v, err := redis.Int(r.getRedis().Get().Do("SISMEMBER", key, lib.Serialize(value)))
	if err != nil {
		return false
	}
	if v==1{
		return true
	}
	return false
}
//SMembers SMembers 命令返回集合中的所有的成员。 不存在的集合 key 被视为空集合。
func (r *Redis)SMembers(key string) [][]byte{
	v, err := redis.ByteSlices(r.getRedis().Get().Do("SMEMBERS", key))
	if err != nil {
		return nil
	}
	return v
}
//SMove SMove 将 member 元素从 source 集合移动到 destination 集合
func (r *Redis)SMove(source, destination string, member interface{}) int{
	v, err := redis.Int(r.getRedis().Get().Do("SMOVE", source, destination, lib.Serialize(member)))
	if err != nil {
		return -1
	}
	return v
}
//SPop SPop 移除并返回集合中的一个随机元素
func (r *Redis)SPop(key string,value interface{}) bool{
	v, err := redis.Bytes(r.getRedis().Get().Do("SPOP", key))
	if err != nil {
		return false
	}
	lib.UnSerialize(v,value)
	return true
}
//SRandMember SRandMember 返回集合中一个或多个随机数
func (r *Redis)SRandMember(key string, count int) [][]byte{
	v, err := redis.ByteSlices(r.getRedis().Get().Do("SRANDMEMBER", key, count))
	if err != nil {
		return nil
	}
	return v
}
//SRem SRem 移除集合中一个或多个成员
func (r *Redis)SRem(key string, members ...interface{}) int{
	var args []interface{}
	args = append(args, key)
	for _, k := range members {
		args = append(args, lib.Serialize(k))
	}
	v, err := redis.Int(r.getRedis().Get().Do("SREM", args...))
	if err != nil {
		return -1
	}
	return v
}
//NewRedis 新建一个redis
func NewRedis()Rediser{
	redisFile:="config/redis.config"
	cfgFile:=file.GetContent(redisFile)
	var config RedisConnect
	if cfgFile==""{
		var slaves []SlaveDB=make([]SlaveDB, 1)
		master:=SlaveDB{
			Name:"name",
			Db:0,
			Password:"password",
			Host:"127.0.0.1",
			Port:"6379",
			MaxIdle:20,
			IdleTimeout:10,
		}
		config=RedisConnect{
			Master:master,
			Slave:slaves,
		}
		file.PutContent(redisFile,fmt.Sprintf("%v",config))
		panic("please setting redis config in config/redis.config file")
	}
	lib.StringToObject(cfgFile, &config)
	var rds Redis
	return rds.SetPool(NewPool(config))
}