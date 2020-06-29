# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (27/2020)!](https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (27/2020)?](https://www.reddit.com/r/rust/comments/hhv5ks/whats_everyone_working_on_this_week_272020/)
- url: https://www.reddit.com/r/rust/comments/hhv5ks/whats_everyone_working_on_this_week_272020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-27-2020/45101?u=llogiq)!
## [3][#[feature(const_if_match)] has been stabilized](https://www.reddit.com/r/rust/comments/hhtl0q/featureconst_if_match_has_been_stabilized/)
- url: https://github.com/rust-lang/rust/pull/72437#issuecomment-650814084
---

## [4][TIL cargo has aliases like `cargo b` for `cargo build`](https://www.reddit.com/r/rust/comments/hhjvaj/til_cargo_has_aliases_like_cargo_b_for_cargo_build/)
- url: https://www.reddit.com/r/rust/comments/hhjvaj/til_cargo_has_aliases_like_cargo_b_for_cargo_build/
---
Might save you a few keystrokes !  


You can also declare custom aliases using `[alias]` in the `.cargo/config` file.  
List of default aliases is:  

    b = "build"
    c = "check"
    t = "test"
    r = "run"
## [5][Frustrated with the state of crates.io](https://www.reddit.com/r/rust/comments/hho5ec/frustrated_with_the_state_of_cratesio/)
- url: https://www.reddit.com/r/rust/comments/hho5ec/frustrated_with_the_state_of_cratesio/
---
I've been looking to create small but useful libraries, and looking for a name has been a very frustrating adventure between all the taken names that have no owners... And I don't mind if the library is well maintained or if it's used.

My current project is a light-weight type-safe Inversion-of-Control library. It could be used like this (actual tests):

```rust
    #[test]
    fn it_works_with_non_copy_types() {
        fn plus_one(a: &amp;Arc&lt;Mutex&lt;u32&gt;&gt;) -&gt; u32 {
            *a.lock().unwrap() + 1
        }
        fn inc(a: &amp;Arc&lt;Mutex&lt;u32&gt;&gt;) -&gt; () {
            *a.lock().unwrap() += 1;
        }

        let mut i = Injector::default();
        i.add_value(Arc::new(Mutex::new(20u32)));

        assert_eq!(i.call(plus_one), 21);
        i.call(inc);
        assert_eq!(i.call(plus_one), 22);
    }
```

Thought it could be useful, as I haven't seen any projects do that yet. I'm planning some stuff like value factories and singletons (right now it's all singletons), adding context and move semantics, etc.

So the first name that came to mind is [Injector](https://crates.io/crates/injector) which stated, two years ago;

&gt; Repository will be filled with existing source code soon. Thanks for your patience.

The repository pointed at doesn't exist, there's no docs, and no dependents.

So I started looking for another good name. [Tapioca](https://crates.io/crates/tapioca)? v0.0.1, 3 years ago, no repo again. [Brioche](https://crates.io/crates/brioche) seems like an actual project, and I'll give it the benefit of the doubt, but there isn't much. Just plain [Ioc](https://crates.io/crates/ioc)? 4 years old, but at least it's used; [di](https://crates.io/crates/di) hasn't been for 5 years.

Geez. This reminds me of the early days of NPM, where people were just reserving names without any projects, just in case...

I'm still looking, if you think of something. /rant
## [6][Even Better TOML - Proper TOML Support for Visual Studio Code](https://www.reddit.com/r/rust/comments/hhq5ar/even_better_toml_proper_toml_support_for_visual/)
- url: https://www.reddit.com/r/rust/comments/hhq5ar/even_better_toml_proper_toml_support_for_visual/
---
_No more "ugh what was that setting..."_

[tl;dr: a TOML extension for VSCode.](https://marketplace.visualstudio.com/items?itemName=tamasfe.even-better-toml)

With Rust Analyzer becoming more awesome than ever, I felt that `Cargo.toml` and TOML in general needed some attention as well.

There has also been a need for a format-preserving TOML parser in Rust, and I haven't really written any parsers in my life so it was time for one (and TOML is simple enough for a start, or so I thought).

I originally wanted to contribute to Rust Analyzer (and still do), so a month ago I started working on a TOML parser in my favourite language. However I am not entirely happy with it yet (it is [Taplo](https://github.com/tamasfe/taplo) by the way), so instead of focusing on making a generic-purpose parser library I decided to make a VSCode extension based on it to test it out and shape it as I go.

This is how [Even Better TOML](https://marketplace.visualstudio.com/items?itemName=tamasfe.even-better-toml) was born.

It is not yet complete nor bug-free, but I consider it somewhat usable in its current state.
I've just added experimental auto completion and such for `Cargo.toml` in it, which I personally was really missing.

I haven't abandoned my original goal of contributing to Rust Analyzer, but I'm a bit hesitant as the code quality might not be up to par with that project's. I will probably contact them in the near future anyway.

For now I invite you guys to check the extension out, and I hope you like it and/or find it useful!
## [7][Show /r/rust: 🔥 in WASM](https://www.reddit.com/r/rust/comments/hhpchk/show_rrust_in_wasm/)
- url: https://smmalis37.github.io/wasm-playground/
---

## [8][Sai - A minimal IoC/DI framework for Rust](https://www.reddit.com/r/rust/comments/hhyft8/sai_a_minimal_iocdi_framework_for_rust/)
- url: https://github.com/zhming0/sai
---

## [9][Manipulating ports, virtual ports and pseudo terminals [Rust Wrocław Webinar, 06/10/2020]](https://www.reddit.com/r/rust/comments/hhvdw2/manipulating_ports_virtual_ports_and_pseudo/)
- url: https://youtu.be/_cYz03jS7tk
---

## [10][rust-analyzer changelog #31](https://www.reddit.com/r/rust/comments/hhzlsq/rustanalyzer_changelog_31/)
- url: https://rust-analyzer.github.io/thisweek/2020/06/29/changelog-31.html
---

## [11][Heliocron 0.4.0 - combine with cron to execute tasks relative to sunset/sunrise. Now has support for civil, nautical and astronomical dusk/dawn](https://www.reddit.com/r/rust/comments/hhnat5/heliocron_040_combine_with_cron_to_execute_tasks/)
- url: https://www.reddit.com/r/rust/comments/hhnat5/heliocron_040_combine_with_cron_to_execute_tasks/
---
Long time Python programmer, relatively new Rust programmer. I've been steadily working on this command line application, Heliocron, which works hand in hand with cron to fire your crontab entries at times relative to various solar events.

Originally supporting just sunrise and sunset (with an optional additional offset), support has now been added for civil, nautical and astronomical dawn and dusk.

For reference

* Sunset/rise is when the top of the sun disappears below the horizon
* Civil dawn/dusk is when the centre of the sun is 6 degrees below the horizon
* Nautical dawn/dusk is when the centre of the sun is 12 degrees below the horizon
* Astronomical dawn/dusk is when the centre of the sun is 18 degrees below the horizon

It's been tricky but very enjoyable learning a much lower level language, and it's been difficult to shake the python idioms and replace them with rust idioms. 

[Here is the github repo to read further about the project.](https://github.com/mfreeborn/heliocron) 

Open to suggestions for improvements both in functionality and code style.
## [12][Rust structs - Wasm Loading dynamically](https://www.reddit.com/r/rust/comments/hhzmao/rust_structs_wasm_loading_dynamically/)
- url: https://www.reddit.com/r/rust/comments/hhzmao/rust_structs_wasm_loading_dynamically/
---
Hey People,

I got a small Problem with WebAssembly and its imports. I want to load the wasm file dynamically to my javascript file. It looks something like this:

    //javascriptfile.js
    
    const rust = import('./../eye-move-image/pkg/eye_move_image_bg.wasm')
    let code;
    
    function setup(m){
        code = m;
        console.log(code.greet());
        let instance = code.Foo.new();
        console.log(code.bar());
    }
    
    rust
        .then(m =&gt; setup(m))
        .catch(console.error)

And here the Rust file:

    //rustfile.rs
    use wasm_bindgen::prelude::*;
    
    #[wasm_bindgen]
    pub fn greet() -&gt; i32{
        42
    }
    
    
    #[wasm_bindgen]
    pub struct Foo{
        val : i32,
    }
    
    #[wasm_bindgen]
    impl  Foo{
        pub fn new() -&gt; Foo{
            Foo {
                val: 3,
            }
        }
    
        pub fn bar(&amp;self) -&gt; i32{
            self.val
        }
    }

I am compiling the rust code with the following Command without flags.

    wasm-pack build

By the wasm\_bindgen Guide( [https://rustwasm.github.io/docs/wasm-bindgen/examples/char.html](https://rustwasm.github.io/docs/wasm-bindgen/examples/char.html) )this should work, shouldn't it? I am kinda frustrated right now.

The console.log(greet()) call works. But The call of the new() methode doesnt, do you have to call the methods in a different way, or what is the problem?  


I am getting the error Message  


    TypeError: Cannot read property 'new' of undefined
        at setup (tets.js:8)
        at eval (tets.js:13)

And i am using the npm start command to start my local server.
