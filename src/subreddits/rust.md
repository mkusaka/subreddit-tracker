# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (45/2020)!](https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (45/2020)?](https://www.reddit.com/r/rust/comments/jmiktw/whats_everyone_working_on_this_week_452020/)
- url: https://www.reddit.com/r/rust/comments/jmiktw/whats_everyone_working_on_this_week_452020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-45-2020/50962?u=llogiq)!
## [3][Bevy 0.3](https://www.reddit.com/r/rust/comments/jnirl3/bevy_03/)
- url: https://bevyengine.org/news/bevy-0-3/
---

## [4][async-std v1.7.0 has been released!](https://www.reddit.com/r/rust/comments/jnueup/asyncstd_v170_has_been_released/)
- url: https://github.com/async-rs/async-std/releases/tag/v1.7.0
---

## [5][With the launch of Bevy Game Engine 0.3, I'm introducing chunk base Tile Maps!](https://www.reddit.com/r/rust/comments/jntjp8/with_the_launch_of_bevy_game_engine_03_im/)
- url: https://www.reddit.com/r/rust/comments/jntjp8/with_the_launch_of_bevy_game_engine_03_im/
---
 [https://crates.io/crates/bevy\_tilemap](https://crates.io/crates/bevy_tilemap) 

This is not intended to be just another Tile Map. It is meant to be a framework and extensible by design, like Bevy. There is an emphasis placed on generic traits to accomplish that. As well as work done to keep it as close to Bevy API as possible. This allows anyone to create their own tiles, chunks and maps and still retain the speed of a handcrafted multi-threaded chunk loader and tile map. 

If you have comments or suggestions do let me know below! Thanks!
## [6][Is Rust worth it for frontend development?](https://www.reddit.com/r/rust/comments/jntj2s/is_rust_worth_it_for_frontend_development/)
- url: https://www.reddit.com/r/rust/comments/jntj2s/is_rust_worth_it_for_frontend_development/
---
Hey everyone!

I'm currently working on a cloud server in Rust and I already have a solid backend. So now the big question is, what to use for the frontend.

Of course I'd like to write the whole thing in Rust but after experimenting with it for a few hours I'm not sure whether that's the best solution. Anyway, here are my thoughts:

Disadvantages of using Rust:

* Additional dev dependencies like wasm toolchain and wasm-bindgen/wasm-pack
* Makefiles or other build system required: wasm-bindgen/wasm-pack needs to be called separately to create "glue code" and requires a lot parameters to work correctly
* Performance: neither Yew nor Seed look very good in benchmarks, at least they are not faster than JavaScript
* WASM bindings don't feel native yet, JavaScript still seems to be a lot easier for interacting with Web-APIs

Advantages of using Rust:

* It's Rust, not JavaScript ;)

I'm curious what you think about about this matter :)

P. S. I'm using actix-web for the server and tried out Seed for the frontend. I've already got a Rust-based frontend setup working with my cloud server and pushed the code to a new branch: [https://github.com/AaronErhardt/Triox/tree/experimental-frontend](https://github.com/AaronErhardt/Triox/tree/experimental-frontend).

EDIT: Thanks a lot for your answers!
## [7][Hardware performance counters for the Rust compiler](https://www.reddit.com/r/rust/comments/jn9892/hardware_performance_counters_for_the_rust/)
- url: https://hackmd.io/sH315lO2RuicY-SEt7ynGA?view
---

## [8][bacon: a background code checker, to keep in a side terminal](https://www.reddit.com/r/rust/comments/jnde63/bacon_a_background_code_checker_to_keep_in_a_side/)
- url: https://github.com/Canop/bacon
---

## [9][This month in Dimforge #2 (October 2020) - New collider shapes, collision filters, restitution, user-defined data, and work-in-progress coupling with fluids](https://www.reddit.com/r/rust/comments/jneh38/this_month_in_dimforge_2_october_2020_new/)
- url: https://www.dimforge.com/blog/2020/11/01/this-month-in-dimforge
---

## [10][rena: a renamer program utility thing (am bad at describing my projects)](https://www.reddit.com/r/rust/comments/jnmius/rena_a_renamer_program_utility_thing_am_bad_at/)
- url: https://github.com/lyssieth/rena
---

## [11][Actix Web Rust](https://www.reddit.com/r/rust/comments/jnlyc6/actix_web_rust/)
- url: https://www.reddit.com/r/rust/comments/jnlyc6/actix_web_rust/
---
I've recently taken an interest in Actix Web, especially with its performance benefits. However, I've noticed it's not well documented. It's usually out of date or isn't comprehensive enough. Is there a way someone can contribute to the official documentation for it? Or a more general way to make a more detailed and example oriented documentation?
## [12][R3: a proof-of-concept static RTOS utilizing compile-time function evaluation for static configuration](https://www.reddit.com/r/rust/comments/jn9ksr/r3_a_proofofconcept_static_rtos_utilizing/)
- url: https://github.com/yvt/r3
---

