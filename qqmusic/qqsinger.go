package qqmusic

import (
	"colly_music/model"
	"colly_music/utils"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// 歌手详情页处理
func Detail_singer(collector *colly.Collector, url string, singer_name string) {
	collector = collector.Clone()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("r.URL: %v\n", r.URL)
	})

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})

	//解析歌手详情页数据提取
	collector.OnHTML("body", func(h *colly.HTMLElement) {

		h.DOM.Find(".mod_data").Each(func(i int, s *goquery.Selection) {
			//简介
			content := s.Find("div.mod_data > div.data__cont > div.data__desc > div").Text()
			//单曲数
			siles := s.Find(" div.mod_data > div.data__cont > ul > li:nth-child(1) > a > strong").Text()
			//专辑数
			als := s.Find(" div.mod_data > div.data__cont > ul > li:nth-child(2) > a > strong").Text()
			//MV数
			ms := s.Find(" div.mod_data > div.data__cont > ul > li.data_statistic__item.data_statistic__item--last > a > strong").Text()
			//关注量
			attention := s.Find(" div.mod_data > div.data__cont > div.data__actions > a.mod_btn > span").Text()
			//图片地址
			//img, _ := s.Find("#app > div > div.main > div.mod_data > a > img").Attr("src")
			
			fmt.Printf("content: %v\n", content)
			singles, _ := strconv.Atoi(siles)
			albums, _ := strconv.Atoi(als)
			mvs, _ := strconv.Atoi(ms)
			//数据提取将关注量转为float类型
			attentions := utils.Tofloat(attention)
			singer := &model.Singer{
				Name:      singer_name,
				Content:   content,
				Singles:   singles,
				Albums:    albums,
				Mvs:       mvs,
				Attention: attentions,
			}

			//将数据插入数据库中
			err := singerD.InsertSinger(singer)
			if err != nil {
				//插入失败，程序暂停
				log.Panic(err)
			}
			time.Sleep(time.Second * 2)

		})
	})

	collector.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Println(r.StatusCode)
			return
		}
	})
	collector.OnScraped(func(r *colly.Response) {
		// fmt.Println("结束", r.Request.URL)
	})
	collector.Visit(url)
	collector.Wait()
}
