package models

import (
	"regexp"

	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) Batch(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"e_r": true,
	}
	havaRep := regexp.MustCompile("^/api/")
	for key, val := range query {
		if havaRep.MatchString(key) {
			data[key] = val
		}
	}
	options := map[string]interface{}{
		"crypto": "eapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
		"url":    "/api/batch",
	}
	return request.CreateRequest(
		"POST", "http://music.163.com/eapi/batch",
		data,
		options)
}
