package models

import "github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"

func (m *MusicObain) AlbumDetailDynamic(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"id": query["id"],
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/api/album/detail/dynamic",
		data,
		options)
}
