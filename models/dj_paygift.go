package models

import (
	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) DjPaygift(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	if val, ok := query["limit"]; ok {
		data["limit"] = val
	} else {
		data["limit"] = 30
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
		"POST", "https://music.163.com/weapi/djradio/home/paygift/list?_nmclfl=1",
		data,
		options)
}
