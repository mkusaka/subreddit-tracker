# golang
## [1][100+ LeetCode solutions in Go](https://www.reddit.com/r/golang/comments/fqu0nf/100_leetcode_solutions_in_go/)
- url: https://www.reddit.com/r/golang/comments/fqu0nf/100_leetcode_solutions_in_go/
---
I'd like to share a set of 100+ solutions to LeetCode problems that are all written in Go. My hope is that these solutions can serve as a helpful reference for others, since there is not always a solution available that's written in Go.

You can check out the solutions at: [https://github.com/austingebauer/go-leetcode](https://github.com/austingebauer/go-leetcode)

I hope this will be a helpful resource to someone :) If you find it useful, a project star is much appreciated.
## [2][Service for markdown sharing in Go](https://www.reddit.com/r/golang/comments/fr5jpl/service_for_markdown_sharing_in_go/)
- url: https://www.reddit.com/r/golang/comments/fr5jpl/service_for_markdown_sharing_in_go/
---
I've written simple service for markdown sharing in Go.
It renders raw markdown to pages and supports CommonMark markdown with additional extensions like tweet embedding.

Further I'm planning to add posts editing support, enhance styles and add minor frontend features, add more extensions to markdown, but main feature is keeping clean and simple.

Comments on service itself, code review, or notes about project architecture will be *very* appreciated!

Code on GitHub:
https://github.com/vdimir/markify

Deployed application:
https://markify.dev/
## [3][formatting library for Node-like (recursive) data structures](https://www.reddit.com/r/golang/comments/fr5606/formatting_library_for_nodelike_recursive_data/)
- url: https://www.reddit.com/r/golang/comments/fr5606/formatting_library_for_nodelike_recursive_data/
---
[https://github.com/roz3x/formatter](https://github.com/roz3x/formatter)

not a common use-case but , i needed something like that. a comment would be great . 👍
## [4][How To Separate Library Packages in Go](https://www.reddit.com/r/golang/comments/fr5r1k/how_to_separate_library_packages_in_go/)
- url: https://qvault.io/2020/03/29/how-to-separate-library-packages-in-go/
---

## [5][GoCache - Cache Server in Go](https://www.reddit.com/r/golang/comments/fqjrat/gocache_cache_server_in_go/)
- url: https://github.com/kadnan/gocache
---

## [6][Does this mean everybody will have to create retry loops around every system call?](https://www.reddit.com/r/golang/comments/fqv036/does_this_mean_everybody_will_have_to_create/)
- url: https://www.reddit.com/r/golang/comments/fqv036/does_this_mean_everybody_will_have_to_create/
---
[Goroutines are asynchronously preemptible](https://golang.org/doc/go1.14#runtime)

[https://golang.org/doc/go1.14#runtime](https://golang.org/doc/go1.14#runtime)
## [7][Best approach to pass flag values to a function while also allowing the function to be called independently ?](https://www.reddit.com/r/golang/comments/fr1xok/best_approach_to_pass_flag_values_to_a_function/)
- url: https://www.reddit.com/r/golang/comments/fr1xok/best_approach_to_pass_flag_values_to_a_function/
---
Basically i am building a cli application. 

I want to pass the user given values that are given to flags to the function for further processing on it. But i also want the function to be called independently by any other library without the cli so what would be the best approach.

Below are couple of approaches i have thought of.

A . Have everything from the flag sent to a variable using the Flag.StringVar and then Pass the variables to the function. 

B. Create a Struct,use the Flag.StringVar(&amp;var) thing and pass the struct to the function. 

C. Pass the Whole Flags() thing to the function for further processing. 

Rn i am thinking B would be the best as i need some default values to be there aswell as the output location if not speciied by the user should default to the current working library.

So basically my flag looks like this     

    imageCmd.Flags().StringVarP(&amp;inloc, "in", "i", "", " image location (stdin by default)") 

    imageCmd.Flags().StringVarP(&amp;outloc, "out", "o", "", "where to put the image ")

    imageCmd.Flags().BoolVarP(&amp;meta, "metad", "m", false, "image metdata")

    func imagelib(inloc string, outloc string, meta string) // processes the image

Now the outloc is optional, so i want it to default to nil but given that optional parameter passing to function is not possible in golang i think rn the struct option would be the best. Since i also want to  acess the function from another lib which means flags wont be used so using struct would mean i can simply pass a struct there . 

    e := mylib.decode(parameters)
## [8][A Guide To Writing Logging Middleware in Go](https://www.reddit.com/r/golang/comments/fqqkwb/a_guide_to_writing_logging_middleware_in_go/)
- url: https://blog.questionable.services/article/guide-logging-middleware-go/
---

## [9][go get -u all: "cannot find package"](https://www.reddit.com/r/golang/comments/fr37qx/go_get_u_all_cannot_find_package/)
- url: https://www.reddit.com/r/golang/comments/fr37qx/go_get_u_all_cannot_find_package/
---
Good day. I installed packages with normal `go get`.

`go version go1.13.8 linux/386`

Full log:

    package github.com/cncf/udpa/test/build: C++ source files not allowed when not using cgo or SWIG: build_test.cc
    package github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go: cannot find package "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go" in any of:
    	/usr/lib/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go (from $GOROOT)
    	/home/vitaly/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go (from $GOPATH)
    package github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go: cannot find package "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go" in any of:
    	/usr/lib/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go (from $GOROOT)
    	/home/vitaly/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go (from $GOPATH)
    package github.com/envoyproxy/protoc-gen-validate/tests/harness/go: no Go files in /home/vitaly/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/go
    package github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/gogo: cannot find package "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/gogo" in any of:
    	/usr/lib/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/gogo (from $GOROOT)
    	/home/vitaly/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/gogo (from $GOPATH)
    package github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/gogo: cannot find package "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/gogo" in any of:
    	/usr/lib/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/gogo (from $GOROOT)
    	/home/vitaly/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/gogo (from $GOPATH)
    package github.com/envoyproxy/protoc-gen-validate/tests/harness/gogo: no Go files in /home/vitaly/go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/gogo
    package github.com/google/go-github/v28/github: cannot find package "github.com/google/go-github/v28/github" in any of:
    	/usr/lib/go/src/github.com/google/go-github/v28/github (from $GOROOT)
    	/home/vitaly/go/src/github.com/google/go-github/v28/github (from $GOPATH)
    # cd /home/vitaly/go/src/github.com/rwcarlsen/goexif; git pull --ff-only
    You are not currently on a branch.
    Please specify which branch you want to merge with.
    See git-pull(1) for details.
    
        git pull &lt;remote&gt; &lt;branch&gt;
    
    package github.com/rwcarlsen/goexif/exif: exit status 1
## [10][Got nil from the channel, which is not possible by code.](https://www.reddit.com/r/golang/comments/fqurch/got_nil_from_the_channel_which_is_not_possible_by/)
- url: https://www.reddit.com/r/golang/comments/fqurch/got_nil_from_the_channel_which_is_not_possible_by/
---
Hi folks - we faced the strange behavior inside our system.  
[https://play.golang.org/p/P3HdG2jNvz7](https://play.golang.org/p/P3HdG2jNvz7)  
we've got a panic line:36  
As for me, that's not possible, can you help?  
maybe a possibility exists?  


that is how it's implemented in the system without the rest of the code.  
I mean, I extracted the main idea of how it's implemented, without details.
