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
## [2][This Week in Rust 337](https://www.reddit.com/r/rust/comments/geagy0/this_week_in_rust_337/)
- url: https://this-week-in-rust.org/blog/2020/05/05/this-week-in-rust-337/
---

## [3][Criticisms of rust](https://www.reddit.com/r/rust/comments/ggyo51/criticisms_of_rust/)
- url: https://www.reddit.com/r/rust/comments/ggyo51/criticisms_of_rust/
---
Rust is on my list of things to try and I have read mostly only good things about it. I want to know about downsides also, before trying. Since I have heard learning curve will be steep.

compared to other languages like Go, I don't know how much adoption rust has. But apparently languages like go and swift get quite a lot of criticism. in fact there is a [github repo](https://github.com/ksimka/go-is-not-good) to collect criticisms of Go.

Are there well written (read: not emotional rant) criticisms of rust language? Collecting them might be a benefit to rust community as well.
## [4][Clivern/Buzzard - All rust features in one file](https://www.reddit.com/r/rust/comments/ggrwjt/clivernbuzzard_all_rust_features_in_one_file/)
- url: https://github.com/Clivern/Buzzard
---

## [5][Do you actually use variable declaration without initialization.?](https://www.reddit.com/r/rust/comments/ggv2i8/do_you_actually_use_variable_declaration_without/)
- url: https://www.reddit.com/r/rust/comments/ggv2i8/do_you_actually_use_variable_declaration_without/
---
This is valid Rust syntax:
```rust
let x;
if condition {
    x = 5;
} else {
    x = 1;
}
```

However, I find that every time I would need to do this, I just bind the `let` to the `if else` directly: 

```rust
let x = if condition {
    5
} else {
    1
};
```

In my opinion, this is both more succinct and more idiomatic. It also leads me to ask the question: are there other Rust features that are allowed but are not used in favor of something better?
## [6][Unzip function slow](https://www.reddit.com/r/rust/comments/ggy0fa/unzip_function_slow/)
- url: https://www.reddit.com/r/rust/comments/ggy0fa/unzip_function_slow/
---
In an effort to use Rayon more effectively I'm switching over some code to a more functional style, and I noticed an weird regression in performance. With some effort I narrowed it down to the unzip function in the standard library. I created the following benchmark to see the severity of the performance difference:  
    #![feature(test)]
    
    extern crate test;
    use test::{Bencher, black_box};
    
    fn run_functional(l: &amp;Vec&lt;(usize, usize)&gt;) -&gt; (Vec&lt;usize&gt;, Vec&lt;usize&gt;) {
        l.iter().copied().unzip()
    }
    
    fn run_imperative(l: &amp;Vec&lt;(usize, usize)&gt;) -&gt; (Vec&lt;usize&gt;, Vec&lt;usize&gt;) {
        let len = l.len();
        let (mut result1, mut result2) = (Vec::with_capacity(len), Vec::with_capacity(len));
        for item in l.iter().copied() {
            result1.push(item.0);
            result2.push(item.1);
        }
        (result1, result2)
    }
    
    #[bench]
    fn bench_functional(b: &amp;mut Bencher) {
        let list = black_box(vec![(1, 2); 256]);
        b.iter(|| run_functional(&amp;list));
    }
    
    #[bench]
    fn bench_imperative(b: &amp;mut Bencher) {
        let list = black_box(vec![(1, 2); 256]);
        b.iter(|| run_imperative(&amp;list));
    }
This yields the following results:

    test bench_functional ... bench:       1,440 ns/iter (+/- 66)
    test bench_imperative ... bench:         443 ns/iter (+/- 43)

So I dove into the standard library code, and it seems that there are two reasons why the imperative method is so much faster. Firstly, it uses `push` instead of `extend`, which makes a significant difference. The rest of the difference is accounted for by the use of `with_capacity` instead of `new` to create the new vector. This is strange to me, because the `size_hint` function of this iterator returns 256, so Rust knows the size of the resulting vectors it will create. Why is the `with_capacity` function not used and why is `vec.push(a)` so much faster than `vec.extend(Some(a))`?
## [7][Announcing the v0.16 release of Yew!](https://www.reddit.com/r/rust/comments/ggi6ix/announcing_the_v016_release_of_yew/)
- url: https://www.reddit.com/r/rust/comments/ggi6ix/announcing_the_v016_release_of_yew/
---
Hello, I'm excited to share with all of you the [latest Yew release](https://github.com/yewstack/yew/releases/tag/0.16.0) :)

*(If you're not familiar, Yew is a framework for building client web apps with Rust &amp; WebAssembly!)*

It's been awhile since I've shared an update but I'll try to be brief...

**Changes since v0.13**

* Split off `stdweb` support into a convenient alias crate called `yew-stdweb` which allows the main yew crate to default to using `web-sys`.
* Keys can be now be assigned to virtual dom nodes to improve render performance for list-type content.
* Renamed the `mounted(..)` lifecycle callback method to `rendered(..)` which will be called after each time Yew renders your `Component` to the screen.
* Many documentation and project template improvements and updated instructions on how to boot up the many example projects.

Read more: [https://github.com/yewstack/yew/releases](https://github.com/yewstack/yew/releases)

**Other News**

* We moved from Gitter to [Discord](https://discord.gg/VQck8X4) (yay reliable notifications!)
* We [summarized](https://github.com/yewstack/yew/wiki/Dev-Survey-%5BSpring-2020%5D) our [recent dev survey](https://www.reddit.com/r/rust/comments/g43ld4/yew_developer_survey/) and used the results to update our [Roadmap](https://yew.rs/docs/more/roadmap). Thank you to the over 100 devs who took the time to share their thoughts and experiences 🙇‍♂️
* We kickstarted a small bounty program for contributions to Yew on [Issuehunt](https://issuehunt.io/r/yewstack/yew/) (expect to see more funded issues soon!)
* Hooks API / Functional Component development is well under development and is making great progress. Follow along [here](https://github.com/yewstack/yew/projects/3)
* Server Side Rendering is in the planning phase and progress can be tracked [here](https://github.com/yewstack/yew/projects/5). Looking for help on this, please reach out if interested!

Lastly, I want to give a huge thanks to the Yew community which makes this all possible, we had 13 first-time contributors submit bug fixes and features and many others who helped improve docs and maintain examples since the last announcement!
## [8][Blocking - crate to wrap blocking calls for async IO](https://www.reddit.com/r/rust/comments/ggongd/blocking_crate_to_wrap_blocking_calls_for_async_io/)
- url: https://twitter.com/stjepang/status/1259201346642853892
---

## [9][Writing Python inside your Rust code — Part 3](https://www.reddit.com/r/rust/comments/ggqzqr/writing_python_inside_your_rust_code_part_3/)
- url: https://blog.m-ou.se/writing-python-inside-rust-3/
---

## [10][A practical introduction to async programming in Rust](https://www.reddit.com/r/rust/comments/ggdw7p/a_practical_introduction_to_async_programming_in/)
- url: http://jamesmcm.github.io/blog/2020/05/06/a-practical-introduction-to-async-programming-in-rust/#en
---

## [11][New version of doh-client (2.2.0)](https://www.reddit.com/r/rust/comments/ggzqfi/new_version_of_dohclient_220/)
- url: https://www.reddit.com/r/rust/comments/ggzqfi/new_version_of_dohclient_220/
---
Hey Rustaceans,  
I publish a new version of doh-client. The new big feature is:

* Code was refactored to increase the readability
* DNS packets are parser by the [dns-message-parser](https://crates.io/crates/dns-message-parser) library to check if the packets are correct
* If control-cache is not present in the HTTP header then smallest TTL in the answer section is used to cache the DNS packet
* Socks5 proxy can be used to connect to the server(--socks5)


[Here is the Github link](https://github.com/LinkTed/doh-client)  

Thank you for your attention.
## [12][SAD! a rust cli tool that does batch file edits, no need for sed, we have sad.](https://www.reddit.com/r/rust/comments/ggm067/sad_a_rust_cli_tool_that_does_batch_file_edits_no/)
- url: https://www.reddit.com/r/rust/comments/ggm067/sad_a_rust_cli_tool_that_does_batch_file_edits_no/
---
Hi guys!

https://github.com/ms-jpq/sad

I made my first project in rust, and I thought it would be cool to share it here.

What it does, is basically **ergonomic batch file edits**.

`sad` will take a list of files, a pattern, a replacement, and generate a nice diff of BEFORE / AFTER the edits.

You can view the changes **before** you apply them.


`sad` is inspired by my usage of `sd`, another cool rust project.


The difference as being while `sd` is more optimized for instream edits, such as
 

`command1 | sd '&lt;pattern&gt;' '&lt;replacement&gt;' | command2`


`sad` does batch edits like


`find "$FIND_ARGS" | sad '&lt;pattern&gt;' '&lt;replacement&gt;'`


of course, the output of `sad` is much more optimized for its own usecase:

https://raw.githubusercontent.com/ms-jpq/sad/master/previews/preview2.gif
