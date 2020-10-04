# golang
## [1][Go, WebAssembly, HTTP requests and Promises - A guide to interacting with JS from Go/Wasm](https://www.reddit.com/r/golang/comments/j4lz94/go_webassembly_http_requests_and_promises_a_guide/)
- url: https://withblue.ink/2020/10/03/go-webassembly-http-requests-and-promises.html
---

## [2][NOEXCITE, a crypto social site built with golang](https://www.reddit.com/r/golang/comments/j4t948/noexcite_a_crypto_social_site_built_with_golang/)
- url: https://www.reddit.com/r/golang/comments/j4t948/noexcite_a_crypto_social_site_built_with_golang/
---
Hi all, my buddy and I built a crypto focused social media site with golang. The frontend was done with React JS and our backend was done with Go/MongoDB.  Check it out if you'd like and give me any feedback or questions you have.

https://www.noexcite.com/
## [3][Experimental Regex Engine](https://www.reddit.com/r/golang/comments/j4npuj/experimental_regex_engine/)
- url: https://www.reddit.com/r/golang/comments/j4npuj/experimental_regex_engine/
---
After listening to a talk by Bwk about successful language design, i decided to experiment with one of the points he made about the success of Awk. So [this package](https://github.com/kazhmir/ATA) is born.  


Everything it does is correlate a pattern with a given action, the pattern is a regular expression and the action is a function of a specific type. It creates a DFA from the regex and appends the function to the accepting state of every pattern, running it through the input linearly and executing whatever you asked it to. For example [here is a brainf\*ck interpreter](https://github.com/kazhmir/ATA/blob/master/docs/examples/bfinterpreter.go).  


Although it uses a very narrow subset of the PCRE standard, i think it can do quite a bit. I am planning on creating a idiotic interpreter for emoji-Assembly (it supports utf8 of course), i would love to see what you guys can come up with, if you're interested in playing with it :D
## [4][I made a proof-of-concept implementation of the Optional[T] type with the go2 generics preview](https://www.reddit.com/r/golang/comments/j4f6ib/i_made_a_proofofconcept_implementation_of_the/)
- url: https://go2goplay.golang.org/p/WyZQeG7OmWI
---

## [5][Thoughts on the ownership concept from Rust](https://www.reddit.com/r/golang/comments/j4w3pk/thoughts_on_the_ownership_concept_from_rust/)
- url: https://www.reddit.com/r/golang/comments/j4w3pk/thoughts_on_the_ownership_concept_from_rust/
---
Hi all, I've been using Go since 3 years now and I love the language. I started playing around with Rust and fell in love with the [ownership](https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html) concept of it. Now I want to know the pros and cons if Go implemented this concept and removed GC completely? Would/Wouldn't that be amazing?

I just want to hear your thoughts and learn something out of this.
## [6][I wanted to learn Golang so made a pastebin server and an on-disk key-value store for the data](https://www.reddit.com/r/golang/comments/j4fcti/i_wanted_to_learn_golang_so_made_a_pastebin/)
- url: https://www.reddit.com/r/golang/comments/j4fcti/i_wanted_to_learn_golang_so_made_a_pastebin/
---
https://github.com/AbhishekBagchi/pastebin-go
https://github.com/AbhishekBagchi/kvdb
## [7][Any good sample example on domain driven design in go](https://www.reddit.com/r/golang/comments/j4nhe9/any_good_sample_example_on_domain_driven_design/)
- url: https://www.reddit.com/r/golang/comments/j4nhe9/any_good_sample_example_on_domain_driven_design/
---
To learn
## [8][Real time problem on Windows](https://www.reddit.com/r/golang/comments/j4mj33/real_time_problem_on_windows/)
- url: https://www.reddit.com/r/golang/comments/j4mj33/real_time_problem_on_windows/
---

Hi guys,

I have a hobby project of an open-source multi-room audio player.
Basically, it streams a list of audio files to the network, and clients are playing the stream time-synchronized.

Everything goes well when both clients and servers are Linux-based. But, on Windows-based computers, the sound is completely chopped.

I suspect the program to don't have access to CPU in time to follow the stream bit rate.

Do you have to deal with similar problems before? Have you any tips to fix this?

I've found this SO question about program priority https://stackoverflow.com/a/4332
Has someone already used something similar?

For further reference here
- The link of the repo: https://github.com/tuarrep/sounddrop
- The link to the issue where I put my research on this: https://github.com/tuarrep/sounddrop/issues/15

Thanks a lot in advance for every input I could get from you.

Stay safe.
## [9][Contributing to GO2](https://www.reddit.com/r/golang/comments/j4ufwd/contributing_to_go2/)
- url: https://www.reddit.com/r/golang/comments/j4ufwd/contributing_to_go2/
---
I would like to share my thoughts and ideas for go2 (maybe 1 one them will be actually good, who knows 🤷).

Where is a good place to post it, where I can have a discussion about my (or others) thoughts and ideas, and maybe even that the developers see it (and hopefully join the discussion or even use it :D).

Thanks in advance for every help and suggestion :D.
## [10][Generate search query for GitHub in Go](https://www.reddit.com/r/golang/comments/j4k117/generate_search_query_for_github_in_go/)
- url: https://github.com/ynqa/github-search-query
---

