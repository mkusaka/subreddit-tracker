# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (27/2020)!](https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (27/2020)?](https://www.reddit.com/r/rust/comments/hhv5ks/whats_everyone_working_on_this_week_272020/)
- url: https://www.reddit.com/r/rust/comments/hhv5ks/whats_everyone_working_on_this_week_272020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-27-2020/45101?u=llogiq)!
## [3][We really need to avoid npm fate](https://www.reddit.com/r/rust/comments/hiew1o/we_really_need_to_avoid_npm_fate/)
- url: https://sambleckley.com/writing/npm.html
---

## [4][Disk space and LTO improvements | Inside Rust Blog](https://www.reddit.com/r/rust/comments/hi80zh/disk_space_and_lto_improvements_inside_rust_blog/)
- url: https://blog.rust-lang.org/inside-rust/2020/06/29/lto-improvements.html
---

## [5][A Few More Reasons Rust Compiles Slowly](https://www.reddit.com/r/rust/comments/hi80xt/a_few_more_reasons_rust_compiles_slowly/)
- url: https://pingcap.com/blog/reasons-rust-compiles-slowly/
---

## [6][Platform keyboard event handling](https://www.reddit.com/r/rust/comments/hi9aq2/platform_keyboard_event_handling/)
- url: https://www.reddit.com/r/rust/comments/hi9aq2/platform_keyboard_event_handling/
---
We just landed [druid#1049](https://github.com/linebender/druid/pull/1049), which significantly reworks and improves keyboard event handling in Druid. In particular, I think it has a really good implementation of platform keyboard handling on Windows, and good enough on macOS, Gtk, and Web. Windows keyboard handling is especially tricky to get right, and in researching this I found a range of implementation quality from hopelessly broken to basically right but quite complex (particularly the browser implementations).

Another nice feature of this patch is that we're moving more towards standard types - it's using the [keyboard-types](https://crates.io/crates/keyboard-types) crate (though doing a few of our own newtypes, and having conversations with upstream about converging even more).

I should also point out that the larger task of keyboard handling is still work in progress. There are some things to fix about hotkeys (especially with non-Latin keyboard layouts), and we don't yet do IME.

The reason I am posting here, rather than just on our Zulip, is that I'd like to open some discussion about how to improve things across the ecosystem, not just in Druid. The simplest thing is for people responsible for other UI platform bindings to be aware of our implementation, study it, and adapt it.

Are there other ways the ecosystem can work together? Should we be investing in "vocabulary types" like keyboard-types? What are the best ways to do this without running into difficult coordination problems?

Discussion welcome.
## [7][rust-analyzer changelog #31](https://www.reddit.com/r/rust/comments/hhzlsq/rustanalyzer_changelog_31/)
- url: https://rust-analyzer.github.io/thisweek/2020/06/29/changelog-31.html
---

## [8][Timetill.rs | Find Rust conferences; Now updated with remote Rust conferences for 2020.](https://www.reddit.com/r/rust/comments/hijdwr/timetillrs_find_rust_conferences_now_updated_with/)
- url: https://timetill.rs/
---

## [9][Announcing const-sha1: a sha1 implementation for use in const contexts 🎉](https://www.reddit.com/r/rust/comments/hi11op/announcing_constsha1_a_sha1_implementation_for/)
- url: https://crates.io/crates/const-sha1
---

## [10][Panic no longer being caught by catch_unwind](https://www.reddit.com/r/rust/comments/hiju01/panic_no_longer_being_caught_by_catch_unwind/)
- url: https://www.reddit.com/r/rust/comments/hiju01/panic_no_longer_being_caught_by_catch_unwind/
---
I'm building a library that interfaces with a C API using FFI. The only I've come up with to handle certain interrupt polling in the host application is to `panic!` and catch the unwind. After some refactoring this is no longer working.

The `catch_unwind` was previously in the main crate, but I now moved it into the child "sys" crate. Both crates have `panic="unwind"` in cargo.toml for the release profile, but for some reason the panic seems to be compiled to an abort now instead of an unwind when I do a release build. Apart from the refactoring I also added a cargo config file to set the default target to x86\_64-pc-windows-gnu where I was previously building with `--target` manually.

In case anyone has a better idea how to handle this, the reason I am using this approach is that the host application performs a `longjmp` on said interrupt, but allows you to pass a callback to run for cleanup before resuming the `longjmp`. If the Rust callback returns normally to C then execution never returns to Rust so I never have a chance to clean up outer scope.

After catching the panic, I am returning an `Err` result, and then allowing the host to resume after the Rust stack has largely finished unwinding.
## [11][IntelliJ Rust Changelog #125](https://www.reddit.com/r/rust/comments/hi0hpg/intellij_rust_changelog_125/)
- url: https://intellij-rust.github.io/2020/06/29/changelog-125.html
---

## [12][crossfire: yet another async mpmc/mpsc based on crossbeam-channel](https://www.reddit.com/r/rust/comments/hi9vhj/crossfire_yet_another_async_mpmcmpsc_based_on/)
- url: https://www.reddit.com/r/rust/comments/hi9vhj/crossfire_yet_another_async_mpmcmpsc_based_on/
---
When last year I began to work on our storage project, there were no mpmc to support async code, and we need something fast enough to serve as backbone I/O pipeline. So I decided to build this base on crossbeam.  Lately I feel it's finally stable enough to open source, and it has been heavily tested along with our storage project. I took the name "crossfire" as a derive from crossbeam.

A few days ago @[stjepang](https://www.reddit.com/user/stjepang/) publish async-channel, it has simular goal.  But crossfire took some different aproach, since waker is wrapped from std waker and require less denpendency. Additionally crossfire support comunication between async code and threaded code. (This feature is not seen elsewhere).

You are welcome to give it a try.

[https://docs.rs/crossfire/](https://docs.rs/crossfire/)
