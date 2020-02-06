# golang
## [1][Yet another O'RLY cover generator, in Go and Vue](https://www.reddit.com/r/golang/comments/ezqno6/yet_another_orly_cover_generator_in_go_and_vue/)
- url: https://www.reddit.com/r/golang/comments/ezqno6/yet_another_orly_cover_generator_in_go_and_vue/
---
Site: https://orly.nanmu.me/

Github: https://github.com/nanmu42/orly

O'RLY Cover Generator is a parody book cover generator, implemented in Golang and Vue.js, supporting a wide range of language including CJK.
## [2][Go 1.14 Release Candidate 1 is released](https://www.reddit.com/r/golang/comments/ezfags/go_114_release_candidate_1_is_released/)
- url: https://groups.google.com/d/msg/golang-announce/mB1Mp9RlQw8/wf1OVYfTAwAJ
---

## [3][VSCode problem matcher for 'go test'](https://www.reddit.com/r/golang/comments/ezsb7i/vscode_problem_matcher_for_go_test/)
- url: https://www.reddit.com/r/golang/comments/ezsb7i/vscode_problem_matcher_for_go_test/
---
Does anyone have a good setup for [VSCode's problem matchers](https://code.visualstudio.com/docs/editor/tasks#_defining-a-problem-matcher) that will capture failures from `go test`?

I've been playing with [multiline problem matchers](https://code.visualstudio.com/docs/editor/tasks#_defining-a-multiline-problem-matcher) but can't get it to pick up the correct relative path to the `*_test.go` file that reported the test failure, such that when I'm in the Problems tab and click on the reported failure, it can't find the file, because it's looking for everything in the workspace root.
## [4][godoc.org will be closed owing to legal reason](https://www.reddit.com/r/golang/comments/ez7m26/godocorg_will_be_closed_owing_to_legal_reason/)
- url: https://groups.google.com/d/msg/golang-dev/mfiPCtJ1BGU/qtCrqlrEEwAJ
---

## [5][Struggle to get nice how to for elemental module based project template with /cmd](https://www.reddit.com/r/golang/comments/ezq609/struggle_to_get_nice_how_to_for_elemental_module/)
- url: https://www.reddit.com/r/golang/comments/ezq609/struggle_to_get_nice_how_to_for_elemental_module/
---
I went through few monthes of go. Love it and start rewritting some of my projects. Started to use modules but still in very basic layout where porject gopath is root. Then I started to need to have separate runables. So i understood /cmd + /pkg or /internal is the way but cant find easi guide how to make this happen. In particular
1) should cmd and pkg contain its own go.mod?
2) how to run right go run /cmd/client/x.go so it sees packages from pkg
3) how to do it nicely in GoLand?
I bet more people would appreciate some small how to...
Yes i went through different talks and medium post but none of them solve particular this..
## [6][Why does http.ListenAndServeTLS take filenames for the cert files instead of either Readers or byte arrays?](https://www.reddit.com/r/golang/comments/eze681/why_does_httplistenandservetls_take_filenames_for/)
- url: https://www.reddit.com/r/golang/comments/eze681/why_does_httplistenandservetls_take_filenames_for/
---
This seems to go against the principles of design in Go, so I expect there's probably a good reason that I just don't understand. Few parts of the stdlib couple implementations to the filesystem like this.
## [7][Image Rendering in 2D Video Games with Ebiten](https://www.reddit.com/r/golang/comments/ezake8/image_rendering_in_2d_video_games_with_ebiten/)
- url: https://medium.com/a-journey-with-go/go-image-rendering-in-2d-video-games-with-ebiten-912cc2360c4f
---

## [8][Do you agree with this Comparative Analysis of Go vs Rust?](https://www.reddit.com/r/golang/comments/ezr43m/do_you_agree_with_this_comparative_analysis_of_go/)
- url: https://www.brsoftech.com/blog/rust-vs-go-which-is-better-programming-language-for-future/
---

## [9][Why Discord is switching from Go to Rust](https://www.reddit.com/r/golang/comments/eywx4q/why_discord_is_switching_from_go_to_rust/)
- url: https://blog.discordapp.com/why-discord-is-switching-from-go-to-rust-a190bbca2b1f
---

## [10][Ideas for versioning multiple binaries under "/cmd" directory](https://www.reddit.com/r/golang/comments/ezlpx7/ideas_for_versioning_multiple_binaries_under_cmd/)
- url: https://www.reddit.com/r/golang/comments/ezlpx7/ideas_for_versioning_multiple_binaries_under_cmd/
---
Hi all,

I have a question regarding versioning when there are multiple binaries in a single project/repo. 

Let's say I have a project call *MyApp* which contains multiple executable binaries and packages. There is a global version for the app.

However, I also want to track changes to individual binaries (and perhaps packages too if possible). Following go best practices, I have a "cmd" directory with sub-directories for each executable binary (ExecA, ExecB etc.).

When I run *ExecA* I would like to print out version info something like:

app version: v1.2.1
execA version: v2.0.0

The relationship between the app version and the executable versions is something I am still working on. The current placeholder system is if an executable's Major version is incremented, then I increment the master's minor version. Kind of silly, I know.

I have been using *git tags* as the global app version for the entire project, and a simple VERSION file for the executables. My Makefile combines these and injects them into the binary before building.

If anyone has any ideas, suggestions or links to informative articles, I'd be grateful. I can provide a few more details too if necessary.

Thanks for reading!
