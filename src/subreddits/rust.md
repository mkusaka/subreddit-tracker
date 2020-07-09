# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (28/2020)!](https://www.reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 346](https://www.reddit.com/r/rust/comments/hnkws3/this_week_in_rust_346/)
- url: https://this-week-in-rust.org/blog/2020/07/08/this-week-in-rust-346/
---

## [3][Compiled First GTK Application for PinePhone](https://www.reddit.com/r/rust/comments/hnsuw0/compiled_first_gtk_application_for_pinephone/)
- url: https://i.redd.it/pev8hzn75q951.jpg
---

## [4][Rust Lang team design meeting update from Inside Rust Blog](https://www.reddit.com/r/rust/comments/hnygqf/rust_lang_team_design_meeting_update_from_inside/)
- url: https://blog.rust-lang.org/inside-rust/2020/07/08/lang-team-design-meeting-update.html
---

## [5][Moving from TypeScript to Rust / WebAssembly](https://www.reddit.com/r/rust/comments/hnv8qj/moving_from_typescript_to_rust_webassembly/)
- url: https://nicolodavis.com/blog/typescript-to-rust/
---

## [6][Serde AMF](https://www.reddit.com/r/rust/comments/ho1pfy/serde_amf/)
- url: https://www.reddit.com/r/rust/comments/ho1pfy/serde_amf/
---
I've uploaded a crate which I've been using for a couple of months.

It's a amf (flash) de/serializer with support for serde. 

Pretty simple and I think fast (but haven't benchmarked) 

check it out.

[serde_amf](https://crates.io/crates/serde_amf)
## [7][Get started with Rust and Evcxr for Rust (Jupyter) Notebooks [Windows 10]](https://www.reddit.com/r/rust/comments/ho1nzx/get_started_with_rust_and_evcxr_for_rust_jupyter/)
- url: https://m8.fyi/evcxr
---

## [8][[ANN] secrecy v0.7.0: a simple secret-keeping library for Rust](https://www.reddit.com/r/rust/comments/hnnm5c/ann_secrecy_v070_a_simple_secretkeeping_library/)
- url: https://www.reddit.com/r/rust/comments/hnnm5c/ann_secrecy_v070_a_simple_secretkeeping_library/
---
- Docs: https://docs.rs/secrecy
- Crate: https://crates.io/crates/secrecy
- GitHub: https://github.com/iqlusioninc/crates/tree/develop/secrecy

The `secrecy` crate is a simple secret-keeping library with the following goals:

- Make secret access explicit and easy-to-audit via a trait ([ExposeSecret](https://docs.rs/secrecy/latest/secrecy/trait.ExposeSecret.html)). This also makes secret values immutable which helps avoid making accidental copies (e.g. reallocating the backing buffer for a `Vec`)
- Prevent accidental leakage of secrets via channels like debug logging
- Ensure secrets are wiped from memory on drop securely (using the `zeroize` crate)

It [provides wrappers for Rust's main collection types (`SecretBox`, `SecretString`, `SecretVec`)](https://docs.rs/secrecy/latest/secrecy/#box-string-and-vec-wrappers) as well as allowing you to define your own `Secret&lt;T&gt;` types (and even works on `no_std`).

This release primarily focused on improving documentation and small API refinements. There are a few remaining features we'd like to add, but we hope to release a v1.0 soon!
## [9][Rust is the only language that gets `await` syntax right](https://www.reddit.com/r/rust/comments/hnbz78/rust_is_the_only_language_that_gets_await_syntax/)
- url: https://www.reddit.com/r/rust/comments/hnbz78/rust_is_the_only_language_that_gets_await_syntax/
---
At first I was weirded out when the familiar `await foo` syntax [got replaced](https://boats.gitlab.io/blog/post/await-decision-ii/) by `foo.await`, but after working with other languages, I've come round and wholeheartedly agree with this decision. Chaining is just much more natural! And this is without even taking `?` into account:

C#: `(await fetchResults()).map(resultToString).join('\n')`

JavaScript: `(await fetchResults()).map(resultToString).join('\n')`

Rust: `fetchResults().await.map(resultToString).join('\n')`

It may not be apparent in this small example, but the absence of extra parentheses really helps readability if there are long argument lists or the chain is broken over multiple lines. It also plain makes sense because all actions are executed in left to right order.

I love that the Rust language designers think things through and are willing to break with established tradition if it makes things truly better. And the solid versioning/deprecation policy helps to do this with the least amount of pain for users. That's all I wanted to say!

More references:

* [Async-await status report: The syntax question](https://smallcultfollowing.com/babysteps/blog/2019/03/01/async-await-status-report/#the-syntax-question)
* [Making progress in await syntax](https://boats.gitlab.io/blog/post/await-syntax/)
* [Update on await syntax](https://boats.gitlab.io/blog/post/await-decision-ii/)
* [A final proposal for await syntax](https://boats.gitlab.io/blog/post/await-decision/)

---

Edit: after posting this and then reading more about how controversial the decision was, I was a bit concerned that I might have triggered a flame war. Nothing of the kind even remotely happened, so kudos for all you friendly Rustaceans too! &lt;3
## [10][Sailfish: Rust's fastest template engine, 200x faster than handlebars](https://www.reddit.com/r/rust/comments/hnhbmv/sailfish_rusts_fastest_template_engine_200x/)
- url: https://github.com/Kogia-sima/sailfish/tree/master/benches
---

## [11][Announcing Rustup 1.22.1](https://www.reddit.com/r/rust/comments/hnjqfo/announcing_rustup_1221/)
- url: https://blog.rust-lang.org/2020/07/08/Rustup-1.22.1.html
---

## [12][Machine Learning with Rust and WebAssembly](https://www.reddit.com/r/rust/comments/ho1img/machine_learning_with_rust_and_webassembly/)
- url: https://www.secondstate.io/articles/machine-learning/
---

