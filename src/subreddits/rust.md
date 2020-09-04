# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (36/2020)!](https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (36/2020)?](https://www.reddit.com/r/rust/comments/ijvxwz/whats_everyone_working_on_this_week_362020/)
- url: https://www.reddit.com/r/rust/comments/ijvxwz/whats_everyone_working_on_this_week_362020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-36-2020/48100?u=llogiq)!
## [3][Leaving the Rust core team](https://www.reddit.com/r/rust/comments/im98yg/leaving_the_rust_core_team/)
- url: http://www.ncameron.org/blog/leaving-the-rust-core-team/
---

## [4][My Rust 2021 roadmap](https://www.reddit.com/r/rust/comments/imd8b8/my_rust_2021_roadmap/)
- url: https://www.reddit.com/r/rust/comments/imd8b8/my_rust_2021_roadmap/
---
As a full-time HPC developer, I want Rust to have better access to the hardware and be better able to interface with optimized system libraries. I want it to be safer and have better tested unsafe code, since I write a lot of it, and also use a lot of crates from crates.io that isn't feasible for me to manually inspect (usually inspect them once initially, but keeping up with updates is hard). I also want some of its APIs to be more ergonomic, and would prefer not to have to use nightly for everything.

- SIMD: C++ compilers like clang and GCC offer quite good language extensions for portable SIMD. The C++ Parallelism TS introduces a portable `std::simd&lt;T&gt;`. These compilers also allow users to pick up the math library and SIMD math libraries to use. Work here seems stalled, and there is no work to support ARM SVE, which we currently have available in some hardware.

- GPU: Rust allows me to almost use 100% of my CPU. In all computers I work on, this represent 1-5% of the actual memory bandwidth and compute available on hardware, 95-99% of it is provided by the GPU. Accessing the GPU from Rust is extremely unsafe and unergonomic, to the point that's it not really worth doing and I need to use C or C++ instead. C++ has SyCL, AMP, and CUDA, which are all efficient and ergonomic. Using Rust's PTX backend is extremely unsafe, and many basic PTX features are not available. Using Rust's to target OpenCL or SPIR-V is very unergonomic as well, essentially relying on DSLs to convert Rust code to some other language, that then gets compiled (there is no native support for a SPIR-V backend in rustc). This essentially means that I can't use Rust for any large application that's compute or memory bound, limiting its usage at best to scaffolding or tooling around the app itself. 

- FFI: it is hard to use Rust to interface with optimized FFI libraries: lack of 128-bit integer, long double, and SIMD type support. Lack of C attributes support, full unwinding support, anonymous unions / structs, lack of dynamically sized types, etc. Its 2020, and after multiple years `std::ffi::CStr` still isn't safe to use in FFI where C strings are expected.. Lack of extern types (although `ffi::c_void` covers most uses). Etc. I would like to incrementally start using Rust in some larger C++ CPU-only projects, and I can't because FFI is not good enough.

- soundness: The only reason I use Rust where I can is because safe Rust _might_ be sound. If it turns out it isn't, there are other languages I'm already using as well that I would use more instead. Safe Rust soundness relies on unsafe code being correct. The documentation that explains what unsafe Rust code is or isn't correct is very poor. It is split into 3 or more inconsistent documents (the nomicon web version has not been updated in years and one needs to read its source code instead!). This means that many unsafe code writers are uninformed about what's allowed or not, and even though they try their bests, it's pretty hard to actually be sure that some unsafe code is correct in practice: too often the answer is "nobody knows" which isn't very reassuring. The UCG RFC was never opened and merged, and the reference is quite incomplete (there is no operational semantics of Rust, no memory model, many docs in the UCGs are PoCs). Testing unsafe code properly is hard: miri is too slow and too incomplete to be practical for real pure-Rust applications. Identifying untested unsafe code is also hard: using code-coverage instrumentation is too cumbersome from Rust (running the tests of a crate should just tell me the code coverage, and warn me if some unsafe code is not covered! this should not require any external tools and be part of rustc and cargo proper!). I wish crates.io had better integration with unsafe code, to be able to tell me whether a crate contains any unsafe code at all. I wish crates.io would have better integration with CVEs, to warn me if a crate version has a CVE open. 

- cargo: I need nightly on many of my crates due to this super old bug: https://github.com/rust-lang/cargo/issues/7916 Nightly is better than the workarounds we used to have, but still. All the problems cargo has with resolving dependencies and features give the feeling that cargo's model isn't actually very sound and more like a ball of ducktape.

- embedded: I find it hard to avoid using floating point on hardware that does not support it. I would prefer compilation to fail instead of using soft-floats. I would prefer to not bundle all the floating point math in those apps as part of libcore. I wish libstd was more modular in general, just like liballoc, so that I can cherry pick the modules of std that I can implement for my target, instead of having to implement all of standard, with 99% of the functionality being `unimplemented!`.

- nightly: almost all the Rust code I write uses nightly due to lack of stable specialization and function traits. Some of this code is over 3 years old, and it still works correctly without any updates. Some of the other features I use are const-generics and inline assembly, which are on the path towards stabilization. For const-generics, I wish they were more powerful, and that they would support `where` clauses, but there are compiler and nightly options to disable some of their limitations.

- libstd APIs: it's impossible to implement an `Iterator` that returns a `Proxy` value, because `Iterator::next() -&gt; &amp;Target` returns a reference to `Target` instead of just `Target`. This is important for many of my applications, and the workarounds are ugly. Custom allocator support on nightly is nice, hope `AllocRef` can be pushed over the finish line next year. 

- edition 2021: one of the changes I've seen in the last couple of years since the last edition, and planned for the next year, are breaking. So I don't understand why a new edition in 2021 makes sense. 

- foundation: It would be cool to be able to sponsor Rust work on any of the areas above. I want my contributions to go towards Rust work, which Mozilla does not allow, and also be tax deductible.

Most issues mentioned have received some work, but except for some notable exceptions, most work in these areas: SIMD, GPU, FFI, soundness, embedded, and some nightly features (specialization, function traits, ATCs), seems stalled (as in, no work happening right now).
## [5][My Rusty PhD thesis is finally done :)](https://www.reddit.com/r/rust/comments/ilwkrl/my_rusty_phd_thesis_is_finally_done/)
- url: https://www.ralfj.de/blog/2020/09/03/phd.html
---

## [6][If you want performance, cheat!](https://www.reddit.com/r/rust/comments/imb1a1/if_you_want_performance_cheat/)
- url: https://vorner.github.io/2020/09/03/performance-cheating.html
---

## [7][Space Factory game progress - Recall Singularity in Sep 2020 (A game in Rust+Godot, much networking discussion)](https://www.reddit.com/r/rust/comments/imca9u/space_factory_game_progress_recall_singularity_in/)
- url: https://medium.com/@recallsingularity/recall-singularity-in-sep-2020-e2f33a85fd7c?source=friends_link&amp;sk=4631322e33e4e49cd65a4726b22be09a
---

## [8][Planning the 2021 Roadmap | Rust Blog](https://www.reddit.com/r/rust/comments/ily2hu/planning_the_2021_roadmap_rust_blog/)
- url: https://blog.rust-lang.org/2020/09/03/Planning-2021-Roadmap.html
---

## [9][OneGraph - A Rust Graph Database](https://www.reddit.com/r/rust/comments/imbd8l/onegraph_a_rust_graph_database/)
- url: https://www.reddit.com/r/rust/comments/imbd8l/onegraph_a_rust_graph_database/
---
Hi everyone,

A few months ago I started to write a graph database in Rust, at the moment it is far from being a functional database.

What I have in mind is to be able to timeline graphs structures with this tool.

I began with an Open Cypher query language implementation.

I'm new to Rust and I wanted to grasp advice on what is already done, any comment is welcome :)

Rust is an amazing tool, and I'm quite happy to have chosen it as the main language for this project.

[https://github.com/alexandre-ricciardi/one-graph](https://github.com/alexandre-ricciardi/one-graph)
## [10][Plotters 0.3 is now released](https://www.reddit.com/r/rust/comments/im0218/plotters_03_is_now_released/)
- url: https://www.reddit.com/r/rust/comments/im0218/plotters_03_is_now_released/
---
Hi everyone, after a long time working, I am very happy to announce the new major update of Plotters, the plotting library in pure Rust.

Here's some details about this major release.

Since this version, some major features are added, including [3D plotting](https://plotters-rs.github.io/plotters-doc-data/3d-plot.svg), some other improvements were made. (You may want to check a [3D WASM demo](https://plumberserver.com/plotters-wasm-demo/index.html#3d-plot)) 

More importantly, since Plotters 0.3, we are able to support third-party crate used as drawing backend which allows anybody targeting Plotters to anything they want very easily.

&amp;#x200B;

Check the [change log](https://github.com/38/plotters/blob/master/CHANGELOG.md) and [release notes](https://github.com/38/plotters/blob/master/RELEASE-NOTES.md) for more details.

Hope you enjoy this new release, and all feedback are warmly welcomed!
## [11][Fontdue: The fastest font renderer in the world, written in pure Rust](https://www.reddit.com/r/rust/comments/ilnd41/fontdue_the_fastest_font_renderer_in_the_world/)
- url: https://github.com/mooman219/fontdue/
---

## [12][Announcing Rust GitHub Template: stop copy pasting md and yml files when creating new projects!](https://www.reddit.com/r/rust/comments/im1fav/announcing_rust_github_template_stop_copy_pasting/)
- url: https://rust-github.github.io
---

