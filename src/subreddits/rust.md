# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (35/2020)!](https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 353](https://www.reddit.com/r/rust/comments/ih8uai/this_week_in_rust_353/)
- url: https://this-week-in-rust.org/blog/2020/08/26/this-week-in-rust-353/
---

## [3][Compiler regression on 1.46.0?](https://www.reddit.com/r/rust/comments/iingya/compiler_regression_on_1460/)
- url: https://www.reddit.com/r/rust/comments/iingya/compiler_regression_on_1460/
---
My code compiled and ran fine on rustc 1.45.2, however after getting 1.46.0, I'm getting this error.

    error: reached the type-length limit while instantiating `std::thread::LocalKey::&lt;std::cel...d::Box&lt;dyn std::error::Error&gt;&gt;&gt;&gt;`
       --&gt; /Users/asdf/.rustup/toolchains/stable-x86_64-apple-darwin/lib/rustlib/src/rust/src/libstd/thread/local.rs:235:5
        |
    235 | /     pub fn with&lt;F, R&gt;(&amp;'static self, f: F) -&gt; R
    236 | |     where
    237 | |         F: FnOnce(&amp;T) -&gt; R,
    238 | |     {
    ...   |
    242 | |         )
    243 | |     }
        | |_____^
        |
        = note: consider adding a `#![type_length_limit="1871104"]` attribute to your crate

After some small code changes, it then gives:

    error: reached the type-length limit while instantiating `&lt;std::slice::Iter&lt;std::string::S...r str) -&gt; std::string::String]]&gt;`
        --&gt; /Users/asdf/.rustup/toolchains/stable-x86_64-apple-darwin/lib/rustlib/src/rust/src/libcore/iter/traits/iterator.rs:2015:5
         |
    2015 | /     fn fold&lt;B, F&gt;(mut self, init: B, mut f: F) -&gt; B
    2016 | |     where
    2017 | |         Self: Sized,
    2018 | |         F: FnMut(B, Self::Item) -&gt; B,
    ...    |
    2024 | |         accum
    2025 | |     }
         | |_____^
         |
         = note: consider adding a `#![type_length_limit="1863488"]` attribute to your crate

I have not been able to find the cause the problem. Is this a regression? And if so, how do I communicate it to the rust team?
## [4][Dwarf Seeks Fortune - 2D platformer game, homage to King's Valley II](https://www.reddit.com/r/rust/comments/iir6qo/dwarf_seeks_fortune_2d_platformer_game_homage_to/)
- url: https://amethyst.rs/posts/showcase-game-5-dwarf-seeks-fortune
---

## [5][rust-analyzer financial report](https://www.reddit.com/r/rust/comments/iianoy/rustanalyzer_financial_report/)
- url: https://rust-analyzer.github.io/blog/2020/08/20/financial-report.html
---

## [6][Rust serialization: What's ready for production today?](https://www.reddit.com/r/rust/comments/iif7hc/rust_serialization_whats_ready_for_production/)
- url: https://blog.logrocket.com/rust-serialization-whats-ready-for-production-today/
---

## [7][how to generate files from procedural macros?](https://www.reddit.com/r/rust/comments/iiolxy/how_to_generate_files_from_procedural_macros/)
- url: https://www.reddit.com/r/rust/comments/iiolxy/how_to_generate_files_from_procedural_macros/
---
hi,

i am using proc macro to generate encode/decode functions for a struct.

this is done using syn and quote crates, and works out well. I get a .so library after building and it can be linked against from C code.

but i also need to generate C header files. i heard that writing a file from proc macro code is a bad idea.

so how to get it done?
## [8][embedded-time v0.10.0 Release](https://www.reddit.com/r/rust/comments/iiiuwo/embeddedtime_v0100_release/)
- url: https://www.reddit.com/r/rust/comments/iiiuwo/embeddedtime_v0100_release/
---
One of the features that has been conspicuously missing is multiplying and dividing *duration*s and *rate*s by an integer. This version adds both operator mul/div and *checked* mul/div.

Coming from C++, I still miss the implicit widening of integers. I decided to incorporate some implicit widening and worked very well, except for the huge (relatively speaking) additional code needed for wider-than-32-bits integers. I have reverted that implicit widening in this version and used an example binary and the cargo-bloat command to verify the huge space savings. I now have a much more favorable view of the strict type behavior of Rust.

See [https://crates.io/crates/embedded-time](https://crates.io/crates/embedded-time) for more information about `embedded-time` if you're interested.
## [9][Best library for lossy image compression?](https://www.reddit.com/r/rust/comments/iiojju/best_library_for_lossy_image_compression/)
- url: https://www.reddit.com/r/rust/comments/iiojju/best_library_for_lossy_image_compression/
---
For my Rust web server I'm looking to compress uploaded images. I looked around and found [oxipng](https://github.com/shssoichiro/oxipng) but it seems to be lossless only.

It seems that lossy compression is pretty good these days. For example, https://tinypng.com (which uses quantization via [pngquant](https://github.com/kornelski/pngquant)) can often achieve compression rates of 70% for the images I tested, compared to the 10% I was getting with oxipng, all while visually looking very close to the original image.

I'd just use pngquant (as they appear to have Rust bindings) but it seems like it forces your project to either be GPL or pay a $&gt;1000 / year license fee, which I cannot afford.

Are there any other recommendations for image compression libraries that can be called from Rust, achieve high compression rates, and have an open license such as MIT or Apache?
## [10][Lifetime issues with rust-sfml](https://www.reddit.com/r/rust/comments/iip5jv/lifetime_issues_with_rustsfml/)
- url: https://www.reddit.com/r/rust/comments/iip5jv/lifetime_issues_with_rustsfml/
---
Hey guys. I'm using rust-sfml for rendering in a project of mine. It's fine for the most part, but I'm finding one thing very hard to solve.

In rust-sfml, the RectangleShape struct has the following (very simplified) form:

    struct RectangleShape&lt;'s&gt; {
        ...
        fn set_texture(&amp;mut self, texture: &amp;'s Texture, reset_rect: bool) {... }
        ...
    }

Note the lifetime on the texture reference. In my project, things like resources are handled using Arcs. A custom widget I'm trying to design has (basically) the following form:

    pub struct RectangularPanel&lt;'s&gt; {
        shape: RectangleShape&lt;'s&gt;,
        texture: Arc&lt;Texture&gt;,
    }

When I try to do the simplest example - i.e. hardcoded, not async loading, non-optional - the entire system shits itself due to lifetime issues. When I instantiate the shape in the RectangularPanel struct, I need to provide a texture reference lifetimed to 's , but suppose I try something very brute force like this:

    pub fn new(widget_table: &amp;Table) -&gt; Result&lt;Self&gt; {
        let handle = Arc::new(Texture::from_file("hardcoded_filename.png").unwrap());
        let mut panel = RectangularPanel {
            shape: RectangleShape::new(),
            texture: handle,
        };
        panel.shape.set_texture(&amp;panel.texture, false);
        Ok(panel)
    }

This doesn't work because I'm returning panel, which references borrowed data. No matter what I do, I can't get the compiler to agree that the texture will be valid for the whole lifetime of the RectangleShape. But I'm pretty sure this lifetime relationship is valid, since because both the RectangleShape and the Arc are owned by the RectangularPanel struct, the shape can always access that texture.

Is there any way to resolve this type of issue? I'm still a bit of a beginner so it's possible my design ideas just suck. It's also possible it's not resolvable since SFML is still chained to C-like ways of doing things.
## [11][#[const_fn] : An attribute for easy generation of const functions with conditional compilations.](https://www.reddit.com/r/rust/comments/iirb88/const_fn_an_attribute_for_easy_generation_of/)
- url: https://github.com/taiki-e/const_fn
---

## [12][Looking for input on emulating my work's C++ framework](https://www.reddit.com/r/rust/comments/iiprte/looking_for_input_on_emulating_my_works_c/)
- url: https://www.reddit.com/r/rust/comments/iiprte/looking_for_input_on_emulating_my_works_c/
---
I'm looking to emulate parts of my work's C++ codebase. Partially as a learning experience, and hopefully make a case to start shifting some work towards Rust. The latter is more of a long-term goal, so C++ integration/API is not really the focus yet.

The C++ applications revolve around spinning up multiple threads connected by single-producer, single-consumer channels. Work loads typically resemble `process(vec&lt;T&gt; x) -&gt; vec&lt;R&gt;` and therefore channel elements are always a vector of some kind i.e. `channel&lt;vec&lt;T&gt;&gt;`. In order to save on re-allocation the vectors get re-used, and they only re-allocate when the capacity needs to grow (there is no shrinking allowed). This is possible because the work that gets done is pretty rinse repeat, so once these vectors have been resized the first time, they pretty much stay the same size for the next. Effectively, each channel is actually a circular-buffer of vectors. Vector sizes vary by work type, but typically range from 10k to 1M elements of basic numeric types.

I don't think such an approach is well suited to Rust's ownership rules (?). And honestly, I dislike using our current C++ framework (even though it is very performant), and it removes the ability to perform work in-place on the current vector i.e. the output vector cannot be the input vector.

I was thinking an alternative is using normal channels and just ignoring allocation cost until it becomes a problem. Then mixing in some form of custom allocator/arena instead?

Any prior art, existing crates or alternative ideas are welcome.
