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
## [3][Iced: a high performance and correct x86 (16/32/64-bit) instruction decoder, disassembler and assembler written in Rust](https://www.reddit.com/r/rust/comments/iw9zyj/iced_a_high_performance_and_correct_x86_163264bit/)
- url: https://github.com/0xd4d/iced
---

## [4][Bevy 0.2](https://www.reddit.com/r/rust/comments/iw1yyp/bevy_02/)
- url: https://bevyengine.org/news/bevy-0-2
---

## [5][Throw-away Code](https://www.reddit.com/r/rust/comments/iwbp4d/throwaway_code/)
- url: https://vorner.github.io/2020/09/20/throw-away-code.html
---

## [6][mouse-rs: Library for controlling your mouse in rust](https://www.reddit.com/r/rust/comments/iwac0l/mousers_library_for_controlling_your_mouse_in_rust/)
- url: https://github.com/DankDumpster/mouse-rs
---

## [7][Ouroboros, a crate for easily creating self-referencing structs.](https://www.reddit.com/r/rust/comments/iw56vl/ouroboros_a_crate_for_easily_creating/)
- url: https://www.reddit.com/r/rust/comments/iw56vl/ouroboros_a_crate_for_easily_creating/
---
[https://crates.io/crates/ouroboros](https://crates.io/crates/ouroboros). Due to the `rental` crate no longer being maintained, I decided to take a stab at making a crate with similar features but with what I believe is a simpler and more flexible API.
## [8][A month and a half ago, I hadn't written a line of Rust code. Today, I published my first app to crates.io!](https://www.reddit.com/r/rust/comments/ivv94g/a_month_and_a_half_ago_i_hadnt_written_a_line_of/)
- url: https://www.reddit.com/r/rust/comments/ivv94g/a_month_and_a_half_ago_i_hadnt_written_a_line_of/
---
What the title says! A month ago, I made a post asking how to get started with Rust development, whether I should prototype in Python, etc etc.

Today, I'm proud to announce that:

a. I did *not* prototype in Python, and

b. I published my first app to [crates.io](https://crates.io).

I'm really happy about this, this is my first ever big software development project (I'm 19) and the first big thing I've written in Rust. I've fallen in love with this language and I expect to be using it a lot more in the future.

What I've written is a TUI app that queries the [Discogs](https://discogs.com) database using a given username and auth token, and displays that information in a TUI. It also stores that data in an sqlite database, which means apart from querying the Discogs API, it is completely offline. The repository is [here](https://github.com/cartoon-raccoon/cogsy).

I'm still extremely new to Rust and my code will probably make you want to take a quick trip to r/eyebleach, but I'm extremely open to constructive criticism and I hope you can take the time to look through my code and give some feedback on how I can make better use of Rust's features. My code is also full of `.clone()`s and `.unwrap()`s so I could use some feedback on how to better leverage the ownership system instead of cloning everywhere.

From here, I plan to write a GraphQL web API using Rocket, mongodb and juniper, and then delve into systems programming with Rust and C. If anyone can give me advice on this path, I'd be extremely grateful.

Thanks for reading!
## [9][Updated "Learn Wgpu" to 0.6](https://www.reddit.com/r/rust/comments/iw1z6g/updated_learn_wgpu_to_06/)
- url: /r/rust_gamedev/comments/iw1yup/updated_learn_wgpu_to_06/
---

## [10][TL;DR Rust](https://www.reddit.com/r/rust/comments/ivv67v/tldr_rust/)
- url: https://christine.website/blog/TLDR-rust-2020-09-19
---

## [11][gRPC services using Rust. I started learning Rust and thought I’d blog about it as I learn](https://www.reddit.com/r/rust/comments/iwb4g2/grpc_services_using_rust_i_started_learning_rust/)
- url: https://www.swiftdiaries.com/rust/
---

## [12][I started a port of Blueprintjs.com to Yew](https://www.reddit.com/r/rust/comments/ivrmbr/i_started_a_port_of_blueprintjscom_to_yew/)
- url: https://github.com/cecton/blueprint-rs
---

