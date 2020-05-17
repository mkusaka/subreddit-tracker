# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (20/2020)!](https://www.reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.43]](https://www.reddit.com/r/rust/comments/g6v98u/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/g6v98u/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/).

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
## [3][Show r/rust: A serde format that uses a subset of Markdown](https://www.reddit.com/r/rust/comments/glc6p1/show_rrust_a_serde_format_that_uses_a_subset_of/)
- url: https://github.com/PurpleMyst/serde-mml
---

## [4][Mun v0.2.0 Released](https://www.reddit.com/r/rust/comments/gkzmab/mun_v020_released/)
- url: https://www.reddit.com/r/rust/comments/gkzmab/mun_v020_released/
---
After half a year of hard work by the Mun Community and Core Team, we are happy to present [Mun v0.2.0](https://github.com/mun-lang/mun/releases/tag/v0.2.0). With the addition of hot reloadable structs, this forms a big step forward in terms of hot reloadability for the language.

**What's Mun?**

[Mun](https://mun-lang.org) is an embeddable programming language empowering creation through iteration. The idea to create Mun originated out of frustration with the Lua dynamic scripting language and a desire to have similar hot reloading functionality available in Rust. As such, it's not a direct competitor with Rust, but instead is intended to be used with Rust (or C/C++) as a host/embedded language pairing. Actually, Mun is completely written in Rust, building on similar crates as [rust-analyzer](https://github.com/rust-analyzer/rust-analyzer) and [rustc](https://github.com/rust-lang/rust). Its key features include:

* **Ahead of time compilation** \- Mun is compiled ahead of time (AOT), as opposed to being interpreted or compiled just in time (JIT).
* **Statically typed** \- Mun resolves types at compilation time instead of at runtime, resulting in immediate feedback when writing code and opening the door for powerful refactoring tools.
* **First class hot-reloading** \- Every aspect of Mun is designed with hot reloading in mind. Hot reloading is the process of changing code and resources of a live application, removing the need to start, stop and recompile an application whenever a function or value is changed.

**What's new?**

Mun v0.2 brings a *ton* of updates. For a full list have a look at the [changelog](https://github.com/mun-lang/mun/releases/tag/v0.2.0), but the main new features are:

* Hot reloadable data structures
* Marshalling of data structures to Rust, C, and C++
* Garbage collection for data structures (with opt-out at the struct-level)
* `loop`, `while`, `break` and explicit`return` expressions
* Full operator support for numeric and boolean types
* Incremental compilation
* Benchmark support

**What's next?**

Mun is still in early development, so there is a lot to do. The Mun Core Team operates according to a tick-tock release cycle. Development of a tick release (Mun v0.2) focuses on implementing new features, whereas a tock release (Mun v0.3) focuses on using and improving existing features.

In particular, we'll focus on making a Rust+Mun game demo that showcases Mun's hot reloading capabilities and we'll be working on a `cargo`\-like application for better Mun project management. For a complete roadmap, please visit our [GitHub Project](https://github.com/mun-lang/mun/projects/1).

As Mun consists of a separate language, runtime, and an ABI that forms the communication protocol between the two; any language can benefit from the Mun Runtime's hot reloading capabilities. This leaves room for others to integrate the Mun ABI with (a semantic subset) of other languages such as Rust.

**How can I learn more?**

For more information, please visit our [website](https://mun-lang.org) or read our [blog](https://mun-lang.org/blog).

To get started, read the [Mun Book](https://docs.mun-lang.org/v0.2) and have a look at our [Rust](https://github.com/mun-lang/mun/tree/master/crates/mun_runtime/examples) and [C++](https://github.com/mun-lang/example-cpp) examples.

If you are interested in helping develop Mun feel free to reach out to us on [Discord](https://discord.gg/SfvvcCU) or [Twitter](https://twitter.com/munlangorg), or pick up one of our [good first issues](https://github.com/mun-lang/mun/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22) on [GitHub](https://github.com/mun-lang/mun).

If you cannot personally contribute but would still like to support our cause, please consider donating to our [Open Collective](https://opencollective.com/mun).

**Acknowledgments**

This release would not have been possible without the incredible contributions from the Mun Community and the generous support of Mozilla - in the form of the [MOSS grant](https://www.mozilla.org/en-US/moss/mission-partners). We are extremely grateful to all of you!
## [5][Projects to do after Rust Book](https://www.reddit.com/r/rust/comments/glczk5/projects_to_do_after_rust_book/)
- url: https://www.reddit.com/r/rust/comments/glczk5/projects_to_do_after_rust_book/
---
I've gone through the Rust Book. Finally finished it and made it to the webserver project. Coming from a good background with Golang, after learning Go with Tour of Go, I could do many things. But after finishing rust book, I feel completely lost. Don't know what to do next.

&amp;#x200B;

I would like to do some project-based learning, so i thought to reach rust community for some advice about how to get going.
## [6][actix-casbin-auth: Casbin actix-web access control middleware](https://www.reddit.com/r/rust/comments/gl9l6q/actixcasbinauth_casbin_actixweb_access_control/)
- url: https://github.com/casbin-rs/actix-casbin-auth
---

## [7][Good ideas for a intermediate-level cli project?](https://www.reddit.com/r/rust/comments/gldqs9/good_ideas_for_a_intermediatelevel_cli_project/)
- url: https://www.reddit.com/r/rust/comments/gldqs9/good_ideas_for_a_intermediatelevel_cli_project/
---
Im new to rust but have been programming in languages like python and c for the past 3 to 4 years. Im looking to learn rust in more detail by starting an intermediate-level cli project but don't really have any good ideas for what i could build. 

What are some of the cool cli projects you guys have written?
## [8][Rust needs a viable puppeteer (headless Chrome)](https://www.reddit.com/r/rust/comments/glep7w/rust_needs_a_viable_puppeteer_headless_chrome/)
- url: https://www.reddit.com/r/rust/comments/glep7w/rust_needs_a_viable_puppeteer_headless_chrome/
---
It seems that rust-headless-chrome has been abandoned by its author.  There's a demand for a Rust puppeteer solution and this library has thus far been it.  The project is useful in its current form but is not feature complete and continues to receive ignored PRs and issues.

This is a very useful project for the Rust ecosystem and worth continuing either in its current architecture or a new async one.
Depending on how things go this year, I may take up the baton, but in the short term I cannot commit.
## [9][Suddenly compile errors on Windows](https://www.reddit.com/r/rust/comments/gle181/suddenly_compile_errors_on_windows/)
- url: https://www.reddit.com/r/rust/comments/gle181/suddenly_compile_errors_on_windows/
---
I am working on a crate but suddenly, compilation fails on 2 of my windows computers. On a macbook the project is compiling fine. I left the project 2 weeks untouched (at that time, it was compiling) but today, I get compilation errors.  Note: this only happens with toolchain 'stable-x86\_64-pc-windows-gnu'. Toolchain 'stable-x86\_64-pc-windows-msvc' compiles the code just fine (but I can not use the debugger, so it is important that the gnu toolchain also works)  

Steps to reproduce:

1. Clone  [https://github.com/Jasperav/cdrs\_query](https://github.com/Jasperav/cdrs_query)
2. There is a tests folder with [progress.rs](https://progress.rs), if you try to run the tests on a windows, it fails (for me on 2 computers).

What's more, strange, I get 2 different compile errors on my 2 different windows computers. All computers are running rustc 1.43.1.

One compilation error is:

    error: could not compile `cdrs_query`.
    
    Caused by:
      process didn't exit successfully: `rustc --crate-name cdrs_query --edition=2018 src\lib.rs --error-format=json --json=diagnostic-rendered-ansi --crate-type lib --emit=dep-info,metadata,link -C debuginfo=2 -C metadata=373c5741dda4140a -C extra-filename=-373c5741dda4140a --out-dir C:\Users\Jasper\CLionProjects\cdrs_query\target\debug\deps -C incremental=C:\Users\Jasper\CLionProjects\cdrs_query\target\debug\incremental -L dependency=C:\Users\Jasper\CLionProjects\cdrs_query\target\debug\deps --extern cdrs_query_proc=C:\Users\Jasper\CLionProjects\cdrs_query\target\debug\deps\cdrs_query_proc-a9f994fa46a9363e.dll --extern proc_macro_hack=C:\Users\Jasper\CLionProjects\cdrs_query\target\debug\deps\proc_macro_hack-5f2bc32e7188aa50.dll` (exit code: 0xc0000005, STATUS_ACCESS_VIOLATION)

The other:

    error: linking with `C:\msys64\mingw64\bin\gcc.exe` failed: exit code: 1
      |
      = note: "C:\\msys64\\mingw64\\bin\\gcc.exe" "-fno-use-linker-plugin" "-Wl,--nxcompat" "-nostdlib" "-m64" "C:\\msys64\\mingw64\\x86_64-w64-mingw32\\lib\\dllcrt2.o" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\rsbegin.o" "-L" "C:\\msys64\\mingw64\\x86_64-w64-mingw32\\lib" "-L" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.10yym40hrkk60fld.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.11pecpngb6c09tbm.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.11wtqrxgr3s4ov3p.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.12ghoy67h40koqw0.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.12ovmk5s9gpmdqxh.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.17xppg7s4qx2oay3.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.19h3c893yr9lq4lg.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1aioe5s3mjfmcijk.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1bvys9b5laompazb.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1cfabbs38usal6uc.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1eupty4xsqt65l69.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1fa8z79xhzg59tjy.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1gcvg4b0jcfo2g3r.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1gr6jpyzh018liof.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1h4h9o8u19tb8taq.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1j0sthr01m0t0p66.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1loeia1abl1ncrrb.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1rrobblqlwm0eswt.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1srmcgeysj8mqz7f.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.1sxyodf0dh6s8fr.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.216ot1a90el4vsqt.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.259z9r0doa6p13ja.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.26f7j8thkx6ca848.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.2au8k5tywv0pwhvl.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.2dyhxakpd7pe4r2i.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.2k3navlzk7nu0var.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.2n4c240qn7dhreq.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.2nocuvuqp8811d57.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.32nuk6guu7trpvu8.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3a8uey96f8czo3y.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3axxuxtmce3cwsb.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3b02k6pfg5ybrg0m.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3fqdbqhg4iwet93x.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3gztn0rgv9ojtu40.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3hg2y89bxf918apt.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3hxuwzt3g2kbxtso.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3isp4zyge98gdj3o.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3itlasudtchyju7b.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3kzwkpa03h83f3tx.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3mpkjmqx5ho0b6bq.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3mwransn8gkx4y42.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3nb6mbyfcnpotx9w.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3s9v6dyll0whqnmb.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3tc37i6jpchn36v7.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3u0y853cauaej683.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3ugnuc6rt2c1ra1l.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3v430r1bdg7lbits.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3vs5vjxh9revqcs0.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3y7625ysoxn6or51.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.3z0al9k7hdf4ciex.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.43c0gaasvhmmfatf.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.49vdy5lkme4nlyih.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.49xoo1quo54uk4p0.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.4adk1z64itthxuse.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.4bjdsfdvztj69hd2.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.4ema0tqt2gkazfjk.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.4gjozm2yizyqfvku.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.4h2lcampdiba4qmw.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.4pr0eteglnm9vcdr.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.4q1qv64aimf2zv74.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.51kohyqi2tptkszg.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.5bb9rqdpnmvljz4m.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.5cwhllrr8u9xf2t0.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.7xnb90q7l8lws4r.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.bj8quq3ibzji58u.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.dxb0oyenmb1e1pn.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.fn94duskpew6cjr.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.jlo2jrtx6aklwyg.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.jycumisnj13furb.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.kqpfjdekq5zx99x.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.lfqugv1ill1nauf.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.mvi9ekd7hsf8vk7.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.o23rvngxdtwt3dw.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.rz2z6cbsegga4xj.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.th35iixpt96qpqz.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.tukmv0c7u59fuhg.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.vug9fo78hwinszg.rcgu.o" "-o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.dll" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.13jybnuwtxuk0gmy.rcgu.o" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.4dizn1lqx1iitrp3.rcgu.o" "-Wl,--gc-sections" "-nodefaultlibs" "-L" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps" "-L" "C:\\Users\\jasper\\.cargo\\registry\\src\\github.com-1ecc6299db9ec823\\winapi-x86_64-pc-windows-gnu-0.4.0\\lib" "-L" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib" "-Wl,-Bstatic" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libcdrs_con-943f2ff06774570b.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libcdrs_transform-32ce559fa8106ecb.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libcdrs-685c6f4349ff6c79.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libcdrs-9837fc72e2bfc612.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libuuid-1464d32ce6c7549f.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libtime-cbd53b6c833a080c.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\librand-016a7d6fdd7f93dd.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libr2d2-bf187d7614b95353.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libscheduled_thread_pool-b463adf965a4ec31.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libparking_lot-3d5c4f226384c651.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libparking_lot_core-719b7eb7e567a0ce.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libwinapi-d58e67ccfbdca0e1.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libsmallvec-ce915d4ae1fc9fcf.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\liblock_api-e4268f5db5cc955d.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libscopeguard-1e3937f3df9e247a.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\liblz4_compress-b3fb6b82de20213f.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libbyteorder-1889e7f61c7b8e10.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\liblog-39a8b9c3c6f79d99.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libcfg_if-5bc5725eeb91809d.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libsnap-32235058df2da79b.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\liblazy_static-e4db2369f50ab5ed.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libbyteorder-5973734557b5ae08.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libsyn-56a3e9b5c1d2d051.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libquote-dbf6c0275016ed4e.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libproc_macro2-a8292559a31f508b.rlib" "C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\libunicode_xid-6c634753621b18e1.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libproc_macro-2b7fef7b67ce54d0.rlib" "-Wl,--start-group" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libstd-b856675eeb99d410.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libpanic_unwind-7d9318240b1ba43e.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libhashbrown-1ea2d9c4ec771682.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\librustc_std_workspace_alloc-d40a2a1378721a73.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libbacktrace-e7c77cc61316712f.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libbacktrace_sys-160c2ebb98665c54.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\librustc_demangle-bd6bce6239a05021.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libunwind-c933f93cd63b2ad0.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libcfg_if-cabd342650cb5eba.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\liblibc-14fea7a91cbd1d65.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\liballoc-13f8aa678796468f.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\librustc_std_workspace_core-d5b2602ca33debca.rlib" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libcore-d0a6bfc5adb4da72.rlib" "-Wl,--end-group" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\libcompiler_builtins-9bf35bb038f7db43.rlib" "-Wl,-Bdynamic" "-lwinapi_advapi32" "-lwinapi_cfgmgr32" "-lwinapi_credui" "-lwinapi_gdi32" "-lwinapi_kernel32" "-lwinapi_msimg32" "-lwinapi_opengl32" "-lwinapi_secur32" "-lwinapi_user32" "-lwinapi_winspool" "-ladvapi32" "-lws2_32" "-luserenv" "-Wl,-Bstatic" "-lgcc_eh" "-lpthread" "-shared" "-Wl,--out-implib,C:\\Users\\jasper\\CLionProjects\\cdrs_query\\target\\debug\\deps\\cdrs_query_proc-a9f994fa46a9363e.dll.lib" "-Wl,-Bdynamic" "-lmingwex" "-lmingw32" "-lgcc" "-lmsvcrt" "-lmsvcrt" "-luser32" "-lkernel32" "C:\\Users\\jasper\\.rustup\\toolchains\\stable-x86_64-pc-windows-gnu\\lib\\rustlib\\x86_64-pc-windows-gnu\\lib\\rsend.o"
      = note: C:/msys64/mingw64/bin/../lib/gcc/x86_64-w64-mingw32/8.3.0/../../../../x86_64-w64-mingw32/bin/ld.exe: Error: export ordinal too large: 72382
              collect2.exe: error: ld returned 1 exit status
              
    
    error: aborting due to previous error

&amp;#x200B;
## [10][Reassign a variable to a value calculated from it inside FnMut](https://www.reddit.com/r/rust/comments/glbelm/reassign_a_variable_to_a_value_calculated_from_it/)
- url: https://www.reddit.com/r/rust/comments/glbelm/reassign_a_variable_to_a_value_calculated_from_it/
---
So I've got something like

    let mut v = something();
    otherthing.onevent(|| {
        v = v.update();
    });

Where onevent takes an FnMut and v.update() moves v.

This gives `cannot move out of v a captured variable in an FnMut closure` which makes sense, but how could I achieve the same effect? The best I've come up with is:

    let mut v = something();
    otherthing.onevent(|| {
        temp = std::mem::replace(v, something());
        temp.update();
        v = temp;
    });

but then I'm allocating a new value everytime which seems wasteful.

Thanks
## [11][get_unchecked makes things slower?](https://www.reddit.com/r/rust/comments/gl5scx/get_unchecked_makes_things_slower/)
- url: https://www.reddit.com/r/rust/comments/gl5scx/get_unchecked_makes_things_slower/
---
I was trying to determine the impact of bounds checking on my code, so I changed all the Vec[] calls to get_unchecked. Much to my surprise, using get_unchecked made it 14% slower on my benchmark. Is this the kind of thing that I should file as a bug? I've been trying to narrow it down to a minimal test case, but it is proving difficult, since the performance is a bit finicky (as is to be expected if it is the result of quirks in the LLVM optimizer as seems likely).

When Googling this, I found a report of a similar issue from 2018 on Github, but that one was marked as fixed, the result of an optimization gone wrong in older versions of LLVM

Update: I still haven't managed to cut it down to a nice repro case. Any attempt to cut out pieces of the code maddeningly seems to make the difference go away. I guess it's one of those really obnoxious "spooky action at a distance" optimizer problems. And yes, everything was timed in `--release` mode.
## [12][SIMD accelerated rank/select data structure](https://www.reddit.com/r/rust/comments/gl6ubm/simd_accelerated_rankselect_data_structure/)
- url: https://github.com/sujayakar/rsdict
---

