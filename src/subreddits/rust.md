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
## [2][What's everyone working on this week (30/2020)?](https://www.reddit.com/r/rust/comments/hurk9f/whats_everyone_working_on_this_week_302020/)
- url: https://www.reddit.com/r/rust/comments/hurk9f/whats_everyone_working_on_this_week_302020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-30-2020/46132?u=llogiq)!
## [3][Rocket can now compile on stable Rust](https://www.reddit.com/r/rust/comments/hviz2q/rocket_can_now_compile_on_stable_rust/)
- url: https://github.com/SergioBenitez/Rocket/issues/19
---

## [4][Tokio: new website &amp; new guides](https://www.reddit.com/r/rust/comments/hvalrb/tokio_new_website_new_guides/)
- url: https://tokio.rs
---

## [5][Why does Rust seem to be pulling so many node developers as opposed to C/C++ developers?](https://www.reddit.com/r/rust/comments/hvpgm4/why_does_rust_seem_to_be_pulling_so_many_node/)
- url: https://www.reddit.com/r/rust/comments/hvpgm4/why_does_rust_seem_to_be_pulling_so_many_node/
---
My day job is C programming on network appliances at a medium sized Cisco competitor, but most people here haven't even heard of Rust (unless I've told them about it) let alone interested in using it. On the other hand I frequently see former or current node developers talk about Rust on this subreddit and other sites. Is there some silent group that isn't getting their word out because coincidentally, if you're good at javascript you can make a website for writing blog posts with much easier than others? (I want to see Rust take off in embedded, lower level systems programming and games, aka performance critical locations, but that doesn't seem to be happening yet)
## [6][Nushell 0.17.0 - now with webassembly support, custom keybindings, benchmarking, and more](https://www.reddit.com/r/rust/comments/hvhsyg/nushell_0170_now_with_webassembly_support_custom/)
- url: https://www.nushell.sh/blog/2020/07/21/nushell_0_17_0.html
---

## [7][This Week in Rust 348](https://www.reddit.com/r/rust/comments/hvjf4i/this_week_in_rust_348/)
- url: https://this-week-in-rust.org/blog/2020/07/21/this-week-in-rust-348/
---

## [8][Scalar code is faster than simd in rust](https://www.reddit.com/r/rust/comments/hvsrl5/scalar_code_is_faster_than_simd_in_rust/)
- url: https://www.reddit.com/r/rust/comments/hvsrl5/scalar_code_is_faster_than_simd_in_rust/
---
I wanted to rewrite some of my old C code in rust as a learning experience. Here are the C and rust versions of the same code. (It's actually a C code i just had to use chrono from c++ std)

[https://gist.github.com/atkurtul/0cb3362c11d944c8a372cda80a3b6ac9](https://gist.github.com/atkurtul/0cb3362c11d944c8a372cda80a3b6ac9)

And of course i needed to benchmark the simd in rust.

Here's the results i get from rust on my pc

    87.662047ms  -&gt; simd 
    33.557277ms  -&gt; scalar
    87.290656ms  -&gt; simd
    32.503664ms  -&gt; scalar

and C results

    18.780300 ms -&gt; simd
    31.580484 ms -&gt; scalar
    17.115702 ms -&gt; simd
    31.227010 ms -&gt; scalar

As you can see scalar versions time almost the same which isnt very surprising. However simd versions make me worry that im missing something critical in the rust version of the code.

Optimization flags in Cargo.toml that i compiled with

    [profile.dev]
    opt-level = 3
    
    [profile.release]
    opt-level = 3

and

    clang++ main.cpp -O3 -mfma

Is the problem in my code or simd in rust just isnt implemented properly?
## [9][Anyone here from the oil&amp;gas industry](https://www.reddit.com/r/rust/comments/hvtfi2/anyone_here_from_the_oilgas_industry/)
- url: https://www.reddit.com/r/rust/comments/hvtfi2/anyone_here_from_the_oilgas_industry/
---
I've been slowly learning and practicing Rust for the past half a year or so. I love the language and the community around it (huge thank to /u/ErichDonGubler for always answering my silly questions). My project is nothing fancy compared to what you geniuses post here about. I am just curious whether anyone at all might be interested in it. I come from the oil&amp;gas reservoir simulation background.  My app is a reader for the common simulation output format: https://github.com/mindv0rtex/eclair. It is mostly a learning vehicle, but so far I managed to convert the format to MessegaPack via `serde`, wrap this functionality with Python via `PyO3`, ever so slighly looked at WASM, and currently I'm working on plotting the results to HTML via `Plotly`.
## [10][Enable compilation with stable Rust. · SergioBenitez/Rocket@56a6172 · GitHub](https://www.reddit.com/r/rust/comments/hviwop/enable_compilation_with_stable_rust/)
- url: https://github.com/SergioBenitez/Rocket/commit/56a617262525de758cb47a0d71049db948f24bbd
---

## [11][Disk Usage Analyzer v2.10: TUI is finally working correctly on Windows](https://www.reddit.com/r/rust/comments/hvp3nf/disk_usage_analyzer_v210_tui_is_finally_working/)
- url: https://github.com/Byron/dua-cli/releases/tag/v2.10.0
---

## [12][Show r/rust: A tiny UNIX shell, from C to Rust, from 1990 to 2020](https://www.reddit.com/r/rust/comments/hv78hc/show_rrust_a_tiny_unix_shell_from_c_to_rust_from/)
- url: https://github.com/zserge/tinysh
---

