# golang
## [1][A slightly better version of Go Playground (maybe)](https://www.reddit.com/r/golang/comments/eqc3sa/a_slightly_better_version_of_go_playground_maybe/)
- url: https://www.reddit.com/r/golang/comments/eqc3sa/a_slightly_better_version_of_go_playground_maybe/
---
Hello folks.

I often find myself thinking that I often encounter a situation when I need to do some small prototyping (playing with goroutines, etc) and Go's playground often is faster solution than a dedicated IDE window. Unfortunately [play.golang.org](https://play.golang.org) is very primitive ([goplay.space](https://goplay.space) is better but not much), so I've decided to try to create something a bit better.

After two days without a sleep managed to create some prototype and I want to share it with you.This playground offers file save and upload feature and some basic code autocomplete (package and function level).

Very interested in your opinion :)

&amp;#x200B;

Demo: [http://goplay.x1unix.com/](http://goplay.x1unix.com/)

Source: [https://github.com/x1unix/go-playground](https://github.com/x1unix/go-playground)
## [2][Go Gazette #90: comparing Python, Go and C++](https://www.reddit.com/r/golang/comments/eqaopk/go_gazette_90_comparing_python_go_and_c/)
- url: https://www.gogazette.com/post/90-comparing-python-go-and-c
---

## [3][Why does "go.mod" have it's own format? It could have been JSON or YAML for example, but nope, it was decided a new syntax. I would really like to know the decision behind that.](https://www.reddit.com/r/golang/comments/eqf3dm/why_does_gomod_have_its_own_format_it_could_have/)
- url: https://www.reddit.com/r/golang/comments/eqf3dm/why_does_gomod_have_its_own_format_it_could_have/
---
Just wondering...
## [4][I made something cool in go, wanted to share it here](https://www.reddit.com/r/golang/comments/epxt1j/i_made_something_cool_in_go_wanted_to_share_it/)
- url: https://www.reddit.com/r/golang/comments/epxt1j/i_made_something_cool_in_go_wanted_to_share_it/
---
I made neko, a neat docker app that can stream a browser to multiple clients, you can take control of the keyboard and mouse and browse the web with friends. The server uses  [pion](https://github.com/pion/webrtc) to stream the desktop through webrtc

https://i.redd.it/4qs4uck1wab41.gif

[https://github.com/nurdism/neko](https://github.com/nurdism/neko)
## [5][GitHub - e-wrks/edh: Đ (Edh) - The next-big-things ought to happen with Haskell not C/C++](https://www.reddit.com/r/golang/comments/eqg4p1/github_ewrksedh_đ_edh_the_nextbigthings_ought_to/)
- url: https://github.com/e-wrks/edh
---

## [6][dealing with database null values, pointers vs null-wrappers](https://www.reddit.com/r/golang/comments/eqaam0/dealing_with_database_null_values_pointers_vs/)
- url: https://www.reddit.com/r/golang/comments/eqaam0/dealing_with_database_null_values_pointers_vs/
---
When dealing with null values from a database,  numbers and time value sometimes have an empty for which Go's zero-value does not make sense. Two ways to deal with this are 1) null wrappers, available in some db libraries and standalone libs (e.g. [this](https://github.com/guregu/null)), and 2) pointers to the primitive value. I wanted to ask if anyone has tried both and which you would recommend
## [7][Come vote for your favourite Golang framework!](https://www.reddit.com/r/golang/comments/eqfkxu/come_vote_for_your_favourite_golang_framework/)
- url: https://www.reddit.com/r/golang/comments/eqfkxu/come_vote_for_your_favourite_golang_framework/
---
Feel free to comment on your favourite framework and we will upvote/downvote it!

No discussion of PROs and CONs here, just vote!

Comment format:

Framework name:  
Website:  
Git URL:
## [8][Understanding the advantages of implicit interfaces](https://www.reddit.com/r/golang/comments/eqcfhz/understanding_the_advantages_of_implicit/)
- url: https://www.reddit.com/r/golang/comments/eqcfhz/understanding_the_advantages_of_implicit/
---
Just started exploring Go. In languages like Java which require interfaces to be explicitly specified, the compiler throws an error if a class doesn't implement atleast one method, which makes sense and is helpful in a lot of cases.

Now coming to Go, if an interface has 5 methods and I implement only 4 of them for a struct, the compiler doesn't complain at all, until I try to assign the struct to a variable of that interface type. I am having a hard time understanding why this decision was taken. Can someone please explain to me how this is considered a better design? 

The only way a developer finds out that he partially implemented the interface is when he actually uses it somewhere, isn't it? Isn't there a reasonable and fairly possible chance that a developer might forget to implement a method or two in case of large interfaces?

Isn't explicitly mentioning the intent of implementing an interface a better approach here? Or am I missing something?
## [9][[ASK] Beginner Microservices in Go](https://www.reddit.com/r/golang/comments/eqd868/ask_beginner_microservices_in_go/)
- url: https://www.reddit.com/r/golang/comments/eqd868/ask_beginner_microservices_in_go/
---
Hi, I am learning microservices, currently trying  to build something like Google's microservices demo ([https://github.com/GoogleCloudPlatform/microservices-demo](https://github.com/GoogleCloudPlatform/microservices-demo)) but with Node.js/Apollo GraphQL instead of frontend. So far, I can communicate between them successfully.

Now I've discovered DDD and CQRS stuffs, still trying to understand what they are. I'm confused about those tools, NATS Streaming, Istio and Event Store.

I don't mind learning all those tools, but I'm afraid some of them might be overlapping. 

Currently I have GraphQL, gRPC, Docker and k8s in my stack. Should I add Istio or NATS, or both to apply DDD and CQRS, will they be overlapping? 

Which stack would be more efficient and apply DDD and CQRS;

gRPC, k8s, Istio - [https://github.com/GoogleCloudPlatform/microservices-demo](https://github.com/GoogleCloudPlatform/microservices-demo)

(or)

NATS, gRPC - [https://github.com/shijuvar/go-distributed-sys](https://github.com/shijuvar/go-distributed-sys)

Please enlighten me.
## [10][How Do I Wrap Errors in recursive functions?](https://www.reddit.com/r/golang/comments/eq82vv/how_do_i_wrap_errors_in_recursive_functions/)
- url: https://www.reddit.com/r/golang/comments/eq82vv/how_do_i_wrap_errors_in_recursive_functions/
---
I'm a junior developer, writing software at a startup. I have 1yr of experience in Go. Unfortunately i cannot provide src code.

Wrapping errors is something i do often. I have never done it in a recursive function tho! The jist is this:

I have a struct. Sometimes it can encounter an error, let's say E, on a call it makes, in which case it needs to retry. It is created with a number N of max consecutive retries. It will fail, F, after exceeding N retries

I would like a wrapped message kinda like:
F @ N retries. Error=E

Anyone have an experience, guides, or wisdom for wrapping errors, in recursive functions? 


TL;DR; I would love some insight in wrapping errors in a recursive function in Golang

Cheers!
