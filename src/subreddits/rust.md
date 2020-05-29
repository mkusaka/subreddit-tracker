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

## [3][Fuchsia Rust Crates](https://www.reddit.com/r/rust/comments/gspk26/fuchsia_rust_crates/)
- url: https://fuchsia.dev/fuchsia-src/development/languages/rust/crates
---

## [4][blocked: A proc-macro that emits a compiler warning when the linked GitHub issue is closed](https://www.reddit.com/r/rust/comments/gsq5gd/blocked_a_procmacro_that_emits_a_compiler_warning/)
- url: https://www.reddit.com/r/rust/comments/gsq5gd/blocked_a_procmacro_that_emits_a_compiler_warning/
---
Inspired by [this post](https://www.reddit.com/r/rust/comments/gs9cgl/blocked_convention_for_marking_code_as_being/) I wrote [my first proc-macro](https://github.com/zacps/blocked)!

    // An attribute-like procedural macro is on the todo-list
    #![feature(proc_macro_hygiene)]
    
    use blocked::blocked;
    
    fn hacky_workaround() {}
    
    fn main() {
        blocked!("1", "This code can be removed when the issue is closed");
        hacky_workaround();
    
        // The reason is optional
        blocked!("1");

        blocked!("rust-lang/rust#123");

        blocked!("https://github.com/rust-lang/rust/issues/123");
    }
    
If you don't specify the repository or organisation it'll infer them from git remotes (`upstream` and `origin`).

I'd like to improve the error messages and warning when an issue is closed. Of course, I also want to add support for Gitlab.

Requires nightly for emitting compiler warnings, and hygiene for function-like proc macros.

Edit: Forgot to link the thing
## [5][binread: Declarative Rust Binary Parsing](https://www.reddit.com/r/rust/comments/gsfb3p/binread_declarative_rust_binary_parsing/)
- url: https://jam1.re/blog/binread-a-declarative-rust-binary-parsing-library
---

## [6][3D boids swimming along in perfect harmony; Implementing the boids flocking algorithm in Rust.](https://www.reddit.com/r/rust/comments/gsldbi/3d_boids_swimming_along_in_perfect_harmony/)
- url: https://www.reddit.com/r/rust/comments/gsldbi/3d_boids_swimming_along_in_perfect_harmony/
---
I mashed together my interests in procedural graphics generation and Rust programming for this experiment. I implemented the boids flocking algorithm to create swirling patterns. The implementation is heavily inspired by Sebastian Lague's coding adventure with boids [video](https://www.youtube.com/watch?v=bqtqltqcQhw). Here's a small demo.

[Swimming along in perfect harmony](https://reddit.com/link/gsldbi/video/l2p1nwaofm151/player)

You can build and run it [source](https://github.com/twitu/rboids) to enjoy these little creatures endlessly. I've also written a blog about it. You can follow along and create your own flock of boids. Any criticism regarding the code and blog is welcome.

1. [A fistful of boids](https://blog.bitsacm.in/a-fistful-of-boids/) \- Setting up the scene and basic animation
2. [For a few boids more](https://blog.bitsacm.in/for-a-few-boids-more/) \- Generating boids and obstacle avoidance
3. [The school, the boid and the Rusty](https://blog.bitsacm.in/the-school-the-boid-and-the-rusty/) \- Simulating a flock, parallelism and benchmarking performance

Many knowledgeable people have written a lot about development in Rust. My conclusion, as a budding Rustaceans, is simple. The Rust ecosystem is not perfect, but that means I have the opportunity to contribute and make a difference.
## [7][Tour of Web Assembly - Chapter 1 Released](https://www.reddit.com/r/rust/comments/gsatru/tour_of_web_assembly_chapter_1_released/)
- url: https://tourofrust.com/webassembly/
---

## [8][Blog Post: coverage marks](https://www.reddit.com/r/rust/comments/gssqpp/blog_post_coverage_marks/)
- url: https://ferrous-systems.com/blog/coverage-marks/
---

## [9][showing rust for a friend](https://www.reddit.com/r/rust/comments/gsomji/showing_rust_for_a_friend/)
- url: https://www.reddit.com/r/rust/comments/gsomji/showing_rust_for_a_friend/
---
I will showcase rust to a friend where i will try to teach him some of the basics of rust and how it works. Are there any good examples on tasks that showcase the different perks of rust, that is easy to show and explain to someone primarly used to C#, Kotlin &amp; some minor experience with C++?
## [10][parser for c3d(biomechanics data) written in rust(early version)](https://www.reddit.com/r/rust/comments/gso9yj/parser_for_c3dbiomechanics_data_written_in/)
- url: https://crates.io/crates/c3d-rs
---

## [11][What is the point of Diesel? What problem is it trying to solve?](https://www.reddit.com/r/rust/comments/gsqq4d/what_is_the_point_of_diesel_what_problem_is_it/)
- url: https://www.reddit.com/r/rust/comments/gsqq4d/what_is_the_point_of_diesel_what_problem_is_it/
---
Do I really need Diesel to communicate with a postgresql database? I don't understand the problem it is trying to solve. Can't I just communicate directly between actix-web and postgresql?
## [12][An introduction to SIMD and ISPC in Rust](https://www.reddit.com/r/rust/comments/gsb0i3/an_introduction_to_simd_and_ispc_in_rust/)
- url: https://state.smerity.com/smerity/state/01E8RNH7HRRJT2A63NSX3N6SP1
---

