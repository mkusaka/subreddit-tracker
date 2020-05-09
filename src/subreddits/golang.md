# golang
## [1][immudb: A lightweight, high-speed immutable database for systems and applications in Go.](https://www.reddit.com/r/golang/comments/ggds5c/immudb_a_lightweight_highspeed_immutable_database/)
- url: https://github.com/codenotary/immudb
---

## [2][I made a CSVtoJSON CLI tool with Golang](https://www.reddit.com/r/golang/comments/gg5m7f/i_made_a_csvtojson_cli_tool_with_golang/)
- url: https://www.reddit.com/r/golang/comments/gg5m7f/i_made_a_csvtojson_cli_tool_with_golang/
---
Hi!

After 7 years of working with Javascript, I decided that it was time to try a new programming language. During the last weeks, I've been learning Golang using this website as my main learning resource: [https://gobyexample.com/](https://gobyexample.com/) .  I thought that the best way to challenge my recently acquired skills was to create a simple program without following any tutorial or course. That's the reason I made this CLI tool that you can check here:

 [https://github.com/Andrew4d3/go-csv2json](https://github.com/Andrew4d3/go-csv2json)

Bear in mind that I'm a complete beginner using Go. So probably the code you're gonna see is not the best out there. *But hey! At least I tried*... I'll be writing a Medium article in the future explaining what I did, but first I want to read your opinions. To see if there's something where I can improve myself.

Any suggestion, advice, or feedback you want to give me is completely welcome. :)
## [3][godoc-static - Generate static Go documentation](https://www.reddit.com/r/golang/comments/gg91mm/godocstatic_generate_static_go_documentation/)
- url: https://gitlab.com/tslocum/godoc-static
---

## [4][Communicate between grpc microservices with jwt](https://www.reddit.com/r/golang/comments/gg99xe/communicate_between_grpc_microservices_with_jwt/)
- url: https://www.reddit.com/r/golang/comments/gg99xe/communicate_between_grpc_microservices_with_jwt/
---
How can we communicate between microservices with jwt? I want to communicate between two microservices first takes care of the username login and password and second microservices should get username and other information of the user.
## [5][[Question] How does `cancel()` result in`&lt;-ctx.Done()` to be the case chosen in a "select" statement?](https://www.reddit.com/r/golang/comments/ggdueg/question_how_does_cancel_result_inctxdone_to_be/)
- url: https://www.reddit.com/r/golang/comments/ggdueg/question_how_does_cancel_result_inctxdone_to_be/
---
I'm trying to understand "select" statements of the following kind:
```go
select {
	case &lt;-ctx.Done():
	...
	// Timeout
	case r:= &lt;-time.After(...):
	...
}
```

Because of how "select" statements work, the channel operands of these two receive operations are evaluated exactly once. i.e., ctx.Done() and time.After(...) are evaluated exactly once. The result is a set of channels to receive from and the corresponding values to send. Since there is no default "case", the "select" statement blocks until at least one of the communications can proceed.

From [here](https://golang.org/pkg/time/#After), I understand that `time.After(...)` will be the case chosen to proceed, if the specified duration has elapsed.

What I've already read but do not completely understand:

From [here](https://golang.org/src/context/context.go?s=2460:6019#L357), I see that `Done()` returns a channel with elements of type `struct{}`.

From [here](https://golang.org/src/context/context.go?s=2460:6019#L389) I see that `cancel()` closes `c.done`.

I know that closing a channel causes receive operations to return the zero value for the channel's element type without blocking. I also know that receiving from a "nil" channel blocks forever; I have a feeling that this fact fits into this picture somehow because I see `if c.done == nil` in the code. 

I can't see the whole 'chain' of how a `cancel()` results in causing `ctx.Done()` to proceed first in the "select" statement.

Please do correct me if I have any misunderstandings.
## [6][HTTP client and REST API](https://www.reddit.com/r/golang/comments/ggf2h6/http_client_and_rest_api/)
- url: https://www.reddit.com/r/golang/comments/ggf2h6/http_client_and_rest_api/
---
There was a question about go vs python for REST services. That was geared toward servers. My question is about client.

Our organization introduced a tool which supports REST. There is language support of Python/NodeJs/Java/C#. But not Go. I would like to write the Go implementation for internal use.

I have it working with Go's standard library net/http client. It works but I would like to make it easier on the user.  I would like to abstract net/http client from the user. Is that a good idea?

Are there any good examples I can refer to?

I looked at this, [https://www.scaledrone.com/blog/creating-an-api-client-in-go/](https://www.scaledrone.com/blog/creating-an-api-client-in-go/). Is this a good idea?
## [7][Web Application Firewall written in Go](https://www.reddit.com/r/golang/comments/gfsns4/web_application_firewall_written_in_go/)
- url: https://github.com/asalih/guardian
---

## [8][Open sourced my Github App](https://www.reddit.com/r/golang/comments/ggdfr8/open_sourced_my_github_app/)
- url: https://www.reddit.com/r/golang/comments/ggdfr8/open_sourced_my_github_app/
---
Hi All, I am still learning Golang, but I have just released the source code of my Github Application ([https://pullassistant.com/](https://pullassistant.com/)) 

Repository: [https://github.com/pullassistant/PA-Monorepo](https://github.com/pullassistant/PA-Monorepo)

If someone wants to build a similar stuff, it might be useful! Thanks!
## [9][Checksum validation in goose](https://www.reddit.com/r/golang/comments/ggadj4/checksum_validation_in_goose/)
- url: https://www.reddit.com/r/golang/comments/ggadj4/checksum_validation_in_goose/
---
 Hi , I just started with golang , have couple of doubts regarding DB  migration , we are using goose for db migration , What I noticed is that  it does not do the checksum validation while running the migration like  done in flyway(java) , suppose for example I have a migration file  which is ran before now if I make any change or update to it , it does  not run the migration or throw any error, for every migration I have to  modify the name of the file or create a new file , is there any thing goose provide out of the box to do this ?
## [10][Developers say Google's Go is 'most sought after' programming language of 2020 | ZDNet](https://www.reddit.com/r/golang/comments/gflyyn/developers_say_googles_go_is_most_sought_after/)
- url: https://www.zdnet.com/article/developers-say-googles-go-is-most-sought-after-programming-language-of-2020/
---

