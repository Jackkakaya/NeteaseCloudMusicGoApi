package models

import (
	"fmt"

	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) Banner(query map[string]interface{}) map[string]interface{} {
	typeMapper := map[string]string{
		"0": "pc",
		"1": "android",
		"2": "iphone",
		"3": "ipad",
	}
	clientType := "pc"
	if val, ok := query["type"]; ok {
		clientType, ok = typeMapper[fmt.Sprintf("%v", val)]
		if !ok {
			clientType = "pc"
		}
	}
	data := map[string]interface{}{
		"clientType": clientType,
	}
	options := map[string]interface{}{
		"crypto": "linuxapi",
		"proxy":  query["proxy"],
	}
	return request.CreateRequest(
		"POST", "https://music.163.com/api/v2/banner/get",
		data,
		options)
}
