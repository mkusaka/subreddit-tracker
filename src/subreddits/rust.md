# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (39/2020)!](https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (39/2020)?](https://www.reddit.com/r/rust/comments/iwxjdh/whats_everyone_working_on_this_week_392020/)
- url: https://www.reddit.com/r/rust/comments/iwxjdh/whats_everyone_working_on_this_week_392020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-39-2020/49088?u=llogiq)!
## [3][My Least Favorite Rust Type](https://www.reddit.com/r/rust/comments/ix751t/my_least_favorite_rust_type/)
- url: https://ridiculousfish.com/blog/posts/least-favorite-rust-type.html
---

## [4][Small tip to transform an Option&lt;impl Iterator&lt;_&gt;&gt; into an impl Iterator&lt;_&gt;](https://www.reddit.com/r/rust/comments/ixj7f3/small_tip_to_transform_an_optionimpl_iterator/)
- url: https://www.reddit.com/r/rust/comments/ixj7f3/small_tip_to_transform_an_optionimpl_iterator/
---
If you ever encounter a situation where you need an `impl Iterator&lt;Item = T&gt;` out of an `Option&lt;impl Iterator&lt;Item = T&gt;&gt;`, doing `iter_option.unwrap_or(iter::empty())` will not work because the types in the `None` and `Some` cases are different.

The tip is to flatten the option: `iter_option.into_iter().flatten()`.
## [5][gui-tools -- A cross-platform drawing kit for creating GUIs](https://www.reddit.com/r/rust/comments/ixh1m3/guitools_a_crossplatform_drawing_kit_for_creating/)
- url: https://www.reddit.com/r/rust/comments/ixh1m3/guitools_a_crossplatform_drawing_kit_for_creating/
---
As of late, I've been examining the model that GTK+ uses to create a GUI interface. One of the most interesting parts is that it separates its code into three parts: the widget toolkit (GTK), the drawing toolkit (GDK), and the runtime for the type system (GLib). 

gui-tools aims to be the drawing kit part of that equation. It provides an event runtime, interfaces to native libraries, and APIs for drawing on surfaces. It doesn't aim to be a full graphics framework; rather, it aims to be an unopinionated part of one.

At the moment, it's in a primitive state. supports the xlib and win32 backends somewhat reliably. Some things I want to do with this in the future:

* Add support for text rendering via the `fontdue` crate.
* Add a backend for OSX's AppKit API.
* Add a backend for the DOM. That way it becomes easy to port apps using `gui-tools` from the desktop to the web.
* Currently, `gui-tools` uses native drawing APIs (e.g. GDI+ and GC), which create unnecessary expenditure on the CPU. If we use OpenGL, we can both transfer that work to the GPU and make the drawing style more consistent.
* Add async support.

https://github.com/not-a-seagull/gui-tools
## [6][A simple library for benchmarking](https://www.reddit.com/r/rust/comments/ixlbqi/a_simple_library_for_benchmarking/)
- url: https://www.reddit.com/r/rust/comments/ixlbqi/a_simple_library_for_benchmarking/
---
It imitates test::bench, but supports async and has beautiful output.

&amp;#x200B;

&gt;I am a rust beginner, please correct me if the code is bad. Thank you

&amp;#x200B;

https://preview.redd.it/cduz8lngmoo51.png?width=1108&amp;format=png&amp;auto=webp&amp;s=6111b6fb13f8c3b2aa1a0df1612b4325038b1421

[https://github.com/juzi5201314/bench-rs](https://github.com/juzi5201314/bench-rs)
## [7][Updating the little book of rust macros to current Rust](https://www.reddit.com/r/rust/comments/ix0jni/updating_the_little_book_of_rust_macros_to/)
- url: https://www.reddit.com/r/rust/comments/ix0jni/updating_the_little_book_of_rust_macros_to/
---
The Little Book of Rust Macros is a great learning resource when it comes to macros, or well it was. Unfortunately the original author has disappear since 2016 which in turn kept the book frozen in time. Rust has changed quite a bit since then, and so did macros. The book is still quite helpful even to this day but it is missing some valuable additions that have been made to the macro system. The practical introduction actually guides the user through building a macro that does not even compile in current Rust anymore.

Due to this I want to revive the project and bring it up to date with current Rust so that new people get a great learning resource back when it comes to learning about macros, especially since it is one of the topics a lot of new people seem to struggle with.

I've already started updating the majority of the book, the methodical introduction, to current Rust here https://github.com/Veykril/tlborm.
Help is deeply appreciated as there is a lot of new stuff to cover!
## [8][IntelliJ Rust Changelog #131](https://www.reddit.com/r/rust/comments/ix0wgl/intellij_rust_changelog_131/)
- url: https://intellij-rust.github.io/2020/09/21/changelog-131.html
---

## [9][rust-analyzer changelog #43](https://www.reddit.com/r/rust/comments/iwzja8/rustanalyzer_changelog_43/)
- url: https://rust-analyzer.github.io/thisweek/2020/09/21/changelog-43.html
---

## [10][robusta — a proc macro to ease Rust⬄Java interop!](https://www.reddit.com/r/rust/comments/ix7f0l/robusta_a_proc_macro_to_ease_rustjava_interop/)
- url: https://github.com/giovanniberti/robusta/
---

## [11][Porting EBU R128 audio loudness analysis from C to Rust](https://www.reddit.com/r/rust/comments/ixn7ju/porting_ebu_r128_audio_loudness_analysis_from_c/)
- url: https://coaxion.net/blog/2020/09/porting-ebu-r128-audio-loudness-analysis-from-c-to-rust/
---

## [12][How to get started with Rust?](https://www.reddit.com/r/rust/comments/ixlkla/how_to_get_started_with_rust/)
- url: https://edfloreshz.blog/how-to-get-started-with-rust
---

