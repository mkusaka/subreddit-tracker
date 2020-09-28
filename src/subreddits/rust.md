# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (39/2020)!](https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.46]](https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting:
# Rules for individuals:

 * Don't create top-level comments; those are for employers.
 * Feel free to reply to top-level comments with on-topic questions.
 * Anyone seeking work should reply to my stickied top-level comment.
 * Meta-discussion should be reserved for the distinguished comment at the very bottom.

# Rules for employers:

 * To find individuals seeking work, see the replies to the stickied top-level comment; you will need to click the "more comments" link at the bottom of the top-level comment in order to make these replies visible.
 * To make a top-level comment you must be hiring directly; no third-party recruiters.
 * One top-level comment per employer. If you have multiple job openings, please consolidate their descriptions or mention them in replies to your own top-level comment.
 * Proofread your comment after posting it and edit it if necessary to correct mistakes.
 * Please base your comment on the following template:

COMPANY: *[Company name; please link to your company's website or careers page.]*

TYPE: *[Full time, part time, internship, contract, etc.]*

DESCRIPTION: *[What does your company do, and what are you using Rust for? How much experience are you seeking and what seniority levels are you hiring for? The more details the better.]*

LOCATION: *[Where are your office or offices located? If your workplace language isn't English, please specify it.]*

REMOTE: *[Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]*

VISA: *[Does your company sponsor visas?]*

CONTACT: *[How can someone get in touch with you?]*
## [3][Rust on the ESP32 &amp; ESP8266 - Building an ecosystem](https://www.reddit.com/r/rust/comments/j18vop/rust_on_the_esp32_esp8266_building_an_ecosystem/)
- url: https://mabez.dev/blog/posts/esp-rust-ecosystem/
---

## [4][Threading WGPU Resource Loading with Rayon](https://www.reddit.com/r/rust/comments/j16a0w/threading_wgpu_resource_loading_with_rayon/)
- url: https://sotrh.github.io/learn-wgpu/intermediate/tutorial13-threading/
---

## [5][Jetbrains is looking for Kotlin + Rust engineers to develop "next-generation IDE platform"](https://www.reddit.com/r/rust/comments/j0rrei/jetbrains_is_looking_for_kotlin_rust_engineers_to/)
- url: https://www.linkedin.com/jobs/view/2151145919
---

## [6][Is there a callback-based IO event library (like LibUV) for Rust (not Mio)?](https://www.reddit.com/r/rust/comments/j17sqs/is_there_a_callbackbased_io_event_library_like/)
- url: https://www.reddit.com/r/rust/comments/j17sqs/is_there_a_callbackbased_io_event_library_like/
---
I'm working on an embedded Linux project which I've started in C, using LibUV for an event-driven design. There are a lot of different varieties of I/O (I2C, USB, flash-based files, serial comms, network, IPC), which is why I went with LibUV. **Not** because of scalability or performance, but because I can write a simple callback for a particular event, register it with the loop, and LibUV takes care of the boilerplate and edge cases. There's also some ZeroMQ in there for IPC. It all fits together quite nicely, especially because anything that doesn't already integrate with LibUV usually has a file descriptor interface I can write a LibUV wrapper around.

I'm starting on some parts that I think would really benefit from being written in Rust, and I'm casting about for how to architect them. All the different areas of responsibility are written as separate processes, so as long as a language has ZeroMQ bindings (Rust does) I'm happy. But I'm struggling to find a Rust crate that does I/O the way I like. (Maybe that's my C brain in action? I'm also very familiar with Python 3.5+'s asyncio if that gives you an idea of how I think.)

I looked at Mio (any Google search for "rust" + "event" or "io" or "async" or just about anything turns it up), but two things suggest it's not what I want:

- It's very much an `epoll()` like interface ie. write your own loop, handle events within it; I strongly prefer writing callbacks against a simple, constrained interface and having the library run the loop and dispatch things for me.
- It explicitly excludes timers, file activity and threading from its concerns. I use timers a lot, I occasionally need to work with file-like things in an event driven way, and if I'm really desperate, I might use threads to wrap up a blocking library to get the job done. (It's a single threaded CPU, so I try not to.)

To put it another way, I struggle to know whether "this Rust crate doesn't do things the way I like" or "this Rust crate does things that are idiomatic in Rust but not in C so it looks strange and confusing to me". Is there a crate that does callback-based event-driven I/O? And if there's not, is it simply because no-one's written one, or because it's not really needed/wise for Rust? If there is, can someone please help me find it?
## [7][Rust 2021 - Don't change!](https://www.reddit.com/r/rust/comments/j1aihh/rust_2021_dont_change/)
- url: https://beyermatthias.de/blog/2020/09/28/rust-2021/
---

## [8][Pandas Dataframe like crate in Rust](https://www.reddit.com/r/rust/comments/j15do8/pandas_dataframe_like_crate_in_rust/)
- url: https://www.reddit.com/r/rust/comments/j15do8/pandas_dataframe_like_crate_in_rust/
---
Hi,
      I was exploring pandas dataframe like lib in Rust and looked at hors, polars but issue is they work with nightly rust. Anything I maybe missing. The goal is to convert json to in memory table for querying.
## [9][frawk: an Awk-like language implemented in Rust](https://www.reddit.com/r/rust/comments/j0y3bh/frawk_an_awklike_language_implemented_in_rust/)
- url: https://www.reddit.com/r/rust/comments/j0y3bh/frawk_an_awklike_language_implemented_in_rust/
---
[repo link](https://github.com/ezrosent/frawk)

The primary goal of this project was to close some gaps in my own usage of Awk, and for me to get some practice in string parsing and compilation. There are still plenty of rough edges, but I figured some folks on here might be interested, or might be able to provide some feedback. Some highlights:

* Both a bytecode interpreter and an LLVM-based JIT.
* Support for CSV and TSV parsing.
* Scripts have [competitive performance](https://github.com/ezrosent/frawk/blob/master/info/performance.md) with purpose-built tools written in Rust and D, even when ignoring compilation time.
* Support for [parallelism](https://github.com/ezrosent/frawk/blob/master/info/parallelism.md), with simple scripts processing CSV at [\&gt;2GB/s](https://github.com/ezrosent/frawk/blob/master/info/performance.md#sum-two-columns) on a laptop.
## [10][Zig's New Relationship with LLVM (creating a new zig-based backend)](https://www.reddit.com/r/rust/comments/j1bvjm/zigs_new_relationship_with_llvm_creating_a_new/)
- url: https://kristoff.it/blog/zig-new-relationship-llvm/
---

## [11][Are we observable yet? An introduction to Rust telemetry (log, tracing)](https://www.reddit.com/r/rust/comments/j0uo0t/are_we_observable_yet_an_introduction_to_rust/)
- url: https://www.lpalmieri.com/posts/2020-09-27-zero-to-production-4-are-we-observable-yet/
---

## [12][Rust Crates that do What the Go Standard library Does](https://www.reddit.com/r/rust/comments/j0ugae/rust_crates_that_do_what_the_go_standard_library/)
- url: https://christine.website/blog/rust-crates-go-stdlib-2020-09-27
---

