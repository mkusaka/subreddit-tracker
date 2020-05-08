# golang
## [1][Developers say Google's Go is 'most sought after' programming language of 2020 | ZDNet](https://www.reddit.com/r/golang/comments/gflyyn/developers_say_googles_go_is_most_sought_after/)
- url: https://www.zdnet.com/article/developers-say-googles-go-is-most-sought-after-programming-language-of-2020/
---

## [2][TermBackTime - Terminal Recorder](https://www.reddit.com/r/golang/comments/gfjw8j/termbacktime_terminal_recorder/)
- url: https://www.reddit.com/r/golang/comments/gfjw8j/termbacktime_terminal_recorder/
---
I started a project a while back called [TermBackTime](https://termbackti.me/). This allows for PTY terminal recordings to be uploaded to [Gist](https://gist.github.com/) for later playback. While I took a break on this project I plan on starting back up, implementing live terminal streaming via WebRTC.

*Example on Windows 10 using WSL:*

Gist: [https://gist.github.com/LouisT/1fc1b6cd6317180d01f60b3011490e75](https://gist.github.com/LouisT/1fc1b6cd6317180d01f60b3011490e75)

Playback: [https://termbackti.me/p/1fc1b6cd6317180d01f60b3011490e75](https://termbackti.me/p/1fc1b6cd6317180d01f60b3011490e75)

CLI:  `termbacktime play 1fc1b6cd6317180d01f60b3011490e75`

[Recording Playback](https://i.redd.it/0g04f3g53gx41.gif)

The web playback version supports some experimental features such as converting to GIF and WebM. So far I've tested on Windows 10 via WSL, OSX 10.12 to 10.14, Ubuntu 14.04.6 to 18.04.4. This was my first "major" project in Go so I'm sure parts are pretty rough. Any feedback would be greatly appreciated.
## [3][Web Application Firewall written in Go](https://www.reddit.com/r/golang/comments/gfsns4/web_application_firewall_written_in_go/)
- url: https://github.com/asalih/guardian
---

## [4][demoinfocs-golang v2.0.0 released - CS:GO Demo Parser Library for Go](https://www.reddit.com/r/golang/comments/gfrtdq/demoinfocsgolang_v200_released_csgo_demo_parser/)
- url: https://github.com/markus-wa/demoinfocs-golang
---

## [5][Tello drone remotely controlled via WebRTC (featuring pion &amp; gobot)](https://www.reddit.com/r/golang/comments/gfrqj0/tello_drone_remotely_controlled_via_webrtc/)
- url: https://www.reddit.com/r/golang/comments/gfrqj0/tello_drone_remotely_controlled_via_webrtc/
---
[https://github.com/oliverpool/tello-webrtc-fpv](https://github.com/oliverpool/tello-webrtc-fpv)

I connected the pion and gobot libraries to let my little Tello drone be remotely controlled (from everywhere on earth).

&amp;#x200B;

https://preview.redd.it/3baoid874jx41.png?width=1919&amp;format=png&amp;auto=webp&amp;s=cc318b088f73c3f29bfeb4e5612fcdf591694c8e

The latency is surprisingly small (clearly under 500ms round-trip).

I managed to have a pure-go program by reading and buffering the h264 frames manually (more info in the README).
## [6][Learn go with tests: HTTP handlers revisited](https://www.reddit.com/r/golang/comments/gf6d3u/learn_go_with_tests_http_handlers_revisited/)
- url: https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/http-handlers-revisited
---

## [7][Functional testing with your database in Go](https://www.reddit.com/r/golang/comments/gfk3od/functional_testing_with_your_database_in_go/)
- url: https://terrastruct.com/blog/functional-testing-database-go/
---

## [8][Live browser sessions](https://www.reddit.com/r/golang/comments/gfpk6k/live_browser_sessions/)
- url: https://www.reddit.com/r/golang/comments/gfpk6k/live_browser_sessions/
---
So I'm doing the classic thing of a webserver in go. I'm sure you're all bored to tears of that by now but please humour me.

For context, I'm very deliberately reinventing the wheel in a lot of places, as a learning exercise. I'm not necessarily looking for code solutions - although those are fine - but tips on general approach and languages or technologies I might need to look into.

I've managed to set up a postgres database, user login, account creation, persistent sessions via cookies, POST actions, etc. I'm interested in how you'd go about handling the more active data exchange that might need to happen for interactive pages. My understanding is that go is very much a server-side operation, and javascript, html etc. are client side. How do you begin to communicate without leaving a page, e.g. for a games that might require a continuous stream of information?

I realise this isn't necessarily a go question per se, but I've made the assumption there would be methods the go community might prefer.

Thanks for you help.
## [9][Kubeless - Build advanced applications with FaaS on top of Kubernetes](https://www.reddit.com/r/golang/comments/gfb2ne/kubeless_build_advanced_applications_with_faas_on/)
- url: https://kubeless.io/
---

## [10][Anyone making a game?](https://www.reddit.com/r/golang/comments/gficdn/anyone_making_a_game/)
- url: https://www.reddit.com/r/golang/comments/gficdn/anyone_making_a_game/
---
This might be a far-fetched question, but I was wondering if there is a video game project that is looking for a contributor. These days I find myself with a lot of time in my hands and I wanted to try something new. Thanks.
