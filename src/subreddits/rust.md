# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (16/2020)!](https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 334](https://www.reddit.com/r/rust/comments/g1fj7p/this_week_in_rust_334/)
- url: https://this-week-in-rust.org/blog/2020/04/14/this-week-in-rust-334/
---

## [3][A thank-you to the people who make all of this possible.](https://www.reddit.com/r/rust/comments/g2r23k/a_thankyou_to_the_people_who_make_all_of_this/)
- url: https://www.reddit.com/r/rust/comments/g2r23k/a_thankyou_to_the_people_who_make_all_of_this/
---
I'm in the middle of writing a chess engine in Rust as an exercise in a class I'm taking. As I build and refactor it, I've also been inspecting the generated assembly with \`cargo asm\` and comparing it to the examples given on the [chess programming wiki](https://www.chessprogramming.org/General_Setwise_Operations#Shifting_Bitboards), and I'm constantly surprised at how much I don't have to worry about manually optimizing it at the instruction level. I can write general-purpose code and the compiler can tell when I'm using a special case and automatically apply optimizations to it.

As an example, here's a general-purpose, direction-independent \`shift\` that moves the bits in a bitboard one square in the given direction:

    const NOT_8: u64 = 0x00ffffffffffffff;
    const NOT_1: u64 = 0xffffffffffffff00;
    const NOT_A: u64 = 0xfefefefefefefefe;
    const NOT_H: u64 = 0x7f7f7f7f7f7f7f7f;
    
    pub fn shift(x: u64, direction: Direction) -&gt; u64 {
        // look up rotate offset, and a mask 
        // to prevent wrap-around on ranks and files.
        let (offset, mask) = match direction {
            Direction::North =&gt; (8, NOT_8),
            Direction::South =&gt; (64 - 8, NOT_1),
            Direction::East =&gt; (1, NOT_H),
            Direction::West =&gt; (64 - 1, NOT_A),
            Direction::NorthEast =&gt; (9, NOT_8 &amp; NOT_H),
            Direction::NorthWest =&gt; (7, NOT_8 &amp; NOT_A),
            Direction::SouthEast =&gt; (64 - 7, NOT_1 &amp; NOT_H),
            Direction::SouthWest =&gt; (64 - 9, NOT_1 &amp; NOT_A),
        }
        (x &amp; mask).rotate_left(offset)
    }
    
    pub fn shift_north(x: u64) -&gt; u64
        shift(x, Direction::North)
    }

Now, if I generate code for \`shift\_north\`, the applied optimizations are incredible:

    shift_north:
        mov rax, rdi
        shl rax, 8
        ret

I expected to at least get constant propagation with the \`Direction\` lookup, but what's this?!  The compiler was able to statically infer that rotating 8 bits up with a pre-mask of 0x00ffff... is equivalent to a left shift by 8 bits, and completely dropped the pre-mask calculation.

The level of detail in both high-level and low-level optimizers is astounding. I've never realized just how much they're able to do, and all of it is the result of countless hours of work from others. I'm impressed, and truly grateful for how much they improve the performance and readability/maintainability of code. Imagine if I needed to manually write and optimize eight functions for every direction-dependent operation, instead of the one I have now.

So, to those of you who work endlessly at making compilers better: Thank you very much. Even though I love taking on big challenges, I can't imagine how much effort it has taken for us to come this far.
## [4][Upgrade Rust's Android SDK to API level 16 (from 14)](https://www.reddit.com/r/rust/comments/g30bpg/upgrade_rusts_android_sdk_to_api_level_16_from_14/)
- url: https://github.com/rust-lang/rust/pull/71123
---

## [5][Announcing Spair, a framework for Single Page Application in Rust](https://www.reddit.com/r/rust/comments/g2xzip/announcing_spair_a_framework_for_single_page/)
- url: https://www.reddit.com/r/rust/comments/g2xzip/announcing_spair_a_framework_for_single_page/
---
Hi crustaceans! I am very happy to share with you my library: A framework for Single Page Application in Rust.

https://github.com/aclueless/spair

Spair directly modifies its current vDOM (starts as an empty vDOM - instead of re-rendering a new vDOM + diffing vs old vDOM + patching). You can build an SPA site in pure Rust with Spair (without using macros. That said, macros can be added in the future for convenience).

I am sorry that there is no document yet.
## [6][Convert strings between kebab-case, snake_case, camelCase, and much more: convert_case released](https://www.reddit.com/r/rust/comments/g30frh/convert_strings_between_kebabcase_snake_case/)
- url: https://www.reddit.com/r/rust/comments/g30frh/convert_strings_between_kebabcase_snake_case/
---
I'd like to share `convert_case`, [a library](https://docs.rs/convert_case/0.1.0/convert_case/index.html) and binary that allows you to convert the casing of strings.  Installing the `convert_case` crate will give you an executable for changing the case of strings called `ccase`.  For example
```
$ ccase -t title super_mario_64
Super Mario 64

$ ccase -f snake -t title 2020-04-15_my_cat_cali
2020-04-16 My Cat Cali

$ ccase -t camel "convert to case"
convertToCase
```
`ccase` converts the input into a new case from the `--to -t` option.  It will determine word boundaries automatically, but you can supply a `--from -f` option to convert from a specific case.  Likewise you can use the library in your rust code.
```
use convert_case::{Case, Casing};

assert_eq!("Super Mario 64", "super_mario_64".to_case(Case::Title));
assert_eq!(
    "2020-04-15 My Cat Cali",
    "2020-04-15_my_cat_cali".from_case(Case::Snake).to_case(Case::Title)
);
assert_eq!("convertToCase", "convert to case".to_case(Case::Camel));
```
Works for a [great variety of cases](https://docs.rs/convert_case/0.1.0/convert_case/enum.Case.html), including kebab and train case!

This is my first crate.  I'd love some feedback!  I hope this serves useful.  I couldn't find a good CLI for converting cases and knew Rust was the right tool for the job.
## [7][5800 line program takes 3 minutes to compile](https://www.reddit.com/r/rust/comments/g2xxhj/5800_line_program_takes_3_minutes_to_compile/)
- url: https://www.reddit.com/r/rust/comments/g2xxhj/5800_line_program_takes_3_minutes_to_compile/
---
I have a fairly small program (5762 lines of Rust code according to Tokei), that takes 3 minutes to compile. This is an incremental release build, so it isn't recompiling any dependencies.

I haven't written any macros but use Serde derive quite a lot. Even so, this seems *very* slow! It's even worse in my case because I have to cross-compile too, so every change takes 6 minutes to build. Is there any way to profile which parts of my code are causing the slow-down? How do people with 100k-line programs deal with hour-long compile times?
## [8][Tutorial, Hash lookups without allocating](https://www.reddit.com/r/rust/comments/g2x6td/tutorial_hash_lookups_without_allocating/)
- url: https://github.com/sunshowers/borrow-complex-key-example/blob/master/README.md
---

## [9][What would you use to build 3D CAD software?](https://www.reddit.com/r/rust/comments/g30nel/what_would_you_use_to_build_3d_cad_software/)
- url: https://www.reddit.com/r/rust/comments/g30nel/what_would_you_use_to_build_3d_cad_software/
---
Hi everyone,

Not sure to ever start such a project, as I don't have a lot of time, know nothing about 3D, not much about GUI, but I was looking for what would be the best 3D rendering library to build a CAD software.

Lots of 3d libraries seem to be used for games, but I think they can do the job for other 3D needs too. 

What might differ is the ergonomics, yet at a lower level it must be the same, but as I said, I know nothing about 3D.

Another type of underrepresented apps in Rust right now are GUI  (I might be wrong, feel free to correct me), but I know bindings exist to GUI libraries and work well. But I don't know what would be a wise choice for a cross platform app.

So what would you use to start such a project ? gfx-rs + gtk-rs? something higher level than gfx-rs? Use electron for the GUI part (if possible)?

Thanks in advance for the time you'll take to answer this post
## [10][How do I get an image from a url?](https://www.reddit.com/r/rust/comments/g2zeps/how_do_i_get_an_image_from_a_url/)
- url: https://www.reddit.com/r/rust/comments/g2zeps/how_do_i_get_an_image_from_a_url/
---
Hey I am new to rust. I want to download am image from a url. If anyone could point me in the right direction it would be appreciated.

Here is the python equivalent just so that it's clear:

    def download_image(url):
        response = requests.get(url)
        return Image.open(BytesIO(response.content))
## [11][regex2fat: convert regexes to FAT32, with regex-automata and a few lines of safe Rust](https://www.reddit.com/r/rust/comments/g2dxpw/regex2fat_convert_regexes_to_fat32_with/)
- url: https://github.com/8051Enthusiast/regex2fat
---

## [12][My blog series about making a web browser from scratch is Rust is finally done.](https://www.reddit.com/r/rust/comments/g2g2gl/my_blog_series_about_making_a_web_browser_from/)
- url: https://joshondesign.com/2020/04/15/next-steps
---

