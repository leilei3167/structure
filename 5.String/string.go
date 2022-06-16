package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/pkg/errors"
)

/*
串的抽象数据类型:
串和线性表在逻辑结构方面很类似,串针对的是字符集,每个元素都是字符,也因此串的基本操作就和线性表有区别
着重于对子串的操作,如查找位置,替换子串等
详见Go的strings包和bytes包

*/
/* 朴素模式匹配 */
/* KMP匹配算法 */
type WebData struct {
	Url           string
	Status        string
	ContentLength string
	//Proto         string
	Title   string      //网页标题
	Icon    string      //网页图标
	Headers http.Header //响应头,对于web端口,将Server作为PortInfo的Service?
	//Dir       []string //用gobuster扫描目录
	//MetaData string //网页源代码
}

func Catch(url string) (WebData, error) {
	var web WebData

	c := colly.NewCollector(colly.MaxDepth(1))
	c.Limit(&colly.LimitRule{
		RandomDelay: 5,
	})

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.WithTransport(
		&http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   2 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:    100,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //暂时不收集证书信息
		})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		web.Title = e.Text
	})
	c.OnHTML("link:contains(.icon)", func(e *colly.HTMLElement) {
		web.Icon = e.Text
	})

	c.OnResponse(func(r *colly.Response) {
		web.Url = r.Request.URL.String()
		web.Title = r.Headers.Get("Server") //html中如果有title会覆盖
		web.Headers = *r.Headers
		web.Status = strconv.Itoa(r.StatusCode)
		web.ContentLength = strconv.Itoa(len(r.Body))

	})

	err := c.Visit(url)
	if err != nil {
		return WebData{}, errors.Wrapf(err, "Visit %v 出错", url)
	}
	if web.Title == "" {
		web.Title = "none"
	}
	return web, nil
}

func main() {
	s, err := Catch("https://www.jianshu.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", s)

}
