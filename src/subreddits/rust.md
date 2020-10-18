# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (42/2020)!](https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.47]](https://www.reddit.com/r/rust/comments/jb4cds/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/jb4cds/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/).

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

COMPANY: *[Company name; optionally link to your company's website or careers page.]*

TYPE: *[Full time, part time, internship, contract, etc.]*

DESCRIPTION: *[What does your company do, and what are you using Rust for? How much experience are you seeking and what seniority levels are you hiring for? The more details the better.]*

LOCATION: *[Where are your office or offices located? If your workplace language isn't English-speaking, please specify it.]*

ESTIMATED COMPENSATION: *[Be courteous to your potential future colleagues by attempting to provide at least a rough expectation of wages/salary. If you are listing several positions in the "Description" field above, then feel free to include this information inline above, and put "See above" in this field. If compensation is negotiable, please attempt to provide at least a base estimate from which to begin negotiations. If compensation is highly variable, then feel free to provide a range. If compensation is expected to be offset by other benefits, then please include that information here as well. If you don't have firm numbers but do have relative expectations of candidate expertise (e.g. entry-level, senior), then you may include that here. If you truly have no information, then put "Uncertain" here.* ***This is a new field in our template; please see the meta comment below to discuss it***.*]*

REMOTE: *[Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]*

VISA: *[Does your company sponsor visas?]*

CONTACT: *[How can someone get in touch with you?]*
## [3][Announcing shared_arena, a thread-safe memory pool](https://www.reddit.com/r/rust/comments/jddens/announcing_shared_arena_a_threadsafe_memory_pool/)
- url: https://github.com/sebastiencs/shared-arena
---

## [4][Announcing yew-state v0.4.0](https://www.reddit.com/r/rust/comments/jd7ty1/announcing_yewstate_v040/)
- url: https://github.com/intendednull/yew-state/releases/tag/v0.4.0
---

## [5][Vimscript Language Server in Rust](https://www.reddit.com/r/rust/comments/jd01h3/vimscript_language_server_in_rust/)
- url: https://github.com/google/vimscript-language-server
---

## [6][How to mutate a global variable with tokio/futures ?](https://www.reddit.com/r/rust/comments/jdeayw/how_to_mutate_a_global_variable_with_tokiofutures/)
- url: https://www.reddit.com/r/rust/comments/jdeayw/how_to_mutate_a_global_variable_with_tokiofutures/
---
So basically I have this non-blocking set\_interval function that does something every 5 seconds.

I'm trying to figure out how I could make it increment the "mutate" integer that should be available throughout the whole main function.

I've read that I'm supposed to use Arc but I have no idea how to implement this myself.

    [dependencies]
    tokio = {version = "0.3.0", features = ["rt", "rt-multi-thread", "macros", "time"]}
    
    use tokio::time::{self, Duration};
    use std::future::Future;
     
    #[tokio::main]
    async fn main() {
    
        let mut mutate = 0; // how do i increment this inside the interval ?
    
        set_interval(|| async {
            mutate = mutate + 1;
            println!("{}", mutate);
        }, 5000);
        
        time::sleep(Duration::from_millis(60000)).await;
    }
     
     
    fn set_interval&lt;F, Fut&gt;(mut task: F, dur: u64)
    where
        F: Send + 'static + FnMut() -&gt; Fut,
        Fut: Future&lt;Output = ()&gt; + Send + 'static,
    {
        let mut interval = time::interval(Duration::from_millis(dur));
        tokio::spawn(async move {
            loop {
                interval.tick().await;
                tokio::spawn(task());
            }
        });
    }
## [7][HPACK: Huffman translation matrix](https://www.reddit.com/r/rust/comments/jdcowt/hpack_huffman_translation_matrix/)
- url: https://kristijansedlak.medium.com/hpack-huffman-translation-matrix-f3cae44bfe8c
---

## [8][Announcement: xshell, ergonomic "bash" scripting in Rust](https://www.reddit.com/r/rust/comments/jctkpi/announcement_xshell_ergonomic_bash_scripting_in/)
- url: https://docs.rs/xshell/0.1.2/xshell/index.html
---

## [9][My first release to crates.io: akima spline interpolation](https://www.reddit.com/r/rust/comments/jd2lhu/my_first_release_to_cratesio_akima_spline/)
- url: https://www.reddit.com/r/rust/comments/jd2lhu/my_first_release_to_cratesio_akima_spline/
---
The akima interpolation is a very pretty one as it tries to mimic the natural flow when drawing lines through points with the hand. This is a modified version that builds on its original philosophy.

I made this, because I didn't like the options available because of either bad interpolation due to ringing or missing extrapolation support.

This really takes a big weight off me, as this is something that I needed since I first began using rust half a year ago. 

Anyway, I just wanted to share this, so if someone has either never seen akima interpolation or desperately needs a quick and easy to use library for pretty interpolation go check it out!

[https://crates.io/crates/makima\_spline](https://crates.io/crates/makima_spline)

Feedback is welcome!
## [10][New crate announcement: enhanced_enum](https://www.reddit.com/r/rust/comments/jd3xhj/new_crate_announcement_enhanced_enum/)
- url: https://www.reddit.com/r/rust/comments/jd3xhj/new_crate_announcement_enhanced_enum/
---
Enhanced Fieldless Enumerations and Associated Array Types

In Rust, enumerations can contain data fields which is a powerful language
feature. However not all enums have data fields. Fieldless enums are simply a
list of variants. This crate provides many features for fieldless enums which
are difficult or impossible to provide for enums with data fields.

This crate contains a single item: the enhanced_enum! macro which generates an enum.

[Documentation](https://docs.rs/enhanced_enum/0.2.1/enhanced_enum/)

Example: a histogram for counting DNA nucleotides. This re-implements the
example from the documentation for the [Trait std::ops::Index](https://doc.rust-lang.org/std/ops/trait.Index.html).

    enhanced_enum::enhanced_enum!(Nucleotide {
        A,
        C,
        G,
        T,
    });
    let nucleotide_count = NucleotideArray::&lt;usize&gt;::new_with(|x| match x {
        Nucleotide::A =&gt; 14,
        Nucleotide::C =&gt; 9,
        Nucleotide::G =&gt; 10,
        Nucleotide::T =&gt; 12
    });
    assert_eq!(nucleotide_count[Nucleotide::A], 14);
    assert_eq!(nucleotide_count[Nucleotide::C], 9);
    assert_eq!(nucleotide_count[Nucleotide::G], 10);
    assert_eq!(nucleotide_count[Nucleotide::T], 12);
## [11][Making a Snake Clone with Bevy](https://www.reddit.com/r/rust/comments/jcv4ck/making_a_snake_clone_with_bevy/)
- url: https://mbuffett.com/posts/bevy-snake-tutorial/
---

## [12][Announcing timed - quick profiling crate](https://www.reddit.com/r/rust/comments/jdgfbh/announcing_timed_quick_profiling_crate/)
- url: https://www.reddit.com/r/rust/comments/jdgfbh/announcing_timed_quick_profiling_crate/
---
[https://github.com/y2kappa/timed](https://github.com/y2kappa/timed)

I built this because I am running rust on aws lambda and I can't really replicate the running environment fully on my machine, as I cannot dump to a file or profile the binary while ssh'd on the machine. I needed something that would print out so I can analyse it later. (Btw, I can run the aws lambda locally just fine, but I cannot fully replicate the environment with the db connections and all the IO dependencies.)

Btw, this is my first ever crate, so any feedback is super useful.

In the [TODO.](https://TODO.MD)md there are a lot of things I want to work on this so I am thinking to also create Issues that people could contribute to.
