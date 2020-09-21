# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (39/2020)!](https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (39/2020)?](https://www.reddit.com/r/rust/comments/iwxjdh/whats_everyone_working_on_this_week_392020/)
- url: https://www.reddit.com/r/rust/comments/iwxjdh/whats_everyone_working_on_this_week_392020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-39-2020/49088?u=llogiq)!
## [3][A New Backend for Cranelift, Part 1: Instruction Selection](https://www.reddit.com/r/rust/comments/iwumjn/a_new_backend_for_cranelift_part_1_instruction/)
- url: https://cfallin.org/blog/2020/09/18/cranelift-isel-1/
---

## [4][Blog Post: Why Not Rust?](https://www.reddit.com/r/rust/comments/iwij5i/blog_post_why_not_rust/)
- url: https://matklad.github.io/2020/09/20/why-not-rust.html
---

## [5][Face Detection with Rust, TensorFlow, and WebAssembly](https://www.reddit.com/r/rust/comments/iwyj85/face_detection_with_rust_tensorflow_and/)
- url: https://www.secondstate.io/articles/face-detection-ai-as-a-service/
---

## [6][rust-analyzer changelog #43](https://www.reddit.com/r/rust/comments/iwzja8/rustanalyzer_changelog_43/)
- url: https://rust-analyzer.github.io/thisweek/2020/09/21/changelog-43.html
---

## [7][Text-Based Web Browser for Rust newbies](https://www.reddit.com/r/rust/comments/iwvn00/textbased_web_browser_for_rust_newbies/)
- url: https://www.reddit.com/r/rust/comments/iwvn00/textbased_web_browser_for_rust_newbies/
---
[LINK TO REPOSITORY](https://github.com/Debmalya99/TermSurf)

Hi all, I created **TermSurf: A text-based web browser in Rust** for learning purposes. I am not a great programmer, so I learnt a lot of new concepts as I made this project. 

A web browser has a javascript engine inside it, and it can render HTML and CSS styles. Now, since my project is text-based, and has no GUI, it **does not feature HTML rendering or CSS.** It simply displays text.

But the problem was with Javascript. I wanted to create functional applications that could be run using the browser, but I had never written a program that had an embedded scripting engine inside it. 

**Here was my plan, that browser will be:**

* A simple Http client that would make web requests to some server.
* The server would return some text response.
* The text response will actually be source code for a scripting language.
* The browser will have a scripting engine embedded into it, so it can run the obtained source code.
* The scripting language will have functions to make web requests to the server for interactivity.

**As an example, I have made some sample servers** all of which can be run in your localhost if you have python and flask installed. All of this server code is included in the repository itself. More information is provided in the README.md

**The Scripting Language that is currently in use:** I have used the [Rhai](https://github.com/jonathandturner/rhai) scripting language for Rust. It is very easy to learn and use, very similar to rust itself, but without the explicit declaration of types.

**Thank you for your time in reading this, if you feel interested please check it out. I am looking for contributors to this project. If you are a Rust newbie like me and would like to help, it will be very helpful.**
## [8][A Bazel Persistent Worker for Rust](https://www.reddit.com/r/rust/comments/iwpg98/a_bazel_persistent_worker_for_rust/)
- url: https://nikhilism.com/post/2020/bazel-persistent-worker-rust/
---

## [9][API documentation... how to improve my experience?](https://www.reddit.com/r/rust/comments/iwqcjv/api_documentation_how_to_improve_my_experience/)
- url: https://www.reddit.com/r/rust/comments/iwqcjv/api_documentation_how_to_improve_my_experience/
---
I'll preface this by explaining that most of my dev work is done in Go, including 80% of my professional work.

I'm trying to learn (and eventually master) Rust, because I think it will have serious value in the work that my organization is trying to accomplish. However, I'm having a hell of a time. My default mode of learning is to reimplement examples I find scattered around the web, using API documentation to understand exactly what the function I'm calling is doing.

Unfortunately... it's extremely hard to find that function, sometimes. With Go's API documentation, the function is always listed under the struct or package that implements it, regardless of whether it's implementing an external interface. That doesn't seem to be the case with Rust. As a (possibly not great) example, the `to_string` method on an `i32`; yes, I (now) understand that it's implementing the `std::string::ToString` trait, but without already knowing that, discovering that that is the source of the `to_string` method on `i32` is not straightforward.

I guess I'm looking for two outcomes on this post: 1) tips on how to more effectively and efficiently use the Rust API documentation to find what I'm looking for, and 2) hopefully either bring attention to this issue to people in a place to do something about it, or at least gather experiences so that I can put together a comprehensive report/request myself.

Thanks for your advice and experience.
## [10][How to set buffer size on Tokio UDP socket?](https://www.reddit.com/r/rust/comments/iwxqul/how_to_set_buffer_size_on_tokio_udp_socket/)
- url: https://www.reddit.com/r/rust/comments/iwxqul/how_to_set_buffer_size_on_tokio_udp_socket/
---
I need to set the receive buffer size on a Tokio UDP socket. There is no way to do this directly.

`from_std()` looks promising, but net2's UdpBuilder also doesn't offer this functionality.

I thought about getting the raw socket handle—and setting the buffer size myself—using the AsRawFd trait, but I couldn't get that to work - the compiler insists there is no method called `as_raw_fd()`. Is it called something else on Windows, perhaps?

Any help much appreciated!
## [11][Unique Array Elements and their Frequency](https://www.reddit.com/r/rust/comments/iwz1xr/unique_array_elements_and_their_frequency/)
- url: https://datacrayon.com/posts/programming/rust-notebooks/unique-array-elements-and-their-frequency/
---

## [12][Throw-away Code](https://www.reddit.com/r/rust/comments/iwbp4d/throwaway_code/)
- url: https://vorner.github.io/2020/09/20/throw-away-code.html
---

