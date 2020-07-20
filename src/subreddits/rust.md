# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (29/2020)!](https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.45]](https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://old.reddit.com/r/rust/comments/gz8ut5/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).
# Rules for individuals:

 * Don't create top-level comments; those are for employers.
 * Feel free to reply to top-level comments with on-topic questions.
 * I will create a stickied top-level comment for individuals looking for work.
 * I will create an additional top-level comment for meta discussion.

# Rules for employers:

 * To find individuals seeking work, see the replies to the stickied top-level comment; you will need to click the "more comments" link at the bottom of the top-level comment in order to make these replies visible.
 * To make a top-level comment you must be hiring directly; no third-party recruiters.
 * One top-level comment per employer. If you have multiple job openings, please consolidate their descriptions or mention them in replies to your own top-level comment.
 * Proofread your comment after posting it and edit it if necessary to correct mistakes.
 * Please base your comment on the following template:

COMPANY: *[Company name; please link to your company's website or careers page.]*

TYPE: *[Full time, part time, internship, contract, etc.]*

DESCRIPTION: *[What does your company do, and what are you using Rust for? How much experience are you seeking and what seniority levels are you hiring for? The more details the better.]*

LOCATION: *[Where are your office or offices located? If your workplace language isn't English, please specify it.]*

REMOTE: *[Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]*

VISA: *[Does your company sponsor visas?]*

CONTACT: *[How can someone get in touch with you?]*
## [3][Blog post: Three Architectures for a Responsive IDE](https://www.reddit.com/r/rust/comments/huja76/blog_post_three_architectures_for_a_responsive_ide/)
- url: https://rust-analyzer.github.io/blog/2020/07/20/three-architectures-for-responsive-ide.html
---

## [4][Dynasm-rs v1.0.0 released targetting 1.45 stable!](https://www.reddit.com/r/rust/comments/hujimd/dynasmrs_v100_released_targetting_145_stable/)
- url: https://www.reddit.com/r/rust/comments/hujimd/dynasmrs_v100_released_targetting_145_stable/
---
Dynasm-rs is a dynamic assembler for rust, providing easy low-level dynamic code generation, currently targetting the x86, AMD64 and aarch64 instruction sets. The crate has been in development for a while and started as a compiler plugin, but thanks to the proc_macro_hygiene stabilizations in rustc 1.45 it finally works with the stable release of rustc. The example below shows what kind of code you can write using this library, but it supports much more complex use cases with ease, as it is used in several leading wasm runtimes.

* [Release notes](https://github.com/CensoredUsername/dynasm-rs/blob/master/doc/releasenotes.md)

* [Documentation](https://censoredusername.github.io/dynasm-rs/language/index.html)

* [Repository](https://github.com/CensoredUsername/dynasm-rs)

Additional thanks go out to the guys at [Wasmer](https://wasmer.io/) for sponsoring the development of the aarch64 backend.

The following is a simple example that prints "Hello world" on an x64 machine:

    use dynasmrt::{dynasm, DynasmApi, DynasmLabelApi};
    use std::{io, slice, mem};
    use std::io::Write;

    fn main() {
        let mut ops = dynasmrt::x64::Assembler::new().unwrap();
        let string = "Hello World!";

        dynasm!(ops
            ; .arch x64
            ; -&gt;hello:
            ; .bytes string.as_bytes()
        );

        let hello = ops.offset();
        dynasm!(ops
            ; .arch x64
            ; lea rcx, [-&gt;hello]
            ; xor edx, edx
            ; mov dl, BYTE string.len() as _
            ; mov rax, QWORD print as _
            ; sub rsp, BYTE 0x28
            ; call rax
            ; add rsp, BYTE 0x28
            ; ret
        );

        let buf = ops.finalize().unwrap();

        let hello_fn: extern "win64" fn() -&gt; bool = unsafe { mem::transmute(buf.ptr(hello)) };

        assert!(hello_fn());
    }

    pub extern "win64" fn print(buffer: *const u8, length: u64) -&gt; bool {
        io::stdout()
            .write_all(unsafe { slice::from_raw_parts(buffer, length as usize) })
            .is_ok()
    }
## [5][Rust Analyzer Changelog #34](https://www.reddit.com/r/rust/comments/hukdvz/rust_analyzer_changelog_34/)
- url: https://rust-analyzer.github.io/thisweek/2020/07/20/changelog-34.html
---

## [6][Clear explanation of Rust’s module system](https://www.reddit.com/r/rust/comments/htzkq7/clear_explanation_of_rusts_module_system/)
- url: http://www.sheshbabu.com/posts/rust-module-system/
---

## [7][Show r/rust: dijo — scriptable, curses-based, digital habit tracker](https://www.reddit.com/r/rust/comments/hu38zj/show_rrust_dijo_scriptable_cursesbased_digital/)
- url: https://i.redd.it/ikku5apvbub51.png
---

## [8][Rust and it's Orphan Rules [Article by Michael Gattozzi]](https://www.reddit.com/r/rust/comments/huhzjp/rust_and_its_orphan_rules_article_by_michael/)
- url: https://blog.mgattozzi.dev/orphan-rules/
---

## [9][Native-Windows-GUI is back and ready to kick some ass (1.0 release)](https://www.reddit.com/r/rust/comments/hu1s7v/nativewindowsgui_is_back_and_ready_to_kick_some/)
- url: https://www.reddit.com/r/rust/comments/hu1s7v/nativewindowsgui_is_back_and_ready_to_kick_some/
---
**Native-Windows-Gui** (aka NWG) is a windows GUI library that runs on top of Win32. Its main goal is to provide a rust-first API that is dead simple to use, keeps compile time low, does not bloat the final binary size, and provides a large number of features. As the first stable release, all those goals are met. 

NWG also comes with a sister project, **Native-Windows-Derive** (also released as 1.0) that uses a procedural macro to build a GUI from a struct, removing almost all the boilerplate. See [this example](https://github.com/gabdube/native-windows-gui#with-native-windows-derive)

NWG is best suited for small to medium size projects that don't require fancy stuff (like animations or custom drawing). 

NWG doesn't enforce any kind of structure, you instance resources and controls the same way you use any other Rust types. That said, NWG defines 2 traits (PartialUI and NativeUI) that makes it easy to manage your GUI applications. All the examples use them.

As you'd expect from a good open source library, NWG comes with a good documentation and plenty of examples. See the links at the end of this post.

[^Follow ^me ^on ^twitter ^while ^you're ^at ^it](https://twitter.com/gdube_dev) 

Here's an overview of the features:

- NWG implements **30** different controls (also called widgets). From simple buttons to tree-views and data grid 
- Menus and menu bar
- Image and font resource (BMP, ICO, CUR, PNG *, GIF *, JPG *, TIFF *, DDS *)
  - * Extended image formats with the Windows Imaging Component (WIC).
- Localization support
  - Wraps the Windows National Language Support (see [Locale](https://docs.rs/native-windows-gui/1.0.1/native_windows_gui/struct.Locale.html)
- Tooltip
- System tray notification
- Cursor handling
- A full , mostly unsafe, clipboard wrapper
- Partial templates support to split large application into chunks
- Dynamic controls support
  - Add/Remove controls at runtime
  - Bind or unbind new events at runtime
- Multithreaded application support
  - Communicate to the GUI thread from another thread
  - Run multiple window on different threads
- Simple layout configurations
  - FlexboxLayout
  - GridLayout
- Drag and drop
  - Drop files from the desktop to a window
- The most common dialog boxes
  - File dialog (save, open, open folder)
  - Font dialog
  - Color dialog
- An canvas that can be used by external rendering APIs
- High-DPI aware
- Support for accessibility functions
- Tab naviguation
- Support for low level system message capture (HWND, MSG, WPARAM, LPARAM)

**Future Development**

NWG 1.0 implements almost every GUI features of win32 and the API has been frozen. 

Short term, the development will be focused on fixing bugs, unsafy behavior, and implementing obscure features that I skipped over. 

Long term, if the library has some success, I'd like work on a canvas api (based on direct 2D), and maybe a WYSIWYG editor. 

**Links**

The github project: https://github.com/gabdube/native-windows-gui

The main documentation: https://gabdube.github.io/native-windows-gui/native-windows-docs/index.html

Docs.rs: https://docs.rs/native-windows-gui/1.0.1/native_windows_gui/

Showcase: https://github.com/gabdube/native-windows-gui/tree/master/showcase
## [10][Generated binary contains irrelevant strings](https://www.reddit.com/r/rust/comments/hufuj1/generated_binary_contains_irrelevant_strings/)
- url: https://www.reddit.com/r/rust/comments/hufuj1/generated_binary_contains_irrelevant_strings/
---
Hi everyone, I'm new to rust and generated binary by

    cargo build --release

and then strip.

strings executed on the binary reveals that it contains entries that are absolutely irrelevant to the project I'm working on. A lot of strings taken from file which contains \\0 separated strings and exists in $HOME/Downloads: this is some dictionary file that's why I immediatelly spotted that some entries are transferred from that file into my binary. Interesting though that file was a subject of my previous project. But no dependency to that project in Cargo.toml of current one. Also naïvely removing that file does not prevent that crap to happen.

Doesn't affect everything else but I would like to know where those strings are taken from and how it is possible. Anyway it is about 10Kb of unrelated stuff that makes me like a mad.

Thanks in advance.

Cargo.toml:  


    [package]
    name = "service"
    version = "0.1.0"
    authors = ["albazzaa"]
    edition = "2018"
    
    [profile.release]
    lto = true
    
    [dev-dependencies]
    criterion = "0.3.3"
    
    [[bench]]
    name = "s1_benchmark"
    harness = false
    
    [dependencies]
    diesel = { version = "1.4.5", features = ["mysql", "r2d2", "chrono", "uuid"] }
    dotenv = "0.15.0"
    dotenv_codegen = "0.15.0"
    chrono = { version="0.4.11", features=["serde"] }
    actix= "0.9.0"
    actix-web = { version = "2.0.0", features = ["openssl"] }
    actix-web-httpauth = "0.4.1"
    actix-sled-cache = "0.2.0"
    actix-rt = "1.1.1"
    actix-ratelimit = "0.2.1"
    serde = "1.0.105"
    openssl = { version = "0.10", features = ["v110"] }
    uuid = { version="0.8.1", features=["v4", "serde"] }
    serde_json = "1.0.55"
    log = "0.4.8"
    env_logger = "0.7.1"
    rust-argon2 = "0.8.2"
    derive_more = "0.99.9"
    actix-cors = "0.2.0"
    jwt = "0.9.0"
    hmac = "0.8.1"
    sha2 = "0.9.1"
    futures-util = "0.3.5"

&amp;#x200B;
## [11][First published create (step one of many)](https://www.reddit.com/r/rust/comments/hud0f8/first_published_create_step_one_of_many/)
- url: https://www.reddit.com/r/rust/comments/hud0f8/first_published_create_step_one_of_many/
---
https://crates.io/crates/knitting_parse

My wife is really into knitting. Patterns come with a DSL'esk representation. They also sometimes come with graphical charts.

These charts are stupid to make. Some people screenshot excel documents where they make them. I had previously made an IDE in QT that you input the DSL, and it generated the chart. This worked, but was impossibly hard to distribute between her friends getting working Windows, Mac, and Linux versions.

I have now decided to try to redo this in rust. I've yet to decide on web or make LSP and integrate into existing IDE.

This is the start of a parser for the DSL'esk part.

As rust is just a very part time hobby, not sure the quality. One thing I am unsure on is the width on stitch. I'd like it to be an enum to match in the chart maker part. If it were a struct, that width would be much easier.

It is using nom 5 with all functions (no macros).
## [12][Rust Research Survey](https://www.reddit.com/r/rust/comments/hua7nz/rust_research_survey/)
- url: https://www.reddit.com/r/rust/comments/hua7nz/rust_research_survey/
---
We invite you to participate in a research study being conducted at University of Maryland, College Park. We are interested in the use and adoption of Rust. Our research relies on the generous help of interested developers! The online survey will take about 30 minutes to complete. It consists of questions about your general programming experience as well as your experience with using Rust. 

&amp;#x200B;

You are eligible for the study if you:

1) Are at least 18 years old

2) Have Rust development experience (especially in a professional setting)

3) Are comfortable completing the study in English

&amp;#x200B;

If you are interested in completing the study and are eligible, please click here to complete the study:

[https://umdsurvey.umd.edu/jfe/form/SV\_87x3o1z8AN54YNn](https://umdsurvey.umd.edu/jfe/form/SV_87x3o1z8AN54YNn)

&amp;#x200B;

Unfortunately, we cannot compensate you for participating in this study, but if you're interested, we will be happy to share more information about our research with you. If you have further questions, fell free to contact one our researchers at [kfulton@cs.umd.edu](mailto:kfulton@cs.umd.edu).  


Thank you in advance for your consideration!
