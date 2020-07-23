# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (30/2020)!](https://www.reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 348](https://www.reddit.com/r/rust/comments/hvjf4i/this_week_in_rust_348/)
- url: https://this-week-in-rust.org/blog/2020/07/21/this-week-in-rust-348/
---

## [3][Rust explained using easy English](https://www.reddit.com/r/rust/comments/hw8rlc/rust_explained_using_easy_english/)
- url: https://github.com/Dhghomon/easy_rust
---

## [4][Giving Rust Another Shot in 2020](https://www.reddit.com/r/rust/comments/hwdaxo/giving_rust_another_shot_in_2020/)
- url: https://sharpend.io/giving-rust-another-shot-in-2020/
---

## [5][UReq - Lightweight alternative to Reqwest](https://www.reddit.com/r/rust/comments/hwdrlc/ureq_lightweight_alternative_to_reqwest/)
- url: https://github.com/algesten/ureq
---

## [6][Rust's CI is moving to GitHub Actions](https://www.reddit.com/r/rust/comments/hwf5zs/rusts_ci_is_moving_to_github_actions/)
- url: https://blog.rust-lang.org/inside-rust/2020/07/23/rust-ci-is-moving-to-github-actions.html
---

## [7][Verify the Linux Kernel pack in 12 seconds with git-oxide*](https://www.reddit.com/r/rust/comments/hw92v0/verify_the_linux_kernel_pack_in_12_seconds_with/)
- url: https://github.com/Byron/git-oxide/issues/1#issuecomment-662168881
---

## [8][Cross-compiling to Redox using Nix](https://www.reddit.com/r/rust/comments/hw54ne/crosscompiling_to_redox_using_nix/)
- url: https://www.redox-os.org/news/redox-plus-nix-0/
---

## [9][[ANN] Sauron 0.28 - a client side web framework library, overtakes some of the popular web framework in terms of performance. Can now be used in stable.](https://www.reddit.com/r/rust/comments/hw07kt/ann_sauron_028_a_client_side_web_framework/)
- url: https://www.reddit.com/r/rust/comments/hw07kt/ann_sauron_028_a_client_side_web_framework/
---
[Project site](https://github.com/ivanceras/sauron)

While some of the popular web framework's performance is dropping, `sauron`'s is continually improving with every new releases.

There are 2 benchmark site, which you can benchmark it yourself:

 - [Benchmark 1](https://ivanceras.github.io/todo-mvc-bench/)
 - [Benchmark 2](https://ivanceras.github.io/todomvc-benchmark/)


The core `virtual-dom` implementation is now decoupled from the main project with a new name [`mt-dom`](https://github.com/ivanceras/mt-dom/). This provides a generic virtual-dom implementation, hence it can be used in variety of ways, such as server-side rendering, polyglot UI implementation such as [`sauron-native`](https://github.com/ivanceras/sauron-native) which is used to do virtual DOM diff on native UI elements.

`sauron` and `sauron-native` is a counterpart to `react` and `react-native`, but written in rust and primarily for desktop applications.
## [10][Rust with Python?](https://www.reddit.com/r/rust/comments/hw4mmj/rust_with_python/)
- url: https://www.reddit.com/r/rust/comments/hw4mmj/rust_with_python/
---
Hello everyone. I apologize for the format, on phone rn. 

I'm a CS student, learning to get into data science and I code in Python. I love front end as well so I use a fair bit of vanilla javascript, html/css for my fun projects. I want to learn a low level language but don't really want to touch C++ ever again and I bumped into Rust in my desperate attempts to find a replacement. After reading multiple articles and being more confused than I was before, I decided to come to all of you for help. 

Most of what I do is apply mathematical concepts using python, build them from scratch, analyse datasets, build websites and wander in the endless desert of weird code that GitHub is. I wanted to write my own mathmatical library and I wanted to know if Rust is something I should learn. It can be done, yes, but... Should I? 

I don't know where I want to go from there but is Rust worth adding to my arsenal when I plan on becoming a data scientist considering I love building stuff as well? What can I do after I learn it? 

There's an endless ocean of things and I don't know what to do. Please guide me dear Rustlings, and perhaps, I may become one of you.
## [11][Condure: a high performance HTTP/WebSocket connection manager](https://www.reddit.com/r/rust/comments/hvzaui/condure_a_high_performance_httpwebsocket/)
- url: https://github.com/fanout/condure
---

## [12][I made a spaced repetition (flashcard) companion for The Rust Book!](https://www.reddit.com/r/rust/comments/hvwd4o/i_made_a_spaced_repetition_flashcard_companion/)
- url: https://www.reddit.com/r/rust/comments/hvwd4o/i_made_a_spaced_repetition_flashcard_companion/
---
tl;dr Free flashcards for The Rust Book. Screenshots of some example cards [here](https://imgur.com/a/jwCvr7e), website [here](https://polysift.com/preview/nyrmfyrwfc8).

Hi all! I'm just finishing up a spaced repetition flashcard deck for The Rust Book, and I thought you guys might be interested in it.

I've read The Rust Book several times, but I find that it just goes in one ear and out the other--I don't internalize it. Since I'm a big flashcards guy, I made some flashcards to bridge the gap between "I vaguely remember something about blanket implementations from The Rust Book" and "I intuitively know what a blanket implementation is and can write one without looking up the syntax."

If you're unfamiliar with it, the basic idea of spaced repetition is to distill the topic that you're learning into flashcards. If you get cards right, you see them less often, and if you get them wrong, you see them more often. If you do reviews every day, you can retain a high percentage of the cards over a very long time span. SRS (spaced repetition systems) are mostly used for stuff like language learning, but I've personally found good success using them for programming languages as well.

[Here](https://imgur.com/a/jwCvr7e) are some demo cards (don't worry, the whole deck is free), and [here's](https://polysift.com/preview/nyrmfyrwfc8) the link to the website. The deck is in active development &amp; I'd really appreciate feedback!

PS: I'm also the guy who made the website, mostly so that I could create decks like this, which aren't possible with older SRS software like Anki. Obviously also open for broader feedback! :)
