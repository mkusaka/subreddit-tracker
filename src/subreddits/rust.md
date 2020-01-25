# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (4/2020)!](https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/)
- url: https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 322](https://www.reddit.com/r/rust/comments/esbf8h/this_week_in_rust_322/)
- url: https://this-week-in-rust.org/blog/2020/01/21/this-week-in-rust-322/
---

## [3][TIL that format! macro has a bunch of features](https://www.reddit.com/r/rust/comments/etgq74/til_that_format_macro_has_a_bunch_of_features/)
- url: https://doc.rust-lang.org/std/fmt/#usage
---

## [4][Build your own block_on()](https://www.reddit.com/r/rust/comments/etivdh/build_your_own_block_on/)
- url: https://stjepang.github.io/2020/01/25/build-your-own-block-on.html
---

## [5][To learn rust I wrote my own HTTP server and JSON parser from scratch. Looking for feedback on how to make it more idiomatic](https://www.reddit.com/r/rust/comments/ethobw/to_learn_rust_i_wrote_my_own_http_server_and_json/)
- url: https://www.reddit.com/r/rust/comments/ethobw/to_learn_rust_i_wrote_my_own_http_server_and_json/
---
Hi everyone ! Long time reddit lurker. I've been wanting to learn rust for a while now, and finally got around to actually writing some code. In hacker fashion, I wrote this project to scratch an itch: I accidentally used up all of the bandwidth on my cellphone for the month and no longer had the means of finding out when the next bus would pass by a stop. Public transport in my city have a website that shows this in real time. I used this as an excuse to code something that reads the schedules from their unofficial REST API and send me the info via SMS using my VoIP's REST API. 

Code can be found here: https://github.com/gelendir/rtcsms

It's not perfect, but i'm pretty glad I got it working for real. Since this the first time i've done any rust i'd appreciate any feedback on how to make the code more idiomatic, or some answers to the following questions:

* I was surprised by how much boilerplate I had to write for creating and wrapping custom errors. I imagine this might be easier with the use of a crate, but would there be a simpler approach to my error handling ?

* I'd like to reduce the friction caused by using `String` and `&amp;str` everywhere.  I've mainly used `String`s because it seemed easier for fixing compiler errors, but because of that i've had to sprinkle a lot of `.to_string()`. I haven't taken the time yet to go over the chapter about lifetimes in the rust book , would that help ?

* Since i'm parsing HTTP requests myself i'm splitting strings and having to use `.next().ok_or(...)?` a lot. An easier way of dealing with that ? Maybe i'm just to used to the ease of `foo.split(" ")[1]` in python

* Still not sure I completely understand why I need to `Box&lt;dyn ReadWrite&gt;` to use a tcp or tls socket interchangeably. Is there a way of using those traits without the `Box` ?

Thanks in advance for any feedback. It took a bit of time before I actually got anything working, but once I got past that it was a pretty interesting experience. Surprisingly, writing the JSON parser wasn't as hard as I expected and once I got all the compiler errors fixed it worked on the first try !
## [6][rweb: generate openapi spec file automatically from rust source code](https://www.reddit.com/r/rust/comments/etcev1/rweb_generate_openapi_spec_file_automatically/)
- url: https://www.reddit.com/r/rust/comments/etcev1/rweb_generate_openapi_spec_file_automatically/
---
I created a web framework for rust called [rweb][].
It gets a feature to generate openapi specification file automatically.

The link below contains examples of using [rweb][].

https://docs.rs/rweb/0.3.0-alpha.1/rweb/openapi/index.html

Repo: https://github.com/kdy1/rweb


[rweb]:https://github.com/kdy1/rweb
## [7][Announcing version 0.5.0 of typed-builder](https://www.reddit.com/r/rust/comments/etk5ka/announcing_version_050_of_typedbuilder/)
- url: https://www.reddit.com/r/rust/comments/etk5ka/announcing_version_050_of_typedbuilder/
---
Crate: https://crates.io/crates/typed-builder  
Changelog: https://github.com/idanarye/rust-typed-builder/blob/master/CHANGELOG.md#050---2020-01-25

I've just released version 0.5.0 of the typed builder proc-macro crate. This release has many breaking changes, so I figured I'd make a post to make them more visible:

* I've moved setter-specific fields to a sub-attribute - `#[builder(setter(...))]` - following the example of [derive_builder](https://crates.io/crates/derive_builder).
* The setter methods no longer automatically convert their argument type to the field's type with `into()`. If you want them to do that, you need to use `#[builder(setter(into))]`.
* I've removed the `#[builder(default_code = "...")]` setting. It was an hack for when Rust's user attributes were more limited and did not support arbitrary expressions. Now this hack is no longer needed, as one could just use `#[builder(default = ...)]`. This was the case for quite a while, but since this is a breaking release I figured now is a good time to remove it.

Thanks to all the contributes that helped with pull requests, opening issues, or even just using this crate and giving feedback!
## [8][Are there any video trackers in Rust?](https://www.reddit.com/r/rust/comments/etouk8/are_there_any_video_trackers_in_rust/)
- url: https://www.reddit.com/r/rust/comments/etouk8/are_there_any_video_trackers_in_rust/
---

## [9][Rust lang in a nutshell - 1 - introduction](https://www.reddit.com/r/rust/comments/etdyie/rust_lang_in_a_nutshell_1_introduction/)
- url: https://www.softax.pl/blog/rust-lang-in-a-nutshell-1-introduction/
---

## [10][Compiling rust+c++ with x86_64-unknown-linux-musl](https://www.reddit.com/r/rust/comments/etlakb/compiling_rustc_with_x86_64unknownlinuxmusl/)
- url: https://www.reddit.com/r/rust/comments/etlakb/compiling_rustc_with_x86_64unknownlinuxmusl/
---
When compiling deno (rust + v8) with x86\_64-unknown-linux-musl  
is there some trick (RUST\_FLAGS or CFLAGS ?) to link in static-libgcc (is that the issue?)  
[https://gist.github.com/hayd/ae5cbe81117863fff98c2f0c877f2b34](https://gist.github.com/hayd/ae5cbe81117863fff98c2f0c877f2b34)

What are some things I should try to compile to avoid the error:  
hidden symbol \`\_\_dso\_handle' isn't defined.

Note: It gets to the very last step in compilation (all the dependencies compile it's creating the bin which fails)...

I'm using [https://github.com/emk/rust-musl-builder](https://github.com/emk/rust-musl-builder) as a docker image base.
## [11][Which free/open licence is common for rust crates?](https://www.reddit.com/r/rust/comments/etmxxp/which_freeopen_licence_is_common_for_rust_crates/)
- url: https://www.reddit.com/r/rust/comments/etmxxp/which_freeopen_licence_is_common_for_rust_crates/
---
Which free/open licence is usual for rust crates? I am thinking of releasing several crates that I am working on. 

Myself, I would prefer LGPL3, but arguably I should use the same licence as everyone else to make it easier to manage. 

Let us know your preference.
## [12][Help me create an executor for my nano-coroutines but in Rust!](https://www.reddit.com/r/rust/comments/etcpin/help_me_create_an_executor_for_my_nanocoroutines/)
- url: https://www.reddit.com/r/rust/comments/etcpin/help_me_create_an_executor_for_my_nanocoroutines/
---
I am interrested in porting in Rust the optimized memory accesses algorithms driven by coroutines and experimented by Gor Nishanov presented at the CppCon 2018.

﻿﻿﻿https://youtu.be/j9tlJAqMV7U?t=732

To make it easier to understand, Gor as made experiments in using experimental C++ coroutines to interleave memory accesses. For every memory lookup the CPU need to fetch the data located at the specified pointer, this lookup can be slow (60ns) but not enough to allow the CPU to context switch to another thread, so Gor tried to make multiple memory prefetches and read all of those refetch location at one time.

This is like doing multitasking on the hardware directly, asking the memory BUS to fetch data in parallel, and make the CPU wait the least amount of time, mind blowing!

[I have made half of the work](https://github.com/Kerollmops/cml/blob/8752abc4c20941321a00d94c297b27337e6df46d/src/main.rs#L41-L62), I have an async function (I tried with Generators too) that can binary search and await (like for an I/O) the memory lookup I ordered.

Now I ask for your help to assist me to create a good executor, one that can execute multiple tasks in parallel with the least amount of overhead possible (multithreading is unecessary). I am a little bit lost with all those `Waker`, `RawWaker` and `parking`. I would like to reproduce something like the [Gor Nishanov scheduler](https://github.com/GorNishanov/await/blob/5dc5d0a7f92a49bd84ef7ca746827c8b94d5402c/2018_CppCon/src/coro_infra.h#L8-L32) and [throttler](https://github.com/GorNishanov/await/blob/5dc5d0a7f92a49bd84ef7ca746827c8b94d5402c/2018_CppCon/src/coro_infra.h#L105-L167).

I am currently using [the little executor SpaceJam talked about on Twitter](https://twitter.com/sadisticsystems/status/1219331853666193408?s=20), and that's not really good in terms of performances.

I would like to create a library that would be able to take a given amount of binary searches to do and execute all those at one time. Eventually using it in [MeiliSearch](https://github.com/meilisearch/MeiliSearch) if possible :)

EDIT: I have finally made it work with generator and [performances are good](https://github.com/Kerollmops/cml/blob/master/README.md)!
