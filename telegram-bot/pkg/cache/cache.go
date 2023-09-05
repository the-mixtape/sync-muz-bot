package cache

type BotCache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Delete(key string)

	//SetExpectedMessageType(userId int64, expectedMsgType int)
	//GetExpectedMessageType(userId int64) int
}
