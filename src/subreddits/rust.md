# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (44/2020)!](https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (44/2020)?](https://www.reddit.com/r/rust/comments/ji8vgk/whats_everyone_working_on_this_week_442020/)
- url: https://www.reddit.com/r/rust/comments/ji8vgk/whats_everyone_working_on_this_week_442020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-44-2020/50643?u=llogiq)!
## [3][Here's a fun little crate I made that generates an interactive ASCII globe in the terminal](https://www.reddit.com/r/rust/comments/jj0pde/heres_a_fun_little_crate_i_made_that_generates_an/)
- url: https://i.redd.it/0qub1980wmv51.gif
---

## [4][Pants 2.0.0 released - Generic build system in Rust](https://www.reddit.com/r/rust/comments/jjcbka/pants_200_released_generic_build_system_in_rust/)
- url: https://blog.pantsbuild.org/introducing-pants-v2/
---

## [5][Any Minecraft fans? Rust and WASM to map minecraft worlds locally](https://www.reddit.com/r/rust/comments/jj6k9b/any_minecraft_fans_rust_and_wasm_to_map_minecraft/)
- url: https://www.reddit.com/r/rust/comments/jj6k9b/any_minecraft_fans_rust_and_wasm_to_map_minecraft/
---
[owengage.com/anvil](https://owengage.com/anvil/)

Hi there, I've been working on a crate called [fastnbt](https://github.com/owengage/fastnbt) for several months now in my spare time. It started as just trying to learn Rust by writing a parser for NBT (a data format specific to Minecraft), and ended up with me doing far more with it.

`fastnbt` now lets you parse 'unstructured' NBT with the original parser, but also supports a serde deserializer, specifically allowing you to avoid doing memory allocations in some places (like for strings and arrays of data).

I eventually wanted to show it off and ended up compiling to to WASM and creating what I linked above.

It renders a world entirely in the browser using the File APIs. It never uploads your world to anywhere to do any processing, it's all in the browser.

I was curious if people would enjoy it/find it useful. But also happy to take feature request/code criticism/suggestions etc.
## [6][LudumDare 47 - The Island](https://www.reddit.com/r/rust/comments/jjet2n/ludumdare_47_the_island/)
- url: https://blog.kuviman.com/2020/10/18/ludumdare47.html
---

## [7][I've needed a crate like this at work lately, so I thought I'd publish it.](https://www.reddit.com/r/rust/comments/jjby3r/ive_needed_a_crate_like_this_at_work_lately_so_i/)
- url: https://www.reddit.com/r/rust/comments/jjby3r/ive_needed_a_crate_like_this_at_work_lately_so_i/
---
[https://crates.io/crates/is\_not](https://crates.io/crates/is_not)

is\_not is a simple little crate that allows crates to pick up feature activations of libraries they depend on.
## [8][Introducing rusteval - add a REPL to your application with a few macros](https://www.reddit.com/r/rust/comments/jj09i8/introducing_rusteval_add_a_repl_to_your/)
- url: https://docs.rs/rusteval/0.1.0/rusteval/
---

## [9][Tokio 0.3 to 0.2 compatibility layer.](https://www.reddit.com/r/rust/comments/jj80co/tokio_03_to_02_compatibility_layer/)
- url: https://github.com/LucioFranco/tokio-compat-02/
---

## [10][fixed-slice-vec: an Embedded Rust no-std Vector](https://www.reddit.com/r/rust/comments/jj537w/fixedslicevec_an_embedded_rust_nostd_vector/)
- url: https://blog.auxon.io/2020/10/26/fixed-slice-vec/
---

## [11][Question about this hypothetical [T; N]::into_iter](https://www.reddit.com/r/rust/comments/jjcf20/question_about_this_hypothetical_t_ninto_iter/)
- url: https://www.reddit.com/r/rust/comments/jjcf20/question_about_this_hypothetical_t_ninto_iter/
---
For a while I've been hearing about 3 major issues preventing iterating over native arrays, once constant generics become widely available:

- Weight: an array iterator would entail copying the entire array every time the iterator is moved.
- Compatibility: array.into_iter() currently falls back to slice iteration; introducing a new meaning could break backwards compatibility.
- Soundness: an array iterator could be unsound WRT accidental double-frees, especially in the context of potential panics.

My question is about the last one. Where does the potential unsoundness come from? I created a hypothetical simple implementation [here](https://gist.github.com/Lucretiel/477bf0b5b222eca8f3b77e412518a59c). Is there any unsoundness in this implementation? It seems like in the worst case you might end up leaking some data, but I don't think there's a way to actually cause memory unsafety with that design.
## [12][Github repo of tutorials how to build something in rust lang](https://www.reddit.com/r/rust/comments/jj29ir/github_repo_of_tutorials_how_to_build_something/)
- url: https://www.reddit.com/r/rust/comments/jj29ir/github_repo_of_tutorials_how_to_build_something/
---
Not so long ago I stumbled across this [**build-your-own-x**](https://github.com/danistefanovic/build-your-own-x) repo. I thought that it might be cool to have that one but only specifically for rust. So I took all the rust based links from there and placed them in my [**repo**](https://github.com/strawberrycheeze/build-your-own-x-in-rust). Also I added a lot of projects that I found on my own.   
I think it might be interesting for anyone who wants to take up rust and try something serious.   
Anyway, any feedback or advice will be appreciated.  
Peace  

repo: [https://github.com/strawberrycheeze/build-your-own-x-in-rus](https://github.com/strawberrycheeze/build-your-own-x-in-rust)
