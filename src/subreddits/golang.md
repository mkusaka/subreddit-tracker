# golang
## [1][Improve your Micro-Services Debugging Environment using Gebug and Telepresence](https://www.reddit.com/r/golang/comments/j365lx/improve_your_microservices_debugging_environment/)
- url: https://www.reddit.com/r/golang/comments/j365lx/improve_your_microservices_debugging_environment/
---
Check out my new blog-post on how to improve your Micro-Services Debugging Environment using Gebug and Telepresence 🎉

[https://medium.com/@moshe.beladev.mb/better-debugging-environment-for-your-micro-services-9420a71b8a37](https://medium.com/@moshe.beladev.mb/better-debugging-environment-for-your-micro-services-9420a71b8a37)
## [2][Faster Literal String Matching in Go](https://www.reddit.com/r/golang/comments/j31dm0/faster_literal_string_matching_in_go/)
- url: https://boyter.org/posts/faster-literal-string-matching-in-go/
---

## [3][Elsa is a minimal, fast and secure runtime for Javascript and Typescript written in Go](https://www.reddit.com/r/golang/comments/j2rjps/elsa_is_a_minimal_fast_and_secure_runtime_for/)
- url: https://github.com/elsaland/elsa
---

## [4][How to upload a file to postgres db with golang](https://www.reddit.com/r/golang/comments/j30hs7/how_to_upload_a_file_to_postgres_db_with_golang/)
- url: https://www.reddit.com/r/golang/comments/j30hs7/how_to_upload_a_file_to_postgres_db_with_golang/
---
Is there any good blogs or tutorial available on how to *upload* and get *uploaded* files using postgres and golang.
## [5][Boilerplating a New Go Program (Microservice) - Qvault](https://www.reddit.com/r/golang/comments/j38fd1/boilerplating_a_new_go_program_microservice_qvault/)
- url: https://qvault.io/2020/10/01/boilerplating-a-new-go-program-microservice/
---

## [6][Jupyter Notebooks for Go](https://www.reddit.com/r/golang/comments/j2n2an/jupyter_notebooks_for_go/)
- url: https://www.reddit.com/r/golang/comments/j2n2an/jupyter_notebooks_for_go/
---
Just found out about this and definitely a pretty cool implementation if you like Jupyter Notebooks.

&amp;#x200B;

https://preview.redd.it/qqp9fam1oaq51.png?width=580&amp;format=png&amp;auto=webp&amp;s=74fce34c4dbcbc59d9c888f53871c4bb1f34e189

It works on Mac, Linux, and Windows. There's also a Docker image that you can easily spin up. [https://github.com/gopherdata/gophernotes](https://github.com/gopherdata/gophernotes)

I'll definitely be working on a YouTube video and some posts about this via Twitter: TheNJDevOpsGuy
## [7][A lib to make parallel go testing less verbose](https://www.reddit.com/r/golang/comments/j33etw/a_lib_to_make_parallel_go_testing_less_verbose/)
- url: https://www.reddit.com/r/golang/comments/j33etw/a_lib_to_make_parallel_go_testing_less_verbose/
---
Repo: [github.com/ysmood/got](https://github.com/ysmood/got)

So that you won't forget to write `t.Parallel()` in every test case.

```go
package main_test

import (
    "testing"

    "github.com/ysmood/got"
)

func Test(t *testing.T) {
    got.Each(t, func(t *testing.T) S {
        t.Parallel()
        return S{got.New(t)}
    })
}

type S struct {
    got.Assertion
}

func (s S) A() { // just like the: func TestA(t *testing.T) {}
    s.Eq(1, 1.0)
}

func (b S) B() {
    b.Eq([]int{1, 2}, []int{1, 2})
}
```

The lib has zero dependencies. It's easy to use it with any other assert lib, such as testify:

```go
package main_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ysmood/got"
)

func Test(t *testing.T) {
    got.Each(t, func(t *testing.T) Suite {
        return Suite{assert.New(t)}
    })
}

type Suite struct {
	*assert.Assertions
}

func (s Suite) A() {
	s.Equal(1, 1)
}

func (s Suite) B() {
	s.Equal(2, 1)
}
```
## [8][ent + gqlgen = &lt;3](https://www.reddit.com/r/golang/comments/j2ljgh/ent_gqlgen_3/)
- url: https://www.reddit.com/r/golang/comments/j2ljgh/ent_gqlgen_3/
---
Hey guys,
a few weeks ago, I mentioned here that we'll OSS our GraphQL integration for [ent](https://github.com/facebook/ent). So, as promised, here's the link to the docs [entgo.io/docs/graphql](https://entgo.io/docs/graphql/), and to the [entgql extension](https://github.com/facebookincubator/ent-contrib) in GitHub. 


What exactly we currently provide in this integration: 

- [Node API](https://entgo.io/docs/graphql/#node-api) support. 

- [Pagination support](https://entgo.io/docs/graphql/#pagination) - Follows the *Relay Cursor Connections Spec*.

- [Connection Ordering](https://entgo.io/docs/graphql/#connection-ordering).

- [Fields Collection](https://entgo.io/docs/graphql/#fields-collection) - (*This one is super cool!*) Auto eager-load relations/edges that exist in the GraphQL request.

- [Transactional Mutations](https://entgo.io/docs/graphql/#transactional-mutations) - A handler that executes each GraphQL mutation in a transaction and either commit or rollback the transaction based on the response. 

- And additional small features that makes the binding between the 2 libraries much friendly. See the docs for more details.
 
We are actively working to improve the integration between the 2 libraries and we expect to OSS additional options in the future. If you're using gqlgen and ent, please give it a look/try and let me know what do you think :)
## [9][Golang Projects to get a job](https://www.reddit.com/r/golang/comments/j2vvxz/golang_projects_to_get_a_job/)
- url: https://www.reddit.com/r/golang/comments/j2vvxz/golang_projects_to_get_a_job/
---
Hey guys, I'm looking for some project recommendations that could get me started on a tech career.

I've been learning Go for two months now, and I am pretty comfortable with the basics up to concurrency and channeling, those are still a challenge. Anyway, I'm looking for a couple of projects that I can do as a beginner to build experience and portfolio, but I'm struggling to find something that could get the attention of recruiters. Can you please suggest something to steer me the right way?

I also have experience with React, but wouldn't like to go for the web backend route of Go. 

Thanks
## [10][What do people say out loud whilst programming in go (just for fun)](https://www.reddit.com/r/golang/comments/j37hyv/what_do_people_say_out_loud_whilst_programming_in/)
- url: https://www.reddit.com/r/golang/comments/j37hyv/what_do_people_say_out_loud_whilst_programming_in/
---
Mine are boring but wondering if others do similar?

I often say semi involuntarily to myself:

1./ Another one bites the dust.
2./ Back in business.
3./ Right, lets get the party started people.
