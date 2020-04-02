# golang
## [1][The best kafka tool to debug brokers written in Go !](https://www.reddit.com/r/golang/comments/ftgxel/the_best_kafka_tool_to_debug_brokers_written_in_go/)
- url: https://github.com/birdayz/kaf
---

## [2][Announcing Go-TinyTime, Go-TinyDate's Sister Package](https://www.reddit.com/r/golang/comments/ftlmxo/announcing_gotinytime_gotinydates_sister_package/)
- url: https://qvault.io/2020/04/02/announcing-go-tinytime-go-tinydates-sister-package/
---

## [3][Code generation in Go - constructors](https://www.reddit.com/r/golang/comments/fthwok/code_generation_in_go_constructors/)
- url: http://domsu.net/posts/code-generation-go-constructors/
---

## [4][Noob Q: How do I use .proto with go mod? (imports)](https://www.reddit.com/r/golang/comments/fthlcv/noob_q_how_do_i_use_proto_with_go_mod_imports/)
- url: https://www.reddit.com/r/golang/comments/fthlcv/noob_q_how_do_i_use_proto_with_go_mod_imports/
---
Hello all,
I've currently got a folder structure as follows:

cmd/
proto/
- A/
- - A.proto
- B/
- - B.proto

in B.proto, i have want to import a message from A.proto.

so do the following in b.proto:

`import "proto/A/A.proto"; `

However, the generated `b.pb.go` file looks as follows:

```
import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "proto/A"
)

...
```

due to the `_ "proto/A"`, I keep getting the error
```
proto/B/b.pb.go:10:2: cannot find package "." in:
        /path/to/repo/vendor/proto/A
``` 

How would I go about using proto files with go mod?

Any help greatly appreciated. Thanks!
## [5][Running Golang on the browser with WebAssembly and TinyGo](https://www.reddit.com/r/golang/comments/fszeix/running_golang_on_the_browser_with_webassembly/)
- url: https://marianogappa.github.io/software/2020/04/01/webassembly-tinygo-cheesse/
---

## [6][How do you like to host a golang webservice](https://www.reddit.com/r/golang/comments/fteuve/how_do_you_like_to_host_a_golang_webservice/)
- url: https://www.reddit.com/r/golang/comments/fteuve/how_do_you_like_to_host_a_golang_webservice/
---
Just curious about how people do this.

I've set up nginx with my golang app serving local host upstream.  

How do you provision your servers? 

Thanks!
## [7][how do i only select the last given flag our of multiple flags ? is there a way to do a list of flags ?](https://www.reddit.com/r/golang/comments/ftlhwv/how_do_i_only_select_the_last_given_flag_our_of/)
- url: https://www.reddit.com/r/golang/comments/ftlhwv/how_do_i_only_select_the_last_given_flag_our_of/
---
for example i could pass in 
    command subcommand --flag1 flag2 --flag3

but i just want to use the last passed flag.  I was hoping i could do something similar to this in golang.

    parser.add_argument('what', choices=['1', 2', '3', '4'])

i haven't been able to find something similar to this in the flags package that golang has built in.
## [8][How to Unit Test a GORM Application With Sqlmock](https://www.reddit.com/r/golang/comments/ftitio/how_to_unit_test_a_gorm_application_with_sqlmock/)
- url: https://medium.com/better-programming/how-to-unit-test-a-gorm-application-with-sqlmock-97ee73e36526
---

## [9][Need help unterstand example](https://www.reddit.com/r/golang/comments/ftiboz/need_help_unterstand_example/)
- url: https://www.reddit.com/r/golang/comments/ftiboz/need_help_unterstand_example/
---
I'm trying to learn go with the book "Learning Go Porgramming".  
So far I had no difficulties, but I'm having a hard time on this example.

[https://github.com/vladimirvivien/learning-go/blob/master/ch09/sync2.go](https://github.com/vladimirvivien/learning-go/blob/master/ch09/sync2.go)

    package main
    
    import (
    	"sync"
    	"time"
    )
    
    type Service struct {
    	started bool
    	stpCh   chan struct{}
    	mutex   sync.Mutex
    }
    
    func (s *Service) Start() {
    	s.stpCh = make(chan struct{})
    	go func() {
    		s.mutex.Lock()
    		s.started = true
    		s.mutex.Unlock()
    		&lt;-s.stpCh // wait to be closed.
    	}()
    }
    func (s *Service) Stop() {
    	s.mutex.Lock()
    	defer s.mutex.Unlock()
    	if s.started {
    		s.started = false
    		close(s.stpCh)
    	}
    }
    
    func main() {
    	s := &amp;Service{}
    	s.Start()
    	time.Sleep(time.Second) // do some work
    	s.Stop()
    }

Which purpose does the stop channel stpCh serve ?Why is it necessary that the annonymous go routine in Service.Start blocks until the service has been stopped ? First I thought it's some kind of protection, which prevents from calling Start() again, but I tried it, and I can call start multiple times before calling Stop() once.
## [10][Fun project ideas for golang beginner](https://www.reddit.com/r/golang/comments/fthxei/fun_project_ideas_for_golang_beginner/)
- url: https://www.reddit.com/r/golang/comments/fthxei/fun_project_ideas_for_golang_beginner/
---
Hey Everyone,

I just started learning golang and its very fun to learn as well. 

I am looking for some project ideas that I can implement as a beginner. 

Thanks..
