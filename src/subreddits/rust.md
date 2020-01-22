# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (4/2020)!](https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/)
- url: https://www.reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here_42020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eo6pjy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (4/2020)?](https://www.reddit.com/r/rust/comments/eraop8/whats_everyone_working_on_this_week_42020/)
- url: https://www.reddit.com/r/rust/comments/eraop8/whats_everyone_working_on_this_week_42020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-4-2020/37181?u=llogiq)!
## [3][std.rs: go directly to stable rust docs!](https://www.reddit.com/r/rust/comments/es35u6/stdrs_go_directly_to_stable_rust_docs/)
- url: https://www.reddit.com/r/rust/comments/es35u6/stdrs_go_directly_to_stable_rust_docs/
---
I just purchased and set up [std.rs](https://std.rs) so you can type that into your URL bar and go straight to https://doc.rust-lang.org/stable/std/. I know, I know, there's already [docs.rs/std](https://docs.rs/std), so why bother with this? Well, it's like... four fewer characters to type in!

^yes ^^I'm ^^^a ^^^^vim ^^^^^user
## [4][[blog post] Why the stabilization of advanced slice patterns is a big deal to me](https://www.reddit.com/r/rust/comments/es1824/blog_post_why_the_stabilization_of_advanced_slice/)
- url: https://www.reddit.com/r/rust/comments/es1824/blog_post_why_the_stabilization_of_advanced_slice/
---
With the impending stabilization of advanced slice patterns (c'mon 1.42!), I wrote a little blog post summarizing what we're getting and why I think it's a big deal:
https://thomashartmann.dev/blog/feature(slice_patterns)
## [5][Saved hours (days?) of frustration and confusion thanks to Rust](https://www.reddit.com/r/rust/comments/es2wrf/saved_hours_days_of_frustration_and_confusion/)
- url: https://www.reddit.com/r/rust/comments/es2wrf/saved_hours_days_of_frustration_and_confusion/
---
At work I work on a mostly C codebase. C was chosen for reasons that made sense in the past, but things have changed now so various other languages have been added, but the center is all still C.

The C code currently uses PCRE2 purely for regex validation, but I wanted to use capture groups. I've never worked with the library before and it looks like a nightmare to use: [https://www.pcre.org/current/doc/html/](https://www.pcre.org/current/doc/html/).

Half a year ago I added Rust with dynamic linking and FFI into the codebase, so I figured it's probably easier to use Rust and FFI with the regex crate. And I was right - it was super easy.

Huge thanks to the Rust team and community for making my life in the C world so much easier :D
## [6][Introducing Fourier: the fastest fast Fourier transform (FFT) in Rust ...and a request for help!](https://www.reddit.com/r/rust/comments/ervc99/introducing_fourier_the_fastest_fast_fourier/)
- url: https://www.reddit.com/r/rust/comments/ervc99/introducing_fourier_the_fastest_fast_fourier/
---
Hi all,

Today I released the initial version of Fourier, the fastest FFT written in pure Rust.

[https://crates.io/crates/fourier](https://crates.io/crates/fourier)

[https://github.com/calebzulawski/fourier](https://github.com/calebzulawski/fourier)

My goal is to provide a competitor to FFTW with a *non-viral free license* (FFTW is GPL licensed, but offers a very expensive commercial license).

This initial release is already very fast, however I'm looking for help from anyone who is interested in bringing it up to par with FFTW.  I would also appreciate help creating a WASM package.  I've opened a few issues in GitHub to get started--just send a PM or open a PR if you're interested!

# Benchmarks

The repository contains benchmarks you can run yourself, but here's the summary of how it performs on my machine (Intel CPU with AVX)

* For `f32` powers of two, about 2x faster than RustFFT, but 3x slower than FFTW
* For `f64` powers of two, about 1.5x faster than RustFFT, but 2x slower than FFTW
* For large prime sizes, sometimes 10x+ (!) faster than RustFFT, but 1.5x slower than FFTW
## [7][What is Rust and why is it so popular? - Stack Overflow Blog](https://www.reddit.com/r/rust/comments/erw561/what_is_rust_and_why_is_it_so_popular_stack/)
- url: https://stackoverflow.blog/2020/01/20/what-is-rust-and-why-is-it-so-popular/
---

## [8][Finished my first Rust project - A website built with the async Rocket branch](https://www.reddit.com/r/rust/comments/es3ya9/finished_my_first_rust_project_a_website_built/)
- url: https://www.reddit.com/r/rust/comments/es3ya9/finished_my_first_rust_project_a_website_built/
---
Hi /r/rust. My friend told me about Rust a bit ago. I had always wanted to learn C++ as I've always had a big performance bias, but was always afraid of the complexity (and not being smart enough for it). So it sat on the backburner for a while, but then I kept seeing Rust posted on Hacker News (Baader Meinhof effect?) and it always being the #1 most loved language on the Stackoverflow Developer survey. The TechEmpower benchmarks pushed me over the edge, so then I devoured Steve and Carol's `The Rust Programming Language`.

To solidify my knowledge I was looking for a project to make and was re-reading Paul Graham's essays at the same time and he mentioned to make something that you yourself would want. I then also remembered a desire I had a while back to be able to find likeminded people around me (for example, someone who also likes Rust!). So I decided to build that as my project.

The core idea is that instead of swiping on people, you instead swipe on concepts and ideas (example: `hunting, vaccines, Christianity, cities, podcasts, etc.`) where swiping right means you like it / identify with it, and left is the converse.

The further you swipe in either direction, the stronger your vote. It then uses the Manhattan distance formula to compute your similarity to others. You can also view statistics like how often a concept is liked or disliked (or neutral), how long on average it takes for people to decide, how that card correlates with other cards (for example, Military and Fracking are highly correlated with one another).

You can then also view clusters of cards on profile pages. These are cards that are found to be clustered together in that their votes are highly correlated with one another. You can see how you and others align to these clusters. It's also a bit of a privacy feature as well as you cannot see how people vote on individual cards, only how they align to clusters. So their cluster alignment somewhat "masks" their individual card votes, or at least provides some plausible deniability. It's also just interesting to find out what groups of cards tend to cluster around each other.

Currently the clustering algorithm is a bit ad-hoc as math is _definitely_ not my strong suit. There are around 250 cards at the moment and originally I wanted to have an exact algorithm for computing similarity that also took into account weights, but I couldn't quite figure out how to have that while also allowing people to sort by similarity quickly at scale.

I found out that it's basically the K-Nearest-Neighbors problem with 250 dimensions and that is a bit tricky (for me at least). So instead I wrote a small algorithm ([which might be able to be replaced with this Rust crate?](https://old.reddit.com/r/rust/comments/dx2k2q/hilbert_curve_transformation_crate_for/)) to create clusters of cards, and then used the [Postgres CUBE data structure](https://www.postgresql.org/docs/9.5/cube.html) to be able to calculate and index someone's alignment in what is now 25-dimensional space (which is much more tractable than 250 dimensional space!)

So, on to the tech stack!

My only two frontend dependencies are React and axios (I'll probably refactor out axios soon). I'm a bit afraid of npm and I like limiting my dependencies. Also, small bundle sizes are great!

The backend is more interesting. I'm using `rand`, `bcrypt`, `serde`, `rusoto`, `oauth2`, `reqwest`, `time`, `rocket` (async branch), `tokio`, `tokio-postgres`, `futures`, `deadpool`, `deunicode`, `pin-project-lite`, and `async-stream`.

I'm then using nginx as a reverse proxy to my Rocket server, and have the server itself currently hosted on EC2 behind CloudFront, with assets on S3.

When I had started the project, async await wasn't quite ready yet, and future combinators were killing me with borrowing errors. Eventually I found out about the async Rocket branch, and Jeb Rosen and Sergio were always _extremely_ accommodating and helpful with all of my newb questions. I also really liked the Rocket syntax so I decided to rewrite it in Rocket!

I was able to get rid of so many clones and lines of code and started feeling really good about the code-base. It was just really clean and elegant. I'm also now confident about the code, which is great. There were so many times when the compiler would refuse to compile and then I'd go, "oh, right, yeah. good catch." I still have some residual PTSD from my last node server which would randomly crash with null reference exceptions due to me missing an edge case.

The only issue was that database access was still synchronous, but recently /u/bikeshedder wrote the amazing `deadpool` library, which I was able to seamlessly integrate and immediately significantly improve my runtime performance. [I wrote about that here](https://old.reddit.com/r/rust/comments/e9n6mx/creating_and_using_a_postgres_database_pool_is/?st=k5ok81ng&amp;sh=440f06b6).

Lastly, I'd just like to [re-emphasize](https://old.reddit.com/r/rust/comments/ereh5h/the_thank_you_thread/ff4f5z2/?st=k5okcaom&amp;sh=8b98f398) my thanking of the Rust community. I truly have not yet had a single bad interaction with anyone. The #beginners discord channel, everyone in IRC, Gitter, Riot, Reddit, etc. have all been extremely welcoming and helpful to a noob like me, and thanks to them I was able to finish this project.

If you'd like to check it out, here's the site: https://www.kardius.com I tried to make as many features available as possible without logging in, so no pressure to create an account at all. You should be able to view the cards on the _Swipe_ and _Cards_ page. I doubt it's good enough yet, but hopefully I can make it better! If you have any suggestions or feedback I'd love to hear it. Thanks for reading!
## [9][This Week in Rust 322](https://www.reddit.com/r/rust/comments/esbf8h/this_week_in_rust_322/)
- url: https://this-week-in-rust.org/blog/2020/01/21/this-week-in-rust-322/
---

## [10][Async Interview #5: Steven Fackler](https://www.reddit.com/r/rust/comments/es48d2/async_interview_5_steven_fackler/)
- url: http://smallcultfollowing.com/babysteps/blog/2020/01/20/async-interview-5-steven-fackler/
---

## [11][C++ for Rust Programmers?](https://www.reddit.com/r/rust/comments/es7oq3/c_for_rust_programmers/)
- url: https://www.reddit.com/r/rust/comments/es7oq3/c_for_rust_programmers/
---
I am in the unique position of knowing rust but not C++ and only kinda C. Are there any good resources for writing high quality C++ that doesn't explode as this subreddit leads me to believe? I have to learn it for $work :(

My biggest issue so far is that there are just no quality resources like there are for rust. And due to rust I have a hard time trusting other languages.

Is there any way to get a rust like experience in modern C++?
## [12][ptable: A complete periodic table for rust](https://www.reddit.com/r/rust/comments/es18dd/ptable_a_complete_periodic_table_for_rust/)
- url: https://crates.io/crates/ptable
---

