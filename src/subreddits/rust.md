# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (13/2020)!](https://www.reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 331](https://www.reddit.com/r/rust/comments/fp9z5t/this_week_in_rust_331/)
- url: https://this-week-in-rust.org/blog/2020/03/24/this-week-in-rust-331/
---

## [3][Goodbye, docs team - Inside Rust](https://www.reddit.com/r/rust/comments/fq7znr/goodbye_docs_team_inside_rust/)
- url: https://blog.rust-lang.org/inside-rust/2020/03/27/goodbye-docs-team.html
---

## [4][A video to help you climb one of Rust’s highest walls... Lifetimes! 🎥](https://www.reddit.com/r/rust/comments/fqeq2v/a_video_to_help_you_climb_one_of_rusts_highest/)
- url: https://youtu.be/tyuDitqoZbM
---

## [5][Writing an OS in Rust: Async/Await](https://www.reddit.com/r/rust/comments/fq083y/writing_an_os_in_rust_asyncawait/)
- url: https://os.phil-opp.com/async-await/
---

## [6][Open sourced my tuntap crate](https://www.reddit.com/r/rust/comments/fqgplt/open_sourced_my_tuntap_crate/)
- url: https://github.com/siegfried/utuntap
---

## [7][Creating 50 stock prices that change over time using threads.](https://www.reddit.com/r/rust/comments/fqjljk/creating_50_stock_prices_that_change_over_time/)
- url: https://www.reddit.com/r/rust/comments/fqjljk/creating_50_stock_prices_that_change_over_time/
---
 I am creating a program which simulates 50 stocks being updated.  Each stock is monitored (incremented) by a thread, updating a listener thread of any change. The listener thread receives the update and prints the value of the stock before waiting for the next update. 

So far what I did was create a threadpool that generates 50 random stock names with their respective random pricing. I send the stock prices and stock names over a channel to the next threapool which is responsible for incrementing the prices with a random value and the thread sleeps with a random sleeptime. Then i send these values to an infinite loop which prints the values out. 

Need advice on whether I am doing it the right way. Here is the [link](https://github.com/yudhiesh1997/50-Stocks-Rust/blob/master/stock/src/main.rs) to my code on Github.
## [8][Nalgebra is incredible and plays nice with wasm](https://www.reddit.com/r/rust/comments/fpvc91/nalgebra_is_incredible_and_plays_nice_with_wasm/)
- url: https://www.reddit.com/r/rust/comments/fpvc91/nalgebra_is_incredible_and_plays_nice_with_wasm/
---
I just wanted to let people know about this crate,and the fantastic experience I had using it. 

I leveraged it to compute the best fit line for a point cloud in my react front end packed with parcel. 

The runtime is absurdly good. The whole cycle takes less than 23 microseconds. I haven't calculated the exact overhead, but it's basically nothing.
## [9][Cross-compile Rust programs to run on Turris Omnia](https://www.reddit.com/r/rust/comments/fqjot7/crosscompile_rust_programs_to_run_on_turris_omnia/)
- url: https://medium.com//cross-compile-rust-programs-to-run-on-turris-omnia-e592b555e2aa?source=friends_link&amp;sk=1f8c6b55e04d67482657e1a8a1e21489
---

## [10][The Book versions](https://www.reddit.com/r/rust/comments/fqiwwt/the_book_versions/)
- url: https://www.reddit.com/r/rust/comments/fqiwwt/the_book_versions/
---
Hi everyone,

I have a problem with understanding the book versioning. As far as i researched there are stable, beta and nightly and also NoStarch first-edition, second-edition and 2018-edition.  
I'm not a fan of reading on screen so i face two options, printing one of the stable, beta, nightly or printing one of the versions of the NoStarch (i have low availability in my country and the prices are around 100$).

So which version should i print and what are the differences?

[View Poll](https://www.reddit.com/poll/fqiwwt)
## [11][Paris 1.4. Now with macros and more control!](https://www.reddit.com/r/rust/comments/fq1ol9/paris_14_now_with_macros_and_more_control/)
- url: https://github.com/0x20F/logger
---

## [12][What will happen if I will Pin something inside of function that returns never?](https://www.reddit.com/r/rust/comments/fqiipy/what_will_happen_if_i_will_pin_something_inside/)
- url: https://www.reddit.com/r/rust/comments/fqiipy/what_will_happen_if_i_will_pin_something_inside/
---
And how actually never works?
