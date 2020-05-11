# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (19/2020)!](https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 337](https://www.reddit.com/r/rust/comments/geagy0/this_week_in_rust_337/)
- url: https://this-week-in-rust.org/blog/2020/05/05/this-week-in-rust-337/
---

## [3][Make LLVM fast again](https://www.reddit.com/r/rust/comments/ghfnqf/make_llvm_fast_again/)
- url: https://nikic.github.io/2020/05/10/Make-LLVM-fast-again.html
---

## [4][Writing A Wayland Compositor In Rust](https://www.reddit.com/r/rust/comments/gh6llk/writing_a_wayland_compositor_in_rust/)
- url: https://wiki.alopex.li/WritingAWaylandCompositorInRust
---

## [5][Yak shaving conditional compilation in Rust](https://www.reddit.com/r/rust/comments/ghkunu/yak_shaving_conditional_compilation_in_rust/)
- url: https://bitshifter.github.io/2020/05/07/conditional-compilation-in-rust/
---

## [6][Comparing the safety and performance of operations in Rust and C++](https://www.reddit.com/r/rust/comments/gh84k2/comparing_the_safety_and_performance_of/)
- url: https://www.viva64.com/en/b/0733/
---

## [7][rust-analyzer changelog #24](https://www.reddit.com/r/rust/comments/ghn4iw/rustanalyzer_changelog_24/)
- url: https://rust-analyzer.github.io/thisweek/2020/05/11/changelog-24.html
---

## [8][LLVM 10 has performance regressions with Rust](https://www.reddit.com/r/rust/comments/gh3thc/llvm_10_has_performance_regressions_with_rust/)
- url: https://lists.llvm.org/pipermail/llvm-dev/2020-May/141482.html
---

## [9][Can I save this clone?](https://www.reddit.com/r/rust/comments/ghirn7/can_i_save_this_clone/)
- url: https://www.reddit.com/r/rust/comments/ghirn7/can_i_save_this_clone/
---
Bellow is the code. Just try to write a game server while learning rust. My intention is to save all Players in the HashMap, but I am confused a bit on the life time if I want to avoid the copy of id. Should I change the `HashMap` key to be a reference or the `id` in struct `Player` to be a reference? 

My own though is if I change the `id` in `Player` to be reference, All other structs owning player will also need the explicitly declare the lifetime.
    
    
```
    use std::collections::HashMap;
    
    #[derive(Debug, Eq, PartialEq, Clone, Ord, PartialOrd, Hash)]
    pub struct Player {
        id: String,
        password: String,
    }
    
    struct Hall {
        players: HashMap&lt;String, Player&gt;,
        rooms: List&lt;Room&gt;,
    }
    
    struct Room {
        players: HashMap&lt;String, Player&gt;,
    }
    
    impl Hall{
    
        fn login(&amp;mut self,  id: String, password: String) {
            let id_copy = id.clone(); // Can I save this clone? 
            let player = Player{
                id: id,
                password: password,
            };
            self.players.insert(id_copy, player);
        }
    }
```
## [10][Enum vs structs implementing a trait](https://www.reddit.com/r/rust/comments/ghk31y/enum_vs_structs_implementing_a_trait/)
- url: https://www.reddit.com/r/rust/comments/ghk31y/enum_vs_structs_implementing_a_trait/
---
Hello! Can someone explain me, why some crates use empty structs instead of plain enums? Like this one: https://docs.rs/uom/0.27.0/uom/si/temperature_interval/index.html
## [11][Criticisms of rust](https://www.reddit.com/r/rust/comments/ggyo51/criticisms_of_rust/)
- url: https://www.reddit.com/r/rust/comments/ggyo51/criticisms_of_rust/
---
Rust is on my list of things to try and I have read mostly only good things about it. I want to know about downsides also, before trying. Since I have heard learning curve will be steep.

compared to other languages like Go, I don't know how much adoption rust has. But apparently languages like go and swift get quite a lot of criticism. in fact there is a [github repo](https://github.com/ksimka/go-is-not-good) to collect criticisms of Go.

Are there well written (read: not emotional rant) criticisms of rust language? Collecting them might be a benefit to rust community as well.
## [12][Rust concurrency: the single-writer principle](https://www.reddit.com/r/rust/comments/ghilsf/rust_concurrency_the_singlewriter_principle/)
- url: https://www.reddit.com/r/rust/comments/ghilsf/rust_concurrency_the_singlewriter_principle/
---
An example of applying the single-writer principle to a concurrent Rust system.

[https://medium.com/@polyglot\_factotum/rust-concurrency-the-single-writer-principle-applied-aada2cdc6fb0?source=friends\_link&amp;sk=cafc8dcf8babf4ec95b1b62ccde7e54b](https://medium.com/@polyglot_factotum/rust-concurrency-the-single-writer-principle-applied-aada2cdc6fb0?source=friends_link&amp;sk=cafc8dcf8babf4ec95b1b62ccde7e54b)
