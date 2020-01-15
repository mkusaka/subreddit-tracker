# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (3/2020)!](https://www.reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here_32020/)
- url: https://www.reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here_32020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (3/2020)?](https://www.reddit.com/r/rust/comments/eo6qjo/whats_everyone_working_on_this_week_32020/)
- url: https://www.reddit.com/r/rust/comments/eo6qjo/whats_everyone_working_on_this_week_32020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-3-2020/36867?u=llogiq)!
## [3][Rust jokes needed!](https://www.reddit.com/r/rust/comments/eozzje/rust_jokes_needed/)
- url: https://www.reddit.com/r/rust/comments/eozzje/rust_jokes_needed/
---
Hey everyone! 

I'm looking for the best rust puns/jokes you can manage! My boyfriend loves rust, as well as a few other languages. Due to a couple different things he's having a rough time in life at the moment. I wanted to make something that would give him a laugh and hopefully help him feel a bit better. My plan is to find an amazing and terrible pun/joke and embroider it on a t-shirt for him! 

I am of course not a programmer so I have no idea where else to turn for these. I've enlisted his friends for final approval and to make sure they actually make sense (because I have no idea!) Hopefully he doesn't follow this sub... 

Thank you for your help!
## [4][Announcing v0.2 of SQLx (the async, pure, safe Rust SQL crate featuring compile-time checked queries without a DSL), now with support for Tokio, TLS, and transactions](https://www.reddit.com/r/rust/comments/ep0mgf/announcing_v02_of_sqlx_the_async_pure_safe_rust/)
- url: https://github.com/launchbadge/sqlx/blob/62e1df5ad9454696fc2c75b78c3190f801fcecde/CHANGELOG.md#020---2020-01-15
---

## [5][Announcing Dashmap v3 - Taking concurrent hashmaps to the next level.](https://www.reddit.com/r/rust/comments/eopxq5/announcing_dashmap_v3_taking_concurrent_hashmaps/)
- url: https://www.reddit.com/r/rust/comments/eopxq5/announcing_dashmap_v3_taking_concurrent_hashmaps/
---
Today I am very happy to annouce v3 of my concurrent hashmap implementation [DashMap](https://github.com/xacrimon/dashmap). I did not make an announcement for v2 as it has some internal flaws that required API redesign.

DashMap is a personal project created by my frustration caused by slowdowns due to
the frequently used `RwLock&lt;HashMap&lt;K, V&gt;&gt;` which I see way to often.
It tries to be close to a drop in replacement and expose a similar API.

So, what changes have been made since I released v1?

Most of the map has undergone redesign but here are the important bits.

- Better iterator API

- Entry API which tries to mimick `std::collections::HashMap`

- Low level raw shard access support

- Massively improved performance (see below)

- Serde support

- Common iterator traits implemented

- Rustified and cleaner API

- More relaxed trait bounds (v1 was much stricter than needed)

So, what about performance you ask?

Well to say that it is better than v1 is a mild understatement.
I've been hard at work squeezing maximum performance. And it has paid off quite well.

Below are some benchmarks I've ran on a i7-6700 with mild heat issues paired with 16 GiB of fancy Corsair memory.

Keep in mind these benchmarks are far from perfect and I'd like more. This is something that I'm not super good at and any help is appreciated with open arms.

https://raw.githubusercontent.com/xacrimon/dashmap/master/assets/bench-insert.svg?sanitize=true

https://raw.githubusercontent.com/xacrimon/dashmap/master/assets/bench-get.svg?sanitize=true

[Google Doc](https://docs.google.com/spreadsheets/d/1q2VR_rMZRzG7YO0ef6V0jMA6hAdkafh_wI8xvY_51fk/edit?usp=sharing)
## [6][Tonic: 0.1 has arrived!](https://www.reddit.com/r/rust/comments/eomgrr/tonic_01_has_arrived/)
- url: https://luciofran.co/tonic-0-1-release/
---

## [7][This Week in Veloren #50: Skeletons and Scrolling Combat Text](https://www.reddit.com/r/rust/comments/eowaxw/this_week_in_veloren_50_skeletons_and_scrolling/)
- url: https://veloren.net/devblog-50/
---

## [8][PoC libs for embedded programming based on HAL, DHT sensors driver and 8segments LED driver for TM16xx MCU](https://www.reddit.com/r/rust/comments/eoza8a/poc_libs_for_embedded_programming_based_on_hal/)
- url: https://www.reddit.com/r/rust/comments/eoza8a/poc_libs_for_embedded_programming_based_on_hal/
---
This one is for reading DHT11, DHT21, DHT22: [https://github.com/rustrum/dht-hal-drv](https://github.com/rustrum/dht-hal-drv)

This one is to output something on 8 segment LED displays based on TM1637 or TM1638: [https://github.com/rustrum/tmledkey-hal-drv](https://github.com/rustrum/tmledkey-hal-drv)

&amp;#x200B;

You could combine these libraries together and create DIY thermometer/clocks on STM32, RaspberryPi or any other platform where is proper HAL implementation available.
## [9][boxing-arena crate - Simple reuse of Box allocation](https://www.reddit.com/r/rust/comments/eozyq5/boxingarena_crate_simple_reuse_of_box_allocation/)
- url: https://github.com/da-x/boxing-arena
---

## [10][How is ownership independent of heap vs stack?](https://www.reddit.com/r/rust/comments/ep0s60/how_is_ownership_independent_of_heap_vs_stack/)
- url: https://www.reddit.com/r/rust/comments/ep0s60/how_is_ownership_independent_of_heap_vs_stack/
---
I came across this post [https://www.reddit.com/r/rust/comments/8kt5eq/ownership\_vs\_the\_stack\_and\_heap/](https://www.reddit.com/r/rust/comments/8kt5eq/ownership_vs_the_stack_and_heap/)

In particular, there are at least two commenters saying "ownership is separate from stack/heap."

&gt;The semantics of a language are generally *not* tied to implementation details such as Stack and Heap. Or otherwise said, the rules should apply whether or not a value is currently stored on the Stack, on the Heap, or maybe in Read-Only Memory.

&amp;#x200B;

&gt;Ownership and borrowing themselves are entirely separate from stack vs heap

From what I gathered, an essential characteristic of ownership is that it "crosses off" everything created in the current scope from the stack whenever it exits that scope. How would you implement ownership with this cross-everything-off-on-scope-out feature if it wasn't relying on the stack architecture? How is ownership not tied to the concepts of stack/heap?

Or maybe I misunderstood what they meant?
## [11][bita 0.5.0 - New release of the low bandwidth file synchronization tool](https://www.reddit.com/r/rust/comments/eorgg0/bita_050_new_release_of_the_low_bandwidth_file/)
- url: https://github.com/oll3/bita
---

## [12][Response from bjorn3 on "Cranelift backend for Rust"](https://www.reddit.com/r/rust/comments/eojzjr/response_from_bjorn3_on_cranelift_backend_for_rust/)
- url: https://github.com/bjorn3/rustc_codegen_cranelift/issues/863
---

