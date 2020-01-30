# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (5/2020)!](https://www.reddit.com/r/rust/comments/eupam3/hey_rustaceans_got_an_easy_question_ask_here_52020/)
- url: https://www.reddit.com/r/rust/comments/eupam3/hey_rustaceans_got_an_easy_question_ask_here_52020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (5/2020)?](https://www.reddit.com/r/rust/comments/eupbbn/whats_everyone_working_on_this_week_52020/)
- url: https://www.reddit.com/r/rust/comments/eupbbn/whats_everyone_working_on_this_week_52020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-5-2020/37515?u=llogiq)!
## [3][In the last year or so, r/rust gained about as many subscribers as it gained in the preceding six years.](https://www.reddit.com/r/rust/comments/ew1i8w/in_the_last_year_or_so_rrust_gained_about_as_many/)
- url: https://www.reddit.com/r/rust/comments/ew1i8w/in_the_last_year_or_so_rrust_gained_about_as_many/
---
From my occasional visits to this subreddit, it seems like there is more "average programming subreddit shenanigans" lately, like hostility toward competing languages, and "fights" in the comments. I don't have much experience with community building, but I hope the mods and citizens can manage to keep Rust's early community spirit alive :)

Data source: [https://subredditstats.com/r/rust](https://subredditstats.com/r/rust)
## [4][GitHub - oreboot/oreboot: oreboot is a fork of coreboot, with C removed, written in Rust.](https://www.reddit.com/r/rust/comments/ew1nje/github_orebootoreboot_oreboot_is_a_fork_of/)
- url: https://github.com/oreboot/oreboot
---

## [5][Lewton 0.10.0 release](https://www.reddit.com/r/rust/comments/ew0pah/lewton_0100_release/)
- url: https://www.reddit.com/r/rust/comments/ew0pah/lewton_0100_release/
---
I'm pleased to announce the release of [lewton](https://github.com/RustAudio/lewton) 0.10.0 - the pure Rust vorbis decoder.

The main changes since the last [reddit-announced release 0.9.1](https://www.reddit.com/r/rust/comments/9i3hu0/lewton_091_released/) of about a year ago ([link to full changelog](https://github.com/RustAudio/lewton/blob/master/CHANGELOG.md#release-0100---january-30-2020)):

* Fix of wrongfully decoded files ([#24](https://github.com/RustAudio/lewton/issues/24), [#26](https://github.com/RustAudio/lewton/issues/26))
* Fixes of bugs on fuzzed inputs, thanks to /u/shnatsel who's [bug reports](https://github.com/RustAudio/lewton/issues?q=is%3Aissue+author%3AShnatsel) were an awesome contribution to lewton (as are their bug reports to the greater rust ecosystem)
* Ability to specify floats as output format. Lewton uses floats internally but used to convert them to i16 unconditionally for the output. Now you can choose to get the floats directly without having to convert anything. See the `read_dec_packet_generic` and `read_audio_packet_generic` functions.
* Addition of a C API. This allows easier inclusion into C/C++ projects. Thanks to /u/lu_zero for their help.
* Increase of the MSRV from 1.20.0 to 1.36.0

Lewton has been getting lots of PRs in the last year. I want to thank the contributors to lewton since the 0.9.1 release:

* [@Luni-4](https://github.com/Luni-4)
* [@lu-zero](https://github.com/lu-zero)
* [@AnthonyMikh](https://github.com/AnthonyMikh)
* [@repi](https://github.com/repi)
* [@lovesegfault](https://github.com/lovesegfault)
* [@BenSandeen](https://github.com/BenSandeen)
* [@Shnatsel](https://github.com/Shnatsel)
## [6][self_update: Self updates for rust executables](https://www.reddit.com/r/rust/comments/ew3f78/self_update_self_updates_for_rust_executables/)
- url: https://github.com/jaemk/self_update
---

## [7][This Week in Rust 323](https://www.reddit.com/r/rust/comments/ew3gri/this_week_in_rust_323/)
- url: https://this-week-in-rust.org/blog/2020/01/28/this-week-in-rust-323/
---

## [8][A new useful crate for chemistry](https://www.reddit.com/r/rust/comments/evzfem/a_new_useful_crate_for_chemistry/)
- url: https://www.reddit.com/r/rust/comments/evzfem/a_new_useful_crate_for_chemistry/
---
Hi everyone. Almost a week ago I published a periodic table here, that was my first crate.

And now I have a crate more called [chembasics](https://crates.io/crates/chembasics) which provides other basic things for chemistry such as some constants, support for empirical and molecular formulas in addition to a SMILES interpreter (chiralide is lack for now, I plan to implement it but pull requests are welcome).

For now it is only that although I plan to implement the ability to graph structures in both 3d and 2d (and a raytracer as an optional feature).

Contributions are welcome and I would also like to receive some suggestions on how to handle the isotopes (I need it for the SMILES interpreter) so far I have not decided on how I should integrate them since I do not want to have to be writing something like Isotope::new(Element::Carbon, 12) every time I want to create a carbon atom.
## [9][Ferrugo: a JVM implementation written in Rust](https://www.reddit.com/r/rust/comments/evnv8z/ferrugo_a_jvm_implementation_written_in_rust/)
- url: https://github.com/maekawatoshiki/ferrugo
---

## [10][Hypothetically, if someone creates a tracing GC pointer, how would it affect implementing traditional Graph-shaped objects?](https://www.reddit.com/r/rust/comments/ew3zvb/hypothetically_if_someone_creates_a_tracing_gc/)
- url: https://www.reddit.com/r/rust/comments/ew3zvb/hypothetically_if_someone_creates_a_tracing_gc/
---
Lifetime, single-ownership and RAII based memory management are what Rust so unique from all the mainstream programming languages. There is no doubt that it helps reduce a big number of potential bugs in compile-time.

However, for general purpose programming, perhaps you do not require real-time performance. There are people who would like to use Rust not as a real-time systems language --some may want to use Rust for a CLI tool, or a GUI backend. Many people love Rust not only for the safety-plus-performance-in-compilitime aspect of Rust. 

My question is, how pragmatic would it be to adopt a pluggable GC from the third party -- e.g., a third party GC pointer. Is there one already?
## [11][evcxr/evcxr_jupyter: Jupyter kernel for rust](https://www.reddit.com/r/rust/comments/evrexn/evcxrevcxr_jupyter_jupyter_kernel_for_rust/)
- url: https://github.com/google/evcxr/tree/master/evcxr_jupyter
---

## [12][Generic types as type parameters](https://www.reddit.com/r/rust/comments/ew023l/generic_types_as_type_parameters/)
- url: https://www.reddit.com/r/rust/comments/ew023l/generic_types_as_type_parameters/
---
I am trying to solve an issue similar to [this one](https://www.reddit.com/r/rust/comments/9oa394/passing_generic_type_as_type_parameters/) where OP was trying to have something like:

    struct MyStruct&lt;K, V, M&gt; { map: M&lt;K,V&gt;; }

Now it _is_ possible to write

    struct MyStruct&lt;K, V, M&gt; where M: Something&lt;K, V&gt; {
        map: M&lt;K, V&gt;,
        _k: PhantomData&lt;K&gt;,
        _v: PhantomData&lt;V&gt;,
        ...
    }

but that leads to repetitive type signatures like `MyStruct&lt;i32, f32, Map&lt;i32, f32&gt;&gt;`. I came up with the following trick involving associated types and dummy structs:

    trait MapType {
        type Key;
        type Value;
        type Map;
    }

    struct MyMapType();
    impl MapType for MyMapType {
        type Key = i32;
        type Value = f32;
        type Map = HashMap&lt;Self::Key, Self::Value&gt;;
    }

    struct MyStruct&lt;MT&gt; where MT: MapType {
        map: MT::Map
    }

and then I can just write `MyStruct&lt;MyMapType&gt;`. Considering that in my case the same signature would be all over the place, it's even a win for readability. Still, I wonder if this is a good idea, and if so, whether it already has a name.
