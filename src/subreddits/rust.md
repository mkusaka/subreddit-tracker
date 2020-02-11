# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (7/2020)!](https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/)
- url: https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (7/2020)?](https://www.reddit.com/r/rust/comments/f1uej1/whats_everyone_working_on_this_week_72020/)
- url: https://www.reddit.com/r/rust/comments/f1uej1/whats_everyone_working_on_this_week_72020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-7-2020/38078?u=llogiq)!
## [3][Scaling back my involvement in Rust](https://www.reddit.com/r/rust/comments/f268p6/scaling_back_my_involvement_in_rust/)
- url: https://internals.rust-lang.org/t/scaling-back-my-involvement-in-rust/11754
---

## [4][Let's Be Real About Dependencies](https://www.reddit.com/r/rust/comments/f1xiub/lets_be_real_about_dependencies/)
- url: https://wiki.alopex.li/LetsBeRealAboutDependencies
---

## [5][Announcing: jlrs. Call arbitrary Julia code from Rust and share data between the two languages](https://www.reddit.com/r/rust/comments/f1v00r/announcing_jlrs_call_arbitrary_julia_code_from/)
- url: https://www.reddit.com/r/rust/comments/f1v00r/announcing_jlrs_call_arbitrary_julia_code_from/
---
After quite some effort I'm finally ready to release first version of `jlrs`, the [Julia](https://julialang.org) bindings I've been working on for the past few weeks. 

This first version brings many nice features. For example:

- Primitive data and strings can be copied from Rust to Julia.

- N-dimensional arrays containing primitive data can be created; their contents can either be completely managed by Julia, take ownership of data from Rust, or mutably borrow data from Rust.

- All of these can be copied back from Julia to Rust.

- You can acquire handles to arbitrary modules, globals and functions, and call the latter.

- You can include and call your own Julia code.

- If a Julia function returns another function, you can call it.

- You can only use Julia data that is guaranteed to be protected from garbage collection.

Of course, not everything is great yet. Some of it is due to limitations of Julia itself: access to the runtime is single-threaded, for example, and that thread will block when it's calling Julia code; you can't simply "pause" a long-running computation to quickly do some smaller amounts of work. Thread support in Julia itself is still in an experimental phase. Most annoyingly, the documentation is not available on docs.rs at this time because the bindings can't be compiled. If anyone could help me fix that or point me in the right direction, your help would be greatly appreciated (and yes, my failed attempts to fix it are the reason several version have been released already... sorry about that).

Luckily, the readme in the Github repository should provide enough details to build this crate, `cargo doc` will work if you're able to build it.

You probably shouldn't use this crate to do many one-off or short-and-simple computations; in the first case, the fact that Julia code is JIT-compiled will bite you, in the latter the overhead of this crate will. Rather, you should use it to call sufficiently complex functions that you will call regularly, or to use functionality that's not available in Rust yet. 

Another possible use case is one that actually motivated me to keep writing this. I work as (mostly) a software engineer in vision and robotics, a recurring task is to convert some specific algorithm written in Python to C++ for both performance reasons and so we can integrate it into the software that actually powers the products we sell. This is an error-prone process (rewriting software always is), takes a lot of time, and you now have two versions of some code to maintain. If I were able to just call the "original" code with great performance, we'd be rid of the split codebase and save a lot of time and money in the process.

[crates.io](https://crates.io/crates/jlrs)

[Github](https://github.com/Taaitaaiger/jlrs)
## [6][rust-analyzer Changelog #11](https://www.reddit.com/r/rust/comments/f1tnem/rustanalyzer_changelog_11/)
- url: https://rust-analyzer.github.io/thisweek/2020/02/10/changelog-11.html
---

## [7][A primer to Rust Async](https://www.reddit.com/r/rust/comments/f1y410/a_primer_to_rust_async/)
- url: https://omarabid.com/async-rust
---

## [8][Quantitative data on the safety of Rust](https://www.reddit.com/r/rust/comments/f1ynel/quantitative_data_on_the_safety_of_rust/)
- url: https://www.reddit.com/r/rust/comments/f1ynel/quantitative_data_on_the_safety_of_rust/
---
While the safety benefits of Rust make a lot of sense intuitively, the presence of `unsafe` makes that intuition less clear-cut. As far as I'm aware there is little hard data on how real-world Rust code performs in terms of security compared to other languages. I've realized that I might just contribute a quantitative data point.

[Fuzzing](https://en.wikipedia.org/wiki/Fuzzing) is quite common in the Rust ecosystem nowadays, largely thanks to the [best-of-breed tooling](https://github.com/rust-fuzz/cargo-fuzz) we have at our disposal. There is also a [trophy case](https://github.com/rust-fuzz/trophy-case) of real-world bugs found in Rust code via fuzzing. It lists ~200 bugs as of commit [17982a8](https://github.com/rust-fuzz/trophy-case/commit/17982a86c7d4696aecfa96e2c271dbdc28318d30), out of which only 5 are security vulnerabilities - or 2.5%. Contrast this with the results from Google's [OSS-fuzz](https://github.com/google/oss-fuzz), which fuzzes high-profile C and C++ libraries: out of [15807](https://bugs.chromium.org/p/oss-fuzz/issues/list?sort=proj&amp;colspec=ID%20Type%20Component%20Status%20Proj%20Reported%20Owner%20Summary&amp;num=100&amp;q=-status%3AWontFix%2CDuplicate%20-Infra%20type%3Abug&amp;can=1)  bugs discovered [3600](https://bugs.chromium.org/p/oss-fuzz/issues/list?sort=proj&amp;colspec=ID%20Type%20Component%20Status%20Proj%20Reported%20Owner%20Summary&amp;num=100&amp;q=-status%3AWontFix%2CDuplicate%20-Infra%20type%3Abug-security&amp;can=1) are security issues. That's a whopping 22%!

OSS-fuzz and Rust ecosystem use the exact same fuzzing backends (afl, libfuzzer, honggfuzz) so these results should be directly comparable. I'm not sure how representative a sample size of 200 is, so I'd appreciate statistical analysis on this data.

Note that this approach only counts the bugs that actually made it into a compiled binary, so it does **not** account for bugs prevented statically. For example, iterators make out-of-bounds accesses impossible, `Option&lt;T&gt;` and `&amp;T` make null pointer dereferences impossible and lifetime analysis makes use-after-frees impossible. All of these bugs were eliminated before the fuzzer could even get to them, so I expect the security defect rate for Rust code to be even lower than these numbers suggest.

**TL;DR:** out of bugs found by the exact same tooling in C/C++ 22% of them pose a security issue while in Rust it's 2.5%. That is about an order of magnitude difference. Actual memory safety defect rates in Rust should be even lower because some bugs are prevented statically and don't make it into this statistic.

This only applies to memory safety bugs, which account for about 70% of all security bugs [according to Microsoft](https://www.zdnet.com/article/microsoft-70-percent-of-all-security-bugs-are-memory-safety-issues/). Mozilla had also [independently arrived to the same estimate](https://hacks.mozilla.org/2019/02/rewriting-a-browser-component-in-rust/).
## [9][Copyright implications of brute forcing all 12-tone major melodies in approximately 2.5 TB.](https://www.reddit.com/r/rust/comments/f1ukv0/copyright_implications_of_brute_forcing_all/)
- url: https://youtu.be/sfXn_ecH5Rw
---

## [10][Debugging Rust in VSCode in 2020](https://www.reddit.com/r/rust/comments/f1qsx9/debugging_rust_in_vscode_in_2020/)
- url: https://jason-williams.co.uk/debugging-rust-in-vscode
---

## [11][Redelete: delete your reddit comments/submissions with optional filters to skip certain ones. Fun project for me to learn async/await, rust, oauth, and testing.](https://www.reddit.com/r/rust/comments/f1zcu4/redelete_delete_your_reddit_commentssubmissions/)
- url: https://github.com/ardeaf/redelete
---

## [12][Lib.rs - crates.io on steroids](https://www.reddit.com/r/rust/comments/f1mr4h/librs_cratesio_on_steroids/)
- url: https://lib.rs/
---

