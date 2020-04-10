# golang
## [1][CLI for ordering Dominos pizza](https://www.reddit.com/r/golang/comments/fy2qya/cli_for_ordering_dominos_pizza/)
- url: https://www.reddit.com/r/golang/comments/fy2qya/cli_for_ordering_dominos_pizza/
---
[https://github.com/harrybrwn/apizza](https://github.com/harrybrwn/apizza)

My hobby project is finally at a place where I'm comfortable sharing it with the world. Tell me what you think.
## [2][GoLand 2020.1 is out now! Check out the Go Modules support and better overall improvements for your workflows!](https://www.reddit.com/r/golang/comments/fxxq1m/goland_20201_is_out_now_check_out_the_go_modules/)
- url: https://twitter.com/GoLandIDE/status/1248313039146844161
---

## [3][Protect CRUD API against third-party app creation](https://www.reddit.com/r/golang/comments/fyf4sr/protect_crud_api_against_thirdparty_app_creation/)
- url: https://www.reddit.com/r/golang/comments/fyf4sr/protect_crud_api_against_thirdparty_app_creation/
---
I'm building a GO CRUD API that I would like to protect as much as possible.  I'm already issuing JWTs and providing access to certain handlers based on that user's level of access.  The problem is that the API is exposed and anyone can sniff out the calls I'm making, reverse engineer how JWTs are created and create their own front end, or automated tool to perform unauthenticated/authenticated calls to my API.  I'd like to avoid third-party apps from using the API, even if they properly authenticate a user and gain a proper JWT.

My first thought is an API key for the web front ends and mobile apps (when that part comes). But even if we're sending the key via HTTPS with the other payload, that can still be sniffed with a MITM attack, so it's pointless. How can I appropriately verify that API calls are coming from a legitimate source before processing the request?
## [4][I made simple and easy to use custom oauth server with example included for mysql driver. Any suggestions?](https://www.reddit.com/r/golang/comments/fyficm/i_made_simple_and_easy_to_use_custom_oauth_server/)
- url: https://github.com/gobeam/golang-oauth
---

## [5][What are some favorite go packages for the following](https://www.reddit.com/r/golang/comments/fyfcxk/what_are_some_favorite_go_packages_for_the/)
- url: https://www.reddit.com/r/golang/comments/fyfcxk/what_are_some_favorite_go_packages_for_the/
---
CLI

Configuration

Serialization

Text processing

Regex

Api
## [6][ldhcpd: a light, minimal dhcpd with a GRPC control plane](https://www.reddit.com/r/golang/comments/fxsyuf/ldhcpd_a_light_minimal_dhcpd_with_a_grpc_control/)
- url: https://github.com/erikh/ldhcpd
---

## [7][Slice operation time complexity](https://www.reddit.com/r/golang/comments/fyckxf/slice_operation_time_complexity/)
- url: https://www.reddit.com/r/golang/comments/fyckxf/slice_operation_time_complexity/
---
I have a question about the time complexity of slice's window operation. Is it always O(1)?

&amp;#x200B;

For example,

If I had 2 lines of code that do this:

`s := make([]int, N)`

`s = s[N/2:]`

Is the 2nd line's time complexity O(1)? or O(N) because I moved up the start of the window by half of N?
## [8][Help me choose a framework for Golang web app](https://www.reddit.com/r/golang/comments/fybza8/help_me_choose_a_framework_for_golang_web_app/)
- url: https://www.reddit.com/r/golang/comments/fybza8/help_me_choose_a_framework_for_golang_web_app/
---
I'm a backend (DB) guy, I've last touched frontend (HTML) when 4.01 was a thing.  


I  love Go, what shall I try for a simple events (a fix has been deployed  happened) view and event inject (sy approves a fix), and a status page  (which fix is in what state) app?  


Macaron, beego, gobuffalo?  


I want sth less opinionated, more opt-in, and I'm lost with javascript, frontend frameworks. But want sth good-looking.  


Please help me!
## [9][Extend go library](https://www.reddit.com/r/golang/comments/fybove/extend_go_library/)
- url: https://github.com/go-board/x-go
---

## [10][GRPC Unidirectional streaming noob question](https://www.reddit.com/r/golang/comments/fyaczb/grpc_unidirectional_streaming_noob_question/)
- url: https://www.reddit.com/r/golang/comments/fyaczb/grpc_unidirectional_streaming_noob_question/
---
Hello all,

I have this piece of code written to receive messages from grpc unidirectional stream
```
func getmsg(cli terminalchat.Terminalchat_JoinClient, msgChan chan *terminalchat.Message) {
	for {
		msg := terminalchat.Message{}
		err := cli.RecvMsg(&amp;msg)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Got msg: ")
		msgChan &lt;- &amp;msg
	}
}
```
the description of `cli.RecvMsg` (generated with protoc btw) says it's supposed to block until it receives a message into m, returns io.EOF when stream completes successfully.

Even though nothing's been sent from the server, it keeps receiving EOF.

How can I stop receiving EOF from the stream?
