# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (20/2020)!](https://www.reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This week in Rust 338](https://www.reddit.com/r/rust/comments/gikfuy/this_week_in_rust_338/)
- url: https://this-week-in-rust.org/blog/2020/05/12/this-week-in-rust-338/
---

## [3][Five Years of Rust | Rust Blog](https://www.reddit.com/r/rust/comments/gk77zt/five_years_of_rust_rust_blog/)
- url: https://blog.rust-lang.org/2020/05/15/five-years-of-rust.html
---

## [4][What's New in IntelliJ Rust](https://www.reddit.com/r/rust/comments/gk8sbp/whats_new_in_intellij_rust/)
- url: https://blog.jetbrains.com/clion/2020/05/whats-new-in-intellij-rust/
---

## [5][NuShell: the shell where traditional Unix meets modern development, written in Rust - An interview with its creators](https://www.reddit.com/r/rust/comments/gjp2xo/nushell_the_shell_where_traditional_unix_meets/)
- url: https://notamonadtutorial.com/nushell-the-shell-where-traditional-unix-meets-modern-development-written-in-rust-caf92c2c7c98
---

## [6][Is it undefined behavior to mem::transmute between const-generic types? If so, Why? If not, why not?](https://www.reddit.com/r/rust/comments/gk3qaw/is_it_undefined_behavior_to_memtransmute_between/)
- url: https://www.reddit.com/r/rust/comments/gk3qaw/is_it_undefined_behavior_to_memtransmute_between/
---
**Answer**

* If you **do not use** `#[repr(C)]` then this **is** undefined behavior.
* If you **do use** `#[repr(C)]` then this **is not** undefined behavior.
* `#[repr(C)]` \[[docs](https://doc.rust-lang.org/nomicon/other-reprs.html)\]

**Original Post**

Here is a simple, minimal example: \[[playground](https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=c433eeb8f746ab0f0df2e5ebc5fd0116)\]

Using const generics, we can split the implementation of `Point`

* only a `Point&lt;{ Type::HasGetX }&gt;` can call `get_x()`
* only a `Point&lt;{ Type::HasGetY }&gt;` can call `get_y()`

We can then create a function, `get_xy_from()` which takes a `Point&lt;const T: Type&gt;` and transmutes it as needed in order to invoke the relevant functions.

This example appears to provide the correct outputs, which is encouraging, but that doesn't prove anything. The behavior could still be undefined and only appear to be correct in this one instance running in the playground.

Could someone with more knowledge about const generics shed some light on whether or not this is okay, and maybe try to explain a bit more about why it is or is not?
## [7][c-closures hits 1.0!](https://www.reddit.com/r/rust/comments/gjz6h5/cclosures_hits_10/)
- url: https://www.reddit.com/r/rust/comments/gjz6h5/cclosures_hits_10/
---
Hi everyone, I've just released version 1.0 of [`c-closures-build`](https://crates.io/crates/c-closures-build) along with [a new tutorial!](https://github.com/Xaeroxe/c-closures-rs/tree/master/example)

`c-closures-build` generates C/C++ FFI compatible `*Closure` structures from rust closures, useful for cross language functional programming.

If you saw my last post about this crate then you should probably look again, a lot of enhancements and improvements have been made to the crate. Also, the `c-closures` crate itself is now deprecated.
## [8][What, *exactly*, are lifetimes, in the mind of Rustc?](https://www.reddit.com/r/rust/comments/gk7aqh/what_exactly_are_lifetimes_in_the_mind_of_rustc/)
- url: https://users.rust-lang.org/t/what-exactly-are-lifetimes-in-the-mind-of-rustc/42662
---

## [9][2x slower cli performance with Clap](https://www.reddit.com/r/rust/comments/gk8q05/2x_slower_cli_performance_with_clap/)
- url: https://www.reddit.com/r/rust/comments/gk8q05/2x_slower_cli_performance_with_clap/
---
Does anyone know why my example command line app would be 2x as slow when I use clap's new derive feature (which is basically just structopt, available on the master branch) when compared to manual command line parsing? Both are build with `--release`

For a 1.2GB csv-like file with 8 columns, two each of float, int, bool, and String, (10M rows), manual command line args parsing makes the example run at [~450MB/s](https://i.imgur.com/jF5IlE3.png) on my machine. With clap, it runs at [~180 MB/s](https://i.imgur.com/0Cu52VS.png).

It also looks like the maximum memory consumed is increased by 700MB


You can see the code with manual command line parsing [here on the master branch](https://github.com/samedhg/sorer) and [here with clap](https://github.com/SamedhG/sorer/tree/clap). The code is in `examples/sorer.rs`.


(By the way, this was a class project so that's the reason for the weird file format that this parses)
## [10][Surprising performance implications of match](https://www.reddit.com/r/rust/comments/gjt98e/surprising_performance_implications_of_match/)
- url: https://www.reddit.com/r/rust/comments/gjt98e/surprising_performance_implications_of_match/
---
An enum can have fake lifetime parameters (and/or type parameters) by using `PhantomData` in one of its variants.  Such an enum can be safely be mapped from one lifetime (and/or type) to another using `match`, giving the `PhantomData` variant a new `PhantomData` value and keeping the original value within any other variant.  One would expect the enum with lifetime `'a` would have the same data layout as the same enum with lifetime `'b`, and as such, a `match` should perform the same function as `mem::transmute`.

However, the code produced by `match` is significantly more complex than the code produced by `mem::transmute`:

    use std::marker::PhantomData;
    use std::mem;
    
    pub enum Foo&lt;'a&gt; {
        Bar(usize),
        Baz(u32),
        Qux(String),
        Quux(Vec&lt;u8&gt;),
        Quuz(PhantomData&lt;&amp;'a str&gt;),
    }
    
    impl&lt;'a&gt; Foo&lt;'a&gt; {
        pub fn match_to_static(self) -&gt; Foo&lt;'static&gt; {
            match self {
                Foo::Bar(v) =&gt; Foo::Bar(v),
                Foo::Baz(v) =&gt; Foo::Baz(v),
                Foo::Qux(v) =&gt; Foo::Qux(v),
                Foo::Quux(v) =&gt; Foo::Quux(v),
                Foo::Quuz(..) =&gt; Foo::Quuz(PhantomData),
            }
        }
    
        pub fn transmute_to_static(self) -&gt; Foo&lt;'static&gt; {
            unsafe { mem::transmute(self) }
        }
    }

`transmute_to_static` is more verbose than I expected but still compact:

    example::Foo::transmute_to_static:
            mov     rax, rdi
            movups  xmm0, xmmword ptr [rsi]
            movups  xmm1, xmmword ptr [rsi + 16]
            movups  xmmword ptr [rdi + 16], xmm1
            movups  xmmword ptr [rdi], xmm0
            ret

When inlined, it becomes a no-op afaik.

However, `map_to_static` produces a lot of code and even a jump table, even on `-C opt-level=3`.  It's probably too long to paste here, so here's a compiler explorer link: https://godbolt.org/z/v44kdG

Is `mem::transmute` doing something evil here, or is the optimizer just not noticing what the `match` is doing?
## [11][Rx 5600 xt problem in Rust.](https://www.reddit.com/r/rust/comments/gk95cq/rx_5600_xt_problem_in_rust/)
- url: https://www.reddit.com/r/rust/comments/gk95cq/rx_5600_xt_problem_in_rust/
---
I just got an rx 5600 xt, and it runs really well when playing Rust. But when I enter the inventory or the menu it seems to be running in 60 hz. But in the game itself, it runs with 144 hz. Does anyone know why?
## [12][[Noob question] Cannot assign to reference](https://www.reddit.com/r/rust/comments/gk46eq/noob_question_cannot_assign_to_reference/)
- url: https://www.reddit.com/r/rust/comments/gk46eq/noob_question_cannot_assign_to_reference/
---
\[Solved, make categories and products mutable\]

I've a product category which contains multiple products. The pictures of the products come from another source.

    pub fn get_categories() -&gt; Result&lt;Vec&lt;Category&gt;, String&gt;
    {
        let categories = get_all();
        
        for category in &amp;categories
        {
            for mut product in &amp;category.products
            {
                let picture_ids: Vec&lt;i32&gt; = ...
                product.picture_ids = picture_ids; // &lt;-- error
            }
        }
    
        return Ok(categories);
    }

This example will give an error `cannot assign to'product.pictures'which is behind a'&amp;'reference'product'is a'&amp;'reference, so the data it refers to cannot be written`. However when I remove the Result and directly return the categories array the error wont popup. How can I solve this error or what am I missing?
