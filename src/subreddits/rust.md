# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (4/2020)!](https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/)
- url: https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 322](https://www.reddit.com/r/rust/comments/esbf8h/this_week_in_rust_322/)
- url: https://this-week-in-rust.org/blog/2020/01/21/this-week-in-rust-322/
---

## [3][ANN: Handlebars 3.0 with performance boost](https://www.reddit.com/r/rust/comments/et4qrz/ann_handlebars_30_with_performance_boost/)
- url: https://www.reddit.com/r/rust/comments/et4qrz/ann_handlebars_30_with_performance_boost/
---
Handlebars was a general purpose template engine written in Rust. It serves [rust-lang.org](https://rust-lang.org), [docs.rs](https://docs.rs) and the Rust book via mdbook. 

My latest 3.0.0 has been released to [crates.io](https://crates.io). I have been working on performance improvements in the last few weeks. The benchmark shows up to 4x boost on rendering large \`each\` block and nested ones. Some logic was moved to template parsing phase, so with 11% overhead on \`parse\_template\` we got big improvements on render phase, without losing features. 

[benchmark](https://preview.redd.it/1gd7i12ecnc41.png?width=737&amp;format=png&amp;auto=webp&amp;s=bc89ce4083ae6f37326d8d6eef9f4883d8277f64)

It is recommend to upgrade you application to use 3.0. If you are using handlebars directly you should not see API changes. Helper developers may need to upgrade your implementations if yours were using \`RenderContext\` deeply.

Feel free to report issue or ask for help on the github repo: [https://github.com/sunng87/handlebars-rust/](https://github.com/sunng87/handlebars-rust/)
## [4][cargo-audit v0.11: Introducing the `fix` feature, yanked crate detection, and more](https://www.reddit.com/r/rust/comments/et04on/cargoaudit_v011_introducing_the_fix_feature/)
- url: https://blog.rust-lang.org/inside-rust/2020/01/23/Introducing-cargo-audit-fix-and-more.html
---

## [5][TIL why the eh_personality language item is called that](https://www.reddit.com/r/rust/comments/estvau/til_why_the_eh_personality_language_item_is/)
- url: https://www.reddit.com/r/rust/comments/estvau/til_why_the_eh_personality_language_item_is/
---
The name of the [language item `eh_personality`](https://doc.rust-lang.org/unstable-book/language-features/lang-items.html#more-about-the-language-items) probably sounds completely mundane to a seasoned systems programmer's ears. But my ears were primed by a degree in English and American Studies / Linguistics rather than CS, so when I first heard it, it sounded like a whimsical pun on Herman Melville's *Bartleby the Scrivener*: ["Ah Bartleby! Ah humanity!"](http://www.gutenberg.org/cache/epub/11231/pg11231.txt) :)

Much to my dismay though, no such backstory for the name turned up when searching online. Actually, no backstory at all. This could only mean one thing -- the name must be perfectly ordinary and transparent in the eyes of the people who came up with it, so much so they didn't even realize it might require explanation for the uninitiated.¹

I was stumped for a while, but I think I've finally managed to elucidate it thanks to [this answer](https://stackoverflow.com/a/329092/1826241) to a SO question about `__gxx_personality_v0`. So here goes, for any fellow Rustaceans with non-CS backgrounds:

- `eh` is not an exclamation but probably an abbreviation of *exception handling*
- `personality` because the corresponding function determines what type of exception handling to perform (and of course, ultimately because tradition: the name comes from C++ and can be traced back to the [Itanium C++ ABI](https://itanium-cxx-abi.github.io/cxx-abi/abi-eh.html#base-personality), as mentioned in [another answer](https://stackoverflow.com/a/329195/1826241) to that SO question)

For anyone hearing about it for the first time: `eh_personality` comes up when implementing [freestanding Rust binaries](https://os.phil-opp.com/freestanding-rust-binary/#the-eh-personality-language-item).

Ah Bartleby! Eh personality!

¹ EDIT: Spoke too soon, thanks to /u/flying-sheep for pointing out [this *is* explained on the Rust side](https://doc.rust-lang.org/1.2.0/std/rt/unwind/index.html), my google-fu was apparently just too weak :)
## [6][Novice question: Will it be a good idea if I approached the embedded programming zone with Rust?](https://www.reddit.com/r/rust/comments/et9jqc/novice_question_will_it_be_a_good_idea_if_i/)
- url: https://www.reddit.com/r/rust/comments/et9jqc/novice_question_will_it_be_a_good_idea_if_i/
---
Heyo, an undergrad CS student here.


I'm also a complete novice who's basically stuck in a capstone project regarding embedded stuff (because democracy is a wonderful concept). Though the coding experience is in Python, C#, and js, was interested in checking out Rust for a while and was wondering if it would be a good idea to approach our project with it :]


If it is a good idea somehow, how would it be different if I initially approached everything with the usual toolchain regarding C? 
Would really appreciate it if you could mention your experiences with rust in the embedded zone.
Also, what kind of resources can an absolute rookie can use to dive into it? Is the Rust Embedded book enough? :[
## [7][An implementation of the game snake for the stm32f3 discovery board!](https://www.reddit.com/r/rust/comments/et6y2e/an_implementation_of_the_game_snake_for_the/)
- url: https://www.reddit.com/r/rust/comments/et6y2e/an_implementation_of_the_game_snake_for_the/
---
Using the stm32f3, an 8x8 LED display and an analog joystick, I implemented snake using Rust's real-time embedded framework for Cortex-M microcontrollers - Real Time For the Masses. For those of you interested, the [code is here.](https://github.com/arosspope/usnake)

This project was primarily a learning exercise in understanding how Rust can be used to solve some of the challenges inherent in embedded application development. In a field where C is still king, It will be interesting to see if Rust can disrupt traditional embedded development practices. I for one greatly enjoyed using it for this project and would do so again in the future.
## [8][Refacture my datastructure (no_std, no alloc)](https://www.reddit.com/r/rust/comments/et79yf/refacture_my_datastructure_no_std_no_alloc/)
- url: https://www.reddit.com/r/rust/comments/et79yf/refacture_my_datastructure_no_std_no_alloc/
---
Hi Rustaceans!  


Please enlighten me with the options of this awesome language, for a better data structure than i currently have (if any such thing).  


Other optimization, comments and low hanging fruits are also more than welcome!  


[https://github.com/BlackbirdHQ/at-rs/issues/9](https://github.com/BlackbirdHQ/at-rs/issues/9)
## [9][Show /r/rust: ytop, a TUI system monitor and rust port of gotop](https://www.reddit.com/r/rust/comments/esvygl/show_rrust_ytop_a_tui_system_monitor_and_rust/)
- url: https://www.reddit.com/r/rust/comments/esvygl/show_rrust_ytop_a_tui_system_monitor_and_rust/
---
Link: https://github.com/cjbassi/ytop

I'm pleased to finally announce a MVP release of ytop! I've been working on this for a few months, and I've been planning it for even longer. A lot of my time has been spent working on [rust-psutil](https://github.com/borntyping/rust-psutil) which I am now a maintainer of, and which just got some preliminary Mac support. Mac support is still a little lacking, but good enough for people to start using ytop. Other than that, ytop has almost feature parity with gotop, but missing a few features and some polish, which I hope to get to in the next few weeks.

I originally wrote gotop right after learning Go a couple of years ago as my first project in that language. It was a great learning experience, and I'm not sure I could have written it in Rust at the time, but after learning Rust, I much prefer writing in it and the library ecosystem is also amazing. Shout out to the structopt, tui, and crossterm crates in particular.
## [10][Secret Scanning Tool: Rusty-Hog - Scan your source and cloud for passwords and API keys. Thanks to the blazingly fast Rust Regex engine I saw 40%+ speed improvements over the Python equivalent of this tool.](https://www.reddit.com/r/rust/comments/eswfmf/secret_scanning_tool_rustyhog_scan_your_source/)
- url: https://crates.io/crates/rusty_hogs/1.0.1
---

## [11][If you’re in the East, please consider Rust!](https://www.reddit.com/r/rust/comments/et9wlh/if_youre_in_the_east_please_consider_rust/)
- url: https://medium.com/@ly.lee/if-youre-in-the-east-please-consider-rust-23fa05873397?source=friends_link&amp;sk=df52d4b5da80006ff42fc036ef01e12c
---

## [12][What's a good first program for writing Rust?](https://www.reddit.com/r/rust/comments/et0loj/whats_a_good_first_program_for_writing_rust/)
- url: https://www.reddit.com/r/rust/comments/et0loj/whats_a_good_first_program_for_writing_rust/
---
I have taken a Rust class, taken Rustlings and read parts of The Book. I am very sold on the language, but I want to showcase its benefits vs. other languages for my colleagues.

I was thinking about writing something in Rust that compares to something in C, but I am open to any idea.

What do you think you is a good first program for writing Rust (primarily a Go dev with a CS background) that also showcases some of its strengths? Thanks!
