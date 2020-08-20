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
## [2][This Week in Rust 352](https://www.reddit.com/r/rust/comments/icwq65/this_week_in_rust_352/)
- url: https://this-week-in-rust.org/blog/2020/08/18/this-week-in-rust-352/
---

## [3][Scaling Bevy](https://www.reddit.com/r/rust/comments/id0qzk/scaling_bevy/)
- url: https://bevyengine.org/news/scaling-bevy/
---

## [4][I wrote about running Rust on the Arduino Uno](https://www.reddit.com/r/rust/comments/id83as/i_wrote_about_running_rust_on_the_arduino_uno/)
- url: https://creativcoder.dev/rust-on-arduino-uno
---

## [5][Comparison of Async Ecosystems](https://www.reddit.com/r/rust/comments/id6fs6/comparison_of_async_ecosystems/)
- url: https://runrust.miraheze.org/wiki/Async_crate_comparison#Comparison_of_Async_Ecosystems
---

## [6][Rustsim becomes Dimforge (the organization developing nalgebra, nphysics, ncollide, simba)](https://www.reddit.com/r/rust/comments/ictbha/rustsim_becomes_dimforge_the_organization/)
- url: https://www.dimforge.com/blog/2020/08/08/rustsim-becomes-dimforge/
---

## [7][Rust and C++ interoperability - The Chromium Projects](https://www.reddit.com/r/rust/comments/icpdfm/rust_and_c_interoperability_the_chromium_projects/)
- url: https://www.chromium.org/Home/chromium-security/memory-safety/rust-and-c-interoperability
---

## [8][Tantivy v0.13 released](https://www.reddit.com/r/rust/comments/id1cz3/tantivy_v013_released/)
- url: https://www.reddit.com/r/rust/comments/id1cz3/tantivy_v013_released/
---
This release happens quite late after tantivy 0.12, as I was stretched thin between my daughter and work during COVID. Things are better now.

The main change is the addition of BlockWAND but that should only 
affect you if you use disjunction queries &amp; BM25. Tbh I suspect this feature is not really useful for most people and will just improve the performance on the tantivy search benchmark.


Here is the changelog
----

&amp;#x200B;
Tantivy 0.13 introduce a change in the index format that will require
you to reindex your index (BlockWAND information are added in the skiplist). 
The index size increase is minor as this information is only added for
full blocks.
If you have a massive index for which reindexing is not an option, please contact me
so that we can discuss possible solutions.

- Bugfix in `FuzzyTermQuery` not matching terms by prefix when it should (@Peachball)
- Relaxed constraints on the custom/tweak score functions. At the segment level, they can be mut, and they are not required to be Sync + Send.
- `MMapDirectory::open` does not return a `Result` anymore.
- Change in the DocSet and Scorer API. (@fulmicoton). 
A freshly created DocSet point directly to their first doc. A sentinel value called TERMINATED marks the end of a DocSet.
`.advance()` returns the new DocId. `Scorer::skip(target)` has been replaced by `Scorer::seek(target)` and returns the resulting DocId.
As a result, iterating through DocSet now looks as follows
```rust
let mut doc = docset.doc();
while doc != TERMINATED {
   // ...
   doc = docset.advance();
}
```
The change made it possible to greatly simplify a lot of the docset's code.
- Misc internal optimization and introduction of the `Scorer::for_each_pruning` function. (@fulmicoton)
- Added an offset option to the Top(.*)Collectors. (@robyoung)
- Added Block WAND. Performance on TOP-K on term-unions should be greatly increased. (@fulmicoton, and special thanks 
to the PISA team for answering all my questions!)
## [9][Everything Old is New Again: Binary Security of WebAssembly](https://www.reddit.com/r/rust/comments/icwme7/everything_old_is_new_again_binary_security_of/)
- url: https://www.reddit.com/r/rust/comments/icwme7/everything_old_is_new_again_binary_security_of/
---
[https://www.usenix.org/system/files/sec20-lehmann.pdf](https://www.usenix.org/system/files/sec20-lehmann.pdf)

A very interesting paper that I believe is relevant to Rust for a number of reasons.

1. It brings up issues of memory safety, which are often of interest to this community.
2. The paper directly discusses the inability for webassembly to protect against source level issues common in languages like C/C++.
3. The authors use C, C++, and Rust, and actually developed their analysis tool in Rust

\&gt; Still, vulnerabilities in memory-unsafe source languages can translate to vulnerabilities in WebAssembly binaries. \[..\] We find that many classic vulnerabilities which, due to common mitigations, are no longer exploitable in native binaries, are completely exposed in WebAssembly. Moreover, WebAssembly enables unique attacks, such as overwriting supposedly constant data or manipulating the heap using a stack overflow
## [10][[NES emulator] Plastic v0.1.0 is out](https://www.reddit.com/r/rust/comments/id5bei/nes_emulator_plastic_v010_is_out/)
- url: https://www.reddit.com/r/rust/comments/id5bei/nes_emulator_plastic_v010_is_out/
---
The first release of [plastic](https://github.com/Amjad50/plastic) is out.

This is my second project in Rust, this project is for learning about emulators and about Rust at the same time. 

Not sure if this design is good, but one pattern I use a lot is `Rc&lt;RefCell&lt;&gt;&gt;`, I can see that the performance of this emulator is not as I expected. Please let me know what do you think about this codebase.

&amp;#x200B;

PRs and suggestions are always welcome.

Note: there are pre-compiled binaries for windows and linux.
## [11][async-rwlock: An async reader-writer lock with a fair locking strategy](https://www.reddit.com/r/rust/comments/icvkar/asyncrwlock_an_async_readerwriter_lock_with_a/)
- url: https://github.com/stjepang/async-rwlock
---

## [12][PSA: libz-sys, flate2, and git2 now support the high-performance zlib-ng](https://www.reddit.com/r/rust/comments/icwh2o/psa_libzsys_flate2_and_git2_now_support_the/)
- url: https://www.reddit.com/r/rust/comments/icwh2o/psa_libzsys_flate2_and_git2_now_support_the/
---
[`libz-sys`](https://github.com/rust-lang/libz-sys/) provides Rust bindings to the zlib compression library, for handling DEFLATE/gzip compression. I've just released `libz-sys` 1.1.0, which provides opt-in support for the high-performance `zlib-ng` library in zlib-compat mode.

zlib-ng provides 2-3x performance improvements for compression, and ~20% improvements for decompression.

Most people use `libz-sys` via the higher-level `flate2` crate; you can opt into zlib-ng for `flate2` with:

`flate2 = { version = "1.0.17", features = ["zlib-ng-compat"], default-features = false }`

The git repository format also uses DEFLATE compression extensively, and the `git2` Rust crate now supports zlib-ng as well:

`git2 = { version = "0.13.9", features = ["zlib-ng-compat"] }`

Note that this will build and link a vendored version of `zlib-ng`, rather than using the system shared version of zlib.

If you build a crate that wraps a C library,  and you use `libz-sys` to provide a version of zlib, please *don't* opt into `zlib-ng` by default; instead, use `libz-sys = { version = "1.1.0", default-features = false, features = ["libc"] }`, which allows zlib-ng but doesn't require it. Optionally, you can provide a `zlib-ng-compat` feature by adding `zlib-ng-compat = ["libz-sys/zlib-ng"]` to your features section; alternatively, people who depend on your crate can also depend on libz-sys themselves and enable the zlib-ng feature.

Please give this a try, and comment if it improves the performance of your own crates! If you encounter any issues, please report them via github on the appropriate crates.
