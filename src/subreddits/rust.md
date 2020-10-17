# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (42/2020)!](https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 360](https://www.reddit.com/r/rust/comments/jbh0ci/this_week_in_rust_360/)
- url: https://this-week-in-rust.org/blog/2020/10/14/this-week-in-rust-360/
---

## [3][Announcement: xshell, ergonomic "bash" scripting in Rust](https://www.reddit.com/r/rust/comments/jctkpi/announcement_xshell_ergonomic_bash_scripting_in/)
- url: https://docs.rs/xshell/0.1.2/xshell/index.html
---

## [4][How to Write Hygienic Rust Macros](https://www.reddit.com/r/rust/comments/jcpowx/how_to_write_hygienic_rust_macros/)
- url: https://gist.github.com/Koxiaet/8c05ebd4e0e9347eb05f265dfb7252e1
---

## [5][Learn Rust With Benford's Law!](https://www.reddit.com/r/rust/comments/jcs1u2/learn_rust_with_benfords_law/)
- url: https://gliderkite.github.io/posts/learn-rust-with-benford/
---

## [6][Should Result::into_iter and Option::into_iter be considered harmful?](https://www.reddit.com/r/rust/comments/jcrijb/should_resultinto_iter_and_optioninto_iter_be/)
- url: https://www.reddit.com/r/rust/comments/jcrijb/should_resultinto_iter_and_optioninto_iter_be/
---
I feel that canonical iterator function `into_iter` which is implemented for both `Result` and `Option` puts developer to a position where they are literally one character away from introducing a serious bug in their code, and should be considered harmful.

I don't disagree that being able to iterate through `Result`/`Option` may be useful in some cases. But the fact that current implementation of this is being automatically picked up by the `for` loop is a really dangerous feature. If developer is not being careful, they may end up iterating over `Result`/`Option` as iterator where their intention was to iterate over underlying content of `Result`/`Option` which may be also implementing an iterator.

For example, this code looks like it should be printing elements one by one, but it's not:

```
    fn get_elements() -&gt; Result&lt;Vec&lt;i32&gt;, Box&lt;dyn Error&gt;&gt; {
        Ok(vec![1, 2, 3])
    }

    for n in get_elements() {
        println!("Element: {:?}", n);
    }
```

Correct version of this loop should look something like:

```
    for n in get_elements()? {
        ...
```

What former code actually does is: if `get_elements()` returns some vector, then loop will run one iteration and print entire vector to the console. If `get_elements` returns an error, then this loop will just skip and do nothing, just silently ignoring the error.

Compiler has little to offer to mitigate this: both of versions of the code are valid and compile without issues.

If developer is using fancy IDE which highlights types like VSCode, they may notice that wrong time was inferred for the loop variable `n: Vec&lt;i32&gt;` instead of `n: i32`, but that won't help if for whatever reason they are ignoring elements in the loop:

```
    for _ in get_elements() {
        println!("one more element");
    }
```

What's are your thoughts? Am I missing some tools/features that can help to detect and prevent this kind of errors?

https://repl.it/repls/InfatuatedZestyObservatory
## [7][The entire minecraft protocol in Rust using macro magic](https://www.reddit.com/r/rust/comments/jcbawo/the_entire_minecraft_protocol_in_rust_using_macro/)
- url: https://www.reddit.com/r/rust/comments/jcbawo/the_entire_minecraft_protocol_in_rust_using_macro/
---
For the last three weeks I've been working on this: an implementation of the Minecraft multiplayer network protocol in Rust. I want to deserialize bytes off the wire into structs and serialize structs into bytes.

I've tried implementing this in other languages but never had a language powerful enough to complete the project in only 3 weeks.

Here's my code, I hope you enjoy:

* mcproto-rs: the protocol itself [https://github.com/Twister915/mcproto-rs](https://github.com/Twister915/mcproto-rs)
* mctokio: tokio I/O stuff [https://github.com/Twister915/mctokio](https://github.com/Twister915/mctokio)
* rustcord: a layer 7 server-switching proxy implementation (WIP): [https://github.com/Twister915/rustcord](https://github.com/Twister915/rustcord)

Oh, and here's the macro magic: [https://github.com/Twister915/mcproto-rs/blob/master/src/v1\_15\_2.rs](https://github.com/Twister915/mcproto-rs/blob/master/src/v1_15_2.rs)
## [8][When reproducible builds?](https://www.reddit.com/r/rust/comments/jct0y4/when_reproducible_builds/)
- url: https://www.reddit.com/r/rust/comments/jct0y4/when_reproducible_builds/
---
In some domains, it is very important to be able to reproducibly rebuild a binary from the sources. This is cargo/rustc pretty bad at. Is there any movement on this issue?

I only have this link (from 2017...) as some sort of reference: [https://users.rust-lang.org/t/testing-out-reproducible-builds/9758](https://users.rust-lang.org/t/testing-out-reproducible-builds/9758)
## [9][Introducing an Event Emitter in rust](https://www.reddit.com/r/rust/comments/jcr42j/introducing_an_event_emitter_in_rust/)
- url: https://github.com/Dylan-Kerler/event_emitter_rs
---

## [10][Are out parameters idiomatic in Rust?](https://www.reddit.com/r/rust/comments/jcdg2e/are_out_parameters_idiomatic_in_rust/)
- url: https://steveklabnik.com/writing/are-out-parameters-idiomatic-in-rust
---

## [11][lfs: a small tool to list filesystems](https://www.reddit.com/r/rust/comments/jcugvo/lfs_a_small_tool_to_list_filesystems/)
- url: https://github.com/Canop/lfs
---

## [12][Blackbody: a thermogram viewer written in Rust, also introducing a new file format parser in Rust](https://www.reddit.com/r/rust/comments/jcu81p/blackbody_a_thermogram_viewer_written_in_rust/)
- url: https://bitbucket.org/nimmerwoner/blackbody/src/master/
---

