# golang
## [1][I made something cool in go, wanted to share it here](https://www.reddit.com/r/golang/comments/epxt1j/i_made_something_cool_in_go_wanted_to_share_it/)
- url: https://www.reddit.com/r/golang/comments/epxt1j/i_made_something_cool_in_go_wanted_to_share_it/
---
I made neko, a neat docker app that can stream a browser to multiple clients, you can take control of the keyboard and mouse and browse the web with friends. The server uses  [pion](https://github.com/pion/webrtc) to stream the desktop through webrtc

https://i.redd.it/4qs4uck1wab41.gif

[https://github.com/nurdism/neko](https://github.com/nurdism/neko)
## [2][go has gone mad and is eating all my disk space!](https://www.reddit.com/r/golang/comments/ept577/go_has_gone_mad_and_is_eating_all_my_disk_space/)
- url: https://imgur.com/a/cghBXUe
---

## [3][Need help getting started with Go.](https://www.reddit.com/r/golang/comments/epzmgh/need_help_getting_started_with_go/)
- url: https://www.reddit.com/r/golang/comments/epzmgh/need_help_getting_started_with_go/
---
Hello guys! I want to build a backend to my application. I will need a database of course, and a Rest API. I saw there are a lot of options and got kinda lost. I'm really just starting with Go and need something simple and straight forward to get started. Any tips on what to use?
## [4][Best practices on error handling in Go](https://www.reddit.com/r/golang/comments/epq5wh/best_practices_on_error_handling_in_go/)
- url: https://www.reddit.com/r/golang/comments/epq5wh/best_practices_on_error_handling_in_go/
---
Hello, since Go 1.13, it seems that the error handling already changes

[https://blog.golang.org/go1.13-errors](https://blog.golang.org/go1.13-errors)

  
I'm new to Go, but is there any resources on how to create/organize error handling in Go based on Go 1.13. 

In addition to that, I have a question. How do you decide when to use sentinel errors vs type struct error?
## [5][Stream Is Deprecating Virtual Go – It’s Time to Move to Go Modules](https://www.reddit.com/r/golang/comments/epk671/stream_is_deprecating_virtual_go_its_time_to_move/)
- url: https://getstream.io/blog/virtual-go-to-go-modules/
---

## [6][To Vendor or not To Vendor with Go Modules?](https://www.reddit.com/r/golang/comments/epzqi8/to_vendor_or_not_to_vendor_with_go_modules/)
- url: https://www.reddit.com/r/golang/comments/epzqi8/to_vendor_or_not_to_vendor_with_go_modules/
---
I created a PR over here. [https://github.com/theupdateframework/notary/pull/1523](https://github.com/theupdateframework/notary/pull/1523)

Now the question is if we should vendor the Go modules or we shouldn't. For both the arguments there is some point.

What does the Go community think about vendoring Go Modules?

What are the tradeoffs?

What are your experiences?
## [7][Reflecting on My Experience With Go, One Year After - Random Bits of Wisdom](https://www.reddit.com/r/golang/comments/epzgfc/reflecting_on_my_experience_with_go_one_year/)
- url: https://p5v.me/2020/01/reflecting-on-my-experience-with-go-one-year-after/
---

## [8][Go Mux - How to create multiple binaries that can communicate with each other?](https://www.reddit.com/r/golang/comments/epxkzr/go_mux_how_to_create_multiple_binaries_that_can/)
- url: https://www.reddit.com/r/golang/comments/epxkzr/go_mux_how_to_create_multiple_binaries_that_can/
---
While using Go Mux (which is a wonderful library), I have defined handerfunctions in different files.

Lets take an example, I have main.go as below:

`package main`

`import (`

    `"../lib/mux"`
    
    `"../handlers"`
    
    `"net/http"`
    
    `"fmt"`

`)`

`func main() {`

    `fmt.Println("hello")`
    
    `r := mux.NewRouter()`
    
    `r.HandleFunc("/{entity}", handlers.HomeHandler).Methods("GET")`
    
    `http.ListenAndServe(":8000", r)`

`}`

and HomeHandler.go

`package handlers

import (`

    "fmt"`
    
    "../lib/mux"`
    
    "net/http"`

)`

`func HomeHandler(w http.ResponseWriter, rr *http.Request) {`

    `//Some code here`
    
    `params := mux.Vars(rr)`
    
    `entityName := params["entity"]`
    
    `w.Header().Set("Content-Type", "text/html")`	
    
    `fmt.Fprintf(w, entityName)`

`}`

When I run **go build main.go**, it created a single executable binary. Is it possible to create individual

binaries based on number of handler functions we have? If possible, can they communicate with each other and also with main binary (main.exe)?

&amp;#x200B;

If in future I create 5000+ routes, I will have to maintain 5000+ handler functions which will get compiled into a single binary which is something I don't want. What is the right approach to following modular design?

I have read about go plugins, they can do exactly what I want however go developers and community has not updated any new feature in that. People also say it is worst with few drawbacks such as plugins work only when they are built on same version on which main program was built.

Another idea is to make use of gRFC but 5000+ gRFCs will consume and block 5000+ port numbers. What is the right approach to follow?
## [9][Screencast of Go binary running on an RG350 handheld (mipsle)](https://www.reddit.com/r/golang/comments/epxfjg/screencast_of_go_binary_running_on_an_rg350/)
- url: https://youtu.be/ngmSXZiqLuQ
---

## [10][online book - Build web application with Golang](https://www.reddit.com/r/golang/comments/epyjtp/online_book_build_web_application_with_golang/)
- url: https://astaxie.gitbooks.io/build-web-application-with-golang/en/
---

