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

## [3][png crate is now ~4x faster, supports APNG](https://www.reddit.com/r/rust/comments/hgrum2/png_crate_is_now_4x_faster_supports_apng/)
- url: https://www.reddit.com/r/rust/comments/hgrum2/png_crate_is_now_4x_faster_supports_apng/
---
[`png`](https://crates.io/crates/png) crate provides a pure-Rust, 100% safe PNG encoder and decoder.

 * Switched from `inflate` to `miniz_oxide` crate for DEFLATE decompression for up to 3x speedup
 * 30% speedup from taking advantage of auto-vectorization in filtering
 * Added support for APNG decoding. `image` crate also updated to support APNG
 * Performed extensive fuzzing, incl. on 32-bit which uncovered some panics and integer overflows
 * Tested the decoder on hundreds of thousands of real-world images, found no decoding failures

This brings `png` crate roughly on par with the C `libpng` in terms of performance! And most of the above has been accomplished nearly single-handedly by /u/HeroicKatora. Kudos!

---

`png` is part of the [image-rs](https://github.com/image-rs) organization that maintains pure-Rust, memory-safe decoders for common image formats. If you have ever loaded an image in Rust, it was through one of these crates.

However, there are still some outstanding issues, and the maintainers have a lot on their plate as it is. If you'd like to get involved, here's a list of self-contained work items that would make for valuable contributions:

* ["no data found" error on decoding a valid JPEG image](https://github.com/image-rs/image/issues/1234)
* ["Index out of range" panic on decoding some GIF files](https://github.com/image-rs/image/issues/1238)
* [infinite loop in into_stream_writer_with_size](https://github.com/image-rs/image/issues/862)
* [Document BitDepth::Sixteen encoding](https://github.com/image-rs/image-png/issues/203)
* [Panic: attempt to subtract with overflow](https://github.com/image-rs/jpeg-decoder/issues/132)

And if you are interested in optimization, help on these issues would be much appreciated:

* [Decoder::decode_internal is slow](https://github.com/image-rs/jpeg-decoder/issues/155) - this is a major performance bottleneck in JPEG decoding, taking up 75% of decoding time.
* [inflate::core::init_tree() is slow](https://github.com/Frommi/miniz_oxide/issues/82) - this slows down decoding of very small PNG images.
## [4][Examining ARM vs X86 Memory Models with Rust](https://www.reddit.com/r/rust/comments/hgkgg2/examining_arm_vs_x86_memory_models_with_rust/)
- url: https://www.nickwilcox.com/blog/arm_vs_x86_memory_model/
---

## [5][Statistics on dynamic linking](https://www.reddit.com/r/rust/comments/hgphem/statistics_on_dynamic_linking/)
- url: https://drewdevault.com/dynlib.html
---

## [6][I found a crate that prints a diff in test failures to help you spot the difference more easily](https://www.reddit.com/r/rust/comments/hgm2kq/i_found_a_crate_that_prints_a_diff_in_test/)
- url: https://github.com/colin-kiegel/rust-pretty-assertions
---

## [7][cargo-fuzz is now 10x faster, better supports sanitizers](https://www.reddit.com/r/rust/comments/hg6x3s/cargofuzz_is_now_10x_faster_better_supports/)
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
## [8][Question to chapter 17.3 of the book](https://www.reddit.com/r/rust/comments/hgpjgl/question_to_chapter_173_of_the_book/)
- url: https://www.reddit.com/r/rust/comments/hgpjgl/question_to_chapter_173_of_the_book/
---
[Here](https://doc.rust-lang.org/book/ch17-03-oo-design-patterns.html) it is said that:

&gt;We need to set `state` to `None` temporarily rather than setting it directly with code like  
`self.state = self.state.request_review();` to get ownership of the `state` value. This ensures `Post` cannot use the old `state` value after we have transformed it into a new state.

My question is, why? If I modify the definition of `Post` to not use `Option` so that the `self.state = ...` is valid, I get the error:

https://preview.redd.it/i3nn58u7te751.png?width=1357&amp;format=png&amp;auto=webp&amp;s=2e59deaf2ea247b9d36c105b6471900e5a362429

Could someone ELI5 to me why taking the ownership is necessary via .take()?
## [9][Rust Stream: Iterators [video]](https://www.reddit.com/r/rust/comments/hgt03b/rust_stream_iterators_video/)
- url: https://youtu.be/lQt0adYPdfQ
---

## [10][French security agency's Rust guide for secured applications](https://www.reddit.com/r/rust/comments/hg85b3/french_security_agencys_rust_guide_for_secured/)
- url: https://anssi-fr.github.io/rust-guide/
---

## [11][Is using an array as a const generic type parameter impossible?](https://www.reddit.com/r/rust/comments/hgju6q/is_using_an_array_as_a_const_generic_type/)
- url: https://www.reddit.com/r/rust/comments/hgju6q/is_using_an_array_as_a_const_generic_type/
---
I recently discovered that const generics made quite a progress in rust and tried to make a tensor (is just a multidimensional array) with it. However, `rustc` panics when compiling the below codes.

    #![feature(const_fn)]
    #![feature(const_generics)]
    #![feature(const_if_match)]
    #![feature(const_loop)]
    
    struct Tensor&lt;const RANK: usize, const SHAPE: [usize; RANK]&gt; {
        shape: [usize; RANK],
        v: [f64; product(SHAPE)] // Remove this, and the error is gone.
    }
    
    const fn product&lt;const N: usize&gt;(v: [usize; N]) -&gt; usize {
        if v.len() &gt; 0 {
            let mut p = 1;
            let mut i = 0;
            while i &lt; v.len() {
                p *= v[i];
                i += 1;
            }
    
            p
        } else {
            0
        }
    }

Error message

    error: internal compiler error: src\librustc_middle\ty\subst.rs:592:17: const parameter `RANK/#0` (Const { ty: usize, val: Param(RANK/#0) }/0) out of range when substituting substs=[]
    
    thread 'rustc' panicked at 'Box&lt;Any&gt;', /rustc/50fc24d8a172a853b5dfe40702d6550e3b8562ba\src\libstd\macros.rs:13:23

I'm confused since this works.

    fn main() {
        const SHAPE: [usize; 5] = [1, 2, 3, 4, 5];
        let v: [f64; product(SHAPE)] = [0f64; product(SHAPE)];
    }

Is this a compiler bug, or I'm doing something wrong here?
## [12][A practical guide to async in Rust](https://www.reddit.com/r/rust/comments/hg7hp8/a_practical_guide_to_async_in_rust/)
- url: http://blog.logrocket.com/a-practical-guide-to-async-in-rust/
---

