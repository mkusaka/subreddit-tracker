# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (19/2020)!](https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (19/2020)?](https://www.reddit.com/r/rust/comments/gd6gvk/whats_everyone_working_on_this_week_192020/)
- url: https://www.reddit.com/r/rust/comments/gd6gvk/whats_everyone_working_on_this_week_192020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-19-2020/42034?u=llogiq)!
## [3][Created a renderer following Vulkan Tutorial](https://www.reddit.com/r/rust/comments/gd7c0u/created_a_renderer_following_vulkan_tutorial/)
- url: https://www.reddit.com/r/rust/comments/gd7c0u/created_a_renderer_following_vulkan_tutorial/
---
https://crates.io/crates/sarekt

* The renderer is designed to be extensible to multithreading, multiple pipelines, have more backends (D3D, etc nothing in the user interface depends on vulkan/ash directly) and be more fleshed out.

* Due to this it isn't so similar to the code in Vulkan Tutorial, but I feel like it's a reasonable first renderer architecture that can be reused in other projects (which I plan to do shortly).

* If you check the tags, however, I tagged when I finished implementing and creating an example of each chapter in the tutorial (numbered and named, eg 17_vertex_buffer_creation).

* There's still a lot wrong with it, and I have tons of issues filed on myself, but I'm really happy I got through it all :)

Special thanks to the whole community and all these fantastic libraries i have leveraged, this would not have been possible without everyone's contribution.  The readme has called out a few libraries in particular.
## [4][Tour of Rust](https://www.reddit.com/r/rust/comments/gcu5l2/tour_of_rust/)
- url: https://richardanaya.github.io/tour_of_rust/
---

## [5][Rust in an instant](https://www.reddit.com/r/rust/comments/gd0hay/rust_in_an_instant/)
- url: https://fnordig.de/2020/05/02/rust-in-an-instant/
---

## [6][Point of WebGPU on native](https://www.reddit.com/r/rust/comments/gdbgoc/point_of_webgpu_on_native/)
- url: http://kvark.github.io/web/gpu/native/2020/05/03/point-of-webgpu-native.html
---

## [7][What's the state of Rust in iOS and Android?](https://www.reddit.com/r/rust/comments/gcym93/whats_the_state_of_rust_in_ios_and_android/)
- url: https://www.reddit.com/r/rust/comments/gcym93/whats_the_state_of_rust_in_ios_and_android/
---
I'm thinking of developing a cross-platform application that covers all 4 main platforms (mac, windows, iOS and Android). The UI would be done in the native technology for those platforms while the business logic would be implemented as a library.

I now need to choose between C++ and Rust. What is the state of the Rust compiler on iOS and Android? Is it safe to go down the Rust road for these platforms or things work because they just happen to work?
## [8][Would you use rust if it weren't as fast as C?](https://www.reddit.com/r/rust/comments/gd8jdd/would_you_use_rust_if_it_werent_as_fast_as_c/)
- url: https://www.reddit.com/r/rust/comments/gd8jdd/would_you_use_rust_if_it_werent_as_fast_as_c/
---

## [9][Magnifying Glasses for Rust Assembly](https://www.reddit.com/r/rust/comments/gd1wtd/magnifying_glasses_for_rust_assembly/)
- url: https://www.justanotherdot.com/posts/magnifying-glasses-for-rust-assembly.html
---

## [10][Rust concurrency: a streaming workflow, served with a side of back-pressure.](https://www.reddit.com/r/rust/comments/gda8gr/rust_concurrency_a_streaming_workflow_served_with/)
- url: https://www.reddit.com/r/rust/comments/gda8gr/rust_concurrency_a_streaming_workflow_served_with/
---
How to write concurrent streaming work pipelines in Rust, with back-pressure.

[https://medium.com/@polyglot\_factotum/rust-concurrency-a-streaming-workflow-served-with-a-side-of-back-pressure-955bdf0266b5](https://medium.com/@polyglot_factotum/rust-concurrency-a-streaming-workflow-served-with-a-side-of-back-pressure-955bdf0266b5)
## [11][rlua is now part of the Amethyst organisation](https://www.reddit.com/r/rust/comments/gcoipo/rlua_is_now_part_of_the_amethyst_organisation/)
- url: https://github.com/amethyst/rlua/issues/174
---

## [12][An unescaped \ immediately before a line-break causes the line-break and all whitespace at the beginning of the next line to be ignored](https://www.reddit.com/r/rust/comments/gcw8yn/an_unescaped_immediately_before_a_linebreak/)
- url: https://doc.rust-lang.org/stable/reference/tokens.html?highlight=esc#string-literals
---

