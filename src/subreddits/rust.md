# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (36/2020)!](https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (36/2020)?](https://www.reddit.com/r/rust/comments/ijvxwz/whats_everyone_working_on_this_week_362020/)
- url: https://www.reddit.com/r/rust/comments/ijvxwz/whats_everyone_working_on_this_week_362020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-36-2020/48100?u=llogiq)!
## [3][const_fn makes it too easy to do mandelbrots](https://www.reddit.com/r/rust/comments/ijpxz2/const_fn_makes_it_too_easy_to_do_mandelbrots/)
- url: https://www.reddit.com/r/rust/comments/ijpxz2/const_fn_makes_it_too_easy_to_do_mandelbrots/
---
It's about 2-3 years since I did my compile time mandelbrot using associated types the hard way ( https://old.reddit.com/r/rust/comments/6zgz3k/rust_compiletime_mandelbrot/ ) - I just rewrote it using const fn; and it's *boring*:

    // (c) Dave Gilbert dave@treblig.org 2020
    const LSIZE: usize = 32;
    const SSIZE: usize = LSIZE*(LSIZE+1);
 
    const fn manpoint(x: i64, y: i64) -&gt; u8 {
        let mut ti : i64 = 0;
        let mut tr : i64 = 0;
        let mut zi : i64 = 0;
        let mut zr : i64 = 0;
        let mut iter : u8 = 128;
 
        while iter &gt; 0  &amp;&amp; (ti + tr &lt; (4 &lt;&lt; 16)) {
            zi = (((((2&lt;&lt;16)*zr) &gt;&gt; 16) * zi)&gt;&gt;16) + y;
            zr = tr-ti+x;
            tr = (zr*zr) &gt;&gt; 16;
            ti = (zi*zi) &gt;&gt; 16;
            iter-=1;
        }
        (iter % 20) + 64
    }
 
    const fn manline() -&gt; [u8;SSIZE] {
        let mut y : i64 = -2 &lt;&lt; 16;
        let xstep : i64 = (4 &lt;&lt; 16) / (LSIZE as i64);
        let ystep : i64 = (4 &lt;&lt; 16) / (LSIZE as i64);
        let mut yi = 0;
        let mut res : [u8;SSIZE] = [32;SSIZE];
 
        while yi &lt; LSIZE {
            let mut x : i64 = -2 &lt;&lt; 16;
            let mut xi = 0;
            while xi &lt; LSIZE {
                res[yi*(LSIZE+1)+xi] = manpoint(x, y);
                x+=xstep;
                xi+=1;
            }
            res[yi*(LSIZE+1)+LSIZE]=10;
            y+=ystep;
            yi+=1;
        }
 
        res
    }
 
    const BROT: [u8;SSIZE] = manline();
    fn main() {
        println!("Const val={}", std::str::from_utf8(&amp;BROT).unwrap());
    }
 
rustc -O const.rs --emit asm; grep -i ascii const.s|tr '\\' '\012'
	.ascii	"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG
    nGGGGGGGGGGGFFFFFFFFFFFGGGGGGGGGG
    nGGGGGGGGGFFFFFFFFFFFFFFFGGGGGGGG
    nGGGGGGGFFFFFFFFFFFFFFFFFFFGGGGGG
    nGGGGGGFFFFFFFFFFFFFFFFFFFFFGGGGG
    nGGGGGFFFFFFFFFFFFFFFFFFFFFFFGGGG
    nGGGGFFFFFEEEEEEEFFFFFFFFFFFFFGGG
    nGGGFFFEEEEEEEDCCDEEFFFFFFFFFFFGG
    nGGGFFEEEEEEDDCBO@DDEFFFFFFFFFFGG
    nGGFEEEEEEEDDDCO@SCDDEEFFFFFFFFFG
    nGGFEEEEEEDDCBAF@NACDEEFFFFFFFFFG
    nGFEEEEEDDCCSHH@@@@KQDEEFFFFFFFFF
    nGEEEEEDCCCBP@@@@@@@SCEEEFFFFFFFF
    nGEEDDCNA@A@@@@@@@@@KCDEEFFFFFFFF
    nGDDDCBS@@DO@@@@@@@@@CDEEFFFFFFFF
    nGDDBAQI@@@B@@@@@@@@KCDEEFFFFFFFF
    nG@@@@@@@@@@@@@@@@@@ACDEEFFFFFFFF
    nGDDBAQI@@@B@@@@@@@@KCDEEFFFFFFFF
    nGDDDCBS@@DO@@@@@@@@@CDEEFFFFFFFF
    nGEEDDCNA@A@@@@@@@@@KCDEEFFFFFFFF
    nGEEEEEDCCCBP@@@@@@@SCEEEFFFFFFFF
    nGFEEEEEDDCCSH@@@@@JQDEEFFFFFFFFF
    nGGFEEEEEEDDCBAF@NACDEEFFFFFFFFFG
    nGGFEEEEEEEDDDCO@SCDDEEFFFFFFFFFG
    nGGGFFEEEEEEDDCBO@DDEFFFFFFFFFFGG
    nGGGFFFEEEEEEEDCCDEEFFFFFFFFFFFGG
    nGGGGFFFFFEEEEEEEFFFFFFFFFFFFFGGG
    nGGGGGFFFFFFFFFFFFFFFFFFFFFFFGGGG
    nGGGGGGFFFFFFFFFFFFFFFFFFFFFGGGGG
    nGGGGGGGFFFFFFFFFFFFFFFFFFFGGGGGG
    nGGGGGGGGGFFFFFFFFFFFFFFFGGGGGGGG
    nGGGGGGGGGGGFFFFFFFFFFFGGGGGGGGGG
    n"
## [4][rust-analyzer changelog #40](https://www.reddit.com/r/rust/comments/ijwdcc/rustanalyzer_changelog_40/)
- url: https://rust-analyzer.github.io/thisweek/2020/08/31/changelog-40.html
---

## [5][Headcrab: August 2020 progress report](https://www.reddit.com/r/rust/comments/ijsa4c/headcrab_august_2020_progress_report/)
- url: https://headcrab.rs/2020/08/31/august-update.html
---

## [6][Rust in Action achieves #1 New Release in "Parallel Computer Programming" on Amazon](https://www.reddit.com/r/rust/comments/ijjofj/rust_in_action_achieves_1_new_release_in_parallel/)
- url: https://i.redd.it/ijp1at2a77k51.png
---

## [7][Rust explained using easy English so second language speakers can learn it too (now completed)](https://www.reddit.com/r/rust/comments/ijd1gs/rust_explained_using_easy_english_so_second/)
- url: https://github.com/Dhghomon/easy_rust/blob/master/README.md
---

## [8][As above, so below, part 2. Bare metal Rust generics.](https://www.reddit.com/r/rust/comments/ijwrxd/as_above_so_below_part_2_bare_metal_rust_generics/)
- url: https://www.ecorax.net/as-above-so-below-2/
---

## [9][Starframe devlog: Architecture (ECS, Graph)](https://www.reddit.com/r/rust/comments/iju3xq/starframe_devlog_architecture_ecs_graph/)
- url: https://moletrooper.github.io/blog/2020/08/starframe-1-architecture/
---

## [10][best way of performance measuring in rust](https://www.reddit.com/r/rust/comments/ijvpez/best_way_of_performance_measuring_in_rust/)
- url: https://www.reddit.com/r/rust/comments/ijvpez/best_way_of_performance_measuring_in_rust/
---
Hi,

I'm quite new to rust and high performance computing in general. For learn- and practice reasons, I am looking for the best way to measure the performance of an algorithm/code snippet, rework it and visualize the improvement (e.g. on a grafana board). 

And how do I make sure, that the performance measuring is not affecting the runtime of the whole program? Is there a way to leave the measuring-code in the later production code, or do i have to remove it?

Thanks for your help!
## [11][Zero To Production #3.5: HTML forms (actix-web/serde), Databases (sqlx), Integration tests](https://www.reddit.com/r/rust/comments/ijosc8/zero_to_production_35_html_forms_actixwebserde/)
- url: https://www.lpalmieri.com/posts/2020-08-31-zero-to-production-3-5-html-forms-databases-integration-tests/
---

## [12][What’s the deal with LTO?](https://www.reddit.com/r/rust/comments/ijwya5/whats_the_deal_with_lto/)
- url: https://www.reddit.com/r/rust/comments/ijwya5/whats_the_deal_with_lto/
---
Below I’ll make a few assertions; please tell me if they’re wrong or inaccurate. Thanks for the help!

- LTO allows optimisations across crates
- There are three settings: `off`, `fat`, `thin` (the default)
- `fat` takes much longer to compile than `thin`, and usually (?) gives better speeds
- With `fat` LTO everything is essentially marked `#[inline]`, and so there is no need to use it any more
