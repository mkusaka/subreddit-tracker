# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (5/2020)!](https://www.reddit.com/r/rust/comments/eupam3/hey_rustaceans_got_an_easy_question_ask_here_52020/)
- url: https://www.reddit.com/r/rust/comments/eupam3/hey_rustaceans_got_an_easy_question_ask_here_52020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/eransa/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 323](https://www.reddit.com/r/rust/comments/ew3gri/this_week_in_rust_323/)
- url: https://this-week-in-rust.org/blog/2020/01/28/this-week-in-rust-323/
---

## [3][Maybe a dumb question, but the page says "150 of 381 crates". How do I find the rest of them? There's no next page link.](https://www.reddit.com/r/rust/comments/ex0t6b/maybe_a_dumb_question_but_the_page_says_150_of/)
- url: https://lib.rs/game-engines
---

## [4][Difference among Deref, Borrow, and AsRef](https://www.reddit.com/r/rust/comments/ex50zg/difference_among_deref_borrow_and_asref/)
- url: https://www.reddit.com/r/rust/comments/ex50zg/difference_among_deref_borrow_and_asref/
---
My impression is that `Borrow&lt;T&gt;` has the same semantics as `Deref&lt;Target=T&gt;` except for the operator overloading part, and `AsRef&lt;T&gt;` has no special semantic requirement.
## [5][Today the Veloren team is proud to announce version 0.5 of the open-source voxel RPG, featuring many new features and dozens of improvements, big and small.](https://www.reddit.com/r/rust/comments/ewt0md/today_the_veloren_team_is_proud_to_announce/)
- url: https://veloren.net/
---

## [6][Nestur: NES emulator in Rust](https://www.reddit.com/r/rust/comments/ewqiuh/nestur_nes_emulator_in_rust/)
- url: https://github.com/spieglt/nestur
---

## [7][Implemented From&lt;SomeT&gt; for MyT. Do I have to implement &amp;SomeT, &amp;mut SomeT manually?](https://www.reddit.com/r/rust/comments/ex3yzw/implemented_fromsomet_for_myt_do_i_have_to/)
- url: https://www.reddit.com/r/rust/comments/ex3yzw/implemented_fromsomet_for_myt_do_i_have_to/
---
Hey guys,

I implemented the `From` trait for my `Session` type like this:

    impl From&lt;Connection&gt; for Session {
        fn from(con: &amp;Connection) -&gt; Result&lt;Self, Error&gt; {
            Session::restore(&amp;con.hostname, &amp;con.session_id)
        }
    }

Creating a `Session` from a `Connection` works as expected but when I try to create a `Session` from a `&amp;mut Connection` I get the message, that the `From` trait is not implemented.

Of course I can just implement `From&lt;&amp;mut Connection&gt; for Session` manually to work around this but I guess there is a method to automatically implement `From` for `&amp;T` and `&amp;mut T` when I implement From for `T`

Thank you in advance :)
## [8][Chunked Http post request with reqwest?](https://www.reddit.com/r/rust/comments/ex4sdl/chunked_http_post_request_with_reqwest/)
- url: https://www.reddit.com/r/rust/comments/ex4sdl/chunked_http_post_request_with_reqwest/
---
I've been looking to send an audio via http in a chunked way with reqwest (mostly because I'm sending something recorded on the moment and I want it to be sent ASAP), so far I haven't seen any way, anyone knows how to do this?
## [9][Writing an OS in Rust: Updates in January 2020](https://www.reddit.com/r/rust/comments/ex654n/writing_an_os_in_rust_updates_in_january_2020/)
- url: https://os.phil-opp.com/status-update/2020-02-01/
---

## [10][Better way to map a result to non-result? Like a question mark operator for non results?](https://www.reddit.com/r/rust/comments/ex0yn8/better_way_to_map_a_result_to_nonresult_like_a/)
- url: https://www.reddit.com/r/rust/comments/ex0yn8/better_way_to_map_a_result_to_nonresult_like_a/
---
Hey all, I have a pattern a lot that results in some ugly code. You can see an example of it in this rust playground:

https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=f4b9a9b1429a6acdef65dbfe035ddd4d

Basically, when you have fn's that return errors, and inside your function you call functions that have errors, the `?` operator is the best thing ever. But what if that outer function does NOT return an error? Is there a way to map the error to a value and return that if there was an error? Thanks!
## [11][Build your own executor](https://www.reddit.com/r/rust/comments/ewodxc/build_your_own_executor/)
- url: https://stjepang.github.io/2020/01/31/build-your-own-executor.html
---

## [12][How to perform square root of a generic type](https://www.reddit.com/r/rust/comments/ewy548/how_to_perform_square_root_of_a_generic_type/)
- url: https://www.reddit.com/r/rust/comments/ewy548/how_to_perform_square_root_of_a_generic_type/
---
Hello,  im currently making a vector crate (for learning purposes) and I came across a problem. I'm using generics for the type that the vector holds,

`pub struct Evec2&lt;T&gt;{pub x: T, pub y: T}`

 and it makes it impossible to perform the square root of the vector's components, because the square root is not a trait defined like other operators (std :: ops). can someone tell me some trick to work around this
