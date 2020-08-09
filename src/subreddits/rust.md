# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (31/2020)!](https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

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
## [3][Announcing Pyo3 Log](https://www.reddit.com/r/rust/comments/i6ef2x/announcing_pyo3_log/)
- url: https://vorner.github.io/2020/08/08/pyo3-log.html
---

## [4][Recommend Rust-friendly microcontroller for home automation](https://www.reddit.com/r/rust/comments/i6645g/recommend_rustfriendly_microcontroller_for_home/)
- url: https://www.reddit.com/r/rust/comments/i6645g/recommend_rustfriendly_microcontroller_for_home/
---
Hi everybody,

I’m looking for recommendations of Rust-friendly microcontroller that potentially can be used with a network stack (ZigBee or BLE). I guess the question is not about microcontroller itself, but more about for which ones there’s the best tooling in Rust.
I’ve decided to get some experience with Rust by getting back to some of my hobby home automation project.
Since I have to go wireless, I’ll need network stack.

Thanks!
## [5][nth_element implementation by Floyd and Rivest.](https://www.reddit.com/r/rust/comments/i6gaso/nth_element_implementation_by_floyd_and_rivest/)
- url: https://www.reddit.com/r/rust/comments/i6gaso/nth_element_implementation_by_floyd_and_rivest/
---
`nth_element` [implementation](https://crates.io/crates/floydrivest) using the Floyd-Rivest algorithm. This is my first crate, I found myself in need of the c++'s  counterpart while porting a bigger project. Here it is... critiques very welcome.  


I'm pretty new to Rust and willing to expand the crate to better suit rustaceans needs.
## [6][How to add a class variable to watch in Visual Studio code when using cppvsdbg?](https://www.reddit.com/r/rust/comments/i6f68o/how_to_add_a_class_variable_to_watch_in_visual/)
- url: https://www.reddit.com/r/rust/comments/i6f68o/how_to_add_a_class_variable_to_watch_in_visual/
---
Hi All,

I am using Visual Studio Code to develop my app and one thing I am missing is a good debugger since my code is huge. Often times I would like to add some variables to watch.

I am able to visualize a local variables in the watch window and in the debug console. However, If i do with struct fields using Arc, I get "Class Struct has no member Field". Any suggestion?

    pub struct mine {
       pub param: f32,
    }
    
    impl mine {
      // pub new() {}
      pub add_one(&amp;self) -&gt; f32 {self.param + 1}
    }
    
    let s1 = Arc::new(RwLock::new(mine::new());
    let s = s1.read();
    // s added to watch or typed in the debug console works
    // s.param added to watch or typed in the debug console
    // says class mine has no member param
    // storing the s.param to some local variable and adding that to watch or typed in debug console works

Likewise,

I am not sure how to call s.add\_one(). Can someone help?

Update: Found out I can get the fields by using C++ syntax.

    s.rwlock-&gt;data.value.param

Now, stuck with calling functions
## [7][Rayon for simple mapping](https://www.reddit.com/r/rust/comments/i6btd7/rayon_for_simple_mapping/)
- url: https://www.reddit.com/r/rust/comments/i6btd7/rayon_for_simple_mapping/
---
Very simple question. 

Should I use rayon for simple task such as restructuring the output of a db query? Or is there a cost to doing that to a simple task. 

Wondering if you guys just literally replace map with rayon’s.
## [8][Using Rust Library in Wasm](https://www.reddit.com/r/rust/comments/i6fddw/using_rust_library_in_wasm/)
- url: https://www.reddit.com/r/rust/comments/i6fddw/using_rust_library_in_wasm/
---
I was going through "Rust and WebAssembly" [book](http://rustwasm.github.io/docs/book/game-of-life/implementing.html), one of the end-chapter exercise requires generation of random numbers. Generating random number using js-sys  works fine but if I use Rust rand library the page does not render. Is there a limitation on Rust library that can be used to generate Wasm?
## [9][Is concurrency actually simultaneous?](https://www.reddit.com/r/rust/comments/i67yl0/is_concurrency_actually_simultaneous/)
- url: https://www.reddit.com/r/rust/comments/i67yl0/is_concurrency_actually_simultaneous/
---
I'm  currently writing a piece of code designed to find out some data for  every number within a range. The problem is for some numbers this is  incredibly intensive, and takes a long time, so I've been trawling the  concurrency documentation to try and see how I could re-implement my  code to use it.

The problem is  that it looks like the documentation is saying that they don't run at  the same time, but rather hand off a single operating system process. As  it says in the docs

&gt;The calls to  
&gt;  
&gt;thread::sleep   
&gt;  
&gt;force a thread to  stop its execution for a short duration, allowing a different thread to  run. The threads will probably take turns, but that isn’t guaranteed: it  depends on how your operating system schedules the threads.

This  seems to imply to me that it won't run the threads simultaneously on  different cores on my machine. Is this actually the case? If so, how can  one run separate blocks of code on different cores?
## [10][SPIR-Q 0.4.2 is now released!](https://www.reddit.com/r/rust/comments/i6hyop/spirq_042_is_now_released/)
- url: /r/rust_gamedev/comments/i6hxh6/spirq_042_is_now_released/
---

## [11][Ergonomic postgres joins?](https://www.reddit.com/r/rust/comments/i68lap/ergonomic_postgres_joins/)
- url: https://www.reddit.com/r/rust/comments/i68lap/ergonomic_postgres_joins/
---
Can anyone point to any examples of ergonomic joins with Diesel or SQLx? As far as I can tell with Diesel joins always produce a tuple of structs corresponding to the tables, so if you had Posts and User  for instance and wanted to do a join this would result in (Post, User) and would need to be converted into a new type to nest user within post and to serialize. Is there a way to make Diesel or SQLx work without needing to map the output type manually?
## [12][Tools for Debugging Memory/Space Leaks?](https://www.reddit.com/r/rust/comments/i6b26z/tools_for_debugging_memoryspace_leaks/)
- url: https://www.reddit.com/r/rust/comments/i6b26z/tools_for_debugging_memoryspace_leaks/
---
Hi all.

I have a small but complex program (heavy usage of futures/streams + pulls in many crates). The program seems to have a memory (or space) leak as it gets OOM after about 24 hours. I'm guessing some of my objects aren't getting dropped at the time I intend.

What are your recommendations for tools to solve this problem? A possible tool would be something that allows me to wrap some of the types I use into newtypes that log when the object dropped. And then the logs can be graphed.
