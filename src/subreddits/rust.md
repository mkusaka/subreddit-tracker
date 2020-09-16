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
## [2][What's everyone working on this week (38/2020)?](https://www.reddit.com/r/rust/comments/ismi9w/whats_everyone_working_on_this_week_382020/)
- url: https://www.reddit.com/r/rust/comments/ismi9w/whats_everyone_working_on_this_week_382020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-38-2020/48763?u=llogiq)!
## [3][Dropbox open sources protobuf codegen!](https://www.reddit.com/r/rust/comments/ittov9/dropbox_open_sources_protobuf_codegen/)
- url: https://www.reddit.com/r/rust/comments/ittov9/dropbox_open_sources_protobuf_codegen/
---
Hey everyone! At Dropbox we built our own protobuf framework to meet our production needs. We're now open sourcing it!

Back in 2015 when we were building our [Storage System](https://dropbox.tech/infrastructure/inside-the-magic-pocket) we needed a framework that supported zero copy de-serialization, which prompted the creation of our own library. Since, we've began using it for several parts of Dropbox, including our [Sync Engine](https://dropbox.tech/infrastructure/rewriting-the-heart-of-our-sync-engine). Along with zero copy de-serialization we also provide a number of "Rustic" proto extensions. 

Feel free to give it a look, file an issue, open a PR, and stay on the lookout for more open source Rust libraries from Dropbox 

[GitHub](https://github.com/dropbox/pb-jelly) | [crates.io](https://crates.io/crates/pb-jelly)

&amp;#x200B;

P.S. proto service generation coming soon...
## [4][Blog Post: Challenging LR Parsing](https://www.reddit.com/r/rust/comments/itsr7u/blog_post_challenging_lr_parsing/)
- url: https://rust-analyzer.github.io/blog/2020/09/16/challeging-LR-parsing.html
---

## [5][Should I publish a library of which I believe that no one would use it?](https://www.reddit.com/r/rust/comments/itr5qn/should_i_publish_a_library_of_which_i_believe/)
- url: https://www.reddit.com/r/rust/comments/itr5qn/should_i_publish_a_library_of_which_i_believe/
---
Hello, I wrote [u16\_tic\_tac\_toe](https://github.com/EikeSchulze/u16_tic_tac_toe), a library that uses an u16 to represent a tic-tac-toe grid.

I personally can not see anyone wanting to use a library for such a trivial problem, so I wanted to get the opinions of the community, whether I should publish this to [crates.io](https://crates.io).

So, what do you think?

[View Poll](https://www.reddit.com/poll/itr5qn)
## [6][Initial C API for hyper (Pull Request)](https://www.reddit.com/r/rust/comments/itgcar/initial_c_api_for_hyper_pull_request/)
- url: https://github.com/hyperium/hyper/pull/2278
---

## [7][js-sandbox: securely embed JavaScript into Rust code](https://www.reddit.com/r/rust/comments/ittnir/jssandbox_securely_embed_javascript_into_rust_code/)
- url: https://www.reddit.com/r/rust/comments/ittnir/jssandbox_securely_embed_javascript_into_rust_code/
---
**js-sandbox** is a crate with a minimal API that allows calling JavaScript from Rust. It is based on the Deno project, which again uses the V8 engine.

Very basic example:

```rust
fn main() -&gt; Result&lt;(), ErrBox&gt; {
    let js_code = "function triple(a) { return 3 * a; }"
    let mut script = Script::from_string(js_code)?;
    
    let arg = 7;
    let result: i32 = script.call("triple", &amp;arg)?;
    
    assert_eq!(result, 21);
    Ok(())
}
```

More examples are available [on the GitHub page](https://github.com/Bromeon/js-sandbox) or [docs.rs](https://docs.rs/js-sandbox/).

Possible use cases are plugin systems or mods for games. JavaScript may not be a typical language for this, however it has been battle-tested for security and there are [many languages that transpile to JS](https://github.com/jashkenas/coffeescript/wiki/List-of-languages-that-compile-to-JS).

Plans for the near future are updating to latest deno\_core version, as well as extending the API to cover basic use cases. I would like to keep js-sandbox as simple and minimal as possible, it's meant to be a high-level abstraction and not a massively configurable interpreter :)
## [8][Which Parsing Approach?](https://www.reddit.com/r/rust/comments/itjxvv/which_parsing_approach/)
- url: https://tratt.net/laurie/blog/entries/which_parsing_approach.html
---

## [9][I've been working on adding WSL support to IntelliJ Rust for the past few months. Should make Windows and Rust development a bit nicer for everyone!](https://www.reddit.com/r/rust/comments/it82wo/ive_been_working_on_adding_wsl_support_to/)
- url: https://github.com/intellij-rust/intellij-rust/pull/5014
---

## [10][Announcing KeySeeBee and Keyberon v0.1.0](https://www.reddit.com/r/rust/comments/ita9va/announcing_keyseebee_and_keyberon_v010/)
- url: https://i.redd.it/zu2iys17rbn51.jpg
---

## [11][Announcing juggle v0.1.0, async cooperative multitasking library.](https://www.reddit.com/r/rust/comments/itg5p1/announcing_juggle_v010_async_cooperative/)
- url: https://crates.io/crates/juggle
---

## [12][Next-Generation Programming: Rust &amp; Elm with Richard Feldman](https://www.reddit.com/r/rust/comments/itufpa/nextgeneration_programming_rust_elm_with_richard/)
- url: https://youtu.be/ukVqQGbxM9A?list=PLEx5khR4g7PL-JwckuOkkc5cR6X5hn6ug
---

