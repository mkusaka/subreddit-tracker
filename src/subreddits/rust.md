# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (43/2020)!](https://www.reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (43/2020)?](https://www.reddit.com/r/rust/comments/jdwv7h/whats_everyone_working_on_this_week_432020/)
- url: https://www.reddit.com/r/rust/comments/jdwv7h/whats_everyone_working_on_this_week_432020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-43-2020/50310?u=llogiq)!
## [3][Facebook is hiring a team to work on the Rust compiler and libraries](https://www.reddit.com/r/rust/comments/jfkmxo/facebook_is_hiring_a_team_to_work_on_the_rust/)
- url: https://twitter.com/nadavrot/status/1319003839018614784
---

## [4][Fighting Rust’s Type System: An OOP programmer learns about FP](https://www.reddit.com/r/rust/comments/jfwhit/fighting_rusts_type_system_an_oop_programmer/)
- url: https://thefuntastic.com/blog/fighting-rusts-type-system
---

## [5][wasm-pack is looking for maintainers](https://www.reddit.com/r/rust/comments/jfub8n/wasmpack_is_looking_for_maintainers/)
- url: https://github.com/rustwasm/wasm-pack/issues/928
---

## [6][Anybody know any research opportunities/graduate programs related with rust?](https://www.reddit.com/r/rust/comments/jfxyrk/anybody_know_any_research_opportunitiesgraduate/)
- url: https://www.reddit.com/r/rust/comments/jfxyrk/anybody_know_any_research_opportunitiesgraduate/
---
Couldn't find any relevant information on this online. Anyone?
## [7][Kata Container Agent 2.0 rewritten in rust](https://www.reddit.com/r/rust/comments/jfy1ar/kata_container_agent_20_rewritten_in_rust/)
- url: https://www.reddit.com/r/rust/comments/jfy1ar/kata_container_agent_20_rewritten_in_rust/
---
some critical parts of Kata 2.0 are now rewritten in rust and show a 10-fold improvement in size...

https://medium.com/kata-containers/kata-containers-version-2-0-e45df4dd328  
https://github.com/kata-containers/kata-containers/releases/tag/2.0.0  
https://katacontainers.io/
## [8][Lunatic: A platform for building fast, secure and scalable actor systems written in Rust](https://www.reddit.com/r/rust/comments/jfi06e/lunatic_a_platform_for_building_fast_secure_and/)
- url: https://kolobara.com/lunatic/index.html
---

## [9][Tip for nightly users: use cargo-sweep after toolchain updates](https://www.reddit.com/r/rust/comments/jfdiao/tip_for_nightly_users_use_cargosweep_after/)
- url: https://www.reddit.com/r/rust/comments/jfdiao/tip_for_nightly_users_use_cargosweep_after/
---
So I found out that there is a tool called cargo-sweep, which will clean up your `target` directories to free disk space. It has an option to delete output from versions of the compiler that you no longer have installed.

With nightly, you probably update your compiler fairly often. It turns out, that your projects pile up build output from each compiler version you used. But you probably only care about the latest nightly and you will never use those previous versions again! It wastes so much disk space!

After running cargo sweep, I freed many gigabytes of disk space.

EDIT: here is the actual command you should run after a toolchain update:

`cargo sweep -i -r -v ~/src`

^ this will search for cargo projects under `~/src` and will clean up all their `target` dirs from unneeded build artifacts

* `-i`: keep build artifacts made by installed versions of the compiler (useful if you also use stable or pinned nightlies)
* `-r`: search recursively
* `-v`: verbose (tell you what it's doing)

If you don't have `cargo sweep`, install it: `cargo install cargo-sweep`.
## [10][Any tips for Node.js developers looking into Rust?](https://www.reddit.com/r/rust/comments/jfckz5/any_tips_for_nodejs_developers_looking_into_rust/)
- url: https://www.reddit.com/r/rust/comments/jfckz5/any_tips_for_nodejs_developers_looking_into_rust/
---
So I’m working through the ‘Rust Book’ and am really enjoying learning rust. I have virtually no exposure to C/Cpp so it’s definitely a learning curve. 

If you came from a node/js background, are there any nuggets of wisdom you can share?
## [11][Performance vs. Readability for Noobs](https://www.reddit.com/r/rust/comments/jfpsmr/performance_vs_readability_for_noobs/)
- url: https://www.reddit.com/r/rust/comments/jfpsmr/performance_vs_readability_for_noobs/
---
Rust noob here, apologizing in advance.

For the past few months, I've been slowly going through the exercises on exercism.  one thing I have noted, consistently, is that the solutions with the most stars are always beautiful - and for the most part, very slow.  (Items that use list comprehensions seem to be particularly slow; from the benchmarking I did, it seems that they are noticably slower than a simple for loop.)

It might just be that I am a noob looking at other noob solutions to problems - but even solutions with a lot of stars tend to be slow.

I understand that tuning code is a complex process in any situation.  I also *generally* prefer readability over performance, since I usually need to go over the same code multiple times.  That said, is there a guide to producing performant code, and a set of functions that we should generally avoid?  Or a set of anti-patterns that should be on the guard for in Rust?
## [12][Any suggestions for implementing vtables with winapi?](https://www.reddit.com/r/rust/comments/jfp431/any_suggestions_for_implementing_vtables_with/)
- url: https://www.reddit.com/r/rust/comments/jfp431/any_suggestions_for_implementing_vtables_with/
---
I'm currently trying to implement a custom event sink to use with the winapi crate. To do this I need to implement [IWbemEventSinkVtbl](https://docs.rs/winapi/0.3.9/winapi/um/wbemprov/struct.IWbemEventSinkVtbl.html) and am not sure how to do it.

For instance do I really need to implement the entire parent hierarchy?

If anyone knows of a good example of how to do this I would appreciate it.
