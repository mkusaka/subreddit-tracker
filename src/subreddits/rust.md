# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (11/2020)!](https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (11/2020)?](https://www.reddit.com/r/rust/comments/fg0q9l/whats_everyone_working_on_this_week_112020/)
- url: https://www.reddit.com/r/rust/comments/fg0q9l/whats_everyone_working_on_this_week_112020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-11-2020/39269?u=llogiq)!
## [3][Bottlerocket, a new Linux-based OS, uses Rust for almost all new components](https://www.reddit.com/r/rust/comments/fgm2st/bottlerocket_a_new_linuxbased_os_uses_rust_for/)
- url: https://github.com/bottlerocket-os/bottlerocket
---

## [4][AVRA-RS: Microchip (Atmel) assembler written entirely in Rust](https://www.reddit.com/r/rust/comments/fguvdu/avrars_microchip_atmel_assembler_written_entirely/)
- url: https://github.com/no111u3/avra-rs
---

## [5][Impressions of Rust as a Swift Developer: Memory Management](https://www.reddit.com/r/rust/comments/fgue77/impressions_of_rust_as_a_swift_developer_memory/)
- url: https://blog.spencerkohan.com/impressions-of-rust-as-a-swift-developer-2/
---

## [6][My exploration of Rust and .NET](https://www.reddit.com/r/rust/comments/fgqz5l/my_exploration_of_rust_and_net/)
- url: https://ericsink.com/entries/dotnet_rust.html
---

## [7][[talk] Rust NYC: Jon Gjengset - Demystifying unsafe code](https://www.reddit.com/r/rust/comments/fglm6k/talk_rust_nyc_jon_gjengset_demystifying_unsafe/)
- url: https://www.youtube.com/watch?v=QAz-maaH0KM
---

## [8][Etsy Open-Sources Their Neuro-Evolution Machine Learning System for Search](https://www.reddit.com/r/rust/comments/fgmlwc/etsy_opensources_their_neuroevolution_machine/)
- url: https://github.com/etsy/Evokit
---

## [9][Trying to understand &amp;str vs. str, &amp;[T] vs [T], &amp;OsStr vs OsStr](https://www.reddit.com/r/rust/comments/fgpdb0/trying_to_understand_str_vs_str_t_vs_t_osstr_vs/)
- url: https://www.reddit.com/r/rust/comments/fgpdb0/trying_to_understand_str_vs_str_t_vs_t_osstr_vs/
---
A someone coming from C++-land and working with Rust recently, I'm trying to get a better understanding of how Rust's slices work. With C++, whenever I see a `std::string_view` or `std::span`, I know that there's a struct/class containing a pointer to the beginning of some contiguous sequence and a size. I understand that Rust's `&amp;str` and `&amp;[T]` primitives serve the same purpose.

What (I think) I know: `&amp;str` and `&amp;[T]` (and `&amp;OsStr`, etc.) behave less like "pointers" or "references", despite the `&amp;`, and more like structures holding a reference to the data and the size. The fact that `mem::size_of::&lt;&amp;str&gt;()` is 16 bytes on my system and that the following code:

    let foo = "Hello, world!";

compiles down to:

    example::test:
            push    rax
            lea     rax, [rip + .L__unnamed_1]
            mov     rdi, rax
            mov     esi, 13
            ...

when optimizations are disabled seems to confirm this.

What I don't understand is what exactly `str` and `[T]` are. The documentation seems to just gloss over them. What do they represent and how are they laid out in memory? Why aren't `str` and `[T]` themselves the structure that holds the pointer and size?

I'm unable to call `std::mem::size_of::&lt;T&gt;()` on `str` or `[T]` or create variables of those types without hitting compile errors. I know that it's possible to use these types though. Looking at the [`OsStr` implementation](https://github.com/rust-lang/rust/blob/15812785344d913d779d9738fe3cca8de56f71d5/src/libstd/ffi/os_str.rs#L106-L108), it seems to boil down to [a struct holding a `[u8]`](https://github.com/rust-lang/rust/blob/15812785344d913d779d9738fe3cca8de56f71d5/src/libstd/sys_common/os_str_bytes.rs#L27-L29). What does it mean for a struct to contain a `[u8]`?

When I see `&amp;str`, my brain is trying to equate that to `std::string_view *` in C++. I know that's not correct and that it actually behaves like a `std::string_view`. Maybe it'd be more intuitive to me if I knew why the `&amp;` is there and what purpose it serves.

I'd appreciate any pointers (ha!) for better understanding these types!
## [10][release of crow 0.4.0, a pixel perfect 2D graphics library](https://www.reddit.com/r/rust/comments/fgeoey/release_of_crow_040_a_pixel_perfect_2d_graphics/)
- url: https://www.reddit.com/r/rust/comments/fgeoey/release_of_crow_040_a_pixel_perfect_2d_graphics/
---
After publishing version [0.3.0](https://www.reddit.com/r/rust/comments/f2wuwc/i_just_published_crow_03_a_pixel_perfect_2d/), this version of `crow` now supports OpenGL version ~~0.3.3~~ `3.3` and has actually decent errors.

Using it in my WIP showcase game, [akari](https://github.com/lcnr/akari), I am really happy with the current API and will keep working on it in the future.

[CHANGELOG](https://github.com/lcnr/crow/blob/master/CHANGELOG.md)
## [11][Rust impl (plus Python bindings) of MIDAS Count-Min Sketch based graph anomaly detector](https://www.reddit.com/r/rust/comments/fgshi9/rust_impl_plus_python_bindings_of_midas_countmin/)
- url: https://www.reddit.com/r/rust/comments/fgshi9/rust_impl_plus_python_bindings_of_midas_countmin/
---
[https://github.com/scooter-dangle/midas\_rs](https://github.com/scooter-dangle/midas_rs)

Original implementation: [https://github.com/bhatiasiddharth/MIDAS](https://github.com/bhatiasiddharth/MIDAS)

Corresponding paper: [https://www.comp.nus.edu.sg/\~sbhatia/assets/pdf/midas.pdf](https://www.comp.nus.edu.sg/~sbhatia/assets/pdf/midas.pdf)

&amp;#x200B;

A teammate wanted to noodle around with the algorithm described in the paper. The original C++ implementation is already pretty straight-forward (which made the port very easy), but both to be able to understand the workings of the algorithm a little better and to have a version that we'd be more comfortable supporting (since we work primarily in Rust and Python) if anyone on the team wanted to play with it, I ported it to Rust.

I haven't done any benchmarking or extensive tests—will develop further based on interest from the team or external users.

I think the original author will be around here in case anyone has questions about the algorithm itself and how to interpret the outputs.
## [12][The 2020 RustConf CFP is Now Open!](https://www.reddit.com/r/rust/comments/fgj1u1/the_2020_rustconf_cfp_is_now_open/)
- url: https://blog.rust-lang.org/2020/03/10/rustconf-cfp.html
---

