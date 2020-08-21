# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 352](https://www.reddit.com/r/rust/comments/icwq65/this_week_in_rust_352/)
- url: https://this-week-in-rust.org/blog/2020/08/18/this-week-in-rust-352/
---

## [3][PinePhone + Rust + Low Level Graphics/Input (no X11 required)](https://www.reddit.com/r/rust/comments/idg7ko/pinephone_rust_low_level_graphicsinput_no_x11/)
- url: https://v.redd.it/rlipzvyid7i51
---

## [4][Linux Terminal Tetris with Rust](https://www.reddit.com/r/rust/comments/idu3pb/linux_terminal_tetris_with_rust/)
- url: https://www.reddit.com/r/rust/comments/idu3pb/linux_terminal_tetris_with_rust/
---
https://preview.redd.it/9h1rssf2zbi51.png?width=973&amp;format=png&amp;auto=webp&amp;s=55a3711b36aa6c69236cd9ceaf701cb8e1aef43b

Actually I'm about to learn some Rust and i wanted to discover how to build a terminal based GUI. First I just used plain Rust without any special crates. It worked good enough to code some list displays, a kind of directory view. Then i tried out Termion which is really easy to use and I ended up making a Tetris game. It took me two programming sessions to complete this actual state, like 4 hours or a bit more in addition. This Tetris is playable with all basic features but that's it. I don't want to add more features like a Score or "next Tetromino Display" or something like that, because it was just a little side project to discover some  Terminal GUI programming with Rust. But feel free to check out my Code on [Github](https://github.com/DrTosh/tetris) and add your desired feature if you want. And don't forget to leave me your thoughts I would appreciate that ;D
## [5][Tge - A lightweight 2D game framework - 0.0.3 released](https://www.reddit.com/r/rust/comments/idsjo4/tge_a_lightweight_2d_game_framework_003_released/)
- url: /r/rust_gamedev/comments/idsh6u/tge_a_lightweight_2d_game_framework_003_released/
---

## [6][Pixels 0.2.0 released](https://www.reddit.com/r/rust/comments/idnq6r/pixels_020_released/)
- url: https://www.reddit.com/r/rust/comments/idnq6r/pixels_020_released/
---
Pixels 0.2.0 just landed. 🎉 This release focuses on further improvements to the custom shader API, and upgrades `wgpu` to 0.6.0.

[https://github.com/parasyte/pixels](https://github.com/parasyte/pixels)

This is the crate to use for emulators and old-school pixel-pushing. Especially if you want the full power of your GPU hardware for scaling and adding special effects to your raw pixel buffer.

## Featured game

[**Sandbox**](https://github.com/JMS55/sandbox) is a sand simulation game using `pixels`. It has a custom shader to add a glow effect. Check it out!

[Sandbox screen shot](https://preview.redd.it/gu2zjkleh9i51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=fa5ba4239fe213fbe2968602b4dc8bd4b5627bde)
## [7][RustConf is online! Official website: https://rustconf.com/](https://www.reddit.com/r/rust/comments/idecaz/rustconf_is_online_official_website/)
- url: https://www.youtube.com/watch?v=ESXMg9OzWrQ&amp;feature=emb_title
---

## [8][[knurling] defmt, a highly efficient Rust logging framework for embedded devices](https://www.reddit.com/r/rust/comments/idwb9c/knurling_defmt_a_highly_efficient_rust_logging/)
- url: https://ferrous-systems.com/blog/defmt/
---

## [9][I wrote about running Rust on the Arduino Uno](https://www.reddit.com/r/rust/comments/id83as/i_wrote_about_running_rust_on_the_arduino_uno/)
- url: https://creativcoder.dev/rust-on-arduino-uno
---

## [10][Problems with compilation](https://www.reddit.com/r/rust/comments/idvpb1/problems_with_compilation/)
- url: https://www.reddit.com/r/rust/comments/idvpb1/problems_with_compilation/
---
SOLVED!

A crate called \`passwors\` broke the compilation process.

&amp;#x200B;

I'm getting this weird error message

    error[E0463]: can't find crate for `syntax`
      --&gt; /home/lifeordeath/.cargo/registry/src/github.com-1ecc6299db9ec823/aster-0.27.0/src/lib.rs:10:1
       |
    10 | extern crate syntax;
       | ^^^^^^^^^^^^^^^^^^^^ can't find crate
    
    error: aborting due to previous error

when trying to compile a program that uses [rocket.rs](https://rocket.rs) and [diesel.rs](https://diesel.rs), but I can't for the life of me figure out what the syntax crate is.
## [11][Discount on the Rust book](https://www.reddit.com/r/rust/comments/idvoix/discount_on_the_rust_book/)
- url: https://twitter.com/nostarch/status/1296575248486273032?s=09
---

## [12][First thoughts on Rust vs OCaml](https://www.reddit.com/r/rust/comments/idh5be/first_thoughts_on_rust_vs_ocaml/)
- url: https://blog.darklang.com/first-thoughts-on-rust-vs-ocaml/
---

