# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (18/2020)!](https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g4nu6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (18/2020)?](https://www.reddit.com/r/rust/comments/g9a7d1/whats_everyone_working_on_this_week_182020/)
- url: https://www.reddit.com/r/rust/comments/g9a7d1/whats_everyone_working_on_this_week_182020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-18-2020/41671?u=llogiq)!
## [3][List of Rust streamers](https://www.reddit.com/r/rust/comments/g9ibs9/list_of_rust_streamers/)
- url: https://github.com/jamesmunns/awesome-rust-streaming
---

## [4][Rust 1.0 T-Shirt](https://www.reddit.com/r/rust/comments/g9lnlj/rust_10_tshirt/)
- url: https://i.redd.it/u078didrpjv41.png
---

## [5][One-liner to (correctly) list all uses of`unsafe` in the entire crate graph](https://www.reddit.com/r/rust/comments/g9mw57/oneliner_to_correctly_list_all_uses_ofunsafe_in/)
- url: https://www.reddit.com/r/rust/comments/g9mw57/oneliner_to_correctly_list_all_uses_ofunsafe_in/
---
`env RUSTFLAGS="-Funsafe-code --cap-lints=warn" cargo check -vv --message-format=json-diagnostic-rendered-ansi 2&gt;/dev/null | grep '^\{' | jq -r 'select(.message?.code?.code? == "unsafe_code").message?.rendered?'`

You may also want to use a temporary `CARGO_TARGET_DIR` that you remove afterwards, to both prevent Cargo from having to rebuild all your deps when using normal `cargo check`, and to re-print all the warnings every time you run the command.

And/or wrap this in something caching the results (presumably some of the Cargo messages tell you which crates are not getting rebuilt because they haven't changed).

A Rust program like `cargo-geiger` could use this method with very little effort, as well, while still retaining the ability to collect various kinds of statistics.

While I have no use for this (at least at the moment, I do however use `#![forbid(unsafe_code)]` everywhere I can in my own projects), I suspect there are other people who do.

Please stay safe ^(and avoid ad-hoc analyses of source code)

(PS: you can use this technique with arbitrary lints, including clippy ones! so the easiest way to have a custom analysis is to add it to clippy as an off-by-default lint then run a command like this)
## [6][First official release of rust-analyzer](https://www.reddit.com/r/rust/comments/g8zzdy/first_official_release_of_rustanalyzer/)
- url: https://rust-analyzer.github.io/blog/2020/04/20/first-release.html
---

## [7][filter_map() slower then filter() then map()](https://www.reddit.com/r/rust/comments/g9lug6/filter_map_slower_then_filter_then_map/)
- url: https://www.reddit.com/r/rust/comments/g9lug6/filter_map_slower_then_filter_then_map/
---
Anyone have comments on the following problem I have outlined in a stackoverflow post?

https://stackoverflow.com/questions/61471978/is-there-a-way-of-making-filter-map-quicker-then-filter-then-map-currentl
## [8][smol (@stjepang) - A small and fast async runtime for Rust](https://www.reddit.com/r/rust/comments/g917ad/smol_stjepang_a_small_and_fast_async_runtime_for/)
- url: https://github.com/stjepang/smol
---

## [9][Ranges: a generic addition to or replacement of core ranges](https://www.reddit.com/r/rust/comments/g9mjx5/ranges_a_generic_addition_to_or_replacement_of/)
- url: https://www.reddit.com/r/rust/comments/g9mjx5/ranges_a_generic_addition_to_or_replacement_of/
---
Hello!

I've been working on this for some time now and I think it's time to spread the word and get some feedback.

&amp;#x200B;

Ranges' most important component is the generic alternative to core/std ranges, aptly named GenericRange. It works mostly the same as the ranges you're used to, but has two requirements to make the generic part work: The start is always less or equal to the end and the type being stored implements the Domain trait, which requires Ord. I think these requirements are reasonable, and the goal is to implement Domain for as many core/std types as possible anyway, so for most use cases it can be used directly without any additional code. By the way, "generic" is a bit of a wordplay. It is of course generic like the already existing ranges are, but it is also "range-generic" and can store all range variants existing in core and even the unimplemented ones like open ranges.

Using the aforementioned GenericRange, Ranges (as in RangeSet) is a vector-backed interval/range set which stores inserted ranges in the most space efficient way possible by unionizing inserted with existing ones. The plan is to also build multi-dimensional versions of this so it works, for example, in 3D space and above.

&amp;#x200B;

There's also 3 feature gates so far:

arbitrary - allowing fuzzers and alike to create random ranges and rangesets which still uphold the main guarantees

noisy\_float - adds implementations of Domain for all 4 noisy\_float variants so now you can have float ranges!

num-bigint - adds implementations of Domain for arbitrary large integers. These showcase truly unbound domains.

&amp;#x200B;

Generally, I plan to make this crate as compatible as possible with core/std to integrate it seamlessly. Maybe even cause some discussion about replacing the existing implementations with it or something close. Additionally I want to support as many custom types as possible (like I already do with noisy-float and bigint).

&amp;#x200B;

Apart from that, please have fun taking the documentation and underlying code to bits - I hope to improve it even more with the feedback I'll get here. Also, this is my first ever release and accompanied post on reddit, so if there's anything I can change to make this even more interesting for potential users, please don't hesitate to tell me :)

&amp;#x200B;

[https://docs.rs/ranges](https://docs.rs/ranges)

[https://lib.rs/crates/ranges](https://lib.rs/crates/ranges)

[https://crates.io/crates/ranges](https://crates.io/crates/ranges)
## [10][A lightweight implementation of 2048, leveraging the expectimax algorithm](https://www.reddit.com/r/rust/comments/g9mjf1/a_lightweight_implementation_of_2048_leveraging/)
- url: https://github.com/adrienball/2048-rs
---

## [11][rust-analyzer Changelog #22](https://www.reddit.com/r/rust/comments/g94g76/rustanalyzer_changelog_22/)
- url: https://rust-analyzer.github.io/thisweek/2020/04/27/changelog-22.html
---

## [12][flatzinc parser 0.1](https://www.reddit.com/r/rust/comments/g9n4du/flatzinc_parser_01/)
- url: https://www.reddit.com/r/rust/comments/g9n4du/flatzinc_parser_01/
---
# A parser for the [FlatZinc](https://www.minizinc.org/doc-2.4.1/en/fzn-spec.html#specification-of-flatzinc) modelling language

[https://github.com/sthiele/flatzinc-parser](https://github.com/sthiele/flatzinc-parser)

I made this parser for one of my projects and maybe it can be useful for others. Also thanks to everybody working on nom.
