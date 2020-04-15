# golang
## [1][Build Your Own Neural Network in Go](https://www.reddit.com/r/golang/comments/g1hkuh/build_your_own_neural_network_in_go/)
- url: https://towardsdatascience.com/neural-network-from-scratch-in-go-language-b98e2abcced3
---

## [2][Go implementation of Lisp (goroutine/channel in Lisp)](https://www.reddit.com/r/golang/comments/g1gtmd/go_implementation_of_lisp_goroutinechannel_in_lisp/)
- url: https://github.com/mattn/golisp
---

## [3][Testing Go Better How To Create Testable Go Code](https://www.reddit.com/r/golang/comments/g1mcba/testing_go_better_how_to_create_testable_go_code/)
- url: https://engineering.kablamo.com.au/posts/2020/testing-go
---

## [4][Is there a framework agnostic (non-sql) migration library](https://www.reddit.com/r/golang/comments/g1rkrw/is_there_a_framework_agnostic_nonsql_migration/)
- url: https://www.reddit.com/r/golang/comments/g1rkrw/is_there_a_framework_agnostic_nonsql_migration/
---
Hi everyone!

I have been building some Node.js apps but am transitioning to Go. For my Node applications I preferred to use [https://github.com/sequelize/umzug](https://github.com/sequelize/umzug) because it enabled me to not only write SQL migrations but also generic migrations that would (for example) call a web API or create some files in the filesystem.

Have you seen such a thing written as Go library?

I had a look at [https://github.com/golang-migrate/migrate#use-in-your-go-project](https://github.com/golang-migrate/migrate#use-in-your-go-project) but that is too SQL specific.

This lists only SQL frameworks: [https://awesome-go.com/](https://awesome-go.com/)

I would actually be interested in creating such a library on my own but as my project is related to the company I work at, there is no budget to do so right now.
## [5][TDD testing framework, with molecularity in mind](https://www.reddit.com/r/golang/comments/g1rhj2/tdd_testing_framework_with_molecularity_in_mind/)
- url: https://github.com/adamluzsi/testcase
---

## [6][🥦Broccoli: We wrote the most efficient static file embedding tool for Go, benefiting from Google's brotli compression.](https://www.reddit.com/r/golang/comments/g12mfh/broccoli_we_wrote_the_most_efficient_static_file/)
- url: https://www.reddit.com/r/golang/comments/g12mfh/broccoli_we_wrote_the_most_efficient_static_file/
---
Hello reddit,

Without the further ado, 🥦[Broccoli](https://github.com/aletheia-icu/broccoli) is the tool we have developed over the course of the last 1-2 weeks. I usually go for go-bindata, or sometimes for parchello, but when I had to embed files into a .wasm binary (WebAssembly target, wasm/js) most of the existing solutions I picked up just didn't work due to some obscure HTTP panic. I told me lads about this (I do mentoring) and the idea to use brotli came up, as well as some other features, unseen in the existing generators.

- The average is 13-25% smaller bundle size due to use of superior compression algorithm. [benchmarks](https://vcs.aletheia.icu/lads/broccoli-bench)
-  Broccoli supports bundling of multiple source directories, only relies on go generate command-line interface and doesn't require configuration files.
- Optional decompression is something you may want; when it's enabled, files are decompressed only when they are read the first time.
- Broccoli can be used as an [http.FileSystem](https://golang.org/pkg/net/http/#FileSystem)
- It's one of the few (including statik) libraries to work well on `wasm/js`!
- There is `-gitignore` option to ignore files, already ignored by your existing .gitignore files.

I posted a link to our repository yesterday and it attracted very unusually high number of dislikes and virtually no comments, so I created this post instead and wish to invite those who feel disdain, to communicate it, so we can have a discussion about the quality of a proposed solution.

-Ian
## [7][Who are some Github residents to follows to discover new exciting projects?](https://www.reddit.com/r/golang/comments/g1pdh8/who_are_some_github_residents_to_follows_to/)
- url: https://www.reddit.com/r/golang/comments/g1pdh8/who_are_some_github_residents_to_follows_to/
---
Title says it all, but to expand:

I don't know how Twitter works (yes I'm one of those) and I'm not on Reddit often, so I usually miss new *exciting* projects like [Twirp RPC](https://github.com/twitchtv/twirp) and the more recent [GoatCounter web analytics platform](https://github.com/zgoat/goatcounter) and [Broccoli file embedder](https://github.com/aletheia-icu/broccoli).

My next favourite source is Github followers because when I get some time to go on there I get a list starred projects to look at, but unfortunately in recent times, that list gets smaller and smaller...

So... who should I *follow* these days? I don't really want to subscribe to any of the *awesome* lists because they're often incomplete and sometimes contain dead projects.
## [8][a job/task queue framework for Go](https://www.reddit.com/r/golang/comments/g1o7vy/a_jobtask_queue_framework_for_go/)
- url: https://www.reddit.com/r/golang/comments/g1o7vy/a_jobtask_queue_framework_for_go/
---
Asynchronous tasks are becoming bigger and bigger, so what can we do? we can use gotasks.

In gotasks, we encourage developer to split tasks into smaller pieces(see the demo bellow) so we can:

\- maintain tasks easily

\- split code into reentrant and un-reentrant pieces, so when reentrant part failed, framework will retry it automatically

see more: [https://github.com/jiajunhuang/gotasks](https://github.com/jiajunhuang/gotasks)
## [9][Containers From Scratch in a Few Lines of Go](https://www.reddit.com/r/golang/comments/g154rq/containers_from_scratch_in_a_few_lines_of_go/)
- url: https://youtu.be/8fi7uSYlOdc
---

## [10][Running Go programs as unikernels - included demo is a functional Gio GUI program](https://www.reddit.com/r/golang/comments/g12jl3/running_go_programs_as_unikernels_included_demo/)
- url: https://www.reddit.com/r/golang/comments/g12jl3/running_go_programs_as_unikernels_included_demo/
---
See the announcement [https://groups.google.com/forum/#!topic/golang-nuts/4cDIL5Vr\_es](https://groups.google.com/forum/#!topic/golang-nuts/4cDIL5Vr_es)

The repo is at   [https://eliasnaur.com/unik](https://eliasnaur.com/unik)
