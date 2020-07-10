# golang
## [1]["Interface smuggling", a Go design pattern for expanding APIs](https://www.reddit.com/r/golang/comments/hoehhv/interface_smuggling_a_go_design_pattern_for/)
- url: https://utcc.utoronto.ca/~cks/space/blog/programming/GoInterfaceSmuggling
---

## [2][Adaptive routing for load balancing](https://www.reddit.com/r/golang/comments/homsvy/adaptive_routing_for_load_balancing/)
- url: https://www.reddit.com/r/golang/comments/homsvy/adaptive_routing_for_load_balancing/
---
Couple years back, I worked on serviceq ([https://github.com/gptankit/serviceq](https://github.com/gptankit/serviceq)) which is a load balancer written in Go, that aims to dynamically adapt to errors in downstream nodes. It does this by   
a) adaptive routing  
b) queueing and retrying failed requests

Have exposed the adaptive routing logic as a package - harmonic ([https://github.com/gptankit/harmonic](https://github.com/gptankit/harmonic)) - which can be used from client apps, load balancers or test suites, if needed.

Comments/suggestions are welcome.
## [3][A useful tip for fetching lockless parallel fetching](https://www.reddit.com/r/golang/comments/hoodnu/a_useful_tip_for_fetching_lockless_parallel/)
- url: https://www.reddit.com/r/golang/comments/hoodnu/a_useful_tip_for_fetching_lockless_parallel/
---
I often use the folloing construction (for instance when I have to fetch a number of things:

    var wg sync.WaitGroup
    results := make([]MyUsefulData, n)
    for i := 0; i &lt; n; i++ {
      // We should use current value of i
      sendToAnyGorutinePoolRealization(i, fund(i int) {
        results[i] = fetchMyUsefulData(i)
        wg.Done()
      })
    }
    wg.Wait()

This code is race-safe and has no locks. #tips_and_tricks
## [4][Draft GORM V2 Documents Online, suggestions, feedback welcome!](https://www.reddit.com/r/golang/comments/ho0ajo/draft_gorm_v2_documents_online_suggestions/)
- url: http://v2.gorm.io/docs/
---

## [5][Does generating a new instance from a struct increase memory consumption?](https://www.reddit.com/r/golang/comments/hokebm/does_generating_a_new_instance_from_a_struct/)
- url: https://www.reddit.com/r/golang/comments/hokebm/does_generating_a_new_instance_from_a_struct/
---
Does generating a new instance from a struct increase memory consumption? Or does it have memory cost?

I need to generate 50 new instances per second, each instance passes through a pipeline and after some modification I leave that instance alone (each instance lives maximum 1min - 2min)

But I'm very worried that this will increase memory consumption (creating 50 instance per second and store them in a slice while they are passing through the pipeline)!
## [6][Go bindings to QuickJS: a fast, small, and embeddable ES2020 JavaScript interpreter.](https://www.reddit.com/r/golang/comments/ho0o9e/go_bindings_to_quickjs_a_fast_small_and/)
- url: https://github.com/lithdew/quickjs
---

## [7][Go-streams v0.5.0 released](https://www.reddit.com/r/golang/comments/ho4dbd/gostreams_v050_released/)
- url: https://github.com/reugn/go-streams
---

## [8][Help deploying Golang app](https://www.reddit.com/r/golang/comments/hofsbs/help_deploying_golang_app/)
- url: https://www.reddit.com/r/golang/comments/hofsbs/help_deploying_golang_app/
---
hey all, I'm a fairly new to Golang and have just created a simple application that will send me an email everyday at a certain time.  Obviously I can't just run the program on my local computer forever, so where is a free service that I can can use to deploy my application? Its not a web app so its not accepting and responding to requests, but just a couple function that are being called in \`main.go\` every day at 8:00 AM.  Any help would be greatly appreciated!
## [9][Gitpod review](https://www.reddit.com/r/golang/comments/hoa1fs/gitpod_review/)
- url: https://www.reddit.com/r/golang/comments/hoa1fs/gitpod_review/
---
Hey there! I use JetBrains' GoLand application as an IDE, however I sometimes need to use other pc than mine. Do you think changing to gitpod would be worth it?

Thanks fou your answers!
## [10][Chat Application (Need some eyes on code)](https://www.reddit.com/r/golang/comments/ho3od7/chat_application_need_some_eyes_on_code/)
- url: https://www.reddit.com/r/golang/comments/ho3od7/chat_application_need_some_eyes_on_code/
---
Hi everyone,  
A while back I thought of writing a chat application with private and group chats, with the backend in Golang.  
My sole intention of picking this up is to learn/work with channels and web-sockets from scratch in Golang, so I'm trying not to make use of services like Pusher or Redis.  


I started by creating a public group where everyone can join and everyone will receive the messages. Building on that, I have added groups to the application so that users connected to one group don't get the messages of another group i.e. the basic functionality.  


What I now need is some review on the code that I've written so far, any mistakes I might have overlooked, "not so good ways to write code", suggestions, stuff like that.  
It's not a lot of code, so I haven't split it into different packages yet. But that will definitely be my next commit.

[Github](https://github.com/Tak1za/ivar)  


Thanks :)
