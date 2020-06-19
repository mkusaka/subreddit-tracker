# golang
## [1][Experimenting with Golang and Webassembly](https://www.reddit.com/r/golang/comments/hbvoq1/experimenting_with_golang_and_webassembly/)
- url: https://morioh.com/p/15d4299249d2
---

## [2][What does 'fatal error: "stack growth after fork"' mean?](https://www.reddit.com/r/golang/comments/hc0dlb/what_does_fatal_error_stack_growth_after_fork_mean/)
- url: https://www.reddit.com/r/golang/comments/hc0dlb/what_does_fatal_error_stack_growth_after_fork_mean/
---
Hi, I am getting this error in my logs and am trying to find more information on it but am having difficulty.  I see it is mention on line 922 [https://golang.org/src/runtime/stack.go?s=32583:32628](https://golang.org/src/runtime/stack.go?s=32583:32628) but am having difficulty understanding it. I think it means that it tried to allocate a new stack because the old one is full, but fails on this if statement \`thisg.m.morebuf.g.ptr().stackguard0 == stackFork\`.

This is an issue with too many variables being assigned on the stack? Would the best solution to debug this be to apply  \`GOTRACEBACK=crash\` and taking a look at the core file generated in delve or would pprof or something else be better in this circumstance?
## [3][A Concise Guide to the Latest Go Generics Draft Design](https://www.reddit.com/r/golang/comments/hbelxf/a_concise_guide_to_the_latest_go_generics_draft/)
- url: https://pmihaylov.com/go-generics-draft-design/
---

## [4][Go library to handle tens of thousands SSH connections for building network device / server automation.](https://www.reddit.com/r/golang/comments/hbh5lu/go_library_to_handle_tens_of_thousands_ssh/)
- url: https://github.com/yahoo/vssh
---

## [5][functional library with go generic](https://www.reddit.com/r/golang/comments/hbzaec/functional_library_with_go_generic/)
- url: https://github.com/smallnest/gofu
---

## [6][Golang's Superior Cache Solution to Memcached and Redis](https://www.reddit.com/r/golang/comments/hbn7xz/golangs_superior_cache_solution_to_memcached_and/)
- url: https://www.mailgun.com/blog/golangs-superior-cache-solution-memcached-redis/
---

## [7][Olric: Distributed caching in Golang](https://www.reddit.com/r/golang/comments/hbynfe/olric_distributed_caching_in_golang/)
- url: https://github.com/buraksezer/olric
---

## [8][EnvGit runs another command with env. variables loaded from Git. Inspired by EnvDir.](https://www.reddit.com/r/golang/comments/hbvqvi/envgit_runs_another_command_with_env_variables/)
- url: https://www.reddit.com/r/golang/comments/hbvqvi/envgit_runs_another_command_with_env_variables/
---
I wrote  [https://github.com/unravela/envgit](https://github.com/unravela/envgit) . It's simple tool in Go that is inspired by EnvDir. I wrote it because I needed some simple, centralised and versioned configuration for few of my apps and I don't want to bother with Consul.  Maybe it would be useful for somebody.
## [9][How can i add the others licence in my project?](https://www.reddit.com/r/golang/comments/hc08ow/how_can_i_add_the_others_licence_in_my_project/)
- url: https://www.reddit.com/r/golang/comments/hc08ow/how_can_i_add_the_others_licence_in_my_project/
---
Guys how to inmplment another one project licence in your project?
## [10][benchmark result of set value of struct](https://www.reddit.com/r/golang/comments/hby5xp/benchmark_result_of_set_value_of_struct/)
- url: https://www.reddit.com/r/golang/comments/hby5xp/benchmark_result_of_set_value_of_struct/
---
    //Reply - 
    type Reply struct {
    	Code    int32
    	Message string
    	Num     uint64
    	ID      string
    }
    
    //Fill -
    func (r *Reply) Fill(code int32, msg string, n uint64, id string) {
    	r.Code = code
    	r.Message = msg
    	r.Num = n
    	r.ID = id
    }
    
    func BenchmarkReply(b *testing.B) {
    	var r Reply
    	for i := 0; i &lt; b.N; i++ {
    		r.Fill(200, "Success", 3394, "827348273")
    	}
    }
    //BenchmarkReply-8    1000000000           0.295 ns/op        0 B/op    0 allocs/op
    
    func BenchmarkReply2(b *testing.B) {
    	var r Reply
    	for i := 0; i &lt; b.N; i++ {
    		r.Code = 200
    		r.Message = "Success"
    		r.Num = 3394
    		r.ID = "827348273"
    	}
    }
    BenchmarkReply2-8    1000000000          0.558 ns/op        0 B/op      0 allocs/op
