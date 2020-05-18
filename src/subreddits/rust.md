# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (21/2020)!](https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (21/2020)?](https://www.reddit.com/r/rust/comments/glvkxz/whats_everyone_working_on_this_week_212020/)
- url: https://www.reddit.com/r/rust/comments/glvkxz/whats_everyone_working_on_this_week_212020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-21-2020/42814?u=llogiq)!
## [3][Rust sessions at Microsoft Build 2020](https://www.reddit.com/r/rust/comments/glv1zk/rust_sessions_at_microsoft_build_2020/)
- url: https://mybuild.microsoft.com/sessions?t=%257B%2522from%2522%253A%25222020-05-19T00%253A00%253A00%252B02%253A00%2522%252C%2522to%2522%253A%25222020-05-21T23%253A59%253A00%252B02%253A00%2522%257D&amp;q=rust&amp;s=%257B%2522name%2522%253A%2522translate.refine.label.sort.relevance%2522%252C%2522type%2522%253A0%257D
---

## [4][Rust/WinRT Brings Microsoft Closer to Adopting Rust Internally](https://www.reddit.com/r/rust/comments/glr1w0/rustwinrt_brings_microsoft_closer_to_adopting/)
- url: https://www.infoq.com/news/2020/05/rust-winrt-microsoft/
---

## [5][Educational Rust live coding - Building a web app - Head first into WASM](https://www.reddit.com/r/rust/comments/glx2ka/educational_rust_live_coding_building_a_web_app/)
- url: https://youtu.be/668vvEzzMnI
---

## [6][Rust Stream: CLI Project - Part 2 [video]](https://www.reddit.com/r/rust/comments/glxg81/rust_stream_cli_project_part_2_video/)
- url: https://youtu.be/cLk47gn6A1M
---

## [7][rust-analyzer changelog #25](https://www.reddit.com/r/rust/comments/gm159f/rustanalyzer_changelog_25/)
- url: https://rust-analyzer.github.io/thisweek/2020/05/18/changelog-25.html
---

## [8][proc_macro_hygiene to land on 1.45.0 stable](https://www.reddit.com/r/rust/comments/gllkal/proc_macro_hygiene_to_land_on_1450_stable/)
- url: https://github.com/rust-lang/rust/pull/68717#issuecomment-629847207
---

## [9][directories-next &amp; dirs-next: The new home for directories &amp; dirs!](https://www.reddit.com/r/rust/comments/gli4pc/directoriesnext_dirsnext_the_new_home_for/)
- url: https://www.reddit.com/r/rust/comments/gli4pc/directoriesnext_dirsnext_the_new_home_for/
---
Hey folks, 

Given that the original maintainer of directories &amp; dirs-rs crate decided to mark it as "abandoned" for a few years due to certain reasons, we have decided to fork it and move it to a new organisation. 

So presenting the first official announcement of our fork!

We have moved the repos (including the sys bindings) to a single github repository for easier maintenance: 

**Links**
Github: [https://github.com/xdg-rs/dirs](https://github.com/xdg-rs/dirs)  

Crates-io:  
directories-next: [https://crates.io/crates/directories-next](https://crates.io/crates/directories-next)  
dirs-next: [https://crates.io/crates/dirs-next](https://crates.io/crates/dirs-next)  
dirs-sys-next: [https://crates.io/crates/dirs-sys-next](https://crates.io/crates/dirs-sys-next)  

**Team**:   
Lzu Tao (lzutao) &amp; Me (Dylan DPC) 

**Want to join us in maintaining the crates?**
Reply to this post or DM me here or on discord/zulip (@dpc ). We would love to give an opportunity to newcomers or beginners in the community.

UPDATE: we received  messages from people and we are happy with the team size at the moment. We are not looking to add more people to the team. Thank you 

**Why the -next convention?**
We decided that to make it easier for people to migrate and knowing the link between the old &amp; the new crates, the best way would be to name it the same way and add `-next` suffix. 

**How to make the change?**
Renaming the dependency in Cargo.toml and adding -next to the name and imports should suffice. We considered aliases similar to how it is used in some other places in the ecosystem, but we concluded that it could confuse people in the future. We are working on a script that makes it easier for people to migrate. (If anyone wants to do it, feel free)

**Has the API changed?**
The API has remained the same. The only change besides organisational changes is the change in minimum rust versions. The old crates had a minimum rust version of 1.13 (ancient :P ) which we have bumped to **1.36**. The crate *may* still work for older versions (going all the back to 1.13) but we are no longer testing it or assure to maintain it for pre-1.36 versions in the future. 

**Was the original author contacted?**
We contacted the original author of the crates and received no reply. Our original idea was to fork and publish to the same crate for minimal disturbance. Sadly, that wasn't possible so we had to release the new crates. You can consider this the de-facto fork going forward.

**The version?**
The last version of the original crate is 2.0.2. We discussed continuing the version, but it could be misleading. So we decided to "reset" it and since the API is largely stable, have gone forward with 1.0.0 

**What's next?**
The API is largely stable. The first thing we want to work on is to update the crates that use this, and polish the codebase to use 2018 edition and polish it with newer features wherever applicable in a backwards-compatible fashion. Post that, the crates will receive maintenance updates. 

**Have a fork?**
It's been a while since the announcement of the original crates was made, so a few users might have forked the original repository and are maintaining it. We are open to discussing the future.
## [10][Learn Rust from PHP](https://www.reddit.com/r/rust/comments/gm0bsl/learn_rust_from_php/)
- url: https://www.reddit.com/r/rust/comments/gm0bsl/learn_rust_from_php/
---
Hi Rustaceans!

I want to learn something new, I'm currently a PHP developer (I use HTML/CSS/JS and MySQL), I'm hesitated between Golang and Rust and I've tested Golang but I don't know if it's for me. So I want to try Rust before.

After reading the book, what can I build (something simple for begining)? Like a cli/tui app.  


Maybe Golang is the best for me but Rust has more interesting stuff in the language, so I don't know.
## [11][Common pitfalls for C# OOP developers](https://www.reddit.com/r/rust/comments/gm0386/common_pitfalls_for_c_oop_developers/)
- url: https://www.reddit.com/r/rust/comments/gm0386/common_pitfalls_for_c_oop_developers/
---
Hi, I'm a long-time C# developer and I want to become a rustacean.

I'm excited to learn about Rust but I'm still worried what pitfalls I will need to overcome.  
For example the ownership paradigm sounds neat but I can see myself not using it correctly because I became lazy during the years developing in a garbage collected world.

Any other .Net developers have some general tips, experience or even caveats during their move to Rust?
## [12][A crate for indexing and quering json objects](https://www.reddit.com/r/rust/comments/glz519/a_crate_for_indexing_and_quering_json_objects/)
- url: https://github.com/mambisi/json_indexer
---

