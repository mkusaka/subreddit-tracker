# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (22/2020)!](https://www.reddit.com/r/rust/comments/gr3r32/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gr3r32/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 340](https://www.reddit.com/r/rust/comments/grs1ql/this_week_in_rust_340/)
- url: https://this-week-in-rust.org/blog/2020/05/27/this-week-in-rust-340/
---

## [3][Rust finally shows up as filter option when searching GitHub](https://www.reddit.com/r/rust/comments/gt2roa/rust_finally_shows_up_as_filter_option_when/)
- url: https://i.redd.it/45dfpuao5s151.png
---

## [4][Auto-Vectorization for Newer Instruction Sets in Rust](https://www.reddit.com/r/rust/comments/gt83zg/autovectorization_for_newer_instruction_sets_in/)
- url: https://www.nickwilcox.com/blog/autovec2/
---

## [5][What does #[cold] actually do?](https://www.reddit.com/r/rust/comments/gtcsem/what_does_cold_actually_do/)
- url: https://www.reddit.com/r/rust/comments/gtcsem/what_does_cold_actually_do/
---
[The Reference's description of the `cold` attribute](https://doc.rust-lang.org/reference/attributes/codegen.html#the-cold-attribute) says almost nothing about it. It sounds useful, but I want to know more before I use it, and information about it seems scarce.

What effect does it have? Where is it effective? I imagine it has an effect on functions, but what about individual statements or blocks? Is it useful to mark an entire module as cold? Should I sprinkle it all over code paths for handling rare errors (e.g. I/O error when writing a file), or would that be unwise or unnecessary? Can it be used for branch weights like the `unlikely` intrinsic (which is unstable)?
## [6][Announcing `float_eq` 0.3.0](https://www.reddit.com/r/rust/comments/gtbuvu/announcing_float_eq_030/)
- url: https://www.reddit.com/r/rust/comments/gtbuvu/announcing_float_eq_030/
---
Hi!

I've just published version 0.3.0 of [float\_eq](https://crates.io/crates/float_eq), for comparing floating point values.

Notable changes include a much more rigorous and consistent treatment of special floating point values, so ULPs checks now respect NaNs, and infinities are correctly handled.

Epsilons for composite types may now be provided on a per-field basis:

    assert_float_eq!(a, b, ulps &lt;= Complex32Ulps { re: 2, im: 4 });
    assert_float_eq!([1.0, -2.0], [-1.0, -2.5], abs &lt;= [2.0, 0.5]);
    assert_float_eq!((1.0f32, 2.0f64), (-1.0, 2.5), abs &lt;= (2.0, 0.5))

Or, where applicable, uniformly across all fields:

    assert_float_eq!(a, b, ulps_all &lt;= 4);
    assert_float_eq!([1.0, -2.0], [-1.0, -2.5], abs_all &lt;= 2.0);

There is also now support for tuples (of up to size 12, inclusive), with more support for standard types planned.

Plus a variety of other fixes and tweaks, see the [CHANGELOG](https://github.com/jtempest/float_eq-rs/blob/master/CHANGELOG.md) for more details.

Constructive criticism and feedback appreciated, I hope you're having a lovely day :)
## [7][zoxide 0.4.1, a faster way to navigate your filesystem - now 2-3x faster](https://www.reddit.com/r/rust/comments/gt42jm/zoxide_041_a_faster_way_to_navigate_your/)
- url: https://github.com/ajeetdsouza/zoxide
---

## [8][Steve Klabnik Gives a Friendly &amp; Thorough Introduction to Async Await](https://www.reddit.com/r/rust/comments/gst7y7/steve_klabnik_gives_a_friendly_thorough/)
- url: https://www.youtube.com/watch?v=NNwK5ZPAJCk
---

## [9][xhacklight - the first time I really needed Rust.](https://www.reddit.com/r/rust/comments/gte05a/xhacklight_the_first_time_i_really_needed_rust/)
- url: https://www.reddit.com/r/rust/comments/gte05a/xhacklight_the_first_time_i_really_needed_rust/
---
I don't get to use Rust at work (which is reasonable for what I do) and also haven't found any personal needs for something more performant than Python or JS. However, this time I just needed a real executable instead of a script - even though it was just for being able to use the SUID bit...

Well, here is [`xhacklight`](https://github.com/SillyFreak/xhacklight), a hacky replacement for `xbacklight` that works at the very least on my machine. I doubt, it will be useful to anyone else, but if anyone wants to critique my code, please do.

Cheers!
## [10][Elektron (synth manufacturer) looking to hire Rust Graphics Engineer](https://www.reddit.com/r/rust/comments/gt0kgw/elektron_synth_manufacturer_looking_to_hire_rust/)
- url: https://www.elektron.se/rust-graphics-engineer/
---

## [11][Implementing a trait only for types satisfying a const-generic predicate! (Nightly Only)](https://www.reddit.com/r/rust/comments/gt067a/implementing_a_trait_only_for_types_satisfying_a/)
- url: https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=39c1d2d4169f182c4fe575dca1286dea
---

## [12][Ncurses update screen](https://www.reddit.com/r/rust/comments/gtchz2/ncurses_update_screen/)
- url: https://www.reddit.com/r/rust/comments/gtchz2/ncurses_update_screen/
---
Hey guys, I'm new to Rust and i'm trying to make a little terminal application for managing my todos, i'm using ncurses. But I have a problem, i'm not able to update the screen, I have the todos listed via a for loop, but whenever the screen is refreshed it moves to the next line and moves to the right, with an offset value of "looks like the length of the previous render text length". Anyone knows how to fix this issue? I'm sorry if i'm being vague, I can also provide a github link for those interested in helping me out. Thank you.
