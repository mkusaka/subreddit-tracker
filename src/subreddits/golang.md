# golang
## [1][A Go Roguelike](https://www.reddit.com/r/golang/comments/f41b1v/a_go_roguelike/)
- url: https://github.com/hunterloftis/notes/blob/master/2020-02-14-golang-roguelike.md
---

## [2][What are some of the must-know topics for Backend Developer interviews (Go specific)?](https://www.reddit.com/r/golang/comments/f3trbd/what_are_some_of_the_mustknow_topics_for_backend/)
- url: https://www.reddit.com/r/golang/comments/f3trbd/what_are_some_of_the_mustknow_topics_for_backend/
---
I have 2 years of experience working in Go and 6 years in total of Backend development. I am based in London.  
I'm thinking of jumping back into the job hunting market after a long time.

What are some of the things I should brush up on?  
What would get me rejected?
## [3][Moving Towards DDD in Go](https://www.reddit.com/r/golang/comments/f480oe/moving_towards_ddd_in_go/)
- url: https://www.calhoun.io/moving-towards-domain-driven-design-in-go/
---

## [4][Best practices or tips writing a CLI tool with golang.](https://www.reddit.com/r/golang/comments/f454o4/best_practices_or_tips_writing_a_cli_tool_with/)
- url: https://www.reddit.com/r/golang/comments/f454o4/best_practices_or_tips_writing_a_cli_tool_with/
---
Hi, what patterns do you recommend to write a good CLI application in golang? 
Next week in our local community we are going to do the exercise, as we are a small group will be nice to have some external feedback. Thanks
## [5][Micro In Action, Part4：Pub/Sub](https://www.reddit.com/r/golang/comments/f3rqzv/micro_in_action_part4pubsub/)
- url: https://medium.com//micro-in-action-part4-pub-sub-564f3b054ecd?source=friends_link&amp;sk=45892c07676feae7086dae696fbf9a46
---

## [6][Structuring Web Applications in Go](https://www.reddit.com/r/golang/comments/f3rqbu/structuring_web_applications_in_go/)
- url: https://www.calhoun.io/structuring-web-applications-in-go
---

## [7][Problem when implementing reverse proxy](https://www.reddit.com/r/golang/comments/f43fp7/problem_when_implementing_reverse_proxy/)
- url: https://www.reddit.com/r/golang/comments/f43fp7/problem_when_implementing_reverse_proxy/
---
Hi all,

When I try to implement a reverse proxy in Go, using the bellow code:

    package main
    
    import (
    	"net/http"
    	"net/http/httputil"
    	url2 "net/url"
    )
    
    func redirect(target string, res http.ResponseWriter, req *http.Request) {
    	targetUrl,_ := url2.Parse(target)
    
    	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
    
    	req.URL.Host = targetUrl.Host
    	req.URL.Scheme = targetUrl.Scheme
    	req.Header.Set("X-Forwarded-Host",req.Header.Get("Host"))
    	req.Host = targetUrl.Host
    	req.URL.Path = targetUrl.Path
    	proxy.ServeHTTP(res,req)
    }
    
    func handleAndRedirect(res http.ResponseWriter, req *http.Request)  {
    	url := "http://localhost:5111/about"
    	redirect(url, res, req)
    }
    
    func main() {
    	http.HandleFunc("/", handleAndRedirect)
    	if err := http.ListenAndServe(":5106",nil); err != nil {
    		panic(err)
    	}
    }

Expected: When I enter localhost:5106 on browser, it returns the result from l**ocalhost:5111/about**

Actual: When I enter localhost:5106 on browser, it returns the result from **localhost:5111**

Is it possible to implement what I expect? If yes, what did I miss?Thanks.
## [8][What do you miss about Go when working in other languages?](https://www.reddit.com/r/golang/comments/f3u3cl/what_do_you_miss_about_go_when_working_in_other/)
- url: https://www.reddit.com/r/golang/comments/f3u3cl/what_do_you_miss_about_go_when_working_in_other/
---
And what do you not miss?
## [9][Any Golang Based E2E Frameworks for Testing ?](https://www.reddit.com/r/golang/comments/f3x2sm/any_golang_based_e2e_frameworks_for_testing/)
- url: https://www.reddit.com/r/golang/comments/f3x2sm/any_golang_based_e2e_frameworks_for_testing/
---
I am fairly new to Golang , Looking for options for an e2e testframework .

I am looking for something which has following features :

\+ Deploy Bunch of Servers and DUT's \[ it would be great if it can talk to git\]

\+ Run Tests on DUT's 

\+ Collect Logs from DUT and Test Reports .

\+ Pretty Display reports ,have a common place to collect them . 

If one framework cannot do this ,please suggest what's best open source software out there that can help me ?
## [10][What are good feature flag tools/libraries that support Go](https://www.reddit.com/r/golang/comments/f41vrl/what_are_good_feature_flag_toolslibraries_that/)
- url: https://www.reddit.com/r/golang/comments/f41vrl/what_are_good_feature_flag_toolslibraries_that/
---
&amp;#x200B;

   Hello, I'm currently working as a senior developer/architect.  I plan on doing a pretty massive code migration in the next few months, and I decided to center my strategy around feature flags.  I've been doing some research for a few months, and I've come across list

&amp;#x200B;

[https://featureflags.io/go-feature-flags/](https://featureflags.io/go-feature-flags/)

&amp;#x200B;

   For the record, I work in an event driven architectures with pollers and streams.  I have found many feature flag libraries to be very oriented towards web based solutions, but I don't interact with the client side at all.  This a purely backend system.   I do plan on trying a few on the list, but I wanted to reach out to the community and get some ideas about your personal accounts and experiences with feature flag tools?
