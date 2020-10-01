# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (40/2020)!](https://www.reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 358](https://www.reddit.com/r/rust/comments/j2yfcu/this_week_in_rust_358/)
- url: https://this-week-in-rust.org/blog/2020/09/30/this-week-in-rust-358/
---

## [3][Introducing binserve - A blazingly fast static web server in a single binary you can set up with zero code](https://www.reddit.com/r/rust/comments/j35c81/introducing_binserve_a_blazingly_fast_static_web/)
- url: https://www.reddit.com/r/rust/comments/j35c81/introducing_binserve_a_blazingly_fast_static_web/
---
Hello, fellow Rustaceans,

I am new to Rust and started reading the Rust Book to learn it. When I lurked through it, I found Rust to have unique principles on how the language is designed and that really got me interested! I decided to choose it as my language of choice from now on and to learn something, you gotta do something so I did! This is it, my first ever project in **Rust**!

Introducing **binserve** ⚡🦀

A blazingly fast static web server in a single binary you can set up with zero code. 🔥

**Features:**

* **Single binary with no dependencies and everything built-in.**
* **Blazingly Fast!** ⚡ - it's built on top of Actix!
* **Handlebars** template engine.
* **Easy Routing**
* **Secure By Design**
* **Config &amp; Static File Assistance**
* **Supports Any Static File** (`Content-Type`)
* **Custom Error Page Support**
* And much more...

**Check it out!**

[https://github.com/mufeedvh/binserve](https://github.com/mufeedvh/binserve)

As this is my first project, I might've made some mistakes in semantics or in implementing things the Rust way, if you're free or want to help, please give me feedback/suggestions, I would really appreciate it! ❤️

`Ok(())` Thanks!
## [4][Colliding galaxies simulation using Rust and WebAssembly 🌌](https://www.reddit.com/r/rust/comments/j2se15/colliding_galaxies_simulation_using_rust_and/)
- url: https://v.redd.it/sy1p08cr1cq51
---

## [5][Mobile operating system SailfishOS now supports Rust](https://www.reddit.com/r/rust/comments/j35d01/mobile_operating_system_sailfishos_now_supports/)
- url: https://forum.sailfishos.org/t/release-notes-pallas-yllastunturi-3-4-0/2258
---

## [6][Announcing the Portable SIMD Project Group | Inside Rust Blog](https://www.reddit.com/r/rust/comments/j2p44j/announcing_the_portable_simd_project_group_inside/)
- url: https://blog.rust-lang.org/inside-rust/2020/09/29/Portable-SIMD-PG.html
---

## [7][Got my first segfault after 18 months of Rust programming](https://www.reddit.com/r/rust/comments/j2qknq/got_my_first_segfault_after_18_months_of_rust/)
- url: https://i.redd.it/dx06hngzkbq51.png
---

## [8][Should I learn rust or c++(or any other recommendations) to get a job in the next 5-10 years in the field of robotics?](https://www.reddit.com/r/rust/comments/j37fli/should_i_learn_rust_or_cor_any_other/)
- url: https://www.reddit.com/r/rust/comments/j37fli/should_i_learn_rust_or_cor_any_other/
---
I have seen rust rising in popularity on the stack overflow survey over the past years, however, it seems to be rarely used? I am 16 at the moment and conflicted on which to learn. This will not be my first language so ease of learning, is not an issue.

Edit: and why should I learn said language?
## [9][Avalog - An experimental logical programming language for higher dimensional mathematical programming](https://www.reddit.com/r/rust/comments/j2zml3/avalog_an_experimental_logical_programming/)
- url: https://crates.io/crates/avalog
---

## [10][Revisiting a 'smaller Rust'](https://www.reddit.com/r/rust/comments/j2l9v9/revisiting_a_smaller_rust/)
- url: https://without.boats/blog/revisiting-a-smaller-rust/
---

## [11][Unit tests uploaded to crates.io](https://www.reddit.com/r/rust/comments/j35wsu/unit_tests_uploaded_to_cratesio/)
- url: https://www.reddit.com/r/rust/comments/j35wsu/unit_tests_uploaded_to_cratesio/
---
Hi  


One thing that I have been thinking about (and exposed to when packaging rust crates for Debian) is that there is quite a few crates uploaded to [crates.io](https://crates.io) that have tests that depend on resources that are not packaged into the crate.  


This isn't a huge problem, as the developer hopefully runs the tests before uploading.  


But it's very useful to also run the tests when packaging a crate for different architectures, it's nice to discover [endian problems](https://github.com/elichai/random-rs/issues/3) or that char is [different on arm64](https://github.com/jeaye/ncurses-rs/pull/194) before they reach production. 

I guess that the problem can be fixed by adding `include` sections in the manifest, but it would be nice if cargo tested this automatically before uploading it.

I'm thinking that adding a step to cargo package that unpacks the .crate into a temporary directory and runs the tests there would improve the ecosystem. Do anyone see any problems with that or have any other feedback?
## [12][Optimization - Making Rust Code Go Brrrr](https://www.reddit.com/r/rust/comments/j2pmd3/optimization_making_rust_code_go_brrrr/)
- url: https://aspenuwu.me/posts/Optimization%20-%20Making%20Rust%20Code%20Go%20Brrrr.html
---

