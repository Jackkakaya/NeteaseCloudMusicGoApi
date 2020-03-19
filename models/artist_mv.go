package models

import "github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"

//歌手mv
func (m *MusicObain) ArtistMv(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"artistId": query["id"],
		"limit":    query["limit"],
		"offset":   query["offset"],
		"total":    true,
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/weapi/artist/mvs",
		data,
		options)
}
