# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (17/2020)!](https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (17/2020)?](https://www.reddit.com/r/rust/comments/g4nupj/whats_everyone_working_on_this_week_172020/)
- url: https://www.reddit.com/r/rust/comments/g4nupj/whats_everyone_working_on_this_week_172020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-17-2020/41247?u=llogiq)!
## [3][rust-analyzer Changelog #21](https://www.reddit.com/r/rust/comments/g4pm86/rustanalyzer_changelog_21/)
- url: https://rust-analyzer.github.io/thisweek/2020/04/20/changelog-21.html
---

## [4][Active Rust developers estimated at 0.6 million (pdf, page 10)](https://www.reddit.com/r/rust/comments/g4kp2f/active_rust_developers_estimated_at_06_million/)
- url: https://s3-eu-west-1.amazonaws.com/vm-blog/uploads/2020/04/DE18-SoN-Digital-.pdf
---

## [5][Working on a build system that can fetch dependencies from online tar files without using Cargo, because Rust feels too stable and I'm bored](https://www.reddit.com/r/rust/comments/g4bzjx/working_on_a_build_system_that_can_fetch/)
- url: https://i.redd.it/fe1ikcajbtt41.png
---

## [6][(FAIL) Amethyst Engine for the Gameshell! Rust/GCC/Fridge/ARM](https://www.reddit.com/r/rust/comments/g4my0e/fail_amethyst_engine_for_the_gameshell/)
- url: https://youtu.be/YVmk82nxahM
---

## [7][MeiliSearch forum on GitHub Discussions (beta hype!)](https://www.reddit.com/r/rust/comments/g4pckj/meilisearch_forum_on_github_discussions_beta_hype/)
- url: https://github.com/meilisearch/MeiliSearch/discussions
---

## [8][Define UB in float-to-int casts to saturate](https://www.reddit.com/r/rust/comments/g4jeb0/define_ub_in_floattoint_casts_to_saturate/)
- url: https://github.com/rust-lang/rust/pull/71269
---

## [9][Multiversion 0.5.0, now "production ready"](https://www.reddit.com/r/rust/comments/g4hgcm/multiversion_050_now_production_ready/)
- url: https://crates.io/crates/multiversion
---

## [10][Wired Logic - a pixel-based digital circuit simulator running in a browser (Rust compiled into WASM).](https://www.reddit.com/r/rust/comments/g478yt/wired_logic_a_pixelbased_digital_circuit/)
- url: https://iostapyshyn.github.io/wired-logic/
---

## [11][clippy-dirty: `cargo clippy` for uncommited files](https://www.reddit.com/r/rust/comments/g4q9ej/clippydirty_cargo_clippy_for_uncommited_files/)
- url: https://www.reddit.com/r/rust/comments/g4q9ej/clippydirty_cargo_clippy_for_uncommited_files/
---
Hi,

I've created a small script that invokes `cargo clippy` and shows only those messages that affect uncommited (i.e. dirty) files:

https://github.com/Patryk27/clippy-dirty

I'm currently using it in a project where applying all changes at once would be quite cumbersome, while using this tool allows me to prepare small improvements as-I-go.
## [12][Building a high performance realtime charting library in the browser using Rust/WASM](https://www.reddit.com/r/rust/comments/g4fly0/building_a_high_performance_realtime_charting/)
- url: https://www.reddit.com/r/rust/comments/g4fly0/building_a_high_performance_realtime_charting/
---
I've written a high performance time series charting library in raw javascript and I'm interested in porting the code over to Rust/WASM to get some further performance enhancements. The charts are all drawn to a html canvas element. The charts are powered by realtime websockets and the chart gets updated in a requestAnimationFrame loop roughly 60 times a second.

My understanding is Rust/WASM cannot directly interact with DOM elements like the canvas yet and instead has to interface with some kind of API calls, which naturally have an overhead. Is there anyway for me to know if Rust/WASM can handle my requirements... Is drawing to a canvas element through Rust 60 times a second feasible?

There is a bunch of computationally heavy code that gets run every animation loop (dataset can have 20k elements, lots of calculations too).

I'm hoping someone here has some experience in this area and can share what they learned! Thanks.

Edit: to also add, the chart currently listens to all kinds of mouse events - like clicking, hover, dragging and window resizing. Can the rust/wasm code listen and respond to these events fast enough?
