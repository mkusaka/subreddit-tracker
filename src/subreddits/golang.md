# golang
## [1][GoLand 2020.1 Reaches Beta with a lot of Go Modules, Completion, Editing, Performance improvements and more!](https://www.reddit.com/r/golang/comments/fpfim3/goland_20201_reaches_beta_with_a_lot_of_go/)
- url: https://blog.jetbrains.com/go/2020/03/26/goland-2020-1-reaches-beta/
---

## [2][Cron job library in GO](https://www.reddit.com/r/golang/comments/fps32l/cron_job_library_in_go/)
- url: https://www.reddit.com/r/golang/comments/fps32l/cron_job_library_in_go/
---
Hi gophers,

[https://github.com/Fred07/gron](https://github.com/Fred07/gron)

This lib forked from [https://github.com/roylee0704/gron](https://github.com/roylee0704/gron), which is not updating since 4 years ago.

I forked it and add some new feature that old one didn't have, like \`GracefullyStop()\`, \`StartAndServe()\`,  handling OS signal interrupt.

check it out and welcome to leave some advice to me :)
## [3][Code Generation with Go and Cuelang](https://www.reddit.com/r/golang/comments/fpt2yy/code_generation_with_go_and_cuelang/)
- url: https://www.reddit.com/r/golang/comments/fpt2yy/code_generation_with_go_and_cuelang/
---
Brand new implementation of our code generation project using CUE here: [https://github.com/hofstadter-io/hof](https://github.com/hofstadter-io/hof) Supports  


* Writing generator specs and "designs" in Cuelang
* Code generation from Cuelang (follow the projects cue files and imports (vendor is cue.mod/pkg))
* Cleanup of any generated code that has been deleted in the design
* Regeneration without losing custom code, write right in the output

Hof uses itself to generate it's CLI layout using [https://github.com/hofstadter-io/cuemod--cli-golang](https://github.com/hofstadter-io/cuemod--cli-golang) \[The only supported generator today, just wait for tomorrow ;\]  
Docs are a bit light right now, it will generate any values that start with `Gen.*` but expects the format seen in `cli-main.cue` and the imported `cuemod--cli-golang/module.cue`. The main engine starts in `hof/lib/gen.go` and is just over 1000LOC.
## [4][PokeTraveler - an old-school Pokémon world in Go (Ebiten)](https://www.reddit.com/r/golang/comments/fpfett/poketraveler_an_oldschool_pokémon_world_in_go/)
- url: https://github.com/Akatsuki-py/PokeTraveler
---

## [5][Question about GOPATH and directory structure with multiple projects/repositories](https://www.reddit.com/r/golang/comments/fplej3/question_about_gopath_and_directory_structure/)
- url: https://www.reddit.com/r/golang/comments/fplej3/question_about_gopath_and_directory_structure/
---
Hi, new to Go here and my Google skills are coming up short.  What is best/preferred practice for directory structure of Go code in the following use cases:

1) A project that has source files of many types, such as Go, CloudFormation YAML, and Ansible scripts?

2) Multiple separate and completely independent entirely Go programs?

The reason I ask is because everything I read indicates a single root directory, say "src" for ALL Go work.  Then sub-directories not at a project level, but based on function.  So for ex.

`/src/cmd/application1`

`/src/cmd/application2`

`/src/pkg/reqforapp1`

`/src/pkg/reqforapp2`

and so forth.  This would make source control repositories by project chaos, and does not work for either use case I mentioned.  I am assuming my naivety with Go is the source of my confusion, looking for someone who has solved this to correct me.  

I guess my basic issue is that everything I read assumes a single GOPATH value across an entire system, and I just don't see how that can work.  If I have a structure of:

`/project1/servicefunctions`

`/project1/service/functions/gosrc`

`/project1/cloudformation`

`/project1/ansible`

I would expect that while I work on that project, my GOPATH = /project1/service/functions/gosrc

If it matters, I'm using GoLand for the IDE.  Any tips are appreciated, critique tolerated.
## [6][Quote : Displays random quotes from collection of quotes.](https://www.reddit.com/r/golang/comments/fprlce/quote_displays_random_quotes_from_collection_of/)
- url: https://www.reddit.com/r/golang/comments/fprlce/quote_displays_random_quotes_from_collection_of/
---
&amp;#x200B;

[Quote](https://preview.redd.it/295zbrkdl5p41.png?width=832&amp;format=png&amp;auto=webp&amp;s=c9cd8450460bd9a6167c206180d5c815032f258a)

I am currently learning Go and made this simple program for fun. Link to Github repository is provide below. Please provide suggestion and new ideas to improve it. Feel free to download, make changes and contribute to it.

[https://github.com/aryanmaurya1/quote](https://github.com/aryanmaurya1/quote)
## [7][I created a linter for organizations who want to standardize on the modules used and be able to recommend alternative modules. Would like some feedback if anyone was interested.](https://www.reddit.com/r/golang/comments/fpo32e/i_created_a_linter_for_organizations_who_want_to/)
- url: https://github.com/ryancurrah/gomodguard
---

## [8][Golang library for fbprophet](https://www.reddit.com/r/golang/comments/fpthh9/golang_library_for_fbprophet/)
- url: https://www.reddit.com/r/golang/comments/fpthh9/golang_library_for_fbprophet/
---
Does anyone know of a native Go library that is equivalent to Facebook's fbprophet?
## [9][How to do testing with zap a popular logging library for Golang](https://www.reddit.com/r/golang/comments/fpsmf0/how_to_do_testing_with_zap_a_popular_logging/)
- url: https://gianarb.it/blog/golang-mockmania-zap-logger
---

## [10][How much performance headroom is left / longer compile times](https://www.reddit.com/r/golang/comments/fp9mrk/how_much_performance_headroom_is_left_longer/)
- url: https://www.reddit.com/r/golang/comments/fp9mrk/how_much_performance_headroom_is_left_longer/
---
Hi all – As you know, the Go compiler is incredibly fast, and generated application code is decently fast. How much do you think application performance could be improved with longer compile times?

Say there was an option for longer, deeper compiles, for contexts where instant compiles were not needed. For example, for the release build of an application I'd want the best possible runtime performance – I'd be fine letting a compiler run over a weekend. Do you think there is still a lot of application performance headroom to be realized with longer build times?

The compiler is somewhat threadbare when it comes to optimization options. It lacks the suite of optimizations that gcc, clang, Visual Studio, and the Intel compiler offer for C/C++ applications. Where I see possible headroom is in targeting modern CPU instructions and vectorization. The last time I looked at the compiler code base I saw a bunch of old assembly targeting old CPUs – is it still not accounting for Haswell+ CPUs? Go support for modern assembly *by the application developer* is excellent, but I'm not sure if the Go compiler itself uses modern instructions in its generated code. Thoughts?

**Edit:** I mentioned vectorization, but I'm mostly thinking about things like Whole Program Optimization / Link Time Optimization (LTO), Profile Guided Optimization (PGO), and superoptimization. For the latter, see Stanford's [STOKE superoptimizing compiler](http://stoke.stanford.edu/) for the x86-64 ISA.

Thanks.
