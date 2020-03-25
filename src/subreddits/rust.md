# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (13/2020)!](https://www.reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (13/2020)?](https://www.reddit.com/r/rust/comments/fnflsh/whats_everyone_working_on_this_week_132020/)
- url: https://www.reddit.com/r/rust/comments/fnflsh/whats_everyone_working_on_this_week_132020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-13-2020/39864?u=llogiq)!
## [3][Learning Rust feels overwhelming](https://www.reddit.com/r/rust/comments/fomxnc/learning_rust_feels_overwhelming/)
- url: https://www.reddit.com/r/rust/comments/fomxnc/learning_rust_feels_overwhelming/
---
Maybe it is because I have worked with JS ( Aah ) mostly in my short coding life. I'm currently in the Ownership section of the Rust book and it totally smashed my head. It's like I need to forget everything I learnt in C classes to understand Rust. I'm up for the challenge though as I will be home for the next 21 days due to Corona Lockdown nationwide here.   


Also, I have huge respect for those programmers who work with Rust daily. You guys really tamed the wild horse.
## [4][Learn Rust building a toy JIRA clone](https://www.reddit.com/r/rust/comments/fog2dj/learn_rust_building_a_toy_jira_clone/)
- url: https://github.com/LukeMathWalker/build-your-own-jira-with-rust
---

## [5][Rust Pie!](https://www.reddit.com/r/rust/comments/fopj3r/rust_pie/)
- url: https://i.redd.it/1ig1geehcto41.jpg
---

## [6][Reducing memory consumption in librsvg, part 3: slack space in Bézier paths](https://www.reddit.com/r/rust/comments/fopa94/reducing_memory_consumption_in_librsvg_part_3/)
- url: https://people.gnome.org/~federico/blog/reducing-memory-consumption-in-librsvg-3.html
---

## [7][Plans for the regex crate](https://www.reddit.com/r/rust/comments/fo8clp/plans_for_the_regex_crate/)
- url: https://github.com/rust-lang/regex/issues/656
---

## [8][My rust-themed home office decoration](https://www.reddit.com/r/rust/comments/fopy3b/my_rustthemed_home_office_decoration/)
- url: https://i.redd.it/gps2brjqhto41.jpg
---

## [9][`zig cc`: Drop-In Replacement for GCC/Clang with built in cross-compilation](https://www.reddit.com/r/rust/comments/fodv5d/zig_cc_dropin_replacement_for_gccclang_with_built/)
- url: https://andrewkelley.me/post/zig-cc-powerful-drop-in-replacement-gcc-clang.html
---

## [10][tokio vs async-std in 2020?](https://www.reddit.com/r/rust/comments/fohfb7/tokio_vs_asyncstd_in_2020/)
- url: https://www.reddit.com/r/rust/comments/fohfb7/tokio_vs_asyncstd_in_2020/
---
I'm aware of https://www.reddit.com/r/rust/comments/dngig6/tokio_vs_asyncstd but since so much has changed since that post I think it would be good to get an update. Thanks!
## [11][Announcing Agnostik: Executor Agnostic Runtime](https://www.reddit.com/r/rust/comments/fo98kc/announcing_agnostik_executor_agnostic_runtime/)
- url: https://www.reddit.com/r/rust/comments/fo98kc/announcing_agnostik_executor_agnostic_runtime/
---
[GitHub](https://github.com/bastion-rs/agnostik) | [Docs](https://docs.rs/agnostik) | [Crates.io](https://crates.io/crates/agnostik)

We are happy to make everything agnostic, so we made Agnostik. Agnostik is a layer between your application and the executor that is used to execute futures. It allows you to switch between the executors smoothly and without having to change much code. You can use Agnostik in every library that requires an executor but wants to let the user decide which executor should be used. You can also use Agnostik in an application if you plan to use multiple executors or want to switch between executors. Soon, the Executor trait will be exposed with a unified async block handle.

&amp;#x200B;

Brought to you by [Bastion](https://bastion.rs)
## [12][Inviting God's Wrath with Cursed Rust](https://www.reddit.com/r/rust/comments/fo487h/inviting_gods_wrath_with_cursed_rust/)
- url: http://troubles.md/abusing-rustc/
---

