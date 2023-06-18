package qqmusic

import (
	"colly_music/dao"
	"context"
	"log"
	"regexp"
	"time"

	"github.com/chromedp/chromedp"
)

var com dao.QQDao

func Detail_comment(url string) {
	// sel := "#comment_box > div:nth-child(3) > ul > li > div:nth-child(1) > p"
	sel := "#app > div > div.main > div.toplist_nav > dl:nth-child(1) > dd:nth-child(3) > a"
	param := `document.querySelector("#mod_comment")`
	src, _ := getHttpHtmlContent(url, sel, param)
	complieRegex := regexp.MustCompile("<span>(.*?)</span>")
	list := complieRegex.FindAllStringSubmatch(src, -1)
	b := com.InsertComment(list)
	if b != true {
		panic("插入失败！")
	}

}

// 获取网站上爬取的数据，动态数据
// htmlContent:上面的 html 页面信息，selector:第一步获取的 selector，sel:
func getHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s  调整等待页面加载时间
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 80*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.Sleep(10*time.Second), //等待10秒

		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return htmlContent, nil
}
