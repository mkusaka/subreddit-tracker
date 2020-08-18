# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.45]](https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://old.reddit.com/r/rust/comments/gz8ut5/official_rrust_whos_hiring_thread_for_jobseekers/).

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
## [3][Amethyst Engine v0.15.1 — last specs release](https://www.reddit.com/r/rust/comments/ibvheq/amethyst_engine_v0151_last_specs_release/)
- url: https://amethyst.rs/posts/release-0.15.1
---

## [4][My first crate: MozWire: A MozillaVPN configurator, finally giving Linux, macOS and BSD users (among others) access to MozillaVPN](https://www.reddit.com/r/rust/comments/iblon7/my_first_crate_mozwire_a_mozillavpn_configurator/)
- url: https://github.com/NilsIrl/MozWire
---

## [5][Announcing async-raft v0.5.0: Rebased onto Tokio &amp; Renamed to `async-raft`](https://www.reddit.com/r/rust/comments/ibsxxd/announcing_asyncraft_v050_rebased_onto_tokio/)
- url: https://www.reddit.com/r/rust/comments/ibsxxd/announcing_asyncraft_v050_rebased_onto_tokio/
---
I am pleased to announce the release of `async-raft` `v0.5.0`.

The TLDR:
* Everything Is Now Based on Tokio
* Singular &amp; Well Defined API (thanks to async/await, async_trait &amp; tokio)
* Automatic Log Compaction
* Joint Consensus Overhaul
* Project Renamed to `async-raft` (edit: was actix-raft previously)

For more info, here are the release notes https://github.com/async-raft/async-raft/releases/tag/v0.5.0
## [6][Does Rust have anything similar to Erlang supervision trees?](https://www.reddit.com/r/rust/comments/ibxv67/does_rust_have_anything_similar_to_erlang/)
- url: https://www.reddit.com/r/rust/comments/ibxv67/does_rust_have_anything_similar_to_erlang/
---
And if it does, how does it integrate with the rest of the ecosystem. I'm working in a new project and the current implementation is one with Go built but I'm working in a proof of concept with Elixir and Rust to see if I can get a better result.

It's a application running on a device at the client so it must be really safe and the supervision tree comes in handy at this moment.
## [7][Control Flow Guard for Clang/LLVM and Rust - Microsoft Security Response Center](https://www.reddit.com/r/rust/comments/ibiy42/control_flow_guard_for_clangllvm_and_rust/)
- url: https://msrc-blog.microsoft.com/2020/08/17/control-flow-guard-for-clang-llvm-and-rust/
---

## [8][Why does MIRI complain about undefined behaviour regarding stacked borrows in this example program?](https://www.reddit.com/r/rust/comments/ibzqe5/why_does_miri_complain_about_undefined_behaviour/)
- url: https://www.reddit.com/r/rust/comments/ibzqe5/why_does_miri_complain_about_undefined_behaviour/
---
I have some unsafe code that I am testing with MIRI, but I don't know what the problem is exactly.

I have narrowed it down to the following example program with a self-referential struct (without moving it):

[https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=5b294efbcb19aab9ddc51dd0d0f67ffd](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=5b294efbcb19aab9ddc51dd0d0f67ffd)

The error I get when running with `cargo +nightly miri run` is:

&amp;#x200B;

```

error: Undefined Behavior: trying to reborrow for SharedReadWrite, but parent tag &lt;3963&gt; does not have an appropriate item in the borrow stack

\--&gt; /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/core/src/ptr/unique.rs:120:18

|

120 |  unsafe { &amp;\*self.as\_ptr() }

|   \^\^\^\^\^\^\^\^\^\^\^\^\^\^\^ trying to reborrow for SharedReadWrite, but parent tag &lt;3963&gt; does not have an appropriate item in the borrow stack

|

= help: this indicates a potential bug in the program: it performed an invalid operation, but the rules it violated are still experimental

= help: see [https://github.com/rust-lang/unsafe-code-guidelines/blob/master/wip/stacked-borrows.md](https://github.com/rust-lang/unsafe-code-guidelines/blob/master/wip/stacked-borrows.md) for further information

= note: inside \`std::ptr::Unique::&lt;A&gt;::as\_ref\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/core/src/ptr/unique.rs:120:18

= note: inside \`alloc::alloc::box\_free::&lt;A&gt;\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/alloc/src/alloc.rs:321:32

= note: inside \`std::intrinsics::drop\_in\_place::&lt;std::boxed::Box&lt;A&gt;&gt; - shim(Some(std::boxed::Box&lt;A&gt;))\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/core/src/ptr/mod.rs:184:1

note: inside \`main\` at src/main.rs:31:1

\--&gt; src/main.rs:31:1

|

31 | }

| \^

= note: inside \`&lt;fn() as std::ops::FnOnce&lt;()&gt;&gt;::call\_once - shim(fn())\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/core/src/ops/function.rs:233:5

= note: inside \`std::sys\_common::backtrace::\_\_rust\_begin\_short\_backtrace::&lt;fn(), ()&gt;\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/std/src/sys\_common/backtrace.rs:137:18

= note: inside closure at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/std/src/rt.rs:66:18

= note: inside \`std::ops::function::impls::&lt;impl std::ops::FnOnce&lt;()&gt; for &amp;dyn std::ops::Fn() -&gt; i32 + std::marker::Sync + std::panic::RefUnwindSafe&gt;::call\_once\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/core/src/ops/function.rs:265:13

= note: inside \`std::panicking::r#try::do\_call::&lt;&amp;dyn std::ops::Fn() -&gt; i32 + std::marker::Sync + std::panic::RefUnwindSafe, i32&gt;\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/std/src/panicking.rs:373:40

= note: inside \`std::panicking::r#try::&lt;i32, &amp;dyn std::ops::Fn() -&gt; i32 + std::marker::Sync + std::panic::RefUnwindSafe&gt;\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/std/src/panicking.rs:337:19

= note: inside \`std::panic::catch\_unwind::&lt;&amp;dyn std::ops::Fn() -&gt; i32 + std::marker::Sync + std::panic::RefUnwindSafe, i32&gt;\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/std/src/panic.rs:394:14

= note: inside \`std::rt::lang\_start\_internal\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/std/src/rt.rs:51:25

= note: inside \`std::rt::lang\_start::&lt;()&gt;\` at /Users/rolftimmermans/.rustup/toolchains/nightly-x86\_64-apple-darwin/lib/rustlib/src/rust/library/std/src/rt.rs:65:5

```

Why does MIRI complain about this *specific* example? It is directly related to the reborrow of `A`, because if I remove that part (line 24-28) then MIRI is happy.

I know that the pointer `ptr` to the array would be invalid if the struct `A` would be moved. But in this case it isn't moved... Am I missing something?

(Some context: I am experimenting with unsafe code to integrate garbage collection into a virtual machine for a toy dynamic language. I'm doing some terribly unsafe things, in part just to better understand what is allowed and what isn't.)
## [9][Actix TaskQueue: A generic task queue for Actix](https://www.reddit.com/r/rust/comments/ibuzk8/actix_taskqueue_a_generic_task_queue_for_actix/)
- url: https://github.com/huytd/actix-taskqueue
---

## [10][Announcing SIMple Physics, a set of educational physics simulators for algebra based physics](https://www.reddit.com/r/rust/comments/ibk2rf/announcing_simple_physics_a_set_of_educational/)
- url: https://www.reddit.com/r/rust/comments/ibk2rf/announcing_simple_physics_a_set_of_educational/
---
SIMple Physics is a set of educational physics simulators meant to help learn/teach algebra based physics. It is written in Rust using `ggez`, `specs`, `imgui-rs`, `nphysics`, and `rlua`.

The goal is to make a simulator for each general section of high school physics:

- [Basic Mechanics](https://github.com/mkhan45/SIMple-Mechanics)

- [Universal Gravitation](https://github.com/mkhan45/SIMple-Gravity)

- Electronics and Magnetism

- Waves/Optics

So far, the [basic mechanics](https://github.com/mkhan45/SIMple-Mechanics) and [universal gravitation](https://github.com/mkhan45/SIMple-Gravity) components are complete.

Find out more at [https://mkhan45.github.io/SIMple-Physics!](https://mkhan45.github.io/SIMple-Physics/tabs/about/)

Here are a few cool GIFs:

[Wave](https://mkhan45.github.io/SIMple-Physics/assets/gifs/misc/mechanics/demo4.gif)

[Spinny](https://mkhan45.github.io/SIMple-Physics/assets/gifs/misc/mechanics/demo1.gif)

[Trippy trails](https://mkhan45.github.io/SIMple-Physics/assets/gifs/misc/gravity/relative_trail.gif)

[Cool gravitational systems](https://mkhan45.github.io/SIMple-Physics/assets/gifs/misc/gravity/saveload.gif)

I chose Rust for this project for its performance, cross platform compilation, and ideally to compile to WASM.

I started the project as my high school senior research project but didn't finish in time due to COVID. However, once it became clear that the next semester of school will be largely online, I decided to finish it this summer. Due to college starting, I won't have much time to work on this for the forseeable future, and, honestly, I want to move on as this has been my primary side project for more than a year. Because of that, contributions would be welcome!
## [11][What's everyone working on this week (34/2020)?](https://www.reddit.com/r/rust/comments/ibtzp9/whats_everyone_working_on_this_week_342020/)
- url: https://www.reddit.com/r/rust/comments/ibtzp9/whats_everyone_working_on_this_week_342020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-33-2020/47095?u=llogiq)!

(created this post since I didn't find the usual one by u/llogiq)
## [12][lifetime mismatch problem](https://www.reddit.com/r/rust/comments/ibys3a/lifetime_mismatch_problem/)
- url: https://www.reddit.com/r/rust/comments/ibys3a/lifetime_mismatch_problem/
---
    fn foo(s: &amp;mut &amp;mut [u8]) {
        let (_, rest) = s.split_at_mut(1);
        *s = rest;
    }

hi,

I find that the above code fails to compile with the following error:

     error[E0623]: lifetime mismatch
     --&gt; src/lib.rs:3:10
      |
    1 | fn foo(s: &amp;mut &amp;mut [u8]) {
      |           --------------
      |           |
      |           these two types are declared with different lifetimes...
    2 |     let (_, rest) = s.split_at_mut(1);
    3 |     *s = rest;
      |          ^^^^ ...but data from `s` flows into `s` here

I dont understand what is going on here. How to fix it?

Further, the following compiles:

    fn foo(s: &amp;mut &amp;[u8]) {
        let (_, rest) = s.split_at(1);
        *s = rest;
    }

I wonder why, as the mutablity should not affect lifetimes.
