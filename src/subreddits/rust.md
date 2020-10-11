# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (41/2020)!](https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 359](https://www.reddit.com/r/rust/comments/j71tq7/this_week_in_rust_359/)
- url: https://this-week-in-rust.org/blog/2020/10/07/this-week-in-rust-359/
---

## [3][Rust FFI: Microsoft Flight Simulator SDK Part 1 [video]](https://www.reddit.com/r/rust/comments/j92jc0/rust_ffi_microsoft_flight_simulator_sdk_part_1/)
- url: https://youtu.be/jNNz4h3iIlw
---

## [4][I made a Gameboy/Gameboy Color Emulator which runs in the browser](https://www.reddit.com/r/rust/comments/j8vw84/i_made_a_gameboygameboy_color_emulator_which_runs/)
- url: https://www.reddit.com/r/rust/comments/j8vw84/i_made_a_gameboygameboy_color_emulator_which_runs/
---
Hello everyone,

I made this a few months back but now finally decided to wrap it up. Mainly during the lockdown I was home a lot more and wanted to do something in Rust. I love using Rust + WASM, so that resulted in a CHIP-8 emulator. Right after that I began work on a gameboy/color emulator.

Initially I just wanted to be able to emulate Pokemon Gold/Silver/Crystal, the first games I ever played on a gameboy emulator when I was a kid. Eventually however it devolved into madness about passing tests and improving accuracy, spending weeks fixing single bugs.

Anyways, I wrote the emu just for myself as a learning experience. So it's lacking in a lot of features you'd expect from one made for others to use, and it's also not mobile friendly unfortunately (but I'll work on that soon). Still, it may be of some interest to someone here :D

[https://github.com/BlueBlazin/gbemu](https://github.com/BlueBlazin/gbemu)

Thanks for reading and feel free to suggest any improvements here!
## [5][Optional parameters in Rust](https://www.reddit.com/r/rust/comments/j8p6fx/optional_parameters_in_rust/)
- url: https://vidify.org/blog/rust-parameters/
---

## [6][New to Rust, and decided to make a game](https://www.reddit.com/r/rust/comments/j8wqfi/new_to_rust_and_decided_to_make_a_game/)
- url: https://www.reddit.com/r/rust/comments/j8wqfi/new_to_rust_and_decided_to_make_a_game/
---
Hey everyone!

I just recently started to learn Rust, and as an exercise I decided to make a small game in theme with October.

Camp Misty is an asymmetric multiplayer game of deduction where a helpless victim is being hunted down by a ruthless killer!

The game is played on the command line, and works on local networks as well as over the internet. Instructions for how to play are included with the game.

[You can find it on GitHub, here.](https://github.com/ReeCocho/camp-misty).

Since I’m fairly new to Rust, I would really appreciate it if you could take a look and provide some pointers on my usage of the language. Also, if you find any bugs (which I am sure there are plenty of) please notify me so I can fix them.

Thanks!
## [7][Introduction of Lazy API in Polars DataFrame library](https://www.reddit.com/r/rust/comments/j93ize/introduction_of_lazy_api_in_polars_dataframe/)
- url: https://www.reddit.com/r/rust/comments/j93ize/introduction_of_lazy_api_in_polars_dataframe/
---
I just released `v0.7.x` for Polars. Polars is a DataFrame library fully implemented in Rust. This release introduces a [lazy AP](https://ritchie46.github.io/polars/polars/lazy/frame/struct.LazyFrame.html)[i](https://ritchie46.github.io/polars/polars/lazy/frame/struct.LazyFrame.html). Polars is completely in-memory so the `eager` API is very much like Pandas (only faster). The lazy API is somewhat in between spark and pandas. There is no distributed computation, but the queries will be optimized to do less work.

* [github](https://github.com/ritchie46/polars)
* [changelog](https://docs.rs/polars/0.7.1/polars/doc/changelog/v0_7/index.html)
* [introductory notebook in Python](https://github.com/ritchie46/polars/blob/master/examples/10_minutes_to_pypolars.ipynb)

I should note that a lot of the lazy API exists due to the excellent work done on [Apache DataFusion](https://arrow.apache.org/blog/2019/02/04/datafusion-donation/) from which I learned a lot.
## [8][📽 Atomatic Rust code coverage in VSCode [follow up]](https://www.reddit.com/r/rust/comments/j91z8r/atomatic_rust_code_coverage_in_vscode_follow_up/)
- url: https://youtu.be/JEgrkmEAHjA
---

## [9][gdbstub 0.4: An ergonomic, #![no_std] implementation of the GDB Remote Serial Protocol in Rust](https://www.reddit.com/r/rust/comments/j8oaxw/gdbstub_04_an_ergonomic_no_std_implementation_of/)
- url: https://www.reddit.com/r/rust/comments/j8oaxw/gdbstub_04_an_ergonomic_no_std_implementation_of/
---
[crates.io](https://crates.io/crates/gdbstub) | [docs](https://docs.rs/gdbstub/) | [repo](https://github.com/daniel5151/gdbstub)

An ergonomic and easy-to-integrate implementation of the [GDB Remote Serial Protocol](https://sourceware.org/gdb/onlinedocs/gdb/Remote-Protocol.html#Remote-Protocol) in Rust, with full `#![no_std]` support. `gdbstub` makes extensive use of Rust's powerful type system + generics to enforce protocol invariants at compile time, minimizing the number of tricky protocol details end users have to worry about.

A lot has changed since my [last post](https://www.reddit.com/r/rust/comments/hpys1x/gdbstub_02_an_ergonomic_allocationfree/) announcing `gdbstub 0.2`!

Version `0.4` includes a _major_ API overhaul, tons of internal optimizations, and a slew of new GDB protocol features, making it the fastest, leanest, and most featureful release of `gdbstub` yet!

It's been absolutely incredible having so many people contribute to the library, and seeing `gdbstub` being used in all sorts of cool projects. Thank you for all the support!

- Lots of new [built-in](https://docs.rs/gdbstub/0.4.0/gdbstub/arch/index.html) architecture definitions:
    -   RISC-V (only integer ISA at the moment) [\#21](https://github.com/daniel5151/gdbstub/pull/21) ([thomashk0](https://github.com/thomashk0))
    -   x86 (i386) [\#23](https://github.com/daniel5151/gdbstub/pull/23) ([jamcleod](https://github.com/jamcleod))
    -   x86_86 [\#11](https://github.com/daniel5151/gdbstub/pull/11) ([jamcleod](https://github.com/jamcleod))
    -   32-bit PowerPC [\#25](https://github.com/daniel5151/gdbstub/pull/25) ([jamcleod](https://github.com/jamcleod))
    -   Mips and Mips64 [\#13](https://github.com/daniel5151/gdbstub/pull/13) ([starfleetcadet75](https://github.com/starfleetcadet75))

- A bunch of newly implemented GDB Protocol Extensions:
    -   [`ExtendedMode`](https://docs.rs/gdbstub/0.4.0/gdbstub/target/ext/extended_mode/index.html) - Allow targets to spawn new processes / attach to existing processes / restart execution.
        -   Includes support for `set disable-randomization`, `set environment`, `set startup-with-shell`, and `set cwd`. 
    -   [`SectionOffsets`](https://docs.rs/gdbstub/0.4.0/gdbstub/target/ext/section_offsets/index.html) - Get section/segment relocation offsets from the target. [\#30](https://github.com/daniel5151/gdbstub/pull/30) ([mchesser](https://github.com/mchesser))
    -   `Target::base` - Add single register read/write methods + `RegId` identifiers [\#22](https://github.com/daniel5151/gdbstub/pull/22) ([thomashk0](https://github.com/thomashk0))
    -   `Target::base` - Add `StopReason::Signal(u8)` variant [\#19](https://github.com/daniel5151/gdbstub/pull/19) ([mchesser](https://github.com/mchesser))

- A Major API Overhaul:
    -   The `Target` API has been rewritten to use ["Inlineable Dyn Extension Traits" (IDETs)](https://docs.rs/gdbstub/0.4.0/gdbstub/target/ext/index.html#how-protocol-extensions-work---inlineable-dyn-extension-traits-idets)
        -   _By breaking up `Target` into smaller pieces which can be mixed-and-matched, not only does has it made it easier than ever to get up-and-running with `gdbstub`, but it's also unlocked a ton of awesome internal optimizations + ergonomic improvements:_
            -   Eliminates the binary-size impact of unimplemented GDB protocol features by _guaranteeing_ dead-code-elimination of all unused parsing/handling code.
            -   Adds compile-time enforcement that certain groups of methods are implemented in-tandem, removing the need for runtime checks (e.g: `add_sw_breakpoint` and `remove_sw_breakpoint`).
            -   Introduces a robust framework for adding new protocol extensions in a backwards-compatible manner.
    -   Update the `Target` API with support for _non-fatal_ error handling.
    -   Update the `Connection` trait with new (mostly optional) methods.
    -   Error types are now marked as `#[non_exhaustive]`.


By the way, if you're taking part in [Hacktoberfest](https://hacktoberfest.digitalocean.com/) this year, there are plenty of ways to contribute to `gdbstub`. There's a whole laundry list of protocol extensions and new architectures to support, so check out the [issue tracker](https://github.com/daniel5151/gdbstub/issues) and consider lending a hand!

Cheers!
## [10][Ocean - A C/C++ project manager/build system written in Rust](https://www.reddit.com/r/rust/comments/j8rdbo/ocean_a_cc_project_managerbuild_system_written_in/)
- url: /r/Cplusplus/comments/j8rd3y/ocean_a_cc_project_managerbuild_system_written_in/
---

## [11][t-rec, a blazingly fast terminal recorder that generates animated gif images for the web written in rust for MacOS](https://www.reddit.com/r/rust/comments/j8tqs9/trec_a_blazingly_fast_terminal_recorder_that/)
- url: https://www.reddit.com/r/rust/comments/j8tqs9/trec_a_blazingly_fast_terminal_recorder_that/
---
&amp;#x200B;

[demo of t-rec](https://i.redd.it/q60pnehlfcs51.gif)

[https://github.com/sassman/t-rec-rs](https://github.com/sassman/t-rec-rs)

[https://crates.io/crates/t-rec](https://crates.io/crates/t-rec)

This is the first release, therefore testing and feedback is very much appreciated. 

Contributions for Linux and Windows support are welcome as well :)
## [12][Are there any good video blogs on rust?](https://www.reddit.com/r/rust/comments/j8y5zx/are_there_any_good_video_blogs_on_rust/)
- url: https://www.reddit.com/r/rust/comments/j8y5zx/are_there_any_good_video_blogs_on_rust/
---
Coming from C++, and looking for something similar like C++ Weekly [https://www.youtube.com/playlist?list=PLs3KjaCtOwSZ2tbuV1hx8Xz-rFZTan2J1](https://www.youtube.com/playlist?list=PLs3KjaCtOwSZ2tbuV1hx8Xz-rFZTan2J1)
