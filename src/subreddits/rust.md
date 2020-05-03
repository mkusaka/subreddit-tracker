# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (18/2020)!](https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g4nu6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.43]](https://www.reddit.com/r/rust/comments/g6v98u/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/g6v98u/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).
# Rules for individuals:

 * Don't create top-level comments; those are for employers.
 * Feel free to reply to top-level comments with on-topic questions.
 * I will create a stickied top-level comment for individuals looking for work.
 * I will create an additional top-level comment for meta discussion.

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
## [3][rlua is now part of the Amethyst organisation](https://www.reddit.com/r/rust/comments/gcoipo/rlua_is_now_part_of_the_amethyst_organisation/)
- url: https://github.com/amethyst/rlua/issues/174
---

## [4][Pijama: A functional programming language written in Rust](https://www.reddit.com/r/rust/comments/gckimp/pijama_a_functional_programming_language_written/)
- url: https://www.reddit.com/r/rust/comments/gckimp/pijama_a_functional_programming_language_written/
---
Hi people!

I've decided to apply some of the concepts I've learned by studying TAPL and write an small functional PL. You can find its source code and more information [here](https://github.com/christianpoveda/pijama/)

I've received a lot of feedback from my blog series about writing an interpreter for the lambda calculus, so I hope you'll enjoy this project too.
## [5][Announcing Routerify: A lightweight and modern router implementation with middleware support for the HTTP library hyper.rs](https://www.reddit.com/r/rust/comments/gco2yr/announcing_routerify_a_lightweight_and_modern/)
- url: https://www.reddit.com/r/rust/comments/gco2yr/announcing_routerify_a_lightweight_and_modern/
---
Hey Rustaceans,

Announcing [Routerify](https://github.com/routerify/routerify): A lightweight and modern **router implementation with middleware** support for the HTTP library [hyper.rs](https://hyper.rs/).

The [Routerify](https://github.com/routerify/routerify) offers the following features:

* 📡 Allows defining **complex routing logic**.
* 🔨 Provides **middleware** support.
* 🌀 Supports **Route Parameters**.
* 🚀 **Fast** as hyper.rs and ready for **production** use.
* 🍺 It supports **any response body type** as long as it implements the [HttpBody](https://docs.rs/hyper/0.13.5/hyper/body/trait.HttpBody.html) trait.
* ❗ Provides a flexible **error handling** strategy.
* 🍗 Exhaustive [examples](https://github.com/routerify/routerify/tree/master/examples) and **well documented**.

&amp;#x200B;

Check it out: [https://github.com/routerify/routerify](https://github.com/routerify/routerify)

Docs: [https://docs.rs/routerify](https://docs.rs/routerify)  


Please **give it a star on Github** if you like it.
## [6][Notice: failure is deprecated](https://www.reddit.com/r/rust/comments/gcbcew/notice_failure_is_deprecated/)
- url: https://www.reddit.com/r/rust/comments/gcbcew/notice_failure_is_deprecated/
---
From the latest release, just a few minutes ago, via issue [#347](https://github.com/rust-lang-nursery/failure/pull/347):

&gt; Notice: failure is deprecated. If you liked failure's API, consider using:

&gt; - [Anyhow](https://github.com/dtolnay/anyhow) is a good replacement for `failure::Error`.
&gt; - [thiserror](https://github.com/dtolnay/thiserror) is a good, near drop-in replacement for `#[derive(Fail)]`.

While it didn't last, I think it was clearly an important step in Rust's error handling evolution.
## [7][Microsoft is making tools for Rust to develop under Windows](https://www.reddit.com/r/rust/comments/gcdrb7/microsoft_is_making_tools_for_rust_to_develop/)
- url: https://www.zdnet.com/article/microsoft-our-rust-programming-language-windows-runtime-library-is-now-in-preview/
---

## [8][An interesting use of where clause with super traits](https://www.reddit.com/r/rust/comments/gcgj4a/an_interesting_use_of_where_clause_with_super/)
- url: https://phaazon.net/blog/where-clause-super-trait
---

## [9][Rust jumps to #21 on the TIOBE Index!](https://www.reddit.com/r/rust/comments/gceiip/rust_jumps_to_21_on_the_tiobe_index/)
- url: https://www.tiobe.com/tiobe-index/
---

## [10][gib: a .gitignore bootstrapper](https://www.reddit.com/r/rust/comments/gckbi3/gib_a_gitignore_bootstrapper/)
- url: https://github.com/DavSanchez/gib
---

## [11][Damascus: a CLI tool for interleaving byte streams](https://www.reddit.com/r/rust/comments/gcpsv3/damascus_a_cli_tool_for_interleaving_byte_streams/)
- url: https://www.reddit.com/r/rust/comments/gcpsv3/damascus_a_cli_tool_for_interleaving_byte_streams/
---
[damascus on github](https://github.com/PicoJr/damascus)

&gt; Note: This a follow up of my [reddit post](https://www.reddit.com/r/rust/comments/gbk1ta/multiplexinginterleaving_2_streams_with_fallback) where I asked for advice.

I'm happy with the result:

* performances are decent: 2.02GiB/s for multiplexing 1.66GiB (= 2 streams) worth of data as 1KB chunks. (test made with `&gt; /dev/null`, no disk IO involved)
* it looks fairly idiomatic: iterate both streams as `chunks`, `zip_longuest` them together, `flat_map` the resulting iterator and write the result to stdout.

I'm still looking forward improving this tool =)

&gt; Note: I had a x2 speedup by wrapping stdout in `BufWriter`
&gt; cf &lt;https://www.reddit.com/r/rust/comments/dogxk8/why_does_buffering_the_already_buffered_stdout/&gt;
## [12][Help with implementing to_string() for a struct from someone else's crate](https://www.reddit.com/r/rust/comments/gcpmli/help_with_implementing_to_string_for_a_struct/)
- url: https://www.reddit.com/r/rust/comments/gcpmli/help_with_implementing_to_string_for_a_struct/
---
Hello, I'm trying to convert a struct from a 3rd party crate (in my case Backtrace) to a string. Looks like it doesn't have the function implemented, so I'm getting this error:

    error[E0599]: no method named `to_string` found for struct `backtrace::capture::Backtrace` in the current scope
      --&gt; src/program.rs:79:34
       |
    79 |     println!("{}", Backtrace::new().to_string());
       |                                     ^^^^^^^^^ method not found in `backtrace::capture::Backtrace`
       |
       = note: the method `to_string` exists but the following trait bounds were not satisfied:
               `backtrace::capture::Backtrace : std::string::ToString`
    error: aborting due to previous error

I'm not exactly sure how to go about implementing the function, this doesn't seem to work:

    use backtrace::capture::Backtrace;
    
    impl ToString for Backtrace {
        fn to_string(&amp;self) {
            println!("aaaa");
        }
    }



    error[E0603]: module `capture` is private
       --&gt; src/program.rs:8:16
        |
    8   | use backtrace::capture::Backtrace;
        |                ^^^^^^^ this module is private
        |
    note: the module `capture` is defined here
       --&gt; /home/xxx/.cargo/registry/src/github.com-1ecc6299db9ec823/backtrace-0.3.46/src/lib.rs:100:9
        |
    100 |         mod capture;
        |         ^^^^^^^^^^^^
    error[E0119]: conflicting implementations of trait `std::string::ToString` for type `backtrace::capture::Backtrace`:
       --&gt; src/program.rs:153:1
        |
    153 | impl ToString for Backtrace {
        | ^^^^^^^^^^^^^^^^^^^^^^^^^^^
        |
        = note: conflicting implementation in crate `alloc`:
                - impl&lt;T&gt; std::string::ToString for T
                  where T: std::fmt::Display, T: ?Sized;
        = note: upstream crates may add a new impl of trait `std::fmt::Display` for type `backtrace::capture::Backtrace` in future versions
    error[E0117]: only traits defined in the current crate can be implemented for arbitrary types
       --&gt; src/program.rs:153:1
        |
    153 | impl ToString for Backtrace {
        | ^^^^^^^^^^^^^^^^^^---------
        | |                 |
        | |                 `backtrace::capture::Backtrace` is not defined in the current crate
        | impl doesn't use only types from inside the current crate
        |
        = note: define and implement a trait or new type instead

Any ideas how to do this right?
