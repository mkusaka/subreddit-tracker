# golang
## [1][Will go ever be good for GUI?](https://www.reddit.com/r/golang/comments/eln3ub/will_go_ever_be_good_for_gui/)
- url: https://www.reddit.com/r/golang/comments/eln3ub/will_go_ever_be_good_for_gui/
---
So far implemented GUI libraries have generally just relied on Qt bindings or browser renderers. When daydreaming, I imagine a Go standard library which includes a convenient, open-source Go GUI lib.
## [2][SimpleS3: AWS S3 Golang Library using REST with V4 Signing (without AWS Go SDK)](https://www.reddit.com/r/golang/comments/elh0cg/simples3_aws_s3_golang_library_using_rest_with_v4/)
- url: https://github.com/rhnvrm/simples3
---

## [3][Go Modules can now be scanned for vulnerabilities in GoCenter!](https://www.reddit.com/r/golang/comments/elh71w/go_modules_can_now_be_scanned_for_vulnerabilities/)
- url: https://www.reddit.com/r/golang/comments/elh71w/go_modules_can_now_be_scanned_for_vulnerabilities/
---
&amp;#x200B;

https://preview.redd.it/6ca59ak92f941.jpg?width=800&amp;format=pjpg&amp;auto=webp&amp;s=98a4cbbea5986a561f1dca2ad6e76b57e97ec665

A really cool security feature just got added to [GoCenter.io](https://GoCenter.io); JFrog Xray’s vulnerability scanning of Go modules! GoCenter automatically scans for known vulnerabilities recognized in the public vulnerability database NVD. Those results are stored in GoCenter and exposed on the Security page of the UI, which will list all vulnerabilities that exist in the module version. Learn more: [https://jfrog.com/blog/gocenter-reveals-go-module-vulnerabilities-with-xray/](https://jfrog.com/blog/gocenter-reveals-go-module-vulnerabilities-with-xray/)

Disclaimer: I am an employee of JFrog, but think this is something that's pretty important for Go!
## [4][Microservice that extends from messages-api is responsible for sending IoT messages through some providers such as NATS or MQTT](https://www.reddit.com/r/golang/comments/elpirn/microservice_that_extends_from_messagesapi_is/)
- url: https://github.com/microapis/messages-iot-api
---

## [5][Error trying to build an go-sdl2 project](https://www.reddit.com/r/golang/comments/elsipt/error_trying_to_build_an_gosdl2_project/)
- url: https://www.reddit.com/r/golang/comments/elsipt/error_trying_to_build_an_gosdl2_project/
---
I recently started programming in Go and I'm looking to set up go-sdl2 since I'm familiar with the library, and have downloaded the dependencies with  
  
    go get -v github.com/veandco/go-sdl2
  
I ran `make.bat` after I added the dependencies from `SDL2-devel-2.0.10-mingw\include\` into the go-sdl2 folder, which built without any errors. Then I have the following code in my `main.go` file:  
  
`package main

    /*
	#include &lt;stdio.h&gt;
	#include &lt;errno.h&gt;
    */
    import "C"
    
    import (
	"fmt"
	"unsafe"
	"github.com/veandco/go-sdl2/sdl"
    )

    var w = *sdl.Wiindow
    var r = *sdl.Renderer

    func main() {
        if w, r, err := SDL_Init(SDL_INIT_EVERYTHING) != nil {
                panic("Failed to init SDL2")
        }
        w, r, err := SDL_CreateWindowAndRenderer(
		800, 600, 
		sdl.WINDOW_SHOWN
	)
	if err != nil {
		return nil, err
	}
    }
  
This gives an error on build:  
  
    could not determine kind of name for C.SDL_AudioCVT
    could not determine kind of name for C.SDL_AudioCallback
    could not determine kind of name for C.SDL_AudioDeviceID
    could not determine kind of name for C.SDL_AudioFilter
    could not determine kind of name for C.SDL_AudioFormat
    could not determine kind of name for C.SDL_AudioSpec
    could not determine kind of name for C.SDL_AudioStatus
    could not determine kind of name for C.SDL_AudioStream
    could not determine kind of name for C.Uint32
    could not determine kind of name for C.Uint8
    could not determine kind of name for C.int
    could not determine kind of name for C.size_t  
  
I'm new to Go, and have tried googleling for help, but couldn't find anything to solve this problem. Also any tutorials on how to use SDL2 with go is greatly appreciated!
## [6][Question: Server.Shutdown -- any way to force an early shutdown?](https://www.reddit.com/r/golang/comments/els1l2/question_servershutdown_any_way_to_force_an_early/)
- url: https://www.reddit.com/r/golang/comments/els1l2/question_servershutdown_any_way_to_force_an_early/
---
Hi,

Does anyone here know much about [https://golang.org/pkg/net/http/#Server.Shutdown](https://golang.org/pkg/net/http/#Server.Shutdown) ? 

I've been banging my head against the wall for a while trying to figure out why an in-flight request was still completing even though I was setting a short context timeout.

I *thought* the behaviour would have dropped any in-flight requests once the context timeout had expired, but I now think that the context timeout is there only for notification purposes, rather than being able to cause the side-effect of stopping the server early.

I now believe the docs suggest golang wont stop the server until the in-flight request has gone 'idle'. Meaning: I can't *force* it to shutdown early, but only be notified that my 'intended' timeout has been passed, and the server is still going to wait for a graceful shutdown.

Just trying to bounce this thought process off some other people who might know about this sort of thing 😬
## [7][How to write API gateway for microservices?](https://www.reddit.com/r/golang/comments/elnw9d/how_to_write_api_gateway_for_microservices/)
- url: https://www.reddit.com/r/golang/comments/elnw9d/how_to_write_api_gateway_for_microservices/
---
I have created microservices. Now I want to add an API gateway which will redirect http request to specific microservice. Should I need any third party library like GoKit? If not then how can I implement it using net/http package?
## [8][autocache: groupcache using memberlist for peer enrollment](https://www.reddit.com/r/golang/comments/elfpg3/autocache_groupcache_using_memberlist_for_peer/)
- url: https://github.com/pomerium/autocache
---

## [9][GoLand 2019.3.1 formatting changes - which setting to fix?](https://www.reddit.com/r/golang/comments/elk9pp/goland_201931_formatting_changes_which_setting_to/)
- url: https://www.reddit.com/r/golang/comments/elk9pp/goland_201931_formatting_changes_which_setting_to/
---
Upgraded GoLand from 2019.2 to 2019.3.1 today and now it corrupts the display of a fmt.Sprintf()

The real code is `fmt.Sprintf("hello %s", name)` but it instead now displays as `fmt.Sprintf("hello #{name}")`

Does anyone know the Setting which turns this corruption off?


Update - Solution:

To remove this; 

1. Goto Settings-&gt;Editor-&gt;General-&gt;Code Folding
2. Scroll down to the options for `Go`, and deselect `Format Strings`
## [10][Best way to read chunks of S3 file into memory](https://www.reddit.com/r/golang/comments/elp5q6/best_way_to_read_chunks_of_s3_file_into_memory/)
- url: https://www.reddit.com/r/golang/comments/elp5q6/best_way_to_read_chunks_of_s3_file_into_memory/
---
I'm having a little trouble finding\\understanding the best method to read an S3 file line by line (or in chunks) without reading the entire file contents into memory. The below code works but I'm not sure it's the most effective. Ideally, I would like to create a data stream. Any help or suggestions would be greatly appreciated. 

`req, err := svc.GetObject(&amp;s3.GetObjectInput{`  
`Bucket: aws.String("my.s3.bucket"),`  
`Key:    aws.String("TEST/test.log"),`  
`})`  
`if err != nil {`  
 `panic(err)`  
`}`  
`var reader io.Reader`  
`reader = req.Body`  
`data, _ := ioutil.ReadAll(reader)`  
`var stringData string`  
`stringData = string(data[:])`  
`scanner := bufio.NewScanner(strings.NewReader(stringData))`  
`for scanner.Scan() {`  
    `fmt.Println(scanner.Text())`  
`}`
