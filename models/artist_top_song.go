package models

import (
	"NeteaseCloudGoApi/pkg/request"
)

// 歌手热门50首
func (m *MusicObain)ArtistTopSong(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {
		"id" : query["id"],
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/api/artist/top/song",
		data,
		options)
}