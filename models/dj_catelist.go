package models

import (
	"NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain)DjCatelist(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/weapi/djradio/category/get" ,
		data,
		options)
}