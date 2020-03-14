package models

import (
	"NeteaseCloudGoApi/pkg/request"
	"fmt"
)

func (m *MusicObain)CommentHot(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {
		"rid" : query["id"],
	}
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
	if val,ok := query["cookie"];ok{
		valMapper := val.(map[string]interface{})
		valMapper["os"] = "pc"
		query["cookie"] = valMapper
	}else {
		query["cookie"] = map[string]interface{}{
			"os":"pc",
		}
	}

	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	queryType := map[string]interface{}{
		"0": "R_SO_4_", //  歌曲
		"1": "R_MV_5_", //  MV
		"2": "A_PL_0_", //  歌单
		"3": "R_AL_3_", //  专辑
		"4": "A_DJ_1_", //  电台,
		"5": "R_VI_62_", //  视频
	}[fmt.Sprintf("%v",query["type"])]
	return request.CreateRequest(
		"POST","https://music.163.com/weapi/v1/resource/hotcomments/"+fmt.Sprintf("%v%v",queryType,query["id"]) ,
		data,
		options)
}