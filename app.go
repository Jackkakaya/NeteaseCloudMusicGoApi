package main

import (
	"NeteaseCloudGoApi/models"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/url"
	"reflect"
	"strings"
	"time"
)

var ModelPath = "./models"

func InitRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	var musicObain models.MusicObain
	store := persistence.NewInMemoryStore(time.Second)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(RequestModifyMiddleware)
	files, err := ioutil.ReadDir(ModelPath)
	if err!=nil{
		log.Fatalln(err)
	}
	filter := map[string] bool {
		"music_obtain.go":true,
	}
	for _,file := range files{
		if filename := file.Name();filter[filename] == false {
			filename = strings.TrimSuffix(filename,".go")
			path := fmt.Sprintf("/%v",strings.ReplaceAll(filename,"_","/"))
			method := strings.ReplaceAll( strings.Title(strings.ReplaceAll(filename,"_"," ")) ," ","")
			r.Any(path, cache.CachePage(store, 2*time.Minute, func(context *gin.Context) {
				query := map[string]interface{}{}
				data, _ := ioutil.ReadAll(context.Request.Body)
				_=json.Unmarshal(data,&query)
				CookieParseMiddleware(context,&query)
				for key,val :=range context.Request.URL.Query(){
					query[key] = val[0]
				}
				for key,val := range context.Request.PostForm{
					query[key] = val
				}
				respRaw := reflect.ValueOf(&musicObain).MethodByName(method).Call([]reflect.Value{reflect.ValueOf(query)})
				resp := respRaw[0].Interface().(map[string]interface{})
				for _,val := range resp["cookie"].([]string) {
					context.Writer.Header().Add("Set-Cookie",val)
				}
				context.JSON(200,resp)
			}))
		}
	}
	return r
}

func RequestModifyMiddleware(context *gin.Context)  {
	if context.FullPath() != "/" && !strings.Contains(context.FullPath(),".") {
		context.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		context.Writer.Header().Set("Access-Control-Allow-Headers'","X-Requested-With,Content-Type")
		context.Writer.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
		context.Writer.Header().Set("Content-Type","application/json; charset=utf-8")
		if context.Request.Header.Get("origin") != ""{
			context.Writer.Header().Set("Access-Control-Allow-Origin",context.Request.Header.Get("origin"))
		}else {
			context.Writer.Header().Set("Access-Control-Allow-Origin","*")
		}
	}
	if context.Request.Method == "OPTIONS" {
		context.Status(204)
		context.Done()
	}
	context.Next()
}

func CookieParseMiddleware(context *gin.Context,queryRaw interface{}) {
	query := queryRaw.(*map[string]interface{})
	cookieMap := map[string]interface{}{}
	for _,val := range context.Request.Cookies(){
		cookieKey ,_ := url.PathUnescape(val.Name)
		cookieVal,_ := url.PathUnescape(val.Value)
		cookieMap[cookieKey] = cookieVal
	}
	(*query)["cookie"] = cookieMap
}

func main() {
	routersInit := InitRouter()
	if err := routersInit.Run(":8080");err==nil{
		log.Printf("[info] start http server listening %s\n")
	}else{
		log.Fatal("[Error]",err)
	}
}
