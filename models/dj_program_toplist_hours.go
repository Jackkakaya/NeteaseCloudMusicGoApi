package models

import (
	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

// 24小时榜单
func (m *MusicObain) DjProgramToplistHours(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	if val, ok := query["limit"]; ok {
		data["limit"] = val
	} else {
		data["limit"] = 100
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/api/djprogram/toplist/hours",
		data,
		options)
}
