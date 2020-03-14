package models

import (
	"NeteaseCloudGoApi/pkg/request"
	"fmt"
)

func (m *MusicObain)ArtistAlbum(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {
		"total":true,
	}
	if val,ok := query["limit"];ok{
		data["limit"] = val
	}else {
		data["limit"] = 25
	}
	if val,ok := query["offset"];ok{
		data["offset"] = val
	}else {
		data["offset"] = 0
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/weapi/artist/albums/"+ fmt.Sprintf("%v", query["id"]),
		data,
		options)
}