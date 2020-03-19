package models

import (
	"fmt"

	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

// todo 测试
func (m *MusicObain) AlbumSub(query map[string]interface{}) map[string]interface{} {
	if fmt.Sprintf("%v", query["t"]) == "1" {
		query["t"] = "sub"
	} else {
		query["t"] = "unsub"
	}
	data := map[string]interface{}{
		"id": query["id"],
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/api/album/"+query["t"].(string),
		data,
		options)
}
