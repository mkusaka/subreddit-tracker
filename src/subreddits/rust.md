# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (27/2020)!](https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 345](https://www.reddit.com/r/rust/comments/hisn3e/this_week_in_rust_345/)
- url: https://this-week-in-rust.org/blog/2020/06/30/this-week-in-rust-345/
---

## [3][Artichoke Ruby is 2.2% faster after reimplementing the symbol table in Rust](https://www.reddit.com/r/rust/comments/hk8nwr/artichoke_ruby_is_22_faster_after_reimplementing/)
- url: https://www.reddit.com/r/rust/comments/hk8nwr/artichoke_ruby_is_22_faster_after_reimplementing/
---
[Artichoke Ruby](https://github.com/artichoke/artichoke) is an alternative Ruby implementation written in Rust. I recently reimplemented the VM's symbol table with [intaglio](https://crates.io/crates/intaglio) to replace the C implementation in the mruby base.

[The resulting Ruby is 2.2% faster when running a portion of the ruby/spec suite](https://github.com/artichoke/artichoke/pull/730).

The Ruby symbol table is a bytestring interner that is used mostly during interpreter boot and source loading.
## [4][React core team discussing re-writing internals in Rust](https://www.reddit.com/r/rust/comments/hkaxgl/react_core_team_discussing_rewriting_internals_in/)
- url: https://www.reddit.com/r/rust/comments/hkaxgl/react_core_team_discussing_rewriting_internals_in/
---
At 39:05 — https://www.twitch.tv/videos/667925289?t=00h39m05s
## [5][Ringbahn II: the central state machine](https://www.reddit.com/r/rust/comments/hk8lab/ringbahn_ii_the_central_state_machine/)
- url: https://without.boats/blog/ringbahn-ii/
---

## [6][Boa v0.9: measureme &amp; optimisations](https://www.reddit.com/r/rust/comments/hkifvb/boa_v09_measureme_optimisations/)
- url: https://boa-dev.github.io/2020/07/03/boa-release-09.html
---

## [7][Compiler team roadmap 2020-2021](https://www.reddit.com/r/rust/comments/hjx82b/compiler_team_roadmap_20202021/)
- url: https://rust-lang.github.io/compiler-team/minutes/design-meeting/2020-05-29-Roadmap-2020-2021/
---

## [8][Finished implementing OpenDataStructures in Rust](https://www.reddit.com/r/rust/comments/hk2nbp/finished_implementing_opendatastructures_in_rust/)
- url: https://www.reddit.com/r/rust/comments/hk2nbp/finished_implementing_opendatastructures_in_rust/
---
I have completed a Rust implementation of OpenDataStructures.

I hope this project will be helpful to anyone trying to learn data structures with Rust. Please use it as an unofficial reference implementation. 
And all contributions are welcome! : [Open Data Structures (in Rust)](https://github.com/o8vm/ods)
## [9][Eye: cross platform camera capture &amp; control](https://www.reddit.com/r/rust/comments/hk7v95/eye_cross_platform_camera_capture_control/)
- url: https://www.reddit.com/r/rust/comments/hk7v95/eye_cross_platform_camera_capture_control/
---
Hi r/rust, I'm working on implementing computer vision libraries and primitives in native Rust. One essential building block is image capture (mostly using USB cameras such as the Intel RealSense D435 I'm using). Naturally, I looked around but could not find much in the way of good, native Rust camera capture libraries.

rscam or webcam\_capture are unfit for the cause: their API is either too limited or the pipeline is not efficient enough for realtime conditions (e.g. by copying buffers around needlessly).

Thus, I decided to start a new crate for cross platform camera capture and control: Eye ([https://github.com/raymanfx/eye-rs](https://github.com/raymanfx/eye-rs)). It features platform abstractions to create capture devices and streams to capture images from those devices. So far, I implemented a Linux backend using my own v4l-rs crate. Zero-copy image streaming is implemented by mapping buffers into userspace and directly returning them to the caller through a layer of abstractions.

Soon, I will leverage another crate of mine ([https://github.com/raymanfx/ffimage](https://github.com/raymanfx/ffimage)) for automatic buffer conversion so one can request e.g. a RGB stream from devices which normally only support YUYV or other formats common for USB webcams such as the Logitech C922 Pro. ffimage features parallel image conversion by using Rayon and is quite fast, on my machine it's actually more than suitable for realtime buffer conversion.

If you are interested in running the code right now (on Linux), the `next` branch features a GTK GUI example: [https://github.com/raymanfx/eye-rs/blob/next/examples/gtk.rs](https://github.com/raymanfx/eye-rs/blob/next/examples/gtk.rs).

You can find version 0.1.0 on [crates.io](https://crates.io): [https://crates.io/crates/eye](https://crates.io/crates/eye).

I'm hoping to find people interested in collaborating and adding Windows and macOS backends. My next steps include automatic format conversion and control parameter support (autofocus, white balance etc).
## [10][Ownership of the standard library implementation | Inside Rust Blog](https://www.reddit.com/r/rust/comments/hjuh2f/ownership_of_the_standard_library_implementation/)
- url: https://blog.rust-lang.org/inside-rust/2020/07/02/Ownership-Std-Implementation.html
---

## [11][A boilerplate template for starting web services using Warp + SQLx (PostgreSQL) + Redis + Juniper (GraphQL)](https://www.reddit.com/r/rust/comments/hk34gw/a_boilerplate_template_for_starting_web_services/)
- url: https://www.reddit.com/r/rust/comments/hk34gw/a_boilerplate_template_for_starting_web_services/
---
[https://github.com/rusty-crab/warp-api-starter-template](https://github.com/rusty-crab/warp-api-starter-template) 

The project is inspired by [Meh's blog post](https://meh.schizofreni.co/2020-04-18/comfy-web-services-in-rust). An attempt to put together what is described in the blog post code excerpts and missing pieces. This is far from finished to be called an actual template for someone to get started. But it has all the pieces from the blog post. I will continue to make it more modular and usable. Meanwhile, contributions, suggestions, directions are welcome (PRs, Issues, etc..). 

Thank you.
## [12][Nick Cameron: Early Impressions of Go from a Rust Programmer](https://www.reddit.com/r/rust/comments/hkjs96/nick_cameron_early_impressions_of_go_from_a_rust/)
- url: https://pingcap.com/blog/early-impressions-of-go-from-a-rust-programmer
---

