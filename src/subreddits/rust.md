# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (4/2020)!](https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/)
- url: https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (4/2020)?](https://www.reddit.com/r/rust/comments/eraop8/whats_everyone_working_on_this_week_42020/)
- url: https://www.reddit.com/r/rust/comments/eraop8/whats_everyone_working_on_this_week_42020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-4-2020/37181?u=llogiq)!
## [3][My FOSS Story](https://www.reddit.com/r/rust/comments/er4qrn/my_foss_story/)
- url: https://blog.burntsushi.net/foss/
---

## [4][How refactoring a macro reduced my Rust project compilation time from 4 hours to 40 seconds](https://www.reddit.com/r/rust/comments/er1xkh/how_refactoring_a_macro_reduced_my_rust_project/)
- url: https://users.rust-lang.org/t/5-hours-to-compile-macro-what-can-i-do/36508
---

## [5][Rust on the ESP32 - SVD's, PAC's and USB flashing](https://www.reddit.com/r/rust/comments/era9uw/rust_on_the_esp32_svds_pacs_and_usb_flashing/)
- url: https://mabez.dev/blog/posts/esp32-rust-svd-pac/
---

## [6][Count cache misses in Rust](https://www.reddit.com/r/rust/comments/erbnz8/count_cache_misses_in_rust/)
- url: https://www.reddit.com/r/rust/comments/erbnz8/count_cache_misses_in_rust/
---
Hello! I'm doing some research using Rust and I've found a significant difference in execution time between two test cases. Both tests use the same algorithm but on different test data but one is about 10x slower than the other. I *think* this is due to the layout of the two tests in memory.

Is there a tool I can use to estimate the number of cache misses?
## [7][A friendly utility crate for DateTime (using chrono underneath). Inspired by Rails.](https://www.reddit.com/r/rust/comments/erbd8e/a_friendly_utility_crate_for_datetime_using/)
- url: https://github.com/jk-gan/aion
---

## [8][Porting Java's ConcurrentHashMap to Rust (part 3) [video]](https://www.reddit.com/r/rust/comments/eqwn0s/porting_javas_concurrenthashmap_to_rust_part_3/)
- url: https://www.youtube.com/watch?v=_z4fZkmlvrs
---

## [9][Announcing Stowaway: A library for packing data into pointers](https://www.reddit.com/r/rust/comments/er26op/announcing_stowaway_a_library_for_packing_data/)
- url: https://www.reddit.com/r/rust/comments/er26op/announcing_stowaway_a_library_for_packing_data/
---
While working on a toy async runtime and `RawWaker`, I found I was getting frustrated working with with vtables and `*mut ()` and ensuring the conversions to my inner context type were being performed correctly. At the same time, I didn't want to have to unconditionally store everything in a `Box` or `Arc` if unnecessary.

This led me to make [stowaway](https://docs.rs/stowaway/0.1.1/stowaway/index.html): a library for a box-like container that packs its data directly into a `*mut ()` if it's small enough, and only allocates for data that is larger than a pointer. It is designed as a way to pass data into any API that expects a "context" pointer, which is a common pattern in C callback APIs, libraries like `boost::context`, and of course, `RawWaker`.

The library is full documented and tested. I'd like to get to a 1.0 release as soon as possible, but first I wanted to ask /r/rust for some light review; as you can imagine, a library like this involves a lot of `unsafe`, and I wanted to make sure it was all being done correctly.

Next steps planned for the library are miri and fuzz tests, to ensure my implementation is correct, and numerous additional trait implementations on `Stowaway`.
## [10][Rust, WebAssembly, and the future of Serverless by Steve Klabnik](https://www.reddit.com/r/rust/comments/er21mr/rust_webassembly_and_the_future_of_serverless_by/)
- url: https://www.youtube.com/watch?v=CMB6AlE1QuI
---

## [11][Which one is better api?](https://www.reddit.com/r/rust/comments/era5pe/which_one_is_better_api/)
- url: https://www.reddit.com/r/rust/comments/era5pe/which_one_is_better_api/
---
I'm currently working on a new http framework. Whiling designing api, I have an issue with api, so I'm writing this post to hear the (potential) user's opinion.

&amp;#x200B;

Which one seems like a better api?

&amp;#x200B;

`/// One box is created per function which accepts Json&lt;T&gt;`

`#[post("/typed")]`

`pub fn typed(body: Json&lt;LoginForm&gt;) {}`

&amp;#x200B;

`/// Everything is statically dispatched`

`#[post("/annotated")]`

`pub fn annotated(#[json] body: LoginForm) {}`

&amp;#x200B;

Also, note that it's easy to support both.
## [12][ttyplot-rs: Plot streaming data from stdin to a tty terminal](https://www.reddit.com/r/rust/comments/er5atl/ttyplotrs_plot_streaming_data_from_stdin_to_a_tty/)
- url: https://github.com/mogenson/ttyplot-rs
---

