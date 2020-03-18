# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (12/2020)!](https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (12/2020)?](https://www.reddit.com/r/rust/comments/fjefxj/whats_everyone_working_on_this_week_122020/)
- url: https://www.reddit.com/r/rust/comments/fjefxj/whats_everyone_working_on_this_week_122020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-12-2020/39545?u=llogiq)!
## [3][Apple hiring Rust engineers for storage and networking groups](https://www.reddit.com/r/rust/comments/fkngza/apple_hiring_rust_engineers_for_storage_and/)
- url: https://twitter.com/benwilliamson/status/1240113606374686721
---

## [4][Rust Search Extension - The ultimate search extension for Rust](https://www.reddit.com/r/rust/comments/fkmdwg/rust_search_extension_the_ultimate_search/)
- url: https://rust-search-extension.now.sh/
---

## [5][COVID-19 data cleaner, converts CSV in to beautiful typed parquet (powered by Rust)](https://www.reddit.com/r/rust/comments/fkhv5k/covid19_data_cleaner_converts_csv_in_to_beautiful/)
- url: https://github.com/assaydepot/dracula-covid19
---

## [6][docs.rs now allows you to choose your build targets | Rust Blog](https://www.reddit.com/r/rust/comments/fk9gey/docsrs_now_allows_you_to_choose_your_build/)
- url: https://blog.rust-lang.org/2020/03/15/docs-rs-opt-into-fewer-targets.html
---

## [7][Post: parallel-stream](https://www.reddit.com/r/rust/comments/fk74ef/post_parallelstream/)
- url: https://blog.yoshuawuyts.com/parallel-stream/
---

## [8][How can I build a raw web request?](https://www.reddit.com/r/rust/comments/fkn83o/how_can_i_build_a_raw_web_request/)
- url: https://www.reddit.com/r/rust/comments/fkn83o/how_can_i_build_a_raw_web_request/
---
Context: I want to test a server on what ciphers it supports in his TLS configurations, this means  sending an *Hello Client* request on port 443 (this is the standard port at least). After the *Server Hello* response I will just close the connection. I don't need any crypto implementation or fancy stuff.

I tried to give a look to *rustls* but I was not able to figure it out by myself.

Hope someone would help me and if you are at home cause COVID-19 and you want to share your knowledge/experience mentoring me in something cooler I would be excited.

PS: I crave for being able to contribute to some project but I have this initial barrier that stops me every time.
## [9][Converting human-like times into time.Duration](https://www.reddit.com/r/rust/comments/fkmvm9/converting_humanlike_times_into_timeduration/)
- url: https://www.reddit.com/r/rust/comments/fkmvm9/converting_humanlike_times_into_timeduration/
---
Hey, Reddit! Today I want to tell you about the new update `ms_converter` library. Now you can convert time not only into i64, but you can also convert it into `time.Duration`

For that, you need to set library version 0.7.0 or higher:

    [dependencies]
    ms-converter = "0.7"

&amp;#x200B;

And now you can use this construction:

    use crate::ms_converter::ms_into_time;
    
    let value = ms_into_time("1d").unwrap();
    assert_eq!(value.as_millis(), 86400000)

Where `value` will be `time.Duration` instance.

&amp;#x200B;

For more info, you can read Readme on [the repository](https://github.com/Mnwa/ms) or read the doc on [doc.rs](https://docs.rs/ms-converter/).

Good luck!
## [10][Governance Working Group Update: Meeting 12 March 2020 | Inside Rust Blog](https://www.reddit.com/r/rust/comments/fkfa2j/governance_working_group_update_meeting_12_march/)
- url: https://blog.rust-lang.org/inside-rust/2020/03/17/governance-wg.html
---

## [11][Postponing RustFest Netherlands](https://www.reddit.com/r/rust/comments/fk35ir/postponing_rustfest_netherlands/)
- url: https://blog.rustfest.eu/postponing-rustfest-nl
---

## [12][What is lock step application in `ndarray::azip`?](https://www.reddit.com/r/rust/comments/fklrs4/what_is_lock_step_application_in_ndarrayazip/)
- url: https://www.reddit.com/r/rust/comments/fklrs4/what_is_lock_step_application_in_ndarrayazip/
---
The `ndarray` describes that `azip` applies **lock step function application.**

*Array zip macro: lock step function application across several arrays and producers.*

Some googling let me to [lockstep\_computing](https://en.wikipedia.org/wiki/Lockstep_(computing)). This describes fault tolerancy by doing redundant operations and comparing them. To me this doesn't seem to be the case for a simple elementwise operation on arrays. What does ndarray mean by lock step computing?
