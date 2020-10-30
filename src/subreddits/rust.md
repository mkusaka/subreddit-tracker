# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (44/2020)!](https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 362](https://www.reddit.com/r/rust/comments/jk35ha/this_week_in_rust_362/)
- url: https://this-week-in-rust.org/blog/2020/10/28/this-week-in-rust-362/
---

## [3][These Weeks In Actix | Sep-Oct '20](https://www.reddit.com/r/rust/comments/jkv5xu/these_weeks_in_actix_sepoct_20/)
- url: https://www.reddit.com/r/rust/comments/jkv5xu/these_weeks_in_actix_sepoct_20/
---
Since the [release of Actix Web v3.0](https://www.reddit.com/r/rust/comments/iqq8k9/announcing_actixweb_v30/) in September, we've been very happy to see the positive response to the numerous internal safety improvements. The uptake of v3 has been swift; over 50% of crate downloads are for the newer versions.

# Notable Changes

## actix-web v3.1 &amp; v3.2

* New `#[route(...)]` multi-method macro.
* A request-local data extractor `web::ReqData`.
* Enable access to client TLS certificates and other advanced use cases where extra socket data is needed (via `HttpServer::on_connect` hook).
* An `exclude_regex` method for the Logger middleware.
* Register function for custom request-derived logging units.

&gt;[Full Changelog](https://github.com/actix/actix-web/blob/HEAD/CHANGES.md)

&amp;#x200B;

## actix-cors v0.5

On the path to stabilization, v0.5 is being considered a v1.0 beta. Includes a new way to filter allowed origins using closures and a major internal rewrite. Any feedback is appreciated.

&gt;[Full Changelog](https://github.com/actix/actix-extras/blob/HEAD/actix-cors/CHANGES.md)

&amp;#x200B;

# An Official Discord Community

&gt;Server Invite: [https://discord.gg/2Kkkc8ZtzF](https://discord.gg/2Kkkc8ZtzF)

We're still keeping Gitter around as a support channel (and I'm hopeful it will improve more quickly under the Matrix banner).
## [4][What are the bad practices in rust that should be avoided?](https://www.reddit.com/r/rust/comments/jkortm/what_are_the_bad_practices_in_rust_that_should_be/)
- url: https://www.reddit.com/r/rust/comments/jkortm/what_are_the_bad_practices_in_rust_that_should_be/
---
I am learning rust and I have read that it is not a good practice to use `unwrap()`,  

&amp;#x200B;

What are the bad practices in rust that should be avoided and why?
## [5][This Week in Veloren #91: Buffs, Alignment, and CI](https://www.reddit.com/r/rust/comments/jkp6kv/this_week_in_veloren_91_buffs_alignment_and_ci/)
- url: https://veloren.net/devblog-91/
---

## [6][http-types v2.7.0 has been released!](https://www.reddit.com/r/rust/comments/jkumqr/httptypes_v270_has_been_released/)
- url: https://github.com/http-rs/http-types/releases/tag/v2.7.0
---

## [7][For Complex Applications, Rust is as Productive as Kotlin](https://www.reddit.com/r/rust/comments/jk8648/for_complex_applications_rust_is_as_productive_as/)
- url: https://ferrous-systems.com/blog/rust-as-productive-as-kotlin/
---

## [8][IntelliJ Rust: New Functionality for Cargo Features](https://www.reddit.com/r/rust/comments/jka6bn/intellij_rust_new_functionality_for_cargo_features/)
- url: https://blog.jetbrains.com/clion/2020/10/intellij-rust-new-functionality-for-cargo-features/
---

## [9][Learning C/C++ as a Rustacean?](https://www.reddit.com/r/rust/comments/jkw8dn/learning_cc_as_a_rustacean/)
- url: https://www.reddit.com/r/rust/comments/jkw8dn/learning_cc_as_a_rustacean/
---
Hello everyone, I would like to hear from those of you who started Rust as your first low-level programming language and then move on to learn C/C++.

If I search for "C"/"C++" in r/rust, almost all threads are about comparisons between C/C++ and Rust and switching *from C/C++ to Rust*. However, my situation is different. I chose to learn Rust as my first low-level language simply because it has great and easily accessible documentation and tutorials, as well as awesome dependency management. However, soon I realized that to become a qualified computational biologist (that's what I plan to pursue in PhD, although my current major is biochemistry which does not require much programming) I have to learn C/C++ because almost all currently available bioinformatics libraries are written in C/++ (either as standalone C/C++ libraries or Python/R extensions). 

I need to learn not only the C/C++ language itself but also how it is used in reality (managing real projects, understanding different ways by which other people manage their projects, idiomatic practices, understanding code that was idiomatic in old standards but not idiomatic anymore in C++20, and understanding why people are re-inventing their own `vec`, `string` and even `bool`).

In addition, I would like to learn best practices for re-writing C/C++ code in Rust  and for replacing part of a complex C/C++ project with Rust (if there are any).

So what's the *least painful* way to learn all these, in November 2020?
## [10][Advanced Cargo [features] Usage | blog.turbo.fish](https://www.reddit.com/r/rust/comments/jkn4gu/advanced_cargo_features_usage_blogturbofish/)
- url: https://blog.turbo.fish/cargo-features/
---

## [11][Why did you start learning Rust?](https://www.reddit.com/r/rust/comments/jkj1sn/why_did_you_start_learning_rust/)
- url: https://www.reddit.com/r/rust/comments/jkj1sn/why_did_you_start_learning_rust/
---
I'm a relatively new Rust developer, I started to learn Rust in mid-2020 and I fell in love with the language, I wanted to ask what made YOU take the decision to learn Rust?
## [12][How heavy is Rc&lt;RefCell&lt;T&gt;?](https://www.reddit.com/r/rust/comments/jkh99u/how_heavy_is_rcrefcellt/)
- url: https://www.reddit.com/r/rust/comments/jkh99u/how_heavy_is_rcrefcellt/
---
Hi all! So i am implementing algorithms for fun in Rust and i am now on graphs and similar, And for mutable data that has references everywhere the only sane way i have found is Rc&lt;RefCell&lt;T&gt;&gt;. Using this however made wonder, how heavy is this reference counting? Or is it basically irrelevant in terms of speed lost? 
