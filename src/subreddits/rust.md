# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (17/2020)!](https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (17/2020)?](https://www.reddit.com/r/rust/comments/g4nupj/whats_everyone_working_on_this_week_172020/)
- url: https://www.reddit.com/r/rust/comments/g4nupj/whats_everyone_working_on_this_week_172020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-17-2020/41247?u=llogiq)!
## [3][Georgia Tech CS-3210: Write an OS for the Raspberry Pi in Rust](https://www.reddit.com/r/rust/comments/g5poww/georgia_tech_cs3210_write_an_os_for_the_raspberry/)
- url: https://tc.gts3.org/cs3210/2020/spring/lab.html
---

## [4][wgpu-rs on the web](https://www.reddit.com/r/rust/comments/g5sfth/wgpurs_on_the_web/)
- url: https://gfx-rs.github.io/2020/04/21/wgpu-web.html
---

## [5][Thank you for ron (Rusty Object Notation)!](https://www.reddit.com/r/rust/comments/g5rlt5/thank_you_for_ron_rusty_object_notation/)
- url: https://www.reddit.com/r/rust/comments/g5rlt5/thank_you_for_ron_rusty_object_notation/
---
Quick shout out to the creators of [ron](https://crates.io/crates/ron) (and `serde`, while we're at it). In my hobby project, I've been bitten by JSON's restrictions:

* Keys in maps must be strings.
* No commenting.
* Trailing commas are not permitted for when a given element is last in a list/map/object.

`ron` washes its hands of these unpleasant constraints. It's a very pleasant config language to work with. Many thanks to its creators and maintainers.
## [6][OrbTk 0.3.1-alpha2 release](https://www.reddit.com/r/rust/comments/g60b03/orbtk_031alpha2_release/)
- url: https://www.reddit.com/r/rust/comments/g60b03/orbtk_031alpha2_release/
---
Next version of OrbTk is released on [crates.io](https://crates.io)

**Release**

[https://crates.io/crates/orbtk/0.3.1-alpha2](https://crates.io/crates/orbtk/0.3.1-alpha2)

**Changelog**

[https://github.com/redox-os/orbtk/blob/develop/CHANGELOG.md](https://github.com/redox-os/orbtk/blob/develop/CHANGELOG.md)  


**Focus next release**  


* Better performance
* Mutli-Window support 
* Improvements state api
* More use of async (under the hood and for users)
## [7][Nushell 0.13.0 - now with aliases, math, startup commands, and more](https://www.reddit.com/r/rust/comments/g5lvct/nushell_0130_now_with_aliases_math_startup/)
- url: https://www.nushell.sh/blog/2020/04/21/nushell_0_13_0.html
---

## [8][projects ideas?](https://www.reddit.com/r/rust/comments/g5y1g8/projects_ideas/)
- url: https://www.reddit.com/r/rust/comments/g5y1g8/projects_ideas/
---
hey rustaceans, just completed the rust programming language book, can anyone suggest few projects i can add my first contribution to or some project idea i should work on to become more efficient at rust?
## [9][Pathfinder (Rust GPU vector graphics renderer) is now on crates.io](https://www.reddit.com/r/rust/comments/g5htil/pathfinder_rust_gpu_vector_graphics_renderer_is/)
- url: https://twitter.com/pcwalton/status/1251254942611849219
---

## [10][RFC: Transition to rust-analyzer as our official LSP implementation](https://www.reddit.com/r/rust/comments/g5ckqi/rfc_transition_to_rustanalyzer_as_our_official/)
- url: https://github.com/rust-lang/rfcs/pull/2912
---

## [11][Show me the most illegal Rust code you've ever seen/written!](https://www.reddit.com/r/rust/comments/g5rsuh/show_me_the_most_illegal_rust_code_youve_ever/)
- url: https://www.reddit.com/r/rust/comments/g5rsuh/show_me_the_most_illegal_rust_code_youve_ever/
---

## [12][Emulsion is a 100% Rust image viewer.](https://www.reddit.com/r/rust/comments/g5d5j9/emulsion_is_a_100_rust_image_viewer/)
- url: https://www.reddit.com/r/rust/comments/g5d5j9/emulsion_is_a_100_rust_image_viewer/
---
Emulsion is fast, small, displays transparency properly, and uses GPU acceleration. It's also minimalistic with the design philosophy of being nothing more than an image viewer but being very good at that.

Check out the website [https://arturkovacs.github.io/emulsion-website/](https://arturkovacs.github.io/emulsion-website/)

Below are some measurements that measure the time from pressing enter on the image file and having the image shown in the Emulsion window. These were made relying on my reflexes and a stopwatch. Originally I intended to make measurements for Windows' Photos but it's suprisingly difficult to get comparable values because the Photos app on my machine takes around 4-5 seconds to start up after a reboot while emulsion is around half of that. But then Photos takes around 1.3 seconds consistently to load almost any image, large and small alike. Photos loads images progressively which Emulsion doesn't do yet so Emulsion falls short in load time for very high resolutions.

    4798 * 4798 JPG    1.07 sec
    2880 * 1800 JPG    0.85 sec
    194 * 255 PNG      &lt;0.5 sec
