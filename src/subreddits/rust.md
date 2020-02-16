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

## [3][Specs entity component system v0.16.0](https://www.reddit.com/r/rust/comments/f4obel/specs_entity_component_system_v0160/)
- url: https://github.com/amethyst/specs/releases/tag/0.16.0
---

## [4][Why is swc fast?](https://www.reddit.com/r/rust/comments/f4nddx/why_is_swc_fast/)
- url: https://swc-project.github.io/blog/2020/02/16/why-is-swc-fast
---

## [5][How does Rust handle multiple template instantiations?](https://www.reddit.com/r/rust/comments/f4p5qq/how_does_rust_handle_multiple_template/)
- url: https://www.reddit.com/r/rust/comments/f4p5qq/how_does_rust_handle_multiple_template/
---
C++ has `extern template` which can be used to avoid the instantiation of the same template in every compilation unit. How does Rust handle this issue?
## [6][The Battle for Rust Knowledge Supremacy](https://www.reddit.com/r/rust/comments/f4gouu/the_battle_for_rust_knowledge_supremacy/)
- url: https://rustbattle.net/battle/straight-finch-8-e4f4
---

## [7][muso: music sorter](https://www.reddit.com/r/rust/comments/f4k6qc/muso_music_sorter/)
- url: https://github.com/quebin31/muso
---

## [8][Rust @ embedded world 2020?](https://www.reddit.com/r/rust/comments/f4oui2/rust_embedded_world_2020/)
- url: https://www.reddit.com/r/rust/comments/f4oui2/rust_embedded_world_2020/
---
Last time I went to the embedded world in 2018, there was a small but nice Rust presence. This time I cannot find any entry in the exhibitors list, FWIW. Is anybody coming?
## [9][Why Rust, Or: A Trip Report from my Satori with Rust and Functional Programming](https://www.reddit.com/r/rust/comments/f4gmzm/why_rust_or_a_trip_report_from_my_satori_with/)
- url: https://christine.website/blog/why-rust-2020-02-15
---

## [10][Help required to write correct Rust program](https://www.reddit.com/r/rust/comments/f4ojsz/help_required_to_write_correct_rust_program/)
- url: https://www.reddit.com/r/rust/comments/f4ojsz/help_required_to_write_correct_rust_program/
---
Hello RustDevs,

Sorry for asking such bad question but even after reading [https://rust-unofficial.github.io/too-many-lists](https://rust-unofficial.github.io/too-many-lists) I am not able to build a correct mental model for programming in Rust :(  I tried writing code logically but I am not able to get pass from borrow checker. That is very demotivating.

Can someone please help me complete my small code? also please explain the reasons for steps.

if my layout is incorrect for what I am trying to achieve with this code please explain that too.

    use std::collections::hash_map::DefaultHasher;
    use std::hash::{Hash, Hasher};
    
    type Link&lt;T&gt; = Option&lt;Box&lt;Node&lt;T&gt;&gt;&gt;;
    
    struct Node&lt;T&gt; {
        item: T,
        key: u64,
        next: Link&lt;T&gt;,
    }
    
    struct CoarseList&lt;T&gt; {
        head: Link&lt;T&gt;,
    }
    
    impl&lt;T: std::hash::Hash&gt; CoarseList&lt;T&gt; {
        pub fn new() -&gt; Self {
            CoarseList { head: None}
        }
    
        pub fn add(&amp;mut self, item: T) -&gt; bool {
            let mut hasher = DefaultHasher::new();
            item.hash(&amp;mut hasher);
            let key =  hasher.finish();
    		
            let mut new_node = Box::new(Node {
                item: item,
                key: key,
                next: None,
            });
    		
            if self.head.is_none() {
                self.head = Some(new_node);
            } else {
    			// iterate over list while node.key &lt; key (newly created key)
    			// if key is already inserted in list return false i.e can't not insert duplicate
    			// else once found correct sorted position add new_node and update cur, prev etc properly.
            }
            true
        }
    }

Thanks in advance!
## [11][Bad performance for certain applications of scoped threads](https://www.reddit.com/r/rust/comments/f4i1m6/bad_performance_for_certain_applications_of/)
- url: https://www.reddit.com/r/rust/comments/f4i1m6/bad_performance_for_certain_applications_of/
---
I've been investigating a performance issue in my computer game project.

What I learnt prompted me to write up the following blog post:

[https://github.com/avl/multi\_iterate\_mut/blob/master/README.md](https://github.com/avl/multi_iterate_mut/blob/master/README.md)

TLDR is that scoped threads (using rayon- or scoped\_threadpool-crates) used wrong can cause huge performance problems if the working set fits in cache but data is artificially ping-ponged between OS-threads because task-scheduling to worker threads is non-deterministic.

In my example running multi-threaded using rayon::scope caused a 3X slowdown over single-threaded.

But with a new, experimental, scoped-thread implementation a 5x speedup could be achieved instead.
## [12][Can there be memory leaks in pure Rust? If yes, how?](https://www.reddit.com/r/rust/comments/f4m4sw/can_there_be_memory_leaks_in_pure_rust_if_yes_how/)
- url: https://www.reddit.com/r/rust/comments/f4m4sw/can_there_be_memory_leaks_in_pure_rust_if_yes_how/
---

