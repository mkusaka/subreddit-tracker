# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (36/2020)!](https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

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
## [3][Rust in 2021: Discoverability](https://www.reddit.com/r/rust/comments/inxwka/rust_in_2021_discoverability/)
- url: https://jyn514.github.io/2020/09/05/Rust-in-2021.html
---

## [4][legion v0.3 released and now part of the Amethyst organization](https://www.reddit.com/r/rust/comments/io6evk/legion_v03_released_and_now_part_of_the_amethyst/)
- url: https://amethyst.rs/posts/legion-ecs-v0.3
---

## [5][Something that blew my mind about scoped variables](https://www.reddit.com/r/rust/comments/inoipw/something_that_blew_my_mind_about_scoped_variables/)
- url: https://www.reddit.com/r/rust/comments/inoipw/something_that_blew_my_mind_about_scoped_variables/
---
To access child rcs/arcs without locking the parent, I used to do this:

    let position;
    {
        let player = player.read().unwrap();
        position = player.position.clone()
    }

But today I learned you can assign variables directly to the scope!

    let position = {
        let player = player.read().unwrap();
        player.position.clone()
    };
## [6][Rust on Haiku: the Case of the Disappearing Deceased Threads](https://www.reddit.com/r/rust/comments/intufd/rust_on_haiku_the_case_of_the_disappearing/)
- url: https://www.haiku-os.org/blog/nielx/2020-09-06_rust_on_haiku_the_case_of_the_disappearing_deceased_threads/
---

## [7][Essay found in 2030: What Killed Haskell Could Kill Rust](https://www.reddit.com/r/rust/comments/io45i3/essay_found_in_2030_what_killed_haskell_could/)
- url: https://gist.github.com/graninas/22ab535d2913311e47a742c70f1d2f2b
---

## [8][Upgrading to a Threadripper for Rust Development](https://www.reddit.com/r/rust/comments/inn005/upgrading_to_a_threadripper_for_rust_development/)
- url: https://bobweb.co/article/threadripper-meets-rustc
---

## [9][Any robust linear optimisation crates that work with Rust?](https://www.reddit.com/r/rust/comments/io5s17/any_robust_linear_optimisation_crates_that_work/)
- url: https://www.reddit.com/r/rust/comments/io5s17/any_robust_linear_optimisation_crates_that_work/
---
The most obvious choice [LpSolve](https://crates.io/crates/lpsolve) is a 404 link.

I've see a couple of research projects [MiniLP](https://crates.io/crates/minilp) and [Rulp](https://crates.io/crates/rulp) but they even state, not for production use.

Anyone know of anything better?
## [10][Translating "Processing" to Rust, adding some OpenGL and OpenCV functionality](https://www.reddit.com/r/rust/comments/inv8sl/translating_processing_to_rust_adding_some_opengl/)
- url: https://www.reddit.com/r/rust/comments/inv8sl/translating_processing_to_rust_adding_some_opengl/
---
I love Processing. It's such a simple language. My exx wanted to learn programming so I gave it to ger and she, a graphic designer, learned everything about programming from Processing in matter of days. I want to translate, note, translate, not bind, Processing to Rust, and add some additional OpenGL and OpenCV functionalities. I've found [this guide](https://craftinginterpreters.com/the-lox-language.html) that helps you design and develop an interpreted language. But is there anything extra I should know? Any books that would help me? Thanks.
## [11][Syntax proposal: Refutable patterns in `for`?](https://www.reddit.com/r/rust/comments/inl1q8/syntax_proposal_refutable_patterns_in_for/)
- url: https://www.reddit.com/r/rust/comments/inl1q8/syntax_proposal_refutable_patterns_in_for/
---
I was hoping to be able to do something like

    for Err(e) in my_results.iter() {
        // automatically skip everything that is not an `Err`
    }

I expect this to be equivalent to:

    for item in my_results.iter() {
        if let Err(e) = item {
            // body
        }
    }

I was surprised that this is not allowed. Has such syntax ever been proposed? I couldn't find any relevant RFC.
## [12][Is there anything missing in the Embedded world of Rust? Libraries, tool chains, support for MCUs?](https://www.reddit.com/r/rust/comments/inuhj2/is_there_anything_missing_in_the_embedded_world/)
- url: https://www.reddit.com/r/rust/comments/inuhj2/is_there_anything_missing_in_the_embedded_world/
---
I’ve got a bit of free time and want to contribute to Rust, particularly in the embedded area. Are there any libraries that don’t exist yet for embedded devices, that you wish you had? Support for any popular tools, missing MCUs, etc.?
