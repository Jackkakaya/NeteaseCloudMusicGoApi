package models

import (
	"NeteaseCloudMusicGoApi/pkg/crypto"
	"NeteaseCloudMusicGoApi/pkg/request"
	"fmt"
	"reflect"
)

func (m *MusicObain)LoginCellphone(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {
		"phone" : query["phone"],
		"countrycode" : query["countrycode"],
		"password" : crypto.GetMd5Hex(fmt.Sprintf("%v",query["password"])),
		"rememberLogin" : true,
	}
	if val,err := query["cookie"];err==false{
		query["cookie"] = map[string]interface{}{
			"os":"pc",
		}
	}else {
		if reflect.ValueOf(val).Kind() == reflect.Map{
			val := val.(map[string]interface{})
			val["os"] = "pc"
			query["cookie"] = val
		}else{
			query["cookie"] = map[string]interface{}{
				"os":"pc",
			}
		}
	}
	options := map[string]interface{}{
		"crypto": "weapi",
		"ua":"pc",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","https://music.163.com/weapi/login/cellphone",
		data,
		options)
}