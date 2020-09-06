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
## [3][Introducing Rune, a new stack-based dynamic programming language for Rust with async support.](https://www.reddit.com/r/rust/comments/in67d3/introducing_rune_a_new_stackbased_dynamic/)
- url: https://rune-rs.github.io/rune/
---

## [4][Syntax proposal: Refutable patterns in `for`?](https://www.reddit.com/r/rust/comments/inl1q8/syntax_proposal_refutable_patterns_in_for/)
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
## [5][parser for Clickhouse SQL dialect.](https://www.reddit.com/r/rust/comments/iniezi/parser_for_clickhouse_sql_dialect/)
- url: https://github.com/superjobru/clickhouse-sql-parser
---

## [6][Announcing rasn: A safe #[no_std] ASN.1 codec framework](https://www.reddit.com/r/rust/comments/ina4ep/announcing_rasn_a_safe_no_std_asn1_codec_framework/)
- url: https://github.com/XAMPPRocky/rasn/releases/tag/v0.2.0
---

## [7][My first month with Rust](https://www.reddit.com/r/rust/comments/ine4fl/my_first_month_with_rust/)
- url: https://www.reddit.com/r/rust/comments/ine4fl/my_first_month_with_rust/
---
Working with Rust gives a feeling of 'we are in it together'. 

It seems that there were two different worlds for general purpose languages. One for systems programming and the other for the programmer. C/C++ dominating the first and was as popular in that era and is even now. Working with them felt as if, the most important user of my program is the computer and compiler itself. I better understand that for it to work. 

And then we lived in an era where programming language evolved to be expressive. An era where the focus was on making the languages as a mechanism to achieve results faster, easier. The programmer took its revenge and said, I am your user, and you better be easy for me, that's when you would be successful. Languages bowed their heads and followed. Whatever you write, they show quick result. Right or wrong, that's a story for another day. The biggest downside though was performance and memory. Doesn't matter how much RAM you put in, doesn't matter how much CPU horsepower you invest in, the computer still hangs. Uninstalling rogue software and apps is a skill that now everyone specializes in.

Rust feels like a partner. When I work with Rust and use the tools around it; Rust seems to say, I know the Computer, you know the User. Let's do it together to solve the problem for which we have met. It seems to say, I would help you understand me better, so you can make the most out of the resources you have to build the solution you know best.

This feels a very different approach from what I have seen so far.
## [8][A four video series I made on using Actix-web and Sqlite with Diesel.](https://www.reddit.com/r/rust/comments/infl59/a_four_video_series_i_made_on_using_actixweb_and/)
- url: https://www.youtube.com/watch?v=HO-KMVXvXdA&amp;list=PLBok0UdvO6519tl0LzER8w_J-0pcro7q9
---

## [9][Microsoft has implemented some safety rules of Rust in their C++ static analysis tool.](https://www.reddit.com/r/rust/comments/imy9lg/microsoft_has_implemented_some_safety_rules_of/)
- url: https://devblogs.microsoft.com/cppblog/new-safety-rules-in-c-core-check/
---

## [10][Announcing serde-query](https://www.reddit.com/r/rust/comments/in0qiu/announcing_serdequery/)
- url: https://crates.io/crates/serde-query
---

## [11][Undefined Behavior: What Happened to My Code?](https://www.reddit.com/r/rust/comments/in9n2c/undefined_behavior_what_happened_to_my_code/)
- url: https://www.reddit.com/r/rust/comments/in9n2c/undefined_behavior_what_happened_to_my_code/
---
The issue of undefined behavior comes up frequently when discussing Rust and comparison of Rust to C or C++, and a lot of people seem to be confused about it or don't fully understand what it means.

While I was reading parts of u/ralphj's PhD thesis (just introduction and conclusion), I came across this excellent article about UB:

Undefined Behavior: What Happened to My Code? [https://people.csail.mit.edu/nickolai/papers/wang-undef-2012-08-21.pdf](https://people.csail.mit.edu/nickolai/papers/wang-undef-2012-08-21.pdf)

It shows several instances of C code triggering UB in real life code (Linux kernel, FreeBSD, PostgreSQL), how C compilers use UB to do optimizations, and written in an easy to understand language. It makes me appreciate Rust more.
## [12][How to directly modify the matrix in OpenCV Rust?](https://www.reddit.com/r/rust/comments/inglxu/how_to_directly_modify_the_matrix_in_opencv_rust/)
- url: https://www.reddit.com/r/rust/comments/inglxu/how_to_directly_modify_the_matrix_in_opencv_rust/
---
I'm getting tired of this. The docs say that [matrix.at](https://matrix.at)::&lt;T&gt;(i, j) = value work, however, the compiler says I'm mixing up lefthand and right hand values. Can someone help me with this? Thanks.
