# golang
## [1][We are the Go Time podcast. Ask us anything. (AMA)](https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/)
- url: https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/
---
Hi everyone! I'm Jon Calhoun, one of the panelists on the Go Time podcast. For those of you unfamiliar - it is a Go podcast that we record live every Tuesday at 3pm ET. We usually have a guest or two on each episode, and discuss a specific topic (defer, testing, databases, infra, etc). You can check it out here: &lt;https://changelog.com/gotime&gt;

This coming episode we want to try something a little different - we want to make a Q&amp;A episode. There are two reasons for doing this:

1. We are hoping this inspires topics for future episodes.
2. We want a venue to discuss questions that don't really fit into an entire episode on their own.

To make this happen we would like everyone here to post any Go-related questions that you would like us to discuss on air, and we will try to get to as many as possible. I'll also try to type up answers here while we discuss them on the episode.

We will be answering questions live tomorrow, Tuesday, Sep 8. We will repeat questions on air, and since we record live you can join in on the Gophers Slack to ask follow-up questions or to elaborate on questions.
## [2][What is ':' and '...'?](https://www.reddit.com/r/golang/comments/iop8rg/what_is_and/)
- url: https://www.reddit.com/r/golang/comments/iop8rg/what_is_and/
---
I was recently working on a project and I came across this:

    package main
    
    import "fmt"
    
    func min(nums ...float64) (min float64) {
    	for _, n := range nums {
    		if min&gt;n {
    			min = n
    		}
    	}
    	return
    }
    
    func main() {
    	s:=[]float64{43.009, 33.112, -3.2, 0.0}
    	fmt.Println(min(s[:]...))
    }

I am **unable to understand the meaning of** `s[:]...` what does that even mean?
## [3][OOP objects v Go Structs](https://www.reddit.com/r/golang/comments/ioqpgk/oop_objects_v_go_structs/)
- url: https://www.reddit.com/r/golang/comments/ioqpgk/oop_objects_v_go_structs/
---
I’m a Go noob but an experienced developer.

In an OOP language I can create an object by passing arguments to its constructor. I can validate these arguments and reason that if my object “Car” exists, it’s make property will always be valid (eg “Ford” or “Ferrari”).

Or, I can create a DB object and inject it into my Repository, and know that when I call repo.db.select(...) the select method will execute against a db connection.

How do you approach this sort of thing idiomatically in Go? If I have a Car struct anyone can create one with arbitrary properties.

Is it simply that I have to get my head around living with structs that could always have invalid values? Do you end up doing nil checks because you can’t guarantee your sub-structs exists/are valid?

Any recommendations for articles/resources targeted at getting out of OOP mindset and into idiomatic Go?

Thanks.
## [4][Beware of upgrading gorilla/handlers to 1.5.0](https://www.reddit.com/r/golang/comments/ios44w/beware_of_upgrading_gorillahandlers_to_150/)
- url: https://www.reddit.com/r/golang/comments/ios44w/beware_of_upgrading_gorillahandlers_to_150/
---
Currently gzip compression is broken, see [https://github.com/gorilla/handlers/issues/194](https://github.com/gorilla/handlers/issues/194). Keep 1.4.2 or use commit hash [d453eff](https://github.com/gorilla/handlers/commit/d453effd20e6817a8acfd9d278dadd21371da0ba).
## [5][What other languages have something akin to Go's interfaces?](https://www.reddit.com/r/golang/comments/iorveh/what_other_languages_have_something_akin_to_gos/)
- url: https://www.reddit.com/r/golang/comments/iorveh/what_other_languages_have_something_akin_to_gos/
---
Well except duck-typed OOP language which do have somewhat similar possibilities but only incidentally.

I've looked at this multiple times in the past, but I can't find another language that has the feature of implicitly satisfying interfaces. Rusts `impl &lt;trait&gt; for &lt;type&gt;` seems to allow similar flexibility but with explicit notation.
## [6][What is New in the go 1.15](https://www.reddit.com/r/golang/comments/iomyzw/what_is_new_in_the_go_115/)
- url: https://www.loginradius.com/engineering/blog/what-is-the-new-go-1.15/
---

## [7][When to use buffered channels in go](https://www.reddit.com/r/golang/comments/iorcmp/when_to_use_buffered_channels_in_go/)
- url: https://medium.com/@ss7loginov/when-to-use-buffered-channels-in-go-317d3593d28f
---

## [8][http.Get returns empty response but curl (on same machine) gets data](https://www.reddit.com/r/golang/comments/iorwxj/httpget_returns_empty_response_but_curl_on_same/)
- url: https://www.reddit.com/r/golang/comments/iorwxj/httpget_returns_empty_response_but_curl_on_same/
---
This issue is in two parts. I'm trying to simply request data from an endpoint (it's a CDN). Originally I was doing my request against `https://endpoint`. It was working with both my Go application as well as curl.

Then that stopped working, for *both* my application *and* curl. Because it is a CDN I figured it was some caching issue. So then I generated a random number and appended it as a query param. The new url would be similar to `https://endpoint/?nocache=1234`.

This worked for both my application as well as curl!!

... and then it stopped working for my application. But it continued to work for curl. This is all on a single host.

I would have my Go application print out the used url (e.g. `https://endpoint/?nocache=9876`) as well as the response.ContentLength (now it is always 0, but used to be -1 when it was actually getting data), and the status (*always* `200 OK`).

Zero data from my Go application. But then I would take the same exact url and curl it: `curl "https://endpoint/9876`. And I receive data.

Has anybody else seen this before? I'm wondering if the edge servers for this CDN have somehow recognized a header or something from my Go application that it is "filtering" of sorts.

Outside of getting out a network tracing tool to see the packets, I'm not sure how to proceed from here. I'd hate to shell out to curl from my application, but this is a pretty hard stop here.

Hoping that maybe I'm not the first to deal with this problem. Thanks in advance!
## [9][Generate RESTful service and swagger documentation with gRPC gateway. This enables us to write code just once but serving both gRPC and HTTP JSON requests at the same time.](https://www.reddit.com/r/golang/comments/ioswkv/generate_restful_service_and_swagger/)
- url: https://youtu.be/Zf9G2KzYs7w
---

## [10][go-pg Postgres ORM v10 with support for map[string]interface{}](https://www.reddit.com/r/golang/comments/ioopad/gopg_postgres_orm_v10_with_support_for/)
- url: https://github.com/go-pg/pg/blob/v10/CHANGELOG.md#changelog
---

## [11][Introducing TA, a technical analysis Go library](https://www.reddit.com/r/golang/comments/iohi6o/introducing_ta_a_technical_analysis_go_library/)
- url: https://www.reddit.com/r/golang/comments/iohi6o/introducing_ta_a_technical_analysis_go_library/
---
I'd like to introduce my newest toy project, [ta](https://github.com/OneOfOne/ta), a technical analysis library with strategy / backtest helpers and csv loader.

Right now it's pre-alpha, the API is unstable, however the functions were tested for accuracy against the python version of [talib](https://github.com/mrjbq7/ta-lib).

All feedback and pull requests are welcome.
