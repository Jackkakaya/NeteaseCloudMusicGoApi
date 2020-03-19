package models

import (
	"NeteaseCloudMusicGoApi/pkg/request"
	"fmt"
)

// 收藏与取消收藏歌手
// todo 待测试
func (m *MusicObain)ArtistSub(query map[string]interface{}) map[string]interface{}  {
	if fmt.Sprintf("%v", query["t"]) == "1" {
		query["t"] = "sub"
	}else {
		query["t"] = "unsub"
	}
	data := map[string]interface{} {
		"artistId" : query["id"],
		"artistIds": fmt.Sprintf("[%v]",query["id"]),
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/weapi/artist/"+query["t"].(string),
		data,
		options)
}