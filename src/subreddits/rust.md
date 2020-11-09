# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (46/2020)!](https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (46/2020)?](https://www.reddit.com/r/rust/comments/jqrmqn/whats_everyone_working_on_this_week_462020/)
- url: https://www.reddit.com/r/rust/comments/jqrmqn/whats_everyone_working_on_this_week_462020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-46-2020/51226?u=llogiq)!
## [3][Why Rust is the Future of Game Development](https://www.reddit.com/r/rust/comments/jqux1s/why_rust_is_the_future_of_game_development/)
- url: https://thefuntastic.com/blog/why-rust-is-the-future-game-dev
---

## [4][rust-analyzer changelog #50](https://www.reddit.com/r/rust/comments/jqwfoj/rustanalyzer_changelog_50/)
- url: https://rust-analyzer.github.io/thisweek/2020/11/09/changelog-50.html
---

## [5][Pijul: Towards 1.0](https://www.reddit.com/r/rust/comments/jquu9r/pijul_towards_10/)
- url: https://pijul.org/posts/2020-11-07-towards-1.0/
---

## [6][I wrote a programming language/interpreter as a Rust library, so that you can easily integrate it into your Rust projects and use it to allow user-definable logic.](https://www.reddit.com/r/rust/comments/jqms89/i_wrote_a_programming_languageinterpreter_as_a/)
- url: https://crates.io/crates/xbasic
---

## [7][Benchmarking LLVM Polly in Rust](https://www.reddit.com/r/rust/comments/jqw0ka/benchmarking_llvm_polly_in_rust/)
- url: https://www.reddit.com/r/rust/comments/jqw0ka/benchmarking_llvm_polly_in_rust/
---
So Polly. It's LLVM's loop fusion and data locality optimizer. LLVM does not include it by default and the [last attempt](https://github.com/rust-lang/rust/pull/51061) to add it to Rust was 2 years ago so I took a shot at it.

[Ported](https://github.com/JRF63/polybench-rs) the benchmark that Polly uses (PolyBench/C). Performance gains are massive in some tests with deep nested loops but there are also slowdowns.

The [PR](https://github.com/rust-lang/rust/pull/78566) on this isn't merged yet. But if anyone cares to try, it's a one-liner change (disregarding the conditional compilation).
## [8][normpath - More reliable path manipulation](https://www.reddit.com/r/rust/comments/jqx8hf/normpath_more_reliable_path_manipulation/)
- url: https://www.reddit.com/r/rust/comments/jqx8hf/normpath_more_reliable_path_manipulation/
---
Hi everyone,

I just released [normpath](https://crates.io/crates/normpath) \- a crate meant to make correct path manipulation easier. It lets you to avoid doing [canonicalization](https://doc.rust-lang.org/std/fs/fn.canonicalize.html) that is often unnecessary on Windows and unreliable: [\#45067](https://github.com/rust-lang/rust/issues/45067), [\#48249](https://github.com/rust-lang/rust/issues/48249), [\#52440](https://github.com/rust-lang/rust/issues/52440), [\#55812](https://github.com/rust-lang/rust/issues/55812), [\#58613](https://github.com/rust-lang/rust/issues/58613), [\#59107](https://github.com/rust-lang/rust/issues/59107), [\#74327](https://github.com/rust-lang/rust/issues/74327). Canonicalization can cause errors for some valid paths (e.g., shared partitions) and always returns a verbatim path, which is usually unexpected. Normalization is probably what you want instead, which this crate provides with [`PathExt::normalize`](https://docs.rs/normpath/0.1.0/normpath/trait.PathExt.html#tymethod.normalize).

It also defines [`BasePath`](https://docs.rs/normpath/0.1.0/normpath/struct.BasePath.html) and [`BasePathBuf`](https://docs.rs/normpath/0.1.0/normpath/struct.BasePathBuf.html), which improve the API of [`Path`](https://doc.rust-lang.org/std/path/struct.Path.html) and [`PathBuf`](https://doc.rust-lang.org/std/path/struct.PathBuf.html) and handle more edge cases when joining paths. For example, you can use them to join verbatim paths without [custom workarounds](https://github.com/rust-lang/cargo/blob/master/src/cargo/util/paths.rs#L61-L86).

[Documentation](https://docs.rs/normpath) | [Repository](https://github.com/dylni/normpath) | [crates.io](https://crates.io/crates/normpath)
## [9][rsys - CLI tool with graphs and more](https://www.reddit.com/r/rust/comments/jqfytu/rsys_cli_tool_with_graphs_and_more/)
- url: https://www.reddit.com/r/rust/comments/jqfytu/rsys_cli_tool_with_graphs_and_more/
---
Hi, rustacenas!

I decided to write a CLI tool with some nice graphs as well as an easy api for gathering system information (dumping info like cpu, network, memory, processes... as JSON, yaml). Maybe some of you might find it interesting or have suggestions on what to improve.

Here is the [repository](https://github.com/wojciechkepka/rsys-cli).

Here is a demo of the graph functionality:

https://i.redd.it/ga4o29p602y51.gif

edit:  
I should probably add that in the current state it's probably only usable on linux.  

## [10][xsv - a command-line toolkit (for CSV data) written in Rust is crushing it in performance](https://www.reddit.com/r/rust/comments/jqiwpj/xsv_a_commandline_toolkit_for_csv_data_written_in/)
- url: https://ezzeddinabdullah.medium.com/how-to-clean-csv-data-at-the-command-line-part-2-207215881c34
---

## [11][Using rust to communicate with gpio. Is that too advanced for a beginner?](https://www.reddit.com/r/rust/comments/jqv3cm/using_rust_to_communicate_with_gpio_is_that_too/)
- url: https://www.reddit.com/r/rust/comments/jqv3cm/using_rust_to_communicate_with_gpio_is_that_too/
---
Recently started with rust and I love it, it's the first language which I (think) I understand and want to stick to it. I also really enjoy making things with a raspberry pi but hate python so I wanted to use rust on the pi.

I know there's already a library out there for gpio and the pi however I would like to put something together myself.

I'm at that awkward stage where I understand the basics of the language but not well enough to implement any real projects. I've made some console apps which have been very helpful
## [12][Rust type checking?](https://www.reddit.com/r/rust/comments/jqm0rv/rust_type_checking/)
- url: https://www.reddit.com/r/rust/comments/jqm0rv/rust_type_checking/
---
I started exploring rust (for one of my courses) and I was wondering how rust does type checking? I looked through official sources but couldn't find anything useful. I hope someone could help me.
