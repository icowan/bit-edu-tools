/**
 * @Time : 2019-09-03 17:14
 * @Author : solacowa@gmail.com
 * @File : main
 * @Software: GoLand
 */

package main

import (
	"flag"
	"fmt"
	"github.com/kplcloud/request"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	httpUrl       = "http://grdms.bit.edu.cn/yjs/dwr/call/plaincall/YYPYCommonDWRController.pyJxjhSelectCourse.dwr"
	cookieContent = "JSESSIONID=57AD326E5B9AD1C486178462B5922B50; td_cookie=2709571256; SECURITY_AUTHENTICATION_COOKIE=49a5e8d8afeeb30a21014b1ded4e1028e8fb15ac51091222e11db182775ad6db670da5c035c4534e; __jsluid_h=b54d83d5ded258b0e0edeb6001da2429; SECURE_AUTH_ROOT_COOKIE=49a5e8d8afeeb30a21014b1ded4e1028e8fb15ac51091222e11db182775ad6db670da5c035c4534e"
	httpLoginUrl  = "https://login.bit.edu.cn/cas/login?service=http%3a%2f%2fgrdms.bit.edu.cn%2fyjs%2flogin_cas.jsp"
)

var (
	fs        = flag.NewFlagSet("bit-edu-tools", flag.ExitOnError)
	username  = fs.String("username", "123123", "学号")
	password  = fs.String("password", "12312", "密码")
	cacheFile = "./cache.json"
)

func main() {

	login()
	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(cacheFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
	return

	var body = []byte(`callCount=1
page=/yjs/yanyuan/py/peiYangJiHua.do?method=initQuery
httpSessionId=57AD326E5B9AD1C486178462B5922B50
scriptSessionId=5BFD077D239E4FE341B2E3D6F644C9B5211
c0-scriptName=YYPYCommonDWRController
c0-methodName=pyGetKkqdCourse4Pyjhxk
c0-id=0
c0-param0=string:3420190175
c0-param1=string:
c0-e1=string:%e7%ae%a1%e7%90%86%e6%b2%9f%e9%80%9a
c0-param2=Object_Object:{t.kczwmc:reference:c0-e1}
batchId=0`)

	b, err := request.NewRequest(httpUrl, http.MethodPost).
		Header("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E)").
		Header("Host", "grdms.bit.edu.cn").
		Header("Referer", "http://grdms.bit.edu.cn/yjs/yanyuan/py/peiYangJiHua.do?method=initQuery").
		Header("Connection", "Keep-Alive").
		Header("Cookie", cookieContent).
		Body(body).Do().Raw()
	if err != nil {
		fmt.Println("err", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(b))
}

func login() {

	param := url.Values{}
	param.Add("username", *username)
	param.Add("password", *password)
	param.Add("lt", "LT-189540-cFbJ13QMl1YForYdf6MA3APnHkFQfy-1567503210392")
	param.Add("execution", "e2s1")
	param.Add("_eventId", "submit")
	param.Add("rmShown", "1")

	result := request.NewRequest(httpLoginUrl, http.MethodPost).
		Header("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3").
		Header("Accept-Language", "zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7,la;q=0.6").
		Header("Cookie", "JSESSIONID=0000DPjjrVXtTkBvl8JSjauzpr4:18bictvom").
		Header("Host", "login.bit.edu.cn").
		Header("Origin", "https://login.bit.edu.cn").
		Header("Referer", httpLoginUrl).
		Header("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E)").
		Header("Sec-Fetch-Mode", "navigate").
		Header("Sec-Fetch-Site", "same-origin").
		Header("Sec-Fetch-User", "?1").
		Header("Upgrade-Insecure-Requests", "1").
		Header("Content-Type", "application/x-www-form-urlencoded").
		Body([]byte(param.Encode())).Do()

	if result.HttpStatusCode() != http.StatusOK {
		fmt.Println(result.Error().Error())
		return
	}
	fmt.Println(result.Error())
	fmt.Println(result.HttpStatusCode())
	fmt.Println(result.Raw())

}
