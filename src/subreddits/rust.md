# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (7/2020)!](https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/)
- url: https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 325](https://www.reddit.com/r/rust/comments/f3a5hu/this_week_in_rust_325/)
- url: https://this-week-in-rust.org/blog/2020/02/11/this-week-in-rust-325/
---

## [3][All chapters from Rust in Action, book published by Manning, are now available](https://www.reddit.com/r/rust/comments/f3pq12/all_chapters_from_rust_in_action_book_published/)
- url: https://www.manning.com/books/rust-in-action?a_aid=rust&amp;a_bid=0367c58f&amp;chan=reddit
---

## [4][tk - a colorful version of hexdump](https://www.reddit.com/r/rust/comments/f3r9dc/tk_a_colorful_version_of_hexdump/)
- url: https://www.reddit.com/r/rust/comments/f3r9dc/tk_a_colorful_version_of_hexdump/
---
Hey all rustys! Just released my first "serious" rust tool, `tk`:

[https://github.com/MrYakobo/tk](https://github.com/MrYakobo/tk)

It was created out of the fact that `hexdump -C`  don't provide enough insight in the patterns of a given file. This new tool makes file "analysis" a breeze. The image below is generated from `tk &lt; /usr/bin/vim`:

&amp;#x200B;

https://preview.redd.it/oz6p65jyrvg41.png?width=784&amp;format=png&amp;auto=webp&amp;s=d1a837fa1c86cb2ecd8938cae9a226f2ffdec258

The header on elf-binaries are totally visible, which I think is cool. [Here](https://github.com/MrYakobo/tk#the-gallery) are some more file types that I tested the tool on.

I initially implemented this in Golang, but I felt lazy doing that and not really empowered to fix the bugs that arose. With rust, I felt sharp all the time, really. This is probably going to be my next go-to language in the future :)
## [5][Sealed Rust Update: The Plan for Safety Critical Rust](https://www.reddit.com/r/rust/comments/f3alfk/sealed_rust_update_the_plan_for_safety_critical/)
- url: https://ferrous-systems.com/blog/sealed-rust-the-plan/
---

## [6][Svgbob 0.5.0-alpha is 3rd architecture rewrite of svgbob to support styling of shapes in the diagrams](https://www.reddit.com/r/rust/comments/f3qwup/svgbob_050alpha_is_3rd_architecture_rewrite_of/)
- url: https://ivanceras.github.io/svgbob-editor/
---

## [7][My first rust crate: Embedded-hal SPI driver for the ADXL355 accelerometer](https://www.reddit.com/r/rust/comments/f3runn/my_first_rust_crate_embeddedhal_spi_driver_for/)
- url: https://jitter.company/blog/2020/02/14/adxl355-embedded-hal-driver-crate/
---

## [8][Rust lang in a nutshell - Part 2 Mini-series - Enums, pattern matching, Options](https://www.reddit.com/r/rust/comments/f3r27g/rust_lang_in_a_nutshell_part_2_miniseries_enums/)
- url: https://www.softax.pl/blog/rust-lang-in-a-nutshell-2-enums-pattern-matching-options/
---

## [9][Odd Missed Optimization](https://www.reddit.com/r/rust/comments/f3fuwv/odd_missed_optimization/)
- url: https://www.reddit.com/r/rust/comments/f3fuwv/odd_missed_optimization/
---
I was optimizing one of my projects (an emulator) and while looking at rustc's assembly output I noticed some very suspicious bounds checks. The original offending code is gone at this point but I recreated it in this example:

    pub fn llvm_test(mut region: u32) -&gt; u32 {
        region &gt;&gt;= 1;
        if region &gt;= 0x08 &amp;&amp; region &lt;= 0x0D {
            if region - 0x08 &gt; 1000 {
                return 0xFFFFFFFF;
            }
            return 66;
        } else {
            return 45;
        }
    }

godbolt: https://rust.godbolt.org/z/t7JTAR

Even at `opt-level=3` the output will always do the `region - 0x08 &gt; 1000` check despite region being in the range [0x08, 0x0D]. I think this might just be an issue with LLVM because clang seems to do the same thing with the equivalent C code (https://rust.godbolt.org/z/6bNPwo). GCC doesn't have this issue and eliminates the check. Is this something I should just ignore? I was concerned because code like this was causing the compiler spit out asm for panics in many places where I knew panics couldn't occur.

EDIT:

If I change the return into a panic instead, the compiler still generates all of the extra panic code but only if I do a right shift. It eliminates the check if I do anything else (https://rust.godbolt.org/z/67YgdC).
## [10][ustr 0.4 - Fast, FFI-Friendly String Interning](https://www.reddit.com/r/rust/comments/f3h2vw/ustr_04_fast_ffifriendly_string_interning/)
- url: https://www.reddit.com/r/rust/comments/f3h2vw/ustr_04_fast_ffifriendly_string_interning/
---
Repo with lots more info in the readme: [https://github.com/anderslanglands/ustr](https://github.com/anderslanglands/ustr)

I've published a new version of my ustr crate to [crates.io](https://crates.io).

Since the first version I posted here I've tried to document the unsafe invariants better, and use checked\_add() to guard against potential overflow. I've also added Miri to the CI as an extra guard.

I've tweaked the hasher and binning to get it even faster, as well as adding serde support. 

I've been using this heavily in another project of mine for a while now and it's working great. I'll probably aim for a 1.0 version soon.
## [11][Some Nuances of Undefined Behavior in Rust](https://www.reddit.com/r/rust/comments/f3ekb8/some_nuances_of_undefined_behavior_in_rust/)
- url: https://typr124.github.io/UB1
---

## [12][How nnethercote optimized the heck out of LEB128 code](https://www.reddit.com/r/rust/comments/f36j05/how_nnethercote_optimized_the_heck_out_of_leb128/)
- url: https://github.com/rust-lang/rust/pull/69050#issuecomment-585508353
---

