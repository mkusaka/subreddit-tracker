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
## [3][Bitfields Forever: Why we need a C-compatible Rust Crate](https://www.reddit.com/r/rust/comments/euxy49/bitfields_forever_why_we_need_a_ccompatible_rust/)
- url: https://immunant.com/blog/2020/01/bitfields/
---

## [4][rust-analyzer Changelog #9](https://www.reddit.com/r/rust/comments/eus0cj/rustanalyzer_changelog_9/)
- url: https://rust-analyzer.github.io/thisweek/2020/01/27/changelog-9.html
---

## [5][Security by design: an introduction to the Rust programming language](https://www.reddit.com/r/rust/comments/ev1zqf/security_by_design_an_introduction_to_the_rust/)
- url: https://medium.com/tadaweb/security-by-design-a-brief-introduction-to-rust-378060e45038
---

## [6][Can you send me snippets of code where you *needed* higher kinded lifetimes? I wish to identify as many cases as possible to categorize them and make rustc understand them and guide you in their direction.](https://www.reddit.com/r/rust/comments/euv4mq/can_you_send_me_snippets_of_code_where_you_needed/)
- url: https://www.reddit.com/r/rust/comments/euv4mq/can_you_send_me_snippets_of_code_where_you_needed/
---
[I asked Twitter this question already](https://twitter.com/ekuber/status/1221515971610832897), but I wish to have as big a corpus as possible. Can you send me snippets of code where you **needed** higher kinded lifetimes? I wish to identify as many cases as possible to categorize them and make `rustc` understand them and guide you in their direction.

To give you an idea of what I'm looking for, the following snippet

    trait Trait&lt;T&gt; {
        fn do_something(&amp;self, value: T);
    }
    
    fn foo(b: Box&lt;Trait&lt;&amp;usize&gt;&gt;) {
        let x: usize = 10;
        b.do_something(&amp;x);
    }

produces the following error

    error[E0597]: `x` does not live long enough
     --&gt; file.rs:7:20
      |
    5 | fn foo(b: Box&lt;Trait&lt;&amp;usize&gt;&gt;) {
      |        - has type `std::boxed::Box&lt;dyn Trait&lt;&amp;'1 usize&gt;&gt;`
    6 |     let x: usize = 10;
    7 |     b.do_something(&amp;x);
      |     ---------------^^-
      |     |              |
      |     |              borrowed value does not live long enough
      |     argument requires that `x` is borrowed for `'1`
    8 | }
      | - `x` dropped here while still borrowed

but if you change it to use Higher-Ranked lifetimes

    trait Trait&lt;T&gt; {
        fn do_something(&amp;self, value: T);
    }
    
    fn foo(b: Box&lt;for&lt;'a&gt; Trait&lt;&amp;'a usize&gt;&gt;) {
        let x: usize = 10;
        b.do_something(&amp;x);
    }

The code actually compiles. Ideally, `rustc` should detect this case and help you get to working code faster.
## [7][Tokio v0.2.11 released, includes new `select!` and `join!` macros.](https://www.reddit.com/r/rust/comments/eushwv/tokio_v0211_released_includes_new_select_and_join/)
- url: https://github.com/tokio-rs/tokio/releases/tag/tokio-0.2.11
---

## [8][Units of Measure in Rust with Refinement Types](https://www.reddit.com/r/rust/comments/eun51n/units_of_measure_in_rust_with_refinement_types/)
- url: https://yoric.github.io/post/uom.rs/
---

## [9][Drue = Drums + Hue Lights](https://www.reddit.com/r/rust/comments/euxkzr/drue_drums_hue_lights/)
- url: https://www.reddit.com/r/rust/comments/euxkzr/drue_drums_hue_lights/
---
# Drue = Hue + Drums
[https://github.com/finnkauski/drue](https://github.com/finnkauski/drue)

Had posted about this before. Here's something worth showing off as far as a workable codebase goes.

First proper hobby project in Rust. To be honest it went from procedural to more struct based to functional and back again once I realised I know nothing about lifetimes. I rewrote it tons of times, learned how to do `macros` in order to avoid the issues I've had. I basically had the best and the *worst* times with the compiler and I'm starting to get it now.
## [10][actix-web or warp?](https://www.reddit.com/r/rust/comments/euyfr2/actixweb_or_warp/)
- url: https://www.reddit.com/r/rust/comments/euyfr2/actixweb_or_warp/
---
I've been wanting to make a web application in Rust. Well, an API/panel that controls something on the server. So I looked into frameworks with built-in websocket support, and came across [actix-web](https://actix.rs/) and [warp](https://github.com/seanmonstar/warp). They both appear to have their ups and downs (warp has much less `unsafe` and is simpler, while actix seems much more mature and fleshed out).

I want to know other developers' experiences with these frameworks, and their opinions too, so I can make an informed decision.
## [11][nd-triangulation brings you arbitrary dimensional Triangulations from cgal (c++) to rust](https://www.reddit.com/r/rust/comments/euspgo/ndtriangulation_brings_you_arbitrary_dimensional/)
- url: https://www.reddit.com/r/rust/comments/euspgo/ndtriangulation_brings_you_arbitrary_dimensional/
---
Hello Everyone,

I have build a very small [crate](https://github.com/Lesstat/nd-triangulation) that brings triangulations in arbitrary dimensions to rust. It uses the fantastic and exceptional [cpp](https://github.com/mystor/rust-cpp) crate to build an API to a few [cgal triangulation classes](https://doc.cgal.org/latest/Triangulation/index.html) and functionalities. 

Right now it is rather limited in various ways :). That is because it's tailored to my specific research, but if any of you guys has ideas/use cases I would be happy to try to implement them.

Also, I would very much appreciate feedback and ideas, as this is my first time publishing a crate. I guess there are some very general things I missed. But to have a very concrete question: I think the crate right now builds only on some linux distros as I do not know how to make a build script that more cleverly detects where the dependencies are located. Any input on that would help me a lot.

I want to give a huge thank you to the rust community. You all have built a community where it is easy to contribute and made it feel worthwhile enough for me to put the extra effort in to publish this :)!
## [12][Kondo 🧹 my first Rust crate.](https://www.reddit.com/r/rust/comments/euinys/kondo_my_first_rust_crate/)
- url: https://www.reddit.com/r/rust/comments/euinys/kondo_my_first_rust_crate/
---
[**Kondo**](https://github.com/tbillington/kondo) is a command line tool for removing things from your drive that don't bring you joy.

It identifies projects on your drive that are using a lot of space that could be saved! Examples include *target* in Cargo projects and *node_modules* in Node projects.

This is especially useful for:

* if you're about to zip up a bunch of projects to move them/back them up
* if you're running low on disk space and just want to free some up
* if you're a serial hoarder and want to keep all those things you cloned from github 1 time without the disk cost...

I've been using Rust on and off for a few years now. Inspired by tools like [ripgrep](https://github.com/BurntSushi/ripgrep) I made a [basic command line tool for helping me at work](https://github.com/tbillington/pj) a while back and was quite happy with how it turned out and was motivated me to make more!

I wanted to make this tool because I've done this manually way too many times to count. It doesn't currently delete anything because I'm paranoid that I'll `rm -rf /` somehow but I do plan to add the feature at some point, perhaps behind a prompt/flag.

You should just be able to run `cargo install kondo` and be on your way!

Sample output:

    $ kondo ~
    Scanning "C:/Users/Trent"
    1000 projects found
    Calculating savings per project
      (redacted 1000~ lines)
      385.6MB UnityTestApp (Unity) C:\Users\Trent\code\UnityTestApp
      458.7MB tokio (Cargo) C:\Users\Trent\code\tokio
        1.5GB ui-testing (Node) C:\Users\Trent\code\ui-testing
        4.0GB rust-analyzer (Cargo) C:\Users\Trent\code\rust-analyzer
    9.5GB possible savings

Any code feedback is greatly appreciated :)
