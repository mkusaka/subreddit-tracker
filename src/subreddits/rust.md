# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (37/2020)!](https://www.reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 355](https://www.reddit.com/r/rust/comments/ippv0q/this_week_in_rust_355/)
- url: https://this-week-in-rust.org/blog/2020/09/09/this-week-in-rust-355/
---

## [3][Rustbot, are you okay?](https://www.reddit.com/r/rust/comments/ipzctt/rustbot_are_you_okay/)
- url: https://i.redd.it/sjul1r9q5am51.jpg
---

## [4][rust-steward](https://www.reddit.com/r/rust/comments/iq17rw/ruststeward/)
- url: https://www.reddit.com/r/rust/comments/iq17rw/ruststeward/
---
Hi, 

Scala has a wonderful tool, [scala-steward](https://github.com/scala-steward-org/scala-steward). It checks a project's dependencies and creates pull requests when updates are available.

Similarly, GitHub has a bot that checks for security issues.

Would anyone be aware of a similar tool for Rust?
## [5][Meuse, a free private crate registry, 1.0.0 release](https://www.reddit.com/r/rust/comments/ipmkqe/meuse_a_free_private_crate_registry_100_release/)
- url: https://mcorbin.fr/posts/2020-09-09-meuse-1.0.0/
---

## [6][Opinions on Rust future](https://www.reddit.com/r/rust/comments/iq1mde/opinions_on_rust_future/)
- url: https://www.reddit.com/r/rust/comments/iq1mde/opinions_on_rust_future/
---
I come from a Java/C/Python background and I'm interested in learning a new programming language. I've always heard about Rust but I started to get interested only recently when I read about its concurrency and security. I know only a little about it but my main fear is that I start learning Rust and then the language won't be used anymore in 5-10 years as it's used today (basically I'm afraid to learn a language that will teach me things but won't give me job opportunities). So my question is: I'm interested in Rust potentialities, but I'm also interested in the long-term use of the language... I know no one can see the future, but has Rust the capability to compete against other famous programming languages (like Javascript, Python, Java, C#,... or the emerging Go, Kotlin,...) or it's a "niche" language and it will only be used in specific areas?
## [7][Make A Language: Part Three](https://www.reddit.com/r/rust/comments/iq2gaa/make_a_language_part_three/)
- url: https://arzg.github.io/lang/3/
---

## [8][Are there cyptolibs written in rust or bindings to other cryptolibs which you consider production ready ?](https://www.reddit.com/r/rust/comments/iq17z0/are_there_cyptolibs_written_in_rust_or_bindings/)
- url: https://www.reddit.com/r/rust/comments/iq17z0/are_there_cyptolibs_written_in_rust_or_bindings/
---
I want to write a simple backup tool in rust which basically should just combine some files and a mariadb sql dump into a tar ball and then encrypt that tarball using a key derived from a password.

With PHP I would use the built-in bindings to libsodium. 

Rust also has bindings for libsodium through the [sodiumoxide](https://github.com/sodiumoxide/sodiumoxide) project but the version of the latest release (0.2.6) makes me hesitant to use it in production. 

So would you use sodiumoxide in production and if not is there a cryptolib written in rust or some bindings to another well known cryptolib that you can recommend ?
## [9][** Rust &amp;&amp; C++ London Joint Meetup**](https://www.reddit.com/r/rust/comments/ipf4sp/rust_c_london_joint_meetup/)
- url: https://www.reddit.com/r/rust/comments/ipf4sp/rust_c_london_joint_meetup/
---
Rust London and C++ London are proud to announce our first joint meetup The Rust &amp;&amp; C++ LDN Talks. This is the first in a series of collaborative events between our two communities which have been compared and contrasted for several years.  


The purpose of these events is to begin an exchange of ideas and experiences between our user groups.  
The Rust London User Group will open up the LDN Talks platform to the C++ London community. We will have speakers from C++ London who will share their perspectives, opinions and learning experiences with the Rust Programming language.

&amp;#x200B;

Check out the Meetup Page below

[https://www.meetup.com/Rust-London-User-Group/events/273056379/](https://www.meetup.com/Rust-London-User-Group/events/273056379/)
## [10][implement trait with different path](https://www.reddit.com/r/rust/comments/iq0be2/implement_trait_with_different_path/)
- url: https://www.reddit.com/r/rust/comments/iq0be2/implement_trait_with_different_path/
---
I try to implement a trait 'path::Example' for a type T. I only succeed to implement the trait 'other::Example' for T. The trait 'Example' is defined once in the file "other[.rs](https://good.rs)". I don't understand how to proceed and why the lib is searching the trait in "path" rather than "other" .

Any suggestions? Thks
## [11][Today I learned how to write procedural macros! My first one writes long enum expressions for tuple enums with differently-typed variants but which the same code would work on. enum_for_matches::run!(e, {TestEnum::I64(i) | TestEnum::U64(i)}, {s = i.to_string();}); Anyone mind reviewing my code?](https://www.reddit.com/r/rust/comments/iq1kuq/today_i_learned_how_to_write_procedural_macros_my/)
- url: https://crates.io/crates/enum_for_matches
---

## [12][[knurling] Learning Embedded Rust with Knurling-rs](https://www.reddit.com/r/rust/comments/ipg76a/knurling_learning_embedded_rust_with_knurlingrs/)
- url: https://ferrous-systems.com/blog/knurling-sessions-introduction/
---

