# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (41/2020)!](https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (41/2020)?](https://www.reddit.com/r/rust/comments/j57zzs/whats_everyone_working_on_this_week_412020/)
- url: https://www.reddit.com/r/rust/comments/j57zzs/whats_everyone_working_on_this_week_412020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-41-2020/49684?u=llogiq)!
## [3][rust-analyzer changelog #45](https://www.reddit.com/r/rust/comments/j6332p/rustanalyzer_changelog_45/)
- url: https://rust-analyzer.github.io/thisweek/2020/10/05/changelog-45.html
---

## [4][Building a Tide Clock in Rust and Raspberry Pi](https://www.reddit.com/r/rust/comments/j625b5/building_a_tide_clock_in_rust_and_raspberry_pi/)
- url: https://www.reddit.com/r/rust/comments/j625b5/building_a_tide_clock_in_rust_and_raspberry_pi/
---
Story of a heartwarming maker project and first foray into Rust: [https://thefuntastic.com/blog/rust-tide-clock](https://thefuntastic.com/blog/rust-tide-clock)

https://preview.redd.it/7bb93q9l2gr51.jpg?width=3696&amp;format=pjpg&amp;auto=webp&amp;s=9e8a6a413a89bc12be91b88daea10cf7dbf61555
## [5][I wrote a 8-byte implementation of Vec lovingly called "MiniVec"](https://www.reddit.com/r/rust/comments/j5v95j/i_wrote_a_8byte_implementation_of_vec_lovingly/)
- url: https://www.reddit.com/r/rust/comments/j5v95j/i_wrote_a_8byte_implementation_of_vec_lovingly/
---
Crate link:  [https://crates.io/crates/minivec](https://crates.io/crates/minivec) 

Hey all,

Long time C++ programmer giving Rust a whirl and I gotta say, this is a cute and powerful language.

It dramatically simplifies many things that are "wrong" in C++ and gives experienced devs some great agility. Overall, I look forward to using Rust more and more.

Download `MiniVec` and let me know what y'all think!
## [6][Ajour – A GUI application in Rust to manage World of Warcraft addons](https://www.reddit.com/r/rust/comments/j63h4b/ajour_a_gui_application_in_rust_to_manage_world/)
- url: https://github.com/casperstorm/ajour
---

## [7][Results of Authoring a JS Library with Rust and Wasm](https://www.reddit.com/r/rust/comments/j62mqp/results_of_authoring_a_js_library_with_rust_and/)
- url: https://nickb.dev/blog/results-of-authoring-a-js-library-with-rust-and-wasm
---

## [8][Which rust libraries do you use to protect your web api?](https://www.reddit.com/r/rust/comments/j63mm5/which_rust_libraries_do_you_use_to_protect_your/)
- url: https://www.reddit.com/r/rust/comments/j63mm5/which_rust_libraries_do_you_use_to_protect_your/
---
My prototype uses a auth microservice which connects to a postgres db. If username&amp;password is correct, it returns a session and access token and stores the session id in the db. Each other microservice validates the access token.

Session token is valid for 7 days and refreshed for every access token refresh, access token is valid for 15 minutes. The session stored in the db let's user identify and blacklist sessions (e.g. Web Login, Firefox, 5 days ago). This works fine (unless one user decides to spam the login endpoint and create endless sessions, but that can be solved with delete all but the 10 most recent).

Now the mantra seems to be: Don't implement it yourself. Is there anything easy like the approach above I can use from preexisting libs? 

&amp;#x200B;

Bonus Question: Let's say a user got access to feature 1-100 except xyz. If I put all of that in my access token it becomes quite huge. But asking another service if user a can do b kind of defeats the "do not hit db until refresh" idea.
## [9][IntelliJ Rust Changelog #132 - make rustfmt your default IDE formatter](https://www.reddit.com/r/rust/comments/j5lv3o/intellij_rust_changelog_132_make_rustfmt_your/)
- url: https://intellij-rust.github.io/2020/10/05/changelog-132.html
---

## [10][Cross compilation instructions of Gtk-rs app(Czkawka) from Ubuntu to Windows for Github CI](https://www.reddit.com/r/rust/comments/j5z48c/cross_compilation_instructions_of_gtkrs/)
- url: https://github.com/qarmin/Instrukcje-i-Tutoriale/blob/master/GtkRsCross.md
---

## [11][Campaign to open-source Sciter GUI Engine (with Rust bindings provided by sciter-rs)](https://www.reddit.com/r/rust/comments/j63sro/campaign_to_opensource_sciter_gui_engine_with/)
- url: https://www.kickstarter.com/projects/c-smile/open-source-sciter-engine/
---

## [12][What are some good projects to learn concurrent programming?](https://www.reddit.com/r/rust/comments/j61fvq/what_are_some_good_projects_to_learn_concurrent/)
- url: https://www.reddit.com/r/rust/comments/j61fvq/what_are_some_good_projects_to_learn_concurrent/
---
I've been programming in different languages for about 3 years now and one thing I have always avoided experimenting with was concurrent programming, considering how unsafe and buggy it is in most languages. 

But now that I have learnt Rust (still a beginner, but I've been able to build some pet projects), I want to get into concurrent/parallel programming. What are some good simple projects that I could try doing? I have spent quite some time looking for ideas but can't seem to find anything that would really benefit from concurrency. 

P.S. I have read Rust documentation about the matter already and checked the suggested "final project", but I don't really want to make a website I have no use for.
