package vk_api

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"golang.org/x/exp/slog"
)

type VkApi struct {
	vk *api.VK
}

type VkUserInfo struct {
	UserId   int
	UserName string
}

func getLoggerMsg(msg string) string {
	return "[VkApi] " + msg
}

func NewVKApi(token string) *VkApi {
	vkApi := VkApi{api.NewVK(token)}

	userInfo, err := vkApi.GetUserIdAndName(nil)
	if err != nil {
		slog.Error(getLoggerMsg("Failed to retrieve user id and name"))
		return nil
	}
	slog.Info(getLoggerMsg("Authorized on account " + userInfo.UserName))

	return &vkApi
}

func (vkApi *VkApi) GetUserIdAndName(id interface{}) (userInfo VkUserInfo, err error) {
	var params api.Params = nil
	if id != nil {
		params = api.Params{"user_ids": id}
	}

	info, err := vkApi.vk.UsersGet(params)
	if err == nil {
		infoResponse := info[0]
		userInfo.UserId = infoResponse.ID
		userInfo.UserName = infoResponse.FirstName + " " + infoResponse.LastName
	}
	return
}
