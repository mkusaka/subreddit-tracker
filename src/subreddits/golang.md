# golang
## [1][go-zero, a cloud native web and rpc framework](https://www.reddit.com/r/golang/comments/inkdyx/gozero_a_cloud_native_web_and_rpc_framework/)
- url: https://www.reddit.com/r/golang/comments/inkdyx/gozero_a_cloud_native_web_and_rpc_framework/
---
## 0. what is go-zero?
go-zero is a web and rpc framework that with lots of engineering practices builtin. It's born to ensure the stability of the busy services with resilience design, and has been serving sites with tens of millions users for years.
go-zero contains simple API description syntax and code generation tool called `goctl`. You can generate Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript from .api files with `goctl`.
Advantages of go-zero:
* improve the stability of the services with tens of millions of daily active users
* builtin chained timeout control, concurrency control, rate limit, adaptive circuit breaker, adaptive load shedding, even no configuration needed
* builtin middlewares also can be integrated into your frameworks
* simple API syntax, one command to generate couple different languages
* auto validate the request parameters from clients
* plenty of builtin microservice management and concurrent toolkits
&lt;img src="https://raw.githubusercontent.com/tal-tech/go-zero/master/doc/images/architecture-en.png" alt="Architecture" width="1500" /&gt;

## 1. Backgrounds of go-zero
At the beginning of 2018, we heavily suffered from frequent downtime. We decided to re-design our system, from monolithic architecture with Java+MongoDB to microservice architecture. After researches and comparison, we chose to:
* Golang based
 * great performance
 * simple syntax
 * proven engineering efficiency
 * extreme deployment experience
 * less server resource consumption
* Self-designed microservice architecture
 * I have rich experience on designing microservice architectures
 * easy to location the problems
 * easy to extend the features

## 2. Design considerations on go-zero
By designing the microservice architecture, we expected to ensure the stability, as well as the productivity. And from just the beginning, we have the following design principles:
* keep it simple
* high availability
* stable on high concurrency
* easy to extend
* resilience design, failure-oriented programming
* try best to be friendly to the business logic development, encapsulate the complexity
* one thing, one way
After almost half a year, we finished the transfer from monolithic system to microservice system, and deployed on August 2018. The new system guaranteed the business growth, and the system stability.

## 3. The implementation and features of go-zero
go-zero is a web and rpc framework that integrates lots of engineering practices. The features are mainly listed below:
* powerful tool included, less code to write
* simple interfaces
* fully compatible with net/http
* middlewares are supported, easy to extend
* high performance
* failure-oriented programming, resilience design
* builtin service discovery, load balancing
* builtin concurrency control, adaptive circuit breaker, adaptive load shedding, auto trigger, auto recover
* auto validation of API request parameters
* chained timeout control
* auto management of data caching
* call tracing, metrics and monitoring
* high concurrency protected

As below, go-zero protects the system with couple layers and mechanisms:
![Resilience](https://raw.githubusercontent.com/tal-tech/go-zero/master/doc/images/resilience-en.png)

## 4. Future development plans of go-zero
* auto generate API mock server, make the client debugging eaisier
* auto generate the simple integration test for the server side just from the .api files

## 5. Installation
Run the following command under your project:
```shell
go get -u github.com/tal-tech/go-zero
```

## 6. Quick Start
0. full examples can be checked out from below:

* [Rapid development of microservice systems](https://github.com/tal-tech/go-zero/blob/master/doc/shorturl-en.md)
* [Rapid development of microservice systems - multiple RPCs](https://github.com/tal-tech/go-zero/blob/master/doc/bookstore-en.md)

1. install goctl
`goctl`can be read as `go control`. `goctl` means not to be controlled by code, instead, we control it. The inside `go` is not `golang`. At the very beginning, I was expecting it to help us improve the productivity, and make our lives easier.
```shell
 GO111MODULE=on go get -u github.com/tal-tech/go-zero/tools/goctl
```
make sure goctl is executable.

2. create the API file, like greet.api, you can install the plugin of goctl in vs code, api syntax is supported.
```go
 type Request struct {
   Name string `path:"name,options=you|me"` // parameters are auto validated
 }
 
 type Response struct {
   Message string `json:"message"`
 }
 
 service greet-api {
   @server(
     handler: GreetHandler
   )
   get /greet/from/:name(Request) returns (Response);
 }
```
the .api files also can be generate by goctl, like below:
```shell
 goctl api -o greet.api
```
3. generate the go server side code
```shell
 goctl api go -api greet.api -dir greet
```
the generated files look like:
```
 ├── greet
 │ ├── etc
 │ │ └── greet-api.yaml // configuration file
 │ ├── greet.go // main file
 │ └── internal
 │ ├── config
 │ │ └── config.go // configuration definition
 │ ├── handler
 │ │ ├── greethandler.go // get/put/post/delete routes are defined here
 │ │ └── routes.go // routes list
 │ ├── logic
 │ │ └── greetlogic.go // request logic can be written here
 │ ├── svc
 │ │ └── servicecontext.go // service context, mysql/redis can be passed in here
 │ └── types
 │ └── types.go // request/response defined here
 └── greet.api // api description file
```
 the generated code can be run directly:
```shell
 cd greet
 go run greet.go -f etc/greet-api.yaml
```
by default, it's listening on port 8888, while it can be changed in configuration file.
you can check it by curl:
```shell
 curl -i http://localhost:8888/greet/from/you
```
the response looks like:
```http
 HTTP/1.1 200 OK
 Date: Sun, 30 Aug 2020 15:32:35 GMT
 Content-Length: 0
```

4. Write the business logic code
* the dependencies can be passed into the logic within servicecontext.go, like mysql, reds etc.
* add the logic code in logic package according to .api file

5. Generate code like Java, TypeScript, Dart, JavaScript etc. just from the api file
```shell
 goctl api java -api greet.api -dir greet
 goctl api dart -api greet.api -dir greet
 …
```

## 7. Benchmark
![benchmark](https://raw.githubusercontent.com/tal-tech/go-zero/master/doc/images/benchmark.png)
[Checkout the test code](https://github.com/smallnest/go-web-framework-benchmark)

## 8. Documents
* [Rapid development of microservice systems](https://github.com/tal-tech/go-zero/blob/master/doc/shorturl-en.md)
* [Rapid development of microservice systems - multiple RPCs](https://github.com/tal-tech/go-zero/blob/master/doc/bookstore-en.md)

## 9. Project Home
[https://github.com/tal-tech/go-zero](https://github.com/tal-tech/go-zero)
## [2][I wrote a fast concurrent fairly strong RNG (based on xoshiro256**). Much faster than the builtin random package.](https://www.reddit.com/r/golang/comments/inb8u8/i_wrote_a_fast_concurrent_fairly_strong_rng_based/)
- url: https://github.com/Villenny/fastrand64-go
---

## [3][Gold, a local docs server and code reader, v0.1.0 released. Added several new features.](https://www.reddit.com/r/golang/comments/ini0xj/gold_a_local_docs_server_and_code_reader_v010/)
- url: https://github.com/go101/gold
---

## [4][The httpbaselinetest package provides a framework for recording requests, responses, and the expected database changes.](https://www.reddit.com/r/golang/comments/inihgd/the_httpbaselinetest_package_provides_a_framework/)
- url: https://github.com/trussworks/httpbaselinetest
---

## [5][I needed an in-browser Go compiler a while back, but couldn't find one, so I made one. Hopefully someone finds it useful. It's based on github.com/ccbrown/wasm-go-playground and currently Work-In-Progess but functional.](https://www.reddit.com/r/golang/comments/in46lo/i_needed_an_inbrowser_go_compiler_a_while_back/)
- url: https://tr-slimey.github.io/IBGC/
---

## [6][Building a game engine with golang!](https://www.reddit.com/r/golang/comments/inlvqi/building_a_game_engine_with_golang/)
- url: https://link.medium.com/bzx0GDeOy9
---

## [7][[Help] http.ListenAndServe is blocking or am I doing something wrong?](https://www.reddit.com/r/golang/comments/inlgto/help_httplistenandserve_is_blocking_or_am_i_doing/)
- url: https://www.reddit.com/r/golang/comments/inlgto/help_httplistenandserve_is_blocking_or_am_i_doing/
---
```
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", nil)
	fmt.Println("should print this")
}

```

The "should print this" is not executed

Help would be appreciated, thank you
## [8][Trubka v3.2.0 is out! 🚀](https://www.reddit.com/r/golang/comments/inhyxq/trubka_v320_is_out/)
- url: https://www.reddit.com/r/golang/comments/inhyxq/trubka_v320_is_out/
---
&amp;#x200B;

[Trubka](https://preview.redd.it/mvw9kylf9hl51.jpg?width=481&amp;format=pjpg&amp;auto=webp&amp;s=b11f11b9b72d67f96ce4d725dfde76acc6c3c9ac)

**Trubka** is a Kafka CLI tool built in [Go](https://go.dev/) which gives you everything you need to

* Manage, query and troubleshoot your Kafka clusters.
* Consume [protocol buffer](https://developers.google.com/protocol-buffers/) and plain text messages from Kafka.
* Publish protocol buffer and plain text messages to Kafka.

&amp;#x200B;

A mini release of Trubka with a few new handy features is now publicly available:

* Like the protocol buffer publisher, the plain text producer also supports data templates to push randomly generated messages to Kafka (Click [here](https://github.com/xitonix/trubka/wiki/Publish-to-Kafka#templates-and-random-data-generators) to find out more about data templates in trubka).
* A new flag has been added to the produce commands (--sleep) which allows us to put a gap between messages. No more for;do;sleep;done in bash ;)
* Last but not least, the produce commands now support **--count=0** to allow publishing to Kafka indefinitely.

Please install the new version, log bugs, suggest new features, share, contribute and have fun. 🥳
## [9][Manejo de errores en Go vs Rust](https://www.reddit.com/r/golang/comments/inlab2/manejo_de_errores_en_go_vs_rust/)
- url: https://emanuelpeg.blogspot.com/2020/09/manejo-de-errores-en-go-vs-rust.html#.X1TXeXfFhfw.reddit
---

## [10][I wrote a piece about (lightweight) design patterns in Go, hope anyone finds it useful](https://www.reddit.com/r/golang/comments/iml937/i_wrote_a_piece_about_lightweight_design_patterns/)
- url: https://link.medium.com/tGADr7pOv9
---

