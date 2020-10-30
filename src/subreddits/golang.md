# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Protobuf vs. Gob benchmark](https://www.reddit.com/r/golang/comments/jku5dl/protobuf_vs_gob_benchmark/)
- url: https://www.reddit.com/r/golang/comments/jku5dl/protobuf_vs_gob_benchmark/
---
Our company receives many huge JSON data and we used [Protocol Buffers](https://developers.google.com/protocol-buffers) (Protobuf) to efficiently persist those JSON on disk. This came with new challenges since Protobuf needs a definition, which needs to be maintained. First, we wanted to generate the definition manually. Then we discovered package [gob](https://golang.org/pkg/encoding/gob/) from the Go standard library.

Gob is a super easy to use binary encoder. You just pass in ordinary Go values and in return, you get the binary encoding of those values.

An important factor for our company is read/write speed of JSON serialization. For that, I created benchmark tests comparing Protobuf read/write and gob read/write I'd like to share: [https://github.com/ndabAP/proto-vs-gob-bench](https://github.com/ndabAP/proto-vs-gob-bench)

Results:

```
BenchmarkGobWrite-8                  121           9679732 ns/op
BenchmarkGobRead-8                    60          19264177 ns/op
BenchmarkProtobufWrite-8              13          88254390 ns/op
BenchmarkProtobufRead-8               26          46792472 ns/op
```
## [3][Go is Gone](https://www.reddit.com/r/golang/comments/jkdzsq/go_is_gone/)
- url: https://www.reddit.com/r/golang/comments/jkdzsq/go_is_gone/
---
Go is being discarded from my team's tool chest. I've been coding in Go every day for 2 years writing command-line utilities accessing REST APIs and gRPC services, and back-ends serving REST APIs and gRPC on both VMs and Google Cloud Platform via App Engine and Cloud Functions. Now it's gone.

Reasons? As officially told - manager wants to:

* standardize on a single language for both front and back ends for sharing of code
* finding Go developers for hire is too tough
* Go's concurrency model is lacking

We're moving to TypeScript because of code-sharing availability of libraries (above), easier to find developers, and concurrency models are better.

Again. This is the "official" line.

I'm currently evaluating my position on this and wondering if I should at least entertain some interviews with companies or teams where Go is a tool or stay around and add TypeScript to my personal tool chest and see if Go can be reintroduced at some point.
## [4][Found a cool cross-platform native desktop GUI library for Go called GoVCL, wraps around Lazarus](https://www.reddit.com/r/golang/comments/jku1vw/found_a_cool_crossplatform_native_desktop_gui/)
- url: https://z-kit.cc/
---

## [5][Complete //go:embed support has just landed!](https://www.reddit.com/r/golang/comments/jkdjgw/complete_goembed_support_has_just_landed/)
- url: https://github.com/golang/go/commit/25d28ec55aded46e0be9c2298f24287d296a9e47
---

## [6][Bubbles: Beautiful components for the BubbleTea TUI framework](https://www.reddit.com/r/golang/comments/jkr2pa/bubbles_beautiful_components_for_the_bubbletea/)
- url: https://github.com/charmbracelet/bubbles
---

## [7][A basic blockchain project that allows automatic peer discovery in golang.](https://www.reddit.com/r/golang/comments/jksgeb/a_basic_blockchain_project_that_allows_automatic/)
- url: https://www.reddit.com/r/golang/comments/jksgeb/a_basic_blockchain_project_that_allows_automatic/
---
👋 Hi, Everyone

I made a blockchain project that allows automatic peer discovery, fault tolerance, and has a CLIUI to see how peers interact with each other in realtime using #golang. It uses libp2p for it's networking and works pretty well 


Github
https://github.com/TheDhejavu/the-crypto-project

Demo:
https://youtu.be/Kg0v6MdVs3I
## [8][I just published my first cli tool written in Go and I would love your feedback! It enables you to manage mysql dbs from thew command-line without the burden of needing to maintain a connection to the db host. It also simplifies managing and connecting to multiple different db hosts from the cli.](https://www.reddit.com/r/golang/comments/jkwmmm/i_just_published_my_first_cli_tool_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/jkwmmm/i_just_published_my_first_cli_tool_written_in_go/
---
[https://github.com/georgiotunson/ssql](https://github.com/georgiotunson/ssql)
## [9][Does anyone do type any interface{}?](https://www.reddit.com/r/golang/comments/jktass/does_anyone_do_type_any_interface/)
- url: https://www.reddit.com/r/golang/comments/jktass/does_anyone_do_type_any_interface/
---
Just wondering if anyone else declares interface{} as 

type any interface{}
## [10][Want to see full body HTTP requests/responses from your Go server using eBPF?](https://www.reddit.com/r/golang/comments/jki4k3/want_to_see_full_body_http_requestsresponses_from/)
- url: https://www.reddit.com/r/golang/comments/jki4k3/want_to_see_full_body_http_requestsresponses_from/
---
At Pixie, we are working on building a platform that gives you automatic observability into your Go apps on K8s clusters without any code changes / instrumentation. 

We're writing a [**series**](https://blog.pixielabs.ai/ebpf) sharing how you can debug applications in prod using eBPF with [**open source code**](https://pixielabs.ai/).  

In today's [**blog post**](https://blog.pixielabs.ai/ebpf-http-tracing/) we show how you can use eBPF to see full body HTTP requests/responses from a Golang HTTP server (simpleHTTP) without any code changes. 

https://i.redd.it/re3e5r07c3w51.gif
## [11][help with handling []primitive.M](https://www.reddit.com/r/golang/comments/jku8mw/help_with_handling_primitivem/)
- url: https://www.reddit.com/r/golang/comments/jku8mw/help_with_handling_primitivem/
---
Hi,

I'm doing a project with the mongodb go driver where I want to access the data from my db. I can get the data as a type of \[\]primitive.M and print it out. But I need to access specific key values in my data. How to I do that?

Thanks!
