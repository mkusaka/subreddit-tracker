# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (10/2020)!](https://www.reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

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
## [3][I wrote a new rust/gtk app to collect info to enter work timesheets](https://www.reddit.com/r/rust/comments/ffqhhv/i_wrote_a_new_rustgtk_app_to_collect_info_to/)
- url: https://github.com/emmanueltouzery/cigale
---

## [4][Help me understand enums in Rust](https://www.reddit.com/r/rust/comments/fft7q5/help_me_understand_enums_in_rust/)
- url: https://www.reddit.com/r/rust/comments/fft7q5/help_me_understand_enums_in_rust/
---
I am going through all the rustlings exercises and today I found a weird one (coming from c++)...

    fn main() {
        let mut res = 42;
        let option = Some(12);
        for x in option {
            res += x;
        }
        println!("{}", res);
    }

Clippy complains that the for loop can be substituted by an `if let`, but I don't really see why. Does Some(12) implement some kind of iterator?

The correct solution is putting `if let Some(x) = option {,` I just can't see how that makes any sense.

source: [https://github.com/rust-lang/rustlings/blob/master/exercises/clippy/clippy2.rs](https://github.com/rust-lang/rustlings/blob/master/exercises/clippy/clippy2.rs)
## [5][Blog Post: Parsing Library in Rust, part 1. - lexer](https://www.reddit.com/r/rust/comments/ffilrz/blog_post_parsing_library_in_rust_part_1_lexer/)
- url: https://blog.frondeus.pl/parser-1/
---

## [6][This TabNine completion flabbergasted me...](https://www.reddit.com/r/rust/comments/ffu74e/this_tabnine_completion_flabbergasted_me/)
- url: https://www.reddit.com/r/rust/comments/ffu74e/this_tabnine_completion_flabbergasted_me/
---
&amp;#x200B;

https://preview.redd.it/ga3q2cz19nl41.png?width=1109&amp;format=png&amp;auto=webp&amp;s=17499293cd4edba8d58ffe050608c908603fdc72

I just enabled the semantic completion in TabNine. I was expecting to get standard code intel, like methods on the type of a variable, in line with what the docs say about the feature, but...

So what's happening here. I have a vector with buffers, but due to the `IoSlice` API, I have to get a `&amp;mut` to each individual buffer. That's quite unusual in Rust and requires a bit of hacking with quite some risk for off-by-one errors, so I thought I'd do a quick `assert` to verify that at least `read_bufs` ends up having the right amount of references, but then TabNine already knew that???

Just to be clear, it's the first time that I use this API, so I don't have other code with this pattern, it's the first time I have a variable called `read_bufs`, ever and I rarily play with variables named `length` either. This is almost freaky. I'm not using the cloud feature. On top of that, `length` here would seem more related to `read_vecs`, but TabNine knew that wasn't the one that needed checking. 

I'm impressed, no, I'm flabbergasted.
## [7][Looking for WASM compatible math library supporting BigInt and arbitrary precision calculations](https://www.reddit.com/r/rust/comments/fftd2m/looking_for_wasm_compatible_math_library/)
- url: https://www.reddit.com/r/rust/comments/fftd2m/looking_for_wasm_compatible_math_library/
---
I am currently exploring opportunities for my next project. The purpose of this project is mainly learning more Rust, exploring a new Frontend (Vue or Yew), and getting familiar with WASM.

My idea was to build an arbitrary precision calculator PWA. Looking for math libraries I found Rug. This crate seems to be well-maintained but since it is some wrapper for some well-known C libraries, it seems that it cannot be compiled to WASM.

Do you know any alternatives? If not, is it feasible to write such a library on my own?
## [8][Logging facade `log` ought to be deprecated in favor of non-global loggers](https://www.reddit.com/r/rust/comments/ffjjjz/logging_facade_log_ought_to_be_deprecated_in/)
- url: https://www.reddit.com/r/rust/comments/ffjjjz/logging_facade_log_ought_to_be_deprecated_in/
---
I believe the \`log\` crate has damaged the Rust ecosystem by introducing a global logging facade that's quite popular (due to it being in the rust-lang-nursey).

&amp;#x200B;

The problem with this facade is that it's global, meaning that multiple instances of objects of a library using the logger won't allow the user to distinguish between the objects, without said objects printing their addresses or some other information.

Another problem is that we can't properly test for personally-identifiable-information (PII) in logs due to the need for running module tests for PII in a single thread to avoid interlaced logs.

&amp;#x200B;

It's especially problematic when you have a tonne of trace logs but want to enable only a single context.

&amp;#x200B;

With a contextual loggers like slog you can tag each logger object and create a filter so you only print the logs you're interested in. This saves on computation time and the need for log post-processing, and is infinitely nestable by creating sub-loggers.

&amp;#x200B;

My goal with this post is to encourage - ideally - the deprecation/change of this entire crate to make the logging macros always take a logging object as their first argument.
## [9][rectangle-pack - a bin packer designed to conform to any two or three dimensional use case](https://www.reddit.com/r/rust/comments/ffmqhc/rectanglepack_a_bin_packer_designed_to_conform_to/)
- url: https://github.com/chinedufn/rectangle-pack
---

## [10][Crate specific profile in workspace](https://www.reddit.com/r/rust/comments/ffu92n/crate_specific_profile_in_workspace/)
- url: https://www.reddit.com/r/rust/comments/ffu92n/crate_specific_profile_in_workspace/
---
So I'm working on an embedded project where I have several crates (related to the full project) and each one's in its own git repo. I was thinking if, for ease of management, I could place all the crates in a single repo, which is when I came across cargo workspaces. The only issue is that the crates have to be compiled with different custom target specs (using xargo). Each crate contains it's required target spec and a .cargo/config pointing to the spec file. When placed in a workspace, it seems like each crates specific target doesn't seem to be used. I get a warning saying ```warning: profiles for the non root package will be ignored, specify profiles at the workspace root``` which seems to suggest that I can't have seperate target specs for each crate, which is not what I want. Is there a way to specify each crates target/profile instead of specifying a single one at the workspace root?
## [11][Announcing Kaos, chaotic testing harness for Rust (Resilience Testing)](https://www.reddit.com/r/rust/comments/ffc81l/announcing_kaos_chaotic_testing_harness_for_rust/)
- url: https://docs.rs/kaos/0.1.1-alpha.2/kaos/
---

## [12][New version of X11 rust bindings - x11rb](https://www.reddit.com/r/rust/comments/fffgyr/new_version_of_x11_rust_bindings_x11rb/)
- url: https://www.reddit.com/r/rust/comments/fffgyr/new_version_of_x11_rust_bindings_x11rb/
---
[A while ago](https://www.reddit.com/r/rust/comments/dgehco/announcing_x11_rust_bindings_x11rb/) I reported that I started working on [x11rb](https://crates.io/crates/x11rb), the X11 rust bindings. A lot has happened since then and today version 0.4.0 was released. Thanks go to [everyone who helped](https://github.com/psychon/x11rb/graphs/contributors).

With today's release, x11rb now supports the full X11 protocol as described by the xcb-proto XML descriptions. This also includes exchanging file descriptors (FDs) with the X11 server (fd passing).

Basically, this thing should be fully functional and only some polishing is missing.

There are two kinds of X11 connections in x11rb:

* `XCBConnection` is based on a FFI with libxcb.
* `RustConnection` is completely Rust-based (and thus does not support FD passing).

Most of the high-level code is shared between the two kinds of connections and they only differ in how the bytes of a request are sent and in how replies and events are received.

If you want to help, here are some questions:

* Is there a nice way to do FD passing over `std::os::unix::net::UnixStream` in a non-ugly way? Ugly ways involve abandoning `Write` and `Read`: [https://github.com/psychon/x11rb/issues/254](https://github.com/psychon/x11rb/issues/254)
* A lot of code is generated from a XML description of the X11 protocol. Generated code does not belong in Git, but still it would be nice to have this in releases. Is there a nice way to integrate this with Cargo so that it only runs the code generator only "builds from git" and ships the generated code with releases? [https://github.com/psychon/x11rb/issues/190](https://github.com/psychon/x11rb/issues/190)
* How would you implement "write and read to a socket at the same time"? To avoid deadlocks, an X11 client may never block in a `write()` call and must always be able to also read from the server. [https://github.com/psychon/x11rb/issues/222](https://github.com/psychon/x11rb/issues/222)

Thanks for any input and happy X11'ing.
