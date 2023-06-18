package utils

import (
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// 可视化数据模块
func CreatePieChart(title string, m map[string]int64, filename string) {
	// 新的饼图
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    title,
				Subtitle: "数据来源于数据库",
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Monthly revenue",
		generatePieItems(m)).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: 200,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      true,
					Formatter: "{b}: {c}",
				},
			),
		)
	f, _ := os.Create("./tempates/" + filename)
	_ = pie.Render(f)
}

func generatePieItems(m map[string]int64) []opts.PieData {
	// subjects := []string{"Maths", "English", "Science", "Computers", "History", "Geography"}
	items := make([]opts.PieData, 0)
	for key, v := range m {
		items = append(items, opts.PieData{
			Name:  key,
			Value: v})
	}
	return items
}

func CreateWordCloud(title string, m map[string]int64, filename string) {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    title,
			Subtitle: "数据来源于数据库",
		}))
	wc.AddSeries("刷新", generateWordCloudData(m)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{40, 80},
					Shape:     "cardioid",
				}),
		)
	f, _ := os.Create("./tempates/" + filename)
	_ = wc.Render(f)
}

// 词云图
func generateWordCloudData(data map[string]int64) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func CreateBarChart(m map[string]int64) {
	bar := charts.NewBar()
	var singers = []string{}
	items := make([]opts.BarData, 0)
	for key, v := range m {
		singers = append(singers, key)
		items = append(items, opts.BarData{Value: v})
	}

	time.Sleep(3 * time.Second)
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "歌手出现的次数",
			Subtitle: "数据来源于数据库",
		}),
	)
	bar.SetXAxis(singers).
		AddSeries("歌手数据展示", items)
	f, _ := os.Create("./tempates/bar.html")
	_ = bar.Render(f)
}
