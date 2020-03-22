# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (12/2020)!](https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.42]](https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://en.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/).

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
## [3][C++20 &amp; Rust on Static vs Dynamic Generics](https://www.reddit.com/r/rust/comments/fmo4zb/c20_rust_on_static_vs_dynamic_generics/)
- url: https://youtu.be/olM7o_oYML0
---

## [4][Memory database](https://www.reddit.com/r/rust/comments/fmvpvs/memory_database/)
- url: https://github.com/bloq/memds
---

## [5][Announcing Tokei 11: Now over 40–60% faster](https://www.reddit.com/r/rust/comments/fmfpfu/announcing_tokei_11_now_over_4060_faster/)
- url: https://github.com/XAMPPRocky/tokei/releases/tag/v11.0.0
---

## [6][Can wrapper types (Box, Rc, RefCell) be implemented in safe Rust?](https://www.reddit.com/r/rust/comments/fmzmsx/can_wrapper_types_box_rc_refcell_be_implemented/)
- url: https://www.reddit.com/r/rust/comments/fmzmsx/can_wrapper_types_box_rc_refcell_be_implemented/
---
My intuition tells me the RefCell is impossible, because of multiple mutable borrowings, but Rc and Box might be possible.

I would like to see code examples, if it's possible
## [7][TerraMach a GUI for devices and desktop](https://www.reddit.com/r/rust/comments/fmnwza/terramach_a_gui_for_devices_and_desktop/)
- url: https://github.com/lykhonis/terramach
---

## [8][Yet another gui crate](https://www.reddit.com/r/rust/comments/fmlzel/yet_another_gui_crate/)
- url: https://www.reddit.com/r/rust/comments/fmlzel/yet_another_gui_crate/
---
I recently published a new desktop gui crate. It’s basically rust bindings to fltk. Here is the link to the repo (it also contains some examples and screenshots):

https://github.com/MoAlyousef/fltk-rs

It’s still in alpha  with a lot of unsafe code due to interfacing with C and C++. 

I would appreciate any code reviews and any testers willing to try it out. 

An advantage of fltk is that it builds into a small and self-contained binary with no external dependencies needed for deployment. 

Downsides include bad right-to-left support, no text alignment inside input and output widgets, and is kinda ugly (think windows 98!), but is customizable. 

However it gets the job done and is used by several applications and Linux distros. https://en.m.wikipedia.org/wiki/FLTK#Use

It also has a permissive modified LGPL license to allow static linking, although the fltk-rs crate is licensed under the MIT license.
## [9][Telegram bot for Reddit weekly popular posts, powered by Rust](https://www.reddit.com/r/rust/comments/fmz309/telegram_bot_for_reddit_weekly_popular_posts/)
- url: https://www.reddit.com/r/rust/comments/fmz309/telegram_bot_for_reddit_weekly_popular_posts/
---
Hi,

I wanted to share a Rust project that I have been working on in my spare time. It's a Telegram Bot that allows you to subscribe to Reddit weekly popular posts of your choice. You can also choose the day of the week and time when you want to receive the posts.

The bot is available here [https://t.me/RedditWeeklyBot](https://t.me/RedditWeeklyBot)  
And the code is available here [https://github.com/aldis-ameriks/reddit-bot](https://github.com/aldis-ameriks/reddit-bot)

I'm rather new to Rust, so feedback is more than welcome.

P.S. I hope you're all doing well. I'm currently in self-isolation with mild symptoms. Writing some Rust and watching Smallville.

Stay safe!
## [10][Secure Rust Guidelines](https://www.reddit.com/r/rust/comments/fmdsc3/secure_rust_guidelines/)
- url: https://anssi-fr.github.io/rust-guide/
---

## [11][`never` crate - stable alternative to !](https://www.reddit.com/r/rust/comments/fmllgx/never_crate_stable_alternative_to/)
- url: https://www.reddit.com/r/rust/comments/fmllgx/never_crate_stable_alternative_to/
---
I'd like to announce the new [never](https://crates.io/crates/never) crate. It provides the `Never` type, which is a stable version of the currently-unstable [`!`](https://doc.rust-lang.org/std/primitive.never.html) type in Rust. It provides much (though not all) of the same functionality.

**Why not use the existing [void](https://crates.io/crates/void) crate?** `void` is great, but unfortunately isn't actively maintained anymore. The last commit was four years ago, and issues have gone unresponded to since August, 2016. It doesn't implement `Hash` or `Eq`, and due to lack of maintenance, never will. We use `never` in production, and plan to actively maintain it until the built-in `!` type is stabilized.

**Why not use the existing standard library [`Infallible`](https://doc.rust-lang.org/nightly/std/convert/enum.Infallible.html) type?** It doesn't implement `Hash`.
## [12][Implementing Hash, Eq of Box&lt;dyn Trait&gt; using hash of Struct&lt;S&gt;: Trait](https://www.reddit.com/r/rust/comments/fmyku6/implementing_hash_eq_of_boxdyn_trait_using_hash/)
- url: https://www.reddit.com/r/rust/comments/fmyku6/implementing_hash_eq_of_boxdyn_trait_using_hash/
---
TL;DR Can I implement `Hash` and `Eq` of `Box&lt;dyn Trait&gt;` by hashing `Struct&lt;S&gt;: Trait`? Is it possible that hashes of `Struct&lt;U&gt;` and `Struct&lt;V&gt;` would collide?

Hi! I want to have a `Vec` or a `HashSet` of `Box&lt;dyn Trait&gt;` and I sometimes need to remove an item, so I need to implement `Hash` and `Eq`. The structs implementing `Trait` are [`actix::Addr`](https://docs.rs/actix/0.10.0-alpha.2/actix/struct.Addr.html) of various actors (that are not the same type). The only relevant difference of structs implementing `Trait` is whether they are an address of the same actor of not. Is it possible that their hashes would collide?
