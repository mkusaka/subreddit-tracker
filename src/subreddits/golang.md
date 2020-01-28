# golang
## [1][bradfitz - Leaving Google](https://www.reddit.com/r/golang/comments/eur3zf/bradfitz_leaving_google/)
- url: https://bradfitz.com/2020/01/27/leaving-google
---

## [2][podlog - merge multiple pod stdio streams to stdout](https://www.reddit.com/r/golang/comments/ev4y3j/podlog_merge_multiple_pod_stdio_streams_to_stdout/)
- url: https://github.com/DoctrinAB/podlog
---

## [3][Thread where beginners can post small questions](https://www.reddit.com/r/golang/comments/ev4uep/thread_where_beginners_can_post_small_questions/)
- url: https://www.reddit.com/r/golang/comments/ev4uep/thread_where_beginners_can_post_small_questions/
---
Hi Gophers! 

I am pretty new to golang and sometimes i have some small questions that I can‘t justify opening a single thread for. I would love to see a (monthly/biweekly?) thread were gophers can ask small questions. 

What do you guys think? Would such a thread be a good idea? 

&amp;#x200B;

Cheers,   
juniorGopher
## [4][Lost converting a project to a Module project](https://www.reddit.com/r/golang/comments/ev4l84/lost_converting_a_project_to_a_module_project/)
- url: https://www.reddit.com/r/golang/comments/ev4l84/lost_converting_a_project_to_a_module_project/
---
I've inherited a Go project (I do not know Go outside of what I picked up the past few days)

It's currently setup with a directory structure like this:

    foo/
        main.go
        foo/foo.go
        bar/bar.go


main imports foo and bar via `github.com/ourdomain/foo/foo` and `github.com/ourdomain/foo/bar` (Private repos in our organization.

This all works when run under in a go workspace though I find the import of internal packages like that odd. 

My need is to move this out of the gopath, our build tool essentially requires the project to live in another directory in order to be mounted into a container running in a virtual env. 

My understanding is that converting this project to a Go Module would be the right way to go about this. When I do however, I get errors that the foo and bar packages cannot be found when I run `go mod tidy` or `go build` etc. 

How should a project like this be set up, where there are multiple packages living under their own sub directories, and how do I import them, particularly under the conditions of a top level Module?

Should that top level Module be in the main package as it is now?

I can change anything thing I need to to get this working in an idiomatic Go structure, I'm just highly confused about the way things are "supposed" to be.
## [5][Supertube: a networked unix pipe utility.](https://www.reddit.com/r/golang/comments/ev3xzh/supertube_a_networked_unix_pipe_utility/)
- url: https://github.com/hobochild/supertube
---

## [6][Asynq - A simple asynchronous task queue library for Go](https://www.reddit.com/r/golang/comments/eunjcl/asynq_a_simple_asynchronous_task_queue_library/)
- url: https://www.reddit.com/r/golang/comments/eunjcl/asynq_a_simple_asynchronous_task_queue_library/
---
Asynq is a simple Go library for queueing tasks and processing them in the background with workers.  
It is backed by Redis and it is designed to have a low barrier to entry. It should be integrated in your web stack easily.

[github.com/hibiken/asynq](https://github.com/hibiken/asynq)
## [7][A Go 1 compatible final value proposal](https://www.reddit.com/r/golang/comments/ev42e3/a_go_1_compatible_final_value_proposal/)
- url: https://github.com/go101/go101/wiki/An-immutable-value-proposal-to-let-Go-support-final-exported-values
---

## [8][Noob question: Object Initialization Types](https://www.reddit.com/r/golang/comments/ev0rst/noob_question_object_initialization_types/)
- url: https://www.reddit.com/r/golang/comments/ev0rst/noob_question_object_initialization_types/
---
**Question:**  
1) What is the difference between Type1 and Type2?

2) Which factor decides whether to use Type 1 or Type 2?

3) Is Type 2 always the preferred way to initialize objects?

**Type 1:**

    type Test struct {
        db *sqlx.DB
        t string
    }
    
    func NewTest(db *sqlx.DB, t string) Test {
        return Test{
            db: db,
            t: t,
        }
    }

**Type 2:**

    type Test struct {
        db *sqlx.DB
        t string
    }
    
    func NewTest(db *sqlx.DB, t string) *Test {
        return &amp;Test{
            db: db,
            t: t,
        }
    }

From what I read I understood that if I use Type1, then if there is change in value of object, then the change doesn't reflect in other places and it leads to inconsistency.

Whereas type 2, since its a pointer it is safe to use and it remains consistent across the application.  
Please do correct me if I'm wrong.
## [9][go modules and internal packages](https://www.reddit.com/r/golang/comments/euwoko/go_modules_and_internal_packages/)
- url: https://www.reddit.com/r/golang/comments/euwoko/go_modules_and_internal_packages/
---
I have so far not migrated towards the new go modules. I am beginning the process. Looking for ideas.

I split my applications into independent packages each in its own directory.

the structure might look like:

top/src

         /pkg1

         /pkg2

the imports are based on relative paths - eg:

import (

    "../pkg1"

)

Questions:

\- there is just 1 go.mod at the top level right?

\- go.mod only specifies what the external packages are. [eg.google.olang.org/package](http://eg.google.olang.org/package)

this appears a bit counter to the go module setup.

what are the recommendations regarding internal packages. 

appreciate pointers.

thanks, srini
## [10][Concise Go](https://www.reddit.com/r/golang/comments/euza0b/concise_go/)
- url: https://www.reddit.com/r/golang/comments/euza0b/concise_go/
---
Hands-on intro to GoLang

 [https://www.amazon.com/Concise-Go-Yaniv-Astamnep-ebook/dp/B0844MWN92](https://www.amazon.com/Concise-Go-Yaniv-Astamnep-ebook/dp/B0844MWN92)
