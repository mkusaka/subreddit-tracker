# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (7/2020)!](https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/)
- url: https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 325](https://www.reddit.com/r/rust/comments/f3a5hu/this_week_in_rust_325/)
- url: https://this-week-in-rust.org/blog/2020/02/11/this-week-in-rust-325/
---

## [3][Blog Post: Why is Rust the Most Loved Programming Language](https://www.reddit.com/r/rust/comments/f41ynd/blog_post_why_is_rust_the_most_loved_programming/)
- url: https://matklad.github.io/2020/02/14/why-rust-is-loved.html
---

## [4][Roa: an async web framework inspired by koajs, lightweight but powerful](https://www.reddit.com/r/rust/comments/f45bck/roa_an_async_web_framework_inspired_by_koajs/)
- url: https://github.com/Hexilee/roa
---

## [5][Nimble: Async friendly binary encoding/decoding](https://www.reddit.com/r/rust/comments/f483xm/nimble_async_friendly_binary_encodingdecoding/)
- url: https://www.reddit.com/r/rust/comments/f483xm/nimble_async_friendly_binary_encodingdecoding/
---
I've just created a crate for binary encoding/decoding of Rust types with support for directly writing values to any type implementing `AsyncWrite` and reading values from any type implementing `AsyncRead`. All of this with `async`/`await` syntax!

- Crates.io: https://crates.io/crates/nimble
- Documentation: https://docs.rs/nimble/0.1.0/nimble/
- Repository: https://github.com/devashishdxt/nimble
## [6][Why doesn't this compile? - value still mutably borrowed after downcast to immutable](https://www.reddit.com/r/rust/comments/f44p9v/why_doesnt_this_compile_value_still_mutably/)
- url: https://www.reddit.com/r/rust/comments/f44p9v/why_doesnt_this_compile_value_still_mutably/
---
[https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=b6b1bcb90a0a904ce4f6853eed5ec2c9](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=b6b1bcb90a0a904ce4f6853eed5ec2c9)

I think I know the answer. It's because the compiler keeps track of what is still borrowed in the return value of a function or expression, but it assumes anything that was mutably borrowed to begin with is still mutably borrowed afterwards.

It seems like having the compiler track when a mutable borrow has been transformed into an immutable one would be reasonably easy, but maybe I'm overlooking something. Is there a reason this isn't a feature?
## [7][All chapters from Rust in Action, book published by Manning, are now available](https://www.reddit.com/r/rust/comments/f3pq12/all_chapters_from_rust_in_action_book_published/)
- url: https://www.manning.com/books/rust-in-action?a_aid=rust&amp;a_bid=0367c58f&amp;chan=reddit
---

## [8][Reverse Question: What do you miss about other languages when working in Rust?](https://www.reddit.com/r/rust/comments/f3ur65/reverse_question_what_do_you_miss_about_other/)
- url: https://www.reddit.com/r/rust/comments/f3ur65/reverse_question_what_do_you_miss_about_other/
---
And what don't you miss?
## [9][Pokérust: a Rust wrapper library for PokéAPI (my first crate)](https://www.reddit.com/r/rust/comments/f3zttk/pokérust_a_rust_wrapper_library_for_pokéapi_my/)
- url: https://www.reddit.com/r/rust/comments/f3zttk/pokérust_a_rust_wrapper_library_for_pokéapi_my/
---
Hello /r\/ustaceans, I recently published my first crate: [Pokérust](https://crates.io/crates/pokerust). Pokérust is a Rust wrapper library for the [PokéAPI](https://pokeapi.co/), which provides "All the Pokémon data you'll ever need in one place,
easily accessible through a modern RESTful API".

Every endpoint and API object should be supported, and API responses are automatically cached in memory (no support for an on-disk cache yet).

Let me know if you have any feedback!
## [10][Question: How to run tokio main loop with dynamically configured server address?](https://www.reddit.com/r/rust/comments/f48mok/question_how_to_run_tokio_main_loop_with/)
- url: https://www.reddit.com/r/rust/comments/f48mok/question_how_to_run_tokio_main_loop_with/
---
Hi.

I want to write a server with tokio with dynamically configured tokio server address.

Problem

* event/command/socket based starting and stopping the tokio server
* multiple tokio server inside application
* main loop for event handling

Can I do it with tokio or must I stay with mio?
## [11][What do you miss about Rust when working in other languages?](https://www.reddit.com/r/rust/comments/f3tx37/what_do_you_miss_about_rust_when_working_in_other/)
- url: https://www.reddit.com/r/rust/comments/f3tx37/what_do_you_miss_about_rust_when_working_in_other/
---
And what do you not miss?
## [12][Upcoming compiler-team design meetings | Inside Rust Blog](https://www.reddit.com/r/rust/comments/f3w4qb/upcoming_compilerteam_design_meetings_inside_rust/)
- url: https://blog.rust-lang.org/inside-rust/2020/02/14/upcoming-compiler-team-design-meetings.html
---

