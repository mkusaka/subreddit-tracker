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

## [3][How to build a new rust toolchain?](https://www.reddit.com/r/rust/comments/hapgv4/how_to_build_a_new_rust_toolchain/)
- url: https://www.reddit.com/r/rust/comments/hapgv4/how_to_build_a_new_rust_toolchain/
---
I want to build a rust toolchain for powerpc64le-unknown-linux-musl . What are the steps I need to take to achieve that?
## [4][Variadic functions can be defined in Rust Nightly (kinda!)](https://www.reddit.com/r/rust/comments/hakjqw/variadic_functions_can_be_defined_in_rust_nightly/)
- url: https://www.reddit.com/r/rust/comments/hakjqw/variadic_functions_can_be_defined_in_rust_nightly/
---
I saw someone post how you could do function overloading in nightly rust.

Well, similarly, variadic functions can work now too by using the `tuple_list` crate.

Example:

    #![feature(fn_traits, unboxed_closures)]
    
    use tuple_list::{TupleList,Tuple};
    
    trait sumTupleList: TupleList {
        fn sum(self) -&gt; i32;
    }
    // induction base case
    impl sumTupleList for () {
        fn sum(self)-&gt; i32 {0}
    }
    impl&lt;Tail&gt; sumTupleList for ( i32, Tail) where
        Self: TupleList,
        Tail: sumTupleList {
        fn sum(self) -&gt; i32 {
            self.0 + self.1.sum()
        }
    }
    struct Foo {}
    impl&lt;'a,T&gt; FnOnce&lt;T&gt; for Foo where
        T:  tuple_list::Tuple ,
        &lt;T as Tuple&gt;::TupleList: sumTupleList
    {
        type Output = i32;
        extern "rust-call" fn call_once(self, args:T ) -&gt; Self::Output {
            sumTupleList::sum(Tuple::into_tuple_list(args))
        }
    }
    const addup: Foo = Foo{};
    fn main() {
        assert!(addup(1,2,4,5,6,7)        == 25);
        assert!(addup(1,2,3)              ==  6);
        assert!(addup(1,2,4,5,6,7,8,9,10) == 52);
    }

It should be noted, the difference between `f` and just calling aliasing `sumTupleList::sum(Tuple::into_tuple_list(...))` is that in the latter, you would have to pass in a tuple  of your items.

(In general, if you  for `Fn_Once&lt;(a,b,c)&gt;`   for `b`; then you can pass an implementation of `b`  three parameters `a,b,c,` not one parameter tuple `(a,b,c).` )

It should also be noted that tuple\_list supports only up to 12 items in a tuple I believe.
## [5][What is the current story for compiling to C?](https://www.reddit.com/r/rust/comments/hapl7w/what_is_the_current_story_for_compiling_to_c/)
- url: https://www.reddit.com/r/rust/comments/hapl7w/what_is_the_current_story_for_compiling_to_c/
---
I need to write some code for a relatively uncommon microprocessor.  The manufacturer provides a proprietary C/C++ compiler.  Obviously, I'd prefer to write this code in Rust and use the manufacturer's compiler as the last step in my build somehow.

I know there used to be an LLVM C backend that would let you do this kind of thing, but as I recall it was removed from the LLVM project.  Have there been any recent developments to improve this situation?
## [6][3K, 60fps, 130ms: achieving it in Rust](https://www.reddit.com/r/rust/comments/ha6kje/3k_60fps_130ms_achieving_it_in_rust/)
- url: https://blog.tonari.no/why-we-love-rust
---

## [7][MeiliSearch v0.11 - Faceted search](https://www.reddit.com/r/rust/comments/hapeaw/meilisearch_v011_faceted_search/)
- url: https://blog.meilisearch.com/whats-new-in-v0-11/
---

## [8][WezTerm: A GPU-accelerated cross-platform terminal emulator and multiplexer implemented in Rust](https://www.reddit.com/r/rust/comments/hach2x/wezterm_a_gpuaccelerated_crossplatform_terminal/)
- url: https://wezfurlong.org/wezterm/
---

## [9][The Rust Programming Language book is part of the Humble Fight for Racial Justice Bundle](https://www.reddit.com/r/rust/comments/hadsoz/the_rust_programming_language_book_is_part_of_the/)
- url: https://www.humblebundle.com/fight-for-racial-justice-bundle
---

## [10][I'm making a mouseless Image editor with Rust. Open source, also documented in video format](https://www.reddit.com/r/rust/comments/haqrel/im_making_a_mouseless_image_editor_with_rust_open/)
- url: https://www.reddit.com/r/rust/comments/haqrel/im_making_a_mouseless_image_editor_with_rust_open/
---
I'm in the early stages of creating an Image Editor, inspired by VIM. The goal I'm aiming towards is to be able to create thumbnails fast and easily.

I document the process on my youtube channel: [I Made a Mouseless Image Editor With RUST](https://youtu.be/2cSY43OcuZc)

Here is my github page:  [https://github.com/TanTanDev/vimnail](https://github.com/TanTanDev/vimnail)

The project is in a very early stage, but if you have criticism surrounding the idea or the code, I highly appreciate it. Especially the code, I'm quite new to Rust!
## [11][Understanding the Rust ecosystem](https://www.reddit.com/r/rust/comments/hamjrm/understanding_the_rust_ecosystem/)
- url: https://joeprevite.com/rust-lang-ecosystem
---

## [12][Rust concurrency: the archetype of a message-passing bug.](https://www.reddit.com/r/rust/comments/han5h8/rust_concurrency_the_archetype_of_a/)
- url: https://www.reddit.com/r/rust/comments/han5h8/rust_concurrency_the_archetype_of_a/
---
A description of a common source of bugs in message-passing code, taking the Rust programming language as an example. 

[https://medium.com/@polyglot\_factotum/817b60efd8f8?source=friends\_link&amp;sk=ad32b77d42eda3dd28a26671282271df](https://medium.com/@polyglot_factotum/817b60efd8f8?source=friends_link&amp;sk=ad32b77d42eda3dd28a26671282271df)
