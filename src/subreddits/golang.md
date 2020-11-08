# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Better error handling with DDD Error](https://www.reddit.com/r/golang/comments/jq9f91/better_error_handling_with_ddd_error/)
- url: https://www.reddit.com/r/golang/comments/jq9f91/better_error_handling_with_ddd_error/
---
Hey, everyone!

Today I just finished my first Go package after more than a year using this amazing lang. I wanted to shared it with you all.

I specifically focused on following the Go’s simplicity principles and keep the library idiomatic. Don’t worry about any good ol’ Java/PHP/C#/Typescript complexities.

`DDD Error` is a _generic domain-driven exception wrapper_ made for Go.

Using existing _validators_ such as playground's implementation is _overwhelming because tag validation_ and the need to rewrite descriptions. 
With `DDD Error`, _you may still use 3rd-party validators_ or _make your own validations_ in your value objects, entities or aggregates.

In addition, infrastructure exceptions were added so you may be able to _catch specific kind of infrastructure errors._

Exceptions _descriptions are based on the [Google Cloud API Design Guidelines](https://cloud.google.com/apis/design/errors)_.

`DDD Error` is compatible with popular error-handling packages such as [Hashicorp's go-multierror](https://github.com/hashicorp/go-multierror)

In conclusion, `DDD Error` aims to _ease the lack of exception handling_ in The Go Programming Language by defining a _wide selection of common exceptions_ 
which happen inside the _domain and/or infrastructure_ layer(s).

_Note: `DDD Error` is dependency-free, it uses built-in packages such as errors package._

**Which use cases you'll find DDD Error useful?**
- Implement retry strategy and circuit breaker resiliency patterns by adding Network exception to the whitelist.
- Not Acknowledging messages from an event bus if got a Network or Infrastructure generic exception.
- Get an HTTP/gRPC/OpenCensus status code from an error.
- Implement multiple strategies when an specific (or generic) type of error was thrown in.
- Fine-grained exception logging on infrastructure layer by using GetParentDescription() function.
- More.

_The only requirement is Go version 1.13 or above._

[Take me to the GitHub repository](https://github.com/NeutrinoCorp/ddderr)

_Note: I've been using similar strategies for quite a time, but have in mind I made this package in just a few hours. Please, feel free to report any bug._
## [3][Today, I presented Go to my team... (the screen recoding)](https://www.reddit.com/r/golang/comments/jpugtg/today_i_presented_go_to_my_team_the_screen/)
- url: https://www.reddit.com/r/golang/comments/jpugtg/today_i_presented_go_to_my_team_the_screen/
---
A few days ago, I posted my experience of [introducing Go to my team](https://www.reddit.com/r/golang/comments/jnhzm5/today_i_presented_go_to_my_team/) and the feedback was overwhelming. I would like to thank you all again for your support!  &lt;3

As some have requested, here is the screen recoding of the presenation for you to watch:

[https://www.youtube.com/watch?v=Mjcw8fHdx8Q](https://www.youtube.com/watch?v=Mjcw8fHdx8Q)

*Make sure to speed it up, if I'm talking too slow :)*

EDIT: New YouTube link (old one had broken audio)

EDIT2: As many have been asking, I've used a personally modified version of [present](https://pkg.go.dev/golang.org/x/tools/present) for making the slides.
You can install it with: `go get golang.org/x/tools/cmd/present`
## [4][If google is supporting Go, is this search result just a joke ? :D](https://www.reddit.com/r/golang/comments/jqbjn2/if_google_is_supporting_go_is_this_search_result/)
- url: https://i.redd.it/iuo2ju8nn0y51.png
---

## [5][Clean Boilerplate of Go, Domain-Driven Design, Clean Architecture, Gin and GORM.](https://www.reddit.com/r/golang/comments/jq1id7/clean_boilerplate_of_go_domaindriven_design_clean/)
- url: https://i.redd.it/mdrc2aexqwx51.png
---

## [6][Tutorial for beginners](https://www.reddit.com/r/golang/comments/jqa93y/tutorial_for_beginners/)
- url: https://www.reddit.com/r/golang/comments/jqa93y/tutorial_for_beginners/
---
Hi everyone! I'm looking for some interesting tutorial for beginners. Suggestions? 🤓
## [7][Go http client closes the underlying tcp connection when it hits a timeout with a request, ignoring the connection pooling settings. Anyone knows the reason why?](https://www.reddit.com/r/golang/comments/jq8j7i/go_http_client_closes_the_underlying_tcp/)
- url: https://v.redd.it/s3xc7z99bzx51
---

## [8][How suitable is Go for handling thousands of busy websocket connections with decent latency?](https://www.reddit.com/r/golang/comments/jq8b7s/how_suitable_is_go_for_handling_thousands_of_busy/)
- url: https://www.reddit.com/r/golang/comments/jq8b7s/how_suitable_is_go_for_handling_thousands_of_busy/
---
Hi there,

I have had a project in mind for a while that I initially wanted to implement with Python. I have a working prototype that I have tested with low player counts.

However you don't have to be a genius to see how hard it will be to scale it. Python's Async / concurrency APIs are not especially mature (and make your codebase quite ugly to be honest).

The requirements for this application server are that it has to be able to handle up to 5k websocket connections, do some simple business logic, and broadcast data to all clients without significant latency.  Has anyone here done something similar with Go?
## [9][Static Dependency Analysis Tool for Go Files](https://www.reddit.com/r/golang/comments/jpm0aj/static_dependency_analysis_tool_for_go_files/)
- url: https://i.redd.it/h71qm0zrhrx51.gif
---

## [10][I started learning Go today... and I really like it!](https://www.reddit.com/r/golang/comments/jq4534/i_started_learning_go_today_and_i_really_like_it/)
- url: https://www.reddit.com/r/golang/comments/jq4534/i_started_learning_go_today_and_i_really_like_it/
---
![](https://pbs.twimg.com/media/Ef-GZDIXoAEnbhE?format=jpg&amp;name=small)

I spent the whole day reading Introducing Go by Caleb Doxsey and I already feel pretty confident writing Go code.

After over a year of working with Rust I wanted to try Go out and I was surprised by how simple it is and also kinda bugged about how verbose it can be sometimes. I really missed Rust's compile time checks and the ownership model, but I also felt that I had more control over what I wanted my code to do, even if it could lead to runtime errors. 

Overall my first experience with the language was very pleasant and I plan to keep on learning, perhaps get a job someday, who knows?
## [11][atomic.Value docs - explain "must not be copied" please](https://www.reddit.com/r/golang/comments/jq1jgd/atomicvalue_docs_explain_must_not_be_copied_please/)
- url: https://www.reddit.com/r/golang/comments/jq1jgd/atomicvalue_docs_explain_must_not_be_copied_please/
---
Hi, the atomic.Value doc (https://golang.org/pkg/sync/atomic/) states

&gt; Once Store has been called, a Value must not be copied.
&gt;
&gt; A Value must not be copied after first use.

What does that even mean? 

Does it just mean that under
`x := y.Load()`
the `x` will not show the atomic behaviour that `y` enjoys? I'd think this is so obvious that I suspect that statement in the docs mean something else.

Considering the following most basic snippet with an atomic.Value `a`, strictly without anything "copied":
```
if len(a.Load().(string)) == 0 {
  panic("empty string!")
}

println(a.Load().(string))
```

In the above snippet, isn't the use of atomic.Value completely useless because between the two Load()'s a Store() might occur from somewhere else, setting `a` to an empty string that would then have bypassed the panic-check?

So what one obviously needs to do is:
```
x := a.Load().(string)

if len(x) == 0 {
  panic("empty string!")
}

println(x)
```

But beware: `x := a.Load().(string)` is copying which *must not be done*... no?  

To me it seems some copying *must totally* be done for atomic.Value to have any use at all.

Can someone please unconfuse me? Thanks.
