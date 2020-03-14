package models

import (
	"NeteaseCloudGoApi/pkg/request"
)

func (m *MusicObain)DjCategoryExcludehot(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","http://music.163.com/weapi/djradio/category/excludehot" ,
		data,
		options)
}