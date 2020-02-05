# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (6/2020)!](https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/)
- url: https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (6/2020)?](https://www.reddit.com/r/rust/comments/ey2xrr/whats_everyone_working_on_this_week_62020/)
- url: https://www.reddit.com/r/rust/comments/ey2xrr/whats_everyone_working_on_this_week_62020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-6-2020/37776?u=llogiq)!
## [3][ZZ is a modern formally provable dialect of C](https://www.reddit.com/r/rust/comments/ez73pi/zz_is_a_modern_formally_provable_dialect_of_c/)
- url: https://github.com/aep/zz
---

## [4][Why Discord is switching from Go to Rust](https://www.reddit.com/r/rust/comments/eytyug/why_discord_is_switching_from_go_to_rust/)
- url: https://blog.discordapp.com/why-discord-is-switching-from-go-to-rust-a190bbca2b1f
---

## [5][Fontfor: CLI tool for render fonts preview in your terminal](https://www.reddit.com/r/rust/comments/ez6poz/fontfor_cli_tool_for_render_fonts_preview_in_your/)
- url: https://www.reddit.com/r/rust/comments/ez6poz/fontfor_cli_tool_for_render_fonts_preview_in_your/
---
Just added tui and font render feature for my tiny project https://github.com/7sDream/fontfor.

It was originally just a script to find fonts for uncommon characters, but now can be used as a simple font viewer!

Some screenshot: [Canvas Render](https://rikka.7sdre.am/files/c7a7a685-3966-4fb4-8e9e-0cc53636e406.png), [Moon Emoji Render](https://rikka.7sdre.am/files/8128a291-27a7-42bd-813e-1136f971cdb3.png).

More screenshot and usage please refer to my [tweet](https://twitter.com/7sDream/status/1224703380871942146) or README.md of the [project](https://github.com/7sDream/fontfor). 

Only tested in macOS and Arch Linux, but feel free to try it and report any issues!

---

Yeah, I  suddenly remember I can show the Moon Emoji Render just use text.

This is render result for character "א", Unicode U+05D0.

🌑🌕🌑🌑🌑🌑🌑🌔🌘🌑🌑🌑

🌓🌕🌘🌑🌑🌑🌑🌕🌕🌕🌗🌑

🌔🌕🌕🌑🌑🌑🌑🌕🌕🌕🌕🌖

🌒🌕🌕🌕🌑🌑🌑🌔🌕🌕🌕🌕

🌑🌔🌕🌕🌕🌑🌑🌑🌔🌕🌕🌖

🌑🌑🌕🌕🌕🌖🌑🌑🌓🌕🌑🌑

🌑🌑🌒🌕🌕🌕🌖🌑🌕🌘🌑🌑

🌑🌒🌕🌔🌕🌕🌕🌖🌕🌑🌑🌑

🌑🌕🌖🌑🌒🌕🌕🌕🌕🌑🌑🌑

🌓🌕🌗🌑🌑🌓🌕🌕🌕🌘🌑🌑

🌔🌕🌕🌑🌑🌑🌔🌕🌕🌕🌑🌑

🌓🌕🌕🌗🌑🌑🌑🌔🌕🌕🌖🌑

🌒🌕🌕🌕🌘🌑🌑🌑🌕🌕🌕🌗

🌑🌕🌕🌕🌖🌑🌑🌑🌑🌕🌕🌕

🌑🌔🌕🌕🌕🌑🌑🌑🌑🌒🌕🌕

🌒🌕🌕🌕🌖🌑🌑🌑🌑🌑🌔🌘
## [6][Why &amp;s[pos..] is slower than s.split_at(pos).1?](https://www.reddit.com/r/rust/comments/ez64g2/why_spos_is_slower_than_ssplit_atpos1/)
- url: https://godbolt.org/z/TBhkjy
---

## [7][Couchbase Rust SDK 1.0.0-alpha.3 Released](https://www.reddit.com/r/rust/comments/ez6hh4/couchbase_rust_sdk_100alpha3_released/)
- url: https://www.reddit.com/r/rust/comments/ez6hh4/couchbase_rust_sdk_100alpha3_released/
---
After lots of rewriting and figuring what approach might be the best one, I've finally managed to get a new alpha out: alpha.3. See [https://github.com/couchbaselabs/couchbase-rs](https://github.com/couchbaselabs/couchbase-rs) if you are interested.

It supports async/await out of the box and you can use it in runtimes like tokio although it has no hard dependency on it. It uses the C-based libcouchbase backend ([https://github.com/couchbase/libcouchbase/](https://github.com/couchbase/libcouchbase/)) so you'll find unsafe code - but only in the place where we need to call the binding. 

Once the lcb-based version is complete we want to start working on a full rust native backend (without changing the API), but this is much more work since lcb does a lot of work in the distributed system context.

We are still extending the API and polishing current ones, but this should be the first release that it is quite close API-wise to a beta/GA one. If you have feedback I'd love to hear it! Oh yeah and we are working on docs as well ;-)
## [8][Exploring some Rust basics with actix-web](https://www.reddit.com/r/rust/comments/ez566b/exploring_some_rust_basics_with_actixweb/)
- url: http://blog.levinotik.com/2020/02/05/exploring-some-rust-basics-with-actix-web/
---

## [9][Rust companies in production list - feel free to add companies you know](https://www.reddit.com/r/rust/comments/ez7m4u/rust_companies_in_production_list_feel_free_to/)
- url: https://github.com/omarabid/rust-companies
---

## [10][Hello, is there any async thread pool?](https://www.reddit.com/r/rust/comments/ez9ag6/hello_is_there_any_async_thread_pool/)
- url: https://www.reddit.com/r/rust/comments/ez9ag6/hello_is_there_any_async_thread_pool/
---
I want pool for diesels `mysqlconnection` and i want it to be async, is there any library for it?
## [11][Cross Compiling Rust on Linux for the Raspberry Pi](https://www.reddit.com/r/rust/comments/eyssub/cross_compiling_rust_on_linux_for_the_raspberry_pi/)
- url: https://chacin.dev/blog/cross-compiling-rust-for-the-raspberry-pi
---

## [12][savefile 0.5.0](https://www.reddit.com/r/rust/comments/eyxoao/savefile_050/)
- url: https://www.reddit.com/r/rust/comments/eyxoao/savefile_050/
---
[https://docs.rs/savefile/0.5.0/savefile/](https://docs.rs/savefile/0.5.0/savefile/)

[https://crates.io/crates/savefile](https://crates.io/crates/savefile)

[https://github.com/avl/savefile](https://github.com//Avl/savefile)

I've been working on a serialization / deserialization crate for rust.

It is made specifically for saving state in complex computer games. It uses a compact binary format, but makes old saves readable by newer versions of the program. This is achieved by a system where structs and enums are annotated with attributes describing the versioning, and using a proc\_macro to generate deserializers for all supported versions.

The reason I'm not just using Serde is because when I started the project, Serde didn't seem to have much support for supporting old file formats when deserializing binary protocols. I haven't followed the story for this with Serde, it may well have improved to the point of doing what savefile does much better :-) .

Even though the file format is fast and efficient, it does (by default) include a schema (at the top of the file, checked once at start of deserialization),  so any mismatch in on-disk format and in-memory format will be detected safely.

There is an auto test suite, and miri does not detect any unsoundness. Of course there may still  be bugs.

One version each of IndexMap, SmallVec, parking\_lot::Mutex, and a few other containers I happened to need are supported out of the box. The story here would probably need to improve to facilitate a broader user base.

Anyway, posting it here since maybe someone might find it interesting, and/or useful!
