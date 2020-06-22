# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (26/2020)!](https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (26/2020)?](https://www.reddit.com/r/rust/comments/hdkuz9/whats_everyone_working_on_this_week_262020/)
- url: https://www.reddit.com/r/rust/comments/hdkuz9/whats_everyone_working_on_this_week_262020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-26-2020/44701?u=llogiq)!
## [3][Announcing Tokei 12: A fast &amp; accurate code analysis tool and library.](https://www.reddit.com/r/rust/comments/hdr6gd/announcing_tokei_12_a_fast_accurate_code_analysis/)
- url: https://github.com/XAMPPRocky/tokei/releases/tag/v12.0.0
---

## [4][My project TermUI on it's way to become an actual terminal emulator in a terminal !](https://www.reddit.com/r/rust/comments/hde9dk/my_project_termui_on_its_way_to_become_an_actual/)
- url: https://v.redd.it/zj5nj132tb651
---

## [5][rust-analyzer changelog #30](https://www.reddit.com/r/rust/comments/hds2vd/rustanalyzer_changelog_30/)
- url: https://rust-analyzer.github.io/thisweek/2020/06/22/changelog-30.html
---

## [6][Implementing a Type-safe printf in Rust](https://www.reddit.com/r/rust/comments/hdlzjq/implementing_a_typesafe_printf_in_rust/)
- url: http://willcrichton.net/notes/type-safe-printf/
---

## [7][Tips for Faster Rust Compile Times](https://www.reddit.com/r/rust/comments/hdb5m4/tips_for_faster_rust_compile_times/)
- url: https://endler.dev/2020/rust-compile-times/
---

## [8][PyO3 now supports stable Rust!](https://www.reddit.com/r/rust/comments/hd5506/pyo3_now_supports_stable_rust/)
- url: https://github.com/PyO3/pyo3/issues/5#issuecomment-647094806
---

## [9][Rust Lints You May Not Know](https://www.reddit.com/r/rust/comments/hdgodw/rust_lints_you_may_not_know/)
- url: https://www.possiblerust.com/pattern/rust-lints-you-may-not-know
---

## [10][nom derive 0.6: deriving binary parsers from structure declaration](https://www.reddit.com/r/rust/comments/hdb5h1/nom_derive_06_deriving_binary_parsers_from/)
- url: https://www.reddit.com/r/rust/comments/hdb5h1/nom_derive_06_deriving_binary_parsers_from/
---
[nom-derive](https://crates.io/crates/nom-derive) 0.6.0 has been released, with many improvements over previous release.

This crate adds a custom derive `#[derive(Nom)]`, which automatically creates a parser from a structure or enum declaration (aka declarative parsing). It is mostly intended for binary formats, and relies on nom, creating fast parsers using zero-copy most of the time.

Trivial example:

    #[derive(Nom)]
    struct S {
      a: u32,
      b: u16,
      c: u16
    }

It also supports custom attributes to control parsing endianness, add verifications or custom code, etc. The objective is to make the 99% of the parser automatically derived, while also being able to switch to custom code and parsing functions for some parts.

* [Documentation of the Nom attribute and options](https://docs.rs/nom-derive/0.6.0/nom_derive/derive.Nom.html)
* [Example of parser, taken from an OSPFv3 protocol](https://github.com/rusticata/ospf-parser/blob/master/src/ospfv3.rs)

Feedback is welcome!
## [11][Rooster - Personal Web Server with Rust](https://www.reddit.com/r/rust/comments/hdlbeu/rooster_personal_web_server_with_rust/)
- url: https://github.com/elasmojs/rooster
---

## [12][Why is assignment an expression in Rust?](https://www.reddit.com/r/rust/comments/hddg7u/why_is_assignment_an_expression_in_rust/)
- url: https://www.reddit.com/r/rust/comments/hddg7u/why_is_assignment_an_expression_in_rust/
---
Since I assume asking about language design itself isn't an easy topic, I ask it here.

Why is assignment in rust an expression and not simply a statement? After all it always returns the unit type `()`, thus it's not really useful for chaining assignments like `a = b = c = 3;`, which if I understood correctly is one of the reasons why languages like C made assignment an expression instead of a statement.

PS: I'm writing a small programming language interpreter/compiler and I am modelling it heavily after rust, swift, etc, thus I'm trying to incorporate their language design (although very simplified). That design decision appears to be also be in swift (where chaining expressions) aren't possible, but assignment is an expression evaluating to the unit type.

PPS: If you have any useful blogs etc on such topics, feel free to post it :)
