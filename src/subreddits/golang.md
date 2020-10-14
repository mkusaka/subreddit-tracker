# golang
## [1][Teles is a small logger for Go. Logging in telegram bot, log file, terminal.](https://www.reddit.com/r/golang/comments/jays0n/teles_is_a_small_logger_for_go_logging_in/)
- url: https://github.com/4FR4KO-POVELECKO/teles
---

## [2][Machine: A zero dependency library for managed goroutines, inspired by errgroup.Group.](https://www.reddit.com/r/golang/comments/jaqypa/machine_a_zero_dependency_library_for_managed/)
- url: https://www.reddit.com/r/golang/comments/jaqypa/machine_a_zero_dependency_library_for_managed/
---
- GitHub: https://github.com/autom8ter/machine
- GoDoc: https://pkg.go.dev/github.com/autom8ter/machine
- import "github.com/autom8ter/machine"


Machine is a zero dependency runtime for managed goroutines. It is inspired by errgroup.Group with extra bells &amp; whistles:

- throttled/max goroutines

- global context cancellation propagation across all goroutines or groups of goroutines(sub machines)

- singe goroutine context cancellation propagation

- goroutines are each given an ID, tags, and record metrics for monitoring/debugging(see machine.Stats)

- pluggable publish/subscribe interface for passing messages between goroutines. An in memory implementation is provided by default, but may be swapped out with redis/nats when scaling horizontally

- middlewares for wrapping/decorating functions(cron, panic-recovery, logging, etc)

- "sub" machines for creating a dependency tree between groups of goroutines

- goroutine leak prevention- machine.Wait() will not unblock until all goroutines exit

- greater than 80% test coverage

- familiar interface for executing goroutines: machine.Go(func(r Routine))

Visit https://github.com/autom8ter/machine/blob/master/examples/README.md for examples that are easy to grok(under &lt; 500 lines of code)
- gRPC Bidirectional Chat Stream Server
- TCP Reverse Proxy
- Concurrent Cron Job Server


Roadmap:
- Prometheus middleware for collecting fine-grained metrics on goroutines
- Opentracing middleware for tracing context propagation and routine execution
- Redis pubsub implementation
- Zap logger middleware for logging when routines start/finish
- Nats pubsub implementation
- Pluggable, concurrency, safe Directed Acyclical Graph storage for storing business logic
- Policies on what routines channels Routines can publish/subscribe to
## [3][ew: Run stuff everywhere](https://www.reddit.com/r/golang/comments/jazfnq/ew_run_stuff_everywhere/)
- url: https://www.reddit.com/r/golang/comments/jazfnq/ew_run_stuff_everywhere/
---
Hi!

As I just tagged the official v1.0.0, I thought I would post it here:

[https://github.com/kernle32dll/ew](https://github.com/kernle32dll/ew)

ew is inspired by [gr](https://github.com/mixu/gr), and basically allows to execute any command on a tagged bunch of folders.

For example, I automated my go dependency updates with this, which now just boil down to `ew @matag sh -c 'go get -u &amp;&amp; go mod tidy'`. It also includes a shorthand to get an overview of your git repositories via `ew status`.

Feedback is appreciated!
## [4][Calling a Rust function from Go](https://www.reddit.com/r/golang/comments/jaw47e/calling_a_rust_function_from_go/)
- url: https://www.reddit.com/r/golang/comments/jaw47e/calling_a_rust_function_from_go/
---
How to call a Rust function from Go?

Are there libraries that are ready for production the most?
## [5][panicwatch: a panicwrap alternative [feedback appreciated]](https://www.reddit.com/r/golang/comments/jaxmpm/panicwatch_a_panicwrap_alternative_feedback/)
- url: https://www.reddit.com/r/golang/comments/jaxmpm/panicwatch_a_panicwrap_alternative_feedback/
---
github: [https://github.com/grongor/panicwatch](https://github.com/grongor/panicwatch)

I created this package because I didn't like the idea that my actual application runs in the background, and the processes that I launched is only a proxy for my application. It is probably perfectly fine, I just didn't like it. So this library achieves the same goal while keeping the application "in front" by running the "monitoring process" in the background. The monitoring process reads stderr of your application (using some dup magic; also solved for Windows, Arm,...) and scans it for panics. When it finds one, it calls the configured callback and exits. If there are any errors in the watcher process, or if the watcher process dies, configurable callbacks are called so that you can handle the situation as you like.

It was more of a challenge than a need for some solution...but hey, it works nicely :)

I'd appreciate any feedback and comments. Thanks!

**tl;dr:**

* achieves the same goal as panicwrap
* instead of being a proxy, it spawns background process that does the monitoring
* works on *most* platforms/architectures, including Linux, Windows, and Arm
## [6][Serverless, Zero-Trust SSH on Microsoft Azure](https://www.reddit.com/r/golang/comments/jaxat0/serverless_zerotrust_ssh_on_microsoft_azure/)
- url: https://github.com/ThalesGroup/sshizzle
---

## [7][How we decreased one of our APIs response time by 87% and used fewer resources](https://www.reddit.com/r/golang/comments/jadhsq/how_we_decreased_one_of_our_apis_response_time_by/)
- url: https://labs.armut.com/how-we-decreased-one-of-our-apis-response-time-by-87-and-used-less-resources-ce847e83308
---

## [8][[SERIOUS] Question: ORM or Not? What do you prefer?](https://www.reddit.com/r/golang/comments/jaryoo/serious_question_orm_or_not_what_do_you_prefer/)
- url: https://www.reddit.com/r/golang/comments/jaryoo/serious_question_orm_or_not_what_do_you_prefer/
---
Coming from Django and Grails.. I am struggling to find a mature ORM in Go. What do you use? Why do you like it? I have tried a few but they all seem to lack something. 

My Use case is using this in a REST API with some business logic thrown in. 

Please don't bother talking about whether an ORM is actually possible in Go or not.
## [9][Working with Date and Time in Golang](https://www.reddit.com/r/golang/comments/jazhye/working_with_date_and_time_in_golang/)
- url: https://www.reddit.com/r/golang/comments/jazhye/working_with_date_and_time_in_golang/
---
&amp;#x200B;

[mohitkhare.com](https://preview.redd.it/o6qnh5wmx1t51.png?width=670&amp;format=png&amp;auto=webp&amp;s=a879e9d7a083e94ce9339c7216b1a01a8bf81561)

Hey Gophers 👋

Wrote a piece on how to work with Date and Time in Golang.

I had to search for this topic again and again. Finally wrote a post to create a complete guide to date and time in Go. It covers the following topics -

* Fetching epoch timestamps
* Time in multiple timezones
* Modifying and working with date

📗 Read: [https://www.mohitkhare.com/blog/date-time-golang/](https://www.mohitkhare.com/blog/date-time-golang/)

Do share if you have some feedbacks 🙂
## [10][I use go for about 5 years now and I still haven't needed generics. What's wrong with me?](https://www.reddit.com/r/golang/comments/janp4h/i_use_go_for_about_5_years_now_and_i_still_havent/)
- url: https://www.reddit.com/r/golang/comments/janp4h/i_use_go_for_about_5_years_now_and_i_still_havent/
---
I can't understand people saying that lack of generics is a big go disadvantage.

I can imagine couple of tasks where generics are useful but not unavoidable: image processing, intensive math computations.

What tasks are really require generics where it is not "usual way of solving it" but the only (almost) option?

*-- Added day later --*

In the next morning I remembered how I implemented skiplist on interfaces and used `sed` to substitute interface{} to a basic\_type. That is definitely the case.

I'm obviously just not meeting such challenges where generics are needed.
