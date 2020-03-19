package models

import (
	"fmt"

	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) DailySignin(query map[string]interface{}) map[string]interface{} {
	queryType := "0"
	if val, ok := query["type"]; ok {
		queryType = fmt.Sprintf("%v", val)
	}
	data := map[string]interface{}{
		"type": queryType,
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/weapi/point/dailyTask",
		data,
		options)
}
