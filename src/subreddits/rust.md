# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (24/2020)!](https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (24/2020)?](https://www.reddit.com/r/rust/comments/gysxgq/whats_everyone_working_on_this_week_242020/)
- url: https://www.reddit.com/r/rust/comments/gysxgq/whats_everyone_working_on_this_week_242020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-24-2020/43988?u=llogiq)!
## [3][New inline assembly syntax available in nightly | Inside Rust Blog](https://www.reddit.com/r/rust/comments/gzil0l/new_inline_assembly_syntax_available_in_nightly/)
- url: https://blog.rust-lang.org/inside-rust/2020/06/08/new-inline-asm.html
---

## [4][Futures and Segmented Stacks](https://www.reddit.com/r/rust/comments/gz5oe0/futures_and_segmented_stacks/)
- url: https://without.boats/blog/futures-and-segmented-stacks/
---

## [5][PNGme: An Intermediate Rust Project](https://www.reddit.com/r/rust/comments/gzehpz/pngme_an_intermediate_rust_project/)
- url: https://picklenerd.github.io/pngme_book/
---

## [6][Why does rust need something like placement new](https://www.reddit.com/r/rust/comments/gzmbpa/why_does_rust_need_something_like_placement_new/)
- url: https://www.reddit.com/r/rust/comments/gzmbpa/why_does_rust_need_something_like_placement_new/
---
Why was there an attempt to introduce new syntax for placement new? Why can't the rust compiler optimize the cases by itself (ever?)?

For e.g.: In `vec.push(MyStruct { foo, bar})`, why can't this be compiled such that MyStruct is constructed in place by itself without any special syntax? Is there a requirement that MyStruct be constructed on stack before moving to heap?

A lot of code is already written like this, which would benefit from implicit optimization rather than creating a new syntax. Is there some technical difficulty in achieving this?

I remember seeing an issue related to `Box::new` not skipping copy even when it's `#[inline(always)]`.

Forgive me if it's a stupid question.
## [7][rust-analyzer changelog #28](https://www.reddit.com/r/rust/comments/gyzk6s/rustanalyzer_changelog_28/)
- url: https://rust-analyzer.github.io/thisweek/2020/06/08/changelog-28.html
---

## [8][Not everything is UTF-8](https://www.reddit.com/r/rust/comments/gz33u6/not_everything_is_utf8/)
- url: https://octobus.net/blog/2020-06-05-not-everything-is-utf8.html
---

## [9][What's the current state of Rust's Sparse Linear Algebra libraries?](https://www.reddit.com/r/rust/comments/gzazna/whats_the_current_state_of_rusts_sparse_linear/)
- url: https://www.reddit.com/r/rust/comments/gzazna/whats_the_current_state_of_rusts_sparse_linear/
---
Hello all,

I've been doing a lot of simulations recently, and I need to solve some ~~large~~ (2000 x 2000) symmetric and non-triangular sparse (~40000 non-zero entries) linear systems of equations where the matrices are changing in a tight loop. (The matrices are Jacobians of a highly non-linear function). The RHS vectors are dense. I currently store the RHS in a `ndarray::Array1`.

What's the best library to use for this?

I've been using `sprs`, but its performance was terrible in my benchmarking.

I understand `nalgebra` has sparse matrix support, but I was under the impression that `nalgebra` isn't meant for high dimensional systems such as mine.

EDIT: As many have pointed out these matrices aren't really large. I'll have larger ones in future iterations of this project, but should have been more careful with my phrasing.
## [10][Prose: a Markdown to HTML parser built using Nom and Yew](https://www.reddit.com/r/rust/comments/gzmeoa/prose_a_markdown_to_html_parser_built_using_nom/)
- url: https://github.com/HGHimself/prose
---

## [11][This Month in Rust GameDev #10 - May 2020](https://www.reddit.com/r/rust/comments/gz59mk/this_month_in_rust_gamedev_10_may_2020/)
- url: https://rust-gamedev.github.io/posts/newsletter-010
---

## [12][unFTP: When you need to FTP, but don't want to.](https://www.reddit.com/r/rust/comments/gz12l4/unftp_when_you_need_to_ftp_but_dont_want_to/)
- url: https://github.com/bolcom/unFTP
---

