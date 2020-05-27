# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (22/2020)!](https://www.reddit.com/r/rust/comments/gr3r32/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gr3r32/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (22/2020)?](https://www.reddit.com/r/rust/comments/gr3s6m/whats_everyone_working_on_this_week_222020/)
- url: https://www.reddit.com/r/rust/comments/gr3s6m/whats_everyone_working_on_this_week_222020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-22-2020/43296?u=llogiq)!
## [3][Multiplatform Minecraft-compatible client written in Rust](https://www.reddit.com/r/rust/comments/gris3i/multiplatform_minecraftcompatible_client_written/)
- url: https://github.com/iceiix/stevenarella
---

## [4][A retrospective on the 2018 rust-lang.org redesign](https://www.reddit.com/r/rust/comments/grb31n/a_retrospective_on_the_2018_rustlangorg_redesign/)
- url: https://blog.rust-lang.org/inside-rust/2020/05/26/website-retrospective.html
---

## [5][easy-parallel: like scoped threads, but simpler](https://www.reddit.com/r/rust/comments/grhq9p/easyparallel_like_scoped_threads_but_simpler/)
- url: https://docs.rs/easy-parallel
---

## [6][Compiling Rust binaries for Windows 98 SE and more: a journey](https://www.reddit.com/r/rust/comments/gr0xqa/compiling_rust_binaries_for_windows_98_se_and/)
- url: https://seri.tools/blog/compiling-rust-for-legacy-windows/
---

## [7][Rust newbie here, need suggestions on how to improve my code](https://www.reddit.com/r/rust/comments/grgf8b/rust_newbie_here_need_suggestions_on_how_to/)
- url: https://www.reddit.com/r/rust/comments/grgf8b/rust_newbie_here_need_suggestions_on_how_to/
---
Hello!

I've made a simple procedural 2D world generator. For now it's very poor, and only generates a tiny ASCII overworld (mountains+sea). There's probably a lot of room for improvement, so I thought it's worth posting here to gain some useful feedback.

What do you guys think?  [https://github.com/jabczyk/worldgen](https://github.com/jabczyk/worldgen)
## [8][Faster Integer Parsing](https://www.reddit.com/r/rust/comments/gr8mho/faster_integer_parsing/)
- url: https://kholdstare.github.io/technical/2020/05/26/faster-integer-parsing.html
---

## [9][macro-compose: a library trying to simplify and organize proc-macros](https://www.reddit.com/r/rust/comments/grhgf3/macrocompose_a_library_trying_to_simplify_and/)
- url: https://github.com/Freax13/macro-compose
---

## [10][Contributing to Rust](https://www.reddit.com/r/rust/comments/gr1oel/contributing_to_rust/)
- url: https://blog.elinvynia.com/posts/2020-05-26-contributing-to-rust.html
---

## [11][Zero Cost Abstractions: Rust vs C++](https://www.reddit.com/r/rust/comments/gqw2gj/zero_cost_abstractions_rust_vs_c/)
- url: https://www.rottedfrog.co.uk/?p=24
---

## [12][[ANN] PWatch - a process watcher 0.2.0](https://www.reddit.com/r/rust/comments/grf0bi/ann_pwatch_a_process_watcher_020/)
- url: https://www.reddit.com/r/rust/comments/grf0bi/ann_pwatch_a_process_watcher_020/
---
**pwatch** is a terminal process watcher. It is akin to **top** utility but its purpose is to watch closely only a few selected processes(filter them with a list of PIDs or regular expression).

Highlights:

* \- CPU and memory historical graphs per processes.
* \- Measure I/O total throughput and memory usage change since a moment in the past
* \- Graphs for dead process are not removed(but stopped updating), so you can examine their CPU/memory usage right before they exited at any moment

Crate: [https://crates.io/crates/pwatch](https://crates.io/crates/pwatch)

Sources: [https://github.com/VladimirMarkelov/pwatch](https://github.com/VladimirMarkelov/pwatch)

Screencast: a few Firefox process after a user clicked button "Refresh" in the browser:

https://i.redd.it/z4mw2dr5e9151.gif
