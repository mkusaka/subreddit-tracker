# golang
## [1][Go update 1.14.3](https://www.reddit.com/r/golang/comments/gk49sx/go_update_1143/)
- url: https://www.reddit.com/r/golang/comments/gk49sx/go_update_1143/
---
[https://golang.org/dl/](https://golang.org/dl/)
## [2][What are some languages written in Go?](https://www.reddit.com/r/golang/comments/gk6860/what_are_some_languages_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/gk6860/what_are_some_languages_written_in_go/
---
Up to now, I only have found out Tnego and Goby.
## [3][KES is a simple, stateless and distributed key-management system](https://www.reddit.com/r/golang/comments/gk3s4u/kes_is_a_simple_stateless_and_distributed/)
- url: https://github.com/minio/kes
---

## [4][Go - from Ideas to Production](https://www.reddit.com/r/golang/comments/gjr80b/go_from_ideas_to_production/)
- url: https://www.reddit.com/r/golang/comments/gjr80b/go_from_ideas_to_production/
---
I've written a [blog](https://www.sarvika.com/2020/05/14/go-from-ideas-to-production/) that's intended for people who are curious about developing in Go that starts with an idea about an application and ends with packaging it in a Docker container that's ready to serve the real world. I'd like some feedback on it and how I can improve as a Go developer.
## [5][Run your Go code in Python using gRPC](https://www.reddit.com/r/golang/comments/gk5woy/run_your_go_code_in_python_using_grpc/)
- url: https://medium.com/swlh/grpc-101-run-your-go-code-in-python-1aab3df732
---

## [6][How do you handle optional sql transactions?](https://www.reddit.com/r/golang/comments/gk78v9/how_do_you_handle_optional_sql_transactions/)
- url: https://www.reddit.com/r/golang/comments/gk78v9/how_do_you_handle_optional_sql_transactions/
---
Let's say I have a store:

    type Store struct {
    	db *db.Connection
    }
    
    func (s Store) Update(ctx context.Context, item *entity.User) (*entity.User, error) {
    	// ...
    }

So far, so good. But how do you go about use-cases where this method needs to be part of a transaction? Where does the transaction start? How do you design your methods so the queries can be run with or without a transaction?
## [7][Auto generation of swagger files](https://www.reddit.com/r/golang/comments/gk8ve6/auto_generation_of_swagger_files/)
- url: https://www.reddit.com/r/golang/comments/gk8ve6/auto_generation_of_swagger_files/
---
Hi there,

Is there a way to auto generate swagger file for an http service, written in golang with gin?  
I prefer not to rewrite all endpoints notation comments in the code base, and I was wondering if this is somehow possible do be done using some sort of a live gin handler that wraps the endpoints and generate the files based on live traffic in a given http server?
## [8][Write your own Go static analysis tool](https://www.reddit.com/r/golang/comments/gk2u3b/write_your_own_go_static_analysis_tool/)
- url: https://about.houqp.me/posts/write-your-own-go-static-analysis-tool/
---

## [9][Mind your Gs, Ms, and Ps in Go](https://www.reddit.com/r/golang/comments/gjw9g8/mind_your_gs_ms_and_ps_in_go/)
- url: https://github.com/golang/go/blob/master/src/runtime/HACKING.md
---

## [10][What is the Golang equivalent of Node.js' read and write streams](https://www.reddit.com/r/golang/comments/gk5v6t/what_is_the_golang_equivalent_of_nodejs_read_and/)
- url: https://www.reddit.com/r/golang/comments/gk5v6t/what_is_the_golang_equivalent_of_nodejs_read_and/
---
I want to serve a relatively large encrypted AWS S3 object without writing to disk or storing all of it in memory and then write to the response.
