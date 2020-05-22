# golang
## [1][10 part series to build web app with Go](https://www.reddit.com/r/golang/comments/gobfe7/10_part_series_to_build_web_app_with_go/)
- url: https://www.gophersumit.com/series/web/
---

## [2][How to do something every 5 minutes in Go?](https://www.reddit.com/r/golang/comments/gododv/how_to_do_something_every_5_minutes_in_go/)
- url: https://www.reddit.com/r/golang/comments/gododv/how_to_do_something_every_5_minutes_in_go/
---
I feel like it's a dummy question, but I want to ask it anyway since I do not have much experience with concurrent programming.

In the app I am currently working on, someone wrote a code that should do something (in this case, printing "TICK!") every, let's say 5 minutes:

    go func() {
        for {
            fmt.Println("TICK!")
            time.Sleep(5 * time.Minute)
        }
    }()

However, from what I have read, the recommended approach is:

    func CallTickerInBackground(done &lt;-chan bool) {
        ticker := time.NewTicker(5 * time.Minute)
        go func() {
            for {
            	select {
            	case &lt;-done:
            		logrus.Info("Stopping ticker")
            		ticker.Stop()
            		return
            	case &lt;-ticker.C:
            		fmt.Println("TICK!")
            	}
            }
        }()
    }

The advantages of the latter are that we have more stable time intervals (instead of processing time + waiting time) and we have control over stopping the background processing. On the other hand, right now a separate channel to stop it needs to be created, so it's slightly more complex.

So which approach is better, and why (or under what circumstances)?
## [3][50+ Technical Interview Problems Solved in Go](https://www.reddit.com/r/golang/comments/gnxfby/50_technical_interview_problems_solved_in_go/)
- url: https://www.reddit.com/r/golang/comments/gnxfby/50_technical_interview_problems_solved_in_go/
---
I have been solving various technical interview problems and coding challenges in Go. The problems vary from simple “Reverse a String” to more involving ones like “Implement the A* Algorithm”

If you have any feedback and/or some suggestions for new problems to solve, I would appreciate it as my goal is to grow this collection and help anyone interested in learning and/or practicing before their interview 

https://github.com/shomali11/go-interview
## [4][GUI packages for Go](https://www.reddit.com/r/golang/comments/gojais/gui_packages_for_go/)
- url: https://golangr.com/gui/
---

## [5][I want to use video streaming in my website , however I got this error "media Stream Error "](https://www.reddit.com/r/golang/comments/goj6jr/i_want_to_use_video_streaming_in_my_website/)
- url: https://www.reddit.com/r/golang/comments/goj6jr/i_want_to_use_video_streaming_in_my_website/
---
 [https://github.com/ali2210/WizDwarf/blob/master/js/webrtc-video.js](https://github.com/ali2210/WizDwarf/blob/master/js/webrtc-video.js) 

This is the code which i find out at youtube
## [6][Replacing net.DefaultResolver with a caching DNS over TLS/HTTPS resolver](https://www.reddit.com/r/golang/comments/gofxbv/replacing_netdefaultresolver_with_a_caching_dns/)
- url: https://www.reddit.com/r/golang/comments/gofxbv/replacing_netdefaultresolver_with_a_caching_dns/
---
DNS caching has been discussed multiple times in the past. The consensus seems to be that Go won't go there: [github.com/golang/go/issues/24796](https://github.com/golang/go/issues/24796#issuecomment-383716244)

I've seen a few DNS caching solutions for Go ([one](https://github.com/rs/dnscache), [two](https://github.com/mercari/go-dnscache)), however, I haven't seen any implementations that allow replacing the `net.DefaultResolver`?

Package [github.com/artyom/dot](https://github.com/artyom/dot) got me thinking if I could do the same for DNS caching, and also DNS over HTTPS.

So [github.com/ncruces/go-dns](https://godoc.org/github.com/ncruces/go-dns) is my attempt.
Replacing `net.DefaultResolver` with a caching DNS over HTTPS resolver using 1.1.1.1 as the name server should be this simple:

    net.DefaultResolver = dns.NewCachingResolver(dns.NewHTTPSResolver(
    	"1.1.1.1", "2606:4700:4700::1111",
    	"1.0.0.1", "2606:4700:4700::1001"))

What do you guys thing? Do you know of any other implementations that allow replacing `net.DefaultResolver`?
## [7][How to convert Map to Slice [3 gotchas]](https://www.reddit.com/r/golang/comments/gofaoa/how_to_convert_map_to_slice_3_gotchas/)
- url: https://web3.coach/golang-how-to-convert-map-to-slice-three-gotchas
---

## [8][Strange code behaviour](https://www.reddit.com/r/golang/comments/gof5w3/strange_code_behaviour/)
- url: https://www.reddit.com/r/golang/comments/gof5w3/strange_code_behaviour/
---
Hi. I've been golang enthusiast for a few months. Lately I wanted to consolidate my knowledge of golang concurrency. The following excerpt is from the book "Concurrency in Go" (by O’Reilly):

 [https://pastebin.com/L24Z8nd1](https://pastebin.com/L24Z8nd1) 

It works as it should (I suppose).   
However if I add a few fmt.Println or another function - it deadlocks. :

 [https://pastebin.com/cAPbMgGv](https://pastebin.com/cAPbMgGv) 

The use of the "useless" function causes deadlock 1/\~10k application runs. On the other hand, if I comment out the useless function and uncomment fmt.Println - deadlock occurs much more often. 

Can somebody explain why is it happening?
## [9][An open-source chatbot powered by an artificial neural network](https://www.reddit.com/r/golang/comments/gnxyk9/an_opensource_chatbot_powered_by_an_artificial/)
- url: https://github.com/olivia-ai/olivia
---

## [10][Three bugs in the Go MySQL Driver](https://www.reddit.com/r/golang/comments/gnm94y/three_bugs_in_the_go_mysql_driver/)
- url: https://github.blog/2020-05-20-three-bugs-in-the-go-mysql-driver/
---

