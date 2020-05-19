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
## [3][Rocket can be compiled on stable Rust 1.45, last blocker has been solved](https://www.reddit.com/r/rust/comments/gmkpsg/rocket_can_be_compiled_on_stable_rust_145_last/)
- url: https://github.com/SergioBenitez/Rocket/issues/19#issuecomment-630650328
---

## [4][A Future for Rust Debugging](https://www.reddit.com/r/rust/comments/gmhdne/a_future_for_rust_debugging/)
- url: https://nbaksalyar.github.io/2020/05/19/rust-debug.html
---

## [5][IntelliJ Rust Changelog #122](https://www.reddit.com/r/rust/comments/gmmumk/intellij_rust_changelog_122/)
- url: https://intellij-rust.github.io/2020/05/19/changelog-122.html
---

## [6][Orb.Farm, a virtual aquatic ecosystem, implemented in Rust](https://www.reddit.com/r/rust/comments/gm4whz/orbfarm_a_virtual_aquatic_ecosystem_implemented/)
- url: https://orb.farm
---

## [7][Rocket will compile with stable Rust 1.45](https://www.reddit.com/r/rust/comments/gmlbg4/rocket_will_compile_with_stable_rust_145/)
- url: https://github.com/SergioBenitez/Rocket/issues/19
---

## [8][rust-analyzer: next few years](https://www.reddit.com/r/rust/comments/gm1nr6/rustanalyzer_next_few_years/)
- url: https://rust-analyzer.github.io/blog/2020/05/18/next-few-years.html
---

## [9][Casbin-RS: An authorization library that supports access control models like ACL, RBAC, ABAC in Rust](https://www.reddit.com/r/rust/comments/gmfdws/casbinrs_an_authorization_library_that_supports/)
- url: https://github.com/casbin/casbin-rs
---

## [10][rust-analyzer changelog #25](https://www.reddit.com/r/rust/comments/gm159f/rustanalyzer_changelog_25/)
- url: https://rust-analyzer.github.io/thisweek/2020/05/18/changelog-25.html
---

## [11][Fwd:AD, a forward auto-differentiation crate](https://www.reddit.com/r/rust/comments/gmbpv5/fwdad_a_forward_autodifferentiation_crate/)
- url: https://www.reddit.com/r/rust/comments/gmbpv5/fwdad_a_forward_autodifferentiation_crate/
---
Hi everyone,

[**Fwd:AD**](https://lib.rs/crates/fwd_ad) is a Rust crate to perform [forward auto-differentiation](https://en.wikipedia.org/wiki/Automatic_differentiation#Automatic_differentiation_using_dual_numbers), with a focus on empowering its users to manage memory location and minimize copying. It is made with the goal of being useful and used, so documentation and examples are considered as important as code during development. Its key selling-points are:

1. **Clone-free** by default. Fwd:AD will never clone memory in its functions and `std::ops` implementations, leveraging Rust's ownership system to ensure correctness memory-wise, and leaving it up to the user to be explicit as to when cloning should happen.
2. **Automatic cloning** on demand. If passed the `implicit-clone` feature, Fwd:AD will implicitly clone when needed. Deciding whether to clone or not is entirely done via the type-system, and hence at compile time.
3. **Generic in memory location**: Fwd:AD's structs are generic over a container type, allowing them to be backed by any container of your choice: `Vec` to rely on the heap, arrays if you're more of a stack-person, or other. For example, it can be used with `&amp;mut [f64]` to allow an FFI API that won't need to copy memory at its frontier.

I've been working on it for the last months and I think it is mature enough to be shared.

I am very eager to get feedback and to see how it could be used by the community so please share any comment or question you might have.

* **lib.rs**: [https://lib.rs/crates/fwd\_ad](https://lib.rs/crates/fwd_ad)
* **gitlab**: [https://gitlab.inria.fr/InBio/Public/fwd\_ad](https://gitlab.inria.fr/InBio/Public/fwd_ad)
* **crates.io**: [https://crates.io/crates/fwd\_ad](https://crates.io/crates/fwd_ad)
* **docs.rs**: [https://docs.rs/fwd\_ad/0.2.0/fwd\_ad/](https://docs.rs/fwd_ad/0.2.0/fwd_ad/)

Thanks to all the Rust community for helping me during the development, you made every step of it enjoyable.
## [12][100k r/rust subscribers – how time flies!](https://www.reddit.com/r/rust/comments/gm3a45/100k_rrust_subscribers_how_time_flies/)
- url: https://www.reddit.com/r/rust/comments/gm3a45/100k_rrust_subscribers_how_time_flies/
---
Feels like yesterday we have reached:

[30k on Oct 31 2017](https://www.reddit.com/r/rust/comments/79xw22/30k/)   
[50k on Jan 15 2019](https://www.reddit.com/r/rust/comments/agc1vq/congratulation_with_50k_rustaceans_on_reddit/)     
[80k on Nov 20 2019](https://www.reddit.com/r/rust/comments/dyx7qd/rrust_hit_80k_subscribers_yesterday/)     
[100k on May 18 2020](https://www.reddit.com/r/rust/comments/gm3a45/100k_rrust_subscribers_how_time_flies/)    

feels good to also look at some [graphs](https://subredditstats.com/r/rust)     

Thanks to all Rustaceans out there making this sub on of the bests on reddit – and especially a big thank you to all the mods making such a great work. I think they deserve some ❤️ and we tend to forget about the people behind the scenes making such a great experience. Let this be also a big thank you because you mods are a big reason to the success of this sub, because this is hard work!
