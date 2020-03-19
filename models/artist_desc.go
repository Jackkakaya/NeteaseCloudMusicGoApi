package models

import "github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"

// 歌手介绍
func (m *MusicObain) ArtistDesc(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"id": query["id"],
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/weapi/artist/introduction",
		data,
		options)
}
