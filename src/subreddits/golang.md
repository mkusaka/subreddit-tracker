# golang
## [1][Caddy 2 is released](https://www.reddit.com/r/golang/comments/gdfnlj/caddy_2_is_released/)
- url: https://caddyserver.com/v2
---

## [2][mathetake/gasm: A WASM Virtual Machine for Gophers](https://www.reddit.com/r/golang/comments/gdwpow/mathetakegasm_a_wasm_virtual_machine_for_gophers/)
- url: https://www.reddit.com/r/golang/comments/gdwpow/mathetakegasm_a_wasm_virtual_machine_for_gophers/
---
Hi all,

I've just published my homemade WASM virtual machine purely written in Go. I did this just for fun and for leaning WASM specification.  [https://github.com/mathetake/gasm](https://github.com/mathetake/gasm)

With gasm, you can embed VMs in your go programs without any dependencies like cgo. Calling host functions from WASM modules and execution of exported functions in WASM modules are supported.

Although it seems to have tons of bugs, I would like to share this personal project as it works fine with basic examples.

Thanks!
## [3][Fairly seasoned developer, interested in the multithreading performance benefits of golang over C++.](https://www.reddit.com/r/golang/comments/gdrm71/fairly_seasoned_developer_interested_in_the/)
- url: https://www.reddit.com/r/golang/comments/gdrm71/fairly_seasoned_developer_interested_in_the/
---
Hi, 

I'm part of a team that maintains a large-ish "AI" service which executes hundreds of thousands of queries every minute. The service is fairly modern, managed and deployed on Azure with Kubernetes and docker containers in datacenters over the world.

The individual backend microservice containers are built in C++ on top of gRPC. Every backend container handles many concurrent requests. 

We place quite a bit of emphasis on performance, and routinely deploy updates that improve scale, which directly impacts our bottom line in terms of margin, or sometimes we pass those savings onto the customer. I myself have optimized some of the core pieces with intrisics that we do have available on some of the high-end azure VMs just to squeeze some 5% more perf. This is to say that we ***do*** spend the 80% extra effort to get the 20% extra benefit.

That said, the multithreading aspect of the service has not been given too much attention. Since each microservice handles many requests, our concurrency is handled by a barebones threadpool which distributes the work to those independent requests. So we have a simple threadpool that puts very optimized tasks on threads. I should also say that like most "AI" stuff, the processing is memory bound, and not CPU bound. 

I'm wondering if I'd get some efficiency out of golang compared to a fairly optimized service written in C++. I might just protoype the service these days just to have some numbers to compare, but wanted to gut check with people first, especially since I'm new to go. I'm also not really sure how much effort this prototyping would be, given the above parameters (e.g. it needs to talk gRPC, and call some very optimized C++ code).

Hoping for a positive answer given that from the little I've read, it looks like golang was designed for my use case.

Thanks!

P.S. Also, as a bonus, could you recommend a good in-depth book coming from a C/C++ backgroud?
## [4][dSock – Distributed WebSocket broker](https://www.reddit.com/r/golang/comments/gdnzzs/dsock_distributed_websocket_broker/)
- url: https://github.com/Cretezy/dSock
---

## [5][Handling Colossal Conditions in Go](https://www.reddit.com/r/golang/comments/gdva0j/handling_colossal_conditions_in_go/)
- url: https://link.medium.com/VtsANFKEd6
---

## [6][Go modules and .proto](https://www.reddit.com/r/golang/comments/gdv1ah/go_modules_and_proto/)
- url: https://www.reddit.com/r/golang/comments/gdv1ah/go_modules_and_proto/
---
Hi everyone, what is the proper way to generate .go files from third-party library, which doesn't contain any .go source code. For example, i want to generate .proto files from [https://github.com/tensorflow/serving/tree/master/tensorflow\_serving/apis](https://github.com/tensorflow/serving/tree/master/tensorflow_serving/apis)

But as i get it, `go mod vendor` doesn't copy this repository to vendor as there are no .go files, and i only can get it to modules using `go get -d` [`github.com/tensorflow/serving`](https://github.com/tensorflow/serving)

So i don't understand how to generate .proto that are located in modules.. Any suggestions?
## [7][I wrote a compiler and VM language with a similar syntax to Javascript in Go](https://www.reddit.com/r/golang/comments/gdfkuu/i_wrote_a_compiler_and_vm_language_with_a_similar/)
- url: https://github.com/actuallyachraf/monkey-giggle
---

## [8][Why `github.com/pkg/errors` is not a golden standard in "Effective Go" ?](https://www.reddit.com/r/golang/comments/gdq6f4/why_githubcompkgerrors_is_not_a_golden_standard/)
- url: https://www.reddit.com/r/golang/comments/gdq6f4/why_githubcompkgerrors_is_not_a_golden_standard/
---
I am doing a deep dive in to Go, on my journey I am currently exploring error handling in Go. I was particularity interested in stack traces. There is very useful package [`github.com/pkg/errors`](https://github.com/pkg/errors`) which nicely handles it.

Why [`github.com/pkg/errors`](https://github.com/pkg/errors`) is not a recommended package to handle errors in "[Effective Go](https://golang.org/doc/effective_go.html)"  ?

I know what pain it is to debug an error with no stack trace. Shouldn't it be encouraged if not forced to always add stack trace ?
## [9][Olric v0.2.0 is out: Distributed cache and in-memory key/value data store. It can be used both as an embedded Go library and as a language-independent service.](https://www.reddit.com/r/golang/comments/gd7v1s/olric_v020_is_out_distributed_cache_and_inmemory/)
- url: https://github.com/buraksezer/olric#olric-
---

## [10][Micro Server - a distributed systems runtime for the Cloud and beyond](https://www.reddit.com/r/golang/comments/gdpmv7/micro_server_a_distributed_systems_runtime_for/)
- url: https://micro.mu/blog/2020/05/04/introducing-micro-server.html
---

