package models

type User struct {
	Id          int64   `db:"id"`
	Username    string  `db:"username" binding:"required"`
	VkId        *int64  `db:"vk_id" binding:"required"`
	VkSync      *string `db:"vk_sync" binding:"required"`
	SyncUTCTime *string `db:"sync_utc_time" binding:"required"`
}
