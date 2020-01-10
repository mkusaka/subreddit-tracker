# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (2/2020)!](https://www.reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here_22020/)
- url: https://www.reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here_22020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/ehk18j/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 320](https://www.reddit.com/r/rust/comments/em9hct/this_week_in_rust_320/)
- url: https://this-week-in-rust.org/blog/2020/01/07/this-week-in-rust-320/
---

## [3][Proof of Concept: Rust code in Unreal Engine](https://www.reddit.com/r/rust/comments/emp0su/proof_of_concept_rust_code_in_unreal_engine/)
- url: https://ejmahler.github.io/rust_in_unreal/
---

## [4][Way Cooler Post Mortem](https://www.reddit.com/r/rust/comments/emjwwp/way_cooler_post_mortem/)
- url: http://way-cooler.org/blog/2020/01/09/way-cooler-post-mortem.html
---

## [5][Towards a Rust foundation](https://www.reddit.com/r/rust/comments/emd0cn/towards_a_rust_foundation/)
- url: http://smallcultfollowing.com/babysteps/blog/2020/01/09/towards-a-rust-foundation/
---

## [6][How Arc works in Rust](https://www.reddit.com/r/rust/comments/emosl0/how_arc_works_in_rust/)
- url: https://link.medium.com/mU0aRn9d82
---

## [7][The panic messages now pointing to the location where they were called, rather than core's internals](https://www.reddit.com/r/rust/comments/em7gav/the_panic_messages_now_pointing_to_the_location/)
- url: https://github.com/rust-lang/rust/pull/67887
---

## [8][Why I like programming for #![no_std] in Rust (even without embedded)](https://www.reddit.com/r/rust/comments/emq98r/why_i_like_programming_for_no_std_in_rust_even/)
- url: https://djugei.github.io/no-std-motivation/
---

## [9][BLAKE3 - Evolution of BLAKE2](https://www.reddit.com/r/rust/comments/emf15c/blake3_evolution_of_blake2/)
- url: https://github.com/BLAKE3-team/BLAKE3-specs
---

## [10][The Library team is moving from IRC to Zulip - announcements](https://www.reddit.com/r/rust/comments/emig2y/the_library_team_is_moving_from_irc_to_zulip/)
- url: https://internals.rust-lang.org/t/the-library-team-is-moving-from-irc-to-zulip/11598
---

## [11][`enum_variant_type` -- proc macro derive to generate structs from enums and convert between them](https://www.reddit.com/r/rust/comments/emnjmm/enum_variant_type_proc_macro_derive_to_generate/)
- url: https://users.rust-lang.org/t/enum-variant-type-proc-macro-derive-to-generate-structs-from-enums-and-convert-between-them/36737
---

## [12][how mutex poisoning is detected in rust by mutex](https://www.reddit.com/r/rust/comments/eml4l8/how_mutex_poisoning_is_detected_in_rust_by_mutex/)
- url: https://www.reddit.com/r/rust/comments/eml4l8/how_mutex_poisoning_is_detected_in_rust_by_mutex/
---
    let mut num = m.lock().unwrap();

&gt;The call to lock would fail if another thread holding the lock panicked. In that case, no one would ever be able to get the lock, so we’ve chosen to unwrap and have this thread panic if we’re in that situation.

However, if another thread holding the mutex guard panics, then objects in that thread should be dropped. **Which should then release the mutex guard**. Then mutex implementation enters a 'poisoned' trait. 

Does this mean there is someway to detect drop is being called as a result of a panic? 

 [https://doc.rust-lang.org/std/sync/struct.MutexGuard.html](https://doc.rust-lang.org/std/sync/struct.MutexGuard.html)  

is this implemented in the  sys\_common::poison?
