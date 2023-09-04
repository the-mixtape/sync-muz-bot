package vk_api

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"golang.org/x/exp/slog"
)

type VkApi struct {
	vk *api.VK
}

func getLoggerMsg(msg string) string {
	return "[VkApi] " + msg
}

func NewVKApi(token string) *VkApi {
	vkApi := VkApi{api.NewVK(token)}

	_, userName, err := vkApi.GetUserIdAndName(nil)
	if err != nil {
		slog.Error(getLoggerMsg("Failed to retrieve user id and name"))
		return nil
	}
	slog.Info(getLoggerMsg("Authorized on account " + userName))

	return &vkApi
}

func (vkApi *VkApi) GetUserIdAndName(id interface{}) (userId int, userName string, err error) {
	var params api.Params = nil
	if id != nil {
		params = api.Params{"user_ids": id}
	}

	info, err := vkApi.vk.UsersGet(params)
	if err == nil {
		userInfo := info[0]
		userId = userInfo.ID
		userName = userInfo.FirstName + " " + userInfo.LastName
	}
	return
}
