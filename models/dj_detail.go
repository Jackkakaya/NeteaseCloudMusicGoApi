package models

import (
	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) DjDetail(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"id": query["rid"],
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/weapi/djradio/get",
		data,
		options)
}
