# golang
## [1][[Q&amp;A] //go:build draft design](https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/
---
I posted a draft design today for updating the // +build lines to fix some usability problems. 

Video: [https://golang.org/s/go-build-video](https://golang.org/s/go-build-video)\
Design: [https://golang.org/s/go-build-design](https://golang.org/s/go-build-design)

As an experiment, let's try doing Q&amp;A about the design here in Reddit.
My hope is that the threading support will help keep questions and answers matched.

**Please start a new top-level comment for each new question.**
## [2][Size doesn't matter](https://www.reddit.com/r/golang/comments/hjuyfv/size_doesnt_matter/)
- url: https://i.redd.it/m5tjssbydf851.jpg
---

## [3][Go is Boring...And That’s Fantastic!](https://www.reddit.com/r/golang/comments/hjqouw/go_is_boringand_thats_fantastic/)
- url: https://www.capitalone.com/tech/software-engineering/go-is-boring/
---

## [4][Good example web applications to learn from? (production-ready, best practice, tech stack)](https://www.reddit.com/r/golang/comments/hjj0l6/good_example_web_applications_to_learn_from/)
- url: https://www.reddit.com/r/golang/comments/hjj0l6/good_example_web_applications_to_learn_from/
---
I'm learning Golang, and right now I'm looking for a good opensource (full-stack) web application that employs best practice and a well-chosen tech stack.

When I got started with Node.js and GraphQL, there was a project called [Spectrum](https://github.com/withspectrum/spectrum), which was tremendously helpful as a reference project. Does the Go community have some projects similar (not necessarily a chat room like Spectrum)?
## [5][Compare chromedp with rod with real examples](https://www.reddit.com/r/golang/comments/hju2j0/compare_chromedp_with_rod_with_real_examples/)
- url: https://github.com/go-rod/rod/tree/master/lib/examples/compare-chromedp
---

## [6][Surge — A fast and efficient binary marshaler for Byzantine networks](https://www.reddit.com/r/golang/comments/hjs0o9/surge_a_fast_and_efficient_binary_marshaler_for/)
- url: https://github.com/renproject/surge
---

## [7][Qvault Classroom Launches Golang Crash Course](https://www.reddit.com/r/golang/comments/hjw7ji/qvault_classroom_launches_golang_crash_course/)
- url: https://qvault.io/2020/07/02/qvault-classroom-launches-golang-crash-course/
---

## [8][Release gopls/v0.4.2 · golang/tools · GitHub](https://www.reddit.com/r/golang/comments/hjbyy4/release_goplsv042_golangtools_github/)
- url: https://github.com/golang/tools/releases/tag/gopls%2Fv0.4.2
---

## [9][Debugging an evil Go runtime bug - marcan.st](https://www.reddit.com/r/golang/comments/hjl0lf/debugging_an_evil_go_runtime_bug_marcanst/)
- url: https://marcan.st/2017/12/debugging-an-evil-go-runtime-bug/
---

## [10][Looking for a library recommendation for a P2P publish/subscribe system](https://www.reddit.com/r/golang/comments/hjumz6/looking_for_a_library_recommendation_for_a_p2p/)
- url: https://www.reddit.com/r/golang/comments/hjumz6/looking_for_a_library_recommendation_for_a_p2p/
---
Dear fellow gophers! I'm looking for recommendations and experiences with using P2P libraries, because I'm really unsure which one to use. I need publishing of data on a topic and subscribing as well as sending messages to individual users via some handle or public key that represents them. It's for a non-open-source hobby project that may later have a potentially large user base. It does not involve any file sharing or low-latency connections, only small data transfers. It should ideally support NAT traversal or at least require only few ports to be opened. The P2P network needs to be able to start by itself, using existing bootstrap nodes.

Currently, I'm considering:

* LibP2P: [https://github.com/libp2p](https://github.com/libp2p)
* Noise: [https://github.com/perlin-network/noise](https://github.com/perlin-network/noise)

LibP2P seems to be the safe choice, it's actively maintained. But it is not well documented and seems  cumbersome to use. The examples seem to all use local peer discovery only and it's not obvious to me how to use the DHT component for bootstrapping the network. Noise, on the other hand, appears to be very easy to use. The problem is just that the author doesn't seem to update it very actively right now. 

Has anyone used the above libraries for non-trivial purposes? Are there other P2P libraries for Go?
## [11][git-get - a better way to clone and organize git repos (inspired by go get)](https://www.reddit.com/r/golang/comments/hjx2l9/gitget_a_better_way_to_clone_and_organize_git/)
- url: https://github.com/grdl/git-get
---

