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

## [3][Taking Advantage of Auto-Vectorization in Rust](https://www.reddit.com/r/rust/comments/gkq0op/taking_advantage_of_autovectorization_in_rust/)
- url: https://nickwilcox.github.io/blog/autovec/
---

## [4][Jonathan Turner teaches Jason Turner (the C++ expert) Rust](https://www.reddit.com/r/rust/comments/gkr311/jonathan_turner_teaches_jason_turner_the_c_expert/)
- url: https://www.youtube.com/watch?v=EzQ7YIIo1rY
---

## [5][A new server-side GraphQL library.](https://www.reddit.com/r/rust/comments/gkokke/a_new_serverside_graphql_library/)
- url: https://www.reddit.com/r/rust/comments/gkokke/a_new_serverside_graphql_library/
---
`Async-graphql` is a GraphQL server-side library implemented in Rust. It is fully compatible with the GraphQL specification and most of its extensions, and offers type safety and high performance, easy to use and contains everything you need about GraphQL.

Repo: [https://github.com/async-graphql/async-graphql](https://github.com/async-graphql/async-graphql)

Examples: [https://github.com/async-graphql/examples](https://github.com/async-graphql/examples)

Benchmark: [https://github.com/async-graphql/benchmark](https://github.com/async-graphql/benchmark)

## Features

* Fully support async/await
* Type safety
* Rustfmt friendly (Procedural Macro)
* Custom scalar
* Minimal overhead
* Easy integration (hyper, actix\_web, tide ...)
* Upload files (Multipart request)
* Subscription (WebSocket transport)
* Custom extension
* Apollo Tracing extension
* Limit query complexity/depth
* Error Extensions
* Apollo Federation

I am a faithful believer of rust. I think rust can change the world! The improvement of ecosystem requires everyone to work together, this is my first small contribution to the rust world, and I am making it better every week and every day.

I hope it will become the best GraphQL library for rust, and I have been working hard!
## [6][Developing a small Windows CLI app using WinRT APIs, with full of rust-analyzer support and quick recompiles](https://www.reddit.com/r/rust/comments/gkiha8/developing_a_small_windows_cli_app_using_winrt/)
- url: https://youtu.be/oLohY8-nKcA
---

## [7][KeePwn: check a KeePass database passwords against "Have I Been Pwned?"](https://www.reddit.com/r/rust/comments/gktqoc/keepwn_check_a_keepass_database_passwords_against/)
- url: https://www.reddit.com/r/rust/comments/gktqoc/keepwn_check_a_keepass_database_passwords_against/
---
Link to KeePwn: https://github.com/gsurrel/keepwn

I wanted to build a small tool that checks a KeePass database against *Have I Been Pwned?* for a long time, but this morning, I got the motivation to finally do it.

Picking up the crates that I felt appropriate, I glued everything together, only to discover late that the [checkpwn](https://crates.io/crates/checkpwn) one is actually a binary crate, rather than a library. For now, I copied the code in a submodule to use it the way I want, but ultimately, it's better to de-couple everything.

Here is the high-level overview of the software: KeePwn opens the database (as read-only), iterate over all entries, retrieve the password, hashes it, get the first chars of the hash to query the HIBP password API, check if the remaining of the hash is present in the API's response, and prints to the terminal the result.

Slow? Yes, because there is a rate-limiting from the API, enforcing a 1.5s delay between requests.

It is not strictly speaking my first Rust code, but maybe the first one that can be interesting to a bigger audience. It's been done in almost exactly 3h of time. Any comments, tips, or debate is welcome :)

*Note:* I've not reviewed the code at all of the used crates, but I *assume* they are safe to use. Yes, it's dealing with passwords and that's sensitive, but it should be OK.
## [8][Five Years of Rust | Rust Blog](https://www.reddit.com/r/rust/comments/gk77zt/five_years_of_rust_rust_blog/)
- url: https://blog.rust-lang.org/2020/05/15/five-years-of-rust.html
---

## [9][cargo-wipe](https://www.reddit.com/r/rust/comments/gktc58/cargowipe/)
- url: https://github.com/mihai-dinculescu/cargo-wipe
---

## [10][State of Web Routing in Rust](https://www.reddit.com/r/rust/comments/gks14y/state_of_web_routing_in_rust/)
- url: https://pksunkara.com/posts/state-of-routing-in-rust/
---

## [11][Tour of Rust: Chapter 3 on data structures released, Spanish translations kickstarted, looking for translator](https://www.reddit.com/r/rust/comments/gku5ds/tour_of_rust_chapter_3_on_data_structures/)
- url: https://tourofrust.com/22_en.html
---

## [12][Function pointer type inference: Is this a bug?](https://www.reddit.com/r/rust/comments/gkikzt/function_pointer_type_inference_is_this_a_bug/)
- url: https://www.reddit.com/r/rust/comments/gkikzt/function_pointer_type_inference_is_this_a_bug/
---
When using total type inference for function pointer parameters, the borrow checker incorrectly evaluates the lifetime of reference parameters, artificially extending their purported lifetimes.

https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=6b7520cf86fd5606987dfb5b7435d2b9

Is this a real bug in type inference or the borrow checker, or just a quirk of the type system?
