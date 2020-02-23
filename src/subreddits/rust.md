# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (8/2020)!](https://www.reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here_82020/)
- url: https://www.reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here_82020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 326](https://www.reddit.com/r/rust/comments/f6sy2s/this_week_in_rust_326/)
- url: https://this-week-in-rust.org/blog/2020/02/18/this-week-in-rust-326/
---

## [3][Discord is not an acceptable choice for Rust](https://www.reddit.com/r/rust/comments/f7qjpb/discord_is_not_an_acceptable_choice_for_rust/)
- url: https://sneak.berlin/20200220/discord-is-not-an-acceptable-choice-for-free-software-projects/
---

## [4][When you learned Rust, what was your “aha!” moment?](https://www.reddit.com/r/rust/comments/f7l84m/when_you_learned_rust_what_was_your_aha_moment/)
- url: https://www.reddit.com/r/rust/comments/f7l84m/when_you_learned_rust_what_was_your_aha_moment/
---
Interested to know. Rust has a lot of concepts that are not found in other popular languages, or are unique to Rust. I’m wondering what your aha moment was, what did you learn / realise?

One of mine was that the value of a variable is moved on assignment to another variable, aka ownership. I had read about it, I kind of understood it. But then I had a moment where I’m like “I get this”. The implication was profound and I saw that this feature is one reason that makes Rust a safe to program in.
## [5][Rust "The Book"-style syntax highlighting for LaTeX](https://www.reddit.com/r/rust/comments/f7ocdx/rust_the_bookstyle_syntax_highlighting_for_latex/)
- url: https://www.reddit.com/r/rust/comments/f7ocdx/rust_the_bookstyle_syntax_highlighting_for_latex/
---
I was writing some rust documentation in LaTeX and discovered this [previous post](https://www.reddit.com/r/rust/comments/5259oo/rust_syntax_highlighting_for_latex/), but didn't like the requirement to use `-shell-escape` to use the minted package to format rust code blocks.

Instead, I wrote a new language definition (see code below) which formats the code in a similar style to the rust book and wanted to share it here.  Here is what it looks like:

https://preview.redd.it/dctmps0hxei41.png?width=1400&amp;format=png&amp;auto=webp&amp;s=c19e0b23b968ed48abebef5832923443d24b4e3b

Add this to your .tex file:

    \usepackage{listings}
    \definecolor{GrayCodeBlock}{RGB}{241,241,241}
    \definecolor{BlackText}{RGB}{110,107,94}
    \definecolor{RedTypename}{RGB}{182,86,17}
    \definecolor{GreenString}{RGB}{96,172,57}
    \definecolor{PurpleKeyword}{RGB}{184,84,212}
    \definecolor{GrayComment}{RGB}{170,170,170}
    \definecolor{GoldDocumentation}{RGB}{180,165,45}
    \lstdefinelanguage{rust}
    {
        columns=fullflexible,
        keepspaces=true,
        frame=single,
        framesep=0pt,
        framerule=0pt,
        framexleftmargin=4pt,
        framexrightmargin=4pt,
        framextopmargin=5pt,
        framexbottommargin=3pt,
        xleftmargin=4pt,
        xrightmargin=4pt,
        backgroundcolor=\color{GrayCodeBlock},
        basicstyle=\ttfamily\color{BlackText},
        keywords={
            true,false,
            unsafe,async,await,move,
            use,pub,crate,super,self,mod,
            struct,enum,fn,const,static,let,mut,ref,type,impl,dyn,trait,where,as,
            break,continue,if,else,while,for,loop,match,return,yield
        },
        keywordstyle=\color{PurpleKeyword},
        ndkeywords={
            bool,u8,u16,u32,u64,u128,i8,i16,i32,i64,i128,char,str,
            Self,Option,Some,None,Result,Ok,Err,String,Box,Vec,Rc,Arc,Cell,RefCell,HashMap,BTreeMap,
            macro_rules
        },
        ndkeywordstyle=\color{RedTypename},
        comment=[l][\color{GrayComment}\slshape]{//},
        morecomment=[s][\color{GrayComment}\slshape]{/*}{*/},
        morecomment=[l][\color{GoldDocumentation}\slshape]{///},
        morecomment=[s][\color{GoldDocumentation}\slshape]{/*!}{*/},
        morecomment=[l][\color{GoldDocumentation}\slshape]{//!},
        morecomment=[s][\color{RedTypename}]{\#![}{]},
        morecomment=[s][\color{RedTypename}]{\#[}{]},
        stringstyle=\color{GreenString},
        string=[b]",
        morestring=[b]'
    }
## [6][Is Rust taught in universities yet?](https://www.reddit.com/r/rust/comments/f7qb4x/is_rust_taught_in_universities_yet/)
- url: https://www.reddit.com/r/rust/comments/f7qb4x/is_rust_taught_in_universities_yet/
---
I personally think it'd be a great way to teach low-level functional programming, memory management and multithreading. Does anyone know of any university undergrad computer science courses use Rust as as a teaching language? Should they?
## [7][Effort estimation to convert my project to rust](https://www.reddit.com/r/rust/comments/f7jzxu/effort_estimation_to_convert_my_project_to_rust/)
- url: https://www.reddit.com/r/rust/comments/f7jzxu/effort_estimation_to_convert_my_project_to_rust/
---
Hi r/rust,

Author of file manager [`nnn`](https://github.com/jarun/nnn) here. I need some suggestion from you guys to plan my next summer.

I'm a C dev for more than a decade. Learnt Python on the go to to write some of my other projects.

Recently I heard some great things about Rust from a long-term dev friend and got interested in the language. So far I have some basic understanding of how the language works from _The Rust Programming Language_ ebook.

I am planning to start writing a rust version of `nnn` during the next summer vacation. The duration is around a month.

The intention of this post is to understand how much time it may take to convert a project like `nnn` to rust based on your experience. For a single dev at a beginner's level with rust. It would be great if you guys can skim through the source (it's one big file... please don't complain about it as we have kept it that way intentionally so far) or at least the list of current features and share your insights.

Many thanks in advance!
## [8][Learning Rust, but don't know what to write](https://www.reddit.com/r/rust/comments/f7lfvg/learning_rust_but_dont_know_what_to_write/)
- url: https://www.reddit.com/r/rust/comments/f7lfvg/learning_rust_but_dont_know_what_to_write/
---
  Hello.    I'm a long time Go developer and I work with Go professionally.   I've been learning Rust on and off for about a year now and I'm ready to jump into a project to get my feet wet.  Particularly a project that's not too crazy or involved.  Maybe 1k to 2k lines of code.  I'm struggling to find project ideas.  If you don't mind, what are some good project ideas I can work on to get more use to the language?  I really am not interested in web apps, as I basically do backend micro services development for a living, and want to do something a little different.  Thanks in advance for your recommendation and suggestions.
## [9][Just used `dua` on my Macbook with a 250GB hard drive to clean up some space](https://www.reddit.com/r/rust/comments/f7d786/just_used_dua_on_my_macbook_with_a_250gb_hard/)
- url: https://www.reddit.com/r/rust/comments/f7d786/just_used_dua_on_my_macbook_with_a_250gb_hard/
---
Man what an impressive terminal app! 

[https://crates.io/crates/dua-cli](https://crates.io/crates/dua-cli)

Here's what I liked:

\- Blazingly fast

\- Using \`tui\` so it looks nice, much better experience than 90% of terminal applications

\- Found files that the "hard drive" cleanup stuff builtin to MacOS couldn't find

\- The percentage-based approach dua uses to show what files are taking up the most space is super intuitive, makes it really clear what's hogging all of the space

In total, I was able to find well over 100GB of data that I didn't need to be storing, (docker was hogging 60GB) that I didn't even know was there! dua also deleted that data way faster than MacOS can.
## [10][Made a rusty rust wallpaper. The other image was larger than 20 MB (24 Mb) and I don't know how to compress PNG images in photoshop.](https://www.reddit.com/r/rust/comments/f796ds/made_a_rusty_rust_wallpaper_the_other_image_was/)
- url: https://i.redd.it/9t12lnng69i41.png
---

## [11][cargo install, /tmp directory, and noexec](https://www.reddit.com/r/rust/comments/f7pqe3/cargo_install_tmp_directory_and_noexec/)
- url: https://www.reddit.com/r/rust/comments/f7pqe3/cargo_install_tmp_directory_and_noexec/
---
If one has, as I had, this 👇 in one's /etc/fstab :

`tmpfs /tmp tmpfs defaults,nosuid,nodev,`**no**`exec,mode=1777 0 0`

... every build is going to fail - always.

I solved it by removing "no" from "noexec". There are probably some workaround to solve this, but it's better to build in RAM (which I have plenty of, \~ 12 GB)  than hammering one's /home SSD with an unhealthy amount of writes.
## [12][Best way to shrink the lifetime of references explicitly on method calls?](https://www.reddit.com/r/rust/comments/f7s14c/best_way_to_shrink_the_lifetime_of_references/)
- url: https://www.reddit.com/r/rust/comments/f7s14c/best_way_to_shrink_the_lifetime_of_references/
---
I'd love to have something like:

    let mut value = MyValue::new();
    let mut obj = MyObj::new();
    obj.do_something::&lt;'a&gt;(&amp;'a mut value);  // like this
    obj.do_other_things();

More specifically, I want to make sure *without looking for the method definition* that the reference passed to `do_something` method won't be used outside of its scope, e.g., not in `do_other_things` method. I found this is possible with a dummy binding after the method call:

    let mut value = MyValue::new();
    let mut obj = MyObj::new();
    obj.do_something(&amp;mut value);
    let _ = &amp;value  // here
    obj.do_other_things();

but this is no good because the original intension is unclear. Is there better way to achieve this? Thanks in advance.
