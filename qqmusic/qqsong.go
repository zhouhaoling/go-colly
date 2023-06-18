package qqmusic

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// 歌曲详情页处理
func Detail_song(collector *colly.Collector, url string, top string, songName string, duration string, vip string, singer string) {
	//存放数据
	song_map := make(map[string]string)
	collector = collector.Clone()

	collector.OnRequest(func(r *colly.Request) {
	})

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})
	collector.OnHTML("#lrc_content", func(h *colly.HTMLElement) {
		ret, _ := h.DOM.Html()
		fmt.Printf("ret: %v\n", ret)
	})
	//解析歌曲详情页数据提取
	collector.OnHTML("body", func(h *colly.HTMLElement) {

		h.DOM.Find(".mod_data").Each(func(i int, s *goquery.Selection) {
			//专辑名
			album_name := s.Find("div.mod_data > div > ul > li:nth-child(1) > a").Text()
			//语种
			la := s.Find("div.mod_data > div > ul > li:nth-child(2)").Text()
			//流派
			ge := s.Find("div.mod_data > div > ul > li:contains(流派：)").Text()
			//发行时间
			isdt := s.Find("div.mod_data > div > ul > li:contains(发行时间：)").Text()
			language := la[10 : len(la)-1]
			issue_date := isdt[15 : len(isdt)-1]
			/* fmt.Printf("album_Name:%v\n", album_name)
			fmt.Printf("language:%v\n", language) */
			if ge != "" {
				genre := ge[9 : len(ge)-1]
				if genre[:1] == " " {
					song_map["Genre"] = genre[1:]
				}
				song_map["Genre"] = genre
				fmt.Printf("genre:%v\n", genre)
			} else {
				genre := "无流派"
				song_map["Genre"] = genre
				// fmt.Printf("ge: %v\n", genre)
			}
			// fmt.Printf("issue_date:%v\n", issue_date)
			song_map["Name"] = songName
			song_map["Top"] = top
			song_map["Duration"] = duration
			song_map["Vip"] = vip
			song_map["Album_name"] = album_name
			song_map["Language"] = language
			song_map["Issue_date"] = issue_date
			song_map["SingerName"] = singer
			err := songD.InsertSong(song_map)
			if err != nil {
				log.Panic(err)
			}

		})
	})

	collector.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Println(r.StatusCode)
			return
		}
	})

	collector.OnScraped(func(r *colly.Response) {
	})

	collector.Visit(url)
	collector.Wait()

}
