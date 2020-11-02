# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (45/2020)!](https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (45/2020)?](https://www.reddit.com/r/rust/comments/jmiktw/whats_everyone_working_on_this_week_452020/)
- url: https://www.reddit.com/r/rust/comments/jmiktw/whats_everyone_working_on_this_week_452020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-45-2020/50962?u=llogiq)!
## [3][Nom 6 released](https://www.reddit.com/r/rust/comments/jmiyg8/nom_6_released/)
- url: https://github.com/Geal/nom/blob/master/CHANGELOG.md#600---2020-10-31
---

## [4][nom parser combinators, version 6.0.0](https://www.reddit.com/r/rust/comments/jmjgug/nom_parser_combinators_version_600/)
- url: https://www.reddit.com/r/rust/comments/jmjgug/nom_parser_combinators_version_600/
---
Hello all!  


I'm happy to tell you that the [nom parser combinators library](https://github.com/Geal/nom) is now out at version 6.0.0! many thanks to [all the contributors](https://github.com/Geal/nom/blob/master/CHANGELOG.md#thanks-1) who helped developing it!  


This release addresses a lot of issues around usability, mainly by relaxing the requirements on parsers (they can be \`FnMut\` now, or even anything that implements the \`nom::Parser\` trait). There was also some work around error management, simplifying the design of custom error types, integration with \`std::error::Error\`, etc.  


This release also integrates with the awesome [bitvec](https://crates.io/crates/bitvec) crate, for easier bit level parsing. Along with the [regex](https://crates.io/crates/regex) and [lexical-core](https://crates.io/crates/lexical-core), it shows that you can easily wrap more specialized library in parsers that will be combined with the rest of nom.  


There are a few [breaking changes](https://github.com/Geal/nom/blob/master/CHANGELOG.md#breaking-changes), but so far upgrades were not too complex changing \`Fn\` to \`FnMut\` here and there, and implementing the \`ContextError\` and \`FromExternalError\` trait if you use a custom error type.  


There might be a few performance gains, but this was not the real focus of this release, I may have good news on that front in a few days/weeks.  


That's all! Happy hacking!
## [5][Rust cheats](https://www.reddit.com/r/rust/comments/jmku1s/rust_cheats/)
- url: https://www.reddit.com/r/rust/comments/jmku1s/rust_cheats/
---
In case you didn't found it already, I have found a very useful cheat sheet for learning Rust.   
[Cheats](https://cheats.rs/)

I'm not connected with the resource, I have just found it interesting and useful.
## [6][The Newtype Pattern in Rust](https://www.reddit.com/r/rust/comments/jmkncv/the_newtype_pattern_in_rust/)
- url: https://www.worthe-it.co.za/blog/2020-10-31-newtype-pattern-in-rust.html
---

## [7][Names are not type safety](https://www.reddit.com/r/rust/comments/jmi0lo/names_are_not_type_safety/)
- url: https://lexi-lambda.github.io/blog/2020/11/01/names-are-not-type-safety/
---

## [8][Are We Rust Yet](https://www.reddit.com/r/rust/comments/jm6v0o/are_we_rust_yet/)
- url: https://github.com/UgurcanAkkok/AreWeRustYet
---

## [9][Precursor: secure open hardware platform for Rust apps](https://www.reddit.com/r/rust/comments/jmh4rv/precursor_secure_open_hardware_platform_for_rust/)
- url: https://www.crowdsupply.com/sutajio-kosagi/precursor
---

## [10][hifitime version 2.1: a scientific time management library with several times systems (UTC, TAI, ...) , duration handling, time units, and build on top of lossless fractions](https://www.reddit.com/r/rust/comments/jmj1nz/hifitime_version_21_a_scientific_time_management/)
- url: https://docs.rs/hifitime/2.1.0/hifitime/#examples
---

## [11][A Continuous Deployment Pipeline For Rust Applications - Zero To Production In Rust #5](https://www.reddit.com/r/rust/comments/jmbxbv/a_continuous_deployment_pipeline_for_rust/)
- url: https://www.lpalmieri.com/posts/2020-11-01-zero-to-production-5-how-to-deploy-a-rust-application/
---

## [12][Compile-Time Reflection in Rust](https://www.reddit.com/r/rust/comments/jm1h2x/compiletime_reflection_in_rust/)
- url: https://www.mn.uio.no/ifi/english/research/groups/psy/completedmasters/2020/gaarde/masterthesis-gaarde.pdf
---

