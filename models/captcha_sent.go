package models

import "NeteaseCloudGoApi/pkg/request"

func (m *MusicObain)CaptchaSent(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {
		"cellphone" : query["phone"],
	}
	if val,ok := query["ctcode"];ok{
		data["ctcode"] = val
	}else {
		data["ctcode"] = "86"
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/weapi/sms/captcha/sent",
		data,
		options)
}