# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (40/2020)!](https://www.reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (40/2020)?](https://www.reddit.com/r/rust/comments/j1jime/whats_everyone_working_on_this_week_402020/)
- url: https://www.reddit.com/r/rust/comments/j1jime/whats_everyone_working_on_this_week_402020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-40-2020/49429?u=llogiq)!
## [3][My frist project in Rust generates mazes and solves them. I'm sure I've missed a lot of Rust patterns and would love some comments on my code (MIC)](https://www.reddit.com/r/rust/comments/j1xzvu/my_frist_project_in_rust_generates_mazes_and/)
- url: https://i.imgur.com/Y8WPfNk.gifv
---

## [4][Rust is a &gt;= $15 unlock for latest Humble Bundle from No Starch Press](https://www.reddit.com/r/rust/comments/j1khe3/rust_is_a_15_unlock_for_latest_humble_bundle_from/)
- url: https://www.humblebundle.com/books/learn-to-code-the-fun-way-no-starch-press-books
---

## [5][What (not so) recently happened in Miri](https://www.reddit.com/r/rust/comments/j1fxd3/what_not_so_recently_happened_in_miri/)
- url: https://www.ralfj.de/blog/2020/09/28/miri.html
---

## [6][which crate provides a similar compression ratio as "deflate 1.2.11" does?](https://www.reddit.com/r/rust/comments/j1syh3/which_crate_provides_a_similar_compression_ratio/)
- url: https://www.reddit.com/r/rust/comments/j1syh3/which_crate_provides_a_similar_compression_ratio/
---
Which Rust crate provides a similar compression ratio as "deflate 1.2.11" does?

A C-based application is using  " deflate 1.2.11 Copyright 1995-2017 Jean-loup Gailly and Mark Adler "; I want to use Rust to replace the application entirely but I can't find a proper crate which can replace deflate 1.2.11.

sample input 1: **01400000000300000f4c150bc06c0000ff**

sample output 1: **6204093023090000**

sample input 2: **01400000000200000f4c150bc06c0000ff**

sample output 2: **6274606060606260e0f711e53e90c3c0f01f00**

I have tried all methods in crate flate2 but most of them only have similar results when the stream is small like above. Once the stream has data more than 500 bytes, the output has a very big gap.
## [7][Put a lifetime in my Error type (anti-pattern?)](https://www.reddit.com/r/rust/comments/j1x5p5/put_a_lifetime_in_my_error_type_antipattern/)
- url: https://www.reddit.com/r/rust/comments/j1x5p5/put_a_lifetime_in_my_error_type_antipattern/
---
I'm writing a parser that tries to allocate as little memory as possible and I want to have an error type that has a string slice indicating the error within the input string.

Is having lifetimes in error types an anti-pattern?
## [8][I was wrong. CRDTs are the future](https://www.reddit.com/r/rust/comments/j1hb3a/i_was_wrong_crdts_are_the_future/)
- url: https://josephg.com/blog/crdts-are-the-future/
---

## [9][Blog Post: How to Make a 💡?](https://www.reddit.com/r/rust/comments/j1gov0/blog_post_how_to_make_a/)
- url: https://rust-analyzer.github.io/blog/2020/09/28/how-to-make-a-light-bulb.html
---

## [10][rust-analyzer changelog #44](https://www.reddit.com/r/rust/comments/j1e7g2/rustanalyzer_changelog_44/)
- url: https://rust-analyzer.github.io/thisweek/2020/09/28/changelog-44.html
---

## [11][Do we have a .rs subdomain for Rust projects ?](https://www.reddit.com/r/rust/comments/j1sio1/do_we_have_a_rs_subdomain_for_rust_projects/)
- url: https://www.reddit.com/r/rust/comments/j1sio1/do_we_have_a_rs_subdomain_for_rust_projects/
---
Similar to [js.org](https://js.org) do we have anything for Rust developer to get subdomain from [xyz.lib.rs](https://xyz.lib.rs) ? or similar !
## [12][inline-proc: write procedural macros directly in your code, instead of having to use another crate](https://www.reddit.com/r/rust/comments/j1j38x/inlineproc_write_procedural_macros_directly_in/)
- url: https://github.com/KaiJewson/inline-proc
---

