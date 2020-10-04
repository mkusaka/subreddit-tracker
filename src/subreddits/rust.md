# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (40/2020)!](https://www.reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

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
## [3][Easy beginner Rust projects to contribute to as an experienced C++ engineer](https://www.reddit.com/r/rust/comments/j4wi0m/easy_beginner_rust_projects_to_contribute_to_as/)
- url: https://www.reddit.com/r/rust/comments/j4wi0m/easy_beginner_rust_projects_to_contribute_to_as/
---
Hi there,
I have been working with C++ for about 5 years and now I would like to learn Rust. What projects can you recommend me to contribute to in order to get in quickly into Rust?
Thanks
## [4][This Month in Rust OSDev (September 2020)](https://www.reddit.com/r/rust/comments/j4xujk/this_month_in_rust_osdev_september_2020/)
- url: https://rust-osdev.com/this-month/2020-09/
---

## [5][ox - An independent Rust text editor that runs in your terminal](https://www.reddit.com/r/rust/comments/j4nbwl/ox_an_independent_rust_text_editor_that_runs_in/)
- url: https://github.com/curlpipe/ox
---

## [6][What are some concrete libraries missing in Rust's ecosystem?](https://www.reddit.com/r/rust/comments/j4xj37/what_are_some_concrete_libraries_missing_in_rusts/)
- url: https://www.reddit.com/r/rust/comments/j4xj37/what_are_some_concrete_libraries_missing_in_rusts/
---
I've heard that one of the biggest issues with Rust is the library ecosystem, but I've not seen many concrete examples of the libraries that are missing.

What are some libraries that you've found yourself wanting that doesn't exist?

Or do people just mean that there are too many unstable (0.\*.\*) libraries when they say this?
## [7][Blog Post: Fast Thread Locals In Rust](https://www.reddit.com/r/rust/comments/j4iy50/blog_post_fast_thread_locals_in_rust/)
- url: https://matklad.github.io/2020/10/03/fast-thread-locals-in-rust.html
---

## [8][Why does the assembly for this code compare a 'char'-value to 0x11_0000?](https://www.reddit.com/r/rust/comments/j4wovc/why_does_the_assembly_for_this_code_compare_a/)
- url: https://www.reddit.com/r/rust/comments/j4wovc/why_does_the_assembly_for_this_code_compare_a/
---
Inspired by one of the replies to an earlier post today, I tried the following code:

[https://play.rust-lang.org/?version=nightly&amp;mode=release&amp;edition=2018&amp;gist=f568fc3563055e460813cddbc8cededf](https://play.rust-lang.org/?version=nightly&amp;mode=release&amp;edition=2018&amp;gist=f568fc3563055e460813cddbc8cededf)

The function convert\_to\_u32 shows how the rust-compiler is smart enough to reuse the memory allocation of a Vec after iterating through it and collecting it to another Vec. In this case, iterating through a Vec&lt;char&gt; and returning a  Vec&lt;u32&gt;.

But the assemly code seemed to contain an unnecessary loop. On inspection, it seemed to compare each element in the Vec&lt;char&gt; to 0x11\_0000, and truncate the Vec at the first such element. The compiled program is still correct, since the value '0x11\_0000' cannot occur in a Vec&lt;char&gt;, and the check thus has no effect. But it is a missed opportunity for optimization, since without it the loop could go away completely.

Note that the code as written is unsound (you're not allowed to use unsafe code to create a 'char' with value 0x11\_0000), so the compiler really is free to do whatever it likes. But the comparison to 0x11\_0000 remains even if the main-function is removed (which makes the code sound again).

Does anyone know why the compiler is comparing to 0x11\_0000?
## [9][What is the Rust story for a package like Go's context? (i.e.deadlines, timeouts, cancellation signals)](https://www.reddit.com/r/rust/comments/j4xx7i/what_is_the_rust_story_for_a_package_like_gos/)
- url: https://www.reddit.com/r/rust/comments/j4xx7i/what_is_the_rust_story_for_a_package_like_gos/
---
The context package in the Go standard library provides an interface for cancellation signals, timeouts, and deadlines. https://golang.org/pkg/context/. As it's a part of the standard library, external Go package authors make consistent use of the package in code they write, meaning you can consistently rely on third-party packages accepting a Context type for cancellation signals.

What is the Rust equivalent for this? Is there a package or a preferred pattern? What do you personally use?
## [10][Is it safe to transmute from Vec&lt;char&gt; to Vec&lt;u32&gt;?](https://www.reddit.com/r/rust/comments/j4qk62/is_it_safe_to_transmute_from_vecchar_to_vecu32/)
- url: https://www.reddit.com/r/rust/comments/j4qk62/is_it_safe_to_transmute_from_vecchar_to_vecu32/
---
I believe that char and u32 have the same layout, and char can safely be transmuted to u32. Does this mean that Vec&lt;char&gt; can be transmuted to Vec&lt;u32&gt;? I ask because clippy throws an error saying they have mismatched layouts.

Playground link: [https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=fb04f03e1384c2ae13273b22e55883ac](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=fb04f03e1384c2ae13273b22e55883ac)
## [11][Rust meets the web - a clash of programming paradigms](https://www.reddit.com/r/rust/comments/j4ot6b/rust_meets_the_web_a_clash_of_programming/)
- url: https://www.jakobmeier.ch/blogging/Rust_on_the_Web.html
---

## [12][Rust Starter Kit 2020](https://www.reddit.com/r/rust/comments/j4gnvh/rust_starter_kit_2020/)
- url: https://wiki.alopex.li/RustStarterKit2020
---

