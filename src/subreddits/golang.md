# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][I saw 14000605 go repositories](https://www.reddit.com/r/golang/comments/jomx7q/i_saw_14000605_go_repositories/)
- url: https://i.redd.it/z70kekw0igx51.jpg
---

## [3][Golang 1.15.4 released](https://www.reddit.com/r/golang/comments/joznmi/golang_1154_released/)
- url: https://www.reddit.com/r/golang/comments/joznmi/golang_1154_released/
---
[https://golang.org/doc/devel/release.html#go1.15.minor](https://golang.org/doc/devel/release.html#go1.15.minor)

&amp;#x200B;

[https://github.com/golang/go/issues?q=milestone%3AGo1.15.4+label%3ACherryPickApproved](https://github.com/golang/go/issues?q=milestone%3AGo1.15.4+label%3ACherryPickApproved)
## [4][Idiomatic Go Resources](https://www.reddit.com/r/golang/comments/jp3om9/idiomatic_go_resources/)
- url: https://medium.com/@dgryski/idiomatic-go-resources-966535376dba
---

## [5][Wake on lan in golang](https://www.reddit.com/r/golang/comments/jp3d6u/wake_on_lan_in_golang/)
- url: https://www.reddit.com/r/golang/comments/jp3d6u/wake_on_lan_in_golang/
---
Hi, I have made a wake on LAN CLI tool and wants to share it to get some feedback on it, I havent been coding in Golang for more than a little under half a year so I'm still a noob so tips to how to make it better would be appreciated [https://github.com/xaner4/Gowakeup](https://github.com/xaner4/Gowakeup)

I also made a blog post about how wake on LAN acctually works with some explanation on how to implement it in Golang [https://blog.xaner.dev/post/wake-on-lan/](https://blog.xaner.dev/post/wake-on-lan/)
## [6][How to use authn/authz with gRPC in the best way?](https://www.reddit.com/r/golang/comments/jp3lqv/how_to_use_authnauthz_with_grpc_in_the_best_way/)
- url: https://www.reddit.com/r/golang/comments/jp3lqv/how_to_use_authnauthz_with_grpc_in_the_best_way/
---
We change our REST backend to gRPC. We use Ory Kratos and Keto for authn&amp;authz. They only work with REST. Now we can either use REST for everything related to the frontend plus a proxy REST-&gt;gRPC, or use gRPC-web. Would you do the first option or opt for the second? How would you implement auth with gRPC-web?
## [7][Combining DDD, CQRS, and Clean Architecture by refactoring a Go project](https://www.reddit.com/r/golang/comments/join3w/combining_ddd_cqrs_and_clean_architecture_by/)
- url: https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/
---

## [8][Rust vs Go](https://www.reddit.com/r/golang/comments/jp4kvp/rust_vs_go/)
- url: https://bitfieldconsulting.com/golang/rust-vs-go
---

## [9][Noob help with interface](https://www.reddit.com/r/golang/comments/jp2iqg/noob_help_with_interface/)
- url: https://www.reddit.com/r/golang/comments/jp2iqg/noob_help_with_interface/
---
Hello all,

I thought I understood go interfaces but this weird issue I'm having is making me question my sanity.

I'm planning to have MapHandlerFunc implement the http.handler interface via creating it's single method, ServeHttp. Ultimate goal is to have it MapHandlerFunc modify the urls in the requests passed into it before serving responses. 

Here's my code:



    package main
    
    import (
    	    "net/http"
    )

    func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	 
        type MapHandlerFunc func(ResponseWriter, *Request)
    
	    func (m MapHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request){

		    http.ServeHTTP(w, r)

	    }

	    return nil
    }

~~~

I'm just mocking the structure so far, hence not returning anything, but the compiler is complaining about the line with my method signature:

syntax error: unexpected http, expecting comma or )


I have no idea what the issue is, I would appreciate any help you can give.
## [10][Linking object file, which is written in any langauge with Go](https://www.reddit.com/r/golang/comments/jp1s1r/linking_object_file_which_is_written_in_any/)
- url: https://www.reddit.com/r/golang/comments/jp1s1r/linking_object_file_which_is_written_in_any/
---
Hello

I'd like to call functions written in C/C++

They have some deep dependencies so I want to compile it into an object file,

then call functions in that object file from GO

Can I do that?

How can I link object files when building a Go binary?

Thank you
## [11][What's the User-Agent for Golang HTTP Client?](https://www.reddit.com/r/golang/comments/jp1r0w/whats_the_useragent_for_golang_http_client/)
- url: https://www.reddit.com/r/golang/comments/jp1r0w/whats_the_useragent_for_golang_http_client/
---
I am using a custom round tripper for Reddit OAuth2. It requires strict User-Agent.

&amp;#x200B;

I don't know what will be the User-Agent header for Golang net/http client
