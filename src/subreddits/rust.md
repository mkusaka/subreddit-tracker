# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (1/2020)!](https://www.reddit.com/r/rust/comments/ehk18j/hey_rustaceans_got_an_easy_question_ask_here_12020/)
- url: https://www.reddit.com/r/rust/comments/ehk18j/hey_rustaceans_got_an_easy_question_ask_here_12020/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

The Rust-related IRC channels on irc.mozilla.org (click the links to open a web-based IRC client):

 - [#rust](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust) (general questions)
 - [#rust-beginners](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-beginners) (beginner questions)
 - [#cargo](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23cargo) (the package manager)
 - [#rust-gamedev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-gamedev) (graphics and video games, and see also [/r/rust_gamedev](https://www.reddit.com/r/rust_gamedev))
 - [#rust-osdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-osdev) (operating systems and embedded systems)
 - [#rust-webdev](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-webdev) (web development)
 - [#rust-networking](https://chat.mibbit.com/?server=irc.mozilla.org%3A%2B6697&amp;amp;channel=%23rust-networking) (computer networking, and see also [/r/rust_networking](https://www.reddit.com/r/rust_networking))

Also check out [last week's thread](https://reddit.com/r/rust/comments/eelr59/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 319](https://www.reddit.com/r/rust/comments/ej4eom/this_week_in_rust_319/)
- url: https://this-week-in-rust.org/blog/2019/12/31/this-week-in-rust-319/
---

## [3][Announcing "Haskell Bits", a Functor/Applicative/Monad hierarchy in Rust](https://www.reddit.com/r/rust/comments/ejubtb/announcing_haskell_bits_a_functorapplicativemonad/)
- url: https://github.com/clintonmead/haskell_bits
---

## [4][Reducing support for 32-bit Apple targets](https://www.reddit.com/r/rust/comments/ejk5s1/reducing_support_for_32bit_apple_targets/)
- url: https://blog.rust-lang.org/2020/01/03/reducing-support-for-32-bit-apple-targets.html
---

## [5][Introduce async-socks5: async/.await version of SOCKS5!](https://www.reddit.com/r/rust/comments/ejwbaq/introduce_asyncsocks5_asyncawait_version_of_socks5/)
- url: https://github.com/ark0f/async-socks5
---

## [6][Has anyone experimented with Qt for WebAssembly using Rust?](https://www.reddit.com/r/rust/comments/ejvfnn/has_anyone_experimented_with_qt_for_webassembly/)
- url: https://www.reddit.com/r/rust/comments/ejvfnn/has_anyone_experimented_with_qt_for_webassembly/
---
I've been trying to develop a cross-platform desktop application that can also target the web. It seems like the two major choices are:

1. A framework that compiles to HTML/CSS/JS, which is then wrapped by Electron for desktop use
2. A framework that compiles to native for desktop, or WebAssembly for web use

For 1, I looked into using [Yew](https://github.com/yewstack/yew) to create the WASM, but I found it difficult to use JS for interactive components, and Yew itself doesn't have the components I need out of the box (floating windows, dockable panels).

For 2, it seems like Qt is a good option, because it compiles either to native or to WebAssembly depending on your target.

Has anyone tried to use Rust bindings for Qt, compiling to WebAssembly? If there's an existing build process which works, I don't want to reinvent the wheel. It would also be great to contribute it back to [rust-qt](https://github.com/rust-qt/examples) for others to use.
## [7][Alex Ionescu (CrowdStrike) hires Rust developers](https://www.reddit.com/r/rust/comments/ejj6p1/alex_ionescu_crowdstrike_hires_rust_developers/)
- url: https://www.reddit.com/r/rust/comments/ejj6p1/alex_ionescu_crowdstrike_hires_rust_developers/
---
I thought you might be interested in this. Alex Ionescu is a legend in the Windows internals world, a co-author of the Windows Internals book, a contributor to the ReactOS open source Windows (wannabe) alternative, and a speaker at well known conferences. And currently a Chief Architect at Crowdstrike. It turned out to be quite a praise :) but that's because I really admire him!  
[https://twitter.com/aionescu/status/1213151075336888325](https://twitter.com/aionescu/status/1213151075336888325)
## [8][Announcing AeroRust - The Unofficial Working Group For Rust in Aerospace](https://www.reddit.com/r/rust/comments/ejdv7w/announcing_aerorust_the_unofficial_working_group/)
- url: https://www.reddit.com/r/rust/comments/ejdv7w/announcing_aerorust_the_unofficial_working_group/
---
Hi Folks,

We are happy to announce the formation of AeroRust Unofficial Working Group. Lately the Aerospace industry is picking up speed in commercialisation and even faster and rapid development in the area.

This Working group is aiming to help push the Open-source community  more into the growing Aerospace industry, by providing information,  materials, tools, crates and etc. to hobbyists and the industry.  


If you are interested, come join us:   


* Github: [https://github.com/AeroRust/Welcome](https://github.com/AeroRust/Welcome)
* Discord: [https://discord.gg/RXNsMXc](https://discord.gg/RXNsMXc)
* Resources: [https://github.com/AeroRust/awesome-space](https://github.com/AeroRust/awesome-space)
## [9][Code Golf now supports rust :-)](https://www.reddit.com/r/rust/comments/ejk773/code_golf_now_supports_rust/)
- url: https://code-golf.io
---

## [10][Announcing Razor, a model-finder for first-order theories](https://www.reddit.com/r/rust/comments/ejld0s/announcing_razor_a_modelfinder_for_firstorder/)
- url: https://www.reddit.com/r/rust/comments/ejld0s/announcing_razor_a_modelfinder_for_firstorder/
---
I’m so excited to have published my first Rust crates and would love feedback:

razor-fol: a library for parsing and syntactic manipulation of first-order (logic) formulae.
razor-chase: a library for constructing models for first-order theories.
razor: a model-finding tool for first-order theories.

Here is a link to the GitHub repo: https://github.com/salmans/rusty-razor

The project is still in its infancy but I’ve had a blast experimenting with Rust. The theorem-proving community cares about correctness and speed, so I’m expecting to see similar project in Rust in the near future.

Any comment, feedback, or suggestion is much appreciated!
## [11][Await-ing within trait function](https://www.reddit.com/r/rust/comments/ejpx88/awaiting_within_trait_function/)
- url: https://www.reddit.com/r/rust/comments/ejpx88/awaiting_within_trait_function/
---
Given I have a trait function that needs to return a `futures::future::Either`, and must do an `await` to read a Stream, how can I `await` within the function? I've tried using `async move {...}` but I cannot get the return type right. This is a simplifed version of what I am trying:
```
async move {
            let body_result: Result&lt;BytesMut, Error&gt; = get_request_body(&amp;mut req).await;
            match body_result {
                Ok(bytes) =&gt; Either::Left(svc.call(req)),
                Err(_) =&gt; Either::Right(ok(req.into_response(
                    HttpResponse::BadRequest()
                        .body("No payload found.")
                        .into_body(),
                ))),
            }
        }
        .boxed_local()
```
For the sake of this example it doesn't do anything with the `BytesMut`, but my problem is the return type. I'd like it to be `LocalBoxFuture&lt;'static, Either&lt;S::Future, Ready&lt;Result&lt;Self::Response, Self::Error&gt;&gt;&gt;&gt;` but this is the error I get:
```
  --&gt; src\custom_middleware\auth.rs:79:12
   |
79 | impl&lt;S, B&gt; Service for AuthenticatorMiddleware&lt;S&gt;
   |            ^^^^^^^ expected enum `futures_util::future::either::Either`, found enum `std::result::Result`
   |
   = note: expected type `futures_util::future::either::Either&lt;&lt;S as actix_service::Service&gt;::Future, futures_util::future::ready::Ready&lt;std::result::Result&lt;actix_web::service::ServiceResponse&lt;B&gt;, actix_http::error::Error&gt;&gt;&gt;`

```

Not sure how I can get there, though. Any suggestions?
## [12][Nvim-rs: Rust library for neovim clients](https://www.reddit.com/r/rust/comments/ejfceo/nvimrs_rust_library_for_neovim_clients/)
- url: /r/neovim/comments/eiy5if/nvimrs_rust_library_for_neovim_clients/
---

