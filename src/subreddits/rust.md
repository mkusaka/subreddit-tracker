# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (25/2020)!](https://www.reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 343](https://www.reddit.com/r/rust/comments/hactqu/this_week_in_rust_343/)
- url: https://this-week-in-rust.org/blog/2020/06/16/this-week-in-rust-343/
---

## [3][String vs &amp;str](https://www.reddit.com/r/rust/comments/hciaol/string_vs_str/)
- url: https://www.reddit.com/r/rust/comments/hciaol/string_vs_str/
---
As a Rust newbie struggling with understanding the basics, I found this article really good at explaining the difference between String and &amp;str in Rust

[https://blog.thoughtram.io/string-vs-str-in-rust/](https://blog.thoughtram.io/string-vs-str-in-rust/)
## [4][Show /r/rust: libps1 a shell prompt for developers](https://www.reddit.com/r/rust/comments/hcfy2d/show_rrust_libps1_a_shell_prompt_for_developers/)
- url: https://github.com/JoshMcguigan/libps1
---

## [5][(Stream) C++ programmer and game developer starts learning Rust this week](https://www.reddit.com/r/rust/comments/hchmsz/stream_c_programmer_and_game_developer_starts/)
- url: https://www.youtube.com/watch?v=1asufxmZW5A
---

## [6][CS196 at Illinois is Teaching Rust Concepts and Sharing on YT](https://www.reddit.com/r/rust/comments/hc1c52/cs196_at_illinois_is_teaching_rust_concepts_and/)
- url: https://www.youtube.com/channel/UCRA18QWPzB7FYVyg0WFKC6g/videos
---

## [7][Tour of Rust - Chapter 8 Released - Smart Pointers - Occidental translations, French translations](https://www.reddit.com/r/rust/comments/hccnbk/tour_of_rust_chapter_8_released_smart_pointers/)
- url: https://tourofrust.com/chapter_8_en.html
---

## [8][I made a pandemic-inspired data structure to practice what I've learned in Rust](https://www.reddit.com/r/rust/comments/hcakf2/i_made_a_pandemicinspired_data_structure_to/)
- url: https://www.reddit.com/r/rust/comments/hcakf2/i_made_a_pandemicinspired_data_structure_to/
---
Need another way to store values that has some unpredictability? Do you agree that Rust just isn't safe when it comes to memory management? Want another 4-letter acronym to keep track of?

Introducing [OIOO](https://github.com/ramirezmike/OIOO) (one-in, one-out), an almost useless data structure I made because I thought it would be fun and good practice. Here's how it works:

    // create an essential oioo with 25% capacity
    let mut oioo = OIOO::&lt;usize&gt;::new(
        Phase::One { 
            occupancy: 12, 
            is_essential: true 
        }
    ); 
    
    // add values with one_in()
    oioo.one_in(10); 
    oioo.one_in(20);
    oioo.one_in(30);
    
    // randomly returns a value with one_out()
    println!("{}", oioo.one_out().unwrap()); 

There's a hard limit based on the initialized occupancy, phase and essentialism, but don't worry because any excess values get stored in a queue outside, which automatically inserts into the main store once room becomes available.

    // add values in an oioo with a max capacity of 3
    oioo.one_in(10); 
    oioo.one_in(20);
    oioo.one_in(30);
    
    // add a value to exceed capacity. 
    // This value gets stored in a separate queue
    oioo.one_in(40); 

    // randomly returns a value from the set [10, 20, 30]
    println!("{}", oioo.one_out().unwrap()); // let's say this outputs 20

    // randomly returns a value from the set [10, 30, 40]
    println!("{}", oioo.one_out().unwrap()); // can potentially return 40

And you don't have to worry about the safety of the memory stored in the OIOO because [they are padded with social distance None values.](https://github.com/ramirezmike/OIOO/raw/master/example.jpg?raw=true). That's right, each value will have plenty of space to stay safe while in the OIOO!

[Further documentation can be found here](https://ramirezmike.github.io/OIOO/target/doc/oioo/) Feedback and criticism is totally welcome! 

^^^^im ^^^^so ^^^^sorry, ^^^^please ^^^^don't ^^^^hurt ^^^^me
## [9][Behold, my first Rust application. It's a CLI that opens Windows paths on Linux to a shared drive. If you can spot some mistakes I made, don't hesitate to tell, I came from PHP.](https://www.reddit.com/r/rust/comments/hciydg/behold_my_first_rust_application_its_a_cli_that/)
- url: https://github.com/ivkovici/sp
---

## [10][I built a fast BLAS-accelerated, Pytorch-style tensor+autograd library in Rust](https://www.reddit.com/r/rust/comments/hcby3r/i_built_a_fast_blasaccelerated_pytorchstyle/)
- url: https://www.reddit.com/r/rust/comments/hcby3r/i_built_a_fast_blasaccelerated_pytorchstyle/
---
https://github.com/bkkaggle/L2/blob/master/screenshot.png

GitHub: https://github.com/bkkaggle/L2

Crates.io: https://crates.io/crates/l2

Twitter thread with images: https://twitter.com/stormtroper1721/status/1273824371056902145

I benchmarked matrix multiplication on two 256x256 matrices, Pytorch (on my MacBook pro's cpu) takes 10us compared to l2 which takes 50us. So it's about 5x slower but I didn't explicitly try too hard to make it super-fast.

I made l2 to get better at using Rust and to learn more about how libraries like Pytorch and Tensorflow work behind the scenes, so don't expect this library to be production-ready :)

I also only picked up rust a few months ago so I’d really appreciate feedback on how to make it more better/more idiomatic :)
## [11][RLS does not provide suggestions for reqwest crate](https://www.reddit.com/r/rust/comments/hcg4ft/rls_does_not_provide_suggestions_for_reqwest_crate/)
- url: https://www.reddit.com/r/rust/comments/hcg4ft/rls_does_not_provide_suggestions_for_reqwest_crate/
---
Is any body having issues with code suggestions for external crates for public methods?

I created an example repo where the issue occurs for the reqwest crate

[https://github.com/ta32/issues/tree/master/rust\_requests](https://github.com/ta32/issues/tree/master/rust_requests)

CLion with the rust plugin can find the definition of structs like Client, but it cannot suggest public methods. This is not happening for other projects that use different crates.

Is it an issue with this crate or an environmental issue - does anyone else experience this?

edit: Clion Rust does not use RLS - i thought it was. 
## [12][smallvec 1.4.0 wont compile on my new laptop but still compiles on my other computer?](https://www.reddit.com/r/rust/comments/hck5ej/smallvec_140_wont_compile_on_my_new_laptop_but/)
- url: https://www.reddit.com/r/rust/comments/hck5ej/smallvec_140_wont_compile_on_my_new_laptop_but/
---
Hi,

So I just got a new laptop and have been setting everything up.

I noticed something very… strange. Some of my projects won't compile anymore because my dependencies have dependencies on `smallvec 1.4.0`. For whatever reason, this compiles just fine on my old computer.

However, on my new computer, I get the following error on the *same project*:

	error[E0554]: `#![feature]` may not be used on the stable release channel
	  --&gt; /Users/xxxxx/.cargo/registry/src/github.com-1ecc6299db9ec823/smallvec-1.4.0/lib.rs:61:32
	   |
	61 | #![cfg_attr(feature = "union", feature(untagged_unions))]
	   |                                ^^^^^^^^^^^^^^^^^^^^^^^^


What the hell? The projects are identical. They're being synced with dropbox, so I know they're identical. Both computers are using the latest stable version of rust (according to rustup, 1.44.1 stable and 1.46.0 nightly).

So how is this possible?

**Bonus question:** Because rustfmt isn't available on the latest version of nightly, when I program on a nightly project rustfmt doesn't work. How can I set it to use the stable rustfmt on this nightly project?
