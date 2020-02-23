# golang
## [1][A non-obtrusive, idiomatic and easy-to-use query and aggregation builder for the official Go client for ElasticSearch.](https://www.reddit.com/r/golang/comments/f7r3ov/a_nonobtrusive_idiomatic_and_easytouse_query_and/)
- url: https://github.com/aquasecurity/esquery
---

## [2][After a month of waiting this guy joined me at work yesterday. Ordered from the Golang Market!](https://www.reddit.com/r/golang/comments/f7c13v/after_a_month_of_waiting_this_guy_joined_me_at/)
- url: https://i.redd.it/qedt8dnfhai41.jpg
---

## [3][Ever dream using WebRTC MediaDevices API in Go?](https://www.reddit.com/r/golang/comments/f7ols0/ever_dream_using_webrtc_mediadevices_api_in_go/)
- url: https://www.reddit.com/r/golang/comments/f7ols0/ever_dream_using_webrtc_mediadevices_api_in_go/
---
Ever dream using WebRTC MediaDevices API in Go? Dream no more, [https://github.com/pion/mediadevices](https://github.com/pion/mediadevices) is what you were looking for! It is easy-to-use, flexible, &amp; extensible.

We support 🎥, 🎙️, &amp; screen-share for Linux + codecs. We're expanding &amp; would ❤️ for you to join us, use, contribute!!

[https://twitter.com/lherman\_cs/status/1231095038911733761](https://twitter.com/lherman_cs/status/1231095038911733761)

&amp;#x200B;

[demo](https://i.redd.it/scsp2bqx1fi41.gif)
## [4][GitHub - inlets/cloud-provision-example: Provision and automate cloud hosts with Go](https://www.reddit.com/r/golang/comments/f7rbh1/github_inletscloudprovisionexample_provision_and/)
- url: https://github.com/inlets/cloud-provision-example
---

## [5][parallel: a Go parallel processing library](https://www.reddit.com/r/golang/comments/f7skxy/parallel_a_go_parallel_processing_library/)
- url: https://github.com/ryanskidmore/parallel
---

## [6][Go lessons learnt by refactoring](https://www.reddit.com/r/golang/comments/f7sirl/go_lessons_learnt_by_refactoring/)
- url: https://anto.pt/post/go-lessons-learnt-by-refactoring/
---

## [7][Your first HTTP Server in Go - Go Web Basics #1](https://www.reddit.com/r/golang/comments/f7shad/your_first_http_server_in_go_go_web_basics_1/)
- url: https://www.youtube.com/watch?v=5BIylxkudaE
---

## [8][I have published Goxygen - a tool that helps generate full stack projects in Go and React. I would love to hear your feedback](https://www.reddit.com/r/golang/comments/f7aco4/i_have_published_goxygen_a_tool_that_helps/)
- url: https://github.com/Shpota/goxygen
---

## [9][Nemesis is a tool for auditing platform configurations for measuring compliance.](https://www.reddit.com/r/golang/comments/f7r30z/nemesis_is_a_tool_for_auditing_platform/)
- url: https://github.com/UnityTech/nemesis
---

## [10][Why does programming with side effects seem to be so prevalent among Go code?](https://www.reddit.com/r/golang/comments/f7i3cz/why_does_programming_with_side_effects_seem_to_be/)
- url: https://www.reddit.com/r/golang/comments/f7i3cz/why_does_programming_with_side_effects_seem_to_be/
---
First, let me get things straight. I've been programming Go long enough to understand why people would pass a pointer as an argument to a function, versus passing, say a struct. Passing by value it is. 

What I don't quite understand is if there is anything wrong with passing a struct by value, modifying the copy inside the function and returning the copy that then gets reassigned to the original variable. I'm talking about side effect free, immutable functions. Coming from other languages, this seems like a no brainer, but apparently, the preferred practice in Go is to pass and modify a pointer instead. Every piece of popular code I've seen, including the language spec itself does that. 

It cannot be simply about copying performance. I have written tens of various benchmarks on this topic, and in none of them did passing a pointer beat significantly passing an entire struct (regardless of its size).

So what is it?
