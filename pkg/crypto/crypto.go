package crypto

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const iv = "0102030405060708"
const presetKey = "0CoJUm6Qyw8W8jud"
const linuxapiKey = "rFgB&h#%2?^eDg:Q"
const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const publicKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----"
const eapiKey = "e82ckenh8dichen8"
const secretKeyLen = 16

func AesEncrypt(buffer []byte, mode string, key []byte, iv []byte) []byte {
	switch mode {
	case "cbc":
		return AesEncryptCBC(buffer, key, iv)
	case "ecb":
		return AesEncryptECB(buffer, key)
	}
	return [] byte{}
}

func RsaEncrypt(buffer []byte, pkKey []byte) []byte {
	data := make([]byte, 128-len(buffer))
	buffer = append(data, buffer...)
	return RsaEncryptRaw(buffer, pkKey)
}

func Weapi(object map[string]interface{}) map[string]interface{} {
	text, _ := json.Marshal(object)
	buf := make([]byte, secretKeyLen)
	_, _ = rand.Read(buf)
	secretKey := make([]byte, secretKeyLen)
	reverseSecretKey := make([]byte, secretKeyLen)
	for index, val := range buf {
		secretKey[index] = byte(base62[val%62])
		reverseSecretKey[secretKeyLen-index-1] = byte(base62[val%62])
	}
	params_ := []byte(
		base64.StdEncoding.EncodeToString(
			AesEncrypt(text,"cbc",[]byte(presetKey),[]byte(iv))))
	params := base64.StdEncoding.EncodeToString(
		AesEncrypt(params_,"cbc",secretKey,[]byte(iv)))
	encSecKey := hex.EncodeToString(RsaEncrypt(reverseSecretKey,[]byte(publicKey)))
	return map[string]interface{}{
		"params":params,
		"encSecKey":encSecKey,
	}
}

func LinuxApi(object interface{}) map[string]interface{}{
	text,_ := json.Marshal(object)
	eparams := hex.EncodeToString(AesEncrypt(text,"ecb",[]byte(linuxapiKey),[]byte{}))
	return map[string]interface{}{
		"eparams":eparams,
	}
}

func EApi(url string , object interface{}) map[string]interface{}{
	var text string
	switch reflect.ValueOf(object).Kind().String() {
	case "string":
		text = object.(string)
		break
	case "map":
		val,_:=json.Marshal(object)
		text = string(val)
		break
	}
	message := fmt.Sprintf("nobody%suse%smd5forencrypt",url,text)
	md := md5.New()
	md.Write([]byte(message))
	digest := hex.EncodeToString(md.Sum(nil))
	data := fmt.Sprintf("%s-36cd479b6b5-%s-36cd479b6b5-%s",url,text,digest)
	params := hex.EncodeToString(AesEncrypt([]byte(data),"ecb",[]byte(eapiKey),[]byte{}))
	params = strings.ToUpper(params)
	return map[string]interface{}{
		"params":params,
	}
}



