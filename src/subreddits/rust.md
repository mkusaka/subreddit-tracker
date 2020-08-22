# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

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
## [3][google/autocxx - calling C++ from Rust in a heavily automated, but safe, fashion](https://www.reddit.com/r/rust/comments/iefeum/googleautocxx_calling_c_from_rust_in_a_heavily/)
- url: https://github.com/google/autocxx
---

## [4][Rust Memory Container Cheat-sheet, publish on GitHub](https://www.reddit.com/r/rust/comments/idwlqu/rust_memory_container_cheatsheet_publish_on_github/)
- url: https://i.redd.it/220xo2f6wci51.png
---

## [5][[@rylev, @fasterthanlime ] Before Main: How Executables Work on Linux](https://www.reddit.com/r/rust/comments/iegw13/rylev_fasterthanlime_before_main_how_executables/)
- url: https://youtu.be/jR2hUhjcAXI
---

## [6][Colors - simple terminal colors fetch](https://www.reddit.com/r/rust/comments/ieftwu/colors_simple_terminal_colors_fetch/)
- url: https://i.redd.it/xs1uegrm0ji51.png
---

## [7][Learning Empathy From Pokemon Blue - RustConf 2020 Closing Keynote](https://www.reddit.com/r/rust/comments/ie8q85/learning_empathy_from_pokemon_blue_rustconf_2020/)
- url: https://youtu.be/RNsEsZbXE-4
---

## [8][Oxigraph, a WIP graph database in Rust implementing RDF/SPARQL](https://www.reddit.com/r/rust/comments/ie3yzi/oxigraph_a_wip_graph_database_in_rust/)
- url: https://github.com/oxigraph/oxigraph
---

## [9][[knurling] defmt, a highly efficient Rust logging framework for embedded devices](https://www.reddit.com/r/rust/comments/idwb9c/knurling_defmt_a_highly_efficient_rust_logging/)
- url: https://ferrous-systems.com/blog/defmt/
---

## [10][How to use buffering with a file that is both write/read?](https://www.reddit.com/r/rust/comments/ie8z2n/how_to_use_buffering_with_a_file_that_is_both/)
- url: https://www.reddit.com/r/rust/comments/ie8z2n/how_to_use_buffering_with_a_file_that_is_both/
---
So, I can create a file:

    let file = OpenOptions::new()
                .read(true)
                .write(true)
                .create(true)
                .open("foo.txt");

Here I can do all, but not use buffering.

But how I can support buffering for write/read? I don't see how to wrap both BufReader/BufWriter

P.D: This is for an interpreter runtime, and [similar to python where this is possible](https://docs.python.org/3/library/functions.html#open):

    open("foo.txt", "+") 

I have thought of wrap in an Rc:

    pub struct File {
        _in: BufReader&lt;Rc&lt;fs::File&gt;&gt;,
        _out: BufWriter&lt;Rc&lt;fs::File&gt;&gt;,
        path: PathBuf,
    }

but not know if this is the best way (or safe!).
## [11][Planning to learn rust, learning resources much appreciated](https://www.reddit.com/r/rust/comments/ieag78/planning_to_learn_rust_learning_resources_much/)
- url: https://www.reddit.com/r/rust/comments/ieag78/planning_to_learn_rust_learning_resources_much/
---
Hello folks, I have been planning to learn rust and hoping to implement the plan this winter. I learn best through courses as a foundation and build on top of it as I go. If there are no such courses, I will learn Go instead. My Interest is Distributed Systems and primary language is JS. Idea is to have a working understanding of Web Assembly and to be able to wield rust.
## [12][Kernel printing with Rust.](https://www.reddit.com/r/rust/comments/ie5fsl/kernel_printing_with_rust/)
- url: https://not-matthias.github.io/kernel-printing-with-rust/
---

