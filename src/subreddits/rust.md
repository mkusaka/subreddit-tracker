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
## [2][This Week in Rust 357](https://www.reddit.com/r/rust/comments/iyjzb8/this_week_in_rust_357/)
- url: https://this-week-in-rust.org/blog/2020/09/23/this-week-in-rust-357/
---

## [3][AMD is looking for a "3D Driver Development Engineer" with Rust experience](https://www.reddit.com/r/rust/comments/iz6ric/amd_is_looking_for_a_3d_driver_development/)
- url: https://jobs.amd.com/job/Boxborough-3D-Driver-Development-Engineer-80489-Mass/677678000/
---

## [4][spotifyd needs your help! The maintainer needs 2-3 devs with rust experience to help keep it going..](https://www.reddit.com/r/rust/comments/izfaob/spotifyd_needs_your_help_the_maintainer_needs_23/)
- url: https://github.com/Spotifyd/spotifyd/issues/682
---

## [5][AMD Is Hiring To Work On New Radeon Driver Tooling Written In Rust](https://www.reddit.com/r/rust/comments/izi8q2/amd_is_hiring_to_work_on_new_radeon_driver/)
- url: https://www.phoronix.com/scan.php?page=news_item&amp;px=AMD-Hiring-Radeon-Rust
---

## [6][Are we IPFS yet? v0.2 of the IPFS crate now released!](https://www.reddit.com/r/rust/comments/izdun1/are_we_ipfs_yet_v02_of_the_ipfs_crate_now_released/)
- url: https://areweipfsyet.rs/
---

## [7][Yoshua Wuyts Rust 2021 - let's take a break](https://www.reddit.com/r/rust/comments/iz7lmv/yoshua_wuyts_rust_2021_lets_take_a_break/)
- url: https://blog.yoshuawuyts.com/rust-2021/
---

## [8][Rust in 2021: Leveraging the Type System for Infallible Message Buffers](https://www.reddit.com/r/rust/comments/izextq/rust_in_2021_leveraging_the_type_system_for/)
- url: https://christian.amsuess.com/blog/website/2020-09-24_rust_2021/
---

## [9][What are some traits every rust developer should know by heart?](https://www.reddit.com/r/rust/comments/iz9kqf/what_are_some_traits_every_rust_developer_should/)
- url: https://www.reddit.com/r/rust/comments/iz9kqf/what_are_some_traits_every_rust_developer_should/
---
I'm a new rust developer and after spending 3 long months reading the book, I can pretty much  code without an IDE.

But the challenge is to memorize the common trait implementations on structs and I hopelessly stare at RLS spitting suggestions in VSCODE.

What are those 5 to 10 odd common traits and the impl fns that I should remember which every hands-on developer would be using regularly?

I came across `Into` , `From`, `Read/Writ` , `BufRead` etc.
## [10][Rust 2021 Roadmap Wishlist](https://www.reddit.com/r/rust/comments/izd2ap/rust_2021_roadmap_wishlist/)
- url: https://aldaronlau.com/rust-2021/
---

## [11][Drop order in Rust: It's tricky](https://www.reddit.com/r/rust/comments/iz5r0y/drop_order_in_rust_its_tricky/)
- url: https://vojtechkral.github.io/blag/rust-drop-order/
---

## [12][Announcing rt-graph](https://www.reddit.com/r/rust/comments/iz3tub/announcing_rtgraph/)
- url: https://www.reddit.com/r/rust/comments/iz3tub/announcing_rtgraph/
---
A real-time graphing experiment written in Rust.

&amp;#x200B;

https://preview.redd.it/w7247opxh5p51.png?width=816&amp;format=png&amp;auto=webp&amp;s=a4ad23bf71d288061b4513d4238ae7d2b6412320

Many other graphing tools do not efficiently update the display when new data is added, for example redrawing the whole screen when only a few pixels of new data are added.

This crate tries to do the minimum incremental work required to update the graph when new data is added: draw the few pixels of new data, and scroll the graph with efficient large copies, which can and should be accelerated by GPU hardware.

As a result of this design rt-graph easily copes with 30k new points per second, at 60 FPS, using just 3% CPU (tested on a Lenovo T460 laptop from 2016 with 2.4 GHz Intel Core i5-6300U, running Ubuntu 18.04.5).

[https://crates.io/crates/rt-graph](https://crates.io/crates/rt-graph)
