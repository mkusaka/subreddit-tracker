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

## [3][hors - instant coding answers via the command line have released v0.6.4, with pretty colorized output by default](https://www.reddit.com/r/rust/comments/i55s5z/hors_instant_coding_answers_via_the_command_line/)
- url: https://i.redd.it/vdjznmmkthf51.png
---

## [4][Propane: an experimental generator syntax for Rust](https://www.reddit.com/r/rust/comments/i51kvi/propane_an_experimental_generator_syntax_for_rust/)
- url: https://without.boats/blog/propane/
---

## [5][Shared Library Comparison: C, C++, Go, Haskell, Rust, and Swift](https://www.reddit.com/r/rust/comments/i4z6aq/shared_library_comparison_c_c_go_haskell_rust_and/)
- url: https://www.reddit.com/r/rust/comments/i4z6aq/shared_library_comparison_c_c_go_haskell_rust_and/
---
As a little unofficial side project of mine, I spent a couple of evenings rewriting a tiny image processing library into different languages to both learn more about those languages and to see if it's worth proposing a rewrite of some bigger libraries (that need refactoring anyway) into a different language. Currently all in C. I can't share the source code (it's an in-house library) but I thought I'd share the results of the evaluation.

Some background info about the library and the evaluation:

* The library processes floating point numbers representing colors.
* The algorithm design was inspired by functional programming (recursion, no side effects, memoisation, ...).
* Explicit vectorisation and concurrency have been turned off for the evaluation (but stuff like loop vectorisation was left up to each compiler's defaults).
* The tested scenario had only a handful of explicit heap memory allocations (for languages where the language lets you decide between the stack and the heap).
* Functions that operated on structs were rewritten as methods where the language allowed it.
* Haskell used algebraic data types instead of structs.
* The code was compiled as a shared library with C representation / FFI.
* Optimisation was enabled.
* All processing happened on the CPU.
* The library calls where accessed from a binary executable written in C.
* OS was CentOS 7.
* The test used a machine with 32-core Intel Xeon E5 3.2GHz, 128GB of RAM and an SSD.

Here are the results for processing a single 4k image/frame (run 50 times over a set of 50 frames, first run discarded) as well as the size of the code (excluding comments and header/interface files):

|Language|Compiler|LoC|Min|Max|Avg|
|:-|:-|:-|:-|:-|:-|
|C|gcc 4.8.5|723|4.112s|4.176s|4.126s|
|C++|clang 10.0.0|643|4.293s|4.397s|4.373s|
|Go^(1)|go 1.14.6|592|15.224s|15.341s|15.237s|
|Haskell^(2)|ghc 7.6.3|460|8.615s|9.002s|8.726s|
|Haskell|ghc 8.8.4|460|9.257s|9.353s|9.281s|
|Rust|rustc 1.45.0|638|5.524s|5.642s|5.540s|
|Swift^(3)|swiftc 5.3-dev|482|9.006s|9.083s|9.065s|
|Swift^(3) (`-Ounchecked`)|swiftc 5.3-dev|482|2.720s|2.775s|2.746s|

^(1)No optimisation flags available as far as I can tell.

^(2)Run on 2k image as 4k gave a stack overflow error

^(3)There's no official way of calling Swift from C. I used `@_cdecl` which is what Swift is using internally.

===

So what have I learned? Well, to begin with, I've learned some Go, Rust, and Swift basics (I write code in C/C++ quite often and did a bit of Haskell as well) but the three were new to me.

And as for the evaluation, who's the winner?

1. C++, probably.
   1. Virtually no speed penalty compared to C (the original implementation).
   2. Great support for dynamic linking (if you know what not to use).
2. Swift is a very close second.
   1. It is the least verbose language. Haskell had less LoCs, but the lines tended to be longer.
   2. With `-Ounchecked`, it is the fastest language, although I still have to research the implications of what unchecked actually means.
   3. It has good support for dynamic linking / shared libraries, although `@_cdecl` being unofficial is a worry.
   4. It was also the easiest to pick up from the Go-Rust-Swift trio.
3. Rust comes third.
   1. The speed penalty could be an acceptable trade off for better memory leak prevention.
   2. Very disappointed with how painful shared libraries are in Rust (no interface generator, or any kind of interface files).
4. Go comes forth.
   1. I triple checked that my Go implementation was correct, and even profiled the code and just couldn't find out why the the performance suffered so much.
   2. It improved substantially for static build (pretty much to Rust levels) rewriting the calling program into Go as well, but I suspect the compiler just did all sorts of inlining that it couldn't do for dynamic linking to help itself out.
   3. Again, shared libraries are not very straightforward in Go.
5. Haskell comes last (but see the edit).
   1. I had big hopes for Haskell, especially because the library design was inspired by Haskell in the first place. But it wouldn't even run for the full 4k. (I got it to work in the end, tweaking the algorithm to be a bit less lazy, but the times were around 23s.)
   2. Static build here helped, but not significantly.
   3. And shared libraries are as painful in Haskell as they are in Rust and Go.

Edit: I got GHC 8.8.4 and rerun the test for Haskell. Now the 4k test works and actually beats Go. It also gets very close to the older version of GHC that run at 2k. Just for comparison, 2k now runs at 1.324s avg (all the other languages were tested at 4k).
## [6][Knurling-rs: Get a grip on bare-metal Rust](https://www.reddit.com/r/rust/comments/i5d7mf/knurlingrs_get_a_grip_on_baremetal_rust/)
- url: https://ferrous-systems.com/blog/knurling-rs/
---

## [7][Why QEMU should move from C to Rust](https://www.reddit.com/r/rust/comments/i4rpyc/why_qemu_should_move_from_c_to_rust/)
- url: http://blog.vmsplice.net/2020/08/why-qemu-should-move-from-c-to-rust.html
---

## [8][DevAndDev: (Yew+Rocket+Diesel) website helps you finding pair programming friends](https://www.reddit.com/r/rust/comments/i51zpa/devanddev_yewrocketdiesel_website_helps_you/)
- url: https://www.reddit.com/r/rust/comments/i51zpa/devanddev_yewrocketdiesel_website_helps_you/
---
Hey rustaceans,

I've created [DevAndDev](http://devand.dev/), a website where you can find pair programming friends, given your favorite languages and your available time slots.

Why am I posting this here? Because it's **completely written in Rust. Even the frontend.** I've used [Rocket](https://rocket.rs/) for the backend, [Yew](https://yew.rs/) for the UI and [Diesel](http://diesel.rs/) to map Rust structs to Postgres. It's my first "big" project in Rust and it has been a great (and challenging) opportunity to learn this beautiful language.

All the code is available on [GitHub](https://github.com/alepez/devand) and building is automated by GitHub actions.

If you want to pair-program with me, you can reach me at [https://devand.dev/u/alepez](https://devand.dev/u/alepez)

&amp;#x200B;

https://preview.redd.it/sqj2pu4bmgf51.png?width=571&amp;format=png&amp;auto=webp&amp;s=cde04af77f65b9c993e4ad5342dc6c1929c6726f
## [9][First Impressions of Rust](https://www.reddit.com/r/rust/comments/i52nf3/first_impressions_of_rust/)
- url: https://john-millikin.com/first-impressions-of-rust
---

## [10][rorshach - A watchman for your files and directories](https://www.reddit.com/r/rust/comments/i58yx7/rorshach_a_watchman_for_your_files_and_directories/)
- url: https://www.reddit.com/r/rust/comments/i58yx7/rorshach_a_watchman_for_your_files_and_directories/
---
Rorshach is a file watching tool for linux, written in rust. It allows a user to watch a directory and execute certain commands based on the changes in files.

https://github.com/sam09/rorshach

I wrote this as a way to get familiar with rust. I am looking to publish this as a crate, but not really sure if it's useful for others too
## [11][Rust webassembly maturity / ease of use](https://www.reddit.com/r/rust/comments/i5bvsj/rust_webassembly_maturity_ease_of_use/)
- url: https://www.reddit.com/r/rust/comments/i5bvsj/rust_webassembly_maturity_ease_of_use/
---
Hi all, apologies for a slightly vague question. I'm primarily interested in writing physics simulations that work in the browser, and Rust - wasm seems like a good fit.

However, I'm getting pretty stuck on some fairly simple bits &amp; pieces and I'm wondering whether to put this aside and come back in a year or two based on maturity and ease of use.

Context: My experience with Rust is \~ 20hrs. I've got lots of years of professional use of C, C#, Java, Kotlin, Python.

This post is partly asking if others have used rust webassembly successfully, and would do so again, partly asking how mature the rust wasm ecosystem is (are my problems due to immaturity), and partly asking for help on the specific problems.

My specific problems are:

* I want a shared buffer, Vec&lt;u8&gt; that is written by rust, but consumed by javascript (transferred to canvas). I'm having a lot of trouble getting the pointer to the memory transferred to javascript. (subproblems: I can't put #\[wasm\_bindgen\] on a vec&lt;u8&gt; for some reason, my simstate object has the vec&lt;u8&gt; as a field, but because of that, it seems not to be #\[wasm\_bindgen\]-able.
* I want a counter++ to be displayed so I can count frames in rust, but static mut counter: i32 = 0; can't be used in a safe function. i can't make functions unsafe because then I can't export them to javascript? I'm not sure if I can just work around this with an unsafe block or something.
* I'm going to want the simulation to run continuously in its own thread, while the display logic runs on a separate thread. Is this possible at the moment in rust wasm?

My question isn't so much about the answer to those specific problems, it's more asking if this sort of problem is common / resolvable / result of inexperience vs immaturity / etc. And any other advice around that type of thing.

Any other comments / questions welcome.
## [12][This Month in Rust GameDev #12 - July 2020](https://www.reddit.com/r/rust/comments/i4uhta/this_month_in_rust_gamedev_12_july_2020/)
- url: https://rust-gamedev.github.io/posts/newsletter-012
---

