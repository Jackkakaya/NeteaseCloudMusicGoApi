package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

// 歌手分类

/*
   categoryCode 取值
   入驻歌手 5001
   华语男歌手 1001
   华语女歌手 1002
   华语组合/乐队 1003
   欧美男歌手 2001
   欧美女歌手 2002
   欧美组合/乐队 2003
   日本男歌手 6001
   日本女歌手 6002
   日本组合/乐队 6003
   韩国男歌手 7001
   韩国女歌手 7002
   韩国组合/乐队 7003
   其他男歌手 4001
   其他女歌手 4002
   其他组合/乐队 4003

   initial 取值 a-z/A-Z
*/

func (m *MusicObain) ArtistList(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"total": true,
	}
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
	if val, ok := query["cat"]; ok {
		data["categoryCode"] = val
	} else {
		data["categoryCode"] = "1001"
	}
	if initial, ok := data["initial"]; ok {
		if val, err := strconv.Atoi(fmt.Sprintf("%v", initial)); err == nil {
			data["initial"] = val
		} else {
			data["initial"] = int(strings.ToUpper(fmt.Sprintf("%v", initial))[0])
		}
	} else {
		data["initial"] = nil
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/weapi/artist/list",
		data,
		options)
}
