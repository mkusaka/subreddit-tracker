# golang
## [1][Micro In Action, part3: Calling a Service](https://www.reddit.com/r/golang/comments/f26gdr/micro_in_action_part3_calling_a_service/)
- url: https://medium.com/@dche423/micro-in-action-part-3-calling-a-service-55d865928f11
---

## [2][CurlyQ - Simple, efficient, and reliable task queue for Go](https://www.reddit.com/r/golang/comments/f1wy28/curlyq_simple_efficient_and_reliable_task_queue/)
- url: https://www.reddit.com/r/golang/comments/f1wy28/curlyq_simple_efficient_and_reliable_task_queue/
---
[CurlyQ](https://github.com/mcmathja/curlyq) provides a simple, easy-to-use interface for performing background processing in Go. It supports scheduled jobs, job deduplication, and configurable concurrent execution out of the box.

I started writing CurlyQ after experimenting with other Redis-backed task queues and finding that none of them had quite what I was looking for in terms of API design and feature mix. CurlyQ aims to be as straightforward as possible for common use cases while exposing a variety of options for fine-tuning and advanced configuration.

I'm nearing a 1.0 release and would love feedback and contributions before locking in the current behavior. Thanks all!
## [3][Go modules proxy](https://www.reddit.com/r/golang/comments/f243lc/go_modules_proxy/)
- url: https://www.reddit.com/r/golang/comments/f243lc/go_modules_proxy/
---
Hey , 

I want to host my own proxy server for go modules. 

I've seen several projects that acts as a proxy/cache server and have some questions-

1. When using GOPROXY environment variable to set a proxy , and the module is not in the proxy , will it go get it from a VCS? 

I've tried to set it up, but it's seems to not work like that and just got an error. There is any extra setup need to be done? 

2. How can I "fill" my proxy cache with popular modules so I'll have it "ready" . 

3. How does the proxy differs from the cache found locally on GOPATH/pkg/mod ? 

Except the fact it can be accessed by multiple people on the network

Thanks!
## [4][Agente - Distributed simple and robust release management and monitoring system.](https://www.reddit.com/r/golang/comments/f2639c/agente_distributed_simple_and_robust_release/)
- url: https://www.reddit.com/r/golang/comments/f2639c/agente_distributed_simple_and_robust_release/
---
Distributed simple and robust release management and monitoring system.

\*\**This project on going work.*

### 

[Agente](https://preview.redd.it/1mzjievfl9g41.png?width=512&amp;format=png&amp;auto=webp&amp;s=5ee0e2e48cca5b2e05d81d2aca83c9ba748248e3)

### Road map

* Core system
* First worker agent
* Management dashboard
* Jenkins vs CI tool extensions
* Management dashboard
* First master agent
* All relevant third-party system integrations (version control, CI, database, queuing etc.)

## Requirements

* Go &gt; 1.11
* Redis or RabbitMQ
* PostgreSQL

## Docker Environment

For PostgreSQL

`docker run --name agente_PostgreSQL -e POSTGRES_PASSWORD=123456 -e POSTGRES_USER=agente -p 5432:5432 -d postgres  docker exec agente_PostgreSQL psql --username=agente -c 'create database agente_dev;'`

For RabbitMQ

`docker run --hostname my-rabbit --name agente_RabbitMQ -e RABBITMQ_DEFAULT_USER=local -e RABBITMQ_DEFAULT_PASS=local -p 5672:5672 -d rabbitmq:3-management`

Development

`git clone -b develop` [`https://github.com/streetbyters/agente`](https://github.com/streetbyters/agente)

`go mod vendor`  

`# Development Mode`

 `go run ./cmd -mode dev -migrate -reset` 

`go run ./cmd -mode dev`  

`# Test Mode` 

`go run ./cmd -mode test -migrate -reset` 

`go run ./cmd -mode test`

## Build

We will release firstly Agente for Linux environment.  
 [See detail](https://github.com/streetbyters/agente/blob/master/docs/build.md)

&amp;#x200B;

GitHub : [https://github.com/streetbyters/agente](https://github.com/streetbyters/agente)
## [5][Can $GOROOT be a totally different directory than $GOPATH?](https://www.reddit.com/r/golang/comments/f26of5/can_goroot_be_a_totally_different_directory_than/)
- url: https://www.reddit.com/r/golang/comments/f26of5/can_goroot_be_a_totally_different_directory_than/
---
Hello there,

In my first baby steps with golang I kept the any new project in the /src directory of my $GOPATH. It was all set inside the directory of the go source code, as i had first downloaded it. Just like most blogposts recommend.

No  I am setting up a new workstation and I thught that maybe I could use different directories for my $GOROOT to keep the source code of the language and another $GOPATH for my projects. I thought that maybe this would help me keep it all clean and nitty, and maybe ease possible back ups in the future ( besides having it all on a git repo). I also added the $GOBIN environmental variable, but I came up with permission errors.

Anyway, do you think that it even worth the mess? Is what I'm doing a good practice or is it totally misguiding? Should I just follow the default practice?
## [6][Just released a library - goptr](https://www.reddit.com/r/golang/comments/f28726/just_released_a_library_goptr/)
- url: https://www.reddit.com/r/golang/comments/f28726/just_released_a_library_goptr/
---
I have a project with some boilerplates related to handling pointers so I've decided to make them into a library named [goptr](https://github.com/noamt/goptr) with the purpose of becoming a utility library.

The first feature is a collection of methods for resolving pointers of primitives or falling back to a default value. For example:

    import "github.com/noamt/goptr"
    
    goptr.String(nil) // Returns an empty string
    goptr.String("foo") // Returns "foo" 
    goptr.StringOrDefault(nil, "foo") // Returns "foo" 
    
    s := "bar"
    goptr.StringOrDefault(&amp;s, "foo") // Returns "bar"
## [7][Micro In Action, Part 3：Calling a Service](https://www.reddit.com/r/golang/comments/f27zd9/micro_in_action_part_3calling_a_service/)
- url: https://medium.com//micro-in-action-part-3-calling-a-service-55d865928f11?source=friends_link&amp;sk=b763ed6cc10dd14132da4191387fcaa4
---

## [8][Build a simple GraphQL API in Golang with MySQL and GORM using Gqlgen](https://www.reddit.com/r/golang/comments/f20nfw/build_a_simple_graphql_api_in_golang_with_mysql/)
- url: https://www.soberkoder.com/go-graphql-api-mysql-gorm/
---

## [9][Show which docker containers are attached to which veth interfaces CLI](https://www.reddit.com/r/golang/comments/f27gqn/show_which_docker_containers_are_attached_to/)
- url: https://www.reddit.com/r/golang/comments/f27gqn/show_which_docker_containers_are_attached_to/
---
Hi everyone :)

I developed a CLI tool that can show which docker containers are attached to which veth interfaces.

It can be used in combination with other commands.  
e.g. Capture any container's packets with termshark.

```sh
sudo termshark -i $(dockerveth -p | fzf | awk "{print \$2}")
```

[https://github.com/skanehira/go-dockerveth](https://github.com/skanehira/go-dockerveth)
## [10][Having trouble understanding the concept of a "package" in Go](https://www.reddit.com/r/golang/comments/f1xwrs/having_trouble_understanding_the_concept_of_a/)
- url: https://www.reddit.com/r/golang/comments/f1xwrs/having_trouble_understanding_the_concept_of_a/
---
I come from the world of Ruby and Rails and I'm trying to build some things in Go. 

It seems like packages are a way of grouping functional code in Go? But I'm trying to wrap my head around 

- how and when to include files in a package? Are all files in a package accessible to each other? 
- Is my package always called `main`? Is that just a convention?
- When to break code into multiple packages? If I'm building a CLI tool, would each main subcommand in the CLI get it's own package and subfolder grouping? 

I've taken a look at the usual getting started docs but I just feel like I'm stumbling here for some reason. Hoping that someone can explain it better.

Thanks!
