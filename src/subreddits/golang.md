# golang
## [1][How to Manage Database Timeouts and Cancellations in Go](https://www.reddit.com/r/golang/comments/g59fdv/how_to_manage_database_timeouts_and_cancellations/)
- url: https://www.alexedwards.net/blog/how-to-manage-database-timeouts-and-cancellations-in-go
---

## [2][a library for creating ISO disk images in Go](https://www.reddit.com/r/golang/comments/g5c9ut/a_library_for_creating_iso_disk_images_in_go/)
- url: https://github.com/kdomanski/iso9660
---

## [3][WebRTC For Gophers: Learn to build sub-second decentralized real-time communications software](https://www.reddit.com/r/golang/comments/g4xsws/webrtc_for_gophers_learn_to_build_subsecond/)
- url: https://www.youtube.com/watch?v=FdgoOrJH8ok&amp;feature=youtu.be&amp;t=989
---

## [4][Go 2019 Survey Results - The Go Blog](https://www.reddit.com/r/golang/comments/g4z83u/go_2019_survey_results_the_go_blog/)
- url: https://blog.golang.org/survey2019-results
---

## [5][What exactly is/was the overhead of defer?](https://www.reddit.com/r/golang/comments/g5bghz/what_exactly_iswas_the_overhead_of_defer/)
- url: https://www.reddit.com/r/golang/comments/g5bghz/what_exactly_iswas_the_overhead_of_defer/
---
In the release notes for Go 1.14 I found this:

&gt; This release improves the performance of most uses of defer to incur almost zero overhead compared to calling the deferred function directly. As a result, defer can now be used in performance-critical code without overhead concerns. As a result, defer can now be used in performance-critical code without overhead concerns.

I don't have a CS background but I know what defer does. What I don't understand is why does defer have overhead. And not just overhead, but enough overhead so that it used to be a somewhat significant performance disadvantage, as the last sentence seems to imply.

In other words: How exactly is `defer fmt.Println("Done")` more overhead intensive than `fmt.Println("Done)` and where does the overhead come from?

Can this be explained in a somewhat layman friendly way?
## [6][It's a map ... it's a heap ... no! It's a non-blocking, persistent treap!](https://www.reddit.com/r/golang/comments/g56pe8/its_a_map_its_a_heap_no_its_a_nonblocking/)
- url: https://www.reddit.com/r/golang/comments/g56pe8/its_a_map_its_a_heap_no_its_a_nonblocking/
---
Hi everyone,

I've been working on a peer-to-peer project that presented me with an interesting problem.  I need to maintain a large collection of structs that look like this:

    type Entry struct {
        TTL time.Time
        SomeData interface{}
    }

The object of the game is:

1. to update the TTLs as messages from peers arrive
2. to expire entries if the TTL is exceeded

TTL updates come in the form of a `(peer.ID, TTL)` tuple, so I needed a datastructure that provides:

1. Fast keyed lookups (like a map)
2. Fast sorting ... but by *TTL*, not by key!

Oh, and the whole damn thing is highly concurrent ...

After fiddling around with some ugly lock-based solutions, I discovered the wonderful world of treaps!  More specifically, the wonderful world of purely-functional, immutable treaps.  After a spectacular performance gain in my project, I decided to move this treap implementation into its own library.

Please let me know if you spot any bugs!

Also, if anyone here is an expert on treaps, I'd really like to optimize this a bit.  Most of the performance benefits I'm getting are related to improvements in time-complexity over my previous approach, as well as the fact that readers no longer wait for writers.  That said, I've done _zero_ code optimization, so I suspect there's some low-hanging fruit.  For example, I'm wondering if the the `upsert` function be re-written iteratively to shave off a few allocations.  Here are my current benchmarks (which are probably shakey):

    $ go test -benchmem github.com/lthibault/treap -bench .
    goos: darwin
    goarch: amd64
    pkg: github.com/lthibault/treap
    BenchmarkInsertSync-4             306808              7915 ns/op            2090 B/op         35 allocs/op
    BenchmarkSplitSync-4              125773              9926 ns/op            4403 B/op         69 allocs/op
    BenchmarkMergeSync-4              621154              6757 ns/op            1627 B/op         25 allocs/op
    BenchmarkDeleteSync-4             162746             12557 ns/op            6115 B/op         96 allocs/op
    BenchmarkPopSync-4                832798              4344 ns/op            1369 B/op         21 allocs/op
    BenchmarkSetWeightSync-4          383072              8460 ns/op            1772 B/op         29 allocs/op
    BenchmarkIterSync-4              2912358               430 ns/op             128 B/op          1 allocs/op
    PASS
    ok      github.com/lthibault/treap      48.998s


# TL;DR

[Here's the GitHub repo](https://github.com/lthibault/treap)!
## [7][Use Anonymous Structs For JSON Marshalling in Go](https://www.reddit.com/r/golang/comments/g5evlp/use_anonymous_structs_for_json_marshalling_in_go/)
- url: https://qvault.io/2020/04/21/use-anonymous-structs-for-json-marshalling-in-go/
---

## [8][In ear monitors](https://www.reddit.com/r/golang/comments/g5e3q1/in_ear_monitors/)
- url: https://www.reddit.com/r/golang/comments/g5e3q1/in_ear_monitors/
---
Almost a year ago I tried to create something like audiofusion in ear system in go. I got there almost. From computer to computer I got a working prototype with almost unnoticeable latency. But I tried computer to mobile using react native and a custom plugin to handle all of the audio buffers and UDP but I got some weird sounds and a big latency so I stopped working on this project. But now i thought let's just trow it on reddit maybe someone is liking it and wants to develop it's further or maybe it gives me more motivation to develop it further

Why do i use UDP? It's faster than TCP.

Why did you use react native with custom plugin? I'm not very into developing apps and react native is almost the same as ReactJS that i'm familiar with

[https://github.com/davisnando/in\_ear\_system](https://github.com/davisnando/in_ear_system)

&amp;#x200B;

Not so clean repo because I was too lazy to clean this. It wasn't on git when developing

[https://github.com/davisnando/in\_ear\_system\_app](https://github.com/davisnando/in_ear_system_app)
## [9][Go assignment by slice, array unpacking or destructuring](https://www.reddit.com/r/golang/comments/g5dv1q/go_assignment_by_slice_array_unpacking_or/)
- url: https://github.com/thedevsaddam/unpack
---

## [10][Making a multiplayer game with Go and gRPC](https://www.reddit.com/r/golang/comments/g4z1u1/making_a_multiplayer_game_with_go_and_grpc/)
- url: https://mortenson.coffee/blog/making-multiplayer-game-go-and-grpc/
---

