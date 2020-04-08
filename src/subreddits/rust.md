# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (15/2020)!](https://www.reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (15/2020)?](https://www.reddit.com/r/rust/comments/fw2i84/whats_everyone_working_on_this_week_152020/)
- url: https://www.reddit.com/r/rust/comments/fw2i84/whats_everyone_working_on_this_week_152020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-15-2020/40539?u=llogiq)!
## [3][Invisible performance wins](https://www.reddit.com/r/rust/comments/fwynp1/invisible_performance_wins/)
- url: https://www.reddit.com/r/rust/comments/fwynp1/invisible_performance_wins/
---
**I want to praise those who work on Rust and LLVM compilers to make our Rust applications even faster to execute!**

I have [a benchmark](https://github.com/frol/completely-unscientific-benchmarks) which my friend and I started to compare a few languages we were interested in (the list is kind of weird: Rust, Kotlin Native, C++, and Python), but once I [published](https://www.reddit.com/r/rust/comments/8jbjku/naive_benchmark_treap_implementation_of_c_rust/) the benchmark to the wild, we received over 40 solutions and now the benchmark has implementations in 24 languages (+4 in PRs). It was 2 years ago, so I used Rust 1.26 back then, and yesterday I decided to re-run the whole benchmark with the recent versions of compilers... Long story short, Rust 1.42 improved the execution time of the same source code by 5-17%! *I want to say this again, we get 5-17% performance win from just compiler version upgrade!*

P.S. Here are some takeaways from this benchmark update:

* With the above optimizations, Rust is now faster than the algorithmically similar implementations in C++ (the matching pairs are: "Rust naive ref-counted" vs "C++ naive java-like", "Rust tuned safe mem::forget" vs "C++ tuned unique\_ptr"). Also, Rust naive idiomatic version is faster than naive "idiomatic" \[I have a hard time defining this term for C++\] solution in C++ (called "C++ naive shared\_ptr"); though, C++ solution has to use defensive programming (using \`shared\_ptr\`) while Rust solution just relies on compile-time ownership checks.
* Most of the compilers (for other languages) did not impact the performance of this unscientific benchmark in the course of the last two years; C++ is still the leader of the naive (and tuned) scoreboards, though I still argue that even though it used to be common to use raw pointers, it is not "idiomatic" C++ nowadays (gcc and clang manage to optimize tail recursion if you pass a mutable out parameter instead of returning a value), and given the previous paragraph, we just need to have algorithmically matching implementations in Rust to be on par with C++.
* The noticeable improvements were in Python (update from 3.5 to 3.8 get 19% improvement, yet 48 times slower than the best naive solution in C++) and Kotlin Native (update from 0.7 to 1.3.71 yield 48% improvement, so it is twice as fast as it used to be, but it is still 16 times slower than the best naive solution)

P.P.S. I wish Rust can optimize the idiomatic version down to the one that uses a simple trick of \`swap\` + \`mem::forget\`. I even tried to pull this myself but failed. Here are some references if you are interested:

* [https://barrielle.cedeela.fr/research\_page/dropping-drops.html](https://barrielle.cedeela.fr/research_page/dropping-drops.html)
* [https://github.com/rust-lang/rfcs/pull/1736](https://github.com/rust-lang/rfcs/pull/1736)
* [https://github.com/rust-lang/rfcs/pull/2490](https://github.com/rust-lang/rfcs/pull/2490) (my duplicate of the #1736)

&amp;#x200B;

P.P.P.S Thank you, Rust designers, developers, and community! You rock!
## [4][Nine patches](https://www.reddit.com/r/rust/comments/fx13f1/nine_patches/)
- url: https://blog.yoshuawuyts.com/nine-patches/
---

## [5][Rust Tutorial: Building a Ray Tracer in One Weekend](https://www.reddit.com/r/rust/comments/fx2jyi/rust_tutorial_building_a_ray_tracer_in_one_weekend/)
- url: https://www.youtube.com/watch?v=_5hD0gxRzzg&amp;list=PLctja8lh-0MdlbFXjmBoFn_1hiK4qkVqe
---

## [6][Error Handling in a Large Correctness-Critical Rust Project](https://www.reddit.com/r/rust/comments/fx5ibw/error_handling_in_a_large_correctnesscritical/)
- url: http://sled.rs/errors
---

## [7][A brief apology of Ok-Wrapping](https://www.reddit.com/r/rust/comments/fwpgsh/a_brief_apology_of_okwrapping/)
- url: https://boats.gitlab.io/blog/post/why-ok-wrapping/
---

## [8][How to share the same dependency build artifacts between different rust projects properly](https://www.reddit.com/r/rust/comments/fx47z7/how_to_share_the_same_dependency_build_artifacts/)
- url: https://www.reddit.com/r/rust/comments/fx47z7/how_to_share_the_same_dependency_build_artifacts/
---
Lets say I am going to build 3 separate rust projects. I don't want to build dependencies for each of them every time I build them, I just want to complile my source code.

How can I copy all the build artifacts generated after `cargo build release` in the target directory so that lets say I have built my project A for the first time and all the deps have been built  and now when I go to build project B and C they simply lookup the copies of those build artifacts and just compile the source code instead of all the deps all over again.

What is the right way to do this?
## [9][Announcing Cooked Wakers 2.0! Now with 100% less unnecessary derives](https://www.reddit.com/r/rust/comments/fwstbs/announcing_cooked_wakers_20_now_with_100_less/)
- url: https://www.reddit.com/r/rust/comments/fwstbs/announcing_cooked_wakers_20_now_with_100_less/
---
Hey everyone! A couple months ago, I [published Cooked Wakers]( https://www.reddit.com/r/rust/comments/f05qiy/announcing_cooked_wakers/), a set of traits designed for making it easy to create `std::task::Waker` objects with 100% safe Rust and without any implementation overhead. One of the frustrating things about that design is that Rust's generic monomorphization rules made it impossible to generically create a `RawWakerVTable`; This meant that the `IntoWaker` trait could only ever be implemented on concrete types, which meant that I had to provide a complex derive for it, rather than a generic implementation. It also meant that downstream users had to create a lot of boilerplate types wrapping `Arc&lt;Notifier&gt;`.

Well, I'm excited to announce that I figured out how to [solve](https://github.com/Lucretiel/cooked-waker/blob/master/src/lib.rs#L193-L217) that issue using an associated `const` on the `IntoWaker` trait. This means that I was able to remove a huge amount of boilerplate from the library and release [version 2.0](https://crates.io/crates/cooked_waker)! Aside from the removal of all the derives, the major change in this library is that it is now possible to call `into_waker` directly on shared pointer types, like so:

    // OLD VERSION /////////////////////////////////////////////////
    struct MyNotifier { ... }

    // Because IntoWaker requires a concrete type, we have to wrap
    // Arc in a struct.
    #[derive(Debug, Clone, WakeRef, Wake, IntoWaker)]
    struct NotifierHandle {
        handle: Arc&lt;MyNotifier&gt;
    }

    let handle = NotifierHandle::new(Arc::new(MyNotifier{...}));
    let waker = handle.clone().into_waker();

    // NEW VERSION ////////////////////////////////////////////////////
    struct MyNotifier { ... }

    // Arc now implements IntoWaker directly!
    let handle = Arc::new(MyNotifier{...}));
    let waker = handle.clone().into_waker();

Like before, this library is based on [`stowaway`](https://docs.rs/stowaway/1.1.2/stowaway/), a crate I designed to pack a `struct` into a `*mut ()` as efficiently as possible, only boxing it if it's larger than a pointer. By using this, plus generic function pointers in the VTable, it allows the creation of wakers that are exactly as efficient as if you had implemented them by hand, without any `unsafe` on your part.

Cooked Wakers: because consuming unsafe RawWakers is bad for your health.

As always, feedback and pull requests are welcome.
## [10][Rust on GPU?](https://www.reddit.com/r/rust/comments/fx0tbt/rust_on_gpu/)
- url: https://www.reddit.com/r/rust/comments/fx0tbt/rust_on_gpu/
---
Hi everyone,

I would love to build in some functionality to my crate which enables it to optionality run on a GPU (shameless plug to [radiate](https://github.com/pkalivas/radiate) \- if you've used it or have experienced issues PLEASE lmk, it is in active development and can use as much input as possible - all comments are totally welcome and will not hurt my feelings). It is a genetic programming engine with an implementation of [NEAT](http://nn.cs.utexas.edu/downloads/papers/stanley.ec02.pdf) with a few personal tweaks and additions to it - check out the above github for more information. LSTM, GRU, traditional neural net training, etc. NEAT Has definitely taken a front seat to development as I've been using it personally, so any type of feedback on the engine or the models is incredibly welcome and appreciated.

But, long story short, I've been using this crate for some personal projects and although it meets my needs, using the engine on the CPU for large problems can take a long time and cause my CPU to burn quite hot (rayon::ThreadPoolBuilder only helps so much). In terms of ML and AI, although I'm not aiming for this crate to be the 'gold standard' by any means, I would like it to be a viable solution for people and would like it to be a professional level crate (I graduated from college a year ago and have only been working in the software industry since Sep '19, so again, any advice or constructive criticism on the crate is 100% welcome and encouraged.). I think there are areas of the crate which can definitely be improved - using the GPU for training being one of them.

I know for languages like Python there are mature libraries for executing code on the GPU through libraries like CUDA. Does something along the same lines exist for rust, and if so could anyone point me in that direction? I've found a few abandoned GPU crates that are a few years old and would like to either use or contribute to one which would be (somewhat) future-proof. 

Thanks.
## [11][Futures Explained in 200 Lines of Rust](https://www.reddit.com/r/rust/comments/fwiet3/futures_explained_in_200_lines_of_rust/)
- url: https://cfsamson.github.io/books-futures-explained/
---

## [12][async-postgres: a runtime-independent, asynchronous PostgreSQL client.](https://www.reddit.com/r/rust/comments/fwyb41/asyncpostgres_a_runtimeindependent_asynchronous/)
- url: https://github.com/Hexilee/async-postgres
---

