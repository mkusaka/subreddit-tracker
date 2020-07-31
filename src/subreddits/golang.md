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
## [2][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [3][Context should go away for Go 2 (2017)](https://www.reddit.com/r/golang/comments/i0vr6h/context_should_go_away_for_go_2_2017/)
- url: https://faiface.github.io/post/context-should-go-away-go2/
---

## [4][Proposal: database/sql: ScannerContext, ValuerContext support, any idea?](https://www.reddit.com/r/golang/comments/i14syh/proposal_databasesql_scannercontext_valuercontext/)
- url: https://github.com/golang/go/issues/40511
---

## [5][hashicorp nomad](https://www.reddit.com/r/golang/comments/i16vm2/hashicorp_nomad/)
- url: https://www.reddit.com/r/golang/comments/i16vm2/hashicorp_nomad/
---
Does anyone here use Nomad? I like the fact its written in Go and very easy to deploy. I use it for Legacy Applications (C++ and Java). I don't need anything fancy like kubernetes. So far I am impressed with it.. Any other good alternatives (hopefully written in Go)?
## [6][Different approaches to HTTP routing in Go](https://www.reddit.com/r/golang/comments/i0s89s/different_approaches_to_http_routing_in_go/)
- url: https://benhoyt.com/writings/go-routing/
---

## [7][Any advice for authentication in a gRPC -&gt; gRPC web project?](https://www.reddit.com/r/golang/comments/i12gbu/any_advice_for_authentication_in_a_grpc_grpc_web/)
- url: https://www.reddit.com/r/golang/comments/i12gbu/any_advice_for_authentication_in_a_grpc_grpc_web/
---
In an ideal world I'd use a Session store and set http only cookies, and validate the cookies on each request. However, I don't think this is possible with gRPC.

Where should I start? Any advice is appreciated.
## [8][Pooling bufio.Reader](https://www.reddit.com/r/golang/comments/i16egt/pooling_bufioreader/)
- url: https://www.reddit.com/r/golang/comments/i16egt/pooling_bufioreader/
---
Need some thoughts on this. Let's say I have an io.Reader object (assume tcp connection) and I assign a new bufio.Reader for buferred i/o, which allocates 4096 bytes of memory by default. If there are multiple such connections (assume 1k per min), Go will allocate a 4096 bytes buffer for each of them. Assuming most connections are short lived, this can hurt performance as we are allocating new memory block every time there is a new connection. Won't it be better to have a pooled buffer, so some of the previous allocations can be reused?

Similar problem can also occur in case there are multiple persistent connections and only few are sending data at a time, thus most bufio buffers are underutilized.
## [9][Is there any example on using MongoDB amboy?](https://www.reddit.com/r/golang/comments/i14q29/is_there_any_example_on_using_mongodb_amboy/)
- url: https://www.reddit.com/r/golang/comments/i14q29/is_there_any_example_on_using_mongodb_amboy/
---
I'm trying to look into some Queue and Job Monitoring tools with set of operations for Pausing/Resuming/Killing/Migrating.

I stumbled with [https://github.com/mongodb/amboy](https://github.com/mongodb/amboy) but I didn't find any examples.

Anyone tried using it or any resources where I can get the example
## [10][GoLand 2020.2 is out! Problems View Window, new inspections and more in the latest release await you!](https://www.reddit.com/r/golang/comments/i0ncyf/goland_20202_is_out_problems_view_window_new/)
- url: https://twitter.com/GoLandIDE/status/1288848194894430209?s=19
---

## [11][Introducing Golang Tweetletter Bot: Get the most popular tweets from the Go Twitterverse](https://www.reddit.com/r/golang/comments/i0z47d/introducing_golang_tweetletter_bot_get_the_most/)
- url: https://twitter.com/GTweetletter/status/1289010731958980609
---

## [12][GoTime: The Latest on Generics](https://www.reddit.com/r/golang/comments/i0v0ot/gotime_the_latest_on_generics/)
- url: https://changelog.com/gotime/140
---

