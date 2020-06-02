# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (23/2020)!](https://www.reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gr3r32/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (23/2020)?](https://www.reddit.com/r/rust/comments/guo51x/whats_everyone_working_on_this_week_232020/)
- url: https://www.reddit.com/r/rust/comments/guo51x/whats_everyone_working_on_this_week_232020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-23-2020/43609?u=llogiq)!
## [3][Rust enters the TIOBE index top 20 for the first time](https://www.reddit.com/r/rust/comments/guzib4/rust_enters_the_tiobe_index_top_20_for_the_first/)
- url: https://www.tiobe.com/tiobe-index/
---

## [4][Rust is amazing! (a C++ developer's experience writing a raytracer in Rust)](https://www.reddit.com/r/rust/comments/gv2act/rust_is_amazing_a_c_developers_experience_writing/)
- url: https://www.reddit.com/r/rust/comments/gv2act/rust_is_amazing_a_c_developers_experience_writing/
---
Background: I've been writing C++ professionally since 2011, and I started to learn Rust last fall by reading a couple of books and solving some Advent of Code 2019 problems (finished 1-11 except 10 in Rust).  I also did the [Ray Tracing in One Weekend](https://raytracing.github.io/) tutorial in both Rust and Go, and just finished the 2nd "Ray Tracking: The Next Week" part of the series in Rust.

Overall, Rust has been an awesome language to learn, and I think it solves many of the common issues/bugs in C++ (which is kind of the whole point).  The concept of lifetimes and the borrow checker really enhances memory safety… our C++ code frequently takes pointers in class constructors and stores them as member variables, and proper lifetime management is totally up to the user (with some `// The object pointed to by this argument must outlive this class.` comments to help).  Many of the bugs we investigate at work are off-by-one or out of bounds accesses as well, so also eliminating those is a huge plus.  I find the if expressions, `match`, implicit returns, and the `?` operator all quite useful, and they help make your code more concise and readable.  I easily managed to get code for the first part (In One Weekend) of the series to run in WebAssembly, although Rayon doesn't work there and it's relatively slow.  It's cool to be able to visualize the ray tracing as it happens, though (I render 1 full frame, display it, render another frame, average it with all the previous ones, display it, and so on, so the image progressively becomes less and less noisy).

I've also found the ecosystem to be quite pleasant to work with, with trivial parallelism available with Rayon (it was super easy to parallelize my raytracer with a simple `into_par_iter`), useful crates for reading .jpg and writing .png files, and good benchmarking/profiling support (`cargo bench`, easy flamegraphs, and the ability to use the standard `perf` tool on binaries).  Performance is also really good… much, much faster than my Go version, and the code is also more readable, thanks to operator overloading.  The compiler and clippy both give really useful hints/tips, and usually make it easy to understand when you're doing something wrong.  I also found code completion in Vim and Visual Studio Code to be dead simple to set up too.

The few minor complaints I have are the lack of inheritance and function overloading, and no implicit type conversions from int to float (you frequently have to do this when using for loop indices in some computation, so I have `i as f32` all over the place).  It's also annoying to have to type 1.0 or 1. instead of simply 1, especially when you're trying to copy and paste some C++ example scene setup.  Compilation is also a bit slow at times (I think using the "image" crate is overkill and pulls in a bunch of dependencies, so I should probably switch to "jpeg_decoder").

[Here's a picture of the final scene from "The Next Week"](https://i.imgur.com/P7e61lb.jpg), which took almost 343 hours of CPU time to render (5120x2880 with 10,000 samples per pixel).  It took 3 hours 37 minutes of wall clock time on a 96 core Google Cloud instance.

Conclusion: Rust is awesome, and I hope to do some more projects in it in the future!                                                                                                                             
                                                                                                                                                                
(I'd post the code, but my employer requires us to go through an open source release approval process to release even personal projects... might do it someday)
## [5][piet-gpu progress report](https://www.reddit.com/r/rust/comments/gv1b95/pietgpu_progress_report/)
- url: https://raphlinus.github.io/rust/graphics/gpu/2020/06/01/piet-gpu-progress.html
---

## [6][Rust Community Wiki (very incomplete)](https://www.reddit.com/r/rust/comments/gv6ktp/rust_community_wiki_very_incomplete/)
- url: https://runrust.miraheze.org
---

## [7][We are the devs behind Lemmy, an open source, Federated alternative to reddit! AMA!](https://www.reddit.com/r/rust/comments/gukx09/we_are_the_devs_behind_lemmy_an_open_source/)
- url: /r/linux/comments/guklhr/we_are_the_devs_behind_lemmy_an_open_source/
---

## [8][Day 1 of learning rust from a programming noob](https://www.reddit.com/r/rust/comments/guspl6/day_1_of_learning_rust_from_a_programming_noob/)
- url: https://www.reddit.com/r/rust/comments/guspl6/day_1_of_learning_rust_from_a_programming_noob/
---
TLDR: I'm trying to learn rust with very little programming experience, and having a blast so far. 

Little background, I'm a controls engineer with a hobby interest in traditional programming. I'm well versed in ladder logic, and know enough Python to write some quick and dirty scripts to automate some of the stuff I do at work.

And that's it.

I kept reading how great rust was, but that it wasn't beginner friendly. I was getting annoyed with finding Python tutorials that interested me, and were at the slightly above basic level I'm at. 

I opened the Rust Book online, and got everything set up per the instructions. 

From my perspective, so far I really enjoy rust and can see what the hype is about. Cargo is fantastic, and the error messages in rust are better than anything else I've ever saw. I feel like I'm learning better habits with rust because I don't really have a choice, whereas with Python just about anything will run. The syntax feels intuitive, and the community is awesome. 

My goal is to be able to contribute either to open source projects or learning resources one day. I've got so much use from both, I want to give back a little.

No, I probably wouldn't recommend others to do what I'm doing, but it's working for me and I'm excited to keep at it.
## [9][A blog from a colleague about his contributions for risc v support in rust, thought it would be worth sharing :)](https://www.reddit.com/r/rust/comments/gv6wl7/a_blog_from_a_colleague_about_his_contributions/)
- url: https://www.codethink.co.uk/articles/2020/improving-risc-v-linux-support-in-rust/
---

## [10][rust-analyzer changelog #27](https://www.reddit.com/r/rust/comments/gumhni/rustanalyzer_changelog_27/)
- url: https://rust-analyzer.github.io/thisweek/2020/06/01/changelog-27.html
---

## [11][IntelliJ Rust Changelog #123](https://www.reddit.com/r/rust/comments/gumy60/intellij_rust_changelog_123/)
- url: https://intellij-rust.github.io/2020/06/01/changelog-123.html
---

## [12][Rust vs FP Scala](https://www.reddit.com/r/rust/comments/gv141b/rust_vs_fp_scala/)
- url: https://www.reddit.com/r/rust/comments/gv141b/rust_vs_fp_scala/
---
Hi all   

I am developing in Scala in FP style and maybe consider to change to Rust. 
What is the pros and cons for Rust over Scala?   

Thanks
