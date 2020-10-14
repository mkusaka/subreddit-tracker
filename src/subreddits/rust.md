# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (42/2020)!](https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (42/2020)?](https://www.reddit.com/r/rust/comments/j9l0o7/whats_everyone_working_on_this_week_422020/)
- url: https://www.reddit.com/r/rust/comments/j9l0o7/whats_everyone_working_on_this_week_422020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-42-2020/49997?u=llogiq)!
## [3][Theory and practice with global application state and statics](https://www.reddit.com/r/rust/comments/jawyeo/theory_and_practice_with_global_application_state/)
- url: https://www.reddit.com/r/rust/comments/jawyeo/theory_and_practice_with_global_application_state/
---
So, good Rust and global statics don't typically go together. When they _are_ used they are typically hidden behind library code to avoid exposing them and instead support RAII-friendly interfaces (e.g. sdl2 bindings wrapping sdl initialization with a unit struct and Drop impl). On the other hand, application code doesn't necessarily have to fit to the conventions that Rust encourages if they don't _need_ to, depending on the needs of the project, or simply because it makes code more readable for the specific use case.

I'm curious to know; for those of you actively developing applications in Rust, do you use statics to represent global application state, or avoid it entirely and prefer other methods to represent that (e.g. dependency injection in web frameworks like actix)?
## [4][Low-Level Academy, an explorable systems programming course](https://www.reddit.com/r/rust/comments/jafssj/lowlevel_academy_an_explorable_systems/)
- url: https://lowlvl.org
---

## [5][Resources about the server side portion of oauth2 in Rust](https://www.reddit.com/r/rust/comments/jax3w3/resources_about_the_server_side_portion_of_oauth2/)
- url: https://www.reddit.com/r/rust/comments/jax3w3/resources_about_the_server_side_portion_of_oauth2/
---
I'm wondering if any of you know of any resources (e.g. videos,  articles, code examples) that helps to explain the server side portion  of oauth2, ideally using rust code and the [oxide-auth](https://github.com/HeroicKatora/oxide-auth) crate as the  examples in that crate are not clear and the other crates that I have found seem to only either be for client side oauth2 or rely on an external service in some way
## [6][Dotenv-linter v2.2.0: find and fix problems in .env files. What's new?](https://www.reddit.com/r/rust/comments/jayzzb/dotenvlinter_v220_find_and_fix_problems_in_env/)
- url: https://evrone.com/dotenv-linter-v220
---

## [7][SixtyFPS: New GUI framework written in Rust (alpha)](https://www.reddit.com/r/rust/comments/jac66t/sixtyfps_new_gui_framework_written_in_rust_alpha/)
- url: https://github.com/sixtyfpsui/sixtyfps/
---

## [8][Miniquad/macroquad/nanoserde got GH sponsors page!](https://www.reddit.com/r/rust/comments/jaj9bf/miniquadmacroquadnanoserde_got_gh_sponsors_page/)
- url: https://github.com/sponsors/not-fl3/
---

## [9][I present you bato](https://www.reddit.com/r/rust/comments/janizv/i_present_you_bato/)
- url: https://www.reddit.com/r/rust/comments/janizv/i_present_you_bato/
---
Hi,

If like me you use a minimalist windows manager on a daily basis (no desktop environment) and your Linux runs on a laptop, this may interest you.

&amp;#x200B;

[bato](https://github.com/doums/bato) is a simple binary that sends desktop notifications on some battery events:

* level full
* level low
* level critical
* charging
* discharging

It works nicely with [Dunst](https://dunst-project.org/).

Configuration in Yaml.

&amp;#x200B;

Hope it will be useful, even a little bit :)
## [10][pin-project 1.0 released](https://www.reddit.com/r/rust/comments/jai12s/pinproject_10_released/)
- url: https://www.reddit.com/r/rust/comments/jai12s/pinproject_10_released/
---
[pin-project](https://crates.io/crates/pin-project) is a library that allows you to use [pin-projection](https://doc.rust-lang.org/nightly/std/pin/index.html#projections-and-structural-pinning) safely and ergonomically.

Mainly changes of 1.0 release are the removal of deprecated APIs. See [release note](https://github.com/taiki-e/pin-project/releases/tag/v1.0.0) for more.
## [11][Building an async-compatible actor system (and act-zero 0.3 released)](https://www.reddit.com/r/rust/comments/jakcca/building_an_asynccompatible_actor_system_and/)
- url: https://github.com/Diggsey/posts/blob/master/actor-systems/README.md
---

## [12][rlua, generate teal types?](https://www.reddit.com/r/rust/comments/jas1qg/rlua_generate_teal_types/)
- url: https://www.reddit.com/r/rust/comments/jas1qg/rlua_generate_teal_types/
---
I recently discovered [teal](https://github.com/teal-language/tl) and it seems like it solves the problems I have with the lua parts of my current project, without requiring me to find another way to embed a scripting language in rust (as rust + rlua + lua are really nice).  


rlua does at first glance seem to make it easy enough to reuse its way to expose methods to generate the types, as pretty much everything is generic.  
However, because rust doesn't allow you to get the names of types I would need a new trait that has a "teal\_type" method, returning the name it should have in teal.  


This however would require more strict type bounds and thus I would need to make my own UserData trait that once implemented also implements the rlua one.  


Is this the accepted way of doing things? Or is there another way I am missing that is nicer to work with? I'm also worried about the foreign trait rule (or whatever it was called exactly) as I need to implement a trait from another crate for everything that implements mine. Which will obviously give problems with the types already present like i8, String, etc.  


in case it helps, this is the documentation (with an example) on how to expose methods/data in rlua to lua [https://docs.rs/rlua/0.17.0/rlua/trait.UserData.html](https://docs.rs/rlua/0.17.0/rlua/trait.UserData.html)
