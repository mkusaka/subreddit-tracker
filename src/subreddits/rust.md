# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (26/2020)!](https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 344](https://www.reddit.com/r/rust/comments/hepkfq/this_week_in_rust_344/)
- url: https://this-week-in-rust.org/blog/2020/06/23/this-week-in-rust-344/
---

## [3][WWM: Windows Window Manager is now feature complete](https://www.reddit.com/r/rust/comments/hg37c5/wwm_windows_window_manager_is_now_feature_complete/)
- url: https://www.reddit.com/r/rust/comments/hg37c5/wwm_windows_window_manager_is_now_feature_complete/
---
Hey!

I am now posting again, after I finally managed to implement all of the features that I needed.  Now is the phase where I need to find bugs and fix them, so I'd appreciate it if you open an issue once you find one.

Since this is my first rust project, I am also asking for some feedback on my source code. I had to refactor it numerous times, because of either borrow checker limitations or windows api related stuff.

I suck at writing documentation, so if there is anything confusing or missing, please let me know.

Overall I'm very happy that I decided to use Rust, because after using it in this project and getting to know the quirks of the language I can say that I really enjoy writing Rust.

[Repository](https://github.com/TimUntersberger/wwm) [Latest Release](https://github.com/TimUntersberger/wwm/releases/tag/v0.9) [Demo](https://gfycat.com/glisteningbrighthound)
## [4][cargo-fuzz is now 10x faster, better supports sanitizers](https://www.reddit.com/r/rust/comments/hg6x3s/cargofuzz_is_now_10x_faster_better_supports/)
- url: https://www.reddit.com/r/rust/comments/hg6x3s/cargofuzz_is_now_10x_faster_better_supports/
---
[Fuzzing](https://en.wikipedia.org/wiki/Fuzzing) is a [highly](https://lcamtuf.coredump.cx/afl/#bugs) [effective](https://github.com/rust-fuzz/trophy-case) way of discovering bugs, including security vulnerabilities. [`cargo-fuzz`](https://github.com/rust-fuzz/cargo-fuzz) is a cargo subcommand that makes fuzzing easy.

`cargo-fuzz` has picked up a lot of improvements since the beginning of the year:

 * **~10x faster execution** for code built in release mode thanks to tweaks to [fuzzing instrumentation](https://github.com/rust-fuzz/cargo-fuzz/pull/229) and [compiler flags](https://github.com/rust-fuzz/cargo-fuzz/pull/215).
 * Code is now compiled in [release mode with debug assertions](https://github.com/rust-fuzz/cargo-fuzz/pull/231) by default instead of debug mode. This brings **another ~10x speedup** to the default configuration.
 * **Added `cargo fuzz fmt` command** to print the fuzzer-generated data via its `Debug` implementation. This is particularly useful in conjunction with [structure-aware fuzzing](https://rust-fuzz.github.io/book/cargo-fuzz/structure-aware-fuzzing.html).
 * **Much better support for Memory Sanitizer.** Now it "Just Works" for pure-Rust code. Code linking to C still requires passing [extra flags](https://doc.rust-lang.org/unstable-book/compiler-flags/sanitizer.html#memorysanitizer) to C compiler.
 * **Support for fuzzing without any sanitizers.** This is useful for testing 100% safe code where you don't have to watch out for memory errors.
 * [Many smaller improvements and fixes.](https://github.com/rust-fuzz/cargo-fuzz/blob/master/CHANGELOG.md)

Fuzzing Rust code has never been easier! Check out the [Rust Fuzz Book](https://rust-fuzz.github.io/book/cargo-fuzz.html) to get started.
## [5][nanoserde: How I cut 50s of clean build time and 1s of incremental build time](https://www.reddit.com/r/rust/comments/hfru5a/nanoserde_how_i_cut_50s_of_clean_build_time_and/)
- url: https://www.reddit.com/r/rust/comments/hfru5a/nanoserde_how_i_cut_50s_of_clean_build_time_and/
---
I tried to benchmark how much of build time I actually saved by migrating from serde to nanoserde. 

https://github.com/not-fl3/nanoserde-bench
https://github.com/not-fl3/nanoserde/

The benchmark is pretty close to my typical serde use case: read some json and send some binary serialized messages.

Turned out that nanoserde reduced clean build time by ~50s, incremental build time by ~1s and stripped ~50Kb from wasm filesize a bonus.
## [6][Yew or Seed for the frontend, which one would you pick and for which reason(s)?](https://www.reddit.com/r/rust/comments/hfzeqm/yew_or_seed_for_the_frontend_which_one_would_you/)
- url: https://www.reddit.com/r/rust/comments/hfzeqm/yew_or_seed_for_the_frontend_which_one_would_you/
---

## [7][Is this a proper functor?](https://www.reddit.com/r/rust/comments/hg3tk1/is_this_a_proper_functor/)
- url: https://www.reddit.com/r/rust/comments/hg3tk1/is_this_a_proper_functor/
---
I got this code compiling on the latest nightly(1.46) and it runs properly too, but I'm not sure if it has the proper functor functionality. Can someone with some fp experience tell me if there's anything wrong with the code?

    #![feature(generic_associated_types)]
    trait Functor {
    	type Unplug;
    	type Plug&lt;B&gt;: Functor;
    
    	fn fmap&lt;B, F&gt;(self, f: F) -&gt; Self::Plug&lt;B&gt;
    	where
    		F: Fn(Self::Unplug) -&gt; B;
    }
    
    #[derive(Debug)]
    struct MyFunctor&lt;A&gt;(A);
    
    impl&lt;A&gt; Functor for MyFunctor&lt;A&gt; {
    	type Unplug = A;
    	type Plug&lt;B&gt; = MyFunctor&lt;B&gt;;
    
    	fn fmap&lt;B, F&gt;(self, f: F) -&gt; MyFunctor&lt;B&gt;
    	where F: Fn(Self::Unplug) -&gt; B {
    	    MyFunctor(f(self.0))
    	}
    }
    
    fn main() {
        let a = MyFunctor(vec!{1,2,3});
        let b = a.fmap(|inner| format!{"{:?}", inner}).fmap(|inner| inner.len());
        println!{"{:?}", b};
    }

[playground](https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=bd66179ca8ee90087d53e76970edf218)
## [8][Crabapple - A framework for creating iOS tweaks using the Rust programming language](https://www.reddit.com/r/rust/comments/hfyjiq/crabapple_a_framework_for_creating_ios_tweaks/)
- url: https://github.com/Crabapple-iOS/crabapple
---

## [9][A practical guide to async in Rust](https://www.reddit.com/r/rust/comments/hg7hp8/a_practical_guide_to_async_in_rust/)
- url: http://blog.logrocket.com/a-practical-guide-to-async-in-rust/
---

## [10][Possible rustc bug with panic=abort flag](https://www.reddit.com/r/rust/comments/hg4ody/possible_rustc_bug_with_panicabort_flag/)
- url: https://www.reddit.com/r/rust/comments/hg4ody/possible_rustc_bug_with_panicabort_flag/
---
Hi all,

I believe I have found a rustc bug, but I wanted to get a second opinion first. Maybe I'm just stupid.

Adding `-Cpanic=abort` to my RUSTFLAGS causes a build which compiled perfectly fine before to exit with SIGABRT. Here is a link to a minimal CI job that reproduces this problem:

[https://github.com/rageagainsthepc/libmicrovmi/runs/810490910#step:5:99](https://github.com/rageagainsthepc/libmicrovmi/runs/810490910#step:5:99)

Unfortunately I was unable to narrow the issue down any further because the build also compiles just fine if I use a subset of those dependencies.

My rust toolchain version:

    stable-x86_64-unknown-linux-gnu (default)
    rustc 1.44.1 (c7087fe00 2020-06-17)

Any help/hints would be greatly appreciated.
## [11][The Matrix homeserver Conduit just got support for cross-signing](https://www.reddit.com/r/rust/comments/hfs8si/the_matrix_homeserver_conduit_just_got_support/)
- url: https://www.youtube.com/watch?v=lPAAxyg2UD4
---

## [12][Web App using React Rocket Framework](https://www.reddit.com/r/rust/comments/hg51sw/web_app_using_react_rocket_framework/)
- url: https://www.reddit.com/r/rust/comments/hg51sw/web_app_using_react_rocket_framework/
---
Getting started with rocket framework understand and build a web app with rocket

[Getting Started with Rust Rocket](https://frogtok.com/extremely-simple-rust-rocket-framework-tutorial/)
