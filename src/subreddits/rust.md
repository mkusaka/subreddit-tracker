# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (42/2020)!](https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (42/2020)?](https://www.reddit.com/r/rust/comments/j9l0o7/whats_everyone_working_on_this_week_422020/)
- url: https://www.reddit.com/r/rust/comments/j9l0o7/whats_everyone_working_on_this_week_422020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-42-2020/49997?u=llogiq)!
## [3][Getting back to C++ after Rust is a pain.](https://www.reddit.com/r/rust/comments/ja5aoe/getting_back_to_c_after_rust_is_a_pain/)
- url: https://www.reddit.com/r/rust/comments/ja5aoe/getting_back_to_c_after_rust_is_a_pain/
---
I would like to share some thoughts about Rust with you guys and maybe hear from you what you think about this.

I'm a software engineer working in the Visual Effects Industry, we write tools for making CG artists work easier and more efficient.  As you can imagine, C++ is dominant there for everything other than the pipeline, which is mostly Python. 

I'm an experienced Python programmer, I write Python for 10+ years now, but I also wrote a good amount of C++ code **with an important note** \- it was mostly plugins for other software with huge and powerful SDKs which cover 95% of your c++ needs. I never had to think about which library to use for parallel programming or string manipulation, logging etc and I basically followed common practices of those mature SDKs (Autodesk Maya, SideEffects Houdini).

About two years ago I started learning Rust and as many of us, fall in love with it. I spent a lot of time learning it and wrote several command-line tools at work (for which I would've previously chose Python). My obsession with Rust even made me write a simple bindings to some of our studio's C++ APIs and I had fun with it. I'm eagerly grasping any opportunity to writing Rust at work, but it's very hard to find application for it other than for simple utilities.

Sadly, Rust has no place in our industry and for me to grow as a developer I need to dive into C++ at a deeper level. I recently made a few attempts and oh boy it's pain. My brain got so Rusty that when I'm staring at some simple C++ code I'm questioning what a hell I'm doing with my life and do I really need all this suffering? :)  Dozens of ways to initialize variables, completely unreadable STL and function signatures, unreadable error messages,  implicit copying,  not clear who owns what, segfaults  and so on....

And now there's C++20. I watched a lot of recent CppCon videos and I'm seeing the shaping of a completely different language within this new standard (which is nonetheless built on top of the same old and crazily complicated codebase). C++20 is finally getting some features we love in Rust - ranges, contracts, async, modules etc. Unfortunately, in visual effects, we won't be able to use the new C++ standard anytime soon (in 5 years maybe?). Will it make C++ easier and safer to use? Maybe. But can you imagine where Rust would be in 5 years? I'm guessing far ahead!

I like my job, and I know that learning C++ helps my carrier, but gosh it's so hard after Rust. I'm finding excuses to not do so and just enjoying Rust when I can.

My takeaway - if you're a good C++ programmer, learning Rust will make you even better coder( or draw away from C++ forever LOL). But if Rust is your first language, learning C++ is a nightmare! Thank you for your attention!
## [4][SixtyFPS: New GUI framework written in Rust (alpha)](https://www.reddit.com/r/rust/comments/jac66t/sixtyfps_new_gui_framework_written_in_rust_alpha/)
- url: https://github.com/sixtyfpsui/sixtyfps/
---

## [5][Async-graphql 2.0.0 version is finally released!](https://www.reddit.com/r/rust/comments/ja7bne/asyncgraphql_200_version_is_finally_released/)
- url: https://www.reddit.com/r/rust/comments/ja7bne/asyncgraphql_200_version_is_finally_released/
---
`Async-graphql` is a high-performance server-side library that supports all GraphQL specifications. The `2.0.0` version has undergone a lot of refactoring, and it now has better performance and easier-to-use API.

Thanks to `Koxiaet`, and everyone who contributed to this library! 😁

Repo: [https://github.com/async-graphql/async-graphql](https://github.com/async-graphql/async-graphql)

Benchmark: [https://github.com/async-graphql/benchmark](https://github.com/async-graphql/benchmark)

## Features

* Fully supports async/await
* Type safety
* Rustfmt friendly (Procedural Macro)
* Custom scalars
* Minimal overhead
* Easy integration (actix\_web, tide, warp, rocket ...)
* Upload files (Multipart request)
* Subscriptions (WebSocket transport)
* Custom extensions
* Apollo Tracing extension
* Limit query complexity/depth
* Error Extensions
* Apollo Federation
* Batch Queries
* Apollo Persisted Queries
## [6][Rust is the second highest paid programming language in 2020](https://www.reddit.com/r/rust/comments/j9w5j2/rust_is_the_second_highest_paid_programming/)
- url: https://expertain.net/what-are-the-highest-paid-programming-languages/
---

## [7][Real-time programming in Rust?](https://www.reddit.com/r/rust/comments/jadbzs/realtime_programming_in_rust/)
- url: https://www.reddit.com/r/rust/comments/jadbzs/realtime_programming_in_rust/
---
I’ve always been interested in doing (interactive) audio coding, such as synthesizers and the like. In fact, it was one of the reasons why I started with C++. Doing real-time (audio) programming almost inevitably involves multithreading, as you generally need one real-time thread to render the audio without any delays. Any work such as loading samples should run on background threads.

As the real-time thread has to return in milliseconds and can never miss a deadline, it often involves lockfree data structures and you’re quite limited in what you can do there. Blocking behaviour is not allowed in the real-time thread, meaning that any function that has an unknown duration cannot be used there. This includes more than you’d initially think:

* Mutexes (which can lead to *priority inversion*)
* Disk access
* Practically all system calls
* Object creation/destruction (as they might call memory allocation functions)

This is quite hard and I was wondering if Rust might make this easier (async/await, no\_std and custom allocators are things that come to mind here). In C++, it’s basically up to you to make sure you’re not using functions which can block in your audio thread. One thing I’ve been thinking about is whether it would be possible to check at compile time whether your real-time code isn't calling any blocking functions (functions could be annotated with "RealTimeSafe"/"NonBlocking" or a subset of the no\_std)?

Also, are there any people here using Rust for lockfree/real-time programming? If so, what is your experience? Are there good/easy libraries for lockfree data structures (Crossbeam)? What is missing/could be improved?
## [8][Pre-RFC: Split up Tier-2 into two tiers](https://www.reddit.com/r/rust/comments/ja2wsk/prerfc_split_up_tier2_into_two_tiers/)
- url: https://internals.rust-lang.org/t/pre-rfc-split-up-tier-2-platforms/13196
---

## [9][Further info on Rust support in SailfishOS](https://www.reddit.com/r/rust/comments/jacvmk/further_info_on_rust_support_in_sailfishos/)
- url: https://www.reddit.com/r/rust/comments/jacvmk/further_info_on_rust_support_in_sailfishos/
---
Taken from [official blog post](https://blog.jolla.com/sailfish-os-pallas-yllastunturi-is-now-available/):

&gt; Developer Offering
&gt; 
&gt; We’ve included experimental support for the Rust programming language in this release. This addition is related to the upcoming browser upgrades, but we will experiment with using Rust in other areas as well. We welcome and appreciate all comments from the community about the behaviour and use of this addition and encourage you to share your experiences on the forum.
&gt; 
&gt; In preparation for future needs we are also releasing experimental support for 64-bit hardware (aarch64). This experiment is aimed at potential upcoming devices and is not currently used by any existing hardware. It should be noted by our development community that we do not consider the aarch64 ABI stable yet. Recompiling applications might be needed later when the 64-bit architecture becomes officially supported.
&gt; 
&gt; We invite our partners and community to take this into use where devices support it (e.g. HW ports) and to share your feedback via our forum (https://forum.sailfishos.org/).
&gt; 
&gt; We’ve also updated our Notification API so that the interface now includes progress information that can be displayed on the Events view.

This is a follow-up post to my post from 12 days ago: https://www.reddit.com/r/rust/comments/j35d01/mobile_operating_system_sailfishos_now_supports/

I'm just happy to see them write about the browser engine (a long standing problem with SFOS), the possibility of future devices with 64bit support, and general support of the Rust language. :)
## [10][blades - a (very) fast static site generator](https://www.reddit.com/r/rust/comments/ja1lpq/blades_a_very_fast_static_site_generator/)
- url: https://github.com/grego/blades
---

## [11][aarch64-pc-windows-msvc has been promoted to Tier 2 Development Platform! 🎉](https://www.reddit.com/r/rust/comments/j9wocy/aarch64pcwindowsmsvc_has_been_promoted_to_tier_2/)
- url: https://github.com/rust-lang/rust/pull/75914
---

## [12][Looking for 2D geometry library in rust](https://www.reddit.com/r/rust/comments/ja8nwi/looking_for_2d_geometry_library_in_rust/)
- url: https://www.reddit.com/r/rust/comments/ja8nwi/looking_for_2d_geometry_library_in_rust/
---
Hi all,

I need a library to deal with basic 2D geometry functions, such as handling shapes, doing boolean operations with polygons (intersection, xor, etc). 

Can anyone suggest a library that implement theses functions for me?
