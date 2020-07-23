# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-embed-reddit).
## [2][[Q&amp;A] //go:embed draft design](https://www.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/
---
I posted a draft design today for embedding files into Go programs.

Video: https://golang.org/s/draft-embed-video

Design: https://golang.org/s/draft-embed-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the io/fs draft design](https://golang.org/s/draft-iofs-reddit).
## [3][fullstorydev/grpcurl: a command-line tool that lets you interact with gRPC servers. It's basically curl for gRPC servers](https://www.reddit.com/r/golang/comments/hwbmdx/fullstorydevgrpcurl_a_commandline_tool_that_lets/)
- url: https://www.reddit.com/r/golang/comments/hwbmdx/fullstorydevgrpcurl_a_commandline_tool_that_lets/
---
[https://github.com/fullstorydev/grpcurl](https://github.com/fullstorydev/grpcurl)
## [4][The official ls-lint documentation is here](https://www.reddit.com/r/golang/comments/hwcl1y/the_official_lslint_documentation_is_here/)
- url: https://ls-lint.org
---

## [5][Exactly 10 years ago, Rob Pike spoke at OSCON about Go. Unfortunately, some slides were barely shown, and the video quality wasn't that great. Hence I recorded 16:9 facsimile slides in 1080p.](https://www.reddit.com/r/golang/comments/hvu1hh/exactly_10_years_ago_rob_pike_spoke_at_oscon/)
- url: https://www.youtube.com/watch?v=ZG_ReFc-bFU
---

## [6][Acker: AMQP producer/consumer CLI](https://www.reddit.com/r/golang/comments/hwcu1f/acker_amqp_producerconsumer_cli/)
- url: https://www.reddit.com/r/golang/comments/hwcu1f/acker_amqp_producerconsumer_cli/
---
I made a small CLI application, named Acker, to produce/consume messages to/from AMQP servers, e.g. RabbitMQ. It can be used for debugging queues, consumers and producers:  
[https://github.com/mostafa/acker](https://github.com/mostafa/acker)
## [7][Monolithic Component Architecture in Go](https://www.reddit.com/r/golang/comments/hwfo4v/monolithic_component_architecture_in_go/)
- url: https://www.reddit.com/r/golang/comments/hwfo4v/monolithic_component_architecture_in_go/
---
Caveat: I'm a long-term Java developer, so a lot of my thinking is coloured by that.

I'm coming back to Go and this time things are making a lot more sense to me - yay :)

What I'm trying to do right now is work out a sensible architecture for my project. It's a monolithic web app (monolith because I want to only install one thing), and is going to have a series of components that work together. As such, things that I want to be able to do are:

* Each component will have an internal API that is the only way for other components to interact with it.
* Many components - though not all - will expose HTTP endpoints for the public API.

So, for example, the `users` component might have use cases for: "GetUser", "UpdateUser" and "CreateUser". The first two of these are available both on the internal and HTTP API, but the third is only on the internal API, because it's only ever called by the code that logs in via Google.

So - presumably what I want is for each component to be it's own package. And for each component to only have as public the bits that are on the internal API. So in this case I would have as public: `GetUserUseCase`, `UpdateUserUseCase`, `CreateUserUseCase`, `UserID`, `UserData` and `UserModel`. The actual service that implements those use cases, the repositories that back it, and the HTTP stuff are all private to the package.

I can then repeat this pattern for each component as needed and it all works.

Is this reasonable? Is this in any way Go-like? Or is this too coloured by my Java experience?

I'm particularly concerned that this will mean each package is going to have files covering many different concerns - API, Service Layer, Repositories, HTTP Endpoints, etc - and that's going to make it harder to work with.

Or is there a more Go-like architecture that's recommended for something like this? Is there a standard way that this kind of thing is/should be structured to better fit with how Go works?

Cheers
## [8][I hope Golang cheat sheet will be useful to you](https://www.reddit.com/r/golang/comments/hvylv9/i_hope_golang_cheat_sheet_will_be_useful_to_you/)
- url: https://simplecheatsheet.com/tag/golang-cheat-sheet/
---

## [9][Implementing "Sign In With Google" in Go From Scratch](https://www.reddit.com/r/golang/comments/hvvu9o/implementing_sign_in_with_google_in_go_from/)
- url: https://qvault.io/2020/07/22/how-to-implement-sign-in-with-google-in-golang/
---

## [10][Go2 playground now supports square brackets for Generics](https://www.reddit.com/r/golang/comments/hvpis5/go2_playground_now_supports_square_brackets_for/)
- url: https://www.reddit.com/r/golang/comments/hvpis5/go2_playground_now_supports_square_brackets_for/
---
[https://go2goplay.golang.org](https://go2goplay.golang.org/) 

&gt;// The playground now supports parentheses or square brackets (only one at  
&gt;  
&gt;// a time) for generic type and function declarations and instantiations.  
&gt;  
&gt;// By default, parentheses are expected. To switch to square brackets,  
&gt;  
&gt;// the first generic declaration in the source must use square brackets.

Sample: [https://go2goplay.golang.org/p/7zFKUcpzhvZ](https://go2goplay.golang.org/p/7zFKUcpzhvZ)
## [11][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [12][Sending custom errors in headers to reduce memory](https://www.reddit.com/r/golang/comments/hwacaf/sending_custom_errors_in_headers_to_reduce_memory/)
- url: https://www.reddit.com/r/golang/comments/hwacaf/sending_custom_errors_in_headers_to_reduce_memory/
---
I have a Microservice that makes multiple API calls and creates one json out of all the responses. The response code for the Microservice is supposed to represent the health of the Microservice, so I can’t use it for the actual error code of the multiple API calls if there is one. I return a 200 no matter what. 

Currently, I’m putting a custom error in the body, but this requires that I read the entire response into memory, but that doesn’t utilize Go’s streaming capabilities. This is because in all cases, I read the whole response and exhaust the stream then create another reader pointing to memory of the response. 

I’m not too familiar with what goes on behind the scenes in the http package, so I was wondering if adding a header named ‘API-Error-Code’ with the response  would solve this problem. If there is an error code, I don’t want to read any of the body into memory. 

Would this solve my problem? I’m not 100% sure what r.Headers.get() will do memory wise.
