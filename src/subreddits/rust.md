# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (46/2020)!](https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 364](https://www.reddit.com/r/rust/comments/jslo80/this_week_in_rust_364/)
- url: https://this-week-in-rust.org/blog/2020/11/11/this-week-in-rust-364/
---

## [3][Crust of Rust: Sorting Algorithms [video]](https://www.reddit.com/r/rust/comments/jua24q/crust_of_rust_sorting_algorithms_video/)
- url: https://youtu.be/h4RkCyJyXmM
---

## [4][ECS scheduler thoughts, part 1](https://www.reddit.com/r/rust/comments/jujamr/ecs_scheduler_thoughts_part_1/)
- url: https://ratysz.github.io/article/scheduling-1/
---

## [5][Render a terminal as HTML with `to-html`](https://www.reddit.com/r/rust/comments/juhbo3/render_a_terminal_as_html_with_tohtml/)
- url: https://www.reddit.com/r/rust/comments/juhbo3/render_a_terminal_as_html_with_tohtml/
---
I made a [program](https://github.com/Aloso/to-html) that can render a terminal on a webpage, including ANSI colors. Demonstration:

https://preview.redd.it/js3aibzsybz51.png?width=1280&amp;format=png&amp;auto=webp&amp;s=e73f3ece6114aca1bd69633cc977cd584e3191eb

As you can see, running `to-html 'your bash command'` will execute the command(s) and print a terminal with the output as HTML. All you need to do is copy+paste it on your website :)

You can see it in action here: [https://aloso.github.io/colo/](https://aloso.github.io/colo/)

**How it works**

to-html consumes 1 or more commands and executes them with [script](https://man7.org/linux/man-pages/man1/script.1.html). This program runs the commands with bash in a pseudoterminal (pty), so the commands emit colors normally. This is necessary because most programs disable colors when they detect that they're not running in a tty.

to-html then parses the ANSI escape codes in the output and converts them to HTML. It also parses the bash commands and applies some syntax highlighting (using [logos](https://github.com/maciejhirsz/logos)). Finally, the bash commands and the output are printed.

I hope some of you will find this useful!
## [6][pacaptr: Pacman-like Syntax Wrapper for Package Managers, Now Reaches v0.8.0](https://www.reddit.com/r/rust/comments/juk5xp/pacaptr_pacmanlike_syntax_wrapper_for_package/)
- url: https://github.com/rami3l/pacaptr
---

## [7][Error handling Isn't All About Errors by Jane Lusby](https://www.reddit.com/r/rust/comments/juc8l0/error_handling_isnt_all_about_errors_by_jane_lusby/)
- url: https://www.youtube.com/watch?v=rAF8mLI0naQ
---

## [8][Benchmarks of Ryzen 5 series CPUs](https://www.reddit.com/r/rust/comments/juetus/benchmarks_of_ryzen_5_series_cpus/)
- url: https://www.reddit.com/r/rust/comments/juetus/benchmarks_of_ryzen_5_series_cpus/
---
Has anyone run the [rustc-perf](https://github.com/rust-lang/rustc-perf/tree/master/collector) benchmark suite on a new (or last gen) Ryzen CPU?

It would be quite useful if there was a comparison of how rustc performed across the 5600x/5800x/5900x/5950x.
## [9][How to work with binary protocols?](https://www.reddit.com/r/rust/comments/jukzbp/how_to_work_with_binary_protocols/)
- url: https://www.reddit.com/r/rust/comments/jukzbp/how_to_work_with_binary_protocols/
---
I am learning network programming (self-taught). I was looking at how binary protocols work and wanted to implement one. I looked at Go's [encoding/binary](https://golang.org/pkg/encoding/binary/) and Rust's [byteorder](https://docs.rs/byteorder/1.3.4/byteorder/) crate. They both provide methods to encode only  numbers. I can't understand how would one encode other data types like strings etc. Am I missing something? Also any resources to learn about binary protocols would be very helpful.

Thanks a lot!
## [10][Make a Language in Rust, Part Ten: Starting Again](https://www.reddit.com/r/rust/comments/ju1vhl/make_a_language_in_rust_part_ten_starting_again/)
- url: https://arzg.github.io/lang/10/
---

## [11][Design Pattern: Scoped Trait Implementations](https://www.reddit.com/r/rust/comments/ju6x5i/design_pattern_scoped_trait_implementations/)
- url: https://jack.wrenn.fyi/blog/private-trait-impls/
---

## [12][Sapling Livestream #3: Building a better code editor in Rust](https://www.reddit.com/r/rust/comments/ju71ty/sapling_livestream_3_building_a_better_code/)
- url: https://www.reddit.com/r/rust/comments/ju71ty/sapling_livestream_3_building_a_better_code/
---
Hi Rustaceans,

I'm streaming some Rust development [here](https://youtu.be/o8IiQPlY-GQ) on my YouTube channel.  I'm building an experimental text editor called [Sapling](https://github.com/kneasle/sapling), which aims to help programmers edit *code* faster rather than just *text*.  I'm still pretty new to streaming, but if you're welcome to come along and chill out whilst we build some Rust code.  Stream starts at 2pm EST every Saturday!

Cheers, and happy coding!
