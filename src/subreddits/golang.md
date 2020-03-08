# golang
## [1][I made a port scanner in go without dependencies!](https://www.reddit.com/r/golang/comments/ff08cd/i_made_a_port_scanner_in_go_without_dependencies/)
- url: https://github.com/R4yGM/netscanner
---

## [2][A Base64 Captchas Manager](https://www.reddit.com/r/golang/comments/ffa7wv/a_base64_captchas_manager/)
- url: https://github.com/clevergo/captchas
---

## [3][My experience with learning Golang](https://www.reddit.com/r/golang/comments/fesci5/my_experience_with_learning_golang/)
- url: https://pmihaylov.com/learning-golang-experience/
---

## [4][Thoughts on Using Micro](https://www.reddit.com/r/golang/comments/fex43s/thoughts_on_using_micro/)
- url: https://www.reddit.com/r/golang/comments/fex43s/thoughts_on_using_micro/
---
Presently at work I am using a mix of a home grown framework for some services and GOKit for others. I’ve been looking to evaluate using GO micro and spent a day looking through examples and documentation. Looking at the documentation I’ve got the impression that micro is fairly opinionated and really works well when you buy into its philosophy and way of doing things. I’m not sure how we’ll GO micro would work in an environment where we already have nearly 100 services with half in GOKit and half in a home grown framework. In addition gRPC doesn’t make sense for our public facing APIs that are called from client side JavaScript applications which GO micro seems to be heavy centered around gRPC. 

Does anyone have experience with using GO micro for production use, particularly in an environment where nearly every other service is not using micro? Is it even worth prototyping?
## [5][CertMagic v0.10.0 released - CertMAutomatic HTTPS for any Go program: fully-managed TLS certificate issuance and renewal.](https://www.reddit.com/r/golang/comments/fet7wz/certmagic_v0100_released_certmautomatic_https_for/)
- url: https://github.com/caddyserver/certmagic
---

## [6][Go Modules VS GOPATH - Creating the Workspace](https://www.reddit.com/r/golang/comments/ff6eet/go_modules_vs_gopath_creating_the_workspace/)
- url: https://www.reddit.com/r/golang/comments/ff6eet/go_modules_vs_gopath_creating_the_workspace/
---
Noobie question, I have been paying attention to Go for a while now but never learned it. I know it used to require all `src/` to be in the GOPATH, but in recent versions, Go has Go Modules to use. I was reading the [How to Code](https://golang.org/doc/code.html) blog post and the [Using Go Modules](https://blog.golang.org/using-go-modules) one and was curious - is it common practice to create all go projects outside of the GOPATH now? 

In the same vein, I never knew how people incorporated their VCS with their projects. Am I supposed to manually name the directory after [`github.com/{username}/{project}`](https://github.com/{username}/{project})? Or is it better if I init it first and retrieve it with either `go get` or `git clone`.

Advice for how to have my workspace would be much appreciated.
## [7][Copy-on-write variables, Go and Newsqueak](https://www.reddit.com/r/golang/comments/ff5ppz/copyonwrite_variables_go_and_newsqueak/)
- url: https://www.reddit.com/r/golang/comments/ff5ppz/copyonwrite_variables_go_and_newsqueak/
---
I have been very interested in the concept of copy-on-write for some time and it jumped out at me when I read about how Newsqueak dispenses with pointers.

I am curious to make a teensy modification to Go syntax with a preprocessor that automates the implementation of this style of variable handling. But before I even start writing a parser for this (it will be a little complicated), I am applying the principle to my code to learn how this pattern works.

Anyone who is a Go programmer and has written any concurrent code would know about the mess of options in front of you - atomics, mutexes, channels... When to use this, when to use that...

If variables passed into functions could not be mutated you wouldn't need atomics or mutexes. It would clean up concurrency so much. Concurrency combined with imperative style code is a quagmire. The only down side is sometimes following this pattern slavishly might result in more copying than is good for performance.

It's upon this point that I need to experiment with it. This is the rules I have got so far from changing about 8 functions in one source file that is all closely connected to this:

1. All parameters are pointers, except slices, interfaces and maps. These get a name with a P suffix to signify they are parameters and are copied into values as required (if it's already a pointer type leave it alone).

2. All return values are values, unless they can't be, for reasons alluded to above and because the pointer is required to call pointer methods.

3. Not all values in parameters have to be values. Probably numbers and booleans can stay as values, since they often don't escape the function anyway.

I have a theory that writing code this way will assist the escape analysis process and improve the number of things that stay on the stack. It should also automatically reduce the amount of problems I have that normally I reach for atomics or mutexes to deal with.

Once I understand this pattern thoroughly and have applied it extensively I will be looking at creating a variant of the Go language that eliminates pointers and does what I have been doing manually to study it automatically. I should start to form some useful rules of thumb about what kinds of variables are dangerous or expensive to work with and how to minimise this in the architecture of the program.
## [8][how to install Golang on mac](https://www.reddit.com/r/golang/comments/ff9oj9/how_to_install_golang_on_mac/)
- url: https://www.youtube.com/watch?v=EPTIZw9pxdw&amp;feature=share
---

## [9][Using templates to render outgoing JSON.](https://www.reddit.com/r/golang/comments/few1fb/using_templates_to_render_outgoing_json/)
- url: https://www.reddit.com/r/golang/comments/few1fb/using_templates_to_render_outgoing_json/
---
I've been toying around with the idea of using Go's templating capabilities to "massage" end-consumer JSON output into the right format. Based on my experience, one always comes to a point, where marshaling the domain data to JSON directly becomes a problem. One either displays too much (certain fields need to be hidden) or too little (additional stuff needs to be added before sending to the client). Sure, one can always use a struct or a map dedicated to represent the end-consumer data to be JSON-ified. This is how I still do it. However, I always find it tedious to convert one struct into another, especially, ones that have nested slices of types that have to be "massaged" on their own.

Somehow, the idea of using templating for this sounded more appealing, because a template has all the necessary means to allow for consistent top-to-bottom rendering. One can easily handle the nested slices too, without breaking the cognitive flow of the conversion. Moreover, templates can be referenced and reused, which can help reduce unnecessary copying and pasting. 

At this point, I'm not worried about performance issues, because I only talk about end-consumer JSON. This isn't supposed to be a means of converting between struct apps inside the application, but between multiple applications. 

Feedback is more than appreciated.
## [10][r8limiter - envoy pluggable rate limiting service - Feedback is needed :)](https://www.reddit.com/r/golang/comments/fet1qj/r8limiter_envoy_pluggable_rate_limiting_service/)
- url: https://github.com/samueltorres/r8limiter
---

