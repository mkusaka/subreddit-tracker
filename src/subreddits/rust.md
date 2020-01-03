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

## [3][Announcing AeroRust - The Unofficial Working Group For Rust in Aerospace](https://www.reddit.com/r/rust/comments/ejdv7w/announcing_aerorust_the_unofficial_working_group/)
- url: https://www.reddit.com/r/rust/comments/ejdv7w/announcing_aerorust_the_unofficial_working_group/
---
Hi Folks,

We are happy to announce the formation of AeroRust Unofficial Working Group. Lately the Aerospace industry is picking up speed in commercialisation and even faster and rapid development in the area.

This Working group is aiming to help push the Open-source community  more into the growing Aerospace industry, by providing information,  materials, tools, crates and etc. to hobbyists and the industry.  


If you are interested, come join us:   


* Github: [https://github.com/AeroRust/Welcome](https://github.com/AeroRust/Welcome)
* Discord: [https://discord.gg/RXNsMXc](https://discord.gg/RXNsMXc)
* Resources: [https://github.com/AeroRust/awesome-space](https://github.com/AeroRust/awesome-space)
## [4][Nvim-rs: Rust library for neovim clients](https://www.reddit.com/r/rust/comments/ejfceo/nvimrs_rust_library_for_neovim_clients/)
- url: /r/neovim/comments/eiy5if/nvimrs_rust_library_for_neovim_clients/
---

## [5][Updating Pernosco to Rust futures 0.3](https://www.reddit.com/r/rust/comments/ejd1wg/updating_pernosco_to_rust_futures_03/)
- url: https://robert.ocallahan.org/2020/01/updating-pernosco-to-rust-futures-03.html?m=1
---

## [6][Update on const generics progress](https://www.reddit.com/r/rust/comments/ej10bh/update_on_const_generics_progress/)
- url: https://github.com/rust-lang/rust/issues/44580#issuecomment-570191702
---

## [7][Podcast: Using Rust for Machine Learning in Search at Etsy](https://www.reddit.com/r/rust/comments/ej7fqb/podcast_using_rust_for_machine_learning_in_search/)
- url: https://changelog.com/practicalai/70
---

## [8][New to rust and loving it. Just published my first crate and would love feedback!](https://www.reddit.com/r/rust/comments/ejevh1/new_to_rust_and_loving_it_just_published_my_first/)
- url: https://github.com/M-J-Hooper/playing-cards/
---

## [9][I'm not feeling the async pressure [the importance of supporting backpressure in async ecosystems]](https://www.reddit.com/r/rust/comments/eiz65y/im_not_feeling_the_async_pressure_the_importance/)
- url: https://lucumr.pocoo.org/2020/1/1/async-pressure/
---

## [10][Adding and Subtracting *](https://www.reddit.com/r/rust/comments/ejb02a/adding_and_subtracting/)
- url: https://www.reddit.com/r/rust/comments/ejb02a/adding_and_subtracting/
---
Full disclosure: this post is probably just as much about my not understanding C++ as it is not understanding Rust. I am porting a C++ function with the following structure: 

```
foo(const uint8_t* src, uint8_t* dst) {

int m = 2;
uint8 n = 3; 

loop {
    // change src, dst according incrementing variables not shown 
    // as examples
    // src += m
    // dst += n
}


memcpy(src, dst,some_num);

}
```

In Rust, I do the following: 

```Rust

fn foo( src: *mut u8, dst: *mut u8) {

// dummy variables for addition below
let n: u8; 
let m: u8; 

loop {
    // increment/decrement src, dst according to loop variables not shown 
    // as examples
    // src += m
    // dst += n
}
std::ptr::copy_nonoverlapping(src, dst, some_num)

}
```

When I try to src/dst += in the Rust for loop, I get type errors (some variation of cannot cast between *-ptr and u8). Cannot chane type signature since those types are needed in the calls to the copy fns. This means I do not know how to +/- when using * int type in Rust. Anyone able to help? If this is too much of a beginner question, happy to move to the appropriate thread.
## [11][Announcing postgres-query: write and execute SQL queries with ease!](https://www.reddit.com/r/rust/comments/ej3thz/announcing_postgresquery_write_and_execute_sql/)
- url: https://github.com/nolanderc/rust-postgres-query
---

## [12][Best error-handling practices - Jan 2020](https://www.reddit.com/r/rust/comments/ej67aa/best_errorhandling_practices_jan_2020/)
- url: https://www.reddit.com/r/rust/comments/ej67aa/best_errorhandling_practices_jan_2020/
---
I found this great survey of error-handling methods here:
https://blog.yoshuawuyts.com/error-handling-survey/

Is there any consensus on best (currently supported) error-handling libraries / practices for:

1) Applications
2) Libraries

?

At the very least, are there any approaches which are now wholly superseded by subsequent solutions?

I recognize this is probably a bit controversial, but it’s also getting a bit overwhelming imho.
