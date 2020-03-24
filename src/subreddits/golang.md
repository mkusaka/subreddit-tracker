# golang
## [1][How do you use context.Context in your applications?](https://www.reddit.com/r/golang/comments/fo395d/how_do_you_use_contextcontext_in_your_applications/)
- url: https://www.reddit.com/r/golang/comments/fo395d/how_do_you_use_contextcontext_in_your_applications/
---
The context in Go is very popular. The most popular use case is the cancellation and setting timeouts. But, we can pass some scope-related data in it. And here's my question: do you do it?

One use case I can imagine is putting the authenticated user's details in the middleware and retrieving them in the final HTTP handler. What's your practice?

I'm asking because I noticed that every time I try to put ANYTHING into the context, I have a very long discussion about it in PRs. Do you have any (in your opinion) good examples of using \`context.WithValue\` and similar? Should we use it this way or not?
## [2][OpenBiohazard2: Open source reimplementation of the original Resident Evil 2 / Biohazard 2](https://www.reddit.com/r/golang/comments/fnxlk6/openbiohazard2_open_source_reimplementation_of/)
- url: https://github.com/samuelyuan/OpenBiohazard2
---

## [3][Event sourcing in Go.](https://www.reddit.com/r/golang/comments/fo2pn6/event_sourcing_in_go/)
- url: https://www.reddit.com/r/golang/comments/fo2pn6/event_sourcing_in_go/
---
I wrote a blog post about how one might go about implementing event sourcing in Go and not use any frameworks.

[https://victoramartinez.com/posts/event-sourcing-in-go/](https://victoramartinez.com/posts/event-sourcing-in-go/)
## [4][Golang and SQLite, when to open connection ?](https://www.reddit.com/r/golang/comments/fnygei/golang_and_sqlite_when_to_open_connection/)
- url: https://www.reddit.com/r/golang/comments/fnygei/golang_and_sqlite_when_to_open_connection/
---
Hi, I'm building a web app in Go in which I have to use a SQLite database.

Should I open a connection every time I have to access the database (in the http Handlers) or should I open a connection at startup and make it a global var ?

Thank you !
## [5][Is Go low level enough that one could write an OS in it without leveraging other languages like C?](https://www.reddit.com/r/golang/comments/fnuiyc/is_go_low_level_enough_that_one_could_write_an_os/)
- url: https://www.reddit.com/r/golang/comments/fnuiyc/is_go_low_level_enough_that_one_could_write_an_os/
---
I’m trying to understand the overhead associate with Golang.

Is Go compiled in a way where it could be use to make an OS without having to lean on C or Assembly to get bootstrapped?

If not, what is the limiting factor? Is there a runtime? Garbage collector missing? What makes it such that you need something else?
## [6][Graphing COVID-19 with Go and chart.js](https://www.reddit.com/r/golang/comments/fnj3xa/graphing_covid19_with_go_and_chartjs/)
- url: https://github.com/kennygrant/coronavirus
---

## [7][Issue with coverage testing integration tests](https://www.reddit.com/r/golang/comments/fo2o55/issue_with_coverage_testing_integration_tests/)
- url: https://www.reddit.com/r/golang/comments/fo2o55/issue_with_coverage_testing_integration_tests/
---
Hi :)

We are about to start collecting coverage analytics from our integration tests. For this I have instrumented the binary with coverage analysis and compiled it with go test. This all works like a charm.

The issue however, and I am not sure of this, I’m not that deep in go, but I believe that the test framework does some buffering on STDIN and STDOUT. So the problem is, in our tests, the binary streams a whole disc partition from one device to another. This works fine for the regular binary, but with the go test compiled binary, I always get an: error: no space left on device stdout :(

I have tried re-opening STDIN and STDOUT, in the test, to no avail :/

How should I approach this problem?

Any suggestions welcome!
## [8][Don't Use Frameworks](https://www.reddit.com/r/golang/comments/fo50l5/dont_use_frameworks/)
- url: https://victoramartinez.com/posts/dont-use-frameworks/
---

## [9][Cannot install srtgo package](https://www.reddit.com/r/golang/comments/fo44ll/cannot_install_srtgo_package/)
- url: https://www.reddit.com/r/golang/comments/fo44ll/cannot_install_srtgo_package/
---
Has anyone tried to use srt package ? I've tried to install it, but getting error:

    # github.com/haivision/srtgo 
    /tmp/go-build815457234/b001/\_cgo\_main.o:/tmp/go-build/cgo-generated-wrappers:5: undefined reference to \`SRT\_INVALID\_SOCK' 
    /tmp/go-build815457234/b001/\_cgo\_main.o:/tmp/go-build/cgo-generated-wrappers:2: undefined reference to \`SRT\_ERROR' 
    collect2: error: ld returned 1 exit status
## [10][[Question] Manipulate binary files ?](https://www.reddit.com/r/golang/comments/fo41sx/question_manipulate_binary_files/)
- url: https://www.reddit.com/r/golang/comments/fo41sx/question_manipulate_binary_files/
---
Hi, I’d like to know if the std library or other library is able to manipulate binary file, for exemple : 
- Open file.bin file
- Read bytes from a certain offset to another one.
- Replace bytes from an offset to another one
- Then save or rewrite the file.

In fact, the thing I’m trying to do can be done by an hex editor, but I want to make it auto :)

Thanks in advance
