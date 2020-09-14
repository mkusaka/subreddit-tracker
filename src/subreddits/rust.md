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
## [3][What I Learnt from Benchmarking Http4k, Ktor (Kotlin) and Actix v2, v3 (Rust) Microservices](https://www.reddit.com/r/rust/comments/is9onc/what_i_learnt_from_benchmarking_http4k_ktor/)
- url: https://matej.laitl.cz/bench-rust-kotlin-microservices/
---

## [4][Rust now no. 18 for September on the TIOBE index](https://www.reddit.com/r/rust/comments/isi49d/rust_now_no_18_for_september_on_the_tiobe_index/)
- url: https://www.tiobe.com/tiobe-index/
---

## [5][Why can't lifetimes be elided from structs containing references?](https://www.reddit.com/r/rust/comments/isgjkj/why_cant_lifetimes_be_elided_from_structs/)
- url: https://www.reddit.com/r/rust/comments/isgjkj/why_cant_lifetimes_be_elided_from_structs/
---
So I've been working more with [lifetimes](https://doc.rust-lang.org/1.9.0/book/lifetimes.html#lifetime-elision) recently, and one detail which is not clear to me is why they can be elided from function signatures but not struct declarations.

What I mean is, when declaring structs with references, I almost always find myself declaring a lifetime for the struct which is shared by the references within:

    struct Foo&lt;'a&gt; {
         bar: &amp;'a Bar
    }

Maybe I'm missing something, but it seems like this is the obvious "default case", so my question is:

1. Could this be assumed by the compiler to save a bit of boilerplate?  If not why not?

2. If it could be assumed as the default, what's the reasoning for requiring this to be typed explicitly each time when it's elided in the case of functions?
## [6][[knurling] One Month of Knurling-rs](https://www.reddit.com/r/rust/comments/isjqo5/knurling_one_month_of_knurlingrs/)
- url: https://ferrous-systems.com/blog/knurling-one-month/
---

## [7][rudiments: A drum machine written in Rust](https://www.reddit.com/r/rust/comments/is0inl/rudiments_a_drum_machine_written_in_rust/)
- url: https://www.reddit.com/r/rust/comments/is0inl/rudiments_a_drum_machine_written_in_rust/
---
I created a drum machine as a vehicle for learning Rust.

[https://github.com/jonasrmichel/rudiments](https://github.com/jonasrmichel/rudiments)

I'm very interested in feedback on overall structure, data flow, and types.

The project includes some example patterns -- give it a try!
## [8][I spy that Pyo3 0.12 is released!](https://www.reddit.com/r/rust/comments/is74w7/i_spy_that_pyo3_012_is_released/)
- url: https://www.reddit.com/r/rust/comments/is74w7/i_spy_that_pyo3_012_is_released/
---
Everyone has been hard at work improving python-rust interop, culminating in a new release:

[https://github.com/PyO3/pyo3/releases/tag/v0.12.0](https://github.com/PyO3/pyo3/releases/tag/v0.12.0)

As usual with each new release, it's simpler and safer than the previous version. I'm constantly impressed by everyone's efforts striving to make python-rust interop as seamless as possible.

&amp;#x200B;

Personally I'm particularly excited to try out the initial support for ADTs/Enums (in the rust world) that map to Unions in python. Read all about that here:

[https://pyo3.rs/v0.12.0/conversions/traits.html#deriving-a-hrefhttpsdocsrspyo3latestpyo3conversiontraitfrompyobjecthtmlfrompyobjecta-for-enums](https://pyo3.rs/v0.12.0/conversions/traits.html#deriving-a-hrefhttpsdocsrspyo3latestpyo3conversiontraitfrompyobjecthtmlfrompyobjecta-for-enums)
## [9][Rust 2021](https://www.reddit.com/r/rust/comments/is3ucs/rust_2021/)
- url: https://hsivonen.fi/rust2021/
---

## [10][Outdated openssl-sys](https://www.reddit.com/r/rust/comments/isjuoy/outdated_opensslsys/)
- url: https://www.reddit.com/r/rust/comments/isjuoy/outdated_opensslsys/
---
In my time with rust I now have come over quite a few crates that depend on an outdated version of openssl-sys(0.9.\*). The problem is that they all fail to build with current versions of libssl-dev.

Is there a better way than just forking all the crates and bumping the version manually?
## [11][Cargo with less noise](https://www.reddit.com/r/rust/comments/is9o7x/cargo_with_less_noise/)
- url: https://github.com/alopatindev/cargo-limit
---

## [12][Bunt: a simple, macro-based terminal coloring/styling library based on `termcolor`](https://www.reddit.com/r/rust/comments/is1x2o/bunt_a_simple_macrobased_terminal_coloringstyling/)
- url: https://github.com/LukasKalbertodt/bunt
---

