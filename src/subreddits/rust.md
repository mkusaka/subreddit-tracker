# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (6/2020)!](https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/)
- url: https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/
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
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.41]](https://www.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/ecxd62/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).

#### Rules for individuals:

* Don't create top-level comments; those are for employers.
* Feel free to reply to top-level comments with on-topic questions.
* I will create a stickied top-level comment for individuals looking for work.
* I will create an additional top-level comment for meta discussion.

#### Rules for employers:

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
## [3][Glad to see that Programming Rust will soon have a Second Edition](https://www.reddit.com/r/rust/comments/f0fjv1/glad_to_see_that_programming_rust_will_soon_have/)
- url: https://i.redd.it/wqa8y53m0kf41.jpg
---

## [4][[Advice] Is Rust suitable for business needs?](https://www.reddit.com/r/rust/comments/f0p05u/advice_is_rust_suitable_for_business_needs/)
- url: https://www.reddit.com/r/rust/comments/f0p05u/advice_is_rust_suitable_for_business_needs/
---
Hello Rustaceans. I will try to keep it short and clean:

* Senior Javascript developer, web front-end
* Soon we will start a new project with my team, possibly in WebAssembly
* For a team with many **juniors** and mid-seniors developers, Javascript starts to be a **mess** over time.
* Despite that, I think Javascript is somehow a good fit for **business logic** (no memory worries, get things done quickly), but for example: dynamic typing doesn't really help keep function definitions as solid as they should be, over time.
* I took a short look over the Rust book, and the language looks amazingly well designed, however as a noob it's not clean to me if it suits well for (non memory related) business logic stuff.

So, excluding Rust's ownership principles and other memory related features, what Rust will give us more over Javascript or Typescript (which is another available choice) in the **long term**?

Thank you!
## [5][Bringing async/await to embedded Rust](https://www.reddit.com/r/rust/comments/f0ckiv/bringing_asyncawait_to_embedded_rust/)
- url: https://ferrous-systems.com/blog/embedded-async-await/
---

## [6][Yet another crate that work with order of bytes (endianness)](https://www.reddit.com/r/rust/comments/f0qr9x/yet_another_crate_that_work_with_order_of_bytes/)
- url: https://www.reddit.com/r/rust/comments/f0qr9x/yet_another_crate_that_work_with_order_of_bytes/
---
Hi, was working with embedded system that communicate via packed struct sent as little-endian bytes. I didn't found any crate that easy allow byteorder-encoding/decoding structs with multiple fields so tried to create own. This crate has **no-dependencies** and can be used in **no-std** environments.   


This is also my first derive crate - It would be nice to hear some constructive criticism or get know it was useful for somebody.   


github: [https://github.com/xoac/endian\_codec](https://github.com/xoac/endian_codec)  
crates-io: [https://crates.io/crates/endian\_codec](https://crates.io/crates/endian_codec)
## [7][Zero cost abstractions in action](https://www.reddit.com/r/rust/comments/f0cjw7/zero_cost_abstractions_in_action/)
- url: https://idursun.com/posts/rust_zero_cost_abstractions_in_action/
---

## [8][Rustacean Station Triple Feature: Rust Release Engineering; Developing Rust's Developer Tools; Rust in Latin America [RustFest 2019 Interviews]](https://www.reddit.com/r/rust/comments/f0kpsx/rustacean_station_triple_feature_rust_release/)
- url: https://rustacean-station.org/episode/012-pietro-pascal-santiago/
---

## [9][JetBrains is doing their 2020 survey](https://www.reddit.com/r/rust/comments/f0cprf/jetbrains_is_doing_their_2020_survey/)
- url: https://surveys.jetbrains.com/s3/a18-developer-ecosystem-survey-2020
---

## [10][2020-02-06 Compiler Team Triage Meeting | Inside Rust Blog](https://www.reddit.com/r/rust/comments/f0fye6/20200206_compiler_team_triage_meeting_inside_rust/)
- url: https://blog.rust-lang.org/inside-rust/2020/02/07/compiler-team-meeting.html
---

## [11][[FOSDEM] cargo deny - Fearlessly update your dependencies](https://www.reddit.com/r/rust/comments/f0br52/fosdem_cargo_deny_fearlessly_update_your/)
- url: https://fosdem.org/2020/schedule/event/rust_cargo_deny/
---

## [12][Release Notes for rs-pbrt v0.8.0 (Rust implementation of Physical Based Rendering)](https://www.reddit.com/r/rust/comments/f0gscr/release_notes_for_rspbrt_v080_rust_implementation/)
- url: https://www.rs-pbrt.org/blog/v0-8-0-release-notes
---

