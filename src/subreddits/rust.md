# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (43/2020)!](https://www.reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 361](https://www.reddit.com/r/rust/comments/jg7hkt/this_week_in_rust_361/)
- url: https://this-week-in-rust.org/blog/2020/10/21/this-week-in-rust-361/
---

## [3][usethe.computer - XMHell: Handling 38GB of UTF-16 XML with Rust](https://www.reddit.com/r/rust/comments/jgixsy/usethecomputer_xmhell_handling_38gb_of_utf16_xml/)
- url: http://usethe.computer/posts/14-xmhell.html
---

## [4][Why I created a package and project manager for C (in Rust ofc !)](https://www.reddit.com/r/rust/comments/jgl1rz/why_i_created_a_package_and_project_manager_for_c/)
- url: https://www.reddit.com/r/rust/comments/jgl1rz/why_i_created_a_package_and_project_manager_for_c/
---
A few months ago, I was wondering why we hoard Makefiles and why it is so painful to use an external library in a C project.

So I had this idea : Creating a project manager &amp; build tool for the C programming language.

I started to write a piece of code in C and it was not functionning properly (Cause I'm one of the worst C developers in this world) but i continued 'till we cannot run that thing.

At the same time, I was learning [Rust](https://rust-lang.org) ; so, I decided to try to rewrite the whole project in that language.

After a few weeks of rewriting, I had a correct product. The 08/10/2020, Wanager 1.0 was released. It had only a few features : project creation and reinitializing, project build and run and header creation. 

At that point, someone called Lockeer told me that it will be cool if we can manage libraries. 

So I wrote a simple system to install libraries hosted on my vps, with a submission system based on mailing. It was working, but limited because of vps bandwidth and the complexity of submiting by email.

So SuperFola poped up :

https://user-images.githubusercontent.com/61330081/96449113-aa418a80-1214-11eb-97e8-32c7afd86ff8.png

At first time, I decided, as he advised me, to use the github api that produces a tar archive of the repo. 
I stucked on that for weeks because the command I was running was producing a corrupted file. 

After raging on that problem, I realised that I'll gain some portability and time with cloning directly with a git command. 

It worked good, so, everything is fine !

But Il_totore opened an issue :

https://user-images.githubusercontent.com/61330081/96449715-8df21d80-1215-11eb-9a22-588c77ce9870.png

So I made python support for build scripts with minimal version of 3.5.x and allowed path specification.

After that, on his advice, I made kind of Python API to have nice build scripts and it produced that :

```py
from wngbuild import * # Import all from wngbuild module

build = BuildProfile(files="src/*.c",output="build/custom/prog.exe" ) # setup a build profile that will compile all files in src/ and place the binary in build/custom/prog.exe
build.cc = "C:\\MinGW\\bin\\gcc.exe" # Setup the compiler (optional, by default "gcc")
build.flags = "-W -Wall -Werror -Wextra" # Setup the flags that the command will be run with (optional)

build.run() # Run the compilation command
build.runOutput() # Run the binary produced by the compilation command (Will raise an error if the compilation command fails)
```

https://github.com/Wmanage/wng

It is still WIP, there are loads of features that I can add to it but I will be more very happy to answer your questions or help you use it.

Thanks for reading and have a nice day.
## [5][Introducing rust-gpu v0.1 🐉 · EmbarkStudios/rust-gpu](https://www.reddit.com/r/rust/comments/jg056t/introducing_rustgpu_v01_embarkstudiosrustgpu/)
- url: https://github.com/EmbarkStudios/rust-gpu/releases/tag/v0.1
---

## [6][Microsoft is hiring for a position "focused on delivering Rust compiler improvements"](https://www.reddit.com/r/rust/comments/jg11lq/microsoft_is_hiring_for_a_position_focused_on/)
- url: https://careers.microsoft.com/us/en/job/917051/Senior-Software-Engineer
---

## [7][Knurling-rs changelog #3](https://www.reddit.com/r/rust/comments/jgmh6d/knurlingrs_changelog_3/)
- url: https://ferrous-systems.com/blog/knurling-changelog-3/
---

## [8][The future of RLSL](https://www.reddit.com/r/rust/comments/jg08jb/the_future_of_rlsl/)
- url: https://maikklein.github.io/rlsl-update3/
---

## [9][Announcing Stability: stability attribute macros for library authors](https://www.reddit.com/r/rust/comments/jgc7hy/announcing_stability_stability_attribute_macros/)
- url: https://github.com/sagebind/stability
---

## [10][mock-io: Mock IO listener and stream for Rust](https://www.reddit.com/r/rust/comments/jgjj0y/mockio_mock_io_listener_and_stream_for_rust/)
- url: https://www.reddit.com/r/rust/comments/jgjj0y/mockio_mock_io_listener_and_stream_for_rust/
---
[https://crates.io/crates/mock-io](https://crates.io/crates/mock-io)
## [11][Why I Wrote zeroconf-rs: A cross-platform wrapper around Bonjour and Avahi](https://www.reddit.com/r/rust/comments/jg42ph/why_i_wrote_zeroconfrs_a_crossplatform_wrapper/)
- url: https://windy1.github.io/posts/2020/10/22/why-i-wrote-zeroconf-rs.html
---

## [12][CRUD with Rocket and Sled](https://www.reddit.com/r/rust/comments/jgdrj5/crud_with_rocket_and_sled/)
- url: https://mbuffett.com/posts/rocket-sled-tutorial/
---

