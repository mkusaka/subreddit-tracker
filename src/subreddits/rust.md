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

## [3][Got this baby in today, gonna enjoy some light reading later](https://www.reddit.com/r/rust/comments/hbtt65/got_this_baby_in_today_gonna_enjoy_some_light/)
- url: https://i.redd.it/plk3qzaxks551.jpg
---

## [4][Small trick destructuring slices with known length](https://www.reddit.com/r/rust/comments/hbxtrm/small_trick_destructuring_slices_with_known_length/)
- url: https://www.reddit.com/r/rust/comments/hbxtrm/small_trick_destructuring_slices_with_known_length/
---
Just want to share a small trick I found on destructuring the contents of a slice into local variables.

Instead of using `if let` like this:

    if let [a, b, c] = k {
        Some(a * x.powi(2) + b * x + c)
    } else {
        None
    }

... you can use the `TryInto` trait together with the `try` (`?`) operator for a more linear code flow:

    use std::convert::TryInto;
    let [a, b, c]: [_; 3] = k.try_into().ok()?;
    Some(a * x.powi(2) + b * x + c)

[Playground example](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=46498dd319aa14d5e5286246e03e0f4e)
## [5][bygge - Build a Rust project without Cargo*](https://www.reddit.com/r/rust/comments/hc07zm/bygge_build_a_rust_project_without_cargo/)
- url: https://fnordig.de/2020/06/19/build-your-project/
---

## [6][rust-highfive's contributions graph - is this because of COVID?](https://www.reddit.com/r/rust/comments/hbghcm/rusthighfives_contributions_graph_is_this_because/)
- url: https://i.redd.it/kcoukwx5qo551.png
---

## [7][80x faster Rust code?!](https://www.reddit.com/r/rust/comments/hbvjer/80x_faster_rust_code/)
- url: https://github.com/choiway/prognog/blob/master/README.md
---

## [8][Announcing Rust 1.44.1](https://www.reddit.com/r/rust/comments/hbgayc/announcing_rust_1441/)
- url: https://blog.rust-lang.org/2020/06/18/Rust.1.44.1.html
---

## [9][Crafting cellular automata in Rust (my first blog post, looking for feedback)](https://www.reddit.com/r/rust/comments/hbu84n/crafting_cellular_automata_in_rust_my_first_blog/)
- url: https://www.reddit.com/r/rust/comments/hbu84n/crafting_cellular_automata_in_rust_my_first_blog/
---
Hey everyone, I just graduated from university and have decided to take the lull in job opportunities to brush back up on my Rust. I am trying to blog about the process of writing a cellular automata library in Rust, but this first piece was more of just a "test" on writing a blog post at all. I'm just looking for basic feedback about the writing, and code style, so anything is appreciated.

[Here's the post](https://oneorten.dev/blog/automata_rust_1/)

Thank you for reading!
## [10][Announcing `float_eq` 0.4.0](https://www.reddit.com/r/rust/comments/hbyq2o/announcing_float_eq_040/)
- url: https://www.reddit.com/r/rust/comments/hbyq2o/announcing_float_eq_040/
---
Hi!

I've just published version 0.4.0 of [float\_eq](https://crates.io/crates/float_eq), for comparing floating point values.

This release adds support for comparing the contents of more standard Rust types - slices, Option, Vec, VecDeque, LinkedList, BTreeMap and HashMap:

    let a = vec![1.0f32, 2.0];
    let b = vec![1.000_000_1, 2.000_000_5];
    
    assert_float_eq!(a, b, ulps &lt;= vec![1, 2]);
    assert_float_eq!(a, b, ulps_all &lt;= 2);

This is in addition to the support added in 0.3.1 for comparing mutable and immutable reference types and the contents of Cell, RefCell, Rc, Arc and Box instances.

You may also now `#[derive]` any or all of the float\_eq traits for non-generic structs and tuple structs, by enabling the optional `"derive"` feature:

    #[derive(
        Debug, PartialEq, FloatUlps, FloatDiff, FloatEq, FloatEqDebug,
        FloatEqAll, FloatEqAllDebug,
    )]
    #[float_eq(ulps = "PointUlps", all_epsilon = "f64")]
    struct Point {
        x: f64,
        y: f64,
    }
    
    let a = Point { x: 1.0, y: -2.0 };
    let b = Point { x: 1.1, y: -2.2 };
    assert_float_eq!(a, b, abs &lt;= Point { x: 0.15, y: 0.25 });
    assert_float_eq!(a, b, abs_all &lt;= 0.25);
    
    let c = Point { x: 1.000_000_000_000_000_9, y: -2.000_000_000_000_001_3 };
    assert_float_eq!(a, c, ulps &lt;= PointUlps { x: 4, y: 3 });
    assert_float_eq!(a, c, ulps_all &lt;= 4);

**How does float\_eq differ from approx?**

Both provide similar comparison capabilities, with different underlying philosophies and APIs, see [this post](https://www.reddit.com/r/rust/comments/fzwxiy/announcing_float_eq_for_explicitly_bounded/fn6nkka/) for a longer discussion.

**Why write yet another floating point equality library?**

To explore providing a different kind of toolkit for working with floating point numbers. Since I am fairly new to Rust (though very experienced with C++) it seemed like a good way to introduce myself to some of the more advanced language and ecosystem features. Finally, I wanted to gain a deeper intuition for floating point representation and usage and you know what they say - don't reinvent the wheel unless you're planning on learning how wheels work :)
## [11][5 years of Rust - a cross-platform &amp; full-stack language of choice for the next decade(s)](https://www.reddit.com/r/rust/comments/hbwpv9/5_years_of_rust_a_crossplatform_fullstack/)
- url: https://gendignoux.com/blog/2020/06/16/rust-5-years-cross-platform.html
---

## [12][HELP Exposing static files with Rocket](https://www.reddit.com/r/rust/comments/hc0d97/help_exposing_static_files_with_rocket/)
- url: https://www.reddit.com/r/rust/comments/hc0d97/help_exposing_static_files_with_rocket/
---
So guys, i have few doubts about exposing static files with rocket.

First one, can i expose files which are not index or .html files?

If so in my code im trying to do something like this in the main():

`rocket::ignite()`  
`.mount("/", routes![index])`  
`.mount("/path",StaticFiles::from("/local/path/files"))`  
`.attach(Template::fairing())`  
`.launch();`

I'm not getting any error from rust, but simply when i try to see if any of the files in that path are availables i get a 404 Errror.

Can i do a "double" mount like i'm doing? Rust seems fine compiling this but i don't know.

How can i expose files like .mp3 or others on /path such i can contact them as (for example) /path/name\_file.mp3 ?

Cheers
