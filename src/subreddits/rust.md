# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (8/2020)!](https://www.reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here_82020/)
- url: https://www.reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here_82020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (8/2020)?](https://www.reddit.com/r/rust/comments/f541u0/whats_everyone_working_on_this_week_82020/)
- url: https://www.reddit.com/r/rust/comments/f541u0/whats_everyone_working_on_this_week_82020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-8-2020/38329?u=llogiq)!
## [3][Learning OpenGL with Rust](https://www.reddit.com/r/rust/comments/f693wh/learning_opengl_with_rust/)
- url: https://www.reddit.com/r/rust/comments/f693wh/learning_opengl_with_rust/
---
Hi, i'm new to Rust but not new to programming. I have recently wanted to start learning OpenGL but at the same time I want to learn Rust. Should i stick it out and do it this way or would it be better for me to learn OpenGL with another language that i'm a bit better with?

The best outcome for me here is that I use Rust.

If anyone has any tutorials or advice, it would be much appreciated.

I program on a Mac btw...
## [4][Rust ghost, signing off](https://www.reddit.com/r/rust/comments/f5ufc1/rust_ghost_signing_off/)
- url: https://quietmisdreavus.net/self/2020/02/17/rust-ghost-signing-off/
---

## [5][PSA: `deny(warnings)` is actively harmful](https://www.reddit.com/r/rust/comments/f5xpib/psa_denywarnings_is_actively_harmful/)
- url: https://www.reddit.com/r/rust/comments/f5xpib/psa_denywarnings_is_actively_harmful/
---
Some crates like to use `#![deny(warnings)]` or `#![cfg_attr(test, deny(warnings)]` as a way of enforcing clean code. This is a lot like GCC's `-Werror` option, and a quick [search](https://duckduckgo.com/?q=don't+use+"-Werror") can show a lot of reasons why that is harmful. It may be OK while you're actively maintaining your code, but once you're not paying attention, small signs of bit-rot like using deprecated APIs will turn into outright errors.

When used as a crate dependency, Cargo will mask this with `--cap-lints allow`/`warn`, so it's usually not the end of the world. However, when the Rust teams want to test the impact of a potentially breaking change, they run mass builds with Crater. If your crate is already failing due to `deny(warnings)`, then any future changes that really do break you won't be seen as a regression.

Case in point, [rust-mac#15](https://github.com/reem/rust-mac/issues/15), their `if_cfg!` macro is broken by the [changes to the `meta` matcher](https://github.com/rust-lang/rust/pull/63674) in Rust 1.40. However, the crater run for that change showed 0 regressions, because `mac`'s tests were already broken by the simple deprecation of `...` in pattern matches.

I don't want to make a mountain out of a molehill, and I certainly don't mean to blame/shame anyone involved in that example. I just want to show this as a concrete way that `deny(warnings)` is actively harmful to the maintenance of the Rust ecosystem. For that matter, any wide net like `deny(deprecated)` would be just as bad. Please don't open yourself to build failures from stuff that should just be a warning!
## [6][Introducing kibi, a text editor in ≤1024 lines of Rust with syntax highlighting, incremental search, UTF-8 support, and more. Contributions welcome!](https://www.reddit.com/r/rust/comments/f5vpco/introducing_kibi_a_text_editor_in_1024_lines_of/)
- url: https://github.com/ilai-deutel/kibi
---

## [7][Announcing Nushell 0.10.0 - now with initial scripting support](https://www.reddit.com/r/rust/comments/f5wdh5/announcing_nushell_0100_now_with_initial/)
- url: https://www.nushell.sh/blog/2020/02/18/nushell-0_10_0.html
---

## [8][Materialize: the streaming data warehouse, written in Rust](https://www.reddit.com/r/rust/comments/f5y0bz/materialize_the_streaming_data_warehouse_written/)
- url: https://materialize.io/blog-introduction/
---

## [9][R we web yet?](https://www.reddit.com/r/rust/comments/f67e99/r_we_web_yet/)
- url: https://www.reddit.com/r/rust/comments/f67e99/r_we_web_yet/
---
R we web yet?
Is Rust good for backend?
Like APIs, etc.
Can someone recommend me something?
## [10][tantivy v0.12 released](https://www.reddit.com/r/rust/comments/f6aig1/tantivy_v012_released/)
- url: https://www.reddit.com/r/rust/comments/f6aig1/tantivy_v012_released/
---
For those who are not familiar with tantivy, it is a search engine library  for rust inspired by Lucene.

* [github]([https://github.com/tantivy-search/tantivy/](https://github.com/tantivy-search/tantivy/))
* [crates.io]([https://crates.io/crates/tantivy](https://crates.io/crates/tantivy))

&amp;#x200B;

**Changelog**

* Removing static dispatch in tokenizers for simplicity. (#762)
* Added backward iteration for `TermDictionary` stream. (@halvorboe)
* Fixed a performance issue when searching for the posting lists of a missing term (@audunhalland)
* Added a configurable maximum number of docs (10M by default) for a segment to be considered for merge (@hntd187, landed by @halvorboe #713)
* Important Bugfix #777, causing tantivy to retain memory mapping. (diagnosed by @poljar)
* Added support for field boosting. (#547, u/fulmicoton)


Thanks to all contributors and all of patreon/github sponsors! (hntd187, mosuka, Colin, Florian, Fredrik, Nate, Sanghyeon, Zak and Zane)
## [11][IntelliJ Rust Changelog #116](https://www.reddit.com/r/rust/comments/f5sgs0/intellij_rust_changelog_116/)
- url: https://intellij-rust.github.io/2020/02/18/changelog-116.html
---

## [12][Crosspost: r/learnrust - Rust tutorials and recipes on YouTube](https://www.reddit.com/r/rust/comments/f64dn4/crosspost_rlearnrust_rust_tutorials_and_recipes/)
- url: /r/learnrust/comments/f639yg/a_list_of_some_youtube_videos_free_that_you_can/
---

