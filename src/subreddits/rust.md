# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (18/2020)!](https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g4nu6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (18/2020)?](https://www.reddit.com/r/rust/comments/g9a7d1/whats_everyone_working_on_this_week_182020/)
- url: https://www.reddit.com/r/rust/comments/g9a7d1/whats_everyone_working_on_this_week_182020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-18-2020/41671?u=llogiq)!
## [3][Why `dirs` and `directories` repositories have been archived?](https://www.reddit.com/r/rust/comments/ga7f56/why_dirs_and_directories_repositories_have_been/)
- url: https://www.reddit.com/r/rust/comments/ga7f56/why_dirs_and_directories_repositories_have_been/
---
[`dirs`](https://github.com/soc/dirs-rs) and [`directories`](https://github.com/soc/directories-rs) have more than 5 million downloads total, and more than 1.5 million downloads recently combined.

What happened? Why were they archived? Is it temporary, or they won't be developed anymore?
## [4][Native Windows GUI 1.0 prerelease](https://www.reddit.com/r/rust/comments/ga024n/native_windows_gui_10_prerelease/)
- url: https://github.com/gabdube/native-windows-gui/tree/1.0-prerelease
---

## [5][From 0 to Faster Than Python in 4 days](https://www.reddit.com/r/rust/comments/ga10tm/from_0_to_faster_than_python_in_4_days/)
- url: https://www.reddit.com/r/rust/comments/ga10tm/from_0_to_faster_than_python_in_4_days/
---
Last Friday, I had enough of how slow my simulation, written in Python, was running. Cython hadn't been good at speeding it up, either, and transforming it into Cython had meant giving up my meticulous type annotations with `Dict[HexIndex, List[Family]]` and my `dataclass` decorators and my generators for some imperceptible speed gains.

So I had a look around. Writing that same thing in C or C++ and having to track pointers and memory? The horror. Haskell? Maybe, what else is there?

Somehow the idea of Rust dropped into my lap. I had seen an impressive demo by /u/theanzelm of /r/citybound at some point where he demonstrates his parallel (etc., but that was what I cared about) library, and the language hat popped up elsewhere in my periphery recently, so I had a look.

And it turns out Rust was exactly what I needed. In the last 10 minutes, while I type this, I have simulated more time steps than in several days during the last week in Python. I have all the types (including, of course, a `Hashmap&lt;H3Index, Vec&lt;&amp;Family&gt;&gt;`, my classes have become nice minimal structs, nothing minds that the order of definitions is in the logical order suggested by the ODD protocol instead of later functions desperately requiring definitions to come before them, and as a bonus I can see which part of my simulation changes which elements.

So. I'm a convert.

But now I need to get my code readable. Ideally, I want to get it to a state where my non-Rustacean colleagues, collaborators, and referees in academia can read the code and point out where the simulation deviates from their concept of the process modeled. But before that, it would be nice if at least fellow Rust programmers were able to understand what's going on, and lead me towards writing better, more idiomatic Rust.

Once I have transferred all my comments from the Python version to the Rust version and written some more tests, where can I get some 800 lines (before comments) of Rust code-reviewed?
## [6][[Example] GitHub Actions with Rust (build, test, coverage, doc, clippy, rustfmt, publish)](https://www.reddit.com/r/rust/comments/ga80lj/example_github_actions_with_rust_build_test/)
- url: https://github.com/TimDiekmann/alloc-compose/tree/master/.github/workflows
---

## [7][Rust and Tell Berlin Meetup - April 2020 [video]](https://www.reddit.com/r/rust/comments/ga5dnk/rust_and_tell_berlin_meetup_april_2020_video/)
- url: https://www.youtube.com/watch?v=yGuxtodWYDs
---

## [8][I made a video on concurrency in Rust with Async/Await](https://www.reddit.com/r/rust/comments/ga6m2d/i_made_a_video_on_concurrency_in_rust_with/)
- url: https://www.reddit.com/r/rust/comments/ga6m2d/i_made_a_video_on_concurrency_in_rust_with/
---
Hello Everyone, I made a video on concurrency in Rust using Async/Await with some real-time examples. I hope you guys like it. I don't have a good audio gear so audio is not so good and I request watch it at 2.X;

[Video Link](https://www.youtube.com/watch?v=hrNoTZMG2MU)
## [9][A simple (and incomplete) implementation of an async Sonic client.](https://www.reddit.com/r/rust/comments/ga78pg/a_simple_and_incomplete_implementation_of_an/)
- url: https://www.reddit.com/r/rust/comments/ga78pg/a_simple_and_incomplete_implementation_of_an/
---
Hi everyone,

&amp;#x200B;

I was trying to use the [Sonic](https://github.com/valeriansaliou/sonic) project and get stuck with the Rust community client, so I decided to start building one using [Tokio](https://tokio.rs).

&amp;#x200B;

Here is the link to the repo: [https://github.com/ceciliacsilva/my\_sonic\_client](https://github.com/ceciliacsilva/my_sonic_client)

&amp;#x200B;

Hopefully it can be helpful for somebody else.
## [10][[Side Project] Help wanted for reviewing code of RustCred](https://www.reddit.com/r/rust/comments/ga8mmx/side_project_help_wanted_for_reviewing_code_of/)
- url: https://www.reddit.com/r/rust/comments/ga8mmx/side_project_help_wanted_for_reviewing_code_of/
---
I have just started learning Rust. I am doing about an hour, or sometimes two, per day and sometime around the fourth week I got tired of reading books, blogs and articles. So I decided to build something!

This past week have been spent building [RustCred (https://github.com/saidaspen/rustcred)](https://github.com/saidaspen/rustcred)! 

It's a small app which generates the HTML for [https://rustcred.dev](https://rustcred.dev/), which is in its infancy at the moment :) A small page to gamify contributions to Rust Open source projects.

I am pretty happy that I could get it to this state. However, since I am all alone, and I don't know a living soul who knows Rust, I have no idea if my code is alright or not. 

Therefore I am wondering if there are any kind souls here, who would like to have a look at the code and give me feedback? It is about 400 lines of code in total.

thanks, and stay safe!
## [11][Is Rust fun?](https://www.reddit.com/r/rust/comments/g9rsaw/is_rust_fun/)
- url: https://www.reddit.com/r/rust/comments/g9rsaw/is_rust_fun/
---
I am a cs student and i really like c and c++. I love to code simple physics stuff and try to get the best performance out of it for fun. In c++ you get a lot of controll and i really love that, but some times is just hate the old and bad parts of c++ that are caused by its long live and backwards compatibility. So my question are:

For my hobby programming, could rust be an upgrade in programming fun for me for my usecases(pathtracing, physics engine)?

Are the compile times better?

How complicated is the build system?

Does rust lack some nice c++ features for example constexpr or templates?

Do i get the same level od memory controll in rust as in c++?

Is rust as fast as c++ or near it in termes of execution time?

Is the syntax generally "better" (in your opinion)?
Edit: a question i forgott: does rust have in language serialisation support?
## [12][sonor: a crate for controlling sonos speakers](https://www.reddit.com/r/rust/comments/g9nh3b/sonor_a_crate_for_controlling_sonos_speakers/)
- url: https://www.reddit.com/r/rust/comments/g9nh3b/sonor_a_crate_for_controlling_sonos_speakers/
---
crates.io: [https://crates.io/crates/sonor](https://crates.io/crates/sonor)

Documentation: [https://docs.rs/sonor/](https://docs.rs/sonor/)

Example usage:

    let speaker = sonor::find("your room name", Duration::from_secs(2)).await?
        .expect("room exists");
    
    println!("The volume is currently at {}", speaker.volume().await?);
    
    match speaker.track().await? {
        Some(track_info) =&gt; println!("- Currently playing '{}", track_info.track()),
        None =&gt; println!("- No track currently playing"),
    }
    
    speaker.clear_queue().await?;
    
    speaker.join("some other room").await?;

You can also take snapshots of the current state (whats playing, volume) and restore that later:

    let snapshot = speaker.snapshot().await?;

    speaker.set_volume(10).await?;
    speaker.set_transport_uri("http://e.g.texttospeech.api", "").await?;
    speaker.play().await?;
    tokio::time::delay_for(Duration::from_secs(3)).await;

    speaker.apply(snapshot).await?;
