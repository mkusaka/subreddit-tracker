# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (40/2020)!](https://www.reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (40/2020)?](https://www.reddit.com/r/rust/comments/j1jime/whats_everyone_working_on_this_week_402020/)
- url: https://www.reddit.com/r/rust/comments/j1jime/whats_everyone_working_on_this_week_402020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-40-2020/49429?u=llogiq)!
## [3][Rust 2021: GUI](https://www.reddit.com/r/rust/comments/j24em4/rust_2021_gui/)
- url: https://raphlinus.github.io/rust/druid/2020/09/28/rust-2021.html
---

## [4][Revisiting a 'smaller Rust'](https://www.reddit.com/r/rust/comments/j2l9v9/revisiting_a_smaller_rust/)
- url: https://without.boats/blog/revisiting-a-smaller-rust/
---

## [5][Rust program runs faster on Linux than Windows?](https://www.reddit.com/r/rust/comments/j2ilkn/rust_program_runs_faster_on_linux_than_windows/)
- url: https://www.reddit.com/r/rust/comments/j2ilkn/rust_program_runs_faster_on_linux_than_windows/
---
I have written a program in Rust which is a physical simulation of a quantum field. It basically means applying a simple algorithm to the same vector milions of times. [Here](https://pastebin.com/5Y0N1306) is the code: the "meat" of the program is the `metropolis` function which, as you can see, runs a 10^(8) long for loop and keeps modifying the `fields.phi` vector. It also writes the result of some averages to a file every `naccu` times.   


Now, here's the questions:

1. Is it normal that this program runs WAY faster on Linux (I'm using [Solus](https://getsol.us/home/)) on a crappy 7+ years old laptop than on Windows 10 on my fairly recent and powerful main PC? On Linux the program halts in about 40 minutes, while on Windows it took 2+ hours. On both OS's I've run the code in release mode and without other flags or optimizations. 
2. Is there a way to optimize this code to run faster? I'm a newbie Rust programer and I already had this code written in C: I basically tried to port it to Rust and see if it runs any faster. The difference is quite small, but the C version is faster (it takes about 10 minutes less on Linux). But, as far as I know, Rust's performance should be equal or even better than C's. Have I written the code in a bad, unoptimized way? Can you suggest ways to improve the code, even if it doesn't really improve performance? Probably I haven't written a very "Rusty" code.
## [6][Fast Electron App with rust](https://www.reddit.com/r/rust/comments/j2g6ry/fast_electron_app_with_rust/)
- url: https://www.reddit.com/r/rust/comments/j2g6ry/fast_electron_app_with_rust/
---
[https://blog.logrocket.com/supercharge-your-electron-apps-with-rust/](https://blog.logrocket.com/supercharge-your-electron-apps-with-rust/)
## [7][Benchmarking vol. 2: Pitting Actix against Rocket v0.4 and v0.5-dev](https://www.reddit.com/r/rust/comments/j2acxl/benchmarking_vol_2_pitting_actix_against_rocket/)
- url: https://matej.laitl.cz/bench-actix-rocket/
---

## [8][My frist project in Rust generates mazes and solves them. I'm sure I've missed a lot of Rust patterns and would love some comments on my code (MIC)](https://www.reddit.com/r/rust/comments/j1xzvu/my_frist_project_in_rust_generates_mazes_and/)
- url: https://i.imgur.com/Y8WPfNk.gifv
---

## [9][The first C++ &amp;Rust LDN *Virtual* Talks tonight](https://www.reddit.com/r/rust/comments/j2hpd4/the_first_c_rust_ldn_virtual_talks_tonight/)
- url: https://www.meetup.com/Rust-London-User-Group/events/273056379/
---

## [10][What would be a good way to have a "plugin" or "extension" system? (more info in body)](https://www.reddit.com/r/rust/comments/j2f549/what_would_be_a_good_way_to_have_a_plugin_or/)
- url: https://www.reddit.com/r/rust/comments/j2f549/what_would_be_a_good_way_to_have_a_plugin_or/
---
Im currently making something similar to nodejs, called [Novel((.)js)](https://github.com/novel-js/runtime) ^(contributions welcome), and one of my ideas is to let packages link with some rust  code to get expanded functionality, for one example, call into an http library, since I haven't implemented that at all yet.

&amp;#x200B;

While sure, they could just push to upstream, but that could be a potentially lengthy process, that might not be very user friendly, So my idea is to have a `config.toml` file, that specifies which plugins to use, and where to fetch them from.

&amp;#x200B;

My current problem, Is that I don't know where to start at all, I don't know whether shared objects is the best way (I heard that they are unreliable when passing complex types on Macos), and other than shared objects, im honestly not sure what other options there are?. Any help on this question, or the project in general are appreciated. (Even if its just a pull request on the Readme to expand plans).

&amp;#x200B;

If you need any new info, comment and ill reply to it.
## [11][Progress report on rustc_codegen_cranelift (Sep 2020)](https://www.reddit.com/r/rust/comments/j20ml4/progress_report_on_rustc_codegen_cranelift_sep/)
- url: https://bjorn3.github.io/2020/09/28/progress-report-sep-2020.html
---

## [12][Panic on unused results](https://www.reddit.com/r/rust/comments/j2i5qy/panic_on_unused_results/)
- url: https://www.reddit.com/r/rust/comments/j2i5qy/panic_on_unused_results/
---
I found a bug in a project today (and multiple bugs in dependencies) due to  unhandled Result.   


The way I found it was by modifying \`libcore\` such that when a \`Result\` is dropped, the program \`panic!\`s if the result contains an error (which means the error is unhandled).

&amp;#x200B;

Why doesn't this happen by default? I'm starting to go through the library, and there are other cases where similar things don't happen (e.g. when a File descriptor is dropped without calling \`close\`), etc.

&amp;#x200B;

I'll maintain my own patched library version, but I wonder why this doesn't happen at least in debug builds or behind some flag.
