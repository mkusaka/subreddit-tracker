# golang
## [1][Why we’re writing machine learning infrastructure in Go, not Python](https://www.reddit.com/r/golang/comments/elzoma/why_were_writing_machine_learning_infrastructure/)
- url: https://towardsdatascience.com/why-were-writing-machine-learning-infrastructure-in-go-not-python-38d6a37e2d76
---

## [2][How to Start with Logging in Go Projects. Part 2](https://www.reddit.com/r/golang/comments/em8uiu/how_to_start_with_logging_in_go_projects_part_2/)
- url: https://blog.maddevs.io/how-to-start-with-logging-in-go-projects-part-2-1e57cc58f329
---

## [3][JWT Authorization in Golang - CloudJourney IO](https://www.reddit.com/r/golang/comments/elwvxm/jwt_authorization_in_golang_cloudjourney_io/)
- url: https://www.cloudjourney.io/articles/security/jwt_in_golang-su/
---

## [4][ditty - Terminal-based music player (beep/oto for audio)](https://www.reddit.com/r/golang/comments/em1et9/ditty_terminalbased_music_player_beepoto_for_audio/)
- url: https://git.sr.ht/~tslocum/ditty
---

## [5][Go Get My/Vulnerabilities](https://www.reddit.com/r/golang/comments/em9iwq/go_get_myvulnerabilities/)
- url: https://www.youtube.com/watch?v=GGQcv7fK0JY
---

## [6][bokeh for golang](https://www.reddit.com/r/golang/comments/em90f6/bokeh_for_golang/)
- url: https://www.reddit.com/r/golang/comments/em90f6/bokeh_for_golang/
---
Does such a tool like bokeh [https://docs.bokeh.org/en/latest/index.html](https://docs.bokeh.org/en/latest/index.html) exist for golang? I think go would be ideal because a single binary can host the webserver and data for visualization. At the moment, I worry about setting up python runtime, webserver, and a proxy which is tedious!
## [7][Help me understand goroutines](https://www.reddit.com/r/golang/comments/em5r4h/help_me_understand_goroutines/)
- url: https://www.reddit.com/r/golang/comments/em5r4h/help_me_understand_goroutines/
---
For some reason I just struggle with goroutines with Channels. In other languages like python or C I’ve had breakthrough moments where I thought “ah ha, that’s how I can use x design pattern or data structure.” For some reason, concurrency/go routines/channels aren’t really clicking for me.

So my question to the gophers is: What breakthrough helped you learn how to use channels and goroutines effectively? How can I get better at “thinking concurrently” if I come from a background of writing procedural and imperative code?
## [8][How to monitor if my Go program is operating optimally?](https://www.reddit.com/r/golang/comments/eltqmr/how_to_monitor_if_my_go_program_is_operating/)
- url: https://www.reddit.com/r/golang/comments/eltqmr/how_to_monitor_if_my_go_program_is_operating/
---
I have a little Go script running on a 16 vCPU Google Compute Engine.  


This script establishes about ~1,300 web sockets to [polygon.io](https://polygon.io/sockets) to monitor stock trades for particular tickers. Behind the scenes this uses [gorilla/websocket](https://github.com/gorilla/websocket). When a trade comes in, some calculations are performed, and the result is stored in a [concurrent-map](https://github.com/orcaman/concurrent-map).

According to Google's dashboard, my CPU utilization for this work is less than 1%. According to `iptraf`, my packet sizes are between 1 and 275 bytes.

Still, I'd like to know if my calculations are operating "fast" enough--let's say, completing a trade calculation within 100ms. I can't shake the feeling that I have a bottleneck somewhere. I believe this is the case because when the script runs "live," nothing happens; but when I reply the day's activities in a debug mode, events that I expect to occur, occur (eg. recommending when to buy a particular stock). How can I ensure CPU throughout or network latency in my Go program are not issues?

TIA
## [9][A Chapter in the Life of Go’s Compiler](https://www.reddit.com/r/golang/comments/elyykn/a_chapter_in_the_life_of_gos_compiler/)
- url: https://medium.com/samsara-engineering/a-chapter-in-the-life-of-gos-compiler-c89b9db74617
---

## [10][How to work with Go modules?](https://www.reddit.com/r/golang/comments/em0zxc/how_to_work_with_go_modules/)
- url: https://www.reddit.com/r/golang/comments/em0zxc/how_to_work_with_go_modules/
---
Hello,

I've started to learn about go modules and I've started with a simple Echo web server.

    // main.go
    package main
    
    import (
    	"net/http"
    
    	"github.com/labstack/echo"
    )
    
    func main() {
    	e := echo.New()
    	e.GET("/", func(c echo.Context) error {
    		return c.String(http.StatusOK, "Hello, World!")
    	})
    	e.Logger.Fatal(e.Start(":1323"))
    }

Using `go get -u github.com/labstack/echo/...` , I can run the above code from within my GOPATH and a server starts with version 4.1.13.

I have created a new folder for Go modules called /gomodules, which is outside my GOPATH. Inside /gomodules I created a folder called /echo (/gomodules/echo).

Inside /echo I created a go.mod file:

    // go.mod
    module testing.com/main
    
    go 1.13

I also copied the main.go file into the /echo folder.

When I run main.go, Go gets v3.3.10 of Echo. I had expected it to get the latest version. Here is the result.

    module testing.com/main
    
    go 1.13
    
    require (
    	github.com/labstack/echo v3.3.10+incompatible // indirect
    	github.com/labstack/gommon v0.3.0 // indirect
    	golang.org/x/crypto v0.0.0-20200108215511-5d647ca15757 // indirect
    )

I've tried manually setting the mod.go file.

    module testing.com/main
    
    go 1.13
    
    require (
    	github.com/labstack/echo v4.1.13 // indirect
    )

As well as:

    module testing.com/main
    
    go 1.13
    
    require (
    	github.com/labstack/echo/v4 v4.1.13 // indirect
    )
    

The later will run but the first wont.  Even after the later runs and the go.mod file looks the same, when running main.go, I am stuck with v3.3.10.

I've read the docs on modules and I feel lost as to how to get this to work. What am I missing?

&amp;#x200B;

Thanks!
