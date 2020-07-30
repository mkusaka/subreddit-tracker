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

## [3][Firefox 79 now supports WebAssembly threads and reference types!](https://www.reddit.com/r/rust/comments/i0g8r7/firefox_79_now_supports_webassembly_threads_and/)
- url: https://www.reddit.com/r/rust/comments/i0g8r7/firefox_79_now_supports_webassembly_threads_and/
---
Firefox 79 includes new WebAssembly functionality:

* First off, seven new built-in operations are provided for [bulk memory operations](https://developer.mozilla.org/docs/WebAssembly/Understanding_the_text_format#Bulk_memory_operations). For example, copying and initializing allow WebAssembly to model native functions such as `memcpy` and `memmove` in a more efficient, performant way.
* The [reference-types proposal](https://github.com/WebAssembly/reference-types) is now supported. It provides a [new type, externref](https://rustwasm.github.io/docs/wasm-bindgen/reference/reference-types.html), which can hold any JavaScript value, for example strings, DOM references, or objects. The [wasm-bindgen](https://rustwasm.github.io/docs/wasm-bindgen/) documentation includes guidance for taking advantage of `externref` from Rust.
* With the return of `SharedArrayBuffer` objects, we’re now also able to support [WebAssembly threads](https://developer.mozilla.org/docs/WebAssembly/Understanding_the_text_format#WebAssembly_threads). Thus, it is now possible for WebAssembly Memory objects to be shared across multiple WebAssembly instances running in separate Web Workers. The outcome? Very fast communication between Workers, as well as significant performance gains in web applications.

(from [https://hacks.mozilla.org/2020/07/firefox-79/](https://hacks.mozilla.org/2020/07/firefox-79/))
## [4][fixed_vec v0.1.0: avoiding extra bounds checks using Ghosts of Departed Proofs](https://www.reddit.com/r/rust/comments/i0k1y6/fixed_vec_v010_avoiding_extra_bounds_checks_using/)
- url: https://github.com/Torrencem/fixed_vec/blob/master/post.md
---

## [5][Deno now uses SWC (a TS/JS compiler in Rust) for Typescript type stripping](https://www.reddit.com/r/rust/comments/i02zu3/deno_now_uses_swc_a_tsjs_compiler_in_rust_for/)
- url: https://www.reddit.com/r/rust/comments/i02zu3/deno_now_uses_swc_a_tsjs_compiler_in_rust_for/
---
&gt; bartlomieju: Update: type stripping has been implemented using SWC and takes impressive ~70ms compared to ~1s before the change.

https://github.com/denoland/deno/issues/5432#issuecomment-665591700
## [6][Is there a trick for checking lifetime annotations?](https://www.reddit.com/r/rust/comments/i0bzw8/is_there_a_trick_for_checking_lifetime_annotations/)
- url: https://www.reddit.com/r/rust/comments/i0bzw8/is_there_a_trick_for_checking_lifetime_annotations/
---
I know that you can check types with 

    let _: () = x;

Is there a simple way to do that but for seeing how lifetimes interact?
## [7][Build for linux running on raspberry pi?](https://www.reddit.com/r/rust/comments/i0lxoi/build_for_linux_running_on_raspberry_pi/)
- url: https://www.reddit.com/r/rust/comments/i0lxoi/build_for_linux_running_on_raspberry_pi/
---
i have ordered an raspberry pi 4 which i will be installing ubuntu server on, but i have no clue how to build for this. i have googled and i cant find any **solid** answers; everything was different, so im not safe on what to do.

I'm on windows and using wsl2
## [8][Calling functions via null pointers](https://www.reddit.com/r/rust/comments/i0dhgu/calling_functions_via_null_pointers/)
- url: https://stackoverflow.com/questions/63164973/wat-code-rust-allows-calling-functions-via-null-pointers
---

## [9][Can we VersionReq to any other type?](https://www.reddit.com/r/rust/comments/i0l8gt/can_we_versionreq_to_any_other_type/)
- url: https://www.reddit.com/r/rust/comments/i0l8gt/can_we_versionreq_to_any_other_type/
---
I tried to use the value of VersionReq in the PackageId::new() but it return an error like **(the trait 'util::to\_semver::ToSemver\` is not implemented for \`&amp;semver::version\_req::VersionReq\`).** Is there any way to convert the VersionReq value into supported type of PackagId?
## [10][Creating Linux Packages For Rust Projects (1/2)](https://www.reddit.com/r/rust/comments/i0465c/creating_linux_packages_for_rust_projects_12/)
- url: https://ebbflow.io/blog/vending-linux-1
---

## [11][Porting NeoCogi's Jude3D to rust](https://www.reddit.com/r/rust/comments/i0k8x9/porting_neocogis_jude3d_to_rust/)
- url: https://twitter.com/weloraiby/status/1288638269597638663
---

## [12][What actually is a closure](https://www.reddit.com/r/rust/comments/i0cbnz/what_actually_is_a_closure/)
- url: https://www.reddit.com/r/rust/comments/i0cbnz/what_actually_is_a_closure/
---
I've recently been working on a project that's involved passing and returning lots of closures, and I realize I don't understand quite how closures are implemented in memory. I was wondering if there was a nice explanation of what it looks like *in memory* when creating or returning a closure? Thanks!
