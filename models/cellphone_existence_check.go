package models

import "NeteaseCloudGoApi/pkg/request"

func (m *MusicObain)CellphoneExistenceCheck(query map[string]interface{}) map[string]interface{}  {
	data := map[string]interface{} {
		"cellphone" : query["phone"],
		"countrycode": query["countrycode"],
	}
	options := map[string]interface{}{
		"crypto": "eapi",
		"url": "/api/cellphone/existence/check",
		"cookie": query["cookie"],
		"proxy": query["proxy"],
	}
	return request.CreateRequest(
		"POST","http://music.163.com/eapi/cellphone/existence/check",
		data,
		options)
}