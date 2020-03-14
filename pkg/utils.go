package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ConvertJsCodeOBJToMap(text string,keysLen int) (interface{} ,error){
	keyRe := regexp.MustCompile("[,|{](\\w+):")
	keys := keyRe.FindAllStringSubmatch(text, keysLen)
	for _, val := range keys {
		text = strings.ReplaceAll(text, val[1], fmt.Sprintf("\"%v\"", val[1]))
	}
	var obj interface{}
	err := json.Unmarshal([]byte(text), &obj)
	if err != nil {
		log.Println(err)
		return map[string]interface{}{},err
	}
	return obj,nil
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}