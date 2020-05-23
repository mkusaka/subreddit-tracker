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

## [3][I did my first project in Rust! A "photo mode" for Yakuza Kiwami 2!](https://www.reddit.com/r/rust/comments/gp1que/i_did_my_first_project_in_rust_a_photo_mode_for/)
- url: https://v.redd.it/q0o5fikhbh051
---

## [4][Daimler internal source code leaked: Some Rust included :-o](https://www.reddit.com/r/rust/comments/gp2uey/daimler_internal_source_code_leaked_some_rust/)
- url: https://www.reddit.com/r/rust/comments/gp2uey/daimler_internal_source_code_leaked_some_rust/
---
Over the weekend, the contents of an internal Gitlab instance from Daimler (Mercedes) was leaked. It contains, among other things, the source code of their internal "onboard logic unit" (OLU) software systems for Mercedes vans.

See this ZDNET post for details: [https://www.zdnet.com/article/mercedes-benz-onboard-logic-unit-olu-source-code-leaks-online/](https://www.zdnet.com/article/mercedes-benz-onboard-logic-unit-olu-source-code-leaks-online/)

Here's a list of Rust repos in the Daimler leak: [https://git.rip/search?utf8=%E2%9C%93&amp;search=rust&amp;group\_id=169&amp;project\_id=&amp;repository\_ref=&amp;nav\_source=navbar](https://git.rip/search?utf8=%E2%9C%93&amp;search=rust&amp;group_id=169&amp;project_id=&amp;repository_ref=&amp;nav_source=navbar)

It seems it's mostly example projects on how to connect to the OLU server from Rust. The main parts of the system are C++. Still interesting!
## [5][Common Rust Lifetime Misconceptions](https://www.reddit.com/r/rust/comments/golrsx/common_rust_lifetime_misconceptions/)
- url: https://github.com/pretzelhammer/rust-blog/blob/master/posts/common-rust-lifetime-misconceptions.md
---

## [6][[Release 1.0] console_engine](https://www.reddit.com/r/rust/comments/gp430d/release_10_console_engine/)
- url: https://www.reddit.com/r/rust/comments/gp430d/release_10_console_engine/
---
Hello everybody,

I'm happy to present you the first stable release of my crate [console\_engine](https://crates.io/crates/console_engine) !

## But what is console_engine ?

It's a Terminal User Interface (tui) library built on top of [crossterm](https://crates.io/crates/crossterm). It's far from being as cool as [tui](https://crates.io/crates/tui) himself but I think that console\_engine can have its own advantages for some use cases.

The crate has two main structs : `ConsoleEngine` and `Screen`. Let's start with `Screen` first :

* a Screen is like a virtual screen that you create with a specific width and height.
* It has many drawing functions that allow you to plot characters wherever you want. You can print text on specific coordinates but also draw lines, rectangles, circles or even triangles.
* You can print a Screen into another one.
* You can resize a Screen if needed
* When your screen is ready, you just have to call the draw function
* Screens are standalone and as such, you can only use it without even using ConsoleEngine !

Now that you're more familiar with Screens, let's take a look into ConsoleEngine :

* It's a structure where you provide a width, height and target frame per seconds, and then you just need to run it in an infinite loop.
* It'll capture keyboard and mouse inputs, and give them back via some functions
* In the meantime, it also has internally a Screen. You can do the same operations to the engine as if it's a Screen.
* each "frame" is drawn of top of each other, so you can do some animations on screen
* terminal resizing is supported : you just need to call one function to auto-resize the engine's screen

I tried to provide many examples and have widely documented most of the functions to help new users to get used to it.

I hope that this crate find its public and that's why I wanted to communicate at least on its 1.0 release.
## [7][Demo Zoo: Zero Cost Abstractions in C++20, Rust, &amp; Zig](https://www.reddit.com/r/rust/comments/gp0vb2/demo_zoo_zero_cost_abstractions_in_c20_rust_zig/)
- url: https://www.youtube.com/watch?v=43X9ia-qpds&amp;feature=share
---

## [8][Which Websocket library to use?](https://www.reddit.com/r/rust/comments/goxm85/which_websocket_library_to_use/)
- url: https://www.reddit.com/r/rust/comments/goxm85/which_websocket_library_to_use/
---
The last time this question was asked was in 2018 it is not 2020 and I was wondering if the information presented in previous questions is still applicable. [Previous Question!](https://www.reddit.com/r/rust/comments/6v44nn/ask_rustaceans_which_web_socket_library_to_use/)

**What initiated this post?**

*  When researching the available crates that can be used to satisfy websocket needs I found there are [84 different crates!](https://lib.rs/web-programming/websocket) and it is not clear what their pros/cons are.
* When this question was last asked the most talked about crates where ``websocket`` and ``ws`` however now it looks like the most used crate is actually ``tungstenite`` or ``tokio-tungstenite``. ``Tungstenite`` also seems to be the most active out of all the crates, perhaps this is the reason for it being most used?
## [9][Tour of Rust - Chapter 4 Released! - Generic Types, new Brazilian Portuguese translation, and more](https://www.reddit.com/r/rust/comments/gp3wc4/tour_of_rust_chapter_4_released_generic_types_new/)
- url: https://tourofrust.com/chapter_4_en.html
---

## [10][Tide and SQLX Rust API](https://www.reddit.com/r/rust/comments/gp3yrl/tide_and_sqlx_rust_api/)
- url: https://www.reddit.com/r/rust/comments/gp3yrl/tide_and_sqlx_rust_api/
---
Hello 🙋‍♂️

I have attempted to create my first Rust API using Tide and SQLX.

[https://github.com/cobbinma/example-rust-api](https://github.com/cobbinma/example-rust-api)

I am at a point now when I would really appreciate feedback. 

My main question is over testing it. What would be the recommended approach to testing an API like this? Would you create a database level trait and mock it or test on an actual database?

Thanks in advance!
## [11][[RANT] the rust ecosystem on Windows....](https://www.reddit.com/r/rust/comments/gortry/rant_the_rust_ecosystem_on_windows/)
- url: https://www.reddit.com/r/rust/comments/gortry/rant_the_rust_ecosystem_on_windows/
---
Okay, so I’ve been trying to write a tool to interact with a kernel driver in Windows recently. My main source has been the `winapi` and `kernel32` crates and it’s been a goddamned nightmare.

Trying to define all the unique Nullable types from rust to do a task as simple as `CreateFileA` to the driver name, and `DeviceIoControl` methods, the argument types and documentation are a steaming pile of undefined behaviour. Never mind the fact that the rust compiler is constantly bitching at me for unsafe types and methods being used all over the place.

I’m just curious, has anyone had any success writing low level windows tools in Rust? And if so... how the fuck are you doing this? My code is littered with `std::mem::zeroes` and `unsafe` C types, plus it doesn’t even work, I suddenly have lost all of the joys of Rusts memory system and compiler warnings. Suddenly I can compile my code and when I run it, it doesn’t work AT ALL. It’s worse than writing it in C.

Question; when will the Rust ecosystem on Windows mature? Is anyone even working on it? I’m really distraught and it’s making me hate my favourite language which has served me so incredibly well over the years of writing it on Linux systems. So sad. Sorry for the rant just needed to get this out there.

Have a lovely weekend y’all!!

EDIT:
Just to be clear, I’m not dissing the maintainers of the crates, I think they’ve done an incredible job at getting access to those API’s . But the API’s are so outdated that rust doesn’t plug into them like we’re used to in Linux. It’s a real shame and I don’t see it changing until the WinAPI is improved.

Edit2: 
And finally, the rust community has proven me wrong as I hoped! The ecosystem isn’t perfect sure, but there are lots of cool ways to make it manageable and some really impressive projects out there to use as a basis. Thank you everyone for listening to my rant and being so helpful.
## [12][What am I missing here??? (I am new to Rust btw)](https://www.reddit.com/r/rust/comments/gp3czx/what_am_i_missing_here_i_am_new_to_rust_btw/)
- url: https://www.reddit.com/r/rust/comments/gp3czx/what_am_i_missing_here_i_am_new_to_rust_btw/
---
This will not compile. I have no hair left so I am hoping someone can point me in the right direction...

`#[test]`  
`fn a_test() {`  
 `let mymarked: Arc&lt;Mutex&lt;bool&gt;&gt; = Arc::new(Mutex::new(false));`  
    `std::thread::spawn(move || {`  
        `*mymarked.get_mut().unwrap() = true;`  
    `});`  
`}`  
The error I get is:

`error[E0596]: cannot borrow data in an \`Arc\` as mutable`

   `--&gt; src\lib.rs:197:14`

`|`

`197 |             *mymarked.get_mut().unwrap() = true;`

`|              ^^^^^^^^ cannot borrow as mutable`

`|`

`= help: trait \`DerefMut\` is required to modify through a dereference, but it is not implemented for \`std::sync::Arc&lt;std::sync::Mutex&lt;bool&gt;&gt;\``

Thanks for any help. Yes I tried RTFM, rustc --explain E0596, etc. I am very new to the language so it is probably something simple.

Thanks,  
SLDR  
(Stephen L. De Rudder)
