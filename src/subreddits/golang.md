# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][distribyted/distribyted: Torrent client with on-demand file downloading as a filesystem in Go](https://www.reddit.com/r/golang/comments/ji9zcb/distribyteddistribyted_torrent_client_with/)
- url: https://github.com/distribyted/distribyted
---

## [3][kboard: a terminal game to practice keyboard typing](https://www.reddit.com/r/golang/comments/jid5l4/kboard_a_terminal_game_to_practice_keyboard_typing/)
- url: https://github.com/CamiloGarciaLaRotta/kboard
---

## [4][go-echarts: the adorable charts library for Golang](https://www.reddit.com/r/golang/comments/jhu7xv/goecharts_the_adorable_charts_library_for_golang/)
- url: https://www.reddit.com/r/golang/comments/jhu7xv/goecharts_the_adorable_charts_library_for_golang/
---
Hi, every gopher, project recommendation!

In the Golang ecosystem, there are not many choices for data visualizing libraries. The development of [go-echarts](https://github.com/go-echarts/go-echarts) aims to provide a simple yet powerful data visualizing library for Golang. [Echarts](https://echarts.apache.org/) is an outstanding charting and visualizing library opensourced by Baidu, it supports adorable chart types and various interactive features. There are many language bindings for Echarts, for example, [pyecharts](https://github.com/pyecharts/pyecharts). go-echarts learns from pyecharts and has evolved a lot.


go-echarts is easy to use and extend,  in this example, we create a simple bar chart with only a few lines of code.


```golang
package main

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// generate random data for bar chart
func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i &lt; 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func main() {
	// create a new bar instance
	bar := charts.NewBar()

	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Bar-basic-example",
			Subtitle: "This is the subtitle.",
		}),
	)

	// iowriter
	f, err := os.Create("bar.html")
	if err != nil {
		panic(err)
	}

	// Put some data in instance
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())

	// Where the magic happens
	bar.Render(f)
}
```

And the generated bar.html is rendered as below

[bar.png](https://user-images.githubusercontent.com/19553554/97107442-85f91880-1702-11eb-8b73-d0d8daedf549.png)

For more information, please refer to [go-echarts/examples](https://github.com/go-echarts/examples) and the [GoDoc](https://pkg.go.dev/github.com/go-echarts/go-echarts/v2).

Project link: https://github.com/go-echarts/go-echarts
## [5][Using Golang to upload and download files in pcloud using cmd](https://www.reddit.com/r/golang/comments/jiacsd/using_golang_to_upload_and_download_files_in/)
- url: https://www.reddit.com/r/golang/comments/jiacsd/using_golang_to_upload_and_download_files_in/
---
Hello gophers,

I am trying to use [https://github.com/yanmhlv/pcloud](https://github.com/yanmhlv/pcloud) project to upload and download files.With this i can upload my files but downloading is not working.Also when using this i can upload small size of files but when it is more than 50mb it is taking more time.

Is there any other method by using golang to connect to pcloud in cmd.Can anybody help me ?
## [6][Generics in Go](https://www.reddit.com/r/golang/comments/jiehkx/generics_in_go/)
- url: https://youtu.be/F8Gl8-3ZW0E
---

## [7][Quick question: Go syntax changes since Go 1.0](https://www.reddit.com/r/golang/comments/jhzjno/quick_question_go_syntax_changes_since_go_10/)
- url: https://www.reddit.com/r/golang/comments/jhzjno/quick_question_go_syntax_changes_since_go_10/
---
For a presentation about Go's simplicity, I would like to know what changes/additions happened to the language/syntax since Go 1.0 (up until Go 1.15).

The only thing I know of is:  
* Type aliasing: `type T1 = T2`
## [8][Reactive planing in Golang. Reach a desired number adding and subtracting random numbers](https://www.reddit.com/r/golang/comments/jicwdt/reactive_planing_in_golang_reach_a_desired_number/)
- url: https://gianarb.it/blog/reactive-plan-golang-example
---

## [9][Am I wrong to use gRPC for this project?](https://www.reddit.com/r/golang/comments/ji9cfs/am_i_wrong_to_use_grpc_for_this_project/)
- url: https://www.reddit.com/r/golang/comments/ji9cfs/am_i_wrong_to_use_grpc_for_this_project/
---
Is gRPC only for microservice for you ?? I plan to use gRPC or GraphQL here but if I use GraphQL I should install many dependency in the Front End and would read more documentation for each dependency , but if I use gRPC web , it was generated code and I should follow the rules how to use it and not too much to configuration things for their dependency on the FE 

Also I was looking how to deploy gRPC to production, and many articles using Kurbenest for the case , I have no idea about Kubernestes :( , that is my problem if I stick with gRPC , can I deploy my go gRPC like usual project using only docker container and run that by docker ?

What you think about gRPC only for microservice ? Should I avoid to gRPC for this ? 
But so far I don't use for bi-direction on gRPC, and might use unary alot for so far
## [10][imgcat - a tool to output images as RGB ANSI graphics on the terminal](https://www.reddit.com/r/golang/comments/jhms24/imgcat_a_tool_to_output_images_as_rgb_ansi/)
- url: https://github.com/trashhalo/imgcat
---

## [11][A generic task queueing system for Go programs](https://www.reddit.com/r/golang/comments/jhvggg/a_generic_task_queueing_system_for_go_programs/)
- url: https://git.sr.ht/~sircmpwn/dowork
---

