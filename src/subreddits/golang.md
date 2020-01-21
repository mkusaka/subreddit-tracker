# golang
## [1][/r/golang hit 100k subscribers yesterday](https://www.reddit.com/r/golang/comments/er8klk/rgolang_hit_100k_subscribers_yesterday/)
- url: https://redditmetrics.com/r/golang
---

## [2][giu - cross platform rapid GUI framework for golang based on imgui](https://www.reddit.com/r/golang/comments/er7js0/giu_cross_platform_rapid_gui_framework_for_golang/)
- url: https://www.reddit.com/r/golang/comments/er7js0/giu_cross_platform_rapid_gui_framework_for_golang/
---
[github.com/AllenDang/giu](https://github.com/AllenDang/giu)

Cross platform rapid GUI framework for golang based on [imgui](https://github.com/ocornut/imgui) and the great golang binding [imgui-go](https://github.com/inkyblackness/imgui-go).

# Features

Compare to other imgui golang bindings, giu has following features:

* Live-update during the resizing of OS window (implemented on GLFW 3.3 and OpenGL 3.2).
* Redraw only when user event occurred. Costs only 0.5% CPU usage with 60FPS.
* Declarative UI (see examples for more detail).
* Drop in usage, no need to implement render and platform.
* Freetype font rendering support.
* OS clipboard support.

https://preview.redd.it/twjsquplvub41.png?width=1200&amp;format=png&amp;auto=webp&amp;s=2ed4bdf30b1f38bb85c7046b25c9f137577c9665
## [3][TutorialEdge/Go: A list of Go Programming Tutorials, their write-ups, their source code and their current build status!](https://www.reddit.com/r/golang/comments/er3dsy/tutorialedgego_a_list_of_go_programming_tutorials/)
- url: https://github.com/TutorialEdge/go
---

## [4][Algorithms with Go (free)](https://www.reddit.com/r/golang/comments/eqvurh/algorithms_with_go_free/)
- url: http://algorithmswithgo.com/
---

## [5][What things are necessary before being able to work on Go in the Go Team.](https://www.reddit.com/r/golang/comments/erc3qy/what_things_are_necessary_before_being_able_to/)
- url: https://www.reddit.com/r/golang/comments/erc3qy/what_things_are_necessary_before_being_able_to/
---
Currently I'm an undergraduate student in Computer Science and had gotten into Go and have loved it ever since. I was wondering, if later down the line I ever wanted to join the Go Team, what do you think would be the appropriate prerequisites/background to have? Are there any computer science/math areas that would be useful to know?
## [6][S-expressions parser (play with LISP-like expressions!)](https://www.reddit.com/r/golang/comments/er8p8o/sexpressions_parser_play_with_lisplike_expressions/)
- url: https://github.com/xiam/sexpr
---

## [7][Slasher-write http servers that respond to slack slash commands](https://www.reddit.com/r/golang/comments/er6vpo/slasherwrite_http_servers_that_respond_to_slack/)
- url: https://www.reddit.com/r/golang/comments/er6vpo/slasherwrite_http_servers_that_respond_to_slack/
---
## Slasher

[Slasher](https://github.com/autom8ter/slasher) makes it easy to write http servers that respond to slack slash commands. Please note, a custom slack app must be created within your target workspace. Visit the following link to learn how to setup a custom slack app: [https://api.slack.com/interactivity/slash-commands](https://api.slack.com/interactivity/slash-commands)

## Example

&amp;#x200B;

    import (
    	"fmt"
    	"github.com/autom8ter/slasher"
    	"github.com/nlopes/slack"
    	"github.com/pkg/errors"
    	"net/http"
    	"os"
    )
    
    
    func helloWorld() slasher.HandlerFunc{
    	return func(s *slasher.Slasher, client *slack.Client, command *slack.SlashCommand) (i interface{}, err error) {
    		script := `echo "hello world!"`
    		output, err := s.ShellScript(script)
    		if err != nil {
    			return nil,  errors.Wrapf(err, "failed to run script: %s", script)
    		}
    		return &amp;slack.Message{
    			Msg: slack.Msg{
    				Text: string(output),
    			},
    		}, nil
    	}
    }
    
    func main() {
    	slash := slasher.NewSlasher(os.Getenv("SLACK_TOKEN"), []string{"autom8ter"})
    	slash.AddHandler("/hello-world", helloWorld())
    	mux := http.NewServeMux()
    	mux.Handle("/slasher", slash.HandlerFunc())
    	if err := http.ListenAndServe(":8080", mux); err != nil {
    		fmt.Printf("failed to start server: %s", err.Error())
    		os.Exit(1)
    	}
    }
## [8][🔌 Fiber is an Expressjs inspired web framework build on 🚀 Fasthttp.](https://www.reddit.com/r/golang/comments/er8qjw/fiber_is_an_expressjs_inspired_web_framework/)
- url: https://www.reddit.com/r/golang/comments/er8qjw/fiber_is_an_expressjs_inspired_web_framework/
---
[github.com/gofiber/fiber](https://github.com/gofiber/fiber)

I've been developing micro web services in node for years, personally and professionally. In 2019 I faced some challenges regarding the performance for a high load micro service on a budget. This was the moment I had to seek for alternatives and porting the project to Go was the first thing that came in mind. Go always intrigued me since I saw Rob Pike gave [this](https://www.youtube.com/watch?v=PAAkCSZUG1c) talk, I recommend watching it if you didn't see it yet.

After porting the node code to Go, I observed three things:

* The community is awesome and so helpful!
* Performance, performance, performance!
* I fell in 💖 with Go and knew this was the beginning of a new chapter

After debugging and bootstrapping my micro service, I decided to rewrite all my projects in Go.This took me some some months, but I thought it was a great way to learn more about the language.

Most of my micro services are mostly build up out of API endpoints, and I always used [Expressjs](https://expressjs.com/) for fast prototyping and ease of use. [NGINX](https://www.nginx.com/) for load balancing and in most cases this made me able to deploy multiple services in a week. This worked seamlessly for me and the companies I worked for, time is 💸 money.

* Was it perfect code?       **No**
* Was it error proof?          **No**
* Was it efficient?                **No**
* Did it get the job done?  **Yes**

All my current projects that I ported to Go are using [Fasthttp](https://github.com/valyala/fasthttp), an amazing package that has some great people that helped me to understand some fundamentals regarding the network part in Go.

I tried to convince some of my co-workers to also switch to Go, some tried but had a hard time with the learning curve. This was my motivation to create a package for them that makes it easier porting [Expressjs](https://expressjs.com/) related projects to Go. In the hope that the performance gain in the end would made them fall in 💖 with Go too. Thats when [Fiber](https://github.com/gofiber/fiber) was born with some extended [Documentation](https://gofiber.github.io/fiber/#/) to help my colleagues out.

It's basically an [Expressjs](https://expressjs.com/) styled router build on top of [Fasthttp](https://github.com/valyala/fasthttp), I made sure almost all express methods are include ([Context API)](https://gofiber.github.io/fiber/#/context). It worked, my colleagues ported some node projects within a day and fell in love on both the language and performance. More gophers, mission complete 💪.

My plan this year is to deploy all our projects in production using Go. I think there is a lot of improvement to be made to this new library, that's one of the reasons I wanted to post it here. I would love to get some feedback, suggestions, pull requests or even criticism!

I would like to take the opportunity to thank some people for helping and inspiring me: [@valyala](https://github.com/valyala)([Fasthttp](https://github.com/valyala/fasthttp/blob/master/LICENSE)) / [@julienschmidt](https://github.com/julienschmidt)([HttpRouter](https://github.com/julienschmidt/httprouter)) / [@savsgio](https://github.com/savsgio)([Atreugo](https://github.com/savsgio/atreugo/blob/master/LICENSE)) / [@vincentLiuxiang](https://github.com/vincentLiuxiang)([Lu](https://github.com/vincentLiuxiang/lu/blob/master/LICENSE)) / [@tj](https://github.com/tj)([Expressjs](https://github.com/expressjs/express)) / [@pillarjs](https://github.com/pillarjs)([Path-to-regexp](https://github.com/pillarjs/path-to-regexp/blob/master/LICENSE)) / [@techempower](https://github.com/TechEmpower)([FrameworkBenchmarks](https://github.com/TechEmpower/FrameworkBenchmarks)) / [@erikdubbelboer](https://github.com/erikdubbelboer)([Fasthttp](https://github.com/valyala/fasthttp/blob/master/LICENSE))

Thanks all! 💖
## [9][Starting Golang: Implementing GNU Tee](https://www.reddit.com/r/golang/comments/ercwcd/starting_golang_implementing_gnu_tee/)
- url: https://medium.com/@sddkal/starting-golang-implementing-gnu-tee-49d33fd60028
---

## [10][Hey all , How to secure the data communicated over a websockets ?](https://www.reddit.com/r/golang/comments/ercja3/hey_all_how_to_secure_the_data_communicated_over/)
- url: https://www.reddit.com/r/golang/comments/ercja3/hey_all_how_to_secure_the_data_communicated_over/
---
I'm planning to use a encryption method like AES or something . Could anyone please help me to follow a good approach ?   


thanks
