# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (16/2020)!](https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (16/2020)?](https://www.reddit.com/r/rust/comments/g0es99/whats_everyone_working_on_this_week_162020/)
- url: https://www.reddit.com/r/rust/comments/g0es99/whats_everyone_working_on_this_week_162020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-16-2020/40880)!
## [3][Better stack fixing for Firefox](https://www.reddit.com/r/rust/comments/g1iw6q/better_stack_fixing_for_firefox/)
- url: https://blog.mozilla.org/nnethercote/2020/04/15/better-stack-fixing-for-firefox/
---

## [4][rust-analyzer Q&amp;A](https://www.reddit.com/r/rust/comments/g1qlrs/rustanalyzer_qa/)
- url: https://www.reddit.com/r/rust/comments/g1qlrs/rustanalyzer_qa/
---
Hey, I'll be doing an webinar/Q&amp;A session about rust-analyzer later today (18:00 CEST, roughly four hours from know). It will be held on zoom, and a recording would be available afterwards. 

Find the link here:

https://rust-experts.com/#qna-ide-1

If you won't be able to participate, but have question, ask it hear and I'll try to answer it later today.
## [5][aeron-rs - efficient reliable UDP and IPC message transport for Rust](https://www.reddit.com/r/rust/comments/g1qrpb/aeronrs_efficient_reliable_udp_and_ipc_message/)
- url: https://www.reddit.com/r/rust/comments/g1qrpb/aeronrs_efficient_reliable_udp_and_ipc_message/
---
Hi, all!

Having fast and reliable messaging is critically important for high performance distributed systems. There are number of message brokers already available for Rust. But what if you need something faster?

There is Aeron system ([https://github.com/real-logic/aeron](https://github.com/real-logic/aeron)) which can transport messages fast through IPC (shared memory) and UDP. But until recently there were no Rust interface for that. Our team from United Traders invested some efforts to obtain so called client library for Aeron written in Rust. So now it is possible to talk from Rust program (through Aeron media driver written in C) with any other Aeron enabled program (written in Rust, C++ or Java).

Our implementation of Aeron client library (called aeron-rs [https://crates.io/crates/aeron-rs](https://crates.io/crates/aeron-rs) ) is a port of C++ library to Rust. While it is completely written in Rust it also inherits design and structure of its C++ ancestor therefore some things are done not in a Rust way. I believe there are plenty improvements which can be done so it will be good if interested people will look at aeron-rs and contribute to it.

As of today aeron-rs 0.1.0 is tested (200+ unit tests, 6 integration test, number of test applications) but was not applied so far in any production environment. You are welcome to give it a try and to propose/implement improvements and fixes.

Comparative performance tests show that (on the same machine) with messages of 32 bytes length C++ Aeron client gives throughput of 6-8 millions messages per second, while aeron-rs just 1-2 millions. On 1K sized message C++ version delivers 500K msgs/sec and aeron-rs 250K msgs/sec. So from performance perspective we also need to catch up and overtake С++ :-)

Thanks!
## [6][Blog Post: From Pratt to Dijkstra](https://www.reddit.com/r/rust/comments/g1p1mn/blog_post_from_pratt_to_dijkstra/)
- url: https://matklad.github.io/2020/04/15/from-pratt-to-dijkstra.html
---

## [7][A Possible New Backend for Rust](https://www.reddit.com/r/rust/comments/g16aje/a_possible_new_backend_for_rust/)
- url: https://jason-williams.co.uk/a-possible-new-backend-for-rust
---

## [8][This Week in Rust 334](https://www.reddit.com/r/rust/comments/g1fj7p/this_week_in_rust_334/)
- url: https://this-week-in-rust.org/blog/2020/04/14/this-week-in-rust-334/
---

## [9][literally... just std::collections literals](https://www.reddit.com/r/rust/comments/g1k2da/literally_just_stdcollections_literals/)
- url: https://www.reddit.com/r/rust/comments/g1k2da/literally_just_stdcollections_literals/
---
Introducing [literally](https://github.com/estk/literally)

I had these macros in many projects, I finally took the time to extract them into a crate. I hope someone else gets some use here.
## [10][Error handling's performance impact on the happy path, the stack, and the heap](https://www.reddit.com/r/rust/comments/g1hj02/error_handlings_performance_impact_on_the_happy/)
- url: https://www.reddit.com/r/rust/comments/g1hj02/error_handlings_performance_impact_on_the_happy/
---
I am currently tuning the performance of a Rust application and have started looking at my error handling code.  My use case for errors is one that I imagine is fairly typical: the overwhelming majority of code paths should not trigger any errors, but nearly all code paths involve some operations that _could_ error.  Based on that, I'm not especially focused on the performance of the program when it encounters an error, but I am _very_ focused on any performance impact the error handling code has on the happy path.

My initial understanding was that the biggest performance pitfall with Rust error handling was with `Box&lt;dyn Error&gt;` – or, more specifically, with the extra allocations that come along with dynamic error types (including, e.g., `anyhow::Error`).  Based on that understanding and a desire to match on many of my error types, I have created custom types for my application's errors in more-or-less the [standard way](https://blog.burntsushi.net/rust-error-handling/#error-handling-with-a-custom-type).

However, Clippy has started to tell me that my Error enums [are too large](https://rust-lang.github.io/rust-clippy/stable/index.html#large_enum_variant).  It suggests boxing them as a way to _improve_ performance (when here I was, thinking that I was being efficient by _not_ allocating them to the heap).  So my specific question is whether I'd be better off following Clippy's advice, better off keeping things as they are, or better off changing my error handling strategy in some other way.

More generally, I'm also curious about how to conceptualize the performance impacts allocating on the heap more broadly.  Before programming in Rust, I mostly worked in JavaScript – I've learned a lot since then, but I'm sure I still have a to learn.  I know that, in general, variables on the stack [will be faster](https://gist.github.com/jFransham/369a86eff00e5f280ed25121454acec1#keep-as-much-as-possible-in-registers) than heap-allocated ones (better cache hit rates, more use of registers, etc).  But Clippy (and logic) seem to suggest that, at some size, allocating to the heap might be more efficient.  For example, I deal with a fairly large data structure (serialized from JSON, ~2kb according to `mem::size_of`) in my program.  Is that the sort of thing that's large enough that I should be storing it on the heap (again, in an effort to prioritize performance, I've been keeping it stack-allocated as much as possible).  I'm less interested in _specific_ answers to these questions than I am on guidance about how to approach these sorts of questions conceptually.

(I know that a big part of the answer is "test and profile" and I'm working on that too.  But I'd like to keep this thread focused, as much as possible, on how to think through these ideas, with the understanding that I'll still need to profile afterwords)
## [11][compress-tools: A Swiss Army Knife for handling compressed data in Rust](https://www.reddit.com/r/rust/comments/g1r2ek/compresstools_a_swiss_army_knife_for_handling/)
- url: https://medium.com/os-systems/compress-tools-a-swiss-army-knife-for-handling-compressed-data-in-rust-9a3c3a75b6b3
---

## [12][Snapshot Testing with Insta 0.16 [Screencast]](https://www.reddit.com/r/rust/comments/g1e0e5/snapshot_testing_with_insta_016_screencast/)
- url: https://www.youtube.com/watch?v=rCHrMqE4JOY
---

