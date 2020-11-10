# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (46/2020)!](https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (46/2020)?](https://www.reddit.com/r/rust/comments/jqrmqn/whats_everyone_working_on_this_week_462020/)
- url: https://www.reddit.com/r/rust/comments/jqrmqn/whats_everyone_working_on_this_week_462020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-46-2020/51226?u=llogiq)!
## [3][aarch64-unknown-linux-gnu is on track to become a Tier 1 target in Rust 1.49!](https://www.reddit.com/r/rust/comments/jr074o/aarch64unknownlinuxgnu_is_on_track_to_become_a/)
- url: https://github.com/rust-lang/rust/pull/78228#event-3971582295
---

## [4][Q: Rust used to have a ton of pointer types; what were they and where did they go?](https://www.reddit.com/r/rust/comments/jrkiuw/q_rust_used_to_have_a_ton_of_pointer_types_what/)
- url: https://www.reddit.com/r/rust/comments/jrkiuw/q_rust_used_to_have_a_ton_of_pointer_types_what/
---
Today's Rust is more streamlined, but you may know that Rust used to have a ton of pointer types:

&gt; Older Rusts had many more pointer types, they’re gone now.
&gt;
&gt; — https://doc.rust-lang.org/guide-pointers.html

I'm looking to understand the chronology of them to understand the language's evolution a little more deeply.

Q1: What were the older pointer types? Ideally I'm looking for actual data from the Rust docs (and not specific memories people have), if possible. Are there previous versions of the Rust docs online that I can look at or do I have to recompile the old docs to do so?

Q2: What became of each pointer type as the language evolved (e.g. became redundant, functionality merged someplace else, dropped as a feature, etc.)?
## [5][My first impression and some questions regarding MVP Clean and Dependency Injection in Rust](https://www.reddit.com/r/rust/comments/jriowl/my_first_impression_and_some_questions_regarding/)
- url: https://www.reddit.com/r/rust/comments/jriowl/my_first_impression_and_some_questions_regarding/
---
I'm currently working on my first GTK application written in Rust and I am trying to figure out how to achieve a [Model-View-Presenter Clean Architecture](https://github.com/ImangazalievM/CleanArchitectureManifest).

I used this pattern for years in Mobile application development and I managed to use it also in my previews GTK apps written in Python.

But I'm frankly struggling to figure out how to implement this pattern in Rust. I'm sure it is mostly due to my lack of knowledge and experience with it (I literally started last week with Rust) but I'm suspecting that some patterns are not really common in this language.

**Is Dependency Injection and Inversion of Control a common pattern in Rust?**

I currently believe that is not, because I was not able to find any useful material about it and there are basically no stable, widely used and currently maintained DI/IOC libraries in [crates.io](https://crates.io/keywords/dependency-injection): most of the libraries are just unmaintained projects made by a single developer that most of the time didn't even reach version 1.0.0 and have, in the best case, a couple of thousands total downloads.

[waiter\_di](https://crates.io/crates/waiter_di) seems a promising new library, but I am a little scared by the fact that almost nobody uses it (only 500 download in total from crates, 4 stars on GitHub and literally [1 single public project](https://github.com/search?l=Rust&amp;q=waiter_di&amp;type=Code) is using it, beside the owner himself) and that it has only a [single contributor](https://github.com/dmitryb-dev/waiter/graphs/contributors).

In Java there is [Dagger](https://dagger.dev/), in Kotlin there are [Koin](https://doc.insert-koin.io/#/) and [Kodein](https://docs.kodein.org/kodein-di/7.1/index.html), in Python there is [Injector](https://injector.readthedocs.io/en/latest/). These are just some of the most popular libraries. Is it really possible that in Rust there is nothing comparable?

I know that you can achieve DI/IOC without the need of any framework but, when applying the MVP Clean pattern, on a medium sized app, you have to create and provide dozens of different UseCases/Interactors. Having to provide these without a proper DI framework is a major annoyance, imho.

Am I missing something? Has anyone here already successfully used MVP Clean with Rust? If yes, is the project public? Can I have a look at it? \^\^
## [6][How to optimize your Rust development workflow on Kubernetes with Okteto](https://www.reddit.com/r/rust/comments/jrjmog/how_to_optimize_your_rust_development_workflow_on/)
- url: https://www.reddit.com/r/rust/comments/jrjmog/how_to_optimize_your_rust_development_workflow_on/
---
You write your code locally and okteto detects the changes and instantly updates your Kubernetes application.

[https://okteto.com/blog/getting-started-with-okteto-and-rust/](https://okteto.com/blog/getting-started-with-okteto-and-rust/)
## [7][Benchmark Godot &amp; Rust](https://www.reddit.com/r/rust/comments/jr5iww/benchmark_godot_rust/)
- url: https://blog.extrawurst.org/general/gamedev/rust/2020/11/07/godot-rust-benchmark.html
---

## [8][A Quick Tour of Tradeoffs Embedding Data in Rust](https://www.reddit.com/r/rust/comments/jrjzfa/a_quick_tour_of_tradeoffs_embedding_data_in_rust/)
- url: https://nickb.dev/blog/a-quick-tour-of-trade-offs-embedding-data-in-rust
---

## [9][Rust tool for highly parallel per-base depth analysis in BAM/CRAM files, with accompanying library for parallel genomic iteration.](https://www.reddit.com/r/rust/comments/jrbfh0/rust_tool_for_highly_parallel_perbase_depth/)
- url: https://github.com/sstadick/perbase
---

## [10][Why Rust is the Future of Game Development](https://www.reddit.com/r/rust/comments/jqux1s/why_rust_is_the_future_of_game_development/)
- url: https://thefuntastic.com/blog/why-rust-is-the-future-game-dev
---

## [11][This Month in Rust GameDev #15 - October 2020](https://www.reddit.com/r/rust/comments/jr2170/this_month_in_rust_gamedev_15_october_2020/)
- url: https://rust-gamedev.github.io/posts/newsletter-015
---

## [12][rust-analyzer changelog #50](https://www.reddit.com/r/rust/comments/jqwfoj/rustanalyzer_changelog_50/)
- url: https://rust-analyzer.github.io/thisweek/2020/11/09/changelog-50.html
---

