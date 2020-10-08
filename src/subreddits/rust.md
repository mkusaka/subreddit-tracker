# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (41/2020)!](https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 359](https://www.reddit.com/r/rust/comments/j71tq7/this_week_in_rust_359/)
- url: https://this-week-in-rust.org/blog/2020/10/07/this-week-in-rust-359/
---

## [3][Woohoo, milestones being made. Updated 2 days ago](https://www.reddit.com/r/rust/comments/j76xgg/woohoo_milestones_being_made_updated_2_days_ago/)
- url: https://i.redd.it/o7w56ku9xsr51.png
---

## [4][Designing a New Rust Class at Stanford: Safety in Systems Programming](https://www.reddit.com/r/rust/comments/j6yht2/designing_a_new_rust_class_at_stanford_safety_in/)
- url: https://reberhardt.com/blog/2020/10/05/designing-a-new-class-at-stanford-safety-in-systems-programming.html
---

## [5][density-mesh - Image to 2D mesh converter](https://www.reddit.com/r/rust/comments/j73ijj/densitymesh_image_to_2d_mesh_converter/)
- url: https://i.redd.it/1dyzc32ftrr51.png
---

## [6][This Month in Rust GameDev #14 - September 2020](https://www.reddit.com/r/rust/comments/j722i5/this_month_in_rust_gamedev_14_september_2020/)
- url: https://rust-gamedev.github.io/posts/newsletter-014
---

## [7][The cost of Rust async/await](https://www.reddit.com/r/rust/comments/j703ge/the_cost_of_rust_asyncawait/)
- url: https://github.com/jkarneges/rust-async-bench
---

## [8][Introducing `code-tour`, an attempt to improve Rust example-based learning approach. Any feedbacks?](https://www.reddit.com/r/rust/comments/j7a5pa/introducing_codetour_an_attempt_to_improve_rust/)
- url: https://www.reddit.com/r/rust/comments/j7a5pa/introducing_codetour_an_attempt_to_improve_rust/
---
Hello folks,

Today I'm releasing [code-tour](https://github.com/Hywan/code-tour-rs), which is my attempt to improve Rust example-based learning approach. Basically, it is `cargo run --example` on steroid. It turns:

    // examples/foo.rs
    
    use code_tour::code_tour;
    
    #[derive(Debug)]
    struct S {
        x: i32,
        y: i32,
    }
    
    #[code_tour]
    fn main() {
        /// Declare something.
        /// Because it's an example.
        let a = S { x: 7, y: 42 };
    
        let b = a.x + a.y;
    
        /// Here is the result!
        let c = b + 1;
    }

into the following when runned with `cargo run --example foo`:

[\`cargo run --example foo\`](https://preview.redd.it/5atxloadbur51.png?width=670&amp;format=png&amp;auto=webp&amp;s=eb91c652fd2b36a3e5880b001ed158c03efa85cf)

The idea is that when a user runs an example, it usually prints nothing except if the author has added `println!` and `dbg!` statements everywhere. Moreover, all comments are useless, except if the user has the example opened side-by-side with the cargo run.

With code-tour, statements (+ doc) are displayed at runtime. The results are twofold:

1. It simplified the *editorial* work for the authors, and
2. It provides a better *experience* to the users!

This crate also provides an *interactive* mode (for long examples), and a *quiet* mode (e.g. for CI).

I'm sharing this here today in order to collect feedbacks:

* Do you think it is useful for you?
* How would you improve this?
* Why 42?

Repository: [https://github.com/Hywan/code-tour-rs](https://github.com/Hywan/code-tour-rs)

On crates.io: [https://crates.io/crates/code-tour](https://crates.io/crates/code-tour)
## [9][Iterators in Rust](https://www.reddit.com/r/rust/comments/j794ro/iterators_in_rust/)
- url: https://blog.thoughtram.io/iterators-in-rust/
---

## [10][Make a Language in Rust, Part Eight: Function Definitions](https://www.reddit.com/r/rust/comments/j7clns/make_a_language_in_rust_part_eight_function/)
- url: https://arzg.github.io/lang/8/
---

## [11][Announcing Rust 1.47.0](https://www.reddit.com/r/rust/comments/j7d49v/announcing_rust_1470/)
- url: https://blog.rust-lang.org/2020/10/08/Rust-1.47.html
---

## [12][Make a Language in Rust, Part Seven: A REPL](https://www.reddit.com/r/rust/comments/j6p8wt/make_a_language_in_rust_part_seven_a_repl/)
- url: https://arzg.github.io/lang/7/
---

