package models

import (
	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) CommentHotwallList(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/api/comment/hotwall/list/get",
		data,
		options)
}
