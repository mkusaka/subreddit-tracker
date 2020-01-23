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

## [3][[ANN] Rust conference in Wrocław (Poland)](https://www.reddit.com/r/rust/comments/esqw30/ann_rust_conference_in_wrocław_poland/)
- url: https://www.reddit.com/r/rust/comments/esqw30/ann_rust_conference_in_wrocław_poland/
---
I am happy to announce, that we are planning the Rust conference in Wrocław, at the end of this year.   


Visit our [website](https://rusty-days.org/), follow us on [Twitter](https://mobile.twitter.com/rdconf), and see you in Poland!  


We will provide more details soon :)
## [4][[ANN] cargo-lock v4.0: list packages, show dependency trees, and translate formats for Cargo.lock files](https://www.reddit.com/r/rust/comments/esl6xu/ann_cargolock_v40_list_packages_show_dependency/)
- url: https://www.reddit.com/r/rust/comments/esl6xu/ann_cargolock_v40_list_packages_show_dependency/
---
The `cargo-lock` crate is a self-contained Cargo.lock parser/serializer based on `serde`. It's usable as a library, notably by the RustSec [cargo-audit](https://github.com/RustSec/cargo-audit) tool, but also as of this release supports a CLI as well.

- GitHub: https://github.com/RustSec/cargo-lock
- Docs.rs: https://docs.rs/cargo-lock/
- Crates.io: https://crates.io/crates/cargo-lock

Among other new features, the [`cargo lock translate` subcommand](https://docs.rs/cargo-lock/4.0.1/cargo_lock/#translate-convert-cargolock-files-between-the-v1-and-v2-formats) can be used to translate your existing Cargo.lock files to the [new merge-friendly V2 lockfile format](https://github.com/rust-lang/cargo/pull/7070).
## [5][Rust 2020 roadmap](https://www.reddit.com/r/rust/comments/esj03j/rust_2020_roadmap/)
- url: https://github.com/rust-lang/rfcs/pull/2857
---

## [6][Introducing faux: A traitless mocking framework](https://www.reddit.com/r/rust/comments/ese4oa/introducing_faux_a_traitless_mocking_framework/)
- url: https://nrxus.github.io/faux/
---

## [7][Rustacean Station Triple Feature: Rust for AAA Game Development; Async Foundations with `async-std`; and Powerful Concurrency Primitives with `crossbeam` [RustFest 2019 Interviews]](https://www.reddit.com/r/rust/comments/eslbuv/rustacean_station_triple_feature_rust_for_aaa/)
- url: https://rustacean-station.org/episode/011-jake-yoshua-stjepan/
---

## [8][Lessons learnt updating a library to std::future](https://www.reddit.com/r/rust/comments/eslvk1/lessons_learnt_updating_a_library_to_stdfuture/)
- url: https://cetra3.github.io/blog/mpart-async-0-3-0/
---

## [9][Post: byte-ordered streams](https://www.reddit.com/r/rust/comments/esk1z4/post_byteordered_streams/)
- url: https://blog.yoshuawuyts.com/byte-ordered-stream-parsing/
---

## [10][Debug Rust+Mynewt Firmware for PineTime on Raspberry Pi](https://www.reddit.com/r/rust/comments/esmf12/debug_rustmynewt_firmware_for_pinetime_on/)
- url: https://medium.com/@ly.lee/debug-rust-mynewt-firmware-for-pinetime-on-raspberry-pi-4b9ac2d093a9?source=friends_link&amp;sk=edb508c31e43d3ec40ecd8554f3405f6
---

## [11][If you see a rust crate you like, please donate if you can!](https://www.reddit.com/r/rust/comments/esg40v/if_you_see_a_rust_crate_you_like_please_donate_if/)
- url: https://www.reddit.com/r/rust/comments/esg40v/if_you_see_a_rust_crate_you_like_please_donate_if/
---
This obviously has the benefit of making you feel nice because you've just done a good deed, but it also supports open-source developers. Most of them are probably doing alright without your donation, but I can speak from experience that someone showing they appreciate your hard work by throwing a couple bucks your way makes the whole experience of open-source development feel 10x more fun, so even from a self-interested perspective if you donate to projects you like you're very likely to motivate those developers to dream big instead of abandoning them.

There are a few weird hairy issues with donating to crate authors, like how you might miss out on donating to their dependencies, but I think we can all agree that some donating to show our appreciation is better than none at all!

But crate authors, please make it easy for us to give you money! I wanted to donate to support [Warp](https://github.com/seanmonstar/warp), but the author doesn't have a sponsorship button set up on their github repo (You can put any link inside the sponsorship button, it doesn't have to be github sponsors). So I went to the author's github profile page, found the link to their website, clicked the "donate" button to get to [this page](https://seanmonstar.com/donate), and clicked the paypal button only to see that **I can't actually donate because their PayPal account has been closed**. Compare this to [Iced](https://github.com/hecrj/iced), where I just click the "Sponsor" button, go to their [Ko Fi page](https://ko-fi.com/), and I easily donate straight away! (which I did). Github Sponsors is free and they match everyone's donations up to $5k I think, so if you want to set up a monthly subscription thing that seems like a good option, and I think Ko-Fi and Patreon are also free as well (although they might take transaction fees, I'm not sure). The one thing I don't like about Github Sponsors is that everyone seems to go straight from the $1/month option straight to $100/month. I'd happily throw a bit more than $1 at [the maintainer of serde](https://github.com/sponsors/dtolnay) every month but $100 is a bit too much. Also, the [main serde repo](https://github.com/serde-rs/serde) also has the issue of there being no obvious way to donate to support its development! 

I hope this inspires everyone who can afford it to think about throwing $5 or so at whichever crates they like the most, and crate authors to make it easy for us :). The Rust community is great partially because of its great libraries and motivated developers, lets support that!
## [12][Trouble getting used to type inference](https://www.reddit.com/r/rust/comments/esp8se/trouble_getting_used_to_type_inference/)
- url: https://www.reddit.com/r/rust/comments/esp8se/trouble_getting_used_to_type_inference/
---
Does anyone else have trouble quickly reading rust code, specifically because type inference? I am coming from a heavy java background and have roughly 30 hours of experience reading/writing rust, so I might just need someone to tell me "don't worry, you'll get used to it." I think I am one of the very few people who really appreciates java verbosity (that's right, I said it).  And when I see non-explicitly declared variables, I get a similar uncomfortable feeling as when I have to read/write python code.

r/rust seems like a very friendly place, so I guess I'm just wondering if anyone has similar experiences to share. Cheers!
