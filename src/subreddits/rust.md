# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (20/2020)!](https://www.reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This week in Rust 338](https://www.reddit.com/r/rust/comments/gikfuy/this_week_in_rust_338/)
- url: https://this-week-in-rust.org/blog/2020/05/12/this-week-in-rust-338/
---

## [3][My team at Nike just open-sourced our Rust wrappers around the AWS Greengrass SDK](https://www.reddit.com/r/rust/comments/gjcsi8/my_team_at_nike_just_opensourced_our_rust/)
- url: https://github.com/Nike-Inc/aws-greengrass-core-sdk-rust
---

## [4][Deno 1.0 - A Secure TypeScript/JavaScript Runtime Built In Rust](https://www.reddit.com/r/rust/comments/gja35o/deno_10_a_secure_typescriptjavascript_runtime/)
- url: https://deno.land/v1
---

## [5][Benefits of full-stack Rust (My experience from implementing a browser game)](https://www.reddit.com/r/rust/comments/gjilkn/benefits_of_fullstack_rust_my_experience_from/)
- url: https://www.jakobmeier.ch/blogging/Paddlers_4.html
---

## [6][Kudos to Rust](https://www.reddit.com/r/rust/comments/gj4uei/kudos_to_rust/)
- url: https://www.reddit.com/r/rust/comments/gj4uei/kudos_to_rust/
---
Hello friends,

Today we ([Wasmer.io](https://Wasmer.io)) are announcing [the first Java library to run WebAssembly, named Wasmer JNI](https://medium.com/wasmer/announcing-the-first-java-library-to-run-webassembly-wasmer-jni-89e319d2ac7c). I know it's the Rust subreddit, but I wanted to thank the Rust community for the maturity of the tooling. Wasmer itself is written in Rust, same for Wasmer JNI which is based on [jni-rs](https://github.com/jni-rs/jni-rs/). There is Java code for the sake of simplicity, but this project would not have been that easy without jni-rs. Kudos!

I take this post as an opportunity to also salute [Pyo3](https://pyo3.rs/) (which we use for our [Python integration](https://github.com/wasmerio/python-ext-wasm)), [Rutie](https://github.com/danielpclark/rutie) (which we use for our [Ruby integration](https://github.com/wasmerio/ruby-ext-wasm)), and [pg-extend](https://github.com/bluejekyll/pg-extend-rs/) (which we use for our [PostgreSQL integration](https://github.com/wasmerio/postgres-ext-wasm)). To be clear, it's not a hidden marketing message: I'm just amazed how easy it is to embed Rust in very various environments. Wasmer is also available in C, [PHP](https://github.com/wasmerio/php-ext-wasm) and so on. How crazy that is!

The Rust teams are doing an amazing job to keep Rust awesome whist being aligned with many standards, making it easy to use or embed it everywhere. The tooling and the ecosystem around Rust is simply absolutely awesome.

I'm programming in Rust full time since a year and a month now. It has changed my life as a developer. Thank you for that.
## [7][Rust: Structuring and handling errors in 2020](https://www.reddit.com/r/rust/comments/gj8inf/rust_structuring_and_handling_errors_in_2020/)
- url: https://nick.groenen.me/posts/rust-error-handling/
---

## [8][Castor: A Browser for the small internet (Gemini, Gopher, Finger)](https://www.reddit.com/r/rust/comments/gizchb/castor_a_browser_for_the_small_internet_gemini/)
- url: https://sr.ht/~julienxx/Castor/
---

## [9][iron-oxide: new Rust bindings to Metal](https://www.reddit.com/r/rust/comments/gje8eu/ironoxide_new_rust_bindings_to_metal/)
- url: https://www.reddit.com/r/rust/comments/gje8eu/ironoxide_new_rust_bindings_to_metal/
---
[Repository](https://github.com/Lonsdaleiter/iron-oxide) and [crate](https://crates.io/crates/iron-oxide).

Vulkan and OpenGL are represented well in Rust with many bindings and wrappers (Vulkano, ash, erupt, luminance,  gl, etc), but there is only one Metal crate: metal-rs. I figured the time was ripe for an alternative crate for programming with Metal in Rust.

Why Metal? It's an elegant, understandable low-level library that perfectly fits the role of a graphics abstraction one would build atop something like Vulkan.

Why iron-oxide and not metal-rs? iron-oxide doesn't use macros for its types, so auto-complete still works. Also, unsafety is used properly. Hence, unsafe rusty metal.
## [10][Mun struct Hot Reloading embedded in Rust](https://www.reddit.com/r/rust/comments/gj2s4s/mun_struct_hot_reloading_embedded_in_rust/)
- url: https://v.redd.it/3lmib1vdaky41
---

## [11][Beginner question please help me understand Rust async, its libraries and history](https://www.reddit.com/r/rust/comments/gjj6tw/beginner_question_please_help_me_understand_rust/)
- url: https://www.reddit.com/r/rust/comments/gjj6tw/beginner_question_please_help_me_understand_rust/
---
I understand that starting with 1.39 Rust has async / await keywords in the language. Before this version, one has to use a library for async programming in Rust. Does it mean I need only Rust standard library for async programming from 1.39 onwards without the need of any third party library? 

There is this "futures" crate under official rust-lang Github repos. Is this crate meant to be used for Rust before 1.39 or it is build for modern rust async?

And then there is tokio and async-std libraries.  Apparently many async frameworks is build on top of either one of these. Why the need of additional library as there is an official "futures" crate? Is it for historical reason? If you were to write a brand new async web framework (like warp or tide), can you just use "futures" crate?
## [12][Loops in Rust; Breaking From Nested Loops](https://www.reddit.com/r/rust/comments/gjlrg7/loops_in_rust_breaking_from_nested_loops/)
- url: https://qvault.io/2020/05/14/loops-in-rust-breaking-from-nested-loops/
---

