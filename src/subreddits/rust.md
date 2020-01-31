# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (5/2020)!](https://www.reddit.com/r/rust/comments/eupam3/hey_rustaceans_got_an_easy_question_ask_here_52020/)
- url: https://www.reddit.com/r/rust/comments/eupam3/hey_rustaceans_got_an_easy_question_ask_here_52020/
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
## [2][This Week in Rust 323](https://www.reddit.com/r/rust/comments/ew3gri/this_week_in_rust_323/)
- url: https://this-week-in-rust.org/blog/2020/01/28/this-week-in-rust-323/
---

## [3][Current state of asynchronous programming in Rust](https://www.reddit.com/r/rust/comments/ewm7gp/current_state_of_asynchronous_programming_in_rust/)
- url: https://www.reddit.com/r/rust/comments/ewm7gp/current_state_of_asynchronous_programming_in_rust/
---
I've read a few chapters from the incomplete [async-book](https://rust-lang.github.io/async-book/) and found that the async ecosystem is quite confusing at the moment. For a new Rustacean (like me), is now the right time to start learning asynchronous programming in Rust? Or come back later when the ecosystem is more mature?

Another question is where (which library) should we start from, here's my understanding and please correct me if I'm wrong: The std `Future` trait and `async/await` syntax were stabilized in Rust 1.39, but most existing libraries were build on the [rust-lang/futures-rs](https://github.com/rust-lang/futures-rs) crate, then eventually the futures-rs crate will be deprecated and the whole ecosystem will migrate to std?
## [4][Pueue: Manage your shell commands. Now rewritten in Rust.](https://www.reddit.com/r/rust/comments/ewklg5/pueue_manage_your_shell_commands_now_rewritten_in/)
- url: https://github.com/nukesor/pueue
---

## [5][Announcing Rust 1.41.0 | Rust Blog](https://www.reddit.com/r/rust/comments/ew8htz/announcing_rust_1410_rust_blog/)
- url: https://blog.rust-lang.org/2020/01/30/Rust-1.41.0.html
---

## [6][Performance comparison of swc and babel](https://www.reddit.com/r/rust/comments/ewkd01/performance_comparison_of_swc_and_babel/)
- url: https://www.reddit.com/r/rust/comments/ewkd01/performance_comparison_of_swc_and_babel/
---
https://swc-project.github.io/blog/2020/01/31/perf-swc-vs-babel

`swc` is a typescript / javascript compiler, which takes next-generation javascript input and emits javascript codes which work on the old browsers.

swc is written in rust, and babel is written in js.


TL;DR;

Even when swc does more work than babel, 
 - swc is 20x faster than babel on a single-core benchmark.
 - swc is 68x faster than babel on a multicore benchmark. (4 physical core, HT enabled)


[swc repository](https://github.com/swc-project/swc)
## [7][Say hello to OpenSK: a fully open-source security key implementation](https://www.reddit.com/r/rust/comments/ewb4y8/say_hello_to_opensk_a_fully_opensource_security/)
- url: https://security.googleblog.com/2020/01/say-hello-to-opensk-fully-open-source.html
---

## [8][First (0.1.0) release of new provenanced YAML library](https://www.reddit.com/r/rust/comments/ewn80v/first_010_release_of_new_provenanced_yaml_library/)
- url: https://www.reddit.com/r/rust/comments/ewn80v/first_010_release_of_new_provenanced_yaml_library/
---
Today I released 0.1.0 of my [`marked-yaml`](https://crates.io/crates/marked-yaml) library ([Github](https://github.com/kinnison/marked-yaml), [docs.rs](https://docs.rs/marked-yaml)).

In brief, it's meant for parsing YAML into a data structure where all the values are marked with provenance of where they're from in the input string.  This lets you read YAML and then give feedback to users when there are issues of where those issues came from in the input they provided.

It's early days and I might change the API dramatically, but I hope this might be of interest/use to others.  I am using it for providing better error messages than I could before in a project I'm working on with a friend.
## [9][My first month of Rust](https://www.reddit.com/r/rust/comments/ewmmt4/my_first_month_of_rust/)
- url: https://engineering.telia.no/blog/my-first-month-of-rust
---

## [10][Build your own executor](https://www.reddit.com/r/rust/comments/ewodxc/build_your_own_executor/)
- url: https://stjepang.github.io/2020/01/31/build-your-own-executor.html
---

## [11][The 2020 Rust Event Lineup](https://www.reddit.com/r/rust/comments/ewmb99/the_2020_rust_event_lineup/)
- url: https://blog.rust-lang.org/2020/01/31/conf-lineup.html
---

## [12][Bastion 0.3.4: Start of the New Era](https://www.reddit.com/r/rust/comments/ewby1r/bastion_034_start_of_the_new_era/)
- url: https://blog.bastion.rs/2020/01/30/new-age-of-bastion.html
---

