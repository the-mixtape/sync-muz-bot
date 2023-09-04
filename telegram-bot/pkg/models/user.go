package models

type User struct {
	Id          int64   `db:"id"`
	Username    string  `db:"username"`
	VkId        *int64  `db:"vk_id"`
	VkSync      *string `db:"vk_sync"`
	SyncUTCTime *string `db:"sync_utc_time"`
}
