# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (38/2020)!](https://www.reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 356](https://www.reddit.com/r/rust/comments/iu3ge0/this_week_in_rust_356/)
- url: https://this-week-in-rust.org/blog/2020/09/16/this-week-in-rust-356/
---

## [3][Rust fullstack developer - Denmark or remote (Denmark, Germany or France)](https://www.reddit.com/r/rust/comments/iuflg1/rust_fullstack_developer_denmark_or_remote/)
- url: https://www.reddit.com/r/rust/comments/iuflg1/rust_fullstack_developer_denmark_or_remote/
---
[Impero](https://impero.com/) is a compliance company: we deliver a cloud-based solution to help companies across the world comply with existing regulations and facilitate audits.

We are looking for a Rust developer to join our team. We have started an incremental rewrite of our product in Rust, and we need someone with a strong backend experience to help us with this. It is an opportunity to take part in the design of our internal and external APIs, and deliver the best experience to our customers.

Location: Aarhus, Denmark, or remote (citizens or foreign nationals with valid work permit in Denmark, Germany or France).

More details on the [job offer page](https://impero.com/blog/full-stack-web-developer-rust/).
## [4][A high performance csv viewer with cjk and emoji support.](https://www.reddit.com/r/rust/comments/iuh7fd/a_high_performance_csv_viewer_with_cjk_and_emoji/)
- url: https://www.reddit.com/r/rust/comments/iuh7fd/a_high_performance_csv_viewer_with_cjk_and_emoji/
---
Hey guys, I wrote a small tool `csview` for viewing csv file from terminal. Thanks to the increasingly rich rust library, I did not use much code to initially complete the core functions.

[screenshot](https://preview.redd.it/9wro9aidwon51.png?width=1888&amp;format=png&amp;auto=webp&amp;s=21befa13fd599a3a0c5c6f53115e39ee53bc9a93)

Project link: [https://github.com/wfxr/csview](https://github.com/wfxr/csview).

Compared with the commonly used alternative [csvlook](https://github.com/wireservice/csvkit), it can correctly handle CJK characters and emoji. And it is about 80 \~ 400 times faster than `csvlook` depends on different inputs. I'm a rust beginner and I really need advices and feedbacks from you guys!
## [5][j4rs: JavaFX support for Rust WIP](https://www.reddit.com/r/rust/comments/iucytc/j4rs_javafx_support_for_rust_wip/)
- url: https://astonbitecode.github.io/blog/post/j4rs_javafx_support/
---

## [6][What happened to the `ocl` crate?](https://www.reddit.com/r/rust/comments/iugwhs/what_happened_to_the_ocl_crate/)
- url: https://www.reddit.com/r/rust/comments/iugwhs/what_happened_to_the_ocl_crate/
---
The [`ocl` crate](https://github.com/cogciprocate/ocl) provides a thin, but friendly interface to OpenCL from Rust, helping our GPGPU story a lot. However, it hasn’t been updated in over a year, and its pull requests and issues have been left stagnant - though some PRs look like needed fixes. The author has also been largely inactive on GitHub in the past year.

Does anyone know what’s happening with `ocl`? On a related note, is there any good way for the community to continue developing a library if the author stops maintaining it / accepting PRs?
## [7][Possible memory leak in Linux binaries?](https://www.reddit.com/r/rust/comments/iu8aj3/possible_memory_leak_in_linux_binaries/)
- url: https://www.reddit.com/r/rust/comments/iu8aj3/possible_memory_leak_in_linux_binaries/
---
Hi all,

TL;DR: getting weird memory footprint results: Mac and Windows are "fine", and Linux seems wrong.

First time poster here.  I have basically been using Rust for every single one of my pet projects recently.  Despite the fact that my grasp is not quite what I'd like it to be, I love it.

Anyway, I have a project that uses some locality sensitive hashing (LSH) to *roughly* correlate the way users on twitter speak.  The project is [here](https://github.com/twitchax/tweet_analyzer) (with an explanation of the process), and I have an instance running [here](https://tweets.home.ajroney.com/) with an assortment of the top 50 most followed twitter accounts, and a few others.

That instance is running on a single node Linux kubernetes cluster (64 cores, 256 GB of memory), and I noticed that the steady state **memory footprint of the application was around 1.2 GB**, which seemed huge.

The application is pretty simple.  It is:

1. A rocket web server that serves static files for the frontend.
2. 3 actor-based tokio threads that (a) send messages to each other, and (b) spin off other tasks to (i) download tweets and insert into a mongo database (feel free to flame me for that :P ), (ii) analyze and shinglize the tweets and insert the results into the database, and (iii) compute a similarity score between one twitter handle and the rest of the twitter handles.

Therefore, once all of the analysis was complete, there should essentially be 3 actor tasks just sitting around doing nothing, so the memory footprint should be in the 10s of MB.   I used a `static once_cell::Lazy&lt;Arc&lt;AtomicU64&gt;&gt;` counter to check that, indeed, this is the case, and it is the case on all 3 platforms: the "task count" goes up to around 40-60, and slowly drains down to 3 as the work completes.

Anyway, I did some analysis, and I found that the steady state footprint is exactly as expected for Windows and Mac, but huge for Linux (checked against stable/nightly and release/debug).  Here are my results on stable/debug.

[Windows](https://preview.redd.it/0kd2ssmxnln51.png?width=1077&amp;format=png&amp;auto=webp&amp;s=c1e5a61d4468580a0da0350c414ee7d260ee76ac)

&amp;#x200B;

[Mac](https://preview.redd.it/t2bbym92oln51.png?width=1892&amp;format=png&amp;auto=webp&amp;s=84aefc5744f9901a78628c951a95f007702aefb1)

&amp;#x200B;

[WSL \(RES is 1160M\)](https://preview.redd.it/l0q122r3oln51.png?width=1222&amp;format=png&amp;auto=webp&amp;s=02ec85b4fd92500e1a0240695245ceffe713f9f7)

&amp;#x200B;

[Linux Container \(see ta\)](https://preview.redd.it/nrg8p7t7oln51.png?width=1479&amp;format=png&amp;auto=webp&amp;s=b1e56ab9cf10632eed3b44db9886f91ce65c8909)

Basically, wondering if there is something I am missing or doing wrong?

EDIT 1: It doesn't appear to be a leak based on the memory profile.  Appears that around 1.2 MB leak, which obviously does not account for 1.2 GB.  


https://preview.redd.it/m3zmsp4x2on51.png?width=3775&amp;format=png&amp;auto=webp&amp;s=efcb441d9bc742d70e3ae81aebe9b5686fbbdd6a

However, the behavior is still weird.  My best guess is that the application is not returning the resources to the OS, but I would like to figure out a way to force that to happen to nail that down as the culprit.  I did try a different allocator (jemalloc), and the behavior is the same.
## [8][Dropbox open sources protobuf codegen!](https://www.reddit.com/r/rust/comments/ittov9/dropbox_open_sources_protobuf_codegen/)
- url: https://www.reddit.com/r/rust/comments/ittov9/dropbox_open_sources_protobuf_codegen/
---
Hey everyone! At Dropbox we built our own protobuf framework to meet our production needs. We're now open sourcing it!

Back in 2015 when we were building our [Storage System](https://dropbox.tech/infrastructure/inside-the-magic-pocket) we needed a framework that supported zero copy de-serialization, which prompted the creation of our own library. Since, we've began using it for several parts of Dropbox, including our [Sync Engine](https://dropbox.tech/infrastructure/rewriting-the-heart-of-our-sync-engine). Along with zero copy de-serialization we also provide a number of "Rustic" proto extensions. 

Feel free to give it a look, file an issue, open a PR, and stay on the lookout for more open source Rust libraries from Dropbox 

[GitHub](https://github.com/dropbox/pb-jelly) | [crates.io](https://crates.io/crates/pb-jelly)

&amp;#x200B;

P.S. proto service generation coming soon...
## [9][Low-Level Academy, an explorable systems programming course using Rust + Wasm](https://www.reddit.com/r/rust/comments/itzhzl/lowlevel_academy_an_explorable_systems/)
- url: https://lowlvl.org/lesson1.html
---

## [10][How to implement a circuit breaker](https://www.reddit.com/r/rust/comments/iugjc9/how_to_implement_a_circuit_breaker/)
- url: https://www.reddit.com/r/rust/comments/iugjc9/how_to_implement_a_circuit_breaker/
---
I'm implementing a Rest service in Rust using async Actix.
My question is, if I need to implement a circuit breaker (https://microservices.io/patterns/reliability/circuit-breaker.html) and if yes, how I would do that?

I heard once that with async this is not needed, but have not found any source for that.
## [11][Why is (LocalType, ForeignType) foreign?](https://www.reddit.com/r/rust/comments/iuhr24/why_is_localtype_foreigntype_foreign/)
- url: https://www.reddit.com/r/rust/comments/iuhr24/why_is_localtype_foreigntype_foreign/
---
Grinds my gears.
## [12][Writing an x86 bootloader in Rust that can launch vmlinux](https://www.reddit.com/r/rust/comments/itvpem/writing_an_x86_bootloader_in_rust_that_can_launch/)
- url: https://vmm.dev/en/rust/krabs.md
---

