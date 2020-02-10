# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (6/2020)!](https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/)
- url: https://www.reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here_62020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.41]](https://www.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/ecxd62/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting (adapted from /r/cpp's jobs thread).

#### Rules for individuals:

* Don't create top-level comments; those are for employers.
* Feel free to reply to top-level comments with on-topic questions.
* I will create a stickied top-level comment for individuals looking for work.
* I will create an additional top-level comment for meta discussion.

#### Rules for employers:

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
## [3][Lib.rs - crates.io on steroids](https://www.reddit.com/r/rust/comments/f1mr4h/librs_cratesio_on_steroids/)
- url: https://lib.rs/
---

## [4][Announcing Seed v0.6](https://www.reddit.com/r/rust/comments/f1nwkr/announcing_seed_v06/)
- url: https://www.reddit.com/r/rust/comments/f1nwkr/announcing_seed_v06/
---
Hi! I'm happy to announce Seed v0.6 release.

 Seed is a Rust framework for creating fast and reliable web apps.

Seed uses an unconventional approach to describe how to display DOM elements. It uses neither completely natural (i.e. macro-free) Rust code, nor an HTML-like abstraction (e.g. JSX or templates). Our intent is to make the code close to natural Rust, while streamlining the syntax in a way suited for creating a visual layout with minimal repetition. The macros used are thin wrappers for constructors, and don't conceal much.

This release focuses mostly on internal refactoring. See [https://seed-rs.org/guide/changelog](https://seed-rs.org/guide/changelog)

We have plenty of examples as well as several starter kits and a welcoming discord chat and forum!

Please visit [https://seed-rs.org/](https://seed-rs.org/) to learn more about building web projects in Rust &amp; WASM!

Cheers,  
Seed core members
## [5][We are now more than 90k Rustaceans on this subreddit](https://www.reddit.com/r/rust/comments/f1nglg/we_are_now_more_than_90k_rustaceans_on_this/)
- url: https://subredditstats.com/r/rust
---

## [6][Pointer-utils: a collection of crates, providing simple custom DSTs, pointer unions, borrowed reference counts, and more!](https://www.reddit.com/r/rust/comments/f1iixm/pointerutils_a_collection_of_crates_providing/)
- url: https://github.com/CAD97/pointer-utils/blob/master/blog/Announcement.md
---

## [7][win32job: a crate for using Windows' Job Objects](https://www.reddit.com/r/rust/comments/f1ob3a/win32job_a_crate_for_using_windows_job_objects/)
- url: https://www.reddit.com/r/rust/comments/f1ob3a/win32job_a_crate_for_using_windows_job_objects/
---
I created the `win32job` [crate](https://github.com/ohadravid/win32job-rs), a small library which exposes a safe API for Windows' Job Objects.

The TL;DR is that this allows a process to force all sort of limits on itself and processes it creates, like limiting the amount of memory used, or forcing all "child" processes to terminate when the "parent" terminates (this relationship is different on Windows compared to Linux/macOS).

This example show how to force any created sub processes to exit when the main process exits:

    use win32job::Job;
    use std::process::Command;
    
    fn main() -&gt; Result&lt;(), Box&lt;dyn std::error::Error&gt;&gt;  {
        let job = Job::create()?;
        
        let mut info = job.query_extended_limit_info()?;
    
        info.limit_kill_on_job_close();
    
        job.set_extended_limit_info(&amp;mut info)?;
        
        job.assign_current_process()?;
    
        Command::new("cmd.exe")
                .arg("/C")
                .arg("ping -n 9999 127.0.0.1")
                .spawn()?;
    
        // The cmd will be killed once we exit, or `job` is dropped.
        
        Ok(())
    }

The API tries to stay close to the original Windows API (hence the long-ish function names).

There a number of existing places using this functionality, all with their own custom impls, so I think this is something worthy of a dedicated crate:

[rustjob](https://github.com/alexcrichton/rustjob/blob/master/src/main.rs#L139) (Wrapper script for buildbots)

[rust-gha](https://github.com/rust-lang/rust-gha/blob/master/src/bootstrap/job.rs) (GitHub Actions tests for rust-lang/rust)

[rustup](https://github.com/rust-lang/rustup/blob/master/src/cli/job.rs) and [cargo](https://github.com/rust-lang/cargo/blob/dc83ead224d8622f748f507574e1448a28d8dcc7/src/cargo/util/job.rs) (You know what this is!)

&amp;#x200B;

&amp;#x200B;

As a little side note, I wanted a better way to test the `README`, and while `skeptic` and `little-skeptic` exists they require a bunch of deps, so I found a [little trick](https://github.com/ohadravid/win32job-rs/blob/aa627a24150681cb712315c7f8663e34fcd3f7f7/src/lib.rs#L87) somewhere in the GitHub comments for various Rust issues and ended up with this:

    #[cfg(debug_assertions)]
    mod test_readme {
        macro_rules! calculated_doc {
            ($doc:expr, $id:ident) =&gt; {
                #[doc = $doc]
                enum $id {}
            }
        }
    
        calculated_doc!(include_str!("../README.md"), _DoctestReadme);
    }

Which will run the `README` as part of the doctests!
## [8][Rust Programming for IOT (Internet of Things)](https://www.reddit.com/r/rust/comments/f1mnf0/rust_programming_for_iot_internet_of_things/)
- url: https://www.youtube.com/watch?v=ro-G5gu5fMk&amp;feature=share
---

## [9][A Rustacean change of heart](https://www.reddit.com/r/rust/comments/f1c3v6/a_rustacean_change_of_heart/)
- url: https://www.reddit.com/r/rust/comments/f1c3v6/a_rustacean_change_of_heart/
---
 

Hi all,

I decided about 2 weeks ago to put my efforts into being a Rustacean after spending six months learning Haskell. I came from 4 years of professional Scala development. Things were going great with Haskell however the tooling(IDE in particular) is generally not up to task with what I would like to ultimately achieve .

This year I will be building a few information system MVP products using Rust , in my spare time after work, and  I see Rust as a language that can help me towards that end. Of course many other languages are also up to the task, but I purposefully wanted a non mainstream language so that I can punt it as a recruitment tool to attract people who are more inclined towards learning. It just so happened that Rust comes with all these other goodies that satisfy my FP inclinations and the memory safety features, which I am eager to start sparring with

So what advice would anyone give towards learning Rust, what frameworks can you guys recommend for basic CRUD operations. And finally what can I do or help out with in the community.
## [10][Partially derive default?](https://www.reddit.com/r/rust/comments/f1nydq/partially_derive_default/)
- url: https://www.reddit.com/r/rust/comments/f1nydq/partially_derive_default/
---
Suppose I have a struct, most of its fields except one can derive default:

    struct MyStruct {
        body: Body, // Body does not have default
        a: u32,
        b: Option&lt;Foo&gt;,
        c: Vec&lt;Bar&gt;,
        // ...
    }

Can I easily derive partial default, so I can easily write something like the following:

    impl MyStruct {
        fn new(body: Body) -&gt; Self {
            // Error, because Self does not implement Default
            Self { body, ..Self::default() }
        }
    }
## [11][Using use keyword inside a function](https://www.reddit.com/r/rust/comments/f1lhp1/using_use_keyword_inside_a_function/)
- url: https://www.reddit.com/r/rust/comments/f1lhp1/using_use_keyword_inside_a_function/
---
Hey there, I'm very new to rust and recently I've been diving into the deep end trying to learn how to use audio API's, my current choice is CPAL and on the [documentation here](https://docs.rs/cpal/0.11.0/cpal/) a number of times the code shows a totally unexplained use of the "use" keyword which I initially thought was just for importing libraries and modules. What is the purpose of writing these "use"s right before these other lines? Do I need to be writing them in the order they appear? Or can I just combine them all at the top of my code?
## [12][Pushrod 0.2.27: SDL2-based GUI for Rust - Call for Help!](https://www.reddit.com/r/rust/comments/f1fcya/pushrod_0227_sdl2based_gui_for_rust_call_for_help/)
- url: https://www.reddit.com/r/rust/comments/f1fcya/pushrod_0227_sdl2based_gui_for_rust_call_for_help/
---
Hi all!

I've been working on [rust-pushrod](https://www.github.com/KenSuenobu/rust-pushrod/) for a little over a year now, and I'm pretty happy with the progress; but there's only so much progress one person can do on the project.

So far, I've had a lot of interest in the project, but would like to see if anyone else is interested in helping out with the project and some of the issues that remain before a 0.5.x release can be achieved.

I have big plans for the 0.5.x release - one of which is to start designing a GUI so that applications can be written easier with a WYSIWYG style of drag-drop designer application.

The latest release of Pushrod includes screenshots for each `Widget` that is currently in the system, a `Tile` widget that can be used as a tile in a group box, and a `TabBarWidget` that displays simple tabs that can be used to navigate pages in an application.

If anyone is interested in helping out - either [monitarily](https://patreon.com/KenSuenobu) or through the [issues navigator](https://www.github.com/KenSuenobu/rust-pushrod/issues/), I would be forever grateful.  I am hoping for adoption of this library, and I have a passion for writing good GUI APIs.

Please feel free to leave feedback - good or bad.  You can [also visit my blog here](https://kensuenobu.github.io/)

Thank you, fellow Rustaceans!
