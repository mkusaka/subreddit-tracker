# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (21/2020)!](https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 339](https://www.reddit.com/r/rust/comments/gmyv8h/this_week_in_rust_339/)
- url: https://this-week-in-rust.org/blog/2020/05/19/this-week-in-rust-339/
---

## [3][Drawing SVG Graphs with Rust](https://www.reddit.com/r/rust/comments/goe4aa/drawing_svg_graphs_with_rust/)
- url: https://cetra3.github.io/blog/drawing-svg-graphs-rust/
---

## [4][A Stable Modular ABI for Rust](https://www.reddit.com/r/rust/comments/goav53/a_stable_modular_abi_for_rust/)
- url: https://internals.rust-lang.org/t/a-stable-modular-abi-for-rust/12347
---

## [5][Microsoft's Safe Systems Programming Languages Effort (Video)](https://www.reddit.com/r/rust/comments/godjki/microsofts_safe_systems_programming_languages/)
- url: https://mybuild.microsoft.com/sessions/61de34c5-b111-4ece-928f-541854875862
---

## [6][New crate: defer-drop. Defer dropping your data to a background thread.](https://www.reddit.com/r/rust/comments/go4xcp/new_crate_deferdrop_defer_dropping_your_data_to_a/)
- url: https://www.reddit.com/r/rust/comments/go4xcp/new_crate_deferdrop_defer_dropping_your_data_to_a/
---
Inspired by the [recent post](https://www.reddit.com/r/rust/comments/gntv7l/dropping_heavy_objects_in_another_thread_can_make/) about how dropping large data structures or some kinds of OS resources (such as Windows files) can be very slow, and potentially you can speed up your main thread by punting the objects to a background thread to be dropped, I made a crate that does exactly that: [defer-drop](https://docs.rs/defer-drop/1.0.0/defer_drop/)

In the docs, I wrote up all the reasons you might *not* want to use this (performance slowdown, nondeterminism, unbounded memory growth), but hopefully someone can use it for the things it's useful for.
## [7][Update to LLVM 10 got merged](https://www.reddit.com/r/rust/comments/gnx7y0/update_to_llvm_10_got_merged/)
- url: https://github.com/rust-lang/rust/pull/67759
---

## [8][RustCrypto: contingency plan for new crate releases](https://www.reddit.com/r/rust/comments/go5pf8/rustcrypto_contingency_plan_for_new_crate_releases/)
- url: https://github.com/RustCrypto/traits/issues/102
---

## [9][How does cargo handle duplicate (indirect) dependencies with same version but different features?](https://www.reddit.com/r/rust/comments/goid6w/how_does_cargo_handle_duplicate_indirect/)
- url: https://www.reddit.com/r/rust/comments/goid6w/how_does_cargo_handle_duplicate_indirect/
---
I realize that if my package depends on packages `alpha` and `beta`, and the former depends on `gamma = "1.0"` and the latter on `gamma = "2.0"`, then two copies of `gamma` will be build and linked into my package. But what if instead `alpha` depends on `gamma = { version: "2.0", features = ["one"] }` (and `beta` still only depends on `gamma = "2.0"`)? Would I still end up with two copies of `gamma`? Or just a single copy, that had the union of `features = ["one"]` and `features = []`?

What if `alpha` depends on `gamma = { version: "2.0", features = ["one"] }`, and `beta` depends on `gamma = { version: "2.0", features = ["two"] }`, and features `"one"` and `"two"` for package `gamma` are meant to be incompatible?
## [10][What is the difference between .to_string() and String::from() and .into()](https://www.reddit.com/r/rust/comments/gohlq2/what_is_the_difference_between_to_string_and/)
- url: https://www.reddit.com/r/rust/comments/gohlq2/what_is_the_difference_between_to_string_and/
---
As titled. Sorry I am a beginner.

I thought implementing .to\_string() automatically derives String::from()
## [11][How can we support optional namespaces for registries that support it in Cargo?](https://www.reddit.com/r/rust/comments/goj86f/how_can_we_support_optional_namespaces_for/)
- url: https://www.reddit.com/r/rust/comments/goj86f/how_can_we_support_optional_namespaces_for/
---
This seems like an issue that's so important to this community.  

    [dependencies]
    foo = { package = "richard/foo", version = "0.2", registry="richard_registry"}

Am I wrong in thinking this is all we need to figure out how to do?  We don't even have to discuss changing crates.io at all.  Just let's focus on specifically, how do we make it even an option.  

I created a feature request if you want to show your support or comment.

https://github.com/rust-lang/cargo/issues/8266
## [12][The Matrix homeserver Conduit just got support for media and end to end encryption](https://www.reddit.com/r/rust/comments/go0jg8/the_matrix_homeserver_conduit_just_got_support/)
- url: https://www.youtube.com/watch?v=HoBslDtqC1k
---

