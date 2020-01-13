# golang
## [1][Golang IO Cookbook (found on HN)](https://www.reddit.com/r/golang/comments/ensegf/golang_io_cookbook_found_on_hn/)
- url: https://github.com/jesseduffield/notes/wiki/Golang-IO-Cookbook
---

## [2][Go-grpc makes use of go-micro plugins to create a better framework for gRPC development.](https://www.reddit.com/r/golang/comments/eo1cxq/gogrpc_makes_use_of_gomicro_plugins_to_create_a/)
- url: https://github.com/vaniila/go-grpc
---

## [3][storage - An application-oriented unified storage layer for Golang](https://www.reddit.com/r/golang/comments/eo12xc/storage_an_applicationoriented_unified_storage/)
- url: https://www.reddit.com/r/golang/comments/eo12xc/storage_an_applicationoriented_unified_storage/
---
Hello everyone, I built a golang lib [storage](https://github.com/Xuanwo/storage) as an application-oriented unified storage layer. 

This lib supports basic operations across all storage services and advanced operations across implemented storage services. All import changes are descibed in [proposals](https://github.com/Xuanwo/storage/tree/master/docs/design). For now, it supports azblob, cos, dropbox, local fs, gcs, kodo, oss, qingstor, s3, uss.

There are many works to do, any comment or advice will be great for me, a star maybe better though (-;
## [4][Vscode can't find package](https://www.reddit.com/r/golang/comments/eo2iti/vscode_cant_find_package/)
- url: https://www.reddit.com/r/golang/comments/eo2iti/vscode_cant_find_package/
---
Mandatory: "I'm going crazy with this"

I'm on go 1.13 and I'm trying to run a simple "hello world" for this UI package 

    https://github.com/gizak/termui

I run the main.go and it says:

    build goChat: cannot load github.com/gizak/termui/v3: cannot find module providing package github.com/gizak/termui/v3

The package is there though. This is my config

    # Go env variables
    export GOPATH="$HOME/golang/packages:$HOME/golang/proj"
    export GOROOT=/usr/local/opt/go/libexec
    export GOPROXY=off

and the files are right there:

    ~/golang/packages/src/github.com/gizak/termui/v3$ ls -l
    total 160
    -rw-r--r--   1 draco  draco   100B 13 Jan 10:12 alignment.go
    -rw-r--r--   1 draco  draco   770B 13 Jan 10:12 backend.go
    -rw-r--r--   1 draco  draco   2.7K 13 Jan 10:12 block.go
    -rw-r--r--   1 draco  draco   1.5K 13 Jan 10:12 buffer.go
    -rw-r--r--   1 draco  draco   753B 13 Jan 10:12 canvas.go
    -rw-r--r--   1 draco  draco   280B 13 Jan 10:12 doc.go
    drwxr-xr-x   3 draco  draco    96B 13 Jan 10:12 drawille/
    -rw-r--r--   1 draco  draco   4.7K 13 Jan 10:12 events.go
    -rw-r--r--   1 draco  draco   216B 13 Jan 10:12 go.mod
    -rw-r--r--   1 draco  draco   641B 13 Jan 10:12 go.sum
    -rw-r--r--   1 draco  draco   3.2K 13 Jan 10:12 grid.go
    -rw-r--r--   1 draco  draco   768B 13 Jan 10:12 render.go
    -rw-r--r--   1 draco  draco   1.3K 13 Jan 10:12 style.go
    -rw-r--r--   1 draco  draco   3.8K 13 Jan 10:12 style_parser.go
    -rw-r--r--   1 draco  draco   1.1K 13 Jan 10:12 symbols.go
    -rw-r--r--   1 draco  draco   549B 13 Jan 10:12 symbols_other.go
    -rw-r--r--   1 draco  draco   522B 13 Jan 10:12 symbols_windows.go
    -rw-r--r--   1 draco  draco   2.7K 13 Jan 10:12 theme.go
    -rw-r--r--   1 draco  draco   4.5K 13 Jan 10:12 utils.go
    drwxr-xr-x  14 draco  draco   448B 13 Jan 10:12 widgets/
    
what to do?

EDIT: This is my go.mod

~/golang/proj/goChat$ cat go.mod
module goChat

go 1.13

**EDIT: SOLVED** with GO111MODULE=off
## [5][Unit test kubernetes client in Go](https://www.reddit.com/r/golang/comments/eo1rjk/unit_test_kubernetes_client_in_go/)
- url: https://gianarb.it/blog/unit-testing-kubernetes-client-in-go
---

## [6][Build a RESTful HTTP API in Golang w/ Mux](https://www.reddit.com/r/golang/comments/eo2e48/build_a_restful_http_api_in_golang_w_mux/)
- url: https://youtu.be/HmiybuiEZI4
---

## [7][Sandy: A tiny sandbox to run untrusted code](https://www.reddit.com/r/golang/comments/enn0k4/sandy_a_tiny_sandbox_to_run_untrusted_code/)
- url: https://github.com/hobochild/sandy
---

## [8][Initializing services in parallel](https://www.reddit.com/r/golang/comments/eo3dq9/initializing_services_in_parallel/)
- url: https://www.reddit.com/r/golang/comments/eo3dq9/initializing_services_in_parallel/
---
Hello all. I have around 10 services being initialized before my echo web server fires up. Some of them are redis, elastic search etc. But two of them are related to each other. Remaining 8 are independent. So I spawned up goroutines for each of them and decreased the boot time of my server. But those remaining 2 are related. When I put them in goroutines with a sync.Waitgroup it throws a panic. When those are running in the main routine everything is fine.

I can't put up links or screenshots because the codebase is proprietary. I regret that.

Anyone came across the same situation? Thank you.

Okat this is my example post. 

    package main
    
    import (
        "fmt"
        "sync"
    )
    
    func hello(wg *sync.WaitGroup) {
        wg.Add(1)
        fmt.Println("Hello world goroutine")
        wg.Done()
    }
    
    func main() {
        var wg sync.WaitGroup
        go hello(&amp;wg)
        wg.Wait()
        fmt.Println("hello main")
    }

How do I make the multiple goroutines run before the main exits. In this example, I need the `hello` function to execute before printing `hello main`
## [9][How to implement &amp; test a server-streaming gRPC API in Go](https://www.reddit.com/r/golang/comments/env6s0/how_to_implement_test_a_serverstreaming_grpc_api/)
- url: https://youtu.be/SBPjEbZcgf8
---

## [10][Hi Gophers, I've just published this package using to make and train multi-layer neural network. Maybe you find it helpful.](https://www.reddit.com/r/golang/comments/ennq4p/hi_gophers_ive_just_published_this_package_using/)
- url: https://github.com/dathoangnd/gonet
---

