# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (19/2020)!](https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (19/2020)?](https://www.reddit.com/r/rust/comments/gd6gvk/whats_everyone_working_on_this_week_192020/)
- url: https://www.reddit.com/r/rust/comments/gd6gvk/whats_everyone_working_on_this_week_192020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-19-2020/42034?u=llogiq)!
## [3][Notes on io-uring](https://www.reddit.com/r/rust/comments/gehh0q/notes_on_iouring/)
- url: https://boats.gitlab.io/blog/post/io-uring/
---

## [4][This Week in Rust 337](https://www.reddit.com/r/rust/comments/geagy0/this_week_in_rust_337/)
- url: https://this-week-in-rust.org/blog/2020/05/05/this-week-in-rust-337/
---

## [5][Rust + Webassembly is dope](https://www.reddit.com/r/rust/comments/gecfvq/rust_webassembly_is_dope/)
- url: https://github.com/justinmimbs/rs-asteroids
---

## [6][This Month in Rust OSDev (April 2020)](https://www.reddit.com/r/rust/comments/gefxh7/this_month_in_rust_osdev_april_2020/)
- url: https://rust-osdev.com/this-month/2020-04/
---

## [7][no_std async/await - soon on stable](https://www.reddit.com/r/rust/comments/ge0pwz/no_std_asyncawait_soon_on_stable/)
- url: https://ferrous-systems.com/blog/stable-async-on-embedded/
---

## [8][MeiliSearch in production: taking it to the next level](https://www.reddit.com/r/rust/comments/gehtao/meilisearch_in_production_taking_it_to_the_next/)
- url: https://blog.meilisearch.com/meilisearch-in-production-taking-it-to-the-next-level/
---

## [9][lz-fear: a pure-Rust no-unsafe LZ4 implementation](https://www.reddit.com/r/rust/comments/ge3kee/lzfear_a_purerust_nounsafe_lz4_implementation/)
- url: https://www.reddit.com/r/rust/comments/ge3kee/lzfear_a_purerust_nounsafe_lz4_implementation/
---
https://github.com/main--/rust-lz-fear

* no unsafe code
* blazingly fast decompression (as fast as the C implementation)
* byte-perfect compressor output (except for some weird edge cases where it's still correct but slightly larger)
* MIT license
* simple and relatively high-level implementation, much easier to read than the heavily optimized C code

Posting here because I wrote this library out of covid-induced boredom and now I don't really know what to do with it. If this sounds like something you would either want to use or work on (or both), here you go.

TODOs:

* add more compression algorithms (LZ4-HC and maybe an even heavier approach like https://create.stephan-brumme.com/smallz4/)
* API work (ergonomics and documentation)
* no_std support? The implementation currently allocates a few `Vec&lt;u8&gt;` buffers, but all of those are basically constant.
* 1.0 release
## [10][Setting up a Rust Dev Environment in 2020?](https://www.reddit.com/r/rust/comments/gegrep/setting_up_a_rust_dev_environment_in_2020/)
- url: https://www.reddit.com/r/rust/comments/gegrep/setting_up_a_rust_dev_environment_in_2020/
---
Hi!

I've recently started my Rust journey and have discovered that many of the tutorials etc. on the internet about setting up the dev environment are quite dated (from 2017/2018) and thus make no mention of new tools and changes (we have rust-analyzer now, some tools have become deprecated etc.)

I noticed that the [Rust-lang Community wiki](https://runrust.miraheze.org/wiki/Main_Page) has a space for a page on "Setup your workspace for Rust" but the page does not currently exist.

Therefore, I thought I'd ask the community about their preferred modern Rust dev environments so I can learn more and set-up my env and also contribute it to the wiki.

**tl;dr:**  

* What is your recommended Rust set-up in 2020?   
* How would one go about setting it up?  

I'll curate the responses and contribute them to the wiki.
## [11][[post] The problems with implementing an async Mutex](https://www.reddit.com/r/rust/comments/ge9qqp/post_the_problems_with_implementing_an_async_mutex/)
- url: https://github.com/Diggsey/posts/tree/master/async-mutexes
---

## [12][SuckIT, a fast, multithreaded website downloader](https://www.reddit.com/r/rust/comments/gdwuat/suckit_a_fast_multithreaded_website_downloader/)
- url: https://www.reddit.com/r/rust/comments/gdwuat/suckit_a_fast_multithreaded_website_downloader/
---
## 

https://preview.redd.it/8bc45135fyw41.png?width=640&amp;format=png&amp;auto=webp&amp;s=a9d830daf5959817be553e09b91e13727221f020

## Introduction

[SuckIT](https://github.com/skallwar/suckit) is a multithreaded, open source web downloader written in Rust. It aims to recursively download webpages and allow offline browsing. Offered by u/Skallwar and u/Arcanin14

## Benchmark

As of right now, it's a *little* faster (about 3460%) than HTTrack on a single core, 60 second run on [http://books.toscrape.com/](http://books.toscrape.com/)

|name|pages downloaded|
|:-|:-|
|suckit|2422|
|httrack --disable-security-limits|70|

[http://books.toscrape.com/](http://books.toscrape.com/) is downloaded in 75s for 1 thread, 37s for 2 threads and 19s for 4 threads. It's more or less a linear time reduction

## Future

Some features are missing, such as a random delay between downloads (to avoid IP ban) or user agent support but they will come soon.

Any feedback, reviews or PRs are welcome !

Enjoy :)
