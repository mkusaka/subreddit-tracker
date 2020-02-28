# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (9/2020)!](https://www.reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here_92020/)
- url: https://www.reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here_92020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 327](https://www.reddit.com/r/rust/comments/fabtwg/this_week_in_rust_327/)
- url: https://this-week-in-rust.org/blog/2020/02/25/this-week-in-rust-327/
---

## [3][Announcing arcs 0.3.0 - now with more algorithms and better docs](https://www.reddit.com/r/rust/comments/fatgr1/announcing_arcs_030_now_with_more_algorithms_and/)
- url: https://docs.rs/arcs/0.3.0/arcs/
---

## [4][Announcing Rust 1.41.1 | Rust Blog](https://www.reddit.com/r/rust/comments/fadbx2/announcing_rust_1411_rust_blog/)
- url: https://blog.rust-lang.org/2020/02/27/Rust-1.41.1.html
---

## [5][Is there any way to rename the main thread?](https://www.reddit.com/r/rust/comments/fauino/is_there_any_way_to_rename_the_main_thread/)
- url: https://www.reddit.com/r/rust/comments/fauino/is_there_any_way_to_rename_the_main_thread/
---
If I spawn threads within my Rust programs, I can name them by using the std::thread::Builder approach. So when the thread panics, it displays a relevant name for the thread.

However, I have several Rust processes running that can panic and be restarted by my supervisor process. The logs are not always super helpful because every Rust process has a thread called 'main', so it's not always clear who died.
## [6][Tokio v0.2.12, includes `Notify`, an async/await synchronization primitive, and `StreamMap` for dynamic merging of streams.](https://www.reddit.com/r/rust/comments/fagnqu/tokio_v0212_includes_notify_an_asyncawait/)
- url: https://github.com/tokio-rs/tokio/releases/tag/tokio-0.2.12
---

## [7][Problem with lifetimes and destructors--is this a bug?](https://www.reddit.com/r/rust/comments/faqcu3/problem_with_lifetimes_and_destructorsis_this_a/)
- url: https://www.reddit.com/r/rust/comments/faqcu3/problem_with_lifetimes_and_destructorsis_this_a/
---
I'm seeing some behavior from the result compiler involving temporaries with destructors that doesn't seem right to me. It happened with real code I was trying to make work, but I've pared it down to something fairly minimal to illustrate the issue. First, here's a wrapper around a reference that implements `Drop`:

    struct SmartPtr&lt;'a, T&gt;(&amp;'a T);
    impl&lt;'a, T&gt; Drop for SmartPtr&lt;'a, T&gt; {
        fn drop(&amp;mut self) {}
    }

The type parameter `T` doesn't really matter; I just added it to show the problem isn't related to any particular type. Here's a function that creates a temporary value holding a `SmartPtr`, which the compiler rejects with the errors in the comments.

    fn broken&lt;T&gt;(x: T) -&gt; T {
        // error[E0505]: cannot move out of `x` because it is borrowed
        if let Some(ptr) = Some(SmartPtr(&amp;x)) {
            //             ------------------
            //             |             |
            //             |             borrow of `x` occurs here
            //             a temporary with access to the borrow is created here ...
            drop(ptr);
            return x;
            //     ^ move out of `x` occurs here
        }
        unimplemented!()
    }
    // ... and the borrow might be used here, when that temporary is dropped and runs the destructor for type `std::option::Option&lt;SmartPtr&lt;'_, T&gt;&gt;`

This seems wrong to me, because I've created and then dropped one instance of `SmartPtr`; after the call to `drop`, there should be no more references to `x`, but the compiler seems to think the temporary `Option` value I created still exists, even though, as far as I can tell, destructuring it *must* have destroyed it, because otherwise there would be a second instance of `SmartPtr` floating around, and that's not possible because it doesn't implement `Copy`.

If I rearrange things to use methods of Option instead of destructuring, everything works as you'd expect:

    fn not_broken&lt;T&gt;(x: T) -&gt; T {
        let opt_ptr = Some(SmartPtr(&amp;x));
        if opt_ptr.is_some() {
            let ptr = opt_ptr.unwrap();
            drop(ptr);
            return x;
        }
        unimplemented!()
    }

I also tried naming the temporary like so, but I got the same errors as before:

    fn broken2&lt;T&gt;(x: T) -&gt; T {
        let tmp = Some(SmartPtr(&amp;x));
        if let Some(ptr) = tmp {
            drop(ptr);
            return x;
        }
        unimplemented!()
    }

I compiled all my samples with rustc 1.43.0-nightly (6d69caba1 2020-02-27). Should I file a bug?
## [8][My first library (for printing a progress bar while iterating )](https://www.reddit.com/r/rust/comments/famvrm/my_first_library_for_printing_a_progress_bar/)
- url: https://www.reddit.com/r/rust/comments/famvrm/my_first_library_for_printing_a_progress_bar/
---
Hello fellow Rustacians.

I just wanted to share my first project I hosted on [crates.io](https://crates.io). 

It is a library for printing a progress bar on stdout, while iterating over any Iterator. The main goal is to make it stupid simple to use, just like the python library tqdm, that I use when in python land.

I'm not a professional rust developer, so I have learned a lot about rust and its awesome eco system from this little project already. I think I have finally found my favorite programming language.

I would love to get some feedback on it, as it not only my first real rust library project but my first real library I have written and published. I have also made many mistakes on the way, but I think it's finally usable, so better use the most recent version.

I'd be happy if you checked it out over at: [https://crates.io/crates/prgrs](https://crates.io/crates/prgrs).  Links to the github repo and documentation can also be found over there.

I'm actively developing on it and always trying to improve, so I'm thankful for every feedback I can get.
## [9][new crate for computational camera geometry](https://www.reddit.com/r/rust/comments/falcf4/new_crate_for_computational_camera_geometry/)
- url: https://github.com/strawlab/cam-geom
---

## [10][Stack Overflow Developer Survey (a.k.a. annual "most loved languages" survey) closes *Friday*](https://www.reddit.com/r/rust/comments/fahwgk/stack_overflow_developer_survey_aka_annual_most/)
- url: https://www.reddit.com/r/rust/comments/fahwgk/stack_overflow_developer_survey_aka_annual_most/
---
Consider taking this if you haven't already. Judging by last year's sample size, +20 responses correspond to moving the "loved" ratio for Rust by up to 0.1 percentage points. I know a lot of people around here are looking forward to using Rust in the next year and it would be great for that to be reflected in the data.

- Survey announcement: https://stackoverflow.blog/2020/02/19/new-decade-new-survey-goals-reminder-to-take-the-survey-before-it-closes-next-week/

- Direct link to survey: https://stackoverflow.com/dev-survey/start
## [11][Announcing the first FFI-unwind project design meeting | Inside Rust Blog](https://www.reddit.com/r/rust/comments/faioti/announcing_the_first_ffiunwind_project_design/)
- url: https://blog.rust-lang.org/inside-rust/2020/02/27/ffi-unwind-design-meeting.html
---

## [12][How do stacktraces work in rust?](https://www.reddit.com/r/rust/comments/fan0fk/how_do_stacktraces_work_in_rust/)
- url: https://www.reddit.com/r/rust/comments/fan0fk/how_do_stacktraces_work_in_rust/
---
Because they don't in `C`. The stacktraces are printed even with a release build, though some function calls are omitted because they're optimized out.

What are the capabilities provided by rust runtime? Apparantly it keeps track of the function names and stack state.
