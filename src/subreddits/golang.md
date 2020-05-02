# golang
## [1][Mid stack inlining in Go](https://www.reddit.com/r/golang/comments/gc23ed/mid_stack_inlining_in_go/)
- url: https://dave.cheney.net/2020/05/02/mid-stack-inlining-in-go
---

## [2][A lib to make web scraping and automation easy](https://www.reddit.com/r/golang/comments/gc3gue/a_lib_to_make_web_scraping_and_automation_easy/)
- url: https://www.reddit.com/r/golang/comments/gc3gue/a_lib_to_make_web_scraping_and_automation_easy/
---
[Project Repo](https://github.com/ysmood/rod)

Rod is a High-level Chrome Devtools driver directly based on [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/).

* Fluent interface design to reduce verbose code
* Chained context design, intuitive to timeout or cancel the long-running task
* Debugging friendly, auto input tracing, remote monitoring headless browser
* Automatically find or download [chrome](https://github.com/ysmood/rod/blob/master/lib/launcher)
* No external dependencies, [CI](https://github.com/ysmood/rod/actions) tested on Linux, Mac, and Windows
* High-level helpers like WaitStable, WaitRequestIdle, GetDownloadFile, Resource
* Two-step WaitEvent design, never miss an event
* Correctly handles nested iframes
* No zombie chrome process after the crash ([how it works](https://github.com/ysmood/leakless))
## [3][Go GraphQL: Introduction to building GraphQL API in Golang #1 [7 Chapter Tutorial Series ]](https://www.reddit.com/r/golang/comments/gc3j12/go_graphql_introduction_to_building_graphql_api/)
- url: https://ednsquare.com/story/go-graphql-introduction-to-building-graphql-api-in-golang------QhS1oC
---

## [4][InfoMark - a real-world app with Docker, Postgres, Elm, Swagger, ...](https://www.reddit.com/r/golang/comments/gc3fax/infomark_a_realworld_app_with_docker_postgres_elm/)
- url: https://www.reddit.com/r/golang/comments/gc3fax/infomark_a_realworld_app_with_docker_postgres_elm/
---
Maybe this app is interesting for some of you. It is a combination of GO for the backend and ELM for the frontend using RESTful JSON.

[https://infomark.org/](https://infomark.org/)

The system allows university students in a course to upload their home programming assignments to the system. These uploads are then tested automatically within a docker sandbox. And the results are visible to students and teaching assistants.

Why this might be interesting for you?

It has unit-tests in the CI process against a full PostgreSQL database (no mocked functions), it generates the Swagger API documentation automatically (with each release), the build process checks if even the comments are correct, it has background workers that act like a docker client to spawn processes within a "sandbox", there are some cronjobs within the app, database migration, database backup/restore, server console, email, Prometheus, Grafana ... and ... and ...  and. And everything is packed into a single binary! That's what I really enjoy when using GO (a previous version written in RoR exists).
## [5][Question about gqlgen testing.](https://www.reddit.com/r/golang/comments/gc64y0/question_about_gqlgen_testing/)
- url: https://www.reddit.com/r/golang/comments/gc64y0/question_about_gqlgen_testing/
---
Does anyone know how to test authentication with [gqlgen](https://gqlgen.com/) client? There is a [recipe](https://gqlgen.com/recipes/authentication/) from their document and says we can get user's information from the context then check authentication, it works fine in the real api, but  when I write some unit testing for the api, I don't konw where to put the user's information to the client's context.

Is there something I was doing wrong or missing? Thanks.
## [6][Writing a protoc plugin with google.golang.org/protobuf](https://www.reddit.com/r/golang/comments/gc4r2f/writing_a_protoc_plugin_with/)
- url: https://medium.com/@tim.r.coulson/writing-a-protoc-plugin-with-google-golang-org-protobuf-cd5aa75f5777
---

## [7][Go: Asynchronous Preemption](https://www.reddit.com/r/golang/comments/gbhnk1/go_asynchronous_preemption/)
- url: https://medium.com/a-journey-with-go/go-asynchronous-preemption-b5194227371c
---

## [8][Elver, a small Advent of Code utility tool for running your solutions](https://www.reddit.com/r/golang/comments/gc3z5a/elver_a_small_advent_of_code_utility_tool_for/)
- url: https://github.com/aod/elver
---

## [9][I'm going to start learning Go, is this resource still valid? I only ask because in JS, books from 5 years ago are pretty useless](https://www.reddit.com/r/golang/comments/gblhly/im_going_to_start_learning_go_is_this_resource/)
- url: https://www.reddit.com/r/golang/comments/gblhly/im_going_to_start_learning_go_is_this_resource/
---
When uni finished next month, I want to learn Go by using the book **The Go Programming Language** [https://www.gopl.io/](https://www.gopl.io/) which seems to be a really popular resource

Is it considered 'up to date'? I ask this because I attempted to learn JS a few years ago with books. This was a bad move in my opinion. When I tried to install libraries and use them in my code, they were effectively out of date and they wouldn't work, it felt like a waste of money.

So how do gophers here feel this book would be for someone looking to start Go (in particularly get into server side development for the web). 

**Bonus question**  
How is the Go ecosystem in general? Does it move as fast as JS? One of the key reasons I am asking this is because I recently attempted to create a REST API in node, but my knowledge of Express was a few years old. This made me run into a lot more problems than I expected.
## [10][Accelerating aggregate MD5 hashing up to 800% with AVX512](https://www.reddit.com/r/golang/comments/gbll24/accelerating_aggregate_md5_hashing_up_to_800_with/)
- url: https://blog.min.io/accelerating-aggregate-md5-hashing-up-to-800-with-avx512-2/
---

