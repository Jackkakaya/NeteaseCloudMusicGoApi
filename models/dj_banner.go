package models

import (
	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) DjBanner(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	if val, ok := query["cookie"]; ok {
		valMapper := val.(map[string]interface{})
		valMapper["os"] = "pc"
		query["cookie"] = valMapper
	} else {
		query["cookie"] = map[string]interface{}{
			"os": "pc",
		}
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}

	return request.CreateRequest(
		"POST", "http://music.163.com/weapi/djradio/banner/get",
		data,
		options)
}
