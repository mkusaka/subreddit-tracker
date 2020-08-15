# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 351](https://www.reddit.com/r/rust/comments/i8iqb9/this_week_in_rust_351/)
- url: https://this-week-in-rust.org/blog/2020/08/11/this-week-in-rust-351/
---

## [3][Nightly stdlib docs now document all keywords!](https://www.reddit.com/r/rust/comments/ia1vlc/nightly_stdlib_docs_now_document_all_keywords/)
- url: https://www.reddit.com/r/rust/comments/ia1vlc/nightly_stdlib_docs_now_document_all_keywords/
---
If you haven't yet noticed, the standard library documentation includes documentation for every [keyword](https://doc.rust-lang.org/stable/std/#keywords).

As of today, the initial documentation for every keyword has been fleshed out, finishing with [unsafe](https://github.com/rust-lang/rust/pull/73943) and closing out the [tracking issue](https://github.com/rust-lang/rust/issues/34601).

I just want to give a congratulations and a huge thanks to everybody who helped write these docs.
## [4][Frustrated? It's not you, it's Rust.](https://www.reddit.com/r/rust/comments/i9sor7/frustrated_its_not_you_its_rust/)
- url: https://fasterthanli.me/articles/frustrated-its-not-you-its-rust
---

## [5][Why We’re Bringing Astropad Cross-Platform with Rust](https://www.reddit.com/r/rust/comments/ia3x9q/why_were_bringing_astropad_crossplatform_with_rust/)
- url: https://astropad.com/why-rust/
---

## [6][Checkpoint/restore applications at 15GB/s with criu-image-streamer, written in Rust](https://www.reddit.com/r/rust/comments/i9ynps/checkpointrestore_applications_at_15gbs_with/)
- url: https://github.com/checkpoint-restore/criu-image-streamer
---

## [7][What is the recommended way to serve Rocket applications?](https://www.reddit.com/r/rust/comments/ia5cg9/what_is_the_recommended_way_to_serve_rocket/)
- url: https://www.reddit.com/r/rust/comments/ia5cg9/what_is_the_recommended_way_to_serve_rocket/
---
So there is a built-in web server in Rocket, however, according to the documentation, its TLS support isn't production ready. This makes me wonder, is the built-in web server meant for development purposes only (the same way as PHP has a built-in web server too), or this is meant for production, it's just still under development and not ready for production use?

And if the built-in web server is made for development purposes only, then what is the correct way to serve Rocket web applications? For example, can nginx handle it? Or is there any better solution?
## [8][Get rustc to emit llvm-bc](https://www.reddit.com/r/rust/comments/ia47lv/get_rustc_to_emit_llvmbc/)
- url: https://www.reddit.com/r/rust/comments/ia47lv/get_rustc_to_emit_llvmbc/
---
I need the llvm bitcode, but the only thing I can get rustc to produce with the --emit=llvm-bc option are object files, which are one step too far in the compilation chain.  The problem occurs when trying to use rustc with cargo to link all crates used in the project.

Any help would be appreceated.
## [9][A full-fledged Shazam client for Linux, written in Rust](https://www.reddit.com/r/rust/comments/i9rcwa/a_fullfledged_shazam_client_for_linux_written_in/)
- url: https://github.com/marin-m/SongRec
---

## [10][[ANN] Sauron 0.30.0 - introduces node! macro which enables you to write html syntax in the Component's view.](https://www.reddit.com/r/rust/comments/ia3tx9/ann_sauron_0300_introduces_node_macro_which/)
- url: https://www.reddit.com/r/rust/comments/ia3tx9/ann_sauron_0300_introduces_node_macro_which/
---
[Project repo](https://github.com/ivanceras/sauron)

Checkout the example
[Todomvc with node macro syntax](https://github.com/ivanceras/sauron/blob/master/examples/todomvc-macro-syntax/src/app.rs)
for the complete usage of `node!` macro.

Alongside with this release is the integration of [markdown syntax](https://github.com/ivanceras/sauron/blob/master/crates/sauron-markdown/src/markdown.rs) which you allows you to parse markdown safely and integrate the parsed Node tree into the Component's `view`.

[Complete Changelog](https://github.com/ivanceras/sauron/blob/master/Changelog.md)

Special thanks to /u/udoprog for the `node!` macro in this release.


Contributions are welcome.
- Documentations: This project needs someone who can rephrase and use a more concise terms for the documentations.
## [11][Against Glob Imports](https://www.reddit.com/r/rust/comments/i9uvek/against_glob_imports/)
- url: https://drs.is/post/against-globs/
---

## [12][Put some system programming project ideas??](https://www.reddit.com/r/rust/comments/ia6vao/put_some_system_programming_project_ideas/)
- url: https://www.reddit.com/r/rust/comments/ia6vao/put_some_system_programming_project_ideas/
---
System programming Project ideas or the concepts that use for developing the system programming application with Rust.

Please give some ideas?
