package cache

type BotCache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Delete(key string)
}
