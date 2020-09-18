# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (38/2020)!](https://www.reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 356](https://www.reddit.com/r/rust/comments/iu3ge0/this_week_in_rust_356/)
- url: https://this-week-in-rust.org/blog/2020/09/16/this-week-in-rust-356/
---

## [3][doukutsu-rs: An reimplementation of Cave Story engine written in Rust](https://www.reddit.com/r/rust/comments/iv2t25/doukutsurs_an_reimplementation_of_cave_story/)
- url: https://github.com/alula/doukutsu-rs
---

## [4][Potential improvements for Rust embedded abstractions](https://www.reddit.com/r/rust/comments/iv31ni/potential_improvements_for_rust_embedded/)
- url: https://tweedegolf.nl/blog/42/potential-improvements-for-rust-embedded-abstractions
---

## [5][Intra-doc links close to stabilization | Inside Rust Blog](https://www.reddit.com/r/rust/comments/iuk0m9/intradoc_links_close_to_stabilization_inside_rust/)
- url: https://blog.rust-lang.org/inside-rust/2020/09/17/stabilizing-intra-doc-links.html
---

## [6][How to get started with Rust? - For beginners](https://www.reddit.com/r/rust/comments/iuzht8/how_to_get_started_with_rust_for_beginners/)
- url: https://edfloreshz.blog/how-to-get-started-with-rust
---

## [7][My Favorite Rust Function Signature](https://www.reddit.com/r/rust/comments/iullue/my_favorite_rust_function_signature/)
- url: https://www.brandonsmith.ninja/blog/favorite-rust-function
---

## [8][Is Rust a Functional Language in Disguise?](https://www.reddit.com/r/rust/comments/iur9ly/is_rust_a_functional_language_in_disguise/)
- url: https://www.reddit.com/r/rust/comments/iur9ly/is_rust_a_functional_language_in_disguise/
---
[https://ceronman.com/2020/09/17/is-rust-a-functional-language-in-disguise/](https://ceronman.com/2020/09/17/is-rust-a-functional-language-in-disguise/)

Hi everyone! This is my first post to this subreddit. I've been learning Rust for the past few weeks. Here is a blog post I wrote with some thoughts about the language and about functional programming. I hope you like it!
## [9][Announcing YBC | Yew Bulma Components](https://www.reddit.com/r/rust/comments/iuuz4g/announcing_ybc_yew_bulma_components/)
- url: https://www.reddit.com/r/rust/comments/iuuz4g/announcing_ybc_yew_bulma_components/
---
YBC is a Yew component library based on the Bulma CSS framework.

- [Github repo](https://github.com/thedodd/ybc)
- [API Docs](https://docs.rs/ybc/0.1.2/ybc/)
- [Bulma CSS framework](https://bulma.io/)
- [The Yew project](https://yew.rs/)

YBC encapsulates all of the structure, style and functionality of the Bulma CSS framework as a set of Yew components. **YBC also ships with support for the Yew Router,** adding Bulma-styled components which wrap the Yew Router components for clean integration.

As a guiding principal, YBC does not attempt to encapsulate every single Bulma style as a Rust type, let alone the many valid style combinations. That would be far too complex, and probably limiting to the user in many ways. Instead, YBC handles strucutre, required classes, functionality, sane defaults and every component can be customized with any additional classes for an exact look and feel.

**To get started with YBC,** have a look at the [Getting Started](https://github.com/thedodd/ybc#getting-started) guide in the README. A few pertinent highlights:

- YBC works out of the box with Bulma CSS. Add `&lt;link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.0/css/bulma.min.css"/&gt;` to your HTML, and then you're ready to start using YBC.
- YBC also supports full customization using Bulma's recommended customization pattern. [Details here](https://github.com/thedodd/ybc#add-bulma-sass-allows-customization--themes). TL;DR, use [Trunk](https://github.com/thedodd/trunk) for building &amp; bundling your app. It will handle compiling your scss/sass, which is what you will use for customizing Bulma.

Let me know what you think. The hope is that this crate will make building web apps with Rust WASM that much easier. Cheers!
## [10][Trunk 0.4.0 | Layered config (Trunk.toml), JS Snippets &amp; release binaries](https://www.reddit.com/r/rust/comments/iuv3sf/trunk_040_layered_config_trunktoml_js_snippets/)
- url: https://www.reddit.com/r/rust/comments/iuv3sf/trunk_040_layered_config_trunktoml_js_snippets/
---
Build, bundle &amp; ship your Rust WASM application to the web.

- [Github repo](https://github.com/thedodd/trunk)
- [Release notes](https://github.com/thedodd/trunk/releases/tag/v0.4.0)
- [Example Trunk.toml](https://github.com/thedodd/trunk/blob/master/Trunk.toml)

### Pertinent Highlights
- In addition to CLI arguments and options, Trunk now supports layered configuration via `Trunk.toml` &amp; environment variables.
- There is an example `Trunk.toml` to the root of the repository showing all possible config values along with their defaults. Link above.
- JS snippets generated by wasm-bindgen are now fully supported by Trunk. They should have been since the `0.1.0`, but I overlooked that feature `:)`. Docs for [wasm-bindgen JS snippets](https://rustwasm.github.io/docs/wasm-bindgen/reference/js-snippets.html).
- Release binaries will now be uploaded to the Github release page for every future release. This is great for CI/CD. Brew formulae for Mac/Linux, and Choco package for Windows are in the works.
## [11][Box Plots at the Olympics with Rust and Plotly](https://www.reddit.com/r/rust/comments/iv52jn/box_plots_at_the_olympics_with_rust_and_plotly/)
- url: https://datacrayon.com/posts/programming/rust-notebooks/box-plots-at-the-olympics/
---

## [12][aarch64-musl now being built by rustup CI, enabling full toolchain for Alpine and postmarketOS on ARM](https://www.reddit.com/r/rust/comments/iumwxu/aarch64musl_now_being_built_by_rustup_ci_enabling/)
- url: https://www.reddit.com/r/rust/comments/iumwxu/aarch64musl_now_being_built_by_rustup_ci_enabling/
---
https://rust-lang.github.io/rustup-components-history/aarch64-unknown-linux-musl.html

https://github.com/rust-lang/rustup/issues/2003

https://github.com/rust-lang/rust/issues/75721

https://github.com/rust-lang/rust/pull/76420

Very useful for anyone who was unable to build with nightly on embedded or mobile (PinePhone, Librem 5).

The `rustup-init.sh` [script](https://static.rust-lang.org/rustup/dist/aarch64-unknown-linux-musl/rustup-init) isn't quite ready but this is a huge step in that direction. Expect to see it soon!
