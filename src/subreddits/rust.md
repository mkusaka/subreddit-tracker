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
## [2][This Week in Rust 349](https://www.reddit.com/r/rust/comments/i094wo/this_week_in_rust_349/)
- url: https://this-week-in-rust.org/blog/2020/07/28/this-week-in-rust-349/
---

## [3][First serious release of Polars DataFrames](https://www.reddit.com/r/rust/comments/i14ab0/first_serious_release_of_polars_dataframes/)
- url: https://www.reddit.com/r/rust/comments/i14ab0/first_serious_release_of_polars_dataframes/
---
I've been working on a DataFrame implementation based on Apache Arrow for the past months. The goal of the project is to be a fast in-memory DataFrame library in Rust.  

Now Apache Arrow 1.0 is out, I could release the first serious version. It has got basic selections, groupbys, joins, arithmetic, sorting, filters/masks, Rust iterators and more. 

Python bindings are being worked on. The goal with regard to Python is to have interaction between Apache arrow and numpy with minimal overhead.

[https://github.com/ritchie46/polars](https://github.com/ritchie46/polars)
## [4][Rust 1.45.1 is released](https://www.reddit.com/r/rust/comments/i0rfzs/rust_1451_is_released/)
- url: https://blog.rust-lang.org/2020/07/30/Rust-1.45.1.html
---

## [5][event-listener: a synchronization primitive to turn non-blocking data structures into async or blocking data structures](https://www.reddit.com/r/rust/comments/i16oly/eventlistener_a_synchronization_primitive_to_turn/)
- url: https://docs.rs/event-listener
---

## [6][geos 7.0 release: More type safety, update dependencies and use std TryFrom](https://www.reddit.com/r/rust/comments/i14vla/geos_70_release_more_type_safety_update/)
- url: https://blog.guillaume-gomez.fr/articles/2020-07-31+geos+7.0+release%3A+More+type+safety%2C+update+dependencies+and+use+std+TryFrom
---

## [7][Edition 2021 and beyond](https://www.reddit.com/r/rust/comments/i0n78c/edition_2021_and_beyond/)
- url: https://github.com/rust-lang/rfcs/pull/2966
---

## [8][tihle, a new emulator targeting TI graphing calculators implemented in C and Rust](https://www.reddit.com/r/rust/comments/i0wi89/tihle_a_new_emulator_targeting_ti_graphing/)
- url: https://www.taricorp.net/2020/introducing-tihle/
---

## [9][Reverse engineering a USB device with Rust](https://www.reddit.com/r/rust/comments/i0moov/reverse_engineering_a_usb_device_with_rust/)
- url: https://gill.net.in/posts/reverse-engineering-a-usb-device-with-rust/
---

## [10][Using the Rust compiler as your integration testing framework](https://www.reddit.com/r/rust/comments/i0tyjs/using_the_rust_compiler_as_your_integration/)
- url: https://blog.logrocket.com/using-the-rust-compiler-as-your-integration-testing-framework/
---

## [11][Firefox 79 now supports WebAssembly threads and reference types!](https://www.reddit.com/r/rust/comments/i0g8r7/firefox_79_now_supports_webassembly_threads_and/)
- url: https://www.reddit.com/r/rust/comments/i0g8r7/firefox_79_now_supports_webassembly_threads_and/
---
Firefox 79 includes new WebAssembly functionality:

* First off, seven new built-in operations are provided for [bulk memory operations](https://developer.mozilla.org/docs/WebAssembly/Understanding_the_text_format#Bulk_memory_operations). For example, copying and initializing allow WebAssembly to model native functions such as `memcpy` and `memmove` in a more efficient, performant way.
* The [reference-types proposal](https://github.com/WebAssembly/reference-types) is now supported. It provides a [new type, externref](https://rustwasm.github.io/docs/wasm-bindgen/reference/reference-types.html), which can hold any JavaScript value, for example strings, DOM references, or objects. The [wasm-bindgen](https://rustwasm.github.io/docs/wasm-bindgen/) documentation includes guidance for taking advantage of `externref` from Rust.
* With the return of `SharedArrayBuffer` objects, we’re now also able to support [WebAssembly threads](https://developer.mozilla.org/docs/WebAssembly/Understanding_the_text_format#WebAssembly_threads). Thus, it is now possible for WebAssembly Memory objects to be shared across multiple WebAssembly instances running in separate Web Workers. The outcome? Very fast communication between Workers, as well as significant performance gains in web applications.

(from [https://hacks.mozilla.org/2020/07/firefox-79/](https://hacks.mozilla.org/2020/07/firefox-79/))
## [12][Installation using Chocolately](https://www.reddit.com/r/rust/comments/i14zz8/installation_using_chocolately/)
- url: https://www.reddit.com/r/rust/comments/i14zz8/installation_using_chocolately/
---
I noticed that you can install Rust using Choco on Windows. Any advantages or disadvantages compared with rustup?

There are two variations, just "rust" and "rust-ms". I am using Microsoft tools, does that mean "rust-ms" is the preferred one?

I especially care about the ability to debug and integrate with Visual Studio Code.
