# golang
## [1][bubbletea: A fun, functional and stateful framework to build terminal apps 🏗](https://www.reddit.com/r/golang/comments/j8rk2i/bubbletea_a_fun_functional_and_stateful_framework/)
- url: https://github.com/charmbracelet/bubbletea
---

## [2][Managing Microservice Schema and Interfaces in Distributed Environments With Protocol Buffers - [Part - I] Feel free to comments your thoughts on this. All examples used in the blog is written in golang](https://www.reddit.com/r/golang/comments/j9095x/managing_microservice_schema_and_interfaces_in/)
- url: https://ednsquare.com/story/managing-microservice-schema-and-interfaces-in-distributed-environments-with-protocol-buffers-part-i-------AkFVJP
---

## [3][statsviz: instant live visualization of your Go application runtime statistics (GC, MemStats, etc.) in the browser](https://www.reddit.com/r/golang/comments/j8u1gk/statsviz_instant_live_visualization_of_your_go/)
- url: https://github.com/arl/statsviz
---

## [4][what do you do with golang in your company?](https://www.reddit.com/r/golang/comments/j91b3v/what_do_you_do_with_golang_in_your_company/)
- url: https://www.reddit.com/r/golang/comments/j91b3v/what_do_you_do_with_golang_in_your_company/
---
I'm newbie to golang and I wonder what kind of project you do with golang. 

I think it'll help motivate to learn golang

can you guys tell me?
## [5][kubecolor: a CLI tool to colorize kubectl output written in Go](https://www.reddit.com/r/golang/comments/j8khzf/kubecolor_a_cli_tool_to_colorize_kubectl_output/)
- url: https://www.reddit.com/r/golang/comments/j8khzf/kubecolor_a_cli_tool_to_colorize_kubectl_output/
---
[https://github.com/dty1er/kubecolor](https://github.com/dty1er/kubecolor)
## [6][Looking for a sample GO script showing how to stream a few mp4s placed in a directory.](https://www.reddit.com/r/golang/comments/j92sll/looking_for_a_sample_go_script_showing_how_to/)
- url: https://www.reddit.com/r/golang/comments/j92sll/looking_for_a_sample_go_script_showing_how_to/
---
Can anybody recommend a sample GO script showing how to stream a few mp4s placed in a directory or similar.
## [7][Experimenting with Golang bufio Reader vs Scanner to understand the difference](https://www.reddit.com/r/golang/comments/j93hpg/experimenting_with_golang_bufio_reader_vs_scanner/)
- url: https://shareablecode.com/snippets/difference-between-golang-bufio-reader-vs-scanner-newreader-vs-newscanner-f8eP-uc1U
---

## [8][Best material to learn golang](https://www.reddit.com/r/golang/comments/j938e6/best_material_to_learn_golang/)
- url: https://www.reddit.com/r/golang/comments/j938e6/best_material_to_learn_golang/
---
what are good udemy or youtube resources available to learn golang, apart from official doc. I have couple of yrs of experience in java and kinda curious about golang, wants to get more insight into it.
## [9][go modules making me rage! How do I fork a module?](https://www.reddit.com/r/golang/comments/j8pqms/go_modules_making_me_rage_how_do_i_fork_a_module/)
- url: https://www.reddit.com/r/golang/comments/j8pqms/go_modules_making_me_rage_how_do_i_fork_a_module/
---
What I'm trying to do, I believe, is extremely standard, yet go modules are causing me to regret every decision.

Here's my app, fooApp. Various 3rd party imports such as "import github.com/authorA/mygolibrary/v3". I did `go mod init` followed by `go mod tidy` and I was able to get a solid build with `go build`, finding all the dependencies and such.

Now, I want to fork authorA's library so that I can make some fixes to eventually submit pull-requests. So, I went to github and forked it to "github.com/mememeB/mygolibrary". I then changed my import lines accordingly "import github.com/mememeB/mygolibrary/v3"

`go build` then fetches my library, but complains that mygolibrary is defined as authorA/mygolibrary but imported as mememeB/mygolibrary

Ok. So, some googling says I need to add a "replace" line in my go.mod file.

`replace github.com/authorA/mygolibrary/v3 =&gt; github.com/mememeB/mygolibrary/v3`

New error. "replacement module without version must be directory path" What? More searching.

`replace github.com/authorA/mygolibrary/v3 =&gt; github.com/mememeB/mygolibrary/v3 v0`

More complaints that mememeB's go.mod file defines the fork as authorA.

Must I really change everything (go.mod, all imports) in my fork to match my "url"? That makes submitting PR's insane because now there's tons more differences than originally needed.

What's the proper way to fork a module and use it in your apps?
## [10][Go might support converting a slice []T to an array pointer *[N]T in Go 1.16](https://www.reddit.com/r/golang/comments/j8lbav/go_might_support_converting_a_slice_t_to_an_array/)
- url: https://github.com/golang/go/issues/395
---

