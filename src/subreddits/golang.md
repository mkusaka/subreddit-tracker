# golang
## [1][cabify/gotoprom: Type-safe Prometheus metrics builder library for golang](https://www.reddit.com/r/golang/comments/ers1df/cabifygotoprom_typesafe_prometheus_metrics/)
- url: https://github.com/cabify/gotoprom
---

## [2][Generate PlantUML Diagrams from Golang sources or directories](https://www.reddit.com/r/golang/comments/ertcxe/generate_plantuml_diagrams_from_golang_sources_or/)
- url: https://github.com/bykof/go-plantuml
---

## [3][Building a global services network using Go, QUIC and Micro](https://www.reddit.com/r/golang/comments/ertnaa/building_a_global_services_network_using_go_quic/)
- url: https://micro.mu/blog/2019/12/05/building-a-microservices-network.html
---

## [4][Learning Go beyond the Syntax](https://www.reddit.com/r/golang/comments/erg0ob/learning_go_beyond_the_syntax/)
- url: https://www.reddit.com/r/golang/comments/erg0ob/learning_go_beyond_the_syntax/
---
I've worked my way through FreeCodeCamp's tutorial on Go and I believe I've got my head around how to read a piece of Go code but I feel like it leaves me without direction on where to go next.

Yes I can write a calculator app, or I could even hack together an API with the http package, but I feel like that's not how real actually go applications are made.

Does anyone know of good resources I can use to learn how to use Go properly, like professional Go programmers who would use it in their day to day would use it? Even what might seem basic, like how do I structure my files, how do I make an API properly so it's not just a massive cascade of switch statements for routing, are there frameworks that are very well used and should be being used?

I've looked all over the internet as much as I can and all the material just seems to teach the syntax in small examples but that's all, leaving me like I can't really build anything on any scale without it being "not the way go is actually written", so any help would be greatly appreciated.
## [5][Looking for Go help growing a new testing technology](https://www.reddit.com/r/golang/comments/erucqq/looking_for_go_help_growing_a_new_testing/)
- url: https://www.reddit.com/r/golang/comments/erucqq/looking_for_go_help_growing_a_new_testing/
---
There's a new technology - [Servirtium](https://servirtium.dev) - that aims be a multi-language standard for service virtualization. Well, a subset of that field really as this focusses on record/playback.  We've a Java version complete enough and in-use for real (can't say "in production" for testing tech's really, can you). A Python version is advanced enough to use. A Ruby and Kotlin version are underway. **I need someone to get involved with a version written in Go for wider use within the comunity**.

We don't like calling these ports as each language team doesn't need to read or copy the source of any pre-existing other-language version. Indeed the Python version is a deal simpler than the Java one, as we wanted **idiomatically correct for the language** as part of the achievement.

Anyone wanting to get involved? Someone who wants to do something small in opensource-land and perhaps go on to talk at conferences about it? We've a list of [baby steps that you need to do](https://github.com/servirtium/README/blob/master/starting-a-new-implementation.md) for easiest tracking to completion.

So I co-created Selenium back in 2004 (which many use for functional web testing) and hope that means I won't get smacked for using "-ium" in a testing-centric tech name :)
## [6][/r/golang hit 100k subscribers yesterday](https://www.reddit.com/r/golang/comments/er8klk/rgolang_hit_100k_subscribers_yesterday/)
- url: https://redditmetrics.com/r/golang
---

## [7][Looking for advice from Go developers that are working on back-end roles.](https://www.reddit.com/r/golang/comments/erlq5n/looking_for_advice_from_go_developers_that_are/)
- url: https://www.reddit.com/r/golang/comments/erlq5n/looking_for_advice_from_go_developers_that_are/
---
Hey all.

I'm currently working as a web &amp; Android developer, and I've been learning Go in my personal time, building services and tools for the past 4-5 months. The more I get involved with Go, the harder my daily job of dealing with JS and Java/Kotlin is becoming. I don't think it's necessary to explicitly state why this is the case, as I'm sure many people here will understand.

I've been thinking about transitioning into a back-end heavy Go developer position, and I want to understand the overall skill-set and responsibilities that the current market demands beyond the regular job listing copy. 

So, if you're in a job position that's similar, it'd be really good to hear your opinions. Which skills should I hone before I go looking for a job like this?
## [8][Bind keys to entry box in Lxn/walk?](https://www.reddit.com/r/golang/comments/ersbga/bind_keys_to_entry_box_in_lxnwalk/)
- url: https://www.reddit.com/r/golang/comments/ersbga/bind_keys_to_entry_box_in_lxnwalk/
---
I want to run a function when I press "Return" ,rather than with a button
## [9][I'm aware of race condition, but should I care?](https://www.reddit.com/r/golang/comments/erjsks/im_aware_of_race_condition_but_should_i_care/)
- url: https://www.reddit.com/r/golang/comments/erjsks/im_aware_of_race_condition_but_should_i_care/
---
I've recently created arcli, CLI tool for Redmine project management tool. During development, I've encountered on one case where I have race condition with multiple writers and single reader. The code is in [this function](https://github.com/mightymatth/arcli/blob/e4ba74bd3a560caa0c684f9270e2db4d217964b1/cmd/status.go#L28-L75).

Let me explain. I have 7 concurrent API calls which update 7 shared variables with screen writer. When one API call get result, it writes to its variable and signalise screen writer to refresh the screen with new values.

It's obvious that sometimes screen writes reads values while one of go routines writes to the same one. Despite that, I think I shouldn't care because the worst thing that may happen is that screen writer could read old variable value in specific cycle, but it would be corrected in the next cycle.

The program works as it should, it asynchronously updates the screen as results arrive to the go routines. When I try to run program in mode with race condition detector, it says the same as I suspected.

When I tried to synchronise it with mutexes, it would fail to asynchronously update values and show it to the screen because every goroutine would lock on the mutex, screen writer would come last, and all the results will print on the screen in the same time.

My question: Is there a smarter way of synchronising this work to have smooth visual updates without race condition? Should I even care about this type of race condition? If so, how to tell the race detector that I'm fine with that? What's your thoughts about this scenario?

Best regards.
## [10][Dabbling with golang, go-kit, and go code generation](https://www.reddit.com/r/golang/comments/erpurn/dabbling_with_golang_gokit_and_go_code_generation/)
- url: https://www.reddit.com/r/golang/comments/erpurn/dabbling_with_golang_gokit_and_go_code_generation/
---
I'm really new to golang but settled on learning it in order to do work with a microservice project I have going on, but go-kit requires a metric buttload of boilerplate. I thought that it'd help if I had a code generator to help me learn, and go-kit cli got me pretty far but I ended up hitting my own walls.

I decided to spend a weekend giving it a go and although what I came up with is extremely basic and not really useful, it does generate a working go-kit service and I thought that was neat.

If anyone wants to check it out it is here,  [https://github.com/Allar/kitgen](https://github.com/Allar/kitgen) 

I was wondering if anyone knew of any code generators for go-kit that are a bit more straightforward and simpler to use than go-kit cli and go-kit/kitgen that are worth looking into? I am really enjoying go-kit but finding general-purpose resources outside of a variety of blog posts and the generators listed above is something I'm struggling with. I also struggle with the go-kit tutorials, but maybe this is because I'm still new to golang?
