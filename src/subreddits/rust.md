# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (30/2020)!](https://www.reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 348](https://www.reddit.com/r/rust/comments/hvjf4i/this_week_in_rust_348/)
- url: https://this-week-in-rust.org/blog/2020/07/21/this-week-in-rust-348/
---

## [3][Serious bug in Rust 1.45 stable](https://www.reddit.com/r/rust/comments/hxlj68/serious_bug_in_rust_145_stable/)
- url: https://www.reddit.com/r/rust/comments/hxlj68/serious_bug_in_rust_145_stable/
---

https://github.com/rust-lang/rust/issues/74739

It was found via a [stackoverflow question](https://stackoverflow.com/q/63087217/493729).

Almost any code could be affected by something like this as it's quite subtle to spot. You should go back to Rust 1.44 if you are doing anything remotely important.

The bug doesn't appear to be present in the most recently Nightly, so I hope it can be fixed quickly. It's still a bit scary that a bug this serious could get past the tests.
## [4][Interactive Chord Diagrams with Rust](https://www.reddit.com/r/rust/comments/hxj9go/interactive_chord_diagrams_with_rust/)
- url: https://shahinrostami.com/posts/programming/rust-notebooks/chord-diagrams/
---

## [5][Langcraft: The LLVM target for Minecraft you've never wanted](https://www.reddit.com/r/rust/comments/hx3we0/langcraft_the_llvm_target_for_minecraft_youve/)
- url: https://github.com/SuperTails/langcraft
---

## [6][Sizedness in Rust](https://www.reddit.com/r/rust/comments/hxips7/sizedness_in_rust/)
- url: https://github.com/pretzelhammer/rust-blog/blob/master/posts/sizedness-in-rust.md
---

## [7][Under the hood of Linkerd's state-of-the-art Rust proxy](https://www.reddit.com/r/rust/comments/hx5dco/under_the_hood_of_linkerds_stateoftheart_rust/)
- url: https://linkerd.io/2020/07/23/under-the-hood-of-linkerds-state-of-the-art-rust-proxy-linkerd2-proxy/
---

## [8][Plotly.rs has now Jupyter support!](https://www.reddit.com/r/rust/comments/hxmni7/plotlyrs_has_now_jupyter_support/)
- url: https://igiagkiozis.github.io/plotly/content/fundamentals/jupyter_support.html
---

## [9][Rust nightly compiler supports AVR](https://www.reddit.com/r/rust/comments/hwyii4/rust_nightly_compiler_supports_avr/)
- url: https://avr-rust.com/
---

## [10][New Rust Project: A free open source accounting software for the command line, inspired by ledger and c](https://www.reddit.com/r/rust/comments/hxb3hn/new_rust_project_a_free_open_source_accounting/)
- url: https://www.reddit.com/r/rust/comments/hxb3hn/new_rust_project_a_free_open_source_accounting/
---
I would like to announce that I started a ledger implementation in Rust at [github.com/rudolfschmidt/acc](https://github.com/rudolfschmidt/acc) and would love to see some traction, feedback, and requests for features that are widely needed and not provided yet.
## [11][Multi Layer Archive - A pure rust encrypted and compressed archive file format](https://www.reddit.com/r/rust/comments/hx58tl/multi_layer_archive_a_pure_rust_encrypted_and/)
- url: https://github.com/ANSSI-FR/MLA
---

## [12][Rust Analyzer or Rust Extension for VSCode?](https://www.reddit.com/r/rust/comments/hxe87f/rust_analyzer_or_rust_extension_for_vscode/)
- url: https://www.reddit.com/r/rust/comments/hxe87f/rust_analyzer_or_rust_extension_for_vscode/
---
Hi guys!

I have been using the `rust-analyzer` for VSCode from some time now, and have found that its pretty expensive to use.

Anytime I add a dependency, this automatically calls `cargo check` and If I disable the option to auto load dependencies, then, the yellow warning in the status bar never goes away.

So at the end of the day, when it comes to autocomplete it works pretty good, Im pretty impressed with the intellisense this extension provides is so good compared to the official extension, the issue comes when trying to remove so many calls to the `cargo check` command.

Any other experiencing the same?

Thanks!
