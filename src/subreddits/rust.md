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
## [2][This Week in Rust 341](https://www.reddit.com/r/rust/comments/gvwvep/this_week_in_rust_341/)
- url: https://this-week-in-rust.org/blog/2020/06/02/this-week-in-rust-341/
---

## [3][Getting started with Rust/WinRT](https://www.reddit.com/r/rust/comments/gy372x/getting_started_with_rustwinrt/)
- url: https://kennykerr.ca/2020/06/05/getting-started-with-rust-winrt/
---

## [4][Animal Fight Chess Game written in rust.](https://www.reddit.com/r/rust/comments/gy7ct1/animal_fight_chess_game_written_in_rust/)
- url: https://www.reddit.com/r/rust/comments/gy7ct1/animal_fight_chess_game_written_in_rust/
---
Here is my first  proj: [https://github.com/netcan/AnimalChess](https://github.com/netcan/AnimalChess)
## [5][First project, written in rust](https://www.reddit.com/r/rust/comments/gxqozw/first_project_written_in_rust/)
- url: https://i.redd.it/393ziqodfa351.gif
---

## [6][Why Rust should be preferred for web assembly?](https://www.reddit.com/r/rust/comments/gyabs4/why_rust_should_be_preferred_for_web_assembly/)
- url: https://www.reddit.com/r/rust/comments/gyabs4/why_rust_should_be_preferred_for_web_assembly/
---

## [7][I finally published cargo-fuzzcheck, a structure-aware fuzzing engine for Rust functions. How to move forward?](https://www.reddit.com/r/rust/comments/gxu0q7/i_finally_published_cargofuzzcheck_a/)
- url: https://www.reddit.com/r/rust/comments/gxu0q7/i_finally_published_cargofuzzcheck_a/
---
Hi everyone,

9 months ago I [posted](https://old.reddit.com/r/rust/comments/d02gxk/fuzzcheck_an_evolutionary_coverageguided_fuzzing/) here about my first Rust project, [fuzzcheck](https://github.com/loiclec/fuzzcheck-rs). For those who don’t know it, it is similar to cargo-fuzz, except that the fuzzing engine itself is fully written in Rust and that it is truly structure-aware. That means that it can fuzz-test values of any type T directly, without having to copy, mutate, and then decode a byte array using the “Arbitrary” trait. I can explain why that is preferable in the comments if someone is interested in the details.

Anyway, since then, fuzzcheck has considerably improved. It is much faster, easier to understand, and faster to compile. So  I [published](https://crates.io/crates/fuzzcheck) it on crates.io! I started this project more than 2 years ago, [initially writing it in Swift](https://github.com/loiclec/FuzzCheck). Looking back, I am really proud that I was able keep slowly working on it, and proud of how much it has improved. But I know it’s nowhere near as good as it should be.

Now that it is the end of the academic year and that my summer course has been canceled due to COVID-19, I will have about two/three months where I can spend some time every week working on it. It is a bit of a critical period, because after more than two years, I am running out of motivation to work alone on this, so I would really love this summer to be the time that fuzzcheck is finally used by someone else! Which is why I’d like to ask you for a little bit of help in making fuzzcheck a reality.

* Is the README.md of the project clear, does it convey the purpose of the tool clearly? If not, what is confusing? What kind of explanation/example would you like to see?
* Do you have a (small) project that could benefit from fuzzing? If so, would you like to work together to test fuzzcheck on it? If not, why wouldn’t fuzzing be applicable to your project?
* Do you think I could realistically be sponsored to write fuzzcheck? If so, what steps should I take to receive funding?
* Do you have some free time to contribute to fuzzcheck directly? If so, please contact me! I can explain the design in more details, and then you can choose to work on one of the many tasks that could make fuzzcheck more viable.

Or simply try to run it on your computer, and report to me if something went wrong :)

Thank you so much for your time!
## [8][capnproto-rust now supports [no_std]](https://www.reddit.com/r/rust/comments/gxsfb9/capnprotorust_now_supports_no_std/)
- url: https://dwrensha.github.io/capnproto-rust/2020/06/06/no-std-support.html
---

## [9][Zero To Production #1: Setup - Toolchain, IDEs, CI](https://www.reddit.com/r/rust/comments/gxqr2l/zero_to_production_1_setup_toolchain_ides_ci/)
- url: https://www.lpalmieri.com/posts/2020-06-06-zero-to-production-1-setup-toolchain-ides-ci/
---

## [10][Reddit images downloader in Rust](https://www.reddit.com/r/rust/comments/gy0a21/reddit_images_downloader_in_rust/)
- url: https://www.reddit.com/r/rust/comments/gy0a21/reddit_images_downloader_in_rust/
---
Hi everyone!

Check out, my program that uses Rust async/await features to concurently download images from Reddit, here is the repo [https://github.com/risboo6909/reddit-dl](https://github.com/risboo6909/reddit-dl)

Initially I was inspired by this [post](https://www.reddit.com/r/golang/comments/gn74r2/i_built_a_reddit_image_downloader_in_go_this_is/) whose author had written the same downloader in Go. 

I'm not a professional Rust developer (however I would like to be) but I use Rust for my pet projects when I have free time, so I strongly decided that Rust must have the same downloader (:

This is actually my first serious experience with async/await and tokio and I'm sure my code is far away from ideal, therefor I would be really appreciated to hear any opinion about the code and the tool itself.

Thanks!
## [11][Which framework/library would you recommend if I like Django?](https://www.reddit.com/r/rust/comments/gxzy1j/which_frameworklibrary_would_you_recommend_if_i/)
- url: https://www.reddit.com/r/rust/comments/gxzy1j/which_frameworklibrary_would_you_recommend_if_i/
---
Hey,

obviously I still use Django for some projects, but I am currently checking out Rust and I really like it. I prefer Django, because it has many useful things built-in (admin, ORM, templating, user-system, migration, etc.), and has terrific documentation.

That does not mean I need all this features (e.g. templating and the ORM are not that important), but it is nice to have them in the package and available.

My research brought me mainly to Rocket and Actic-Web, but I am not sure which to choose (I will probably try out some of them anyway).

Cheers and thx for the discussion.
## [12][What's Functional Programming All About?](https://www.reddit.com/r/rust/comments/gxw1sk/whats_functional_programming_all_about/)
- url: https://www.lihaoyi.com/post/WhatsFunctionalProgrammingAllAbout.html
---

