package models

import (
	"NeteaseCloudGoApi/pkg/request"
	"fmt"
)

func (m *MusicObain)CommentEvent(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {}
	if val,ok := query["limit"];ok{
		data["limit"] = val
	}else {
		data["limit"] = 20
	}
	if val,ok := query["offset"];ok{
		data["offset"] = val
	}else {
		data["offset"] = 0
	}
	if val,ok := query["before"];ok{
		data["beforeTime"] = val
	}else {
		data["beforeTime"] = 0
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/weapi/v1/resource/comments/"+fmt.Sprintf("%v",query["threadId"]) ,
		data,
		options)
}