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

## [3][Introducing Tree-Buf](https://www.reddit.com/r/rust/comments/gud6c1/introducing_treebuf/)
- url: https://www.reddit.com/r/rust/comments/gud6c1/introducing_treebuf/
---
Tree-Buf is an experimental serialization system for data sets (not messages) that is on track to be the fastest, most compact self-describing serialization system ever made. I've been working on it for a while now, and it's time to start getting some feedback.

Tree-Buf is smaller and faster than ProtoBuf, MessagePack, XML, CSV, and JSON for medium to large data.

It is possible to read any Tree-Buf file - even if you don't have a schema.

Tree-Buf is easy to use, only requiring you to decorate your structs with \`#\[Read, Write\]\`

Even though it is the smallest and the fastest, Tree-Buf is yet un-optimized. It's going to get a lot better as it matures.

You can read more about how Tree-Buf works under the hood at this [README](https://github.com/That3Percent/tree-buf/blob/master/README.md).
## [4][Scary Acronyms (and Super Creeps) [Rust Wrocław Webinar, 27.05.2020]](https://www.reddit.com/r/rust/comments/gugi8b/scary_acronyms_and_super_creeps_rust_wrocław/)
- url: https://www.youtube.com/watch?v=6Qi5-VU-kS0
---

## [5][A New Better CSS Preprocessor](https://www.reddit.com/r/rust/comments/guf3ys/a_new_better_css_preprocessor/)
- url: https://www.reddit.com/r/rust/comments/guf3ys/a_new_better_css_preprocessor/
---
Hey everyone. My name is Arijit, for the past two montha I have been working on a new preprocessor for CSS. It is called Pbss. I started writing it in Python but now I have switched to Rust. I want everyone to please check it out on my Github and start contributing. Please check it out:- [Github](https://github.com/arijit79/pbss)

It currently supports the following
* Variables
* Including code from.other filea without editing HTML. It is similar to the use statement of Sass
* Error catching of common typos and lines
* Arithmetic

I request the community to please join and contribute on this project
## [6][Announcing: jlrs 0.5: Windows support, A more consistent API, custom derives for Julia tuples and structs, system image support, and more!](https://www.reddit.com/r/rust/comments/guih33/announcing_jlrs_05_windows_support_a_more/)
- url: https://www.reddit.com/r/rust/comments/guih33/announcing_jlrs_05_windows_support_a_more/
---
After a few weeks of work, a new version of jlrs has been released. This version introduces a few new features. Since the previous version Windows is officially supported (but I don't think I ever announced that).

First of all, the `TryUnbox`-trait has been removed in favour of `Cast`. This new trait works the same way as the old one, but can also be used to cast more complex data involving lifetimes. For example, if a function returns a `Module` or `Array` as a `Value`, it can be cast to these types. Similarly, `Value::is` now uses the trait `JuliaTypecheck` rather than `JuliaType` and supports a much larger range of checks; for example, you can check if the value is a mutable struct or a tuple, rather than just a handful of explicit types.  

If you were using jlrs to work with multidimensional arrays that borrowed or moved data from Rust, you needed to include `jlrs.jl` in order to create those arrays. This problem has been solved, the only things defined in there now are two functions that will let you print or acquire the backtrace if an exception is thrown. Another limitation that has been removed is the lack of support for working with arrays that contain bools or chars.

The major additions to this version are two custom derives, `JuliaStruct` and `JuliaTuple`, which can be used to work with Julia tuples and structs for which the condition `isbitstype(&lt;ty&gt;`) holds true in Julia. Deriving one of these traits provides the same features as supported by primitive types. 

If you were bothered by long compilation times in Julia, it is now possible to initialize Julia with a custom system image. These system images contain precompiled packages and functions which can get rid of the issue that Julia functions are slow the first time they're called.

What's next? A few things. First of all, I lost about a week of work on jl-sys because I used git submodules incorrectly, so most of that has to be redone. Additionally, I want to take a look at "partially inlined" `Value`s and take another shot at providing experimental async-await support.

[Github](https://github.com/Taaitaaiger/jlrs)

[Docs.rs](https://docs.rs/jlrs/0.5.0/jlrs/index.html)

[Crates.io](https://crates.io/crates/jlrs)
## [7][Rust Cookbook 2018 update](https://www.reddit.com/r/rust/comments/guaznl/rust_cookbook_2018_update/)
- url: https://www.reddit.com/r/rust/comments/guaznl/rust_cookbook_2018_update/
---
[https://rust-lang-nursery.github.io/rust-cookbook/intro.html](https://rust-lang-nursery.github.io/rust-cookbook/intro.html)

Shout out to [pickfire](https://github.com/pickfire) who helped the Rust Cookbook get updated to the 2018 version.  

I wrote about some of the problems we had completing this [http://www.yetanother.site/rust/2020/05/31/Rust-Cookbook-Reaches-2018.html](http://www.yetanother.site/rust/2020/05/31/Rust-Cookbook-Reaches-2018.html)
## [8][Oxidised eBPF II: Taming LLVM](https://www.reddit.com/r/rust/comments/guhvdn/oxidised_ebpf_ii_taming_llvm/)
- url: https://blog.redsift.com/labs/oxidised-ebpf-ii-taming-llvm/
---

## [9][Adding symbol stripping to Cargo](https://www.reddit.com/r/rust/comments/gu1b1d/adding_symbol_stripping_to_cargo/)
- url: https://blog.gabrielmajeri.ro/2020/05/31/adding-symbol-stripping-to-cargo.html
---

## [10][Concerns about input latency with Amethyst](https://www.reddit.com/r/rust/comments/gueafx/concerns_about_input_latency_with_amethyst/)
- url: /r/rust_gamedev/comments/gu8nko/id_like_to_try_out_amethyst_but_i_have_some/
---

## [11][Mutex in async world](https://www.reddit.com/r/rust/comments/guivuf/mutex_in_async_world/)
- url: https://kitsu.me/posts/2020_06_01_mutex_in_async_world
---

## [12][Self-referential structs, is there really no better way?](https://www.reddit.com/r/rust/comments/gudr22/selfreferential_structs_is_there_really_no_better/)
- url: https://www.reddit.com/r/rust/comments/gudr22/selfreferential_structs_is_there_really_no_better/
---
Hi, I've decided this spring to really dive deep into Rust and make a complete game with it. It took a while but I think I'm getting the hang of the borrow checker and lifetime analysis. Unfortunately, I'm finding that what seem to be very common, even "pedestrian" design cases, are difficult or impossible to model in Rust. I'll show you one example of what I mean, and perhaps someone can show me where I've gone wrong.

I'm using the SDL2 lib, and I'd like a data structure that caches all the different font textures, so that I don't have to recreate them every frame. It looks like this:

    pub struct Fonts&lt;'a&gt; {
        ttf_context: Sdl2TtfContext,
        cache: HashMap&lt;&amp;'static str, Font&lt;'a, 'static&gt;&gt;,
    }

What's bad about this, is that we have to use the lifetime `'a`, even though `'a` is guaranteed to be the same as the lifetime of the `ttf_context` (that's what the internal refs inside of Font points to), which Fonts already owns. I'm surprised there isn't a special lifetime (`'struct`), so I could write something like this:

    pub struct Fonts {
        ttf_context: Sdl2TtfContext,
        cache: HashMap&lt;&amp;'static str, Font&lt;'struct, 'static&gt;&gt;,
    }

Anyway, what I ended up doing was this crazy hack (notice the use of `std::mem::transmute` to trick the compiler into accepting my lifetimes):

    pub struct Fonts {
        ttf_context: Sdl2TtfContext,
        cache: HashMap&lt;&amp;'static str, Font&lt;'static, 'static&gt;&gt;,
    }
    
    impl Fonts {
        fn new() -&gt; Result&lt;Fonts, String&gt; {
            let ttf_context = sdl2::ttf::init().map_err(|e| e.to_string())?;
            let mut fonts = Fonts {
                    ttf_context,
                    cache: HashMap::new(),
            };
            Ok(fonts)
        }
    
        fn get_cached_font&lt;'a&gt;(&amp;'a mut self, name: &amp;'static str) -&gt; &amp;Font&lt;'a, 'static&gt; {
            if !self.cache.contains_key(name) {
                unsafe { 
                    let raw_font = Self::make_font(&amp;self.ttf_context, name);
                    // change Font&lt;'a, 'static&gt; to Font&lt;'static, 'static&gt;
                    let font = std::mem::transmute(raw_font); 
                    self.cache.insert(name, font);
                }
            }
            self.cache.get(name).unwrap()
        }
    
        fn make_font&lt;'a&gt;(ttf_context: &amp;'a Sdl2TtfContext, name: &amp;'static str) -&gt; Font&lt;'a, 'static&gt; {
             ttf_context.load_font(name, 128).unwrap()
        }
    }

Is there really no better way to do this? After years of rust usage and thousands of contributors, I would suspect that the notion of a "cache where the items in the cache all point at a central source" would be very common... and yet this abuse of "transmute" is the best solution I've found. Is there a better way to structure this problem, and my thinking is not "rusty" enough, or is there some third party crate that does exactly what I want?
