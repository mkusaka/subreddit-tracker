# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (27/2020)!](https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 345](https://www.reddit.com/r/rust/comments/hisn3e/this_week_in_rust_345/)
- url: https://this-week-in-rust.org/blog/2020/06/30/this-week-in-rust-345/
---

## [3][Linus Torvalds: "the kernel team is looking at having interfaces to do [drivers], for example, in Rust... I'm convinced it's going to happen."](https://www.reddit.com/r/rust/comments/hiyfhq/linus_torvalds_the_kernel_team_is_looking_at/)
- url: https://www.theregister.com/2020/06/30/hard_to_find_linux_maintainers_says_torvalds/
---

## [4][Convincing my company to consider Rust in the tech plan](https://www.reddit.com/r/rust/comments/hj66pc/convincing_my_company_to_consider_rust_in_the/)
- url: https://www.reddit.com/r/rust/comments/hj66pc/convincing_my_company_to_consider_rust_in_the/
---
My company is going to open a debate period in order to update its tech plan in a few weeks. During the last period (about 2 years ago) I pushed to get Rust into the tech plan as a substitute for C++ (which is admittedly there for legacy reasons) but the proposal didn't get traction.

A lot of good things have happened in the Rust community in the last two years, and I would like to try this again, arguing once again that we should replace C++ with Rust (adding a new tech implies removing one from the tech plan in my company, and going after C++ seems like a good idea since new projects aren't using it). These are the arguments that I think can use to convince people to take a second look at Rust:

- Greatly improved ergonomics with the addition of NLL and async/await

- Huge adoption boost in the past years: Microsoft, Amazon, Facebook and Google amongst other

- Increased security and robustness when compared to C++

- One of the best languages for WASM

- Great support for interacting with Kubernetes

- Faster and leaner for microservices when compared to JVM-based services

Any other major points I might be missing?
## [5][7 Things I Learned From Porting a C Crypto Library to Rust](https://www.reddit.com/r/rust/comments/hj9pvu/7_things_i_learned_from_porting_a_c_crypto/)
- url: https://sharpend.io/7-things-I-learned-from-porting-a-c-crypto-library-to-rust/
---

## [6][griddle: A HashMap variant that spreads resize load across inserts.](https://www.reddit.com/r/rust/comments/hiukis/griddle_a_hashmap_variant_that_spreads_resize/)
- url: https://github.com/jonhoo/griddle
---

## [7][After a positive response on my friend's work improving Risc-V support in Rust, he's written a more technical follow up article on his work. Hope you enjoy :)](https://www.reddit.com/r/rust/comments/hio31k/after_a_positive_response_on_my_friends_work/)
- url: https://www.codethink.co.uk/articles/2020/fixing-rusts-test-suite-on-risc-v/
---

## [8][Firefox 78 stable release enables WebRender for Windows users with Intel GPUs](https://www.reddit.com/r/rust/comments/hiqukg/firefox_78_stable_release_enables_webrender_for/)
- url: https://www.mozilla.org/en-US/firefox/78.0/releasenotes/
---

## [9][Implementing WebSockets in Rust](https://www.reddit.com/r/rust/comments/hj2sxu/implementing_websockets_in_rust/)
- url: http://subhojit777.in/implementing-websockets-in-rust
---

## [10][diffy: tools for finding and manipulating differences between files](https://www.reddit.com/r/rust/comments/hiwgrn/diffy_tools_for_finding_and_manipulating/)
- url: https://www.reddit.com/r/rust/comments/hiwgrn/diffy_tools_for_finding_and_manipulating/
---
Before I became heavily interested in Rust I was a contributor to git. In fact, I was first introduced to Rust when I asked a fellow git contributor what language they would implement git in if they had the opportunity to start from scratch. Anyway, since then I've always been interested in source control and wanted to see parts of git's functionality implemented in Rust. Over the last few years i've tinkered with implementing bits and pieces here and there but I finally found some time recently to devote to releasing something. Instead of focusing on functionality specific to git I thought I'd work on something more common all VCS's, manipulating differences between different versions of a file.

[diffy](https://docs.rs/diffy) is a library intended to be a collection of tools used to find and manipulate differences between line-based files inspired by [LibXDiff](http://www.xmailserver.org/xdiff-lib.html) and [GNU Diffutils](https://www.gnu.org/software/diffutils/). The initial release provides basic functionality for producing diffs or patches between two files and displaying them in the unified diff format as well as the ability to merge two files given a common ancestor.

In the future I intend to also add support for parsing unified diffs into patches as well as applying a patch. For right now the library only operates on utf8 strings but it shouldn't be too much work to also handle non-utf8 input (which would be needed if this was to actually be used in a VCS).
## [11][Question about Rust AST](https://www.reddit.com/r/rust/comments/hj720a/question_about_rust_ast/)
- url: https://www.reddit.com/r/rust/comments/hj720a/question_about_rust_ast/
---
Hello,

Which part of the Rust tooling can, from a given piece of Rust code (whether it's a while program, a single function or a single statement/expression) generate a kind of AST that includes the original variable names and detailed info about their state in memory (size, stack vs heap etc.) ? I'm investigating the possibility of representing memory "usage" from Rust code. I know there are various optimisations passes, but I'm interested in the code's intent.

Thanks in advance.
## [12][Getting started with WebAssembly and Rust](https://www.reddit.com/r/rust/comments/hivf01/getting_started_with_webassembly_and_rust/)
- url: https://blog.logrocket.com/getting-started-with-webassembly-and-rust/
---

