# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (25/2020)!](https://www.reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 343](https://www.reddit.com/r/rust/comments/hactqu/this_week_in_rust_343/)
- url: https://this-week-in-rust.org/blog/2020/06/16/this-week-in-rust-343/
---

## [3][PyO3 now supports stable Rust!](https://www.reddit.com/r/rust/comments/hd5506/pyo3_now_supports_stable_rust/)
- url: https://github.com/PyO3/pyo3/issues/5#issuecomment-647094806
---

## [4][Installed and Compiled Rust on Raspberry Pi Zero W](https://www.reddit.com/r/rust/comments/hcyolu/installed_and_compiled_rust_on_raspberry_pi_zero_w/)
- url: https://i.redd.it/u5fo8jzi86651.jpg
---

## [5][Announcing lexical-sort, a crate for sorting strings properly (and quickly) 🎉](https://www.reddit.com/r/rust/comments/hcw2re/announcing_lexicalsort_a_crate_for_sorting/)
- url: https://www.reddit.com/r/rust/comments/hcw2re/announcing_lexicalsort_a_crate_for_sorting/
---
The standard library implements `Ord` for strings, but the resulting ordering is unsuitable for many use cases. This is why I wrote the [lexical-sort](https://github.com/Aloso/lexical-sort) crate, which provides better string comparison functions.

It can sort strings **lexicographically**, which means that Unicode characters such as äßàćñæΩµð are transliterated to ASCII and converted to lowercase before the comparison. In effect, the comparison is case insensitive and works as expected for many languages, including English, German, French, Spanish, Portuguese, Italian, etc.

It can sort strings **naturally** (sometimes called "human sort"), which parses whole numbers instead of just digits, so "50" comes before "100". There's also an option to ignore non-alphanumeric characters.

According to my benchmark, its performance is very competitive. If all strings are ASCII, sorting them lexicographically is only 2–3 times as slow as the standard library. It's slower for non-ASCII strings, which is not surprising since it does much more work.

One use case for this is sorting files. Note file managers as well as tools like GNU `ls` usually sort files lexicographically.

Note that it will not return the desired order in some languages such as Swedish. Implementing correct lexicograpical order for every language is much harder. Also, just transliterating characters to ASCII is a bit of a hack, so if you have suggestions for improvement, they are very appreciated!
## [6][Announcing genco: A whitespace-aware quasiquoter for beautiful code generation (nightly only)](https://www.reddit.com/r/rust/comments/hd5g1v/announcing_genco_a_whitespaceaware_quasiquoter/)
- url: https://github.com/udoprog/genco
---

## [7][I decided to learn Rust as a first* language.](https://www.reddit.com/r/rust/comments/hcyp93/i_decided_to_learn_rust_as_a_first_language/)
- url: https://www.reddit.com/r/rust/comments/hcyp93/i_decided_to_learn_rust_as_a_first_language/
---
I have been wanting to actually learn to code for a while. I am fairly proficient in R for data analysis and visualization (including some ML) and write SQL queries daily for work, but I feel that most of the expertise won't really translate. I searched most loved programming languages and ended up here.  

Seems like Rust is not an easy language to begin with, but the documentation seems fantastic and the community seems friendly, so fuck it, we are doing it live. Wish me luck!
 (1 hour in I am already a bit frustrated about "size not known at compile time")
## [8][TIL: You can add strings in Rust](https://www.reddit.com/r/rust/comments/hcvagr/til_you_can_add_strings_in_rust/)
- url: https://www.reddit.com/r/rust/comments/hcvagr/til_you_can_add_strings_in_rust/
---
I've been using Rust for years now I only just learned that you can use `+` and `+=` for string concatenation.  Did anyone else not know this?

Edit:
FYI `+` has been around since 1.0 and `+=` was added in 1.14 (which was back in 2016).
## [9][Does Rust have anything like placement new?](https://www.reddit.com/r/rust/comments/hd3myx/does_rust_have_anything_like_placement_new/)
- url: https://www.reddit.com/r/rust/comments/hd3myx/does_rust_have_anything_like_placement_new/
---
I'm once again considering porting from C++ to something else and I had a quick question about Rust. As the title says does Rust have anything like a placement new?
## [10][Writing Non-Trivial Macros in Rust](https://www.reddit.com/r/rust/comments/hcqn81/writing_nontrivial_macros_in_rust/)
- url: http://adventures.michaelfbryan.com/posts/non-trivial-macros/?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=non-trivial-macros
---

## [11][Second project : TermUI : A Window manager for the command line.](https://www.reddit.com/r/rust/comments/hcyxoa/second_project_termui_a_window_manager_for_the/)
- url: https://www.reddit.com/r/rust/comments/hcyxoa/second_project_termui_a_window_manager_for_the/
---
It's all here : [https://github.com/Teln0/TermUI](https://github.com/Teln0/TermUI). Feel free to contribute !
## [12][Why the `finish` function on `std::hash::Hasher` returns a fixed `u64` type?](https://www.reddit.com/r/rust/comments/hd35ri/why_the_finish_function_on_stdhashhasher_returns/)
- url: https://www.reddit.com/r/rust/comments/hd35ri/why_the_finish_function_on_stdhashhasher_returns/
---
Hello rustaceans! I'm an intermediate rustacean, recently checking out detailed Rust library designs. I saw this function in the standard library: 

```rust
fn finish(&amp;self) -&gt; u64
```

(Link: https://doc.rust-lang.org/nightly/std/hash/trait.Hasher.html#tymethod.finish)

I noticed that this function returns a fixed `u64` type, instead of a generic type or associated type. I'm wondering how the authors discussed and analyzed on this design at that time. Just curious to discuss on this detailed API design and/or its history, didn't mean any harms to anyone. Thanks! :)
