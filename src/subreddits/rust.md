# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (21/2020)!](https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.43]](https://www.reddit.com/r/rust/comments/g6v98u/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/g6v98u/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/fjsj1l/official_rrust_whos_hiring_thread_for_jobseekers/).

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
## [3][My first Rust project! A CLI that sends a system notification 100 times evenly over a user-specified timespan, updating you with the time elapsed as a percentage.](https://www.reddit.com/r/rust/comments/gq4ocy/my_first_rust_project_a_cli_that_sends_a_system/)
- url: https://v.redd.it/2as6sgr4cu051
---

## [4][rust-analyzer: Changelog #26](https://www.reddit.com/r/rust/comments/gq9ppx/rustanalyzer_changelog_26/)
- url: https://rust-analyzer.github.io/thisweek/2020/05/25/changelog-26.html
---

## [5][A new fast regexp engine in Rust, supporting JS syntax (backreferences + lookarounds)](https://www.reddit.com/r/rust/comments/gq1evg/a_new_fast_regexp_engine_in_rust_supporting_js/)
- url: https://github.com/ridiculousfish/regress
---

## [6][Zero to Production in Rust #0: Foreword](https://www.reddit.com/r/rust/comments/gpydni/zero_to_production_in_rust_0_foreword/)
- url: https://www.lpalmieri.com/posts/2020-05-24-zero-to-production-0-foreword/
---

## [7][How is the Rust compiler able to tell the visible width of unicode characters?](https://www.reddit.com/r/rust/comments/gpw2ra/how_is_the_rust_compiler_able_to_tell_the_visible/)
- url: https://www.reddit.com/r/rust/comments/gpw2ra/how_is_the_rust_compiler_able_to_tell_the_visible/
---
Here is an example: \[[playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=a66dc3d831b1546c01c089aa1439449d)\]

Take this text: `读文`

(I don't speak this language, but Google translate says it means "reading," so I figure it's not anything bad.)

If I were trying to highlight a compiler error for this text as an invalid value, I would naively choose to count the characters using `chars().count()`.

As you can see by running the program, this comes up short.  

    读文;
    ^^
But if you uncomment the line with the invalid  assignment, you will see that the Rust compiler correctly highlights the entirety of the text. 

    读文;
    ^^^^

How does the compiler compute the visible width of these characters?

I have tried to look this up, but my efforts are coming up just as short as my own implementation ☹️️

My search results are almost all about UTF-8 byte width, not visible width.
## [8][AccessQueue - limit concurrent accesses (async IO backpressure)](https://www.reddit.com/r/rust/comments/gq9ds3/accessqueue_limit_concurrent_accesses_async_io/)
- url: https://github.com/withoutboats/access-queue
---

## [9][Hey y'all, I wrote a tmux workspace manager called DMUX using Rust, and I'd love if you checked it out! It's focused on easily selecting and setting up different working environments. Any feedback on either the UX or my code is very welcome!](https://www.reddit.com/r/rust/comments/gq26fj/hey_yall_i_wrote_a_tmux_workspace_manager_called/)
- url: https://github.com/zdcthomas/dmux
---

## [10][How do you effectively and iteratively develop in Rust without a REPL?](https://www.reddit.com/r/rust/comments/gq2ia0/how_do_you_effectively_and_iteratively_develop_in/)
- url: https://www.reddit.com/r/rust/comments/gq2ia0/how_do_you_effectively_and_iteratively_develop_in/
---
my only experience is with languages that feature REPLs, and so it's honestly a little confusing learning rust without a REPL.

for example, i don't know how to effectively test out a new library that i am unfamiliar with or play around with some new techniques.
## [11][[New Crate] Include files in Markdown docs](https://www.reddit.com/r/rust/comments/gq7kqp/new_crate_include_files_in_markdown_docs/)
- url: https://www.reddit.com/r/rust/comments/gq7kqp/new_crate_include_files_in_markdown_docs/
---
[https://github.com/martypapa/md-inc](https://github.com/martypapa/md-inc)

I'd love your feedback!

Essentially, it expands inline:

    Look at the following rust code:
    &lt;!--{ "file.rs" | code: rust }--&gt;
    &lt;!--{ end }--&gt;
    This will print 'Hello World' to the console.

Into:

    Look at the following rust code:
    &lt;!--{ "file.rs" | code: rust }--&gt;
    ```rust
    fn main() {
        println!("Hello, World!");
    }
    ```
    &lt;!--{ end }--&gt;
    This will print 'Hello World' to the console.

Follow link for more config options.
## [12][What am I doing wrong with Option matching ?](https://www.reddit.com/r/rust/comments/gq82jd/what_am_i_doing_wrong_with_option_matching/)
- url: https://www.reddit.com/r/rust/comments/gq82jd/what_am_i_doing_wrong_with_option_matching/
---
I have the following code, which is based upon standard Thread examples and is using flume for the channel:

    use flume::{Sender, Receiver};
    use std::thread;
    use std::time::Duration;
    
    fn main() {
    
        let (tx, rx) = flume::bounded::&lt;Option&lt;_&gt;&gt;(0);

        let producer = thread::spawn(move || {
            for sent in 1..10 {
                tx.send(Some(sent));
            }
            tx.send(None);
        });
    
        let consumer = thread::spawn(move || {
            loop {
                match rx.iter().next() {
                    None =&gt; break,
                    Some(received) =&gt; {
                        println!("Received {}", received.unwrap());
                    }
                }
            }
        });
    
        producer.join();
        consumer.join();
    }

Everything works except for the final message to end the process - the None is blowing up as it is entering the Some match arm (at least on my system):

&gt;thread '&lt;unnamed&gt;' panicked at 'called `Option::unwrap()` on a `None` value', src\main.rs:21:45

It's likely something stupid I'm doing, but I can't see what. I would expect the None to be matched and break the loop, not enter the subsequent arm.

Thanks.
