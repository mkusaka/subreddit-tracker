# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (31/2020)!](https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 349](https://www.reddit.com/r/rust/comments/i094wo/this_week_in_rust_349/)
- url: https://this-week-in-rust.org/blog/2020/07/28/this-week-in-rust-349/
---

## [3][Go vs Rust: Writing a CLI tool](https://www.reddit.com/r/rust/comments/i3fxwg/go_vs_rust_writing_a_cli_tool/)
- url: https://cuchi.me/posts/go-vs-rust
---

## [4][Why async version of simple program use to 50x more memory](https://www.reddit.com/r/rust/comments/i3fync/why_async_version_of_simple_program_use_to_50x/)
- url: https://www.reddit.com/r/rust/comments/i3fync/why_async_version_of_simple_program_use_to_50x/
---
Asked on stackoverflow:

https://stackoverflow.com/questions/63243354/why-async-version-of-simple-program-use-to-50x-more-memory

The program is an echo server written using standard library, tokio and async_std. So the question is that why does the async version using tokio/async_std consume 55x more memory than synchronous one?

    Memory usage:
    std: 11 MB
    tokio: 607 MB
    async_std: 607 MB
## [5][Announcing Rust 1.45.2](https://www.reddit.com/r/rust/comments/i2zp72/announcing_rust_1452/)
- url: https://blog.rust-lang.org/2020/08/03/Rust-1.45.2.html
---

## [6][A Webserver in Rust](https://www.reddit.com/r/rust/comments/i3gka2/a_webserver_in_rust/)
- url: https://www.reddit.com/r/rust/comments/i3gka2/a_webserver_in_rust/
---
Hi!

I figured that a Rust Web **server** (not a Framework) would be nice.

[https://github.com/User65k/flash\_rust\_ws/](https://github.com/User65k/flash_rust_ws/)

Some Features are still missing but what do you think?  
Feedback, Issues and Pull-requests are very welcome :-)
## [7][rust-analyzer changelog #36](https://www.reddit.com/r/rust/comments/i2wic3/rustanalyzer_changelog_36/)
- url: https://rust-analyzer.github.io/thisweek/2020/08/03/changelog-36.html
---

## [8][A gentle intro to assembly with Rust](https://www.reddit.com/r/rust/comments/i32wf3/a_gentle_intro_to_assembly_with_rust/)
- url: https://lfn3.net/2020/08/03/a-gentle-intro-to-assembly-with-rust/
---

## [9][Help me please for a function generic to both u32 and u64](https://www.reddit.com/r/rust/comments/i3hjf2/help_me_please_for_a_function_generic_to_both_u32/)
- url: https://www.reddit.com/r/rust/comments/i3hjf2/help_me_please_for_a_function_generic_to_both_u32/
---
Hi everyone,  


I have a quick question about generic functions. Here is what I want: to simply define a function taking as input slices of either u32 or u64 without wrapping into a struct or anything that could slow down the computation. I would love to have something like that (since u32 and u64 implement the same operators that I will have to use):

```rust
pub fn print_t&lt;T: u32 OR u64&gt;(t: &amp;[T]) {
    for el in t.iter() {
        println!("{}", el);
        // el.wrapping_add(4), el.wrapping_mul ...
    }
}

fn main() {
    let mut t32: Vec&lt;u32&gt; = vec![0; 10];
    let mut t64: Vec&lt;u64&gt; = vec![0; 10];

    print_t&lt;u32&gt;(t32);
    print_t&lt;u64&gt;(t64);
}
```
Is there a simple way to do that?
## [10][Building a Brainf*ck Compiler with Rust and LLVM](https://www.reddit.com/r/rust/comments/i3iy5e/building_a_brainfck_compiler_with_rust_and_llvm/)
- url: https://benkonz.github.io/building-a-brainfuck-compiler-with-rust-and-llvm/
---

## [11][Rust for JavaScript Developers](https://www.reddit.com/r/rust/comments/i3i07z/rust_for_javascript_developers/)
- url: https://vilagedeveloper.hashnode.dev/rust-for-javascript-developers-ckagxyab500tzlns1wd78lexq
---

## [12][rust-analyser: how do I force it to use stable Rust in vscode?](https://www.reddit.com/r/rust/comments/i3gd6h/rustanalyser_how_do_i_force_it_to_use_stable_rust/)
- url: https://www.reddit.com/r/rust/comments/i3gd6h/rustanalyser_how_do_i_force_it_to_use_stable_rust/
---
Stability seems not good with \`default nightly\` Rust. I need to stop and restart vscode constantly to remove visible errors after minor edits. Removing and re-installing the rust-analyzer plugin does not help.
