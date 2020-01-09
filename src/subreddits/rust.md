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
## [2][What’s everyone working on this week (2/2020)?](https://www.reddit.com/r/rust/comments/ekpr6w/whats_everyone_working_on_this_week_22020/)
- url: https://www.reddit.com/r/rust/comments/ekpr6w/whats_everyone_working_on_this_week_22020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-2-2020/36546?u=llogiq)!
## [3][The panic messages now pointing to the location where they were called, rather than core's internals](https://www.reddit.com/r/rust/comments/em7gav/the_panic_messages_now_pointing_to_the_location/)
- url: https://github.com/rust-lang/rust/pull/67887
---

## [4][FFI like it's 2020: Announcing *safe* FFI for Rust &lt;-&gt; C++](https://www.reddit.com/r/rust/comments/elvfyn/ffi_like_its_2020_announcing_safe_ffi_for_rust_c/)
- url: https://github.com/dtolnay/cxx
---

## [5][Rust reimplemented xv6 buddy allocator](https://www.reddit.com/r/rust/comments/em6xhz/rust_reimplemented_xv6_buddy_allocator/)
- url: https://github.com/jjyr/buddy-alloc
---

## [6][This Week in Rust 320](https://www.reddit.com/r/rust/comments/em9hct/this_week_in_rust_320/)
- url: https://this-week-in-rust.org/blog/2020/01/07/this-week-in-rust-320/
---

## [7][Intimidated by advanced Rust](https://www.reddit.com/r/rust/comments/elzas3/intimidated_by_advanced_rust/)
- url: https://www.reddit.com/r/rust/comments/elzas3/intimidated_by_advanced_rust/
---
Hi,

&amp;#x200B;

I was wondering whether some of you felt similarly. I really enjoy learning Rust. I use it pretty much exclusively for my personal projects and I feel my understanding of the language and its idioms increases with each project. That was until I started working on an implementation of a linked list to better understand RefCell and friends. To this end, I read [https://rust-unofficial.github.io/too-many-lists/](https://rust-unofficial.github.io/too-many-lists/), and at chapter 'A Bad Safe Deque', I got to a point (Iter and IterMut) where I feel really discouraged. I'm not sure I'll ever get to a level where I can reason like the author does in 'Iteration'. Holy moly, the compiler errors are dense and intimidating. Has anybody else been there and how did you overcome it? Just table more advanced topics and accept that this process takes time to fully master?

Not trying to cry on your guys' shoulders but looking for a little guidance...

&amp;#x200B;

Thanks,
## [8][Update of our Rust Boilerplate server with GraphQL (Async and Actix-web 2)](https://www.reddit.com/r/rust/comments/em8bx9/update_of_our_rust_boilerplate_server_with/)
- url: https://www.reddit.com/r/rust/comments/em8bx9/update_of_our_rust_boilerplate_server_with/
---
Repo: [https://github.com/clifinger/canduma](https://github.com/clifinger/canduma)

Changelist:

* use async/await
* use actix-web version 2
* use anyhow + thiserror in place of failure
* structopt
## [9][Manning's Rust book and video course are half off today](https://www.reddit.com/r/rust/comments/em5ls8/mannings_rust_book_and_video_course_are_half_off/)
- url: https://i.redd.it/h6vqe78p1p941.png
---

## [10][Help debug a performance degradation in parking_lot that only happen some processors](https://www.reddit.com/r/rust/comments/em0ru7/help_debug_a_performance_degradation_in_parking/)
- url: https://github.com/Amanieu/parking_lot/issues/201
---

## [11][Videos from Rust Belt Rust 2019 are now available!](https://www.reddit.com/r/rust/comments/elw26j/videos_from_rust_belt_rust_2019_are_now_available/)
- url: https://www.youtube.com/playlist?list=PLgC1L0fKd7UkVwjVlOySfMnn80Qs5TOLb
---

## [12][Rust 2020: a conference in Latin America](https://www.reddit.com/r/rust/comments/em0ru8/rust_2020_a_conference_in_latin_america/)
- url: https://www.reddit.com/r/rust/comments/em0ru8/rust_2020_a_conference_in_latin_america/
---
Hello Rust community!

Rust LATAM is the conference working on growing the interest around Rust in Latin America; the next iteration of the conference will happen in **Mexico City on May, 22nd-23rd 2020** 🇲🇽

Rust LATAM is not only a conference, it's a community event and that means that is made by **YOU** all participating and spreading the word 💪

The Rust LATAM team has opened the "Call for Papers" for the next LATAM2020 conference. We are accepting proposals for talks (or, if you have a ✨ cool idea ✨ for workshop, feel free to submit that, too!). So, get involved and get your chance to show something about your project, your experiments, a new development or about an aspect of Rust that improved your life as a developer (or not as much as you expected!).

Do you have in mind a talk not strictly technical? That's great, we'd love to hear from you, too!

Submit your proposals at [https://cfp.rustlatam.org](https://cfp.rustlatam.org)

Are you a company working with Rust? Do you want to get in contact and support the conference? Contact us for a sponsorship opportunity at [sponsors@rustlatam.org](mailto:sponsors@rustlatam.org)

The conference is made by and for the community, so don't miss this opportunity, save the date and take part to Rust LATAM 2020: help us make it happen! 🎉

Check out our website [https://www.rustlatam.org](https://www.rustlatam.org), it will be updated as we will be getting closer to the conference.

Follow us on Twitter @ RustLatamConf or join our Discord channel at: [https://discord.gg/b6MBzy](https://discord.gg/b6MBzy)

See you there!
