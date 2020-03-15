# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (11/2020)!](https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 329](https://www.reddit.com/r/rust/comments/fhcuec/this_week_in_rust_329/)
- url: https://this-week-in-rust.org/blog/2020/03/10/this-week-in-rust-329/
---

## [3][Kondo 🧹 v0.3 released, now with GUI 🎉!](https://www.reddit.com/r/rust/comments/fixsyj/kondo_v03_released_now_with_gui/)
- url: https://i.redd.it/q80b5uregsm41.png
---

## [4][Out of the Box Dynamic Dispatch](https://www.reddit.com/r/rust/comments/fiprk7/out_of_the_box_dynamic_dispatch/)
- url: https://llogiq.github.io/2020/03/14/ootb.html
---

## [5][Reducing the size of `Option&lt;T&gt;` by adding members to `T`?](https://www.reddit.com/r/rust/comments/fimbo3/reducing_the_size_of_optiont_by_adding_members_to/)
- url: https://www.reddit.com/r/rust/comments/fimbo3/reducing_the_size_of_optiont_by_adding_members_to/
---
Suppose I have a struct `BoolPadded`

    struct BoolPadded {
        _x: u16,
        _y: u8,
        _z: bool,
    }

Then `Option&lt;BoolPadded&gt;` takes up 4 bytes in memory. I assume this is because the compiler is smart enough to use the invalid bit patterns of `bool` to represent the enum tag of `Option&lt;T&gt;`.

However, if I remove the bool member:

    struct Unpadded {
        _x: u16,
        _y: u8,
    }

Now `Option&lt;Unpadded&gt;` takes up 6 bytes in memory. It seems as though now the compiler can't determine that `Unpadded`, when aligned at a 2-byte boundary, will have unused padding bits at the end.

Is there some sort of memory layout or ABI stability guarantee that is preventing this kind of optimization?

[Playground Example](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=14ff93f1a0b204cba11bb3b47c7a9c58)
## [6][Flume, a 100% safe MPSC that's faster than std and gives crossbeam a run for its money](https://www.reddit.com/r/rust/comments/fj17z6/flume_a_100_safe_mpsc_thats_faster_than_std_and/)
- url: https://github.com/zesterer/flume
---

## [7][Redox OS - pkgar introduction](https://www.reddit.com/r/rust/comments/fim6ju/redox_os_pkgar_introduction/)
- url: https://www.redox-os.org/news/pkgar-introduction/
---

## [8][Parsing Library in Rust pt. 2 - Testing](https://www.reddit.com/r/rust/comments/fj086p/parsing_library_in_rust_pt_2_testing/)
- url: https://blog.frondeus.pl/parser-2/
---

## [9][StaticVec 0.9.0: now with a fixed-capacity StaticHeap based on std::collections::BinaryHeap, a const-context-compatible macro constructor for StaticString, const trait impls, and more!](https://www.reddit.com/r/rust/comments/fis6dd/staticvec_090_now_with_a_fixedcapacity_staticheap/)
- url: https://github.com/slightlyoutofphase/staticvec
---

## [10][TpNote - fast note talking with templates and filename-synchronization V1.0.0 is out](https://www.reddit.com/r/rust/comments/fiyq5s/tpnote_fast_note_talking_with_templates_and/)
- url: https://www.reddit.com/r/rust/comments/fiyq5s/tpnote_fast_note_talking_with_templates_and/
---
Markup languages like *Markdown*, *ReStructuredText*, *asciidoc*, *textile*,
*txt2tags* or *mediawiki* are perfectly suited for fast note-taking. Type
your notes with your favourite editor and chose your favourite markup
language[^1].

_Tp-Note_ helps you to quickly get started writing notes with its powerful
template system. As _Tp-Note_ takes care that the note's filename is always
synchronized with its document title, you will find back your notes easily.

_Tp-Note_ is available for Linux, Windows and iOS. The manual illustrates
its main use-cases and how to get started:

1. Fast start note-taking (when the lecture starts).
2. Take a note about an existing (downloaded) file.
3. Bookmark and comment a hyperlink.

If you want to customize _Tp-Note_ with own templates, another markup
language, please consult the [man-page] for more technical details.

The project is hosted on Github:
[getreu/tp-note](https://github.com/getreu/tp-note). The project's webpage is on
[http://blog.getreu.net](http://blog.getreu.net/projects/tp-note/).
The documentation of this project is dived into tow parts:

* User manual

  [tp-note user manual - html](https://blog.getreu.net/projects/tp-note/tp-note--manual.html)\
  [tp-note user manual - pdf](https://blog.getreu.net/_downloads/tp-note--manual.pdf)

* Unix man-page (more technical)

  [tp-note manual page - html](https://blog.getreu.net/projects/tp-note/tp-note--manpage.html)\
  [tp-note manual page - pdf](https://blog.getreu.net/_downloads/tp-note--manpage.pdf)

Download _tp-note_:

* Binaries and packages:

  - Executable for Windows:

    [x86_64-pc-windows-gnu/release/tp-note.exe](https://blog.getreu.net/projects/tp-note/_downloads/x86_64-pc-windows-gnu/release/tp-note.exe)

  - Binary for Linux:

    [x86_64-unknown-linux-gnu/release/tp-note](https://blog.getreu.net/projects/tp-note/_downloads/x86_64-unknown-linux-gnu/release/tp-note) \
    [x86_64-unknown-linux-musl/release/tp-note](https://blog.getreu.net/projects/tp-note/_downloads/x86_64-unknown-linux-musl/release/tp-note)

  - Package for Debian and Ubuntu:

    [x86_64-unknown-linux-gnu/debian/tp-note_1.0.0_amd64.deb](https://blog.getreu.net/projects/tp-note/_downloads/x86_64-unknown-linux-gnu/debian/tp-note_1.0.0_amd64.deb)

As this project follows [Semantic Versioning](https://semver.org/), version 1.0.0 comes with a stable API.
## [11][Optimizations That Aren't, or are they?](https://www.reddit.com/r/rust/comments/fill1q/optimizations_that_arent_or_are_they/)
- url: https://oribenshir.github.io/afternoon_rusting/blog/copy-on-write
---

## [12][Tuple structs are functions](https://www.reddit.com/r/rust/comments/firk1r/tuple_structs_are_functions/)
- url: https://www.reddit.com/r/rust/comments/firk1r/tuple_structs_are_functions/
---
In a project I'm working on I have a few lines of code like this:

    str::from_utf8(&amp;string_bytes).map_err(
        |utf8_err| DeserializationError::InvalidUTF8(utf8_err),
    )?

where `DeserializationError` is an enum and `InvalidUTF8` is a tuple variant. Clippy corrected this to the following, which also compiles:

    str::from_utf8(&amp;string_bytes)
        .map_err(DeserializationError::InvalidUTF8)?

`map_err` takes a `FnOnce(E) -&gt; F` , so I was curious why the name of the enum variant also works. Is the construction of a tuple variant like calling a function? Yes, as it turns out!

From [this 2017 thread](https://users.rust-lang.org/t/enum-variants-and-newtype-as-functions-closures/14706) in the rust-lang users forum:

&gt;Unit structs, newtypes and tuple structs all live simultaneously in both  the type/module namespace and the value/function namespace.  Only standard, braced structs live exclusively in the type/module namespace.

    struct Unit; // exists as a value
    struct Unit2(); // exists in value namespace as a 0-arg function
    struct Newtype(u8); // exists in value namespace as a 1-arg function
    struct Tuple(u8, u8); // exists in value namespace as a 2-arg function
    struct Regular { value: u8 }
    
    // fn Unit() {} // the name `Unit` is defined multiple times
    // fn Unit2() {} // the name `Unit2` is defined multiple times
    // fn Newtype() {} // the name `Newtype` is defined multiple times
    // fn Tuple() {} // the name `Tuple` is defined multiple times
    fn Regular() {} // ok

Thought that was a neat syntactical coincidence, though it tripped me up for a second.
