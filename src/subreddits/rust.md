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

## [3][Ownership of the standard library implementation | Inside Rust Blog](https://www.reddit.com/r/rust/comments/hjuh2f/ownership_of_the_standard_library_implementation/)
- url: https://blog.rust-lang.org/inside-rust/2020/07/02/Ownership-Std-Implementation.html
---

## [4][Cargo and rustc running natively on an iPhone 11](https://www.reddit.com/r/rust/comments/hjk4vw/cargo_and_rustc_running_natively_on_an_iphone_11/)
- url: https://i.redd.it/hkuyjj5tib851.png
---

## [5][Google: "Over time we will continue to invest in Rust and see which [Android] system components are better off being written in Rust"](https://www.reddit.com/r/rust/comments/hjejsp/google_over_time_we_will_continue_to_invest_in/)
- url: https://www.reddit.com/r/rust/comments/hjejsp/google_over_time_we_will_continue_to_invest_in/
---
Sudhi Herle, Head of Android Platform Security, says in yesterday's Android Developer weekly video:

&gt; Over time we will continue to invest in Rust and see which system components are better off being written in Rust. We believe Rust will end up fundamentally making the platform safe for all of our users.

https://www.youtube.com/watch?v=MNkFSCRUk6E&amp;feature=youtu.be&amp;t=727

[Kudos to /u/devsquid for finding this!](https://www.reddit.com/r/rust/comments/hisn3e/this_week_in_rust_345/fwjvy47/)
## [6][Transpiling A Kernel Module to Rust: The Good, the Bad and the Ugly](https://www.reddit.com/r/rust/comments/hjkuwf/transpiling_a_kernel_module_to_rust_the_good_the/)
- url: https://immunant.com/blog/2020/06/kernel_modules/
---

## [7][`dua 2.8` (disk usage analyzer) is now ~430kb smaller and compiles about a minute more quickly!](https://www.reddit.com/r/rust/comments/hjn64e/dua_28_disk_usage_analyzer_is_now_430kb_smaller/)
- url: https://github.com/byron/dua-cli#installation
---

## [8][DUSTR: Call rust from dart seamlessly](https://www.reddit.com/r/rust/comments/hjlses/dustr_call_rust_from_dart_seamlessly/)
- url: https://github.com/mqnfred/dustr
---

## [9][Appreciation for The Book and Rustlings](https://www.reddit.com/r/rust/comments/hjecij/appreciation_for_the_book_and_rustlings/)
- url: https://www.reddit.com/r/rust/comments/hjecij/appreciation_for_the_book_and_rustlings/
---
Not 100% sure if this is allowed, but I wanted to make a post to show a little appreciation.

Last week I worked through The Book alongside Rustlings, and I must say it was truly a pleasure. I've paid for textbooks that don't hold a candle to The Book, and having all of the Rustlings exercises freely available is incredible, and just adds to the experience. Much better experience than MATLAB or Python.

So thanks again; those of you that helped contribute, your effort does not go unappreciated. I look forward to being a part of the community and using the language.

Edit: damn phone typing
## [10][v0.10 of Gleam, a statically typed language written in Rust for the Erlang VM, is out](https://www.reddit.com/r/rust/comments/hjhr45/v010_of_gleam_a_statically_typed_language_written/)
- url: https://lpil.uk/blog/gleam-v0.10-released/
---

## [11][7 Things I Learned From Porting a C Crypto Library to Rust](https://www.reddit.com/r/rust/comments/hj9pvu/7_things_i_learned_from_porting_a_c_crypto/)
- url: https://sharpend.io/7-things-I-learned-from-porting-a-c-crypto-library-to-rust/
---

## [12][Enforcing value based constraints in the type system?](https://www.reddit.com/r/rust/comments/hjsler/enforcing_value_based_constraints_in_the_type/)
- url: https://www.reddit.com/r/rust/comments/hjsler/enforcing_value_based_constraints_in_the_type/
---
I have a question about the type system. One thing I really like about Rust is that it can be used to make certain things impossible (like in a state machine you can enforce certain transitions by making states types, and only allowing going from one type to another).

I have a similar situation. It's a bit complicated, so I'll try to get an analogy going.

Let's say I have cows and pigs. Pigs can breed with other pigs, but not with cows, and vice versa. So I would have something like:

    impl Pig {
        pub fn breed(other: &amp;Pig) -&gt; Pig {
             // ...
        }
    }

And similar for the cow.

However, I have two or more (potentially unlimited) cow farms! The external library for animals in this case has this implemented as an enum:

    enum Animal {
        Pig,
        Cow(usize)
    }

with the `usize` indicating which farm is used.

I now want to build a breeding wrapper around this enum like I did for pigs in my first code example. I want to be have pigs breed with pigs and cows with cows (on the same farm). 
Can I enforce that only cows from the same farm can breed at compile time, or will this have to be a runtime check?
