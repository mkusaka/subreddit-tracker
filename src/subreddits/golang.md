# golang
## [1][TinyGo: it just works!](https://www.reddit.com/r/golang/comments/hq0bau/tinygo_it_just_works/)
- url: https://www.reddit.com/r/golang/comments/hq0bau/tinygo_it_just_works/
---
I have a couple of years of embedded experience, both fiddling around at home and professionally at work. I've used a number of different frameworks: C++/Arduino, Rust,  STM32HAL/FreeRTOS, MicroPython, and even AVR Assembly.

Most of these have a steep learning curve involving a lot of frustration at first.

Since I also use Go at work (for host-side development), I decided to give TinyGo a spin. I grabbed an Adafruit Itsy-Bitsy M0 and an SSD1306 I had lying around, copied the TinyGo example code and ... wow, it just worked!

So, kudos to the TinyGo team for an amazing initial experience!
## [2][Golang Integration Testing](https://www.reddit.com/r/golang/comments/hq5ciy/golang_integration_testing/)
- url: https://hackandsla.sh/posts/2020-07-12-golang-integration-testing/
---

## [3][Making Pong with Go and WebAssembly](https://www.reddit.com/r/golang/comments/hpzy7f/making_pong_with_go_and_webassembly/)
- url: https://dstoiko.github.io/posts/go-pong-wasm/
---

## [4][Code review for handling timeout](https://www.reddit.com/r/golang/comments/hqbpy2/code_review_for_handling_timeout/)
- url: https://www.reddit.com/r/golang/comments/hqbpy2/code_review_for_handling_timeout/
---
As a learning process, I would like to learn the correct way to use `context` package by making a program.

I've made an aggregation server which collects response from other servers and returns to client. It has been implemented by handling timeout with  \`context\` package.

The internal servers can include *fast* and *slow* handler as this link: [https://play.golang.org/p/1eN6kwyrSH3](https://play.golang.org/p/1eN6kwyrSH3)

And the aggregation server raises `internal error` whenever one of endpoints failed as this link: [https://play.golang.org/p/0XttDep7kNO](https://play.golang.org/p/0XttDep7kNO)

I’m looking for some feed-backs on 2 above programs.

Is it a correct way to handle timeout ? Is there any way to improve the programs.

Thank you
## [5][Learning Go is in Early Release](https://www.reddit.com/r/golang/comments/hq7qw3/learning_go_is_in_early_release/)
- url: https://www.reddit.com/r/golang/comments/hq7qw3/learning_go_is_in_early_release/
---
I'm happy to announce that my upcoming book, Learning Go, is in Early Release at O'Reilly Safari:

[https://learning.oreilly.com/library/view/learning-go/9781492077206/](https://learning.oreilly.com/library/view/learning-go/9781492077206/)

Learning Go is an in-depth guide for developers that teaches you how to write idiomatic Go. It's due out in early 2021. Take a look!
## [6][How would I make more "chrome like" requests with golang?](https://www.reddit.com/r/golang/comments/hq9d1s/how_would_i_make_more_chrome_like_requests_with/)
- url: https://www.reddit.com/r/golang/comments/hq9d1s/how_would_i_make_more_chrome_like_requests_with/
---
There are WAFs (like cloudflare) that block requests that seem to be made from bots. How would I make my requests more chrome like? I know I can change the ClientHello fingerprint, and change TLS versions, but is there more I can do? I know electron (NodeJS) managed to import the entire chrome networking library. Is that possible with golang?
## [7][MasterPlan, a Graphical Idea-board / Planning Software written in Go reaches v0.4!](https://www.reddit.com/r/golang/comments/hq517n/masterplan_a_graphical_ideaboard_planning/)
- url: https://www.reddit.com/r/golang/comments/hq517n/masterplan_a_graphical_ideaboard_planning/
---
Hello!

I'm SolarLune, and I wanted to post about a graphical free-form idea-board / planning software that I'm writing in Go. The program is called MasterPlan, and has reached v0.4.0, with the latest feature being undo and redo of actions and automatic backups. It runs on Windows, Mac, and Linux, of course.

I actually posted about this project right here [once before](https://www.reddit.com/r/golang/comments/dkv6dy/masterplan_a_wip_of_a_project_management_software/) and got great feedback, but that was around 8 months ago, before I made the software and code publicly available. Since then, I've done so, and have been steadily working on it. I've made quite a bit of progress since then, so I figured it would be good to post about it again.

There's still quite a bit I want to implement (including more idiomatic / smart Go code, alongside other "actual" features), but it's pretty solid right now. I originally started working on it because I wanted something to help me plan out projects, but found most existing applications (Trello, Notion, Hack n Plan, etc) to be not quite right for independent developers.

You can find the Github repository [here](https://github.com/SolarLune/masterplan), and a write-up and screenshots on the market page [here](https://solarlune.itch.io/masterplan). I'm trying something a bit different for the licensing, as the source code is available on GitHub, but the code is still privately licensed.

If anyone has any questions or suggestions, please feel free to let me know! I would also happily answer any technical questions people might have (though I don't feel like it's well-written enough to look to as an example, haha).
## [8][Going Driverless (Oscar Söderlund, GopherCon EU 2020)](https://www.reddit.com/r/golang/comments/hpt2c4/going_driverless_oscar_söderlund_gophercon_eu_2020/)
- url: https://www.youtube.com/watch?v=IbggJHJUv0U
---

## [9][Go Channel Use Cases](https://www.reddit.com/r/golang/comments/hpnlc4/go_channel_use_cases/)
- url: https://go101.org/article/channel-use-cases.html
---

## [10][Generative art with fractals and geometry in golang.](https://www.reddit.com/r/golang/comments/hpwegv/generative_art_with_fractals_and_geometry_in/)
- url: https://github.com/0x0f0f0f/tripbot9000
---

