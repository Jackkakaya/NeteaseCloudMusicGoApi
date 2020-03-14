package models

import "NeteaseCloudGoApi/pkg/request"

func (m *MusicObain)AlbumNewest(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/api/discovery/newAlbum",
		data,
		options)
}
