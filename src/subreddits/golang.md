# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][What’s it like (Career-wise) as a Go programmer?](https://www.reddit.com/r/golang/comments/jqrks8/whats_it_like_careerwise_as_a_go_programmer/)
- url: https://www.reddit.com/r/golang/comments/jqrks8/whats_it_like_careerwise_as_a_go_programmer/
---
I’d love to know what you day to say is. What is your career, how much time do you actually spending coding? Do you recommend Go? What made you interested in it?

And also, have you done any cool (personal) projects?

Don’t need to answer above, I’d just like to read your experience!
## [3][If google is supporting Go, is this search result just a joke ? :D](https://www.reddit.com/r/golang/comments/jqbjn2/if_google_is_supporting_go_is_this_search_result/)
- url: https://i.redd.it/iuo2ju8nn0y51.png
---

## [4][Upcoming Go 1.16 reduces significantly reduces sizes for built executables](https://www.reddit.com/r/golang/comments/jqdkc6/upcoming_go_116_reduces_significantly_reduces/)
- url: https://www.reddit.com/r/golang/comments/jqdkc6/upcoming_go_116_reduces_significantly_reduces/
---
Just tried building [VictoriaMetrics project](https://github.com/VictoriaMetrics/VictoriaMetrics/) with Go tip (aka upcoming Go 1.16) and noticed significant reduction in produced binary size comparing to Go 1.15.4:

```
$ ls -lGg --si bin/
total 34M
-rwxr-xr-x 1 19M Nov  7 01:19 victoria-metrics-go1.15.4
-rwxr-xr-x 1 16M Nov  7 01:21 victoria-metrics-go-tip
```

VictoriaMetrics executable size has been reduced from 19MB to 16MB - almost 16% reduction. This looks awesome!

It would be great to know which particular changes in Go 1.16 are responsible for such great binary size reduction.
## [5][Why does the log package always output to stderr?](https://www.reddit.com/r/golang/comments/jqvaft/why_does_the_log_package_always_output_to_stderr/)
- url: https://www.reddit.com/r/golang/comments/jqvaft/why_does_the_log_package_always_output_to_stderr/
---
I'm just starting out with Go and wondered why the log package always outputs to stderr? 

I would expect some functions to output to stdout and others to stderr. For example calls to log.Fatal would go to stderr, calls to log.Print to stdout. 

Is there some idiomatic way to handle logging in Go?
## [6][gRPC](https://www.reddit.com/r/golang/comments/jqsf0n/grpc/)
- url: https://www.reddit.com/r/golang/comments/jqsf0n/grpc/
---
Is anyone using gRPC or still using Restful APIs? I tried using it a while back (\~ 2 years) but had to use [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to be able to invoke from react (web)/ mobile etc but we quickly shifted back to Restful APIs.
## [7][I have rewritten "Spinning Donut" to Golang](https://www.reddit.com/r/golang/comments/jqxm6d/i_have_rewritten_spinning_donut_to_golang/)
- url: https://www.reddit.com/r/golang/comments/jqxm6d/i_have_rewritten_spinning_donut_to_golang/
---
Implementation of "Spinning Donut" on Golang, not much different but cool

Left: Go  
Right: C  
Source: [https://github.com/Vlad-Shevliakov/ASCII-render](https://github.com/Vlad-Shevliakov/ASCII-render)

https://reddit.com/link/jqxm6d/video/9m2dr0o5v7y51/player
## [8][I added syntax highlighting to Go's present tool. Inspired by flippeeer's presentation.](https://www.reddit.com/r/golang/comments/jqkl5m/i_added_syntax_highlighting_to_gos_present_tool/)
- url: https://github.com/soypat/go-presentx
---

## [9][Simple Modern OpenGL Renderer Written in Go](https://www.reddit.com/r/golang/comments/jqfswr/simple_modern_opengl_renderer_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/jqfswr/simple_modern_opengl_renderer_written_in_go/
---
Hi folks,

I want to share my recent side project, nothing fancy but I'm looking forward to your feedback also maybe some of you will find it interesting to play around with the code.

hope you like it, Thanks.

[https://github.com/dkvilo/andromeda](https://github.com/dkvilo/andromeda)
## [10][Go Offline Development. Best Practices/Tips](https://www.reddit.com/r/golang/comments/jquf12/go_offline_development_best_practicestips/)
- url: https://www.reddit.com/r/golang/comments/jquf12/go_offline_development_best_practicestips/
---
\[Solved\]

Seems like this tutorial will do the trick [https://medium.com/@bnprashanth256/dependency-management-and-vendoring-in-go-golang-5ec6d6b7590e](https://medium.com/@bnprashanth256/dependency-management-and-vendoring-in-go-golang-5ec6d6b7590e)

&amp;#x200B;

Hey guys,

I've got an interesting problem. I need to write code in Go in an airgapped environment. That means no internet connection and only file transfer via secure USB-Sticks.

**How can I download and resolve my dependencies for 3rd-party librarys without an internet connection?**

I can't seem to find anything online about this topic. I need to be able to use the `gopkg.in/yaml.v2` Libary. Simply downloading the .zip from GitHub and copying it manually to the GOPATH is not feasible. Most library use a decent amount of other library so i need a way (that can be automated) to download all sub-dependencies to.

Thanks in advance for all your input.
## [11][Machine is a zero dependency library for highly concurrent Go applications](https://www.reddit.com/r/golang/comments/jql79e/machine_is_a_zero_dependency_library_for_highly/)
- url: https://www.reddit.com/r/golang/comments/jql79e/machine_is_a_zero_dependency_library_for_highly/
---
Update: v1.1.0 

&amp;#x200B;

https://preview.redd.it/ybk93cwmg3y51.jpg?width=800&amp;format=pjpg&amp;auto=webp&amp;s=61deeda7df783a2d0904790f4d066a27787b9682

* GitHub: [https://github.com/autom8ter/machine](https://github.com/autom8ter/machine)
* GoDoc: [https://pkg.go.dev/github.com/autom8ter/machine](https://pkg.go.dev/github.com/autom8ter/machine)
* import "[github.com/autom8ter/machine](https://github.com/autom8ter/machine)"

Machine is a zero dependency runtime for managed goroutines. It is inspired by [errgroup.Group](https://errgroup.Group) with extra bells &amp; whistles:

* throttled goroutines
* self-cancellable goroutines with \[*\`*Context\*\`\*\]([*https://golang.org/pkg/context#Context*](https://golang.org/pkg/context#Context))
* global-cancellable goroutines with context (see \[Cancel\]([*https://pkg.go.dev/github.com/autom8ter/machine#Machine.Cancel*](https://pkg.go.dev/github.com/autom8ter/machine#Machine.Cancel)))
* goroutines have IDs and optional tags for easy debugging (see \[Stats\]([*https://pkg.go.dev/github.com/autom8ter/machine#Machine.Stats*](https://pkg.go.dev/github.com/autom8ter/machine#Machine.Stats)))
* native \[publish/subscribe\]([*https://pkg.go.dev/github.com/autom8ter/machine/pubsub#PubSub*](https://pkg.go.dev/github.com/autom8ter/machine/pubsub#PubSub)) implementation for broadcasting messages to active goroutines
* \[middlewares\]([*https://pkg.go.dev/github.com/autom8ter/machine#Middleware*](https://pkg.go.dev/github.com/autom8ter/machine#Middleware)) for wrapping/decorating functions
* "sub" machines for creating a dependency tree between groups of goroutines
* goroutine leak prevention
* native pprof &amp; golang execution tracer integration
* greater than 80% test coverage

## Use Cases

[Machine](https://pkg.go.dev/github.com/autom8ter/machine#Machine) is meant to be completely agnostic and dependency free- its use cases are expected to be emergent. Really, it can be used **anywhere** goroutines are used.

Highly concurrent and/or asynchronous applications include:

* gRPC streaming servers
* websocket servers
* pubsub servers
* reverse proxies
* cron jobs
* custom database/cache
* ETL pipelines
* log sink
* filesystem walker
* code generation

Visit [https://github.com/autom8ter/machine/blob/master/examples/README.md](https://github.com/autom8ter/machine/blob/master/examples/README.md) for examples that are easy to grok(under &lt; 500 lines of code)

* gRPC Bidirectional Chat Stream Server
* TCP Reverse Proxy
* Concurrent Cron Job Server
