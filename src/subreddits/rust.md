# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (36/2020)!](https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (36/2020)?](https://www.reddit.com/r/rust/comments/ijvxwz/whats_everyone_working_on_this_week_362020/)
- url: https://www.reddit.com/r/rust/comments/ijvxwz/whats_everyone_working_on_this_week_362020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-36-2020/48100?u=llogiq)!
## [3][Fontdue: The fastest font renderer in the world, written in pure Rust](https://www.reddit.com/r/rust/comments/ilnd41/fontdue_the_fastest_font_renderer_in_the_world/)
- url: https://github.com/mooman219/fontdue/
---

## [4][Just made my own little Ferris :)](https://www.reddit.com/r/rust/comments/ildkmw/just_made_my_own_little_ferris/)
- url: https://i.redd.it/p8vpfjrfcsk51.jpg
---

## [5][Deserializing JSON really fast](https://www.reddit.com/r/rust/comments/ilgfgu/deserializing_json_really_fast/)
- url: https://blog.datalust.co/deserializing-json-really-fast/
---

## [6][100 binaries: /u/wezm's highlights of awesome open source rusty applications](https://www.reddit.com/r/rust/comments/ilnur3/100_binaries_uwezms_highlights_of_awesome_open/)
- url: https://mobile.twitter.com/search?q=from%3A%40wezm%20%23100binaries&amp;f=live
---

## [7][Why are std traits only implemented on arrays of lengths 0..=32?](https://www.reddit.com/r/rust/comments/ilqfrn/why_are_std_traits_only_implemented_on_arrays_of/)
- url: https://www.reddit.com/r/rust/comments/ilqfrn/why_are_std_traits_only_implemented_on_arrays_of/
---
While playing around with a project I am working on to learn Rust, I came across what I thought was a bug, but turned out to be a feature.

In the standard library traits like `IntoIterator` are only implemented on the primitive type `array` for lengths `0..=32`. [See the `LengthAtMost32` trait here.](https://doc.rust-lang.org/src/core/array/mod.rs.html#390)

Why is the `array` type restricted in this way? Is there a technical reason behind it? 

It says [here](https://doc.rust-lang.org/src/core/array/mod.rs.html#395) that this is a temporary step until const generics are stable; what do const generics provide in this particular scenario that means the `array` length limit can be dropped?

Thanks!


Disclaimer: I'm a self taught hobbyist programmer, please forgive my ignorance if the explanation of this is some fairly trivial CompSci 101 thing.
## [8][Flume 0.8, a fast &amp; lightweight MPMC, released: Performance improvements, async support, select API, eventual fairness, Send + Sync + Clone, multiple receivers, no unsafe, timeout/deadline support](https://www.reddit.com/r/rust/comments/il4nq5/flume_08_a_fast_lightweight_mpmc_released/)
- url: https://github.com/zesterer/flume
---

## [9][You can now use clap in shell scripts](https://www.reddit.com/r/rust/comments/ilt8cc/you_can_now_use_clap_in_shell_scripts/)
- url: https://github.com/agnipau/shlap
---

## [10][TIL that Rust allows to modify a vector inside a loop that is iterating on it, if the modification immediately precedes the break](https://www.reddit.com/r/rust/comments/il9g5g/til_that_rust_allows_to_modify_a_vector_inside_a/)
- url: https://www.reddit.com/r/rust/comments/il9g5g/til_that_rust_allows_to_modify_a_vector_inside_a/
---
I put the following code in the playground as an experiment:

    fn replace_with&lt;T: Eq&gt;(list: &amp;mut Vec&lt;T&gt;, old: T, new: T) {
        for (i, value) in list.iter().enumerate() {
            if value == &amp;old {
                list[i] = new;
                break;
            }
        }
    }

To my surprise it compiled.

Is this behavior explained somewhere in the documentation?
## [11][Google's XLS: Build custom hardware using Rust-like syntax](https://www.reddit.com/r/rust/comments/il9mfd/googles_xls_build_custom_hardware_using_rustlike/)
- url: https://google.github.io/xls/
---

## [12][Crate level docs not documented in Rustdoc Book](https://www.reddit.com/r/rust/comments/ilnbq8/crate_level_docs_not_documented_in_rustdoc_book/)
- url: https://www.reddit.com/r/rust/comments/ilnbq8/crate_level_docs_not_documented_in_rustdoc_book/
---
So as usual I was checking out the books in the Rust website's catalogue and I was reading the Rustdoc book. An interesting thing I noticed there was that there was no information about how to write documentation at the crate level. It did not say anything about when you want to write a crate level docs, you should use this syntax

///! 

Pretty unfortunate since they have one of the best documentations out there

I thought to write an issue or pull request but I couldn't find any information where to file it

If anyone can write this information in the book, that would be amazing otherwise someone please help me where to file a pull request on it. Maybe I will write it for them

EDIT:  There was a typo, you need to write //! instead of ///! which I wrote above. Thanks to u/Zenithsiz for pointing it out
