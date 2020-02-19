# golang
## [1][Build complex APIs in GO without writing code.](https://www.reddit.com/r/golang/comments/f67vrz/build_complex_apis_in_go_without_writing_code/)
- url: https://www.reddit.com/r/golang/comments/f67vrz/build_complex_apis_in_go_without_writing_code/
---
Super Graph is a GraphQL to SQL compiler that offers features like nested queries, nested inserts and updates, cursor pagination, full-text search, and a lot more right out of the box. It works with Postgres and highly scalable Yugabyte DB.

Say you had to fetch the current user, his info, the top 20 threads along with a cursor to fetch more, the votes related to each thread and all the posts on each thread also sorted along with another cursor to paginate the posts and the author of each post. 

This is a heavy request and will need a lot of code in any language or framework. With Super Graph the below GraphQL will be converted into a single efficient SQL query (1 query only) without you writing a line of code. 

It's a pet project, open source and entirely in GO and not a startup it's taken me a lot of time and effort to get it to where it is today I'd love to have more people use it and contribute. It's well documented even got an internals doc for those who want to contribute code. 

FYI Super Graph literally saved me many hundreds of man hours on projects I only focus on the React or Mobile UI while Super Graph is the entire backend.

[https://github.com/dosco/super-graph](https://github.com/dosco/super-graph)

```graphql
query { 
 user : me {
   slug 
   firstName : first_name 
   lastName : last_name 
   pictureURL : picture_url 
   bio
  }
  thread(first: 20, after: $cursor, order_by: { cached_votes_total: desc }) {
    slug
    title
    published
    createdAt : created_at
    totalVotes : cached_votes_total
    totalPosts : cached_posts_total
    vote : thread_vote(where: { user_id: { eq: $user_id } }) {
      created_at
    }
    topics {
      slug
      name
    }
    author : me {
      slug
    }
    posts(first: 20, order_by: { cached_votes_total: desc }) {
      slug
      body
      published
      createdAt : created_at
      totalVotes : cached_votes_total
      vote : post_vote(where: { user_id: { eq: $user_id } }) {
        created_at
      }
      author : user {
        slug 
        firstName : first_name
        lastName : last_name
      }
    }
  }
}
```
## [2][Testing in Go: Clean Tests Using t.Cleanup](https://www.reddit.com/r/golang/comments/f61egk/testing_in_go_clean_tests_using_tcleanup/)
- url: https://ieftimov.com/post/testing-in-go-clean-tests-using-t-cleanup/
---

## [3][go-cov-rollup - roll up concatenated go coverage reports (like from go-acc)](https://www.reddit.com/r/golang/comments/f68ff2/gocovrollup_roll_up_concatenated_go_coverage/)
- url: https://www.reddit.com/r/golang/comments/f68ff2/gocovrollup_roll_up_concatenated_go_coverage/
---
[https://github.com/redstarnv/go-cov-rollup](https://github.com/redstarnv/go-cov-rollup)

The other day we ran into a problem with go modules coverage and Code Climate. We use `go-acc` to get an accurate coverage report for a project split into multiple modules that has integration tests covering more than one module. It works well for local coverage reports via `go tool cover`, but once we tried uploading it to Code Climate - the coverage numbers were off there.

After looking into it a but further, it turned out that as concatenated report produced by `go-acc` contains multiple entries for the same statement with different hit counts – e.g. one for every time it was touched by a test – Code Climate doesn't correctly roll those up, and reports only the first hit, which is wrong.

To work around it, I wrote a quick tool that takes coverage report(s) from stdin, rolls up hit counts and outputs summarised report to stdout. This can be easily plugged into CI pipelines before the final report is sent off to an external service like Code Climate.

Hope someone else finds it useful, too.
## [4][The Goyave framework replaces gorilla/mux with his own router, which is twice as fast and uses 3 times less memory](https://www.reddit.com/r/golang/comments/f6aktu/the_goyave_framework_replaces_gorillamux_with_his/)
- url: https://www.reddit.com/r/golang/comments/f6aktu/the_goyave_framework_replaces_gorillamux_with_his/
---
Hello ! The v2.6.0 release is out today. The focus was on routing optimization. This is why I got rid of gorilla/mux for routing and implemented my own router. This release doesn't break backwards compatibility and opens new possibilities for the framework. Also, the performance greatly increased: the new router is twice as fast and uses about 3 times less memory.

&amp;#x200B;

**Benchmark:**

OS: Ubuntu 19.10

Arch: amd64

CPU: Intel Core i7-6700HQ @ 8x 3.5GHz (Laptop)

RAM: 16GiB

**Before** (gorilla/mux)**:**

    BenchmarkRouteRegistration-8               12631             92962 ns/op           60023 B/op        814 allocs/op
    BenchmarkRootLevelNotFound-8            10202770               121 ns/op              48 B/op          1 allocs/op
    BenchmarkRootLevelMatch-8                2993304               392 ns/op             112 B/op          3 allocs/op
    BenchmarkRootLevelPostMatch-8            2586350               462 ns/op             112 B/op          3 allocs/op
    BenchmarkRootLevelPostParamMatch-8       1252915               973 ns/op             416 B/op          4 allocs/op
    BenchmarkSubrouterMatch-8                 912388              1302 ns/op             112 B/op          3 allocs/op
    BenchmarkSubrouterPostMatch-8             897088              1314 ns/op             112 B/op          3 allocs/op
    BenchmarkSubrouterNotFound-8              662738              1746 ns/op             113 B/op          3 allocs/op
    BenchmarkParamMatch-8                     678423              1650 ns/op             436 B/op          5 allocs/op
    BenchmarkParamPutMatch-8                  564192              2016 ns/op             436 B/op          5 allocs/op
    BenchmarkParamDeleteMatch-8               445140              2267 ns/op             436 B/op          5 allocs/op
    BenchmarkMatchAll-8                        90562             12947 ns/op            2347 B/op         35 allocs/op

**After** (new router):

    BenchmarkRouteRegistration-8               43315             27346 ns/op           21765 B/op        278 allocs/op
    BenchmarkRootLevelNotFound-8            18301303                59.4 ns/op             0 B/op          0 allocs/op
    BenchmarkRootLevelMatch-8                7645866               151 ns/op              16 B/op          1 allocs/op
    BenchmarkRootLevelPostMatch-8            5557107               216 ns/op              16 B/op          1 allocs/op
    BenchmarkRootLevelPostParamMatch-8       1538215               765 ns/op             368 B/op          3 allocs/op
    BenchmarkSubrouterMatch-8                3730909               312 ns/op              32 B/op          2 allocs/op
    BenchmarkSubrouterPostMatch-8            2679608               460 ns/op              48 B/op          3 allocs/op
    BenchmarkSubrouterNotFound-8             1290444               910 ns/op              48 B/op          2 allocs/op
    BenchmarkParamMatch-8                    1306718               878 ns/op             388 B/op          4 allocs/op
    BenchmarkParamPutMatch-8                 1056933              1147 ns/op             421 B/op          5 allocs/op
    BenchmarkParamDeleteMatch-8               825115              1328 ns/op             453 B/op          6 allocs/op
    BenchmarkMatchAll-8                       173403              6490 ns/op            1798 B/op         27 allocs/op

You can run the benchmark yourself, it's available on the [Goyave repository.](https://github.com/System-Glitch/goyave) I suspect the performance improvements to be even more significant on large applications with many more routes. I am open to any suggestion (and contribution) for improvements on the benchmark and the router implementation.
## [5][Bloom effect using Go](https://www.reddit.com/r/golang/comments/f5tvng/bloom_effect_using_go/)
- url: https://remy.io/blog/bloom-effect-in-go
---

## [6][API made with Go - is a systemd service the best route to provide uptime, automatic restarts etc?](https://www.reddit.com/r/golang/comments/f6aini/api_made_with_go_is_a_systemd_service_the_best/)
- url: https://www.reddit.com/r/golang/comments/f6aini/api_made_with_go_is_a_systemd_service_the_best/
---
I've develop a small API for the sake of learning the workflow, it works just fine when executed directly in the shell but now I'd like to learn how to make the binary "persistent" across reboots or crashes.

For what it's worth in working under Ubuntu server 18.04 and from my understanding a systemd service is the way to go (opinion due to mostly digital ocean tutorials).

1. Am I pursuing the right path?
1. Most importantly If I coded in the program to read some values from a config.yml in the same folder of the binary, how this should be handled with a service?
## [7][gorilla/mux v1.7.4](https://www.reddit.com/r/golang/comments/f6ahmi/gorillamux_v174/)
- url: https://github.com/gorilla/mux/releases/tag/v1.7.4
---

## [8][Is concurrent read okay in go maps? Just reading no writing](https://www.reddit.com/r/golang/comments/f6af6h/is_concurrent_read_okay_in_go_maps_just_reading/)
- url: https://www.reddit.com/r/golang/comments/f6af6h/is_concurrent_read_okay_in_go_maps_just_reading/
---

## [9][Testing for SSL related outages with Go](https://www.reddit.com/r/golang/comments/f67w1d/testing_for_ssl_related_outages_with_go/)
- url: https://medium.com/@noamt/simulating-ssl-outages-with-go-8f14e5ef0621
---

## [10][Gontainer: a simple and rudimental container for Linux written in GO](https://www.reddit.com/r/golang/comments/f5pda3/gontainer_a_simple_and_rudimental_container_for/)
- url: https://www.reddit.com/r/golang/comments/f5pda3/gontainer_a_simple_and_rudimental_container_for/
---
[https://github.com/alegrey91/Gontainer](https://github.com/alegrey91/Gontainer)
