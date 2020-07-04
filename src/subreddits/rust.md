# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (27/2020)!](https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 345](https://www.reddit.com/r/rust/comments/hisn3e/this_week_in_rust_345/)
- url: https://this-week-in-rust.org/blog/2020/06/30/this-week-in-rust-345/
---

## [3][so: A terminal interface for StackOverflow](https://www.reddit.com/r/rust/comments/hktgsu/so_a_terminal_interface_for_stackoverflow/)
- url: https://github.com/samtay/so
---

## [4][How do you avoid getting google results about the game instead of the language?](https://www.reddit.com/r/rust/comments/hkueg9/how_do_you_avoid_getting_google_results_about_the/)
- url: https://www.reddit.com/r/rust/comments/hkueg9/how_do_you_avoid_getting_google_results_about_the/
---
I'm (very) new to rust. My problem is how do you avoid ending up on a search page about the game instead of the language? 

e.g: when I type "rust \[something here\]", I get results about the game. I tried typing "rustlang etc.." but it doesn't always help. 

Pardon my blatant ignorance.
## [5][autodiscovery-rs: library to help discover other instance on the net, and connect to them](https://www.reddit.com/r/rust/comments/hkzh50/autodiscoveryrs_library_to_help_discover_other/)
- url: https://www.reddit.com/r/rust/comments/hkzh50/autodiscoveryrs_library_to_help_discover_other/
---
Hi all,

I uploaded a new library to [Github](https://github.com/over-codes/autodiscover-rs) today, and would love some feedback. Short summary from the README:

autodiscover-rs implements a simple algorithm to detect peers on an IP network, connects to them, and calls back with the connected stream.  The algorthm supports both UDP broadcast and multicasting.

It can be found at [https://github.com/over-codes/autodiscover-rs](https://github.com/over-codes/autodiscover-rs)
## [6][log only info and error to terminal, everything else to file](https://www.reddit.com/r/rust/comments/hl1kzb/log_only_info_and_error_to_terminal_everything/)
- url: https://www.reddit.com/r/rust/comments/hl1kzb/log_only_info_and_error_to_terminal_everything/
---
I'm using simplelog but don't care which logger to use at this point.

```
    CombinedLogger::init(vec![
        TermLogger::new(LevelFilter::Info, Config::default(), TerminalMode::Mixed).unwrap(),
        WriteLogger::new(
            LevelFilter::Debug,
            Config::default(),
            std::fs::File::create("/tmp/dfclient.log").unwrap(),
        ),
    ])
    .unwrap();
```

This is wrong since it logs warnings to the terminal as well... which is kinda annoying in a network / p2p app with disconnects and so on happening often. How can I get the warnings not to appear on the terminal? About to just write my own logger at this point.
## [7][RustPython: Building a Python 3 interpreter in Rust (FOSDEM 2019)](https://www.reddit.com/r/rust/comments/hkjpgx/rustpython_building_a_python_3_interpreter_in/)
- url: https://www.youtube.com/watch?v=nJDY9ASuiLc
---

## [8][Graceful keyboard shutdown of thread pool in Rust](https://www.reddit.com/r/rust/comments/hl2kdm/graceful_keyboard_shutdown_of_thread_pool_in_rust/)
- url: https://killavus.github.io/posts/thread-pool-graceful-shutdown/
---

## [9][What are the benefits of using Rust async HTTP clients?](https://www.reddit.com/r/rust/comments/hkw74o/what_are_the_benefits_of_using_rust_async_http/)
- url: https://www.reddit.com/r/rust/comments/hkw74o/what_are_the_benefits_of_using_rust_async_http/
---
Generally it seems like async only has a real benefit when there are many connections.

Is there a scenario where using an HTTP client implemented with Rust's async features gives a substantive performance benefit?
## [10][Nick Cameron: Early Impressions of Go from a Rust Programmer](https://www.reddit.com/r/rust/comments/hkjs96/nick_cameron_early_impressions_of_go_from_a_rust/)
- url: https://pingcap.com/blog/early-impressions-of-go-from-a-rust-programmer
---

## [11][RustyPipe - A youtube extractor and frontend made in Rust](https://www.reddit.com/r/rust/comments/hkr9bx/rustypipe_a_youtube_extractor_and_frontend_made/)
- url: https://www.reddit.com/r/rust/comments/hkr9bx/rustypipe_a_youtube_extractor_and_frontend_made/
---
Website Link: [https://rustypipe.deepraven.co](https://rustypipe.deepraven.co)

Linux AppImage Link: [https://github.com/deep-gaurav/rusty\_pipe\_front/releases/tag/20200703205931-4819d47](https://github.com/deep-gaurav/rusty_pipe_front/releases/tag/20200703205931-4819d47)

&amp;#x200B;

Some Story,

So as a second project in **learning rust** I decided to rewrite [NewPipeExtractor](https://github.com/TeamNewPipe/NewPipeExtractor) in rust.I rewrote the youtube extraction part it's available at [https://github.com/deep-gaurav/rusty\_pipe](https://github.com/deep-gaurav/rusty_pipe).

My initial plans were to make a graphql api of rustypipe and use it to make website/app like [invidio.us](https://invidio.us).

So I made a graphql server using juniper avaiable at [rustypipe.herokuapp.com](https://rustypipe.herokuapp.com) ([https://github.com/deep-gaurav/rusty\_pipe\_server](https://github.com/deep-gaurav/rusty_pipe_server)).

When making frontend website to server I decided to use [yew.rs](https://yew.rs) but when implementing graphql client, i realised that i could use extracter in browser itself instead of relying on the graphql server which was slow and added quite burden to server, moreover heroku minutes were limited.But Cors was problem so i implemented a simple cors proxy using vercel serverless function and deployed the site on vercel.

So the current version does all the extraction and parsing on client side and only relies on vercel serverless function as a cors proxy.

This also led me to make a desktop version by wrapping it in cordova, since there is no cors in cordova/electron, it does not relies on any external server (Except youtube ofcourse) and behaves pretty much like newpipe for desktop.

I also made a custom video element to be used with app in typescript.It's highly work in progress, would love any sort of feedback/criticism on it.

P.S: If on linux consider using desktop release, it's faster and should be able to play even encrypted videos, like songs from official channels which web version cannot ( just like [invidio.us](https://invidio.us) ) as url is ip bound (and ip is of cors proxy).

If on chrome desktop it'll also show a Picture in Picture mode button, and on chrome android it'll use mediasession api to customize notification
## [12][Canrun: A logic programming library inspired by the *Kanren family](https://www.reddit.com/r/rust/comments/hkp1hi/canrun_a_logic_programming_library_inspired_by/)
- url: https://esimmler.com/announcing-canrun/
---

