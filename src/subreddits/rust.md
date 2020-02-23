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

## [3][Rust/WinRT coming soon](https://www.reddit.com/r/rust/comments/f85cp7/rustwinrt_coming_soon/)
- url: https://kennykerr.ca/2020/02/22/rust-winrt-coming-soon/
---

## [4][Programming Language Warts](https://www.reddit.com/r/rust/comments/f7vimo/programming_language_warts/)
- url: https://www.reddit.com/r/rust/comments/f7vimo/programming_language_warts/
---
What are some rust warts ("A small, crocky feature that sticks out of an otherwise clean design.") that you've encountered (Either because they annoyed you in a practical setting or just realized they were there)?

&amp;#x200B;

Bonus points for not well known things. What I have in mind is mostly things that are already in stable rust, that either cannot be fixed without a new edition because of compatibility issues, or that are so core so what makes rust rust that they won't ever be "fixed".

&amp;#x200B;

I'll start with these two:

Arrays do not implement IntoIterator ([https://github.com/rust-lang/rust/issues/25725](https://github.com/rust-lang/rust/issues/25725)). And, currently, calling into\_iter on an array will actually call it on a slice of the array, which gives you references (And can be confusing, specially for beginners).

&amp;#x200B;

There is no good way to sort a &amp;mut VecDeque&lt;T&gt; in std ([https://github.com/rust-lang/rust/issues/27322](https://github.com/rust-lang/rust/issues/27322)). Because std sort operates on slices (Which must be contiguous and from the indices must be in order), you cannot use it to sort some things that are sortable and provide good random access. You can currently collect into a Vec, sort that, and then turn that into a VecDeque (But that is less than optimal), or you can move the VecDeque using From to make a Vec, sort it, and go back to VecDeque, which according to docs will not reallocate, but it involves some additional initial moving of data and also requires ownership of the VecDeque (So it cannot be used on &amp;mut VecDeque&lt;T&gt;)
## [5][Why isn't there a compilation cache in ~/.cache?](https://www.reddit.com/r/rust/comments/f8289h/why_isnt_there_a_compilation_cache_in_cache/)
- url: https://www.reddit.com/r/rust/comments/f8289h/why_isnt_there_a_compilation_cache_in_cache/
---
I've been getting into Rust, and the big thing that stands out to me is that `cargo` likes to build all my dependencies, and all their dependencies, and so on from source. Even though I have an Intel i7-7500U, this takes a significant amount of time.

Even worse, it doesn't cache the final binaries between projects. So if I build 2 different projects that have similar dependencies, it actually rebuilds all their dependencies from source twice, which _really_ hurts in compilation times.

Why doesn't Rust cache final compiled binaries in `~/.cache/rustc` or something? It would also be quite trivial I would guess to implement a least-recently-used (LRU) cache or something to keep the cache directory under a certain size too.

Rust is great, but compilation times hurt badly (I assume that every effort is being made to optimise the Rust compiler's speed).
## [6][Dynamic symbol resolution (Making our own executable packer: Part 7)](https://www.reddit.com/r/rust/comments/f88wl2/dynamic_symbol_resolution_making_our_own/)
- url: https://fasterthanli.me/blog/2020/dynamic-symbol-resolution/
---

## [7][Have you tried to introduce Rust in the workplace? How was your experience?](https://www.reddit.com/r/rust/comments/f83dhr/have_you_tried_to_introduce_rust_in_the_workplace/)
- url: https://www.reddit.com/r/rust/comments/f83dhr/have_you_tried_to_introduce_rust_in_the_workplace/
---
&amp;#x200B;

   For the record, I don't plan to introduce Rust to my workplace.  As I can honestly say I'm not good with Rust (though I hope to greatly improve on that).  However in my journey learning Rust, it reminds me a lot of my journey when I was trying to learn Go around 6 years ago.  And I remember around 4 years ago introducing it my job at the time instead of bootstapping another Java or Python app which is all we did at the time.   Rust isn't a big industry language right now, but that could change in a few years.    I remember there were hardly any Go jobs 4 years ago.  

So I present the following question

Have you been successful in introducing Rust to your workplace?  

If so, any resistance? (PMs, Manager, other co-workers, devs entranced in other languages/technologies)

&amp;#x200B;

If you were able to migrate a project, or bootstrap a new project in Rust, were the results satisfactory?  Did you see benefits in the move?

&amp;#x200B;

If you've thought about introducing Rust to your job, what is the primary driver?  What languages or platforms are most used in your current work/project?
## [8][Discord is not an acceptable choice for Rust](https://www.reddit.com/r/rust/comments/f7qjpb/discord_is_not_an_acceptable_choice_for_rust/)
- url: https://sneak.berlin/20200220/discord-is-not-an-acceptable-choice-for-free-software-projects/
---

## [9][Giron: A javascript parser that outputs estree](https://www.reddit.com/r/rust/comments/f830am/giron_a_javascript_parser_that_outputs_estree/)
- url: https://www.reddit.com/r/rust/comments/f830am/giron_a_javascript_parser_that_outputs_estree/
---
[https://github.com/BlueBlazin/giron](https://github.com/BlueBlazin/giron)

**TL;DR: My first open source project. Lots to improve, hoping for contributions!**

Hi fellow rustaceans,

Giron is an ECMAScript (javascript) parser that outputs JSON in the [ESTree](https://github.com/estree/estree/) format. This is the most popular AST spec for open source JS parsers.

After reading the book I wanted to build something non trivial in rust. So I settled on an ECMAScript parser since I knew Rust compiled very well to WebAssembly.

Anyways, several months later I finally have something that can be put on a public github repo. Currently there is still a lot to be done. So contributions are very welcome!

I should mention that this is my first time working on any open source project. Go easy on me please, lol.

I still have to do lots of testing for correctness, and becoming compliant with the ES2020 spec. Once that is done, we can focus on adding JSX support and optimizations. The wasm module works but I also need to publish it on npm.

Any feedback is welcome. Thank you!
## [10][Code review request: ruut](https://www.reddit.com/r/rust/comments/f84q5e/code_review_request_ruut/)
- url: https://www.reddit.com/r/rust/comments/f84q5e/code_review_request_ruut/
---
Hi!

I recently made my first (non-tutorial led) program in Rust and was wondering if I could
get some feedback?

It’s a CLI utility to easily print `tree(1)`-like diagrams for tree-like structures that are easy to type or easily extracted from elsewhere.

I think the README might do a better job of explaining...

https://github.com/HarrisonB/ruut
## [11][Is working on a wrapper for a C library a good first time project?](https://www.reddit.com/r/rust/comments/f8961n/is_working_on_a_wrapper_for_a_c_library_a_good/)
- url: https://www.reddit.com/r/rust/comments/f8961n/is_working_on_a_wrapper_for_a_c_library_a_good/
---
I just finished reading The book, and looking to build something small as my first project.

I know of a C library, which has a Python wrapper that we use at work. Now I am thinking that writing a Rust wrapper might be a good first project for me to take on...

but I would want to know if this is a good idea for a beginners project? I would like the opinion of more experienced Rust developers...

I mean, would the fact that I am interfacing with a C library bring about it's own unique challenge that won't really help ground my knowledge of Rust? Or would it be accurate to say the skills gained from interfacing with Rust is best left at intermediate level and not as a first project? 

Or you think I should just go for it and the learning experience would be worth it?
## [12][Building a Terminal-Based Finder](https://www.reddit.com/r/rust/comments/f801fo/building_a_terminalbased_finder/)
- url: https://www.reddit.com/r/rust/comments/f801fo/building_a_terminalbased_finder/
---
https://preview.redd.it/9ggco03pxji41.png?width=2628&amp;format=png&amp;auto=webp&amp;s=39710807bea2a3aed848073785a0dd67577df14f

I had an idea to sort of replicate the macOS Finder inside of the shell.  The goal is to make it work across all platforms (though I've only tested on macOS currently).  I'm only 2 weeks into the project and all of my work so far has been laying the foundation down.  It is not yet interactive, but I will be working on that later after finishing a handful of GH Issues.  For now, it simply prints a structure that is similar to that of the Finder.  There are a handful of options that can be set to change the behavior of certain things: different modes for truncating file names (no truncation, truncation based on a constant value, truncation based on statistically measurements such as average and by outlier), the characters used to print the DirectoryContainers can be customized, the vertical and horizontal padding between DirectoryContainers is customizable, and other items as well.  There is color coding based on the type of item (directory, file, symlink), the currently selected item, and the directories in the current path (these colors will be customizable later).

In the end, I want it to be able to do a few things:

1. Be able to launch applications based on the currently selected file, this can be terminal applications or UI applications.
2. Be able to return the currently selected path back to the shell after the application is exited, in the case that someone was just using it to explore files and found an item they want to interact with in the shell outside of TermiFind.

I mainly started this to get back into Rust after a long hiatus, but I could find myself using it.  There is tons of work to be done, but I was wanting to see what others think of it and just get some opinions, received constructive criticism, get advice on some of the GH Issues, or possibly even collect a star or two!

[https://github.com/JosephTLyons/termifind](https://github.com/JosephTLyons/termifind)
