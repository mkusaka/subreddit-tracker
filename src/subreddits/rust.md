# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (23/2020)!](https://www.reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gr3r32/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (23/2020)?](https://www.reddit.com/r/rust/comments/guo51x/whats_everyone_working_on_this_week_232020/)
- url: https://www.reddit.com/r/rust/comments/guo51x/whats_everyone_working_on_this_week_232020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-23-2020/43609?u=llogiq)!
## [3][What features would you like to see in a Rust IDE?](https://www.reddit.com/r/rust/comments/gvnwcx/what_features_would_you_like_to_see_in_a_rust_ide/)
- url: https://users.rust-lang.org/t/what-features-would-you-like-to-see-in-a-rust-ide/43697
---

## [4][Rust seems like the most consistent language I've ever used? What's inconsistent about it from your experience?](https://www.reddit.com/r/rust/comments/gvqhea/rust_seems_like_the_most_consistent_language_ive/)
- url: https://www.reddit.com/r/rust/comments/gvqhea/rust_seems_like_the_most_consistent_language_ive/
---
Almost every language I've encountered there have been plenty of questionable (but understandable) decisions (cpp no ABI break , vector&lt;bool&gt;, pythons GIL, java type erasure on generics, etc.). I've yet to encounter a single issue like that in Rust. Everything about it seems consistent and have one and only one way to do it and if there is another way to do something, it seems fulfills a very specific purpose.

To be honest, I'm almost certain there must be some problems like the ones mentioned earlier though so I'd like to hear them.
## [5][This is a neat trick for getting good runtime performance with Development mode builds](https://www.reddit.com/r/rust/comments/gvrgca/this_is_a_neat_trick_for_getting_good_runtime/)
- url: https://www.reddit.com/r/rust/comments/gvrgca/this_is_a_neat_trick_for_getting_good_runtime/
---
You can tell Cargo to compile all your dependencies with full optimisation (like in Release mode), while compiling your own code with few optimisations (like in Development mode). Since your dependencies change rarely you can take the hit of compiling them in Release mode once without affecting your day-to-day workflow much. Potentially this can give you the best of both worlds - fast run time and fast(ish) compile times.

    [profile.dev.package."*"]
    # Set the default for dependencies in Development mode.
    opt-level = 3

    [profile.dev]
    # Turn on a small amount of optimisation in Development mode.
    opt-level = 1

Of course the effectiveness will depend heavily on your own individual program, but worth trying I think.

As the [Cargo profiles](https://doc.rust-lang.org/cargo/reference/profiles.html#overrides-and-generics) page shows, if you make heavy use of generics from 3rd party crates then you won't get the full benefit, but in most cases I would not be surprised if 90% of the Rust code in your program is actually from crates and can benefit from this. It certainly worked fantastically well on a game I was writing.
## [6][Parcel and Rust: A WASM Romcom](https://www.reddit.com/r/rust/comments/gvqje8/parcel_and_rust_a_wasm_romcom/)
- url: https://dev.p.ota.to/post/parcel-and-rust-a-wasm-romcom-4kgsjnrdm5t/
---

## [7][DOOM Fire implemented in rust](https://www.reddit.com/r/rust/comments/gvcj6d/doom_fire_implemented_in_rust/)
- url: https://www.reddit.com/r/rust/comments/gvcj6d/doom_fire_implemented_in_rust/
---
I was looking for a simple project to try out different graphics libraries from the rust ecosystem.

I ended up implementing the DOOM fire effect (based on [Fabien Sanglard's blog post](https://fabiensanglard.net/doom_fire_psx/)) and tried it using different libraries:

* [minifb](https://github.com/emoon/rust_minifb)
* [pixels](https://github.com/parasyte/pixels)
* [sdl2 bindings](https://github.com/Rust-SDL2/rust-sdl2)

The code can be found [here](https://github.com/r-marques/doomfire)

This could be helpful for someone new to rust and trying to get into game development and looking for the right libraries to use.
## [8][Why F#, Rust and Others Use Option Type Instead Of Nullable types like C# 8 Or TypeScript?](https://www.reddit.com/r/rust/comments/gvjccd/why_f_rust_and_others_use_option_type_instead_of/)
- url: https://softwareengineering.stackexchange.com/questions/410724/why-f-rust-and-others-use-option-type-instead-of-nullable-types-like-c-8-or-t
---

## [9][A new Wasm/rust newsletter. Focus on the server-side.](https://www.reddit.com/r/rust/comments/gvqdwq/a_new_wasmrust_newsletter_focus_on_the_serverside/)
- url: /r/WebAssembly/comments/gv01ds/a_new_webassembly_newsletter_focus_on_the/
---

## [10][Rust is amazing! (a C++ developer's experience writing a raytracer in Rust)](https://www.reddit.com/r/rust/comments/gv2act/rust_is_amazing_a_c_developers_experience_writing/)
- url: https://www.reddit.com/r/rust/comments/gv2act/rust_is_amazing_a_c_developers_experience_writing/
---
Background: I've been writing C++ professionally since 2011, and I started to learn Rust last fall by reading a couple of books and solving some Advent of Code 2019 problems (finished 1-11 except 10 in Rust).  I also did the [Ray Tracing in One Weekend](https://raytracing.github.io/) tutorial in both Rust and Go, and just finished the 2nd "Ray Tracking: The Next Week" part of the series in Rust.

Overall, Rust has been an awesome language to learn, and I think it solves many of the common issues/bugs in C++ (which is kind of the whole point).  The concept of lifetimes and the borrow checker really enhances memory safety… our C++ code frequently takes pointers in class constructors and stores them as member variables, and proper lifetime management is totally up to the user (with some `// The object pointed to by this argument must outlive this class.` comments to help).  Many of the bugs we investigate at work are off-by-one or out of bounds accesses as well, so also eliminating those is a huge plus.  I find the if expressions, `match`, implicit returns, and the `?` operator all quite useful, and they help make your code more concise and readable.  I easily managed to get code for the first part (In One Weekend) of the series to run in WebAssembly, although Rayon doesn't work there and it's relatively slow.  It's cool to be able to visualize the ray tracing as it happens, though (I render 1 full frame, display it, render another frame, average it with all the previous ones, display it, and so on, so the image progressively becomes less and less noisy).

I've also found the ecosystem to be quite pleasant to work with, with trivial parallelism available with Rayon (it was super easy to parallelize my raytracer with a simple `into_par_iter`), useful crates for reading .jpg and writing .png files, and good benchmarking/profiling support (`cargo bench`, easy flamegraphs, and the ability to use the standard `perf` tool on binaries).  Performance is also really good… much, much faster than my Go version, and the code is also more readable, thanks to operator overloading.  The compiler and clippy both give really useful hints/tips, and usually make it easy to understand when you're doing something wrong.  I also found code completion in Vim and Visual Studio Code to be dead simple to set up too.

The few minor complaints I have are the lack of inheritance and function overloading, and no implicit type conversions from int to float (you frequently have to do this when using for loop indices in some computation, so I have `i as f32` all over the place).  It's also annoying to have to type 1.0 or 1. instead of simply 1, especially when you're trying to copy and paste some C++ example scene setup.  Compilation is also a bit slow at times (I think using the "image" crate is overkill and pulls in a bunch of dependencies, so I should probably switch to "jpeg_decoder").

[Here's a picture of the final scene from "The Next Week"](https://i.imgur.com/P7e61lb.jpg), which took almost 343 hours of CPU time to render (5120x2880 with 10,000 samples per pixel).  It took 3 hours 37 minutes of wall clock time on a 96 core Google Cloud instance.

Conclusion: Rust is awesome, and I hope to do some more projects in it in the future!                                                                                                                             
                                                                                                                                                                
(I'd post the code, but my employer requires us to go through an open source release approval process to release even personal projects... might do it someday)
## [11][(help) Type error and Result](https://www.reddit.com/r/rust/comments/gvtuj6/help_type_error_and_result/)
- url: https://www.reddit.com/r/rust/comments/gvtuj6/help_type_error_and_result/
---
I'm a new rust programmer and I thought I'd start a project to learn it a bit better. I'm trying to create a header structure but when implementing a trait function I get an error that I dont understand how to resolve.

This is a function in my trait. 

        fn get_e_entry_value(&amp;self) -&gt; Result&lt;u32, u64&gt; {
            match self.e_entry {
                Addr::Elf32(e_entry) =&gt; Ok(e_entry),
                Addr::Elf64(e_entry) =&gt; Ok(e_entry),
            }
        }

when compilingwith the above I get the following error

    error[E0308]: mismatched types
       --&gt; src/elf/structures.rs:258:40
        |
    258 |             Addr::Elf64(e_entry) =&gt; Ok(e_entry),
        |                                        ^^^^^^^ expected `u32`, found `u64`
        |
    help: you can convert an `u64` to `u32` and panic if the converted value wouldn't fit
        |
    258 |             Addr::Elf64(e_entry) =&gt; Ok(e_entry.try_into().unwrap()),
        |                                        ^^^^^^^^^^^^

Which is complaining that entry is expected to be a u32 even though the return type is actually a u64 inside the Result and the e\_entry type for the second match arm is a u64 (wrapped in an Ok).

So, the compiler knows its a u64 as evidenced by the 'found \`u64\`' statement. I have no idea what is going on. Can any of you help me figure out just what is going on?
## [12][Rust Community Wiki (very incomplete)](https://www.reddit.com/r/rust/comments/gv6ktp/rust_community_wiki_very_incomplete/)
- url: https://runrust.miraheze.org
---

