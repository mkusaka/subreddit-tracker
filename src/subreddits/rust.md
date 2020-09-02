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
## [3][Flume 0.8, a fast &amp; lightweight MPMC, released: Performance improvements, async support, select API, eventual fairness, Send + Sync + Clone, multiple receivers, no unsafe, timeout/deadline support](https://www.reddit.com/r/rust/comments/il4nq5/flume_08_a_fast_lightweight_mpmc_released/)
- url: https://github.com/zesterer/flume
---

## [4][PSA: What is `const fn`?](https://www.reddit.com/r/rust/comments/iksmgk/psa_what_is_const_fn/)
- url: https://www.reddit.com/r/rust/comments/iksmgk/psa_what_is_const_fn/
---
Recently there's been a lot of discussion about what `const fn` actually is. For example, see the [recent comments on the Rust 1.46 announcement](https://old.reddit.com/r/rust/comments/ihnnnz/announcing_rust_1460_rust_blog/). This post is my personal answer to the question but it's by no means authoritative so please provide corrections!

So why is `const fn` useful? One thing you cannot ordinarily do with a function is call it in a constant position. Let's see what I mean:

```

    // a normal function
    fn foo(n: usize) -&gt; usize {
        n + 1
    }

    fn main() {
        const BAR: usize = foo(5);
        let array = [0_u8; foo(7)];
    }

```

Both of the lines in `main` will cause an error that may look something like this:

    error[E0015]: calls in constants are limited to constant functions, tuple structs and tuple variants
     --&gt; src/main.rs:7:28
      |
    7 |         const BAR: usize = foo(5);
      |                            ^^^^^^

As the error says, this problem can be solved by using a "constant function" aka `const fn`. In this case all we need to do is change the definition of `foo` to:

```

    // a constant function
    const fn foo(n: usize) -&gt; usize {
        n + 1
    }

```

And now `main` will work! But there's a catch. `const fn` is currently much more limited than a normal function. This simple example works and you can also now use `loop` and `if` in constant functions but there's a lot that won't work. Perhaps most notably, iterators won't work at all. Which amongst other things means no `for` loop. Work is continuing to allow more features in `const fn` but there will always be some functions that cannot be `const fn` (e.g. ones that rely on platform APIs).

In summary, that's all `const fn` means. You can use it in more places than an ordinary function. It can be used as a `static` variable, a `const` variable, as a length for an array, etc.

So should you mark every function as `const fn` so long as it compiles? No! `const fn` is a contract between you and the people that call your function. You are declaring that you will never change the function in ways that are invalid for `const fn`. This contract may prevent some future optimizations so think carefully before using it.

-----

At this point you might be wondering why I haven't mentioned the phrase "compile time". That's because it's not what `const fn` is really about. Just about all functions can be run either partially or fully at compile time (the optimizer permitting) and a `const fn` may be run at runtime. Of course sometimes you do want to make 100% sure a function is run at compile time which is a useful side-effect of doing `const BAR: usize = foo();`. Note though that if you just do `let bar = foo();` (where `foo` is a `const fn`) it's pretty much the same as using an ordinary function which may or may not be run at compile time.

In the future there might be [inline `const`](https://github.com/ecstatic-morse/rfcs/blob/inline-const/text/0000-inline-const.md) that explicitly declares that you want to run part of your program at compile time. For example:

```

    let bar = const {
        // The compiler will run this whole block at compile time
        // or else it will produce an error
    };

``` 

But this is very far from being stabilized (it's not even implemented yet).
## [5][VSCode, multiple project roots, rust-analyzer, clippy, and rustfmt](https://www.reddit.com/r/rust/comments/ikzbdj/vscode_multiple_project_roots_rustanalyzer_clippy/)
- url: https://www.reddit.com/r/rust/comments/ikzbdj/vscode_multiple_project_roots_rustanalyzer_clippy/
---
I'm sure it's intuitive to most, but to spare anyone in the future from spending an hour trying to get VSCode to work with multiple project roots in your workspace and rust-analyzer with clippy and rustfmt on save, here are the little pieces from 10 different websites that I pieced together to make it work.

# Multiple project roots in workspace:

If you're like me and want a `rust` folder as the root of your VSCode workspace to contain multiple crates for testing or generally messing around with code examples, you want to set up a **cargo workspace**.

[Here's the documentation on how](https://doc.rust-lang.org/book/ch14-03-cargo-workspaces.html).

It boils down to manually creating a `Cargo.toml` file in the root of your VSCode workspace. It only needs one section, `[workspace]` containing a list of each crate you created with `cargo run &lt;crate name&gt;` and then running `cargo build` in the root directory. Here's an example:


    [workspace]
    
    members = [
        "fun-project",
        "other-project",
    ]

# rust-analyzer:

[Here's](https://marketplace.visualstudio.com/items?itemName=matklad.rust-analyzer) the Visual Studio Marketplace page with the latest (as of the time of this post) rust-analyzer plugin for VSCode. 

You can hit `Ctrl + P` in VSCode and paste in `ext install matklad.rust-analyzer` to install it. After installation, you will be prompted to install the language server in VSCode. Install it by clicking `Download now`.

# Clippy

To install clippy, run `rustup component add clippy`

# Rustfmt

To install rustfmt, run `rustup component add rustfmt`

# Putting it all together

Add these lines to your VSCode `settings.json` file and save it. You may be prompted to restart rust-analyzer after save:

    "rust-analyzer.checkOnSave.command": "clippy", // Replaces the default `cargo check` command with `cargo clippy` on Rust file save
    "editor.formatOnSave": true,                   // General VSCode setting to allow a formatter, in this case rustfmt, to format text on save
    "rust-analyzer.cargo.allFeatures": true,       // For fun. Activates all features of rust-analyzer. You don't need this one.
    "rust-analyzer.updates.channel": "nightly",    // For fun. Enables the nightly updates channel for rust-analyzer since they ask nicely in their documentation. You don't need this one. You will be prompted to download the nightly install after save.

&amp;nbsp;

This should be everything you need to get coding!
## [6][Show r/rust: a `no_std` NonEmpty types inspired by Haskell](https://www.reddit.com/r/rust/comments/il3pdy/show_rrust_a_no_std_nonempty_types_inspired_by/)
- url: https://crates.io/crates/oom
---

## [7][Bottlerocket is a free and open-source Linux-based operating system meant for hosting containers](https://www.reddit.com/r/rust/comments/ikqqm8/bottlerocket_is_a_free_and_opensource_linuxbased/)
- url: https://github.com/bottlerocket-os/bottlerocket
---

## [8][Refactoring Rust Transpiled from C](https://www.reddit.com/r/rust/comments/ikqzni/refactoring_rust_transpiled_from_c/)
- url: https://immunant.com/blog/2020/09/transpiled_c_safety/
---

## [9][Crepe: fast, compiled Datalog in Rust](https://www.reddit.com/r/rust/comments/ikszdg/crepe_fast_compiled_datalog_in_rust/)
- url: https://www.reddit.com/r/rust/comments/ikszdg/crepe_fast_compiled_datalog_in_rust/
---
I'm excited to share a recent project on implementing compiled Datalog in Rust, for areas like static analysis and database systems. [Crepe](https://github.com/ekzhang/crepe/) is a procedural macro that lets you easily write declarative Datalog-like code, and it interoperates fully with larger programs. As a simple example, you can compute reachability in a directed graph like so:

    use crepe::crepe;
    
    crepe! {
        @input
        struct Edge(i32, i32);
    
        @output
        struct Reachable(i32, i32);
    
        Reachable(x, y) &lt;- Edge(x, y);
        Reachable(x, z) &lt;- Edge(x, y), Reachable(y, z);
    }
    
    fn main() {
        let mut runtime = Crepe::new();
        runtime.extend(&amp;[Edge(1, 2), Edge(2, 3), Edge(3, 4), Edge(2, 5)]);
    
        let (reachable,) = runtime.run();
        for Reachable(x, y) in reachable {
            println!("node {} can reach node {}", x, y);
        }
    }

The advantage of this approach over something like [Souffle](https://souffle-lang.github.io/) is that you get seamless integration and portability with pure Rust code (no need for C++ environment setup), and unlike the low-level [Datafrog](https://github.com/rust-lang/datafrog) engine, you can write ergonomic, high-level specifications.

We also support semi-naive evaluation and stratified negation (see [Chapter 3 of this book](http://blogs.evergreen.edu/sosw/files/2014/04/Green-Vol5-DBS-017.pdf) for more info). What I think is really fun is the ability to extend Datalog rules with **arbitrary Rust syntax**, by allowing any Rust expression to be inlined into the bodies and conditions of rules. For example, we can use arithmetic to compute all paths in a graph with length less than 20:

    use crepe::crepe;
    
    const MAX_PATH_LEN: u32 = 20;
    
    crepe! {
        @input
        struct Edge(i32, i32, u32);
    
        @output
        struct Walk(i32, i32, u32);
    
        @output
        struct NoWalk(i32, i32);
    
        struct Node(i32);
    
        Node(x) &lt;- Edge(x, _, _);
        Node(x) &lt;- Edge(_, x, _);
    
        Walk(x, x, 0) &lt;- Node(x);
        Walk(x, z, len1 + len2) &lt;-
            Edge(x, y, len1),
            Walk(y, z, len2),
            (len1 + len2 &lt;= MAX_PATH_LEN);
    
        NoWalk(x, y) &lt;- Node(x), Node(y), !Walk(x, y, _);
    }

In particular, any part of a rule can have inline Rust expressions, so you could even add a clause like `PrimeWalk(x, y) &lt;- Walk(x, y, len), (primes::is_prime(len as u64));` to compute **all paths of prime length** in a weighted graph! (here relying on the [primes](https://docs.rs/primes/0.3.0/primes/) crate)

As mentioned before, this is not the first time Datalog and its extensions have been compiled for static analysis purposes (see [Souffle](https://souffle-lang.github.io/), [Formulog](https://github.com/HarvardPL/formulog)). However, I think implementing it as a procedural macro makes for a really ergonomic, intuitive way to play with Datalog code and integrate it into larger applications.

This is my first published crate, happy to receive any advice you have! Code: [https://github.com/ekzhang/crepe/](https://github.com/ekzhang/crepe/)
## [10][Are you interested in a tutorial series on making your own language using Rust?](https://www.reddit.com/r/rust/comments/ikiag7/are_you_interested_in_a_tutorial_series_on_making/)
- url: https://www.reddit.com/r/rust/comments/ikiag7/are_you_interested_in_a_tutorial_series_on_making/
---
Hi,

Would anyone be interested in a tutorial series on creating an interpreted language using Rust? My plan is to start the tutorials as simple and easy to follow as possible, and gradually work upwards from there to improve the language. I’ve only created one language (still WIP at the moment), so I’m not very experienced in this area, but I have a strong grasp on the fundamentals as a result of rewriting my pet language from scratch many times. Creating the series will hopefully also be a learning experience for me :)

At first I thought that the tutorials should be in text form, but then I remembered how much more engaged and interested I am when watching tutorial content in video form. I believe that this may also be easier to make for me, as I find things easier to explain with speech rather than writing. However, I’ve heard quite a few people say they dislike video content, as it doesn’t allow them to skim and slows them down. Another reason that text might be better is because I’m 16, so it would probably be strange watching a tutorial where the person teaching isn’t an adult (I know I would find it strange).

So, would you be interested?

[View Poll](https://www.reddit.com/poll/ikiag7)
## [11][Introducing Curio - A blazing fast HTTP client.](https://www.reddit.com/r/rust/comments/ikr16f/introducing_curio_a_blazing_fast_http_client/)
- url: https://www.reddit.com/r/rust/comments/ikr16f/introducing_curio_a_blazing_fast_http_client/
---
Hey there folks, i would like to present to you my current project [Curio](https://github.com/fatalcenturion/curio), a blazing fast HTTP Client, written entirely in Rust.

&amp;#x200B;

Some of the things it currently supports:

* HTTPS automatic redirect
* GET, POST, OPTIONS, DELETE, HEAD methods

&amp;#x200B;

You can view benchmarks against other popular HTTP clients [here](https://github.com/fatalcenturion/curio#benchmarks). Independant benchmarkers to help verify these would be appreciated.

You can begin using Curio by checking out some of the examples on the repository [here](https://github.com/fatalcenturion/curio) or by reading the documentation on [docs.rs](https://docs.rs/curio/0.0.3/curio/).

Overall, i would appreciate if you kind folks could check it out and provide me with some feedback. Please let me know what you think of my project, and how i could improve it!

Edit: Benchmark utility can be found [here](https://github.com/fatalcenturion/curio-benchmarks)
## [12][Is it possible to generate a type from within a pattern-matching macro?](https://www.reddit.com/r/rust/comments/il4d6n/is_it_possible_to_generate_a_type_from_within_a/)
- url: https://www.reddit.com/r/rust/comments/il4d6n/is_it_possible_to_generate_a_type_from_within_a/
---
I'm attempting to implement a DSL using macros which would include a short-hand for struct definitions.  So for instance, I want to be able to have a block like this:

    my_macro! {
        Point2D -&gt; i32, i32
        Point3D -&gt; i32, i32, i32
    }

Which would create these types:

    struct Point2D {
        v0: i32,
        v1: i32
    }

    struct Point3D {
        v0: i32,
        v1: i32,
        v2: i32
    }


Is this possible using match rules?

I was trying something like this:

    macro_rules! context_free_grammar {
        ($new_type:ident -&gt; $other_type:ty) =&gt; { ... }
    }

But it seems like it wants the new_type identity to exist outside of the macro block.
