# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (2/2020)!](https://www.reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here_22020/)
- url: https://www.reddit.com/r/rust/comments/ekpqp7/hey_rustaceans_got_an_easy_question_ask_here_22020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/ehk18j/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.40]](https://www.reddit.com/r/rust/comments/ecxd62/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/ecxd62/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/dvxw6u/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust every week, for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).

#### Rules for individuals:

* Don't create top-level comments; those are for employers.
* Feel free to reply to top-level comments with on-topic questions.
* I will create a stickied top-level comment for individuals looking for work.
* I will create an additional top-level comment for meta discussion.

#### Rules for employers:

* To find individuals seeking work, see the replies to the stickied top-level comment; you will need to click the "more comments" link below that comment in order to see them.
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
## [3][Cranelift backend for Rust](https://www.reddit.com/r/rust/comments/enxgwh/cranelift_backend_for_rust/)
- url: https://www.reddit.com/r/rust/comments/enxgwh/cranelift_backend_for_rust/
---
For any that don't know, there's an ongoing work to create a [backend](https://github.com/bjorn3/rustc_codegen_cranelift) for Rust built on cranelift.  This comes with a promise of [drastically](https://www.reddit.com/r/rust/comments/9xobgz/the_cranelift_backend_for_rustc_now_produces_code/) reduced debug compilation times.

A fantastic amount of progress has been made, with all but 57 tests from the rustc test suite passing.

&amp;#x200B;

Almost all of this work has been done single-handedly by [bjorn3](https://github.com/bjorn3).  I'm here hoping to raise awareness of the project and ask any who are capable and willing to support by [donating to bjorn3](https://liberapay.com/bjorn3/) or contributing via [github](https://github.com/bjorn3/rustc_codegen_cranelift).  Thank you!!

&amp;#x200B;

I firmly believe that getting compilation times down is vital for Rust's continued success - so I'm very excited to see this project completed!
## [4][Update on rustysd: a minimal systemd compatible service manager written in rust](https://www.reddit.com/r/rust/comments/eo276v/update_on_rustysd_a_minimal_systemd_compatible/)
- url: https://www.reddit.com/r/rust/comments/eo276v/update_on_rustysd_a_minimal_systemd_compatible/
---
Hello r/rust!

I have posted here a few weeks back about my newest project: rustysd. Since then I have pushed the project further and have reached a point where I think it's worthwhile to share it again. [The repo is located here](https://github.com/KillingSpark/rustysd/tree/v0.1.0).

The short summary of rustysd goes like this:
"Rustysd is a service manager that tries to replicate systemd behaviour for a subset of the configuration possibilities. It focuses on the core functionality of a service manager."
(So it is only minimal in the amount of features it tries to provide. The binaries can be somewhat small if built with musl and strip'ed, but generally rust produces relatively big binaries)

This means rustysd can read (a subset of) systemd unit files and run them as if it were systemd, as far as the services will know. Since rustysd does not need to be PID1 it can provide this functionality to linux distros that don't use systemd and also to FreeBSD (other BSDs are untested right now). Additionally it can be used in docker as PID1 so you can run services that require to be run by systemd in a container.

The biggest usability issue right now is probably the rsdctl tool which is the corresponding tool for systemds systemctl. I don't have a lot of experience in making good UIs so if you want to help me out with that I would be very thankful.

The next steps will probably be:
* make rsdctl a good tool
* get rustysd to act as the init system in a VM
* develop a test-suite to be able to catch regressions / bugs

I would be glad about any feedback, especially regarding the unsafe blocks I had to use where I needed to use libc::* directly. (This is contained in the platform module if you are interested)

If you have any use for this I'd like to hear from you! I have started this project for the fun of it but if anyone actually wants this I would love to make it as easy as possible for you!
## [5][PROST! (a Protocol Buffers implementation) v0.6.0 released](https://www.reddit.com/r/rust/comments/enwvap/prost_a_protocol_buffers_implementation_v060/)
- url: https://github.com/danburkert/prost/releases/tag/v0.6.0
---

## [6][Shamir's Secret Sharing written in Rust](https://www.reddit.com/r/rust/comments/eo36fp/shamirs_secret_sharing_written_in_rust/)
- url: https://www.reddit.com/r/rust/comments/eo36fp/shamirs_secret_sharing_written_in_rust/
---
Hi all, I've just published my first Rust crate so I would love to have other rustaceans opinions: [https://crates.io/crates/sharks](https://crates.io/crates/sharks)

Basically it is a crate which implements Shamir's Secret Sharing ([https://en.wikipedia.org/wiki/Shamir's\_Secret\_Sharing](https://en.wikipedia.org/wiki/Shamir's_Secret_Sharing)), an algorithm I find very very interesting.

I wanted it to be small, fast and secure, so I made it as idiomatic as possible and keeping external dependencies to a minimum. It only uses rand and num, so it is super easy to read! :)

Thank you very much for checking it out!
## [7][Awesome wgpu-rs. List of learning resources, games, etc.](https://www.reddit.com/r/rust/comments/eo4pj8/awesome_wgpurs_list_of_learning_resources_games/)
- url: https://github.com/rofrol/awesome-wgpu-rs
---

## [8][An update on Rust Database Connectivity (RDBC) and news about the move to Tokio](https://www.reddit.com/r/rust/comments/enphc2/an_update_on_rust_database_connectivity_rdbc_and/)
- url: https://andygrove.io/2020/01/rust-database-connectivity-rdbc/
---

## [9][git-poly : A minimalist multi repo tool written in rust](https://www.reddit.com/r/rust/comments/eo1j54/gitpoly_a_minimalist_multi_repo_tool_written_in/)
- url: https://www.reddit.com/r/rust/comments/eo1j54/gitpoly_a_minimalist_multi_repo_tool_written_in/
---
Hey,  
 I wanted to share a small project I've been working on that helps with working on multiple git repos at once.

[https://luke\_titley.gitlab.io/git-poly](https://luke_titley.gitlab.io/git-poly)

I wrote it for myself mostly, to help on a project at my work.

It's in the same sort of space as these projects:

* git repo
* git subtree
* git slave
* git submodule

It's written in rust.
## [10][Show /r/rust: Optimath, linear algebra with const generics and SIMD. [no_std]](https://www.reddit.com/r/rust/comments/eo4ury/show_rrust_optimath_linear_algebra_with_const/)
- url: https://djugei.github.io/optimath-0-3-0/
---

## [11][How to export C++ struct into Rust?](https://www.reddit.com/r/rust/comments/eo43dg/how_to_export_c_struct_into_rust/)
- url: https://www.reddit.com/r/rust/comments/eo43dg/how_to_export_c_struct_into_rust/
---
I know `extern "C"` can extern functions from c++ and after building and linking this code, I can use it in Rust via libc crate and extern:

    extern crate libc;
    
    extern {
        fn mai() -&gt; libc::c_int;
    }
    
    pub fn tar() {
        unsafe {
            mai()
        };
    }

But imagine I have a C++ struct:

    struct Order {
        // Some fields and code
    }

And I need to use this struct directly in Rust. How to do that?
## [12][Has there been any community project started to revive CEF support in Servo?](https://www.reddit.com/r/rust/comments/eo2mwu/has_there_been_any_community_project_started_to/)
- url: https://www.reddit.com/r/rust/comments/eo2mwu/has_there_been_any_community_project_started_to/
---
For reference CEF (Chromium Embedded Framework) is what's used to embed a chromium based browser in applications. It's what's used in Steam and Spotify for example.
Some people in the Rust team were working on CEF support basically making it so Servo (Mozillas web-rendering engine written in Rust) could easily be used instead. 
However about 2 years ago they dropped it to focus on other stuff hoping it would be a 3rd party project.

https://groups.google.com/forum/#!topic/mozilla.dev.servo/5MHDnfMya3s

Having looked into it recently as I'm leaning Rust and was looking into GUI options for potential projects I've found nothing that indicates it has been revived sadly.
I'm mainly interested in this because as of now GUI development frameworks for Rust are seemingly in their infancy.
Additionally Servo seems to be coming along well but I'd love it if there was a way of bringing it's speed to the nowadays oh so common web technology based desktop applications where chromium dominates trough both CEF and Electron. I think it would generate a lot of interest in Rust as a language.
