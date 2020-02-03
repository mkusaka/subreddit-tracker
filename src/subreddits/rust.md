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
## [3][Update on rawloader rustc benchmarks from 1.20 to 1.41 and current beta/nightly](https://www.reddit.com/r/rust/comments/ey4thl/update_on_rawloader_rustc_benchmarks_from_120_to/)
- url: https://www.reddit.com/r/rust/comments/ey4thl/update_on_rawloader_rustc_benchmarks_from_120_to/
---
I've been running this benchmark to try and catch regressions in either my code or rustc versions. The latest version is always here:

http://chimper.org/rawloader-rustc-benchmarks/

The general trend is quite good. Over these two and a half years of rust versions the improvements are:

- 31% faster compile time
- 23% smaller binary size
- 9% faster runtime

Great work by the rust team, thanks!

There have also been a few regressions over time and recently I've noticed three:

- The current beta has a 14% increase in compile time which seems mostly fixed in nightly. Something similar happened in 1.31 and 1.32 with a 11% increase followed by a 15% decrease that more than made up for it.
- 1.36 introduced a 29% increase in binary size that was never recovered.
- There is one specific file that has had a ~5-15% runtime regression ever since 1.21 for which I should probably try and isolate a test case.

I'm particularly curious if anyone knows what the binary size increase was due to.
## [4][std.rs 2.0: Electric Boogaloo](https://www.reddit.com/r/rust/comments/exus04/stdrs_20_electric_boogaloo/)
- url: https://www.reddit.com/r/rust/comments/exus04/stdrs_20_electric_boogaloo/
---
Hey, everyone! I just published a new version of std.rs which has built-in search now (the top rated request from the last thread). That means you can now do all of the following redirects:

| From                       | To                                                      |
| -------------------------- | ------------------------------------------------------- |
| `https://std.rs`           | `https://doc.rust-lang.org/stable/std/`                 |
| `https://n.std.rs`         | `https://doc.rust-lang.org/nightly/std/`                |
| `https://std.rs/n`         | `https://doc.rust-lang.org/nightly/std/`                |
| `https://std.rs/println`   | `https://doc.rust-lang.org/stable/std/?search=println`  |
| `https://n.std.rs/println` | `https://doc.rust-lang.org/nightly/std/?search=println` |
| `https://std.rs/n/println` | `https://doc.rust-lang.org/nightly/std/?search=println` |

This new version is built on CloudFlare workers, so if there are any other redirects which may be helpful, feel free to open a pull request or issue! The source code is available here: https://github.com/knpwrs/std.rs

Original post: https://www.reddit.com/r/rust/comments/es35u6/stdrs_go_directly_to_stable_rust_docs/
## [5][Did anyone else catch the Upwork commercial that name-dropped Rust?](https://www.reddit.com/r/rust/comments/ey0sk3/did_anyone_else_catch_the_upwork_commercial_that/)
- url: https://www.reddit.com/r/rust/comments/ey0sk3/did_anyone_else_catch_the_upwork_commercial_that/
---
Did a double take at first: https://www.ispot.tv/ad/ZCUI/upwork-upwork-is-how-developer

(Hopefully this isn’t too off topic)
## [6][FOSDEM rust devroom videos](https://www.reddit.com/r/rust/comments/ey6m5x/fosdem_rust_devroom_videos/)
- url: https://video.fosdem.org/2020/K.3.401/
---

## [7][Rust: Built to Last](https://www.reddit.com/r/rust/comments/ey12r0/rust_built_to_last/)
- url: https://www.welcometothejungle.com/en/articles/btc-rust-programming-language
---

## [8][Synchrony: A distributed job execution &amp; scheduling engine](https://www.reddit.com/r/rust/comments/ey16qg/synchrony_a_distributed_job_execution_scheduling/)
- url: https://www.reddit.com/r/rust/comments/ey16qg/synchrony_a_distributed_job_execution_scheduling/
---
[Synchrony](https://github.com/Protryon/synchrony) is a project I threw together over the weekend that is designed to be a partial drop-in replacement for Sidekiq (a similar ruby project), and a general job execution/scheduler tool that scales horizontally.

It uses Redis by default, but is designed to be able to use different store types. It has a fully implemented &amp; tested HTTP API powered by Iron.

This is my first large-ish Rust project, so any feedback &amp; criticism is appreciated!
## [9][My newest and most absurd Rust project](https://www.reddit.com/r/rust/comments/exrqdl/my_newest_and_most_absurd_rust_project/)
- url: https://github.com/adam-mcdaniel/free
---

## [10][WASM and Native Gameboy Emulator written in Rust](https://www.reddit.com/r/rust/comments/expoza/wasm_and_native_gameboy_emulator_written_in_rust/)
- url: https://github.com/benkonz/gameboy_emulator
---

## [11][Rust HTTP Testing with Mockito](https://www.reddit.com/r/rust/comments/ey51i4/rust_http_testing_with_mockito/)
- url: https://zupzup.org/rust-http-testing/
---

## [12][I had a crack at an example GraphQL Webserver in Rust!](https://www.reddit.com/r/rust/comments/ey501c/i_had_a_crack_at_an_example_graphql_webserver_in/)
- url: https://www.reddit.com/r/rust/comments/ey501c/i_had_a_crack_at_an_example_graphql_webserver_in/
---
For any interested, it uses:

\- actix-web

\- juniper

\- dataloader-rs

&amp;#x200B;

Would love any feedback, and happy to discuss :)

[https://github.com/jayy-lmao/rust-graphql-docker](https://github.com/jayy-lmao/rust-graphql-docker)

&amp;#x200B;

Edit: Oooh and it uses a minimal docker container, see references for details ;)
