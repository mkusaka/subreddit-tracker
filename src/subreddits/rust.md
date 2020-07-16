# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (29/2020)!](https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 347](https://www.reddit.com/r/rust/comments/hrc4dt/this_week_in_rust_347/)
- url: https://this-week-in-rust.org/blog/2020/07/14/this-week-in-rust-347/
---

## [3][Benchmarking gRPC in Rust and Go](https://www.reddit.com/r/rust/comments/hs5k36/benchmarking_grpc_in_rust_and_go/)
- url: https://medium.com/@Rustling_gopher/benchmarking-grpc-in-rust-go-184545e7688a
---

## [4][Blog Post: Two Beautiful Rust Programs](https://www.reddit.com/r/rust/comments/hrwckz/blog_post_two_beautiful_rust_programs/)
- url: https://matklad.github.io//2020/07/15/two-beautiful-programs.html
---

## [5][Shipping Const Generics in 2020](https://www.reddit.com/r/rust/comments/hs9hbn/shipping_const_generics_in_2020/)
- url: https://without.boats/blog/shipping-const-generics/
---

## [6][Writing a kernel driver with Rust.](https://www.reddit.com/r/rust/comments/hrwyl8/writing_a_kernel_driver_with_rust/)
- url: https://not-matthias.github.io/kernel-driver-with-rust/
---

## [7][j4rs v0.12.0: Java to Rust direction](https://www.reddit.com/r/rust/comments/hs4677/j4rs_v0120_java_to_rust_direction/)
- url: https://astonbitecode.github.io/blog/post/j4rs_0.12.0/
---

## [8][Function overloading in Rust](https://www.reddit.com/r/rust/comments/hrxzpi/function_overloading_in_rust/)
- url: https://medium.com/@nrabulinski/function-overloading-in-rust-d591aff64a03
---

## [9][Building and debugging a high-throughput daemon in Rust](https://www.reddit.com/r/rust/comments/hrvu2t/building_and_debugging_a_highthroughput_daemon_in/)
- url: https://brokenco.de/2020/07/15/high-throughput-in-rust.html
---

## [10][IntelliJ Rust 0.3: New Macro Expansion Engine](https://www.reddit.com/r/rust/comments/hrknxy/intellij_rust_03_new_macro_expansion_engine/)
- url: https://blog.jetbrains.com/clion/2020/07/intellij-rust-0-3-new-macro-expansion-engine/
---

## [11][How to handle asynchrony in otherwise blocking code for WASM?](https://www.reddit.com/r/rust/comments/hs8tre/how_to_handle_asynchrony_in_otherwise_blocking/)
- url: https://www.reddit.com/r/rust/comments/hs8tre/how_to_handle_asynchrony_in_otherwise_blocking/
---
Hey,

in recent efforts on progressing my rust-skills I have been working on a little toy project called [poppy](https://github.com/liona24/poppy), which is basically a simple framework to play poker. Originally I planned on re-implementing the [Pluribus AI](https://science.sciencemag.org/content/365/6456/885), but time is scarce. Eventually we will get there.

Anyway I wanted to make a little demo build targeting WASM. This turned out to be quite challenging since I made some unfortunate design decisions: Currently the game is owning each player and hence *awakes* each one when it is his/her time to shine. This is a simple blocking design and seemed fine at that point in time since I wanted to implement an AI anyway.

Now however I am failing to implement user interaction in a javascript-callback-world, since I cannot simply return from the invoked function until I (the player) know what to do. I was wondering how to solve this problem? Is this something I should have foreseen in the first place or are there some patterns which can be used here? 

I suspect in a multi-threaded environment one could create a polling thread or fallback to actual async/await operations though I am not sure if this is the right place here. Is it generally advisable to implement anything that might require user interaction in a callback-oriented design?
## [12][Why even unused data needs to be valid](https://www.reddit.com/r/rust/comments/hrlozh/why_even_unused_data_needs_to_be_valid/)
- url: https://www.ralfj.de/blog/2020/07/15/unused-data.html
---

