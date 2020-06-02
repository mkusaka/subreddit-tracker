# golang
## [1][Go 1.14.4 is downloadable](https://www.reddit.com/r/golang/comments/gv282q/go_1144_is_downloadable/)
- url: https://www.reddit.com/r/golang/comments/gv282q/go_1144_is_downloadable/
---
[https://golang.org/dl/](https://golang.org/dl/)
## [2][I made a simple tutorial on using Go and GraphQL](https://www.reddit.com/r/golang/comments/guum44/i_made_a_simple_tutorial_on_using_go_and_graphql/)
- url: https://www.youtube.com/watch?v=YZjy3Wj0BZg&amp;t=1495s
---

## [3][Source code for progapanda.org: a severely over-engineered backend developer's personal website built with tview and websockets [issue help wanted!]](https://www.reddit.com/r/golang/comments/gv6ovx/source_code_for_progapandaorg_a_severely/)
- url: https://github.com/progapandist/progapanda.org
---

## [4][Smocker, a simple and efficient HTTP mock server for functional tests](https://www.reddit.com/r/golang/comments/gulpux/smocker_a_simple_and_efficient_http_mock_server/)
- url: https://www.reddit.com/r/golang/comments/gulpux/smocker_a_simple_and_efficient_http_mock_server/
---
[Smocker](https://github.com/Thiht/smocker)

Hey guys :)

[My colleague](https://github.com/gwleclerc) and [I](https://github.com/Thiht) just released [Smocker](https://smocker.dev) on [GitHub](https://github.com/Thiht/smocker). It’s a mock server we use at our company to write functional tests of our microservices.

Smocker is tailored for our main use-case: a microservice environment.

After a few months of use in production, we’re pretty confident in its stability and proud of its main features:

* Mock declaration as code in YAML instead of JSON. It’s very handy to declare JSON payloads because you don’t need to escape anything,
* Great user interface, extremely useful to figure out what’s going on at any point, in the call history,
* Dynamic responses system, to compute mock responses dynamically using templates or Lua,
* Session system, to be able to import and export the current history and mock state, a very useful debugging feature.
* Proxy feature, when some calls don’t need to be mocked but simply forwarded,
* Nice documentation, we tried to polish it as much as possible,
* Good overall performances, we used other solutions before and sometimes had latency or crash issues from the mock server.

Note that it’s written in Go, the codebase might be worth a read :)

We hope it can be useful to others!
## [5][Side-gigs, freelancing on the side etc. - where, and how?](https://www.reddit.com/r/golang/comments/gv7dtm/sidegigs_freelancing_on_the_side_etc_where_and_how/)
- url: https://www.reddit.com/r/golang/comments/gv7dtm/sidegigs_freelancing_on_the_side_etc_where_and_how/
---
I've got like 5 years of experience as a software engineer, the last two primarily using Go. I have quite a lot of free-time and will the next few months. I wouldn't mind pocketing some extra cash for my savings in this time as my wife unfortunately was struck by a lay-off :-/

Any sites you people recommend? I've checked out Upwork and Fiverr.. But they seem kind of like race-to-the-buttom kind of places. I live in a medium CoL place and my hourly rate for this to be worth it would be like 40USD.. Does it even make sense to pursue this idea then?
## [6][Errors with stack trace](https://www.reddit.com/r/golang/comments/gv69tg/errors_with_stack_trace/)
- url: https://rafallorenz.com/go/go-error-stack-trace/
---

## [7][Pakkero: a packer written in Go](https://www.reddit.com/r/golang/comments/gulnwo/pakkero_a_packer_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/gulnwo/pakkero_a_packer_written_in_go/
---
Hi all, in my free time I've developed this project to learn Golang :)

It is a packer that launches another program that is stored compressed and encrypted inside the binary itself, without passing for the storage.

All the functionalities are explained in the README, I'm curious to hear from you comments and improvements on the project.

https://github.com/89luca89/pakkero
## [8][AsyncWait: Wait don't sleep in your tests](https://www.reddit.com/r/golang/comments/gv498z/asyncwait_wait_dont_sleep_in_your_tests/)
- url: https://github.com/rShetty/asyncwait
---

## [9][What makes writing sync style api handlers affordable in Go?](https://www.reddit.com/r/golang/comments/gv0zzm/what_makes_writing_sync_style_api_handlers/)
- url: https://www.reddit.com/r/golang/comments/gv0zzm/what_makes_writing_sync_style_api_handlers/
---
It must be a noob question, but I have some experience building servers with node / scala / java and clients using swift / js / dart which all uses some notion of async operation like callbacks and Future/Promises. 

I know that go http runs the handlers in goroutines but it seems to me that it might be less scheduler-friendly for devs to write them in sync style (which seems to be the recommended way) rather than explicitly stating which operations are async and which are sync. 

For example, in a code that goes like

sync1
sync2
async1
sync3
async2
sync4

I feel like the scheduler in a explicit async style won’t make the execution yield after sync3 before async2 but with goroutines it might as it doesn’t have any distinction between the two. 

Am I missing something here? What separates goroutines from other language’s async primitives that we’re allowed to just write everything in sync style?
## [10][I made a simple serverless framework to scale big projects using lambda](https://www.reddit.com/r/golang/comments/gv02px/i_made_a_simple_serverless_framework_to_scale_big/)
- url: https://github.com/msfidelis/serverless-architecture-boilerplate-go
---

