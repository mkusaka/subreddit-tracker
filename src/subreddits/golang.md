# golang
## [1][Wrote an article/tutorial on Concurrency. Beginner Friendly!](https://www.reddit.com/r/golang/comments/j1xij2/wrote_an_articletutorial_on_concurrency_beginner/)
- url: https://medium.com/@yashaswi_nayak/go-a-tale-of-concurrency-a-beginners-guide-b8976b26feb
---

## [2][A simple and lightweight Static Site Generator written in Go](https://www.reddit.com/r/golang/comments/j1y78h/a_simple_and_lightweight_static_site_generator/)
- url: https://www.reddit.com/r/golang/comments/j1y78h/a_simple_and_lightweight_static_site_generator/
---
Hi all, I'm happy to announce my new Static Site Generator! It is designed for Markdown-based content with focus on simplicity and performance. Some key features are:  
 

* Flexible templating through an own theming system.
* Interchangeable themes due to pre-defined template variables.
* Local development without a third-party webserver.
* Creating projects &amp; themes within a single command.
* High performance, parallelized site builds.
* Optional RSS feeds and overview pages for tags.
* Auto-generated, overwritable overview pages for all directories.
* Pre-build hooks for generating assets.
* Simple configuration via YAML, TOML or JSON in a central file.

Check it out: [https://github.com/verless/verless](https://github.com/verless/verless)

verless is in active development and pre-build hooks specifically for themes are coming soon, but the public API already is stable.
## [3][A not so simple TCP socket implementation of TCP server/client ?](https://www.reddit.com/r/golang/comments/j1wry6/a_not_so_simple_tcp_socket_implementation_of_tcp/)
- url: https://www.reddit.com/r/golang/comments/j1wry6/a_not_so_simple_tcp_socket_implementation_of_tcp/
---
Im planning to build a TCP server &amp; client in Go where the focus is on speed, efficiency, stable connections ( recognize drops, reconnects, error handling ) and a proper concurrency model that can be scaled well, my plan is to try to get protocol implementation done via streaming protobuf messages. 

I know some have jumped into big products with Go, and theres NatsIO etc out there and that we normally say Go is perfect for networking and concurrency.

But in my search trying to see 'how others done it" its always 'a simple server' implementation - i remember from my CPP days that there was alot of practices that needed to be followed to achieve fast and stable performance and best practice.

Anyone have some directions to finding some information on some proper practices and some more complete approaches to handling all levels of socket communication in Go - ex the proper reconnect/connection dropped/connection disconnected ? the proper model of for ex handling reads/write on sockets ? 

Any books, totourials, projects, approaches you know of you could help me point into the right direction or would it be smarter just to stick to CPP ?
## [4][Data structures and algorithms in go](https://www.reddit.com/r/golang/comments/j1t0ob/data_structures_and_algorithms_in_go/)
- url: https://www.reddit.com/r/golang/comments/j1t0ob/data_structures_and_algorithms_in_go/
---
Hi guys.

I want to learn data structures and algorithms and implement all of them in Go. Any good youtube playlist or book you can refer for the same

I just have been procrastinating all my life for wanting to learn Data structures and algorithms.started off with java and now have moved to golang. I just want to stop this procrastination and finally start working on it.

Thanks
Demo
## [5][Do you use gccgo?](https://www.reddit.com/r/golang/comments/j1g1z6/do_you_use_gccgo/)
- url: https://www.reddit.com/r/golang/comments/j1g1z6/do_you_use_gccgo/
---
Hi,

I recently tried gccgo on a small project and I am surprised how small the executables can become. Standard Go compiler: almost 7 MiB, gccgo: 220 KiB. And it runs out of the box, no hassle. So why and when should you use the standard Go compiler at all?
## [6][A way to fetch human-readable string from binary file in Golang?](https://www.reddit.com/r/golang/comments/j1xpad/a_way_to_fetch_humanreadable_string_from_binary/)
- url: https://www.reddit.com/r/golang/comments/j1xpad/a_way_to_fetch_humanreadable_string_from_binary/
---
Hi all! I'm working on a small program that at some point needs to read a binary file and fetch all the human-readable text from it, basically the equivalent of the \`stings\` bash command. Is there any built in function / way to do so that I'm missing? All the solutions I found so far include reading the file, but I couldn't find a way to distinguish binary data to human readable string. 

Any help would be appreciated.
## [7][AppEngine with Go issue deploying](https://www.reddit.com/r/golang/comments/j1ydle/appengine_with_go_issue_deploying/)
- url: https://stackoverflow.com/questions/64119757/appengine-deploy-cannot-find-go-packages
---

## [8][Optimize your development workflow on Kubernetes](https://www.reddit.com/r/golang/comments/j1xrod/optimize_your_development_workflow_on_kubernetes/)
- url: https://www.reddit.com/r/golang/comments/j1xrod/optimize_your_development_workflow_on_kubernetes/
---
This blog post explains how to optimize your development workflow when building Go applications on Kubernetes with okteto:

[https://okteto.com/blog/how-to-develop-go-apps-in-kubernetes/](https://okteto.com/blog/how-to-develop-go-apps-in-kubernetes/)
## [9][Different ways to use environment variables in Golang](https://www.reddit.com/r/golang/comments/j1cxw4/different_ways_to_use_environment_variables_in/)
- url: https://www.loginradius.com/engineering/blog/environment-variables-in-golang/
---

## [10][Gio Simple GUI](https://www.reddit.com/r/golang/comments/j1utsd/gio_simple_gui/)
- url: https://www.reddit.com/r/golang/comments/j1utsd/gio_simple_gui/
---
Hello,

I'm trying to create a simple GUI with Gio, but I have one issue.

I want a label with a text field next to it (not below it), but for some reason the text field only has the width of the hint. Below is the code. How can I use the entire window width for the text field?

```
// SPDX-License-Identifier: Unlicense OR MIT

package main

// A simple Gio program. See https://gioui.org for more information.

import (
	"log"
	"os"
	"image/color"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/unit"
	"gioui.org/text"

	"gioui.org/font/gofont"
)

var(
	labelInvoer = "Input:"
	labelField1 = "Field 1:"
	lineEditor = &amp;widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	button = new(widget.Clickable)
	list = &amp;layout.List{
		Axis: layout.Vertical,
	}
)

func main() {
	go func() {
		minSize := app.MinSize(unit.Dp(300),unit.Dp(300))
		maxSize := app.MaxSize(unit.Dp(600),unit.Dp(600))
		size := app.Size(unit.Dp(400),unit.Dp(400))
		title := app.Title("Test123")
		
		w := app.NewWindow(size,minSize, maxSize, title)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := &lt;-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&amp;ops, e)
			
			widgets := []layout.Widget{
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Alignment: layout.Start}.Layout(gtx,
						layout.Rigid(
							material.H4(th, labelInvoer).Layout,
						),
					)
				},
				func(gtx layout.Context) layout.Dimensions {					
					return layout.Flex{Alignment: layout.Start}.Layout(gtx,
						layout.Rigid(
							material.H6(th, labelField1).Layout,
						),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: unit.Dp(18), Top: unit.Dp(0)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								e := material.Editor(th, lineEditor, "Hint")
								e.Font.Style = text.Italic
								border := widget.Border{Color: color.RGBA{A: 0xff}, CornerRadius: unit.Dp(8), Width: unit.Px(2)}
					
								return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
								})
							})
						}),		
					)
				},
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Alignment: layout.Start}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, button, "Click me!").Layout(gtx)
						}),
					)
				},
			}
			
			list.Layout(gtx, len(widgets), func(gtx layout.Context, i int) layout.Dimensions {
				return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
			})
			
			e.Frame(gtx.Ops)
		}
	}
}
```

Any tips are welcome, thx!
