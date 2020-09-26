# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (39/2020)!](https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.46]](https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting:
# Rules for individuals:

 * Don't create top-level comments; those are for employers.
 * Feel free to reply to top-level comments with on-topic questions.
 * Anyone seeking work should reply to my stickied top-level comment.
 * Meta-discussion should be reserved for the distinguished comment at the very bottom.

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
## [3][Amazon has the most Job Descriptions for Rust](https://www.reddit.com/r/rust/comments/izzyb5/amazon_has_the_most_job_descriptions_for_rust/)
- url: https://news.efinancialcareers.com/us-en/3004512/rust-vs-c-hedge-fund-jobs
---

## [4][Rust in 2021: We've done a lot, more to de done](https://www.reddit.com/r/rust/comments/j01n3v/rust_in_2021_weve_done_a_lot_more_to_de_done/)
- url: https://popzxc.github.io/rust-2021
---

## [5][A high performance code minimap render](https://www.reddit.com/r/rust/comments/j03qew/a_high_performance_code_minimap_render/)
- url: https://www.reddit.com/r/rust/comments/j03qew/a_high_performance_code_minimap_render/
---
I wrote a tool [`code-minimap`](https://github.com/wfxr/code-minimap) for generating text minimaps for code files.

It is streaming rendering, consumes very little fixed memory, and supports arbitrary scaling. You can use it to implement IDE-like minimap for the terminal text editor with high speed. All the [rust source code](https://github.com/rust-lang/rust/tree/1.46.0) (latest stable version 1.46, over 1,000,000 lines) as input only takes 323ms in my PC (detailed benchmark can be found in the repository).

I used it to create a vim plugin [`minimap.vim`](https://github.com/wfxr/minimap.vim):

[minimap.vim](https://i.redd.it/rk0099mxygp51.gif)

**Project link**:

code-minimap: [https://github.com/wfxr/code-minimap](https://github.com/wfxr/code-minimap)

minimap.vim: [https://github.com/wfxr/minimap.vim](https://github.com/wfxr/minimap.vim)

Issues and feedbacks are welcomed!

&amp;#x200B;

**EDIT:**

This tool is for **text files**, not for games.
## [6][FOSDEM 2021 - Online only](https://www.reddit.com/r/rust/comments/j03i80/fosdem_2021_online_only/)
- url: https://www.reddit.com/r/rust/comments/j03i80/fosdem_2021_online_only/
---
[FOSDEM 2021 will be online only](https://fosdem.org/2021/). The organizers of the previous RUST devroom (2018, 2019, 2020) decided that we won't be organizing it this year as a huge part of FOSDEM is getting people in the same room and being able to create face to face discussions on the back of the talks. After all, DEM = DEveloper Meeting. There are plenty of great online conferences and recorded talks devoted to Rust so it seems superfluous to us.

We thought it would be polite to inform the community in case someone else would like to submit a devroom proposal and would have otherwise assumed @lucab, @itkovian, and myself would be submitting the proposal this year.  

If you're interested, we can provide you with previous proposal templates, talk review google sheets, and other material we use to manage the devroom.

Stay safe!

NB: [This is a cross post](https://users.rust-lang.org/t/rust-2021-online-only/49333)
## [7][Rust 2021: Continue](https://www.reddit.com/r/rust/comments/izk75l/rust_2021_continue/)
- url: https://www.reddit.com/r/rust/comments/izk75l/rust_2021_continue/
---
I only have a little thought on Rust 2021, not worth a separate blog post I guess. There are 'only' 3 things I'm waiting for:

* specialization
* const generics
* generators

These are the only things that I ever missed. Other features (tbh - including async) were, for me, just nice to have additions.

As far as I remember, there were plans to close pending features in 2020. We could see some really nice progress but in my opinion, there are still too many things open. So 2021 needs, again in my opinion, only to continue ideas from 2020.
## [8][has anyone compiled a gtk-rs program for Windows lately?](https://www.reddit.com/r/rust/comments/j04t2y/has_anyone_compiled_a_gtkrs_program_for_windows/)
- url: https://www.reddit.com/r/rust/comments/j04t2y/has_anyone_compiled_a_gtkrs_program_for_windows/
---
I've built a few simple GTK programs for Linux recently and I've been trying to build Windows versions. I'm following the [tutorial](https://gtk-rs.org/docs-src/tutorial/cross) and I've got as far as installing MinGW and setting up the `x86_64-pc-windows-gnu` target, but the build process fails at the linking stage.

I'm not sure if the version of MinGW packaged for my OS (Solus) is missing anything - I tried in an Ubuntu VM and didn't get any further. The tutorial basically points at Arch being the best-supported platform so I'm trying to get that set up in a VM now. am I better off just building MinGW from source?

I'm also a bit confused about the `PKG_CONFIG_PATH` environment variable - it points to `/usr/i686-w64-mingw32/lib/pkgconfig` in the example. on my system MinGW is installed at /usr/share/mingw-w64 (fine, I can change the path) but how come it points to i686 instead of x86\_64? is this for backwards compatibility? Also, I understand that pkg-config is a tool for managing libraries for compilation purposes, but there's no `pkgconfig` file/folder inside my MinGW installation. Am I missing something here?

The tutorial mentions a precompiled GTK .dll but the link is broken, which makes me wonder if the tutorial is out of date. Does anyone know where I can find the .dll?

Is building *on* Windows an option? I still have Windows 8.1 on a laptop and could build on there if needs be, but I'd rather cross-compile.

Obviously, no hard feelings toward the developers of gtk-rs - I like it a lot so far. I'd be interested to hear your experiences in cross-compiling GTK programs or GUI programs in general!
## [9][I just released the first version of Audiobench, a free/open-source modular synthesizer written almost entirely in Rust.](https://www.reddit.com/r/rust/comments/izkey2/i_just_released_the_first_version_of_audiobench_a/)
- url: https://www.reddit.com/r/rust/comments/izkey2/i_just_released_the_first_version_of_audiobench_a/
---
Downloads and a getting started guide [can be found here](https://bit.ly/audio_bench). The GitHub repository [can be found here](https://github.com/joshua-maros/audiobench).

[A screenshot of the main interface.](https://preview.redd.it/6r4l0fdi0bp51.png?width=1356&amp;format=png&amp;auto=webp&amp;s=b6bfc6f610c1875ef733c152f4cf3e696d26bbbc)
## [10][Knurling-rs changelog #2](https://www.reddit.com/r/rust/comments/izlqet/knurlingrs_changelog_2/)
- url: https://ferrous-systems.com/blog/knurling-changelog-2/
---

## [11][Rust API Challenges](https://www.reddit.com/r/rust/comments/izoxn4/rust_api_challenges/)
- url: https://www.reddit.com/r/rust/comments/izoxn4/rust_api_challenges/
---
Here are two challenges.

They may look similar, but the first is much easier than the second. Both are solvable.

If you're a beginning Rustacean, treat the first as the challenge, and ignore the second. If you're an advanced Rustacean, treat the first as a warmup, and the second as the challenge.

Good luck!

# Challenge 1

Implement this API:

    pub type Challenge1&lt;A, B&gt;;
    
    impl&lt;A, B&gt; Challenge1&lt;A, B&gt; {
        pub fn left(a: A) -&gt; Self;
        pub fn right(b: B) -&gt; Self;
        pub fn consume&lt;T&gt;(self, f: impl FnOnce(A) -&gt; T, g: impl FnOnce(B) -&gt; T) -&gt; T;
    }

Your implementation should satisfy these rules:

    Challenge1::left(a).consume(f, g) = f(a)
    Challenge1::right(b).consume(f, g) = g(b)

# Challenge 2

Implement this API:

    pub type Challenge2&lt;'r, A, B&gt;;
    
    impl&lt;'r, A, B&gt; Challenge2&lt;'r, A, B&gt; {
        pub fn left(self) -&gt; A;
        pub fn right(self) -&gt; B;
        pub fn new&lt;T: 'r&gt;(t: T, f: impl 'r + FnOnce(T) -&gt; A, g: impl 'r + FnOnce(T) -&gt; B) -&gt; Self;
    }

Your implementation should satisfy these rules:

    Challenge2::new(t, f, g).left() = f(t)
    Challenge2::new(t, f, g).right() = g(t)
## [12][Custom user defined literals using proc_macro](https://www.reddit.com/r/rust/comments/izt63q/custom_user_defined_literals_using_proc_macro/)
- url: https://www.5snb.club/posts/2020/09/25/custom-literals-in-rust.html
---

