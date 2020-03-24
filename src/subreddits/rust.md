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
## [3][Tokio's new internal semaphore based on intrusive lists](https://www.reddit.com/r/rust/comments/fo066j/tokios_new_internal_semaphore_based_on_intrusive/)
- url: https://github.com/tokio-rs/tokio/commit/acf8a7da7a64bf08d578db9a9836a8e061765314
---

## [4][How I structure my apps (in Rust and other languages)](https://www.reddit.com/r/rust/comments/fnyird/how_i_structure_my_apps_in_rust_and_other/)
- url: https://dpc.pw/how-i-structure-my-apps-in-rust-and-other-languages
---

## [5][Broom: An ergonomic, easy to use garbage collector for your toy programming languages](https://www.reddit.com/r/rust/comments/fnscsx/broom_an_ergonomic_easy_to_use_garbage_collector/)
- url: https://github.com/zesterer/broom
---

## [6][Inviting God's Wrath with Cursed Rust](https://www.reddit.com/r/rust/comments/fo487h/inviting_gods_wrath_with_cursed_rust/)
- url: http://troubles.md/abusing-rustc/
---

## [7][gfx-hal-0.5 is released!](https://www.reddit.com/r/rust/comments/fnq64x/gfxhal05_is_released/)
- url: https://www.reddit.com/r/rust/comments/fnq64x/gfxhal05_is_released/
---
This is mostly a polishing release, see the [CHANGELOG](https://github.com/gfx-rs/gfx/blob/fdfe887f5afab7ca057e3e824817d0bcdcf794ba/CHANGELOG.md#hal-050-23-03-2020). The existing APIs are yak-shaved, a few more is added. At the same time, we are moving closer to WebGPU needs, in particular with:

  * read-only storage binding flags. Specifying them is an optimization on DX and Metal. Users can always keep them off as in Vulkan mode.
  * sizes are passed to vertex and index buffer bindings. Again, this can be an optimization, and users can always pass them without bounds as in Vulkan.
  * support for NDC space that is Y-flipped, to match DX and Metal.

[gfx-ocean](https://github.com/gfx-rs/gfx-ocean) demo has been moved under gfx-rs organization, it's a great demo of using gfx-hal directly.
## [8][A Rust library built on top of libbpf that allows to write and load bpf programs (no bcc dependency).](https://www.reddit.com/r/rust/comments/fnzktq/a_rust_library_built_on_top_of_libbpf_that_allows/)
- url: https://github.com/uccidibuti/rebpf
---

## [9][Rust Async and the Terrible, Horrible, No Good, Very Bad Day](https://www.reddit.com/r/rust/comments/fnj12j/rust_async_and_the_terrible_horrible_no_good_very/)
- url: https://medium.com/@KevinHoffman/rust-async-and-the-terrible-horrible-no-good-very-bad-day-348ebc836274
---

## [10][SIMD Array-of-Structures-of-Arrays in nalgebra and comparison with ultraviolet](https://www.reddit.com/r/rust/comments/fnt1ru/simd_arrayofstructuresofarrays_in_nalgebra_and/)
- url: https://www.rustsim.org/blog/2020/03/23/simd-aosoa-in-nalgebra/
---

## [11][Rust analyser weekly changelog](https://www.reddit.com/r/rust/comments/fnl8s1/rust_analyser_weekly_changelog/)
- url: https://rust-analyzer.github.io/thisweek/2020/03/23/changelog-17.html
---

## [12][Former Haskellers: Do you miss purity in rust? Why.](https://www.reddit.com/r/rust/comments/fnl4px/former_haskellers_do_you_miss_purity_in_rust_why/)
- url: https://www.reddit.com/r/rust/comments/fnl4px/former_haskellers_do_you_miss_purity_in_rust_why/
---

