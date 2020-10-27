# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (44/2020)!](https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (44/2020)?](https://www.reddit.com/r/rust/comments/ji8vgk/whats_everyone_working_on_this_week_442020/)
- url: https://www.reddit.com/r/rust/comments/ji8vgk/whats_everyone_working_on_this_week_442020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-44-2020/50643?u=llogiq)!
## [3][Here's a fun little crate I made that generates an interactive ASCII globe in the terminal](https://www.reddit.com/r/rust/comments/jj0pde/heres_a_fun_little_crate_i_made_that_generates_an/)
- url: https://i.redd.it/0qub1980wmv51.gif
---

## [4][Cranelift has just been successfully merged as an optional backend for rustc](https://www.reddit.com/r/rust/comments/jil2rs/cranelift_has_just_been_successfully_merged_as_an/)
- url: https://github.com/rust-lang/rust/pull/77975
---

## [5][Introducing rusteval - add a REPL to your application with a few macros](https://www.reddit.com/r/rust/comments/jj09i8/introducing_rusteval_add_a_repl_to_your/)
- url: https://docs.rs/rusteval/0.1.0/rusteval/
---

## [6][smolscale: A global, auto-scaling, preemptive scheduler based on async-executor](https://www.reddit.com/r/rust/comments/jipcgz/smolscale_a_global_autoscaling_preemptive/)
- url: https://docs.rs/smolscale/0.1.0/smolscale/
---

## [7][Chunked Template engine](https://www.reddit.com/r/rust/comments/jiyx9z/chunked_template_engine/)
- url: https://www.reddit.com/r/rust/comments/jiyx9z/chunked_template_engine/
---
Hello,

I'm new to Rust, and so far I had a lot of fun porting my code from PHP/C to Rust/Rocket/Tera/Diesel, but one feature i have in C particularly is chunked templating. it's particularly nice for hi-load services for the customer (and google bot) to have the first byte of the page asap. I know Rocket itself allow Stream, but i didn't found any template engine that use it for allowing Futures in context for example.

Do you know any ?
## [8][Bytecode Alliance: One year update](https://www.reddit.com/r/rust/comments/jiic8m/bytecode_alliance_one_year_update/)
- url: https://bytecodealliance.org/articles/1-year-update
---

## [9][Text layout is a loose hierarchy of segmentation](https://www.reddit.com/r/rust/comments/jin8wh/text_layout_is_a_loose_hierarchy_of_segmentation/)
- url: https://raphlinus.github.io/text/2020/10/26/text-layout.html
---

## [10][Librsvg's test suite is now in Rust](https://www.reddit.com/r/rust/comments/jim2l1/librsvgs_test_suite_is_now_in_rust/)
- url: https://people.gnome.org/~federico/blog/librsvg-test-suite-is-now-in-rust.html
---

## [11][rust-analyzer changelog #48](https://www.reddit.com/r/rust/comments/jifcky/rustanalyzer_changelog_48/)
- url: https://rust-analyzer.github.io/thisweek/2020/10/26/changelog-48.html
---

## [12][cargo:rerun-if-changed=build.rs](https://www.reddit.com/r/rust/comments/jj0jpf/cargorerunifchangedbuildrs/)
- url: https://www.reddit.com/r/rust/comments/jj0jpf/cargorerunifchangedbuildrs/
---
Is this working for you guys?

`build.rs`:
```rust
fn main() {
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:warning=Only once?");
}
```

Warning appears on every `cargo build`...
Same with build script in `build/mod.rs`.

What am I doing wrong?
