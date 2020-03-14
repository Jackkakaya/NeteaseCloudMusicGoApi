package request

import (
	"NeteaseCloudGoApi/pkg/crypto"
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/willf/pad"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	netUrl "net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func encodeURIComponent(str string) string {
	r := netUrl.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}
func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func ChooseUserAgent(ua string) string{
	userAgentList := []string{
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
		"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	}
	rand.Seed(time.Now().UnixNano())
	index := 0
	if ua == "<nil>" {
		index = rand.Int() % len(userAgentList)
	}else if ua == "mobile" {
		index = rand.Int() % 7
	}else if ua == "pc" {
		index  = (rand.Int() % 5) + 8
	}else {
		return ua
	}
	return userAgentList[index]
}

func CreateRequest(method string, url string, data map[string]interface{}, options map[string]interface{}) map[string]interface{}{
	headers := make(map[string]interface{})
	headers["User-Agent"] = ChooseUserAgent(fmt.Sprintf("%v", options["ua"]))
	if strings.ToUpper(method) == "POST"{
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	if strings.Contains(url,"music.163.com") {
		headers["Referer"] = "https://music.163.com"
	}
	// todo: 待测试
	value,ok:= options["cookie"]
	if ok && reflect.ValueOf(value).Kind() == reflect.Map{
		cookies := ""
		val,ok := value.(map[string]interface{})
		if ok {
			for k,v := range val {
				cookies += encodeURIComponent(k) + "=" + encodeURIComponent( fmt.Sprintf("%v", v) ) + "; "
			}
		}
		headers["Cookie"] = strings.TrimRight(cookies,"; ")
	} else if ok {
		headers["Cookie"] = options["cookie"]
	}
	if fmt.Sprintf("%v",options["crypto"]) == "weapi"{
		reg := regexp.MustCompile("_csrf=([^(;|$)]+)")
		value,ok:= headers["Cookie"]
		data,url = func() (map[string]interface {},string) {
			csrfTokenRaw := reg.FindStringSubmatch(fmt.Sprintf("%v",value))
			if ok && len(csrfTokenRaw) >= 2{
				data["csrf_token"] = reg.FindStringSubmatch(fmt.Sprintf("%v",value))[1]
			}else{
				data["csrf_token"] = ""
			}
			re := regexp.MustCompile(`\w*api`)
			return crypto.Weapi(data),re.ReplaceAllString(url,"weapi")
		}()
	}else if fmt.Sprintf("%v",options["crypto"]) == "linuxapi"{
		re := regexp.MustCompile(`\w*api`)
		data = crypto.LinuxApi(map[string]interface{}{
			"method":method,
			"url":re.ReplaceAllString(url,"api"),
			"params":data,
		})
		headers["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36"
		url = "https://music.163.com/api/linux/forward"
	} else if fmt.Sprintf("%v",options["crypto"]) == "eapi"{
		value,ok := options["cookie"].(map[string]interface{})
		cookie := map[string]interface{}{}
		if ok {
			cookie = value
		}

		csrfValue,isok := cookie["__csrf"]
		csrfToken := ""
		if isok {
			csrfToken = fmt.Sprintf("%v",csrfValue)
		}
		header := make(map[string]interface{})
		keys := [...]string{"osver","deviceId","mobilename","channel"}
		for _,val := range keys{
			value,ok := cookie[val]
			if ok{
				header[val] = value
			}
		}
		header["appver"] = func() string {
			val,ok := cookie["appver"]
			if ok{
				return fmt.Sprintf("%v",val)
			}
			return "6.1.1"
		}()
		header["versioncode"] = func() string {
			val,ok := cookie["versioncode"]
			if ok{
				return fmt.Sprintf("%v",val)
			}
			return "140"
		}()
		header["buildver"] = func() string {
			val,ok := cookie["buildver"]
			if ok{
				return fmt.Sprintf("%v",val)
			}
			return strconv.FormatInt(time.Now().Unix(),10)[0:10]
		}()
		header["resolution"] = func() string {
			val,ok := cookie["resolution"]
			if ok{
				return fmt.Sprintf("%v",val)
			}
			return "1920x1080"
		}()
		header["os"] = func() string {
			val,ok := cookie["os"]
			if ok{
				return fmt.Sprintf("%v",val)
			}
			return "android"
		}()
		header["requestId"] = fmt.Sprintf("%s_%s",strconv.FormatInt(time.Now().UnixNano()/1000000,10),
			pad.Left(strconv.Itoa(rand.New(rand.NewSource(time.Now().Unix())).Int() % 1000),4,"0" ))
		header["__csrf"] = csrfToken
		cookieMusicU,ok := cookie["MUSIC_U"]
		if ok{
			header["MUSIC_U"] = cookieMusicU
		}
		cookieMusicA,ok := cookie["MUSIC_A"]
		if ok{
			header["MUSIC_A"] = cookieMusicA
		}
		data,url = func() (map[string]interface {},string) {
			toRet := ""
			for key,val := range header{
				toRet += encodeURIComponent(key) + "=" + encodeURIComponent(fmt.Sprintf("%v", val)) + "; "
			}
			for _,val := range keys{
				toRet += encodeURIComponent(val) + "=" + "undefined" + "; "
			}
			headers["Cookie"] = strings.TrimRight(toRet,"; ")
			data["header"] = header
			re := regexp.MustCompile(`\w*api`)
			return crypto.EApi(options["url"].(string),data),re.ReplaceAllString(url,"eapi")
		}()
	}
	valBody := netUrl.Values{}
	for keyData,valData := range data {
		valBody.Add(keyData,fmt.Sprintf("%v",valData))
	}
	// It's not necessary here in go , go http request return away byte
	//if fmt.Sprintf("%v",options["crypto"]) == "eapi"{
	//	settings["encoding"] = nil
	//}
	// todo 增加proxy设置
	answer := map[string] interface{}{
		"status":500,
		"body":map[string]interface{}{},
		"cookie": []string{},
	}

	request , err:= http.NewRequest(method,url,bytes.NewBuffer([]byte(valBody.Encode())))
	if err != nil {
		answer["status"] = 502
		answer["body"] = map[string]interface{}{
			"code": 502,
			 "msg": err.Error(),
		}
		return answer
	}

	for key,value := range headers {
		request.Header.Set(key,fmt.Sprintf("%v",value))
	}
	client := http.Client{}
	resp,err := client.Do(request)
	if err != nil {
		answer["status"] = 502
		answer["body"] = map[string]interface{}{
			"code": 502,
			"msg": err.Error(),
		}
		return answer
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		answer["status"] = 502
		answer["body"] = map[string]interface{}{
			"code": 502,
			"msg": err.Error(),
		}
		return answer
	}
	re := regexp.MustCompile("\\s*Domain=[^(;|$)]+;*")
	if value,ok := resp.Header["Set-Cookie"];ok{
		var cookie []string
		for _,val := range value{
			cookie = append(cookie,re.ReplaceAllString(val,"") )
		}
		answer["cookie"] = cookie
		//cookieB ,_:= json.Marshal(cookie)
		//fmt.Printf("%s\n", pretty.Color(pretty.Pretty(cookieB), pretty.TerminalStyle))
	}
	if options["crypto"] == "eapi" {
		bodyZip := body
		zipReader, err := zip.NewReader(bytes.NewReader(bodyZip), int64(len(body)))
		// err occur
		if err != nil {
			bodyMapData := map[string]interface{}{}
			unmarshalErr := json.Unmarshal(body,&bodyMapData)
			if unmarshalErr !=nil {
				eapiKey := "e82ckenh8dichen8"
				bodyJson := map[string]interface{}{}
				unzippedFileBytesAes := crypto.AesDecryptECB(body,[]byte(eapiKey))
				if err := json.Unmarshal(unzippedFileBytesAes,&bodyJson);err==nil{
					answer["body"] = bodyJson
					if val,ok:=bodyJson["code"];ok{
						answer["status"] = val
					}else {
						answer["status"] =  resp.StatusCode
					}
				}
				return answer
			}
			answer["body"] = bodyMapData
			answer["status"] = resp.StatusCode
			return answer
		}
		for _, zipFile := range zipReader.File {
			fmt.Println("Reading file:", zipFile.Name)
			unzippedFileBytes, err := readZipFile(zipFile)
			if err != nil {
				log.Println(err)
				continue
			}
			eapiKey := "e82ckenh8dichen8"
			bodyJson := map[string]interface{}{}
			unzippedFileBytesAes := crypto.AesDecryptECB(unzippedFileBytes,[]byte(eapiKey))
			if err := json.Unmarshal(unzippedFileBytesAes,&bodyJson);err==nil{
				answer["body"] = bodyJson
				if val,ok:=bodyJson["code"];ok{
					answer["status"] = val
				}else {
					answer["status"] =  resp.StatusCode
				}
			}else{
				_ = json.Unmarshal(unzippedFileBytes,&bodyJson)
				answer["body"] = bodyJson
				answer["status"] =  resp.StatusCode
			}
			if answer["status"].(int) <= 100 || answer["status"].(int) >= 600{
				answer["status"] = 400
			}
			return answer
		}
	} else{
		bodyMapData := map[string]interface{}{}
		unmarshalErr := json.Unmarshal(body,&bodyMapData)
		if unmarshalErr !=nil {
			answer["body"] = string(body)
			answer["status"] = resp.StatusCode
			return answer
		}
		answer["body"] = bodyMapData
		if val,ok:=bodyMapData["code"];ok{
			if fmt.Sprintf("%v", val) == "502"{
				answer["status"] = 200
			}else{
				answer["status"] = val
			}
		}else {
			answer["status"] =  resp.StatusCode
		}
	}
	return answer
}