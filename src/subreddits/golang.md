# golang
## [1][Learn Go test-first with 'For the Love of Go: Fundamentals'](https://www.reddit.com/r/golang/comments/jelx8o/learn_go_testfirst_with_for_the_love_of_go/)
- url: https://bitfieldconsulting.com/books/fundamentals
---

## [2][What is the most production ready GraphQL server library written for Golang?](https://www.reddit.com/r/golang/comments/jemm91/what_is_the_most_production_ready_graphql_server/)
- url: https://www.reddit.com/r/golang/comments/jemm91/what_is_the_most_production_ready_graphql_server/
---

## [3][We're Hiring Remotely @Spacelift! 🚀](https://www.reddit.com/r/golang/comments/je3qug/were_hiring_remotely_spacelift/)
- url: https://www.reddit.com/r/golang/comments/je3qug/were_hiring_remotely_spacelift/
---
Hey, we’re hiring at [spacelift.io](https://spacelift.io)!

100% Remote (office in Warsaw if you fancy, but we’re working in a remote-first fashion). We’re looking for developers in Europe, which are available most of the working day in CET.

$50-120k per annum  
Mid-Senior

We’re a VC-funded startup building an automation platform for Infrastructure-as-Code, adding a Policy-as-Code layer above it, in order to make IaC usable in bigger companies, where you have to take care of state consistency, selective permissions, a usable git flow, etc.

On the backend we’re using 100% Go with AWS primitives.  
We’re looking for developers who like doing DevOps’ish stuff sometimes (because in a way it’s the spirit of our company) , though that’s not required. Ideally you’d have experience working with an IaC tool, i.e. Terraform, Ansible, CloudFormation, or SaltStack.

Overall we have a deeply technical product, trying to build something customers love to use, and already have multiple happy and satisfied potential clients.

We promise interesting work, the ability to open source parts of the project which don’t give us a business advantage, as well as sane working hours.

If that sounds like fun to you, or you have any questions, send me a private message with your email, and I’ll get back to you! You can also directly email me at kubam at spacelift dot io.

PS: Fear not, no whiteboarding!
## [4][Determine prominent colors in a picture, your first AWS Lambda in Go](https://www.reddit.com/r/golang/comments/jem8ri/determine_prominent_colors_in_a_picture_your/)
- url: https://buddy.works/tutorials/determine-prominent-colors-in-a-picture-your-first-aws-lambda-in-go
---

## [5][BigBro: An extensible monitoring tool for user defined services and protocols](https://www.reddit.com/r/golang/comments/je8rfd/bigbro_an_extensible_monitoring_tool_for_user/)
- url: https://www.reddit.com/r/golang/comments/je8rfd/bigbro_an_extensible_monitoring_tool_for_user/
---
Hi everyone, I would share this beta project that can be used to monitoring services (and different protocols).

The tool is lightweight and therefore it can be used also in embedded devices, such as Raspberry Pi (zero to 4) or similar.

In addition, the tool is extensible in terms of custom services/protocols implementation.

The project is actually divided in two parts:

* BigBro ([https://github.com/bigbroproject/BigBro](https://github.com/bigbroproject/BigBro)) - is the main tool and provides an user interface via webserver and it is based on BigBroCore
* BigBroCore ([https://github.com/bigbroproject/BigBroCore](https://github.com/bigbroproject/BigBroCore)) - is the core of the project, basically is BigBro without user interface.

Let us know about suggestions, issues or anything about the project!
## [6][What is a "sentinel" element?](https://www.reddit.com/r/golang/comments/jehceb/what_is_a_sentinel_element/)
- url: https://www.reddit.com/r/golang/comments/jehceb/what_is_a_sentinel_element/
---
Hey! I'm looking at the source code of the `container/list` package in the standard library and came across [this line](https://golang.org/src/container/list/list.go#L49).

I'm curious about the idea behind this implementation where we keep track of an element that's not actually an element in the linked list. Can someone explain?
## [7][Golang library for Google Analytics](https://www.reddit.com/r/golang/comments/jedg5u/golang_library_for_google_analytics/)
- url: https://www.reddit.com/r/golang/comments/jedg5u/golang_library_for_google_analytics/
---
Hey folks! I threw together a quick library that supports sending metrics to Google Analytics from Go applications as well as generating URLs that you can embed in emails and other applications. No more needing to remember or look up the parameters. Just put in your information and go...

Let me know what you think!

[https://github.com/mjpitz/go-ga](https://github.com/mjpitz/go-ga)
## [8][VSCode Debugging](https://www.reddit.com/r/golang/comments/jej4rp/vscode_debugging/)
- url: https://www.reddit.com/r/golang/comments/jej4rp/vscode_debugging/
---
Delve doesn't seem to work for me on my Ubuntu virtual machine (not WSL, an actual VM).

I get a "Failed to continue" error, and my debug console says "exit status 2, Process exiting with code: 1"

The program executes fine without the debugger.

I tried looking at the following [https://github.com/golang/vscode-go/blob/master/docs/debugging.md](https://github.com/golang/vscode-go/blob/master/docs/debugging.md)

Ran the following in terminal: go get github.com/go-delve/delve/cmd/dlv

And then updated dlv from the command palette, but it's still not working. I should also say, I do have the Go extension installed.

Has anyone ever experienced something similar? I'm not getting any substantial error messages, so I have not been able to find much by just googling.
## [9][Implementing fast lightweight containers in Go with bst and btrfs (Part 1)](https://www.reddit.com/r/golang/comments/je1kwx/implementing_fast_lightweight_containers_in_go/)
- url: https://snai.pe/posts/lightweight-containers-part-1
---

## [10][Rewrite Raspberry Pi Python project in Go? (GPIO)](https://www.reddit.com/r/golang/comments/je38sl/rewrite_raspberry_pi_python_project_in_go_gpio/)
- url: https://www.reddit.com/r/golang/comments/je38sl/rewrite_raspberry_pi_python_project_in_go_gpio/
---
I'm thinking about rewriting an application in Go, that I built with Python and PyQt for the Raspberry Pi. It needs to access the GPIO pins and has a GUI to control them. Is this generally a good idea? I would use Go because of the easier use of concurrency and because I just like the language. For the GUI I would use fyne, but which library should I use for the GPIOs? I've seen a lot of different ones, some of them aren't maintained anymore, so which GPIO library would you recommend?
