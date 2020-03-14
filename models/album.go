package models

import (
	"NeteaseCloudGoApi/pkg/request"
	"fmt"
)

func (m *MusicObain) Album (query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/weapi/v1/album/"+ fmt.Sprintf("%v", query["id"]),
		data,
		options)
}
