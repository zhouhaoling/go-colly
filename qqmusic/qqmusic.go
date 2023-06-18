package qqmusic

import (
	"bytes"
	"colly_music/dao"
	"colly_music/utils"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const letterBytes = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"

var singerD dao.QQDao
var songD dao.QQDao

func Colly_qqmusic() {
	url := "https://y.qq.com/n/ryqq/toplist/26"
	var songerList = []string{}
	c := colly.NewCollector(
		colly.Async(true),
	)

	siteCookie := c.Cookies(url)
	c.SetCookies(url, siteCookie)

	c.OnRequest(func(r *colly.Request) {
		//Request头部信息设定
		r.Headers.Set("User-Agent", utils.RandomString(letterBytes))
		r.Headers.Set("Host", "y.qq.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "https://y.qq.com")
		r.Headers.Set("Referer", "https://y.qq.com/n/ryqq/toplist/26")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")

		fmt.Printf("r.URL: %v\n", r.URL)
	})

	//对响应的HTML元素进行处理
	c.OnHTML("title", func(h *colly.HTMLElement) {
		fmt.Printf("h.Text: %v\n", h.Text)
	})

	c.OnHTML(".topList_mod_songlist", func(h *colly.HTMLElement) {
		//基础地址
		base_url := "https://y.qq.com"

		h.DOM.Find("#app > div > div.main > div.mod_toplist > div.topList_mod_songlist > ul.songlist__list > li").Each(func(i int, s *goquery.Selection) {
			var buf1 bytes.Buffer
			var buf2 bytes.Buffer

			//拿到热歌榜的歌曲排名、vip标志、歌手、时长、歌曲url
			top := s.Find(".songlist__number").Text()
			songName := s.Find(" li> div > div.songlist__songname > span").Text()
			song_url, _ := s.Find("li> div > div.songlist__songname > span >a:nth-child(2)").Attr("href")
			singer := s.Find("li > div > div.songlist__artist").Text()
			singer_url, _ := s.Find("li > div > div.songlist__artist > a").Attr("href")
			duration := s.Find("li > div > div.songlist__time").Text()
			vip, exists := s.Find(" li > div > div.songlist__songname > i").Attr("title")
			if exists != true {
				vip = "NO VIP"
				exists = true
			}

			fmt.Printf("top: %v, %v, %v, %v, %v\n", top, songName, vip, singer, duration)

			//地址拼接，得到可以访问的歌曲地址和歌手地址
			buf1.WriteString(base_url)
			buf1.WriteString(song_url)
			dsong_url := buf1.String()
			// fmt.Printf("songName_url: %v\n", dsong_url)

			buf2.WriteString(base_url)
			buf2.WriteString(singer_url)
			dsinger_url := buf2.String()
			for i := range songerList {
				if dsinger_url != songerList[i] {
					songerList = append(songerList, dsinger_url)
				}
			}
			// fmt.Printf("singer_url: %v\n", dsinger_url)

			//分别处理歌曲页面信息和歌手页面信息
			Detail_song(c, dsong_url, top, songName, duration, vip, singer)
			time.Sleep(2 * time.Second)
			Detail_singer(c, dsinger_url, singer)
			time.Sleep(2 * time.Second)
			//评论爬取
			Detail_comment(url)
			// fmt.Printf("top: %v, %v, %v, %v\n", top, songName, vip, singer)
		})
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("r.StatusCode: %v\n", r.StatusCode)
		if r.StatusCode != 200 {
			log.Println(r.StatusCode)
			return
		}
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second,
	})

	c.Visit("https://y.qq.com/n/ryqq/toplist/26")
	c.Wait()
}

