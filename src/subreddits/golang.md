# golang
## [1][Diago, a visualization tool for profiles and heap snapshots generated with pprof](https://www.reddit.com/r/golang/comments/fl8gbv/diago_a_visualization_tool_for_profiles_and_heap/)
- url: https://remy.io/blog/how-to-use-diago-to-diagnose-cpu-and-memory-usage-in-go-programs/
---

## [2][I wrote a guide for beginners wanting to get started with Go](https://www.reddit.com/r/golang/comments/fkw2rx/i_wrote_a_guide_for_beginners_wanting_to_get/)
- url: https://www.reddit.com/r/golang/comments/fkw2rx/i_wrote_a_guide_for_beginners_wanting_to_get/
---
My recent experience getting started with another language made me realized that the first 1-2 weeks of learning a new language is not easy.

As soon as you lose the hand-holding of the classic resources out there, you know that day when you wake up in the morning and say: alright, today I start my project. It's usually when people get discouraged because of road bumps.

When it's the first time you're going to write in a new language "on your own", it's challenging. 

I wrote a guide for that exact moment. What they should read first, how to structure their first projects, what not to be afraid of that might be common to some programmers, listing some useful packages that they will probably have to use on day one, having them used to write tests, a general development flow, etc.

There's not much code example, but it's more like "last time" hand-holding, or recomforting guide that tries to address what I think can block new programmers from writing their first program in Go.

This is the URL: https://dominicstpierre.com/getting-started-with-go-guide

I hope it can help some programmers wanting to jump in Go pass the Tour of Go and initial tutorial(s) that tells them to write X, Y, Z.
## [3][Implementations of reinforcement learning processes in Go.](https://www.reddit.com/r/golang/comments/fl5pq1/implementations_of_reinforcement_learning/)
- url: https://github.com/eltorocorp/reinforcement-learning
---

## [4][Golang Mutexes - What Is RWMutex For?](https://www.reddit.com/r/golang/comments/fl9cgc/golang_mutexes_what_is_rwmutex_for/)
- url: https://qvault.io/2020/03/19/golang-mutexes-what-is-rwmutex-for/
---

## [5][Untyped constants in Go - 2mins intro with quiz inside](https://www.reddit.com/r/golang/comments/fl7ouv/untyped_constants_in_go_2mins_intro_with_quiz/)
- url: https://medium.com/golangspec/untyped-constants-in-go-2c69eb519b5b
---

## [6][Help with error goroutine [running] sync.(*Mutex).Lock(...)](https://www.reddit.com/r/golang/comments/fl9ww9/help_with_error_goroutine_running_syncmutexlock/)
- url: https://www.reddit.com/r/golang/comments/fl9ww9/help_with_error_goroutine_running_syncmutexlock/
---
I was testing deadlocks in Go and now whenever I execute code it gives me an error that I cannot fix :(. There is no error when I want to print a simple message, it's just when I have database query. I tried restarting computer but that did not help, even when I removed all code related to Mutex/Deadlocks.

&amp;#x200B;

    panic: runtime error: invalid memory address or nil pointer dereference
    [signal 0xc0000005 code=0x1 addr=0x20 pc=0x4b60f8]
    
    goroutine 1 [running]:
    sync.(*Mutex).Lock(...)
            D:/Software/Go/src/sync/mutex.go:74
    database/sql.(*DB).conn(0x0, 0x51e5e0, 0xc0000120a0, 0x1c0001, 0x0, 0x46dad9, 0x0)
            D:/Software/Go/src/database/sql/sql.go:1132 +0x48
    database/sql.(*DB).query(0x0, 0x51e5e0, 0xc0000120a0, 0xc00000a300, 0x2c, 0x0, 0x0, 0x0, 0x4fe201, 0xa, ...)
            D:/Software/Go/src/database/sql/sql.go:1565 +0x6d
    database/sql.(*DB).QueryContext(0x0, 0x51e5e0, 0xc0000120a0, 0xc00000a300, 0x2c, 0x0, 0x0, 0x0, 0x0, 0xc00008bee0, ...)
            D:/Software/Go/src/database/sql/sql.go:1547 +0xd8
    database/sql.(*DB).QueryRowContext(...)
            D:/Software/Go/src/database/sql/sql.go:1648
    database/sql.(*DB).QueryRow(0x0, 0xc00000a300, 0x2c, 0x0, 0x0, 0x0, 0x2c)
            D:/Software/Go/src/database/sql/sql.go:1659 +0x92
    main.myTest(0x4fe2c4, 0xa)
            D:/x/api.go:126 +0xe3
    main.main()
            D:/x/api.go:41 +0x3d
## [7][Convert Apple CoreML to any format which I can use with Go.](https://www.reddit.com/r/golang/comments/fl9qyh/convert_apple_coreml_to_any_format_which_i_can/)
- url: https://www.reddit.com/r/golang/comments/fl9qyh/convert_apple_coreml_to_any_format_which_i_can/
---
Is there a way to convert (or use) CoreML format to anything like Haar or Tensorflow? What is the best way to do it and which direction shall I go? I really like [https://gocv.io/](https://gocv.io/) and I would like to use it if possible. For now I'm trying to use CoreML as I'm having a problem to train my own model by yolov3.

Thanks!
## [8][Linux Window Managers in Go?](https://www.reddit.com/r/golang/comments/fl9lxg/linux_window_managers_in_go/)
- url: https://www.reddit.com/r/golang/comments/fl9lxg/linux_window_managers_in_go/
---
Has anyone come across any Linux Window managers written in Go?

A tiling window manger written in Go would be cool. Like i3 etc..
## [9][A GoMock Quick Start Guide](https://www.reddit.com/r/golang/comments/fl3gqo/a_gomock_quick_start_guide/)
- url: https://medium.com/better-programming/a-gomock-quick-start-guide-71bee4b3a6f1
---

## [10][Handy wrapper for MongoDB driver (repositories, query builder, aggregate pipeline builder)](https://www.reddit.com/r/golang/comments/fl49t0/handy_wrapper_for_mongodb_driver_repositories/)
- url: https://www.reddit.com/r/golang/comments/fl49t0/handy_wrapper_for_mongodb_driver_repositories/
---
I wrote this package which we use internally at our company for creating a more organized way to manage a MongoDB client and collection specific methods (called repositories). We've also written some handy functions for creating queries and aggregate pipelines without dealing with the bson types directly which sometimes can be a little tricky. Anyways, looking for feedback on the project. Thanks!

[https://github.com/clarkmcc/gomongo](https://github.com/clarkmcc/gomongo)
