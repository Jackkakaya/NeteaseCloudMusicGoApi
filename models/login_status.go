package models

import (
	"regexp"

	neteasePkg "github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg"
	"github.com/Jackkakaya/NeteaseCloudMusicGoApi/pkg/request"
)

func (m *MusicObain) LoginStatus(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	options := map[string]interface{}{
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	res := request.CreateRequest(
		"GET", "https://music.163.com",
		data,
		options)
	profileRe := regexp.MustCompile("GUser\\s*=\\s*([^;]+);")
	bindingsRe := regexp.MustCompile("GBinds\\s*=\\s*([^;]+);")

	text := res["body"]
	profileRaw := profileRe.FindStringSubmatch(text.(string))
	bindingsRaw := bindingsRe.FindStringSubmatch(text.(string))
	body := map[string]interface{}{}
	if lenProfile := len(profileRaw); lenProfile >= 2 {
		profile, _ := neteasePkg.ConvertJsCodeOBJToMap(profileRaw[1], 100)
		body["profile"] = profile
	}
	if lenBindings := len(bindingsRaw); lenBindings >= 2 {
		bindings, _ := neteasePkg.ConvertJsCodeOBJToMap(bindingsRaw[1], 100)
		body["bindings"] = bindings
	}
	res["body"] = body
	return res
}
