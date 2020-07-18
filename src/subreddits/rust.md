# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (29/2020)!](https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 347](https://www.reddit.com/r/rust/comments/hrc4dt/this_week_in_rust_347/)
- url: https://this-week-in-rust.org/blog/2020/07/14/this-week-in-rust-347/
---

## [3][Simple Rewrite of a Crate 🦀: OnceCell [video]](https://www.reddit.com/r/rust/comments/htd97r/simple_rewrite_of_a_crate_oncecell_video/)
- url: https://youtu.be/YBG8QTO8fNI
---

## [4][RFC: Promote aarch64-unknown-linux-gnu to a Tier-1 Rust target](https://www.reddit.com/r/rust/comments/hsxrbt/rfc_promote_aarch64unknownlinuxgnu_to_a_tier1/)
- url: https://github.com/rust-lang/rfcs/pull/2959
---

## [5][CLI to alert you on Telegram when your builds finish](https://www.reddit.com/r/rust/comments/ht9dye/cli_to_alert_you_on_telegram_when_your_builds/)
- url: https://github.com/Hermitter/tepe
---

## [6][Guidelines on Benchmarking and Rust](https://www.reddit.com/r/rust/comments/ht4uf6/guidelines_on_benchmarking_and_rust/)
- url: https://nbsoftsolutions.com/blog/guidelines-on-benchmarking-and-rust
---

## [7][My Bet on Rust has been Vindicated](https://www.reddit.com/r/rust/comments/hsw959/my_bet_on_rust_has_been_vindicated/)
- url: https://nbsoftsolutions.com/blog/my-bet-on-rust-has-been-vindicated.html
---

## [8][Traits working group 2020 sprint 3 summary](https://www.reddit.com/r/rust/comments/ht0pz7/traits_working_group_2020_sprint_3_summary/)
- url: https://blog.rust-lang.org/inside-rust/2020/07/17/traits-sprint-3.html
---

## [9][What physically owns the Rust organization that services/products are donated to? (And general concerns)](https://www.reddit.com/r/rust/comments/hta9ft/what_physically_owns_the_rust_organization_that/)
- url: https://www.reddit.com/r/rust/comments/hta9ft/what_physically_owns_the_rust_organization_that/
---
I saw recently that ARM donated a c2.large.arm system to the core team, but what is the legal entity that owns this system? If some organization approached such legal entity with an attractive proposal to buy that legal entity, what could happen? Unlike C++ or C there's no ISO standard that defines what Rust actually IS. Isn't the rust language compiler an open source project like any other that could get bought by some company?
## [10][Tutorials or documentation for home automation (in Rust)?](https://www.reddit.com/r/rust/comments/htcizm/tutorials_or_documentation_for_home_automation_in/)
- url: https://www.reddit.com/r/rust/comments/htcizm/tutorials_or_documentation_for_home_automation_in/
---
This summer I finally have time to take over Rust (now I have 0 experience) and I want to program some simple application in a RPi to get info from a movement sensor. 

This is pure hobby and it is a chance for me to learn Rust. 

I would appreciate any Rust information focused on the RaspberryPi, in home automation or server/client apps. Networking is also interesting for me. 

Thanks in advance for the help.
## [11][High level design questions for a Rust websocket server and async](https://www.reddit.com/r/rust/comments/htfblr/high_level_design_questions_for_a_rust_websocket/)
- url: https://www.reddit.com/r/rust/comments/htfblr/high_level_design_questions_for_a_rust_websocket/
---
I've got a fun/learning project I am working on, making a sort of text based MUD, web based with WASM on the client side, so using websockets. The basic flow on the server will be:

A thread accepts connections and puts them into a Lobby.  The Lobby will accept a name from each client to log them in, and when the lobby fills up with N players, it passes the clients on to a game thread, and then the Lobby begins collecting the next N players.

So at any time there will be 1 Lobby thread in play, and a few Game threads in play.  Then Game threads will be accepting messages from each client as they come in, and sending messages out to each client as necessary.  It will not be a uniform broadcast to all clients. I don't want to send information to a client that they can't see, as that is a cheating vector.

So, is this something where async gets me a big win?  Part of me feels like since the threads are nicely isolated with groups of clients, that a non blocking websocket api would be easier/simpler to work with, but the websocket crates out there don't seem to work that way.  It seems like they are either async or blocking, am I wrong about that?  

&amp;#x200B;

If I do use async, what is a good way to structure this? How do I get those futures to complete appropriately?  I'm having trouble holding in my head how I call an async send method on a client's websocket, but get the future to complete 'now' without also blocking the game thread.  Do I need to pump messages to be sent to a separate task and call send().await on each?  Could a slow client block up the sends for others?

&amp;#x200B;

Apologies for the long message, I'm new to Rust async and trying to wrap my head around it.
## [12][Programming languages: Now Rust project looks for a way into the Linux kernel | ZDNet](https://www.reddit.com/r/rust/comments/hstd2v/programming_languages_now_rust_project_looks_for/)
- url: https://www.zdnet.com/article/programming-languages-now-rust-project-looks-for-a-way-into-the-linux-kernel/
---

