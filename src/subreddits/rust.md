# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (6/2020)!](https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/)
- url: https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (6/2020)?](https://www.reddit.com/r/rust/comments/ey2xrr/whats_everyone_working_on_this_week_62020/)
- url: https://www.reddit.com/r/rust/comments/ey2xrr/whats_everyone_working_on_this_week_62020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-6-2020/37776?u=llogiq)!
## [3][Announcing rust-search-extension v0.7! Search std docs, crates, error codes, and attributes in your address bar instantly!](https://www.reddit.com/r/rust/comments/eymfxu/announcing_rustsearchextension_v07_search_std/)
- url: https://www.reddit.com/r/rust/comments/eymfxu/announcing_rustsearchextension_v07_search_std/
---
Hi all! About three months ago, **rust-search-extension** got a pretty good first impression in [this reddit post](https://www.reddit.com/r/rust/comments/dp1ru6/i_published_a_handy_chromefirefox_extension_to/). Now, I bring a brand-new version of **rust-search-extension** to all of you!

Here are the major features:

## Supports search std docs

Offline mode is supported.

https://preview.redd.it/sl9pzlp8zue41.png?width=1802&amp;format=png&amp;auto=webp&amp;s=e791a5d7fc74f52a3e3bc4a57798f83ddc6b7ba8

## Supports search top 10K crates

We build top 10K crates (35+k in total on crates.io) index hourly to help you search crates instantly. You can open the popup page to sync the latest crates index automatically.

What more, you can prefix an **!** (exclamation mark) before the keyword to search crates exclusively! 

https://preview.redd.it/ntx23kmazue41.png?width=1800&amp;format=png&amp;auto=webp&amp;s=64c691ceaa09238e2cb71847e40652770e1d3e6a

## Supports search doc with type mode

You can prefix a type to get the exclusive search result. Those type keywords including:

* mod
* struct
* enum
* fn
* trait
* type
* static
* macro
* const For example, `fn:asref` to search function or `trait:iterator` to search trait. 

https://preview.redd.it/jx6ysliczue41.png?width=1796&amp;format=png&amp;auto=webp&amp;s=e4cdbdf1f1ee72fd4652efbd536acbf6afb49320

## Supports search builtin attributes

Search builtin attributes (such as `derive`, `non_exhaustive`) are also supported, by default, the result is mixed with other related docs or crates though.

If you prefer to search attribute exclusively, prefix a **#** (pound sign) before the keyword to narrow the result. 

https://preview.redd.it/lxtntewdzue41.png?width=1802&amp;format=png&amp;auto=webp&amp;s=21e7c86aa5c134524eb1b3066b47995db7e4eeda

## Supports search error codes

https://preview.redd.it/y1ca793fzue41.png?width=1800&amp;format=png&amp;auto=webp&amp;s=8cd08be25f7feb6d103a29bbbde45ae1432844bf

## Brand-new official website

[https://rust-search-extension.now.sh](https://rust-search-extension.now.sh)

## Github

**rust-search-extension** is open-sourced on Github ([https://github.com/Folyd/rust-search-extension](https://github.com/Folyd/rust-search-extension)), welcome to submit issue or PR!

# Installation

\- [Chrome Web Store](https://chrome.google.com/webstore/detail/rust-search-extension/ennpfpdlaclocpomkiablnmbppdnlhoh)

\- [Firefox](https://addons.mozilla.org/en-US/firefox/addon/rust-search-extension/)
## [4][C++ ABI: Now or Never](https://www.reddit.com/r/rust/comments/eydy51/c_abi_now_or_never/)
- url: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p1863r1.pdf
---

## [5][highlight-stderr: Run a command and highlight its stderr, preserving the order of stdout and stderr](https://www.reddit.com/r/rust/comments/eyj15s/highlightstderr_run_a_command_and_highlight_its/)
- url: https://github.com/joshtriplett/highlight-stderr
---

## [6][My first crate: const-buffer, a fixed-capacity memory buffer allocated on the stack using const generics](https://www.reddit.com/r/rust/comments/eyjoqk/my_first_crate_constbuffer_a_fixedcapacity_memory/)
- url: https://www.reddit.com/r/rust/comments/eyjoqk/my_first_crate_constbuffer_a_fixedcapacity_memory/
---
[crates.io](https://crates.io/crates/const-buffer)

[docs.rs](https://docs.rs/const-buffer/0.1.0/const_buffer/)

[GitHub](https://github.com/timvermeulen/const-buffer)

`ConstBuffer&lt;T, N&gt;` is a block of memory `N` times the size of `T` that you can read from and write to arbitrarily. It's essentially a glorified wrapper around `[MaybeUninit&lt;T&gt;; N]`.

Most const generics `Vec`\-like types that have been going around store either a `MaybeUninit&lt;[T; N]&gt;` or `[MaybeUninit&lt;T&gt;; N]` and use lots of pointer arithmetic in order to do reads and writes. This crate is an attempt to make it easier and less error-prone to write such data structures (it's very appropriate for ring buffers, too). The source code is very simple, but it improves readability significantly. The [docs.rs](https://docs.rs/const-buffer/0.1.0/const_buffer/) page has a very basic example of what using it is like.

I would appreciate any feedback you might have, and please don't hesitate to open issues or submit PRs.
## [7][Rust Analyzer Changelog #10](https://www.reddit.com/r/rust/comments/ey8twg/rust_analyzer_changelog_10/)
- url: https://rust-analyzer.github.io/thisweek/2020/02/03/changelog-10.html
---

## [8][Announcing async-std v1.5.0](https://www.reddit.com/r/rust/comments/ey98zw/announcing_asyncstd_v150/)
- url: https://github.com/async-rs/async-std/releases/tag/v1.5.0
---

## [9][dprint - JS, Typescript and JSONC formatter written in Rust](https://www.reddit.com/r/rust/comments/eyncld/dprint_js_typescript_and_jsonc_formatter_written/)
- url: https://www.reddit.com/r/rust/comments/eyncld/dprint_js_typescript_and_jsonc_formatter_written/
---
https://github.com/dsherret/dprint

Recently, [Deno replaced](https://github.com/denoland/deno/pull/3820) the prettier plugin for its built-in formatter with dprint.
## [10][win-win: semi-opinionated for window creation in Windows](https://www.reddit.com/r/rust/comments/eyco38/winwin_semiopinionated_for_window_creation_in/)
- url: https://www.reddit.com/r/rust/comments/eyco38/winwin_semiopinionated_for_window_creation_in/
---
I've been exploring soundness issues in window creation on Windows, and have found a bunch. A particularly common one (it might exist in all existing code) is use-after-free when `DestroyWindow` is called inside a window procedure, causing a drop in the reentrant call.

Partly to document this, and partly to make a good basis for experimentation and partly to explore whether there's appetite for cross-ecosystem low level infrastructure, I've created the [win-win](https://crates.io/crates/win-win) crate.

In my experience, UI code tends to play especially fast and loose with soundness, with lots of lifetime and thread safety issues. Also, the lower levels were done in a time predating our modern understanding of safety. I believe it's possible to do better, to build stuff on a sound foundation, then use the safety properties of Rust at the higher levels.
## [11][Export Rust functions to Java using JNI (WebAssembly comparison)](https://www.reddit.com/r/rust/comments/eyppst/export_rust_functions_to_java_using_jni/)
- url: https://www.reddit.com/r/rust/comments/eyppst/export_rust_functions_to_java_using_jni/
---
I wrote a blog post last week on importing Rust WebAssembly functions inside a JavaFX project. I have investigated this week how to use the Java Native Interface directly on a Rust dylib and compared the two approaches in convenience and speed.

[http://vegapit.com/article/export-rust-java-jni-webassembly-comparison](http://vegapit.com/article/export-rust-java-jni-webassembly-comparison)
## [12][Learning Rust: Memory, Ownership and Borrowing](https://www.reddit.com/r/rust/comments/eyepht/learning_rust_memory_ownership_and_borrowing/)
- url: https://www.youtube.com/watch?v=8M0QfLUDaaA
---

