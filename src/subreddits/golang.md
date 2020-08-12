# golang
## [1][Golang 1.15 released!](https://www.reddit.com/r/golang/comments/i80s89/golang_115_released/)
- url: https://golang.org/doc/go1.15
---

## [2][Is there a tool that warns about variable shadowing?](https://www.reddit.com/r/golang/comments/i8ch3t/is_there_a_tool_that_warns_about_variable/)
- url: https://www.reddit.com/r/golang/comments/i8ch3t/is_there_a_tool_that_warns_about_variable/
---
Just spent some time trying to figure out some very unexpected behavior. A simplified version is available on playground [https://play.golang.org/p/SxK\_XO-NvM1](https://play.golang.org/p/SxK_XO-NvM1)

    package main
    
    import (
    	"fmt"
    	"strconv"
    )
    
    func add(w int) int {
    	return w + 1
    }
    
    func doStuff(w int) {
    	fmt.Println(strconv.Itoa(w))
    }
    
    func main() {
    	i := 5
    	if i == 3 {
    		doStuff(i)
    	}
    
    	for z := 0; z &lt; 4; z++ {
    		i := add(i)
    		doStuff(i)
    	}
    
    	doStuff(i)
    }

The error is an extra `:` present inside the forloop which creates a variable with the same name.

&amp;#x200B;

Is there any tool that warns about this?

&amp;#x200B;

We currently use golang.org/x/lint/golint, honnef.co/go/tools/cmd/staticcheck and go vet but none of them complains. Or at least none of them complained about the more complex case.
## [3][How Go 1.15 improved converting small integer values to interfaces](https://www.reddit.com/r/golang/comments/i8bqyj/how_go_115_improved_converting_small_integer/)
- url: https://utcc.utoronto.ca/~cks/space/blog/programming/Go115InterfaceSmallInts
---

## [4][ote updates a packages' go.mod file with a comment next to all dependencies that are test dependencies; identifying them as such.](https://www.reddit.com/r/golang/comments/i8c042/ote_updates_a_packages_gomod_file_with_a_comment/)
- url: https://github.com/komuw/ote
---

## [5][oscar: Next generation building tool for nothing 🐶 .](https://www.reddit.com/r/golang/comments/i81vx3/oscar_next_generation_building_tool_for_nothing/)
- url: https://github.com/chenjiandongx/oscar
---

## [6][Wrapping up our series on value semantics vs pointer semantics!](https://www.reddit.com/r/golang/comments/i8drnb/wrapping_up_our_series_on_value_semantics_vs/)
- url: https://www.youtube.com/watch?v=579aeQSJFBs&amp;t=12s
---

## [7][I want to learn go and grpc by making microservices but i am getting clueless whenever i have to start. I want communication between two microservices but i don't know how to make can anyone give me some help or suggestions so that i can get started with it ?](https://www.reddit.com/r/golang/comments/i8cq19/i_want_to_learn_go_and_grpc_by_making/)
- url: https://www.reddit.com/r/golang/comments/i8cq19/i_want_to_learn_go_and_grpc_by_making/
---

## [8][Production ready Golang microservice?](https://www.reddit.com/r/golang/comments/i8cmfx/production_ready_golang_microservice/)
- url: https://www.reddit.com/r/golang/comments/i8cmfx/production_ready_golang_microservice/
---
Where can i find an example of production ready Golang microservice?

With proper auth, comunication protocols, good code structure, good practices etc
## [9][Email Disrupt](https://www.reddit.com/r/golang/comments/i8benr/email_disrupt/)
- url: https://github.com/monopolly/mail
---

## [10][Centrifuge – a real-time messaging core of Centrifugo server as standalone library for Go](https://www.reddit.com/r/golang/comments/i7xec5/centrifuge_a_realtime_messaging_core_of/)
- url: https://github.com/centrifugal/centrifuge
---

