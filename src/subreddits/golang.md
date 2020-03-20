# golang
## [1][gopls 0.3.4 update release notes](https://www.reddit.com/r/golang/comments/fln72e/gopls_034_update_release_notes/)
- url: https://github.com/golang/go/issues/33030#issuecomment-601280048
---

## [2][Go 1.14.1 and Go 1.13.9 are released](https://www.reddit.com/r/golang/comments/flhjz4/go_1141_and_go_1139_are_released/)
- url: https://groups.google.com/d/msg/golang-announce/Ix2U_8WWmXo/a2nJkNW5AAAJ
---

## [3][Buffalo Tests with Cookie](https://www.reddit.com/r/golang/comments/fls9n1/buffalo_tests_with_cookie/)
- url: https://www.reddit.com/r/golang/comments/fls9n1/buffalo_tests_with_cookie/
---
Hi !

  
I am trying to set up few basic tests on my Buffalo API project, but I can't make it works with Cookies..  
I have a Middleware checking Cookie content, and I want to make sure it is properly tested

```go
const Cookie string  = "somename=somevalue"

func (as *ActionSuite) Test_APIAuthorizer_Success() {
	req := as.JSON("/api/v1/foo/bar")
	req.Headers["Cookie"] = Cookie
	res := req.Get()
	as.Equal(200, res.Code)
}
```

Any clue ?
## [4][Illustrated Tales of Go Runtime Scheduler.](https://www.reddit.com/r/golang/comments/flek7p/illustrated_tales_of_go_runtime_scheduler/)
- url: https://medium.com/@ankur_anand/illustrated-tales-of-go-runtime-scheduler-74809ef6d19b
---

## [5][Building and Testing a REST API in Go with Gorilla Mux](https://www.reddit.com/r/golang/comments/flb4bo/building_and_testing_a_rest_api_in_go_with/)
- url: https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
---

## [6][[Recommendation] Does golang have any third-party stateful, structured logging packages](https://www.reddit.com/r/golang/comments/flhx95/recommendation_does_golang_have_any_thirdparty/)
- url: https://www.reddit.com/r/golang/comments/flhx95/recommendation_does_golang_have_any_thirdparty/
---
I am looking for a logger similar to logrus but that lets me set/remove fields that persist and pass the logger instance around in my program. Also it needs to be concurrency safe. I tried building [https://github.com/nitishm/logger](https://github.com/nitishm/logger) but I am running into some significant race conditions.
## [7][Diago, a visualization tool for profiles and heap snapshots generated with pprof](https://www.reddit.com/r/golang/comments/fl8gbv/diago_a_visualization_tool_for_profiles_and_heap/)
- url: https://remy.io/blog/how-to-use-diago-to-diagnose-cpu-and-memory-usage-in-go-programs/
---

## [8][Gophercon 2020 Update](https://www.reddit.com/r/golang/comments/flaz8s/gophercon_2020_update/)
- url: https://blog.gopheracademy.com/gophercon-2020-news/
---

## [9][Linux Window Managers in Go?](https://www.reddit.com/r/golang/comments/fl9lxg/linux_window_managers_in_go/)
- url: https://www.reddit.com/r/golang/comments/fl9lxg/linux_window_managers_in_go/
---
Has anyone come across any Linux Window managers written in Go?

A tiling window manger written in Go would be cool. Like i3 etc..
## [10][A library for initializing heaps from arbitrary slices](https://www.reddit.com/r/golang/comments/flkiw3/a_library_for_initializing_heaps_from_arbitrary/)
- url: https://www.reddit.com/r/golang/comments/flkiw3/a_library_for_initializing_heaps_from_arbitrary/
---
I've written a library to help initializing heaps from arbitrary slices without having to implement all the methods required by `heap.Interface`.

The repo on Github: [https://github.com/suzaku/heaptools](https://github.com/suzaku/heaptools)

Since `reflect` is used, it's slower than when you implement `heap.Interface` manually, the benchmark result is pasted below:

```
BenchmarkNewSliceHeap/init-12            4749110               247 ns/op
BenchmarkNewSliceHeap/init-12            4773699               246 ns/op
BenchmarkNewSliceHeap/init-12            4856407               248 ns/op
BenchmarkNewSliceHeap/init-12            4832068               245 ns/op
BenchmarkNewSliceHeap/init-12            4827453               248 ns/op
BenchmarkNewSliceHeap/push-12            8304612               140 ns/op
BenchmarkNewSliceHeap/push-12            8990208               132 ns/op
BenchmarkNewSliceHeap/push-12            9116899               133 ns/op
BenchmarkNewSliceHeap/push-12            8923878               135 ns/op
BenchmarkNewSliceHeap/push-12            9216303               131 ns/op
BenchmarkNewSliceHeap/pop-12             3839652               350 ns/op
BenchmarkNewSliceHeap/pop-12             3912520               339 ns/op
BenchmarkNewSliceHeap/pop-12             3878684               342 ns/op
BenchmarkNewSliceHeap/pop-12             3833668               357 ns/op
BenchmarkNewSliceHeap/pop-12             3407191               350 ns/op
BenchmarkExplicitImplementation/init-12                 18090414                67.3 ns/op
BenchmarkExplicitImplementation/init-12                 15764011                70.1 ns/op
BenchmarkExplicitImplementation/init-12                 17032052                67.0 ns/op
BenchmarkExplicitImplementation/init-12                 18332876                64.7 ns/op
BenchmarkExplicitImplementation/init-12                 18457176                63.7 ns/op
BenchmarkExplicitImplementation/push-12                 41596758                48.5 ns/op
BenchmarkExplicitImplementation/push-12                 43841671                36.6 ns/op
BenchmarkExplicitImplementation/push-12                 44319072                35.5 ns/op
BenchmarkExplicitImplementation/push-12                 45409448                33.1 ns/op
BenchmarkExplicitImplementation/push-12                 39200691                27.2 ns/op
BenchmarkExplicitImplementation/pop-12                   6022366               236 ns/op
BenchmarkExplicitImplementation/pop-12                   5759280               243 ns/op
BenchmarkExplicitImplementation/pop-12                   5334996               230 ns/op
BenchmarkExplicitImplementation/pop-12                   6122968               229 ns/op
BenchmarkExplicitImplementation/pop-12                   6127668               232 ns/op
```
