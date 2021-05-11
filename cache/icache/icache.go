package icache

//Master 从接口
type ICacher interface{
	Set(key string, value interface{}, overtime int64) bool
	Get(key string, value interface{}) bool
	Del(key string) bool
	FlushDB() bool
}