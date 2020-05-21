# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (21/2020)!](https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/glvkc5/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 339](https://www.reddit.com/r/rust/comments/gmyv8h/this_week_in_rust_339/)
- url: https://this-week-in-rust.org/blog/2020/05/19/this-week-in-rust-339/
---

## [3][Support cargo.toml for GitHub dependency graph](https://www.reddit.com/r/rust/comments/gnryht/support_cargotoml_for_github_dependency_graph/)
- url: https://github.community/t5/How-to-use-Git-and-GitHub/Feature-request-Support-Cargo-toml-for-dependency-graph/m-p/26122
---

## [4][Dropping heavy objects in another thread can make your code faster](https://www.reddit.com/r/rust/comments/gntv7l/dropping_heavy_objects_in_another_thread_can_make/)
- url: https://abramov.io/rust-dropping-things-in-another-thread
---

## [5][Announcement: Electron-like library for rust that does not bundle chrome](https://www.reddit.com/r/rust/comments/gnvkf0/announcement_electronlike_library_for_rust_that/)
- url: https://www.reddit.com/r/rust/comments/gnvkf0/announcement_electronlike_library_for_rust_that/
---
Using electron, desktop apps can be created for Mac, Windows and Linux easily using JS and HTML. But in electron the chrome browser is bundled so each app has it's own copy of chrome. Eventhough webview exists, in Windows it uses Internet Explorer so modern JS,HTML or CSS cannot be used. Also different browsers have to be supported.

Webview can be used to create desktop apps but I came across a go library called [lorca](https://github.com/zserge/lorca).It uses the existing chrome installation and communicates with it using the devtools protocol. So I created a lorca like library for rust called [alcro](https://github.com/Srinivasa314/alcro).

I plan to publish it to crates.io but before that I wanted it tested on windows and macos. Alcro will search for the chrome installation and display a message box if not found. The messagebox asks the user whether they want to download chrome and opens the webpage to download chrome if the user agrees.I want some users to check if the following work

* Whether it compiles on win/mac
* Whether it finds the chrome installation
* If you comment out the path of your chrome installation in `src/locate.rs` does it display a messagebox
* If you click yes does a webpage open to download chrome
* Do documentation tests work
* Do the examples work
* Read the documentation ( using cargo doc --open as of now ) and see if all the features work
## [6][Oxidised eBPF, part I.: Building a toolchain](https://www.reddit.com/r/rust/comments/gnu5td/oxidised_ebpf_part_i_building_a_toolchain/)
- url: https://blog.redsift.com/labs/oxidised-ebpf-part-i-building-a-toolchain/
---

## [7][Things I hate about Rust](https://www.reddit.com/r/rust/comments/gnd4bd/things_i_hate_about_rust/)
- url: https://blog.yossarian.net/2020/05/20/Things-I-hate-about-rust
---

## [8][Is actix-web still a good bet for a new production website?](https://www.reddit.com/r/rust/comments/gnm0em/is_actixweb_still_a_good_bet_for_a_new_production/)
- url: https://www.reddit.com/r/rust/comments/gnm0em/is_actixweb_still_a_good_bet_for_a_new_production/
---
With all the fuss about the main developer leaving and the repository getting deleted for a bit, is it still likely that actix-web will see continued development? If I'm making a new production website for a business, should I have reservations about using it?
## [9][Oxidizing the technical interview](https://www.reddit.com/r/rust/comments/gnbblm/oxidizing_the_technical_interview/)
- url: https://blog.mgattozzi.dev/oxidizing-the-technical-interview/
---

## [10][Create Cookies with Warp](https://www.reddit.com/r/rust/comments/gnui36/create_cookies_with_warp/)
- url: https://www.reddit.com/r/rust/comments/gnui36/create_cookies_with_warp/
---
I want to create cookies on my server using \`warp\`. Now in documentation I could only find the \`cookie\` filter which lets me access a cookie but not create one.

No any other documentation I could find related to cookies in the documentation.

&amp;#x200B;

Any help would be really appreciated.
## [11][*mut Type vs &amp;mut Type . What's the difference between them?](https://www.reddit.com/r/rust/comments/gnr9th/mut_type_vs_mut_type_whats_the_difference_between/)
- url: https://www.reddit.com/r/rust/comments/gnr9th/mut_type_vs_mut_type_whats_the_difference_between/
---

## [12][3d6: Rendering electron orbitals of hydrogen](https://www.reddit.com/r/rust/comments/gnhkxn/3d6_rendering_electron_orbitals_of_hydrogen/)
- url: https://github.com/cbeuw/iiiD6
---

