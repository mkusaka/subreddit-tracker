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
## [3][Popol: Minimal non-blocking I/O with Rust](https://www.reddit.com/r/rust/comments/htyyf3/popol_minimal_nonblocking_io_with_rust/)
- url: https://cloudhead.io/popol/
---

## [4][Clear explanation of Rust’s module system](https://www.reddit.com/r/rust/comments/htzkq7/clear_explanation_of_rusts_module_system/)
- url: http://www.sheshbabu.com/posts/rust-module-system/
---

## [5][Penrose - a tiling window manager in the style of xmonad](https://www.reddit.com/r/rust/comments/htny8s/penrose_a_tiling_window_manager_in_the_style_of/)
- url: https://youtu.be/buBzyw2Vivs
---

## [6][Booting to 'Hello Rust!' on x86_64](https://www.reddit.com/r/rust/comments/htheec/booting_to_hello_rust_on_x86_64/)
- url: https://micouy.github.io/posts/low-level-pt-1/
---

## [7][Catching panics from smol tasks](https://www.reddit.com/r/rust/comments/htz5p8/catching_panics_from_smol_tasks/)
- url: https://www.reddit.com/r/rust/comments/htz5p8/catching_panics_from_smol_tasks/
---
I  created this gist with  one of my trials.

[https://gist.github.com/arkaitzj/7782ab9f3e3b6ffac3fa821c70fb3ce2](https://gist.github.com/arkaitzj/7782ab9f3e3b6ffac3fa821c70fb3ce2)

Essentially I have an  already created smol::Task whose futures can  panic and I am trying to find a way to catch that panic to do some reporting.  

I am not interested in recovering that particular task but I'd like to relaunch a new instance of it without crashing the process.

I want to target that task, not a general panic catch.

Is this possible?
## [8][Tree-Buf v0.10.0 released!](https://www.reddit.com/r/rust/comments/htn0b7/treebuf_v0100_released/)
- url: https://www.reddit.com/r/rust/comments/htn0b7/treebuf_v0100_released/
---
Since the announcement of Tree-Buf about [and a half month ago](https://www.reddit.com/r/rust/comments/gud6c1/introducing_treebuf/), I've been focused on speed, writing smaller files, and reliability. The MessagePack vs Tree-Buf for GraphQL [benchmark](https://github.com/That3Percent/tree-buf#tree-buf-vs-messagepack-for-graphql) now clocks in at reading and writing **2x as fast** as MessagePack, producing a file that is only **1/17th of the size** making Tree-Buf by far the fastest and smallest self-describing serialization format available.

Here's a summary of the important changes since the announcement:

* Renamed Read/Write to Encode/Decode to not conflict with std traits
* Support reading and writing for more fixed array length types in macros to support Ethereum data types
* Created the [firestorm](https://github.com/That3Percent/firestorm) crate to replace inferno for intrusive profiling to get better performance and ergonomics (a separate announcement will be made for this)
* Implement the fast\_size\_for method for all encoders to support faster measurements when deciding which encoders to use on samples of the data
* Avoid unnecessarily buffering data in special situations
* Fixed performance problem when writing that run-length encodings were being evaluated deeply nested
* Added delta zigzag support for u32 (more types need to be added)
* Added an unchecked writer to [simple-16](https://github.com/That3Percent/simple-16) and use it within Tree-Buf for much faster writing
* Fixed panic in the Gorilla float compressor when running in Debug

For more information check out the [README](https://github.com/That3Percent/Tree-Buf) but please note that Tree-Buf is not yet production-ready and the format is still changing. You might find basic usability issues - like i32 not yet being supported. At 177 commits and counting, this turned out to be more work than I expected. Patience, things like i32, huge file support, SIMD, and more may come in time.
## [9][I have a strange problem like 'PC offset is too large' in rust linking, what could be the reason to it?](https://www.reddit.com/r/rust/comments/htu8zm/i_have_a_strange_problem_like_pc_offset_is_too/)
- url: https://www.reddit.com/r/rust/comments/htu8zm/i_have_a_strange_problem_like_pc_offset_is_too/
---
To illustrate this problem, I have minimized the project: [https://github.com/luojia65/runtime-test](https://github.com/luojia65/runtime-test). This project depends on `riscv-rt` crate 0.8.0.

Clone the project down and try to compile. After compile using `cargo build`, it shows problems like this:

```text
"rust-lld" "-flavor" "gnu" "--eh-frame-hdr" "-L" /* ... many rlib files ... */ "-Tmemory.x" "-Tlink.x" "-Bdynamic"
  = note: rust-lld: error: .. /runtime-test/target/riscv64imac-unknown-none-elf/debug/deps/libriscv_rt-7ea7a4b3a0900dbe.rlib(riscv-rt.o):(.eh_frame): PC offset is too large: 0x80000018
```

I've uploaded full error message into ubuntu pastebin: https://paste.ubuntu.com/p/jdmhvmy3WS/. It shows that the linking process has some error like `(.eh_frame): PC offset is too large: 0x80000018`.

It's very strange. I configure the memory settings in `memory.x`. When SRAM base is `0x80000000`, the error occurres. But when I set it to a slightly bigger value like `0x80000134`, or sligntly smaller `0x7fffffe0`, it compiles normally without any errors.

My target (embedded) platform has memory located exactly in `0x80000000` where I must flash my program into. 

So what could have happened? How could `PC offset is too large` be generated? Does this error happen elsewhere? Thank you all redditors! :)

My Rust version:

```
❯ rustc --version --verbose
rustc 1.47.0-nightly (39d5a61f2 2020-07-17)
binary: rustc
commit-hash: 39d5a61f2e4e237123837f5162cc275c2fd7e625
commit-date: 2020-07-17
host: x86_64-apple-darwin
release: 1.47.0-nightly
LLVM version: 10.0
```

Edit: tried on latest stable `rustc 1.45.0 (5c1f21c3b 2020-07-13)`, the error does not occur again.
## [10]["Sure, you can build a web application with Rust but..."](https://www.reddit.com/r/rust/comments/htwluy/sure_you_can_build_a_web_application_with_rust_but/)
- url: https://www.reddit.com/r/rust/comments/htwluy/sure_you_can_build_a_web_application_with_rust_but/
---
As a newbie web developer I have very recently started learning Rust exclusively for building web applications. I know that it is risky and not that mature but I have the time and energy and I want to do something new and exciting instead of learning old stuff.

So please tell me why I can build Rust web applications but...

What are your buts? What are the things that I am going to miss? Are they really a deal breaker?

By the way, I already know about [Are we web yet?](https://www.arewewebyet.org/) I am just curious about everyday practical side of things.
## [11][C programming's -O3 equivalent compiler flag in rust](https://www.reddit.com/r/rust/comments/htyvkw/c_programmings_o3_equivalent_compiler_flag_in_rust/)
- url: https://www.reddit.com/r/rust/comments/htyvkw/c_programmings_o3_equivalent_compiler_flag_in_rust/
---
Rustaceans,

I was wondering what is the equivalent flag for `-O3` in rust? I searched compiler flags for optimization [here](https://www.mankier.com/1/rustc#-O) however I could only find `-O2` equivalent (`-O` in rust). Could someone enlighten?
## [12][Referencing self in a function without the &amp;self argument?](https://www.reddit.com/r/rust/comments/htpork/referencing_self_in_a_function_without_the_self/)
- url: https://www.reddit.com/r/rust/comments/htpork/referencing_self_in_a_function_without_the_self/
---
I'm trying to create a new() function for a struct, but I am struggling to find a way to have a member of my function point to another member of my function. Here's my code:

    struct test&lt;'b&gt;
    {
    	x: i32,
    	y: &amp;'b i32,
    }
    
    impl&lt;'b&gt; test&lt;'b&gt;
    {
    	pub fn new(arg1: i32) -&gt; Self
    	{
    		test
    		{
    			x: arg1,
    			y: &amp;x
    		}
    	}
    }

I understand that if I change my function to

    impl&lt;'b&gt; test&lt;'b&gt;
    {
    	pub fn new(&amp;'b self, arg1: i32) -&gt; Self
    	{
    		test
    		{
    			x: arg1,
    			y: &amp;self.x
    		}
    	}
    }

Then this would work! But since I'm trying to make a new object, it doesn't really make sense to require an already existing object. I guess this is where I need to dynamically allocate memory? Or is there a way that I can just have my pointer point to x in the constructor?
