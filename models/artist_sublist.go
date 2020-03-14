package models

import "NeteaseCloudGoApi/pkg/request"

// 关注歌手列表
// todo 待测试
func (m *MusicObain)ArtistSublist(query map[string]interface{}) map[string]interface{}  {
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
		"POST","https://music.163.com/weapi/artist/sublist",
		data,
		options)
}