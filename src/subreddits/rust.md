# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (12/2020)!](https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.42]](https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://en.reddit.com/r/rust/comments/eyw94s/official_rrust_whos_hiring_thread_for_jobseekers/).

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
## [3][`template-rust-backend-with-electron-frontend` released](https://www.reddit.com/r/rust/comments/flsoa0/templaterustbackendwithelectronfrontend_released/)
- url: https://www.reddit.com/r/rust/comments/flsoa0/templaterustbackendwithelectronfrontend_released/
---
Hi Rust users, I published the template project for a Rust backend with Electron frontend project. I believe this template is useful for now maybe until a GUI Toolkit libraries will be grows well that has a high Rust purity and high licensing flexibility.

[https://github.com/usagi/template-rust-backend-with-electron-frontend](https://github.com/usagi/template-rust-backend-with-electron-frontend)
## [4][Alchemy: A New Rust GUI Framework](https://www.reddit.com/r/rust/comments/flsrq5/alchemy_a_new_rust_gui_framework/)
- url: https://alchemy.rs/
---

## [5][Cheatsheet: Option (in Rust) vs Maybe (in Haskell)](https://www.reddit.com/r/rust/comments/fli7wk/cheatsheet_option_in_rust_vs_maybe_in_haskell/)
- url: https://notes.iveselov.info/cheatsheet-rust-option-vs-haskell-maybe
---

## [6][How does Rust's scalability compare to Erlang?](https://www.reddit.com/r/rust/comments/flf2ld/how_does_rusts_scalability_compare_to_erlang/)
- url: https://www.reddit.com/r/rust/comments/flf2ld/how_does_rusts_scalability_compare_to_erlang/
---
The ErlangVM (BEAM) is well regarded as a battle tested, no downtime runtime that can scale gracefully to handling orders of magnitude more concurrent tasks without spiking.
What kind of mechanisms need to be in place to enable a similar amount of graceful scaling in Rust?

I have been keeping an eye on the actor frameworks (such as actix, bastion), but I haven't seen something that would allow you to schedule and prefer short running processes rather than long running processes as Erlang does. I believe that if a process handlers runs for a long time in Erlang, the scheduler will pause that process to allow shorter handlers to take priority. This leads to ensuring that there will seldom be anything blocking your program (unless it's a long running NIF).

Does anyone have thoughts about if these scalability characteristics are already enabled today with Rust, or what would have to be done to achieve them?

Thanks!
## [7][WASM built with Emscripten and loaded with js-sys?](https://www.reddit.com/r/rust/comments/flravn/wasm_built_with_emscripten_and_loaded_with_jssys/)
- url: https://www.reddit.com/r/rust/comments/flravn/wasm_built_with_emscripten_and_loaded_with_jssys/
---
I have been playing with [https://github.com/rustwasm/wasm-bindgen/tree/master/examples/wasm-in-wasm](https://github.com/rustwasm/wasm-bindgen/tree/master/examples/wasm-in-wasm) in order to see how I can load in a WASM module.  I have some C code that I would like to use so I built a WASM module from this code using Emscripten.  Unfortunately it panics when I run the code and I am now asking myself if this is possible at all.  Is this possible?
## [8][Trace caller’s source location in the stable Rust](https://www.reddit.com/r/rust/comments/flnns3/trace_callers_source_location_in_the_stable_rust/)
- url: https://www.reddit.com/r/rust/comments/flnns3/trace_callers_source_location_in_the_stable_rust/
---
[https://blog.knoldus.com/trace-callers-source-location-in-the-stable-rust/](https://blog.knoldus.com/trace-callers-source-location-in-the-stable-rust/)
## [9][This Week in Rust 330](https://www.reddit.com/r/rust/comments/flb5tn/this_week_in_rust_330/)
- url: https://this-week-in-rust.org/blog/2020/03/17/this-week-in-rust-330/
---

## [10][Announcing Rust IPFS, and a call for contributors](https://www.reddit.com/r/rust/comments/fl6ij9/announcing_rust_ipfs_and_a_call_for_contributors/)
- url: https://blog.ipfs.io/2020-03-18-announcing-rust-ipfs/
---

## [11][Recommendations on this data structure](https://www.reddit.com/r/rust/comments/fltzwg/recommendations_on_this_data_structure/)
- url: https://www.reddit.com/r/rust/comments/fltzwg/recommendations_on_this_data_structure/
---
Any recommendations on how to design the following data structure:

The biggest issue I have is using the RefCell to be able to mutate the structure inside the hashmap. Is there a better solution to this?
            
    pub struct AAA {
        pub(crate) one: Option&lt;u8&gt;,
    }
    
    pub struct BBB {
        pub(crate) one: Option&lt;u8&gt;,
        pub(crate) two: Option&lt;u8&gt;,
        pub(crate) three: Option&lt;u8&gt;,
        pub(crate) four: Option&lt;u8&gt;,
        pub(crate) five: Option&lt;u32&gt;,
        pub(crate) six: Option&lt;u32&gt;,
    }
    
    pub struct Union {
        pub(crate) b: Option&lt;AAA&gt;,
        pub(crate) a: Option&lt;BBB&gt;,
    }
    
    pub struct Hash {
        hash: HashMap&lt;u32, RefCell&lt;Union&gt;&gt;,
    }
## [12][As a programming language, how does Rust compare to Python?](https://www.reddit.com/r/rust/comments/fld46b/as_a_programming_language_how_does_rust_compare/)
- url: https://www.reddit.com/r/rust/comments/fld46b/as_a_programming_language_how_does_rust_compare/
---
Now I know that Rust is incredibly good at memory management and is also quite fast(?), and companies like Discord are using it in their backend systems. On the other hand, Python is the go to language for machine learning and other AI related stuff.

I know it seems like I answered my own question, but what I really want to know is that being a student, should I focus on Rust as well, or just continue with Python until the right time comes, when Rust becomes more popular than it is now?

 I don't know as much about Rust as I would like to, so please correct me if I'm wrong about anything.
