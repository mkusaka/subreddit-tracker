# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (13/2020)!](https://www.reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (13/2020)?](https://www.reddit.com/r/rust/comments/fnflsh/whats_everyone_working_on_this_week_132020/)
- url: https://www.reddit.com/r/rust/comments/fnflsh/whats_everyone_working_on_this_week_132020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-13-2020/39864?u=llogiq)!
## [3][Two Years With Rust](https://www.reddit.com/r/rust/comments/fnf7ni/two_years_with_rust/)
- url: http://brooker.co.za/blog/2020/03/22/rust.html
---

## [4][bat v0.13.0 released](https://www.reddit.com/r/rust/comments/fn9jrg/bat_v0130_released/)
- url: https://github.com/sharkdp/bat/releases/tag/v0.13.0
---

## [5][Rust Async and the Terrible, Horrible, No Good, Very Bad Day](https://www.reddit.com/r/rust/comments/fnj12j/rust_async_and_the_terrible_horrible_no_good_very/)
- url: https://medium.com/@KevinHoffman/rust-async-and-the-terrible-horrible-no-good-very-bad-day-348ebc836274
---

## [6][async functions for no_std now available on nightly](https://www.reddit.com/r/rust/comments/fn6zzq/async_functions_for_no_std_now_available_on/)
- url: https://www.reddit.com/r/rust/comments/fn6zzq/async_functions_for_no_std_now_available_on/
---
I noticed today a really interesting PR got resolved on Rust nightly! It used to not be possible with vanilla rust to write **async** functions blocks in **#\[no\_std\]** libraries. Now with this resolved its one step closer to better async/await in embedded and no\_std web assembly.

[https://github.com/rust-lang/rust/issues/56974](https://github.com/rust-lang/rust/issues/56974)

example: [https://github.com/richardanaya/js\_ffi/blob/master/examples/helloworld\_async/src/lib.rs](https://github.com/richardanaya/js_ffi/blob/master/examples/helloworld_async/src/lib.rs)

I was able to shave about 10kb off this wasm module.
## [7][Terminal Tamagotchi in Rust - Beginner Friendly](https://www.reddit.com/r/rust/comments/fnhu78/terminal_tamagotchi_in_rust_beginner_friendly/)
- url: https://github.com/trevarj/termagotchi
---

## [8][cargo-feature - Don't suffer from adding or removing feature](https://www.reddit.com/r/rust/comments/fnj8xp/cargofeature_dont_suffer_from_adding_or_removing/)
- url: https://www.reddit.com/r/rust/comments/fnj8xp/cargofeature_dont_suffer_from_adding_or_removing/
---
[https://github.com/Riey/cargo-feature](https://github.com/Riey/cargo-feature)

&amp;#x200B;

Just type

    cargo feature serde +derive

&amp;#x200B;

You don't need typing this by hand!

    [dependencies]
    serde = "1"

&amp;#x200B;

    [dependencies]
    serde = { version = "1", features = ["derive"] }
## [9][New version of unsafe Qt bindings for Rust is released](https://www.reddit.com/r/rust/comments/fn9dqh/new_version_of_unsafe_qt_bindings_for_rust_is/)
- url: https://rust-qt.github.io/blog/qt_crates_release_0_5/
---

## [10][I've written a Python VM in Rust and open for any feedback](https://www.reddit.com/r/rust/comments/fn45fk/ive_written_a_python_vm_in_rust_and_open_for_any/)
- url: https://github.com/FKLC/PyVM
---

## [11][Explain — a tool to visualize postgresql explain](https://www.reddit.com/r/rust/comments/fni8c7/explain_a_tool_to_visualize_postgresql_explain/)
- url: https://www.reddit.com/r/PostgreSQL/comments/fn08ff/another_explain_visualizer/
---

## [12][Rust bindings for the Philips Hue API](https://www.reddit.com/r/rust/comments/fnc8i0/rust_bindings_for_the_philips_hue_api/)
- url: https://github.com/yuqio/huelib-rs
---

