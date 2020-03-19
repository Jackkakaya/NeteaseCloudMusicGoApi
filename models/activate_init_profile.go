package models

import "github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"

func (m *MusicObain) ActivateInitProfile(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"nickname": query["nickname"],
	}
	options := map[string]interface{}{
		"crypto": "eapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
		"url":    "/api/activate/initProfile",
	}
	return request.CreateRequest(
		"POST", "http://music.163.com/eapi/activate/initProfile",
		data,
		options)
}
