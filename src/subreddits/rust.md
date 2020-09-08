# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (37/2020)!](https://www.reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (37/2020)?](https://www.reddit.com/r/rust/comments/ioc6ld/whats_everyone_working_on_this_week_372020/)
- url: https://www.reddit.com/r/rust/comments/ioc6ld/whats_everyone_working_on_this_week_372020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-37-2020/48445?u=llogiq)!
## [3][How to speed up the Rust compiler one last time](https://www.reddit.com/r/rust/comments/ioh9yi/how_to_speed_up_the_rust_compiler_one_last_time/)
- url: https://blog.mozilla.org/nnethercote/2020/09/08/how-to-speed-up-the-rust-compiler-one-last-time/
---

## [4][On finally learning to program at the age of 40](https://www.reddit.com/r/rust/comments/iopo8m/on_finally_learning_to_program_at_the_age_of_40/)
- url: https://github.com/Dhghomon/programming_at_40/blob/master/README.md
---

## [5][Juniper - fast &amp; type-safe GraphQL servers in Rust](https://www.reddit.com/r/rust/comments/ioskx6/juniper_fast_typesafe_graphql_servers_in_rust/)
- url: https://blog.graphqleditor.com/graphql-rust-juniper/
---

## [6][I Made A Rust Quine](https://www.reddit.com/r/rust/comments/iocgex/i_made_a_rust_quine/)
- url: https://www.reddit.com/r/rust/comments/iocgex/i_made_a_rust_quine/
---
A quine is a simple script that prints it's source. After playing around with the RustBot in the Rust Programming Discord, I realized it actually marks all the output code blocks with the \`rust\` language tag, giving it Rust markdown. I then knew what I had to do next.

    fn main() {
        macro_rules! script {() =&gt; {"fn main() {{\n\tmacro_rules! script {{() =&gt; {{{:?}}}}}\n\tprintln!(script!(), script!());\n}}"}}
        println!(script!(), script!());
    }

It's as simple as it gets. Macro is used so we can reference the string twice (println! requires a string literal as the first argument), and then we format it with itself.

Try it out for yourself [here](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=6145572e99a52183c707980d3618d648)!
## [7][swc v1.2.23 is out! Much better performance, jsx pragma, etc...](https://www.reddit.com/r/rust/comments/ioka4e/swc_v1223_is_out_much_better_performance_jsx/)
- url: https://swc-project.github.io/blog/2020/09/08/swc-1.2.23
---

## [8][Macro to chain block effects such as unsafe, async, etc.](https://www.reddit.com/r/rust/comments/ioqx0u/macro_to_chain_block_effects_such_as_unsafe_async/)
- url: https://crates.io/crates/block_effects
---

## [9][Is there some guide for Linux From Scratch but with with only Rust code mostly?](https://www.reddit.com/r/rust/comments/ioqr7t/is_there_some_guide_for_linux_from_scratch_but/)
- url: https://www.reddit.com/r/rust/comments/ioqr7t/is_there_some_guide_for_linux_from_scratch_but/
---
There is [https://github.com/mesalock-linux/mesalock-distro](https://github.com/mesalock-linux/mesalock-distro) but last commit is from 2018-11-12.

I would like to try such distro using only what is needed to run all applications written in Rust.

Some low-level stuff would need to be in C/C++ but we already have nushell, smithay etc.  


My goal would be to have cargo manage dependencies and some minimal/existing? pacakge manager for C/C++ deps.
## [10][Is there any way to represent some_slice[1..33] as a [T; 32] instead of just a [T]?](https://www.reddit.com/r/rust/comments/ioogfn/is_there_any_way_to_represent_some_slice133_as_a/)
- url: https://www.reddit.com/r/rust/comments/ioogfn/is_there_any_way_to_represent_some_slice133_as_a/
---

## [11][generic-simd: easy, safe, cross-platform SIMD](https://www.reddit.com/r/rust/comments/ioe0u6/genericsimd_easy_safe_crossplatform_simd/)
- url: https://crates.io/crates/generic-simd
---

## [12][This Month in Rust GameDev #13 - August 2020](https://www.reddit.com/r/rust/comments/ioftll/this_month_in_rust_gamedev_13_august_2020/)
- url: https://rust-gamedev.github.io/posts/newsletter-013
---

