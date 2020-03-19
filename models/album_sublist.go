package models

// 已收藏专辑列表
import "github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"

func (m *MusicObain) AlbumSublist(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"total": true,
	}
	if val, ok := query["limit"]; ok {
		data["limit"] = val
	} else {
		data["limit"] = 25
	}
	if val, ok := query["offset"]; ok {
		data["offset"] = val
	} else {
		data["offset"] = 0
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/weapi/album/sublist",
		data,
		options)
}
