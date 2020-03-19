package models

import (
	"NeteaseCloudMusicGoApi/pkg/request"
	"fmt"
	"strconv"
)

func (m *MusicObain)CheckMusic(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	if val, err := strconv.Atoi(fmt.Sprintf("%v", query["id"])); err == nil {
		data["ids"] = fmt.Sprintf("[%v]", val)
	}
	if val, ok := query["br"]; ok {
		data["br"] = val
	} else {
		data["br"] = 999000
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	resp := request.CreateRequest(
		"POST", "https://music.163.com/weapi/song/enhance/player/url",
		data,
		options)
	playable := false
	if val, ok := resp["body"]; ok {
		valMapper := val.(map[string]interface{})
		if fmt.Sprintf("%v", valMapper["code"]) == "200" {
			playable = true
		}
	}
	if playable {
		resp["body"] = map[string]interface{}{"success": true, "message": "ok"}
	} else {
		resp["status"] = 404
		resp["body"] = map[string]interface{}{"success": false, "message": "亲爱的,暂无版权"}
	}
	return resp
}