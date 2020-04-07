# golang
## [1][GoLang Game Development : devlog 04 (pixel lib)](https://www.reddit.com/r/golang/comments/fwimmw/golang_game_development_devlog_04_pixel_lib/)
- url: https://www.youtube.com/watch?v=XBQ6jMGGk_Y
---

## [2][week 2 from my learning go series is up! check it out 👍🏻](https://www.reddit.com/r/golang/comments/fwizn4/week_2_from_my_learning_go_series_is_up_check_it/)
- url: https://martincartledge.io/learning-go-week-2
---

## [3][A lightweight job scheduler for Golang web apps](https://www.reddit.com/r/golang/comments/fwdsat/a_lightweight_job_scheduler_for_golang_web_apps/)
- url: https://www.reddit.com/r/golang/comments/fwdsat/a_lightweight_job_scheduler_for_golang_web_apps/
---
Hi Gophers,

I have written a job scheduler in Go which can be used in any Golang web apps. It does not have any dependency on any third-party library. It is very small and lightweight. It uses a priority queue to distribute requests to workers. The library makes all the backend calls asynchronously with retry, timeout and context cancellation functionality. It also provides very easy semantics to join multiple data sources based on their output and input types, at the same time having no coupling between the data sources. 

[https://github.com/susamn/rio](https://github.com/susamn/rio)

The included example shows how to use it. 

&amp;#x200B;

Reviews are welcome !!!
## [4][NATS Messaging - Part 1 - choria.io/blog](https://www.reddit.com/r/golang/comments/fwg5na/nats_messaging_part_1_choriaioblog/)
- url: https://choria.io/blog/post/2020/03/23/nats_patterns_1/
---

## [5][Example Rest API](https://www.reddit.com/r/golang/comments/fw0pv7/example_rest_api/)
- url: https://www.reddit.com/r/golang/comments/fw0pv7/example_rest_api/
---
Hello gophers ✌️

I've been working on an example Rest API to show how I'm currently building them in production.

Hopefully someone will find this useful. All feedback, suggestions and discussions welcome.

[https://github.com/cobbinma/example-go-api](https://github.com/cobbinma/example-go-api)

Thanks!

&amp;#x200B;

https://preview.redd.it/k7lh4l8av7r41.jpg?width=800&amp;format=pjpg&amp;auto=webp&amp;s=53f7feddc7ac999271a7efafba28868ba079be56
## [6][Are there any gophers here who are looking for a side project?](https://www.reddit.com/r/golang/comments/fwi52w/are_there_any_gophers_here_who_are_looking_for_a/)
- url: https://www.reddit.com/r/golang/comments/fwi52w/are_there_any_gophers_here_who_are_looking_for_a/
---
I'd like to bring in a few low-middle experience gophers who want to learn and contribute to an open source network routing project, and at least one expert-level gopher who has some extra time on their hands.
## [7][Micro In Action: Error Handling Across Service Boundaries](https://www.reddit.com/r/golang/comments/fwkehd/micro_in_action_error_handling_across_service/)
- url: https://medium.com/@dche423/micro-in-action-error-handling-across-services-boundary-2f9f27821bd5?source=friends_link&amp;sk=5e665efa74853455391d877d8aaf500f
---

## [8][Building Go Documentation](https://www.reddit.com/r/golang/comments/fwg32p/building_go_documentation/)
- url: https://www.reddit.com/r/golang/comments/fwg32p/building_go_documentation/
---
TL;DR: Are there any good static site documentation generators that take inline docs and spit out either HTML or PDF's like rust's [cargo doc command](https://doc.rust-lang.org/cargo/commands/cargo-doc.html) or the [pdoc3 package](https://pdoc3.github.io/pdoc/) for python? I am aware of [godoc static](https://gitlab.com/tslocum/godoc-static?0.1.5) but it does not give you per-function documentation, just per-package.

&amp;#x200B;

For reference here is the file I am looking to document (more to test it out on a simple program than anything):  [https://github.com/Descent098/simple-otp/blob/master/go/otp.go](https://github.com/Descent098/simple-otp/blob/master/go/otp.go) 

&amp;#x200B;

I have been messing around a bit with go recently trying to see if I want to invest time into learning it deeply. I had used it in the past and one of the things that blew me away was that it was the first language I used that included an easy to use documentation system (godoc). Admittedly I seemed to remember it being able to output to static documentation, but I think I may just have rose tinted nostalgia glasses on since it doesn't seem to have ever been the case.

&amp;#x200B;

Anyway since using go originally I have been using rust and cargo's doc tool and the pdoc3 module for python have spoil me with generating nice documentation from docstrings in programs. Now however it looks like:

1. godoc is no longer included in go (I can live with go getting it even though google's results gave me the wrong answer a few times)
2. Even when I got it setup and use [godoc static](https://gitlab.com/tslocum/godoc-static?0.1.5) I cannot for the life of me wrap my head around getting it to generate good docs on a single file program since godoc static builds package, but not function docs.

&amp;#x200B;

Am I missing something here or is the commonplace to generate manual docs in an external system for API's, scripts and such? 

&amp;#x200B;

P.S. This post is not a dig at go, just a genuine question so I know what the standard practices are in the community going forward it seems like a great language and I want to make sure to get this right from the beginning.
## [9][Tests fail when run individually but pass when run together](https://www.reddit.com/r/golang/comments/fwkc53/tests_fail_when_run_individually_but_pass_when/)
- url: https://www.reddit.com/r/golang/comments/fwkc53/tests_fail_when_run_individually_but_pass_when/
---
Seeing some strange behavior from my tests. When I run all my tests using \`go test ./...\` they all pass, however I have two tests that when run individually fail with the error 

    Failed to run &lt;myDependency&gt;` exec: \"&lt;myDependency&gt;\": executable file not found in $PATH

The app runs fine when built and the dependencies are in the $PATH,  but something about the testing is off. Any ideas?
## [10][Q : database/sql - can I create a Row struct from my code ?](https://www.reddit.com/r/golang/comments/fwhisv/q_databasesql_can_i_create_a_row_struct_from_my/)
- url: https://www.reddit.com/r/golang/comments/fwhisv/q_databasesql_can_i_create_a_row_struct_from_my/
---
Hello gophers,

I am posting here a question I also asked on StackOverflow : [https://stackoverflow.com/q/61076580/86072](https://stackoverflow.com/q/61076580/86072)

**Context**

I am trying to write code that would allow "nested" transactions, using PostgreSQL savepoints.

I would like to have a common interface, which would look like :

    type interface Trx{
      Begin() (Trx, error)
      Commit() error
      Rollback() error

      Query(...) (*sql.Rows, error)
      QueryRow(...) *sql.Row
      Exec(...) (sql.Result, error)
    }

I would like my code would introduce new error cases when making queries : a transaction should not be used while a "sub transaction" is active.

**Issue**

My issue is with the `QueryRow()` method :

  * as `Query()` and `Exec()` methods return their error in a separate value, I can return any error I want from my implementation, along with the zero value for the result part ;
  * for `QueryRow()`, however, my custom error should be wrapped in as `*sql.Row{}` struct, so that the next call to `.Scan()` returns the error.

`sql.Row` is a struct, not an interface, and I see no way to create something else than the zero value from my code.

**Question**

Is there a way to create a `sql.Row{}` struct, that would not trigger any call to any underlying driver function, and just wrap a known error ?
