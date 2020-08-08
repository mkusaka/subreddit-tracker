# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (31/2020)!](https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 350](https://www.reddit.com/r/rust/comments/i4a3e0/this_week_in_rust_350/)
- url: https://this-week-in-rust.org/blog/2020/08/04/this-week-in-rust-350/
---

## [3][A Terminal User Interface for Taskwarrior written in Rust](https://www.reddit.com/r/rust/comments/i5lx46/a_terminal_user_interface_for_taskwarrior_written/)
- url: https://i.redd.it/jibjh3kq8nf51.gif
---

## [4][Small WIP Browser based on Matt Brubeck - Let's build a browser engine!](https://www.reddit.com/r/rust/comments/i5y4h0/small_wip_browser_based_on_matt_brubeck_lets/)
- url: https://git.sr.ht/~bo41/prowser
---

## [5][Knurling-rs: Get a grip on bare-metal Rust](https://www.reddit.com/r/rust/comments/i5d7mf/knurlingrs_get_a_grip_on_baremetal_rust/)
- url: https://ferrous-systems.com/blog/knurling-rs/
---

## [6][smol vs tokio vs async-std;](https://www.reddit.com/r/rust/comments/i5hppj/smol_vs_tokio_vs_asyncstd/)
- url: https://www.reddit.com/r/rust/comments/i5hppj/smol_vs_tokio_vs_asyncstd/
---
Hello!

I'm trying to understand the motivation behind `smol` (and related crates) a little better, as compared with `tokio` and `async-std`. More generally, I want to make sure that have a good enough understanding of the current world of async!

Here's my current understanding in the form of numbered points (to hopefully make them easier to reply to!):

1. Futures need to be polled to completion. This is the job of an _executor_. Some futures additionally need to wait for events from the kernel to know when there might be data ready to read from a file, or somesuch. A _reactor_ handles this (by using `mio`, or `polling` for instance to register for events from the kernel and know when things might be able to progress).

2. `tokio` has an executor and reactor bundled within it. Futures that rely on the `tokio::io/fs` need to be run inside the context of a tokio runtime (which makes the tokio reactor available to them and allows spawning), and so you must remember to start one up before using tokio related bits. These futures can be run on any executor, though, I think.

3. `async-std` and `smol` both use the same underlying executor and reactor code now.

4. `smol` is really just a light wrapper around `async-executor`, and doesn't come with a reactor itself. Crates like `async-io` (which `async-net` builds on) start up a reactor on-demand when it's needed by certain futures (for async io and timers). Futures that rely on these underlying crates like `async-net` for instance, don't care about the executor that runs them or about any reactor existing or being in scope (it'll start as needed).

5. Spawning futures: `tokio`, `async-std` and `smol` all start up an executor (or multiple of them), and if you try to spawn a future, you'll need to spawn it into one of these executors (ie, there is no generic way to spawn a future onto "whatever is available").

6. `smol` and `async-std` can be asked to start up a tokio runtime so that tokio related futures will run and can be spawned without issue. Tokio bits will then run inside a separate tokio runtime that lives alongside the bits `smol` spins up.

7. If I want to write a library that's generic over whether it's run by `tokio`, `async-std` etc, and don't want to use feature flags to conditionally code for each one, then I need to: 
  a. avoid spawning futures in my library (which then ties me to a given executor)
  b. either make users kick off a `tokio` runtime, or base the library on something like `async-io`/`async-net` which will spin up a runtime behind the scenes as necessary, or write my own runtime and spin that up as needed.

8. If I want to write application code that doesn't care whether the future it runs relies on `tokio` or `async-std` features, using `smol` or `async-std` at the top level are probably the easiest way to do this; either will spin up a `tokio` runtime as needed, and`smol`+`async-std` are compatible with each other and rely on the same fundamentals now.

9. `smol` takes a slightly different direction than `tokio` by splitting up the async primitives that you may need (eg executor and reactor) into separate crates and expecting that users should pick and mix between these different crates as needed. The observable impact of this for me is that futures written in this way don't depend on (for instance) a global reactor, or a global thread-pool for blocking operations, and instead will spin them up as needed (rather than the `tokio` approach of expecting these things to exist when the future runs). I feel like there's something fundamental I might be missing here though?

10. When `smol` makes the claim that "All async libraries work with smol out of the box." in its README, it is specifically referring to `tokio` and `async-std` based libraries. Is there a more fundamental claim though that's being made here though? I can see that `smol` encourages futures to pull in and spin up things like reactors as needed, which in turn makes them more portable, but is there more to it?

I'm hoping that I've generally got the gist here; I guess I have a few questions over `smol` and its philosophy, and am interested to know if it is doing something fundamnetally different which could help bridge the gap between different async ecosystems (eg tokio and async-std). I'm also interested in making sure that I use the right building blocks if I create my own async libraries.

Thanks for reading; I'm looking forward to being corrected :)
## [7][What Is The Minimal Set Of Optimizations Needed For Zero-Cost Abstraction?](https://www.reddit.com/r/rust/comments/i5ila5/what_is_the_minimal_set_of_optimizations_needed/)
- url: https://robert.ocallahan.org/2020/08/what-is-minimal-set-of-optimizations.html
---

## [8][Can't find crate when compiling](https://www.reddit.com/r/rust/comments/i5syo2/cant_find_crate_when_compiling/)
- url: https://www.reddit.com/r/rust/comments/i5syo2/cant_find_crate_when_compiling/
---
Sorry in advance I'm very new to programming in general. when I run my code the crate works fine, but when trying to compile it using the *rustc* method below. It says it can't find the crate. can someone please explain this to me? Thank you

https://preview.redd.it/ap4dqj76hpf51.png?width=668&amp;format=png&amp;auto=webp&amp;s=3181cdf19565ff80203a29febfc323104e79b2ed
## [9][Code review of interprocess-channel implementation](https://www.reddit.com/r/rust/comments/i5z3gc/code_review_of_interprocesschannel_implementation/)
- url: https://www.reddit.com/r/rust/comments/i5z3gc/code_review_of_interprocesschannel_implementation/
---
Heyho :D I'm currently working through "Inside the message passing interface" and try implementing stuff from it in rust (rather than C which is used in the book), but in a more modern way.

I started by writing up the very basic communication mechanism they use for multiprocess communication (mmap based) and their benchmarks and I found it hugely encouraging that this could work out when I saw that the rust version outperformed the C version from the book by a factor of 5.

But I always hated MPI for it's type system and "yeah you allocate a buffer and then give it to me" system; so I tried packing this up in a channel-interface. This actually worked out quite well performancewise and in the very basic benchmarks I've done it actually comes out about an order of magnitude faster than the servo ipc-channels for equally sized messages, and if I ramp the sizes of the data being transfered up (the servo channels are limited to 32 element arrays because they use serde) it actually comes out about 1000 times faster (in terms of data rate, not latency) compared to the servo channels. The servo channels of course have other **big** advantages so I'm not claiming to be better than them in anything but raw throughput performance.

My problem is: I don't like the code at all. It contains tons of unsafe bits (and as such probably also some nasty bugs that I just haven't found until now) and has to deviate from the "normal" channel interface to please the borrow checker (I think writing the normal `channel` function for my `Sender` and `Receiver` structs is literally impossible), it has ugly places like using a fixed size buffer on the receiving end rather than a "fitting" one because of current limitations in rust (it'd have to const evaluate a generic function for an array size which according to the associated error "requires changes to the language design" or something). This is the first time I've really used lower level primitives and unsafe things in rust which probably also contributed to the code's ugliness.

&amp;#x200B;

So all in all: I'd really appreciate it if someone could look at the code and give me some pointers for improvements :) I've thought about using async for some parts (because the communication is blocking and the sender/receiver has to poll a piece of memory, I kinda feel like this would be a nice place for async); what do you guys think about this? Would this be terribly detremental to the performance?

&amp;#x200B;

If anyone wants to do have a look at the code, here's the basis structure:

At the lowest level there's the `TransferBuffer`; this is just a mmap-based raw byte buffer with an "owner" mechanism. This mechanism reserves a bit of memory for synchronisation of buffer reads and writes without using locks.

One level up there's the `Sender&lt;'a, T&gt;` and `Receiver&lt;T&gt;` structs. Every `Receiver` owns a `TransferBuffer` and the accompanying `Sender` holds a mutable reference to it. This is kinda going in an unsafe direction, but since they'll end up in different processes and synchronize themselves it's fine I think. The synchronization works such that the Sender "hands-off" it's access-rights to the receiver via the transferbuffers "owner" mechanism when it's done sending and vice versa. The "serialization" (if you want to call it that) of the values to be sent is just a literal copy of the bytes. Since this serialized representation is not going to be moved to another machine or something I think this is fine? But this is a part I'm quite unsure about.

&amp;#x200B;

EDIT: yeah I of course forgot to link to the code: [https://github.com/SV-97/MPI-rs](https://github.com/SV-97/MPI-rs)
## [10][How do you select a random enum value without manually matching](https://www.reddit.com/r/rust/comments/i5n2zm/how_do_you_select_a_random_enum_value_without/)
- url: https://www.reddit.com/r/rust/comments/i5n2zm/how_do_you_select_a_random_enum_value_without/
---
Title. I have a enum in rust who's values are not sequential (A = 1, B = 2, C = 5, D = 100, E = 300)


And I just want to pick one of those at random. I don't want to manually have to pattern match every value because I feel that defeats the purpose of having the enum.
## [11][Code Review needed](https://www.reddit.com/r/rust/comments/i5x4sy/code_review_needed/)
- url: https://www.reddit.com/r/rust/comments/i5x4sy/code_review_needed/
---
Hey guys just whipped up a Sudoku solver using rust

Used backtracking for solving a Sudoku

Code : [https://github.com/Dravigon/rust\_sudoku\_solver](https://github.com/Dravigon/rust_sudoku_solver)

Have coded rust before but actually want someone to look at the code and give me code reviews .

Thanks beforehand
## [12][Question about rust-analyzer](https://www.reddit.com/r/rust/comments/i5lkc2/question_about_rustanalyzer/)
- url: https://www.reddit.com/r/rust/comments/i5lkc2/question_about_rustanalyzer/
---
I was using RLS until now and decided to check rust-analyzer.

Is it not possible to check for errors as you type like what RLS does. I want to keep rust-analyzer if I can get this working since it has better auto-completion.

I'm using vim with coc-rust-analyzer. Also checked this in VSCode and doesn't work there either.
