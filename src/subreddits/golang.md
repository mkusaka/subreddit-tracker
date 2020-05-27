# golang
## [1][Diving into Go by building a CLI application](https://www.reddit.com/r/golang/comments/grgvlu/diving_into_go_by_building_a_cli_application/)
- url: https://eryb.space/2020/05/27/diving-into-go-by-building-a-cli-application.html
---

## [2][[go-quartz] Simple, zero-dependency scheduling library for Go](https://www.reddit.com/r/golang/comments/grft7q/goquartz_simple_zerodependency_scheduling_library/)
- url: https://github.com/reugn/go-quartz
---

## [3][Why are so many Golang programmers against using the goto statement?](https://www.reddit.com/r/golang/comments/grglzm/why_are_so_many_golang_programmers_against_using/)
- url: https://www.reddit.com/r/golang/comments/grglzm/why_are_so_many_golang_programmers_against_using/
---
In my opinion, there are cases where using the goto statement results in cleaner code.
## [4][Robust gRPC communication on Google Cloud Run (but not only!)](https://www.reddit.com/r/golang/comments/grhm3e/robust_grpc_communication_on_google_cloud_run_but/)
- url: https://threedots.tech/post/robust-grpc-google-cloud-run/
---

## [5][Robotgo v0.90.0 is released, Go desktop automation. Huge updated!](https://www.reddit.com/r/golang/comments/gqzk9w/robotgo_v0900_is_released_go_desktop_automation/)
- url: https://www.reddit.com/r/golang/comments/gqzk9w/robotgo_v0900_is_released_go_desktop_automation/
---
[https://github.com/go-vgo/robotgo/releases/tag/v0.90.0](https://github.com/go-vgo/robotgo/releases/tag/v0.90.0)
## [6][SDNS v1.0.0 🎉released. Performance improvement 50% more, stability also improvement.](https://www.reddit.com/r/golang/comments/gri8ci/sdns_v100_released_performance_improvement_50/)
- url: https://github.com/semihalev/sdns
---

## [7][How to print nicely a nested struct](https://www.reddit.com/r/golang/comments/gritgv/how_to_print_nicely_a_nested_struct/)
- url: https://www.reddit.com/r/golang/comments/gritgv/how_to_print_nicely_a_nested_struct/
---
Hello everyone! I'm trying to learn go, and to do so I've created a little quiz game to study psychopathology for an exam. My problem is that I don't like the code I've wrote, so I'm trying to do better, hoping to learn good go practices along the way. 

My question is: if I have a nested struct, with maps and stuff like this...:

```go
type Disorders struct {
    Disorders []Disorder
}

type Disorder struct {
    Name     string
    Symptoms Symptom
    // More stuff
}

type Symptom struct {
    Cognitive map[string]string
    Emotional map[string]string
    // More stuff
}
```

...what is the cleanest way to write a ```.String()``` method for everyone of them, that lets me specify the kind of symptom, the key of the map, etc.

I load the values from a JSON, so it looks like I can't use the [stringer tool](https://godoc.org/golang.org/x/tools/cmd/stringer)

Do you see a better solution to writing every ```.String()``` method manually? 

Btw I wanted to thank this community, you're all very inclusive and supportive, and it's one of the subs I read from most gladly!
## [8][Logging without losing money or context.](https://www.reddit.com/r/golang/comments/gr238r/logging_without_losing_money_or_context/)
- url: https://www.komu.engineer/blogs/log-without-losing-context/log-without-losing-context
---

## [9][Testing in Go: philosophy and tools](https://www.reddit.com/r/golang/comments/gr4d45/testing_in_go_philosophy_and_tools/)
- url: https://lwn.net/SubscriberLink/821358/429b0d0527749a3c/
---

## [10][Bigbucket – Serverless Bigtable-like database, backed by Cloud Storage](https://www.reddit.com/r/golang/comments/grbfms/bigbucket_serverless_bigtablelike_database_backed/)
- url: https://github.com/adrianchifor/Bigbucket
---

