# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (8/2020)!](https://www.reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here_82020/)
- url: https://www.reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here_82020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (8/2020)?](https://www.reddit.com/r/rust/comments/f541u0/whats_everyone_working_on_this_week_82020/)
- url: https://www.reddit.com/r/rust/comments/f541u0/whats_everyone_working_on_this_week_82020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-8-2020/38329?u=llogiq)!
## [3][Why does Rust infer `()` when the generic type cannot be really inferred?](https://www.reddit.com/r/rust/comments/f5nw1o/why_does_rust_infer_when_the_generic_type_cannot/)
- url: https://www.reddit.com/r/rust/comments/f5nw1o/why_does_rust_infer_when_the_generic_type_cannot/
---
     #![feature(type_name_of_val)]
    
    fn foo&lt;T: Default&gt;() -&gt; Result&lt;T, String&gt; {
        Ok(Default::default())
    }
    
    fn bar() -&gt; Result&lt;String, String&gt; {
        let baz = foo()?; // WTF is type of baz?
        println!("baz: {:?}", std::any::type_name_of_val(&amp;baz));
        Ok(String::default())
    }
    
    fn main() {
      let _ = bar();
    }

[https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=fee0ee95cba86cc77ba001f7a16fcf7f](https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=fee0ee95cba86cc77ba001f7a16fcf7f)

I would expect the above to cause compile error but surprisingly, it compiles and the type of baz is `()`. Is there any rule that allows the compiler to do so or it is a bug?
## [4][Source code for all examples from Rust in Action, book published by Manning](https://www.reddit.com/r/rust/comments/f5g7gq/source_code_for_all_examples_from_rust_in_action/)
- url: https://github.com/rust-in-action/code/
---

## [5][Rust support for openapi/swagger](https://www.reddit.com/r/rust/comments/f5m80w/rust_support_for_openapiswagger/)
- url: https://www.reddit.com/r/rust/comments/f5m80w/rust_support_for_openapiswagger/
---
Hey guys, 
As I continue to learn rust here I keep wanting to see what’s in the eco system of like packages I’ve been able to use in languages like go. 

Was wonder what’s in rust that support openapi/swagger 3.0? Be great if there is a contract drive approach where the spec is defined first and some level of codegen then happens. Anyway love to know the best way to get swagger support for a future public api’s I’m sure to build. 
Thanks
## [6][UniverCity open source release (university management sim in Rust)](https://www.reddit.com/r/rust/comments/f5fiii/univercity_open_source_release_university/)
- url: https://github.com/Thinkofname/UniverCity
---

## [7][rust-analyzer Changelog #12](https://www.reddit.com/r/rust/comments/f5a0ku/rustanalyzer_changelog_12/)
- url: https://rust-analyzer.github.io/thisweek/2020/02/17/changelog-12.html
---

## [8][Create a SVG donut chart in Rust WebAssembly (rustwasm)](https://www.reddit.com/r/rust/comments/f5ibv4/create_a_svg_donut_chart_in_rust_webassembly/)
- url: https://www.webassemblyman.com/rustwasm/rustwasm_svg_donut_chart_webcomponent.html
---

## [9][Rust's runtime removes CLI quotes?](https://www.reddit.com/r/rust/comments/f5r7ld/rusts_runtime_removes_cli_quotes/)
- url: https://www.reddit.com/r/rust/comments/f5r7ld/rusts_runtime_removes_cli_quotes/
---
I've got this simple program:

    fn main() {
        std::env::args().enumerate().for_each(|(num, arg)| {
            println!("arg {} is `{}`", num, arg);
        });
    }

When I run it (**not** via `cmd.exe`) with the args `'See one coronation' "and you've seen them all"` (confirmed with ProcessExplorer that the double quotes are indeed passed to the program), it will print

    arg 0 is `&lt;skipped&gt;.exe`
    arg 1 is `'See`
    arg 2 is `one`
    arg 3 is `coronation'`
    arg 4 is `and you've seen them all`

Is it Rust removing the double quotes and treating the contents as a single arg, or is that Windows runtime?

I really just need to pass and interpret all and any symbols in CLI args as is.

PS. Somebody told me it's possible to get the as-is CLI args in Windows via `GetCommandLineW`, so I guess it must be Rust's stdlib? Would it be possible without too much pain to access the same info from Rust?
## [10][A Pragmatic Approach To Global State](https://www.reddit.com/r/rust/comments/f5a2ld/a_pragmatic_approach_to_global_state/)
- url: http://adventures.michaelfbryan.com/posts/pragmatic-global-state/?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=pragmatic-global-state
---

## [11][Weird memory unsafety problem in safe Rust](https://www.reddit.com/r/rust/comments/f51fem/weird_memory_unsafety_problem_in_safe_rust/)
- url: https://www.reddit.com/r/rust/comments/f51fem/weird_memory_unsafety_problem_in_safe_rust/
---
Recently, I updated Rust to the latest stable version (1.41.0) on my macOS system and noticed that the tests of my [pet project](https://github.com/dfyz/osm-renderer) started crashing with `capacity overflow` errors coming from `alloc::raw_vec::capacity_overflow()`. My gut reaction was that I must have been bitten by UB in my unsafe code, but surprisingly this doesn't seem to be the case here.

I managed to reproduce the problem with [a tiny program](https://github.com/dfyz/rust-segfault) without any unsafe code at all (and no external crates). When I run it with `cargo run --release` on either Ubuntu 19.10 or macOS 10.15, I get something like this:

```
$ cargo run --release
[...]
0 0 3 18446744073709551615
[1]    19196 segmentation fault  cargo run --release
```

Basically, a slice obtained from an array of small length somehow has length `2**64 - 1`, which allows you to read from memory you don't have access to without bounds checking. The expected output is something like this:
```
[...]
0 0 3 0
thread 'main' panicked at 'index out of bounds: the len is 0 but the index is 16777216', src/main.rs:13:35
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace.
```

The problem goes away if I do any of the following:
  * remove either of the two `do_test(...);` calls in `main()`;
  * remove the `for _ in 0..1 {` loop;
  * replace the `for y in 0..x {` loop with `for y in 0..1 {`;
  * remove the `z.extend(std::iter::repeat(0).take(x));` line or replace it with `z.extend(std::iter::repeat(0).take(1));`;
  * replace the `for arr_ref in arr {` loop with `let arr_ref = &amp;arr[0];`;
  * specify `RUSTFLAGS="-C opt-level=2"` (the default optimization level for release builds is 3);
  * specify `RUSTFLAGS="-C codegen-units=1"` (the default is 16).

I can't reproduce the problem:
  * in the Rust playground, presumably because it uses `codegen-units = 1` in its Cargo.toml;
  * on a Windows 10 machine (I have no idea why).

`cargo-bisect-rustc` says the regression first happened in the `2019-12-12` nightly, specifically [in this commit](https://github.com/rust-lang/rust/commit/033662dfbca088937b9cdfd3d9584015b5e375b2), which is a rollup of multiple commits. Neither of them seems related to the problem I'm seeing.

My best guess so far is `-C opt-level=3` enables a problematic optimization pass in LLVM, which results in miscompilation. This is corroborated by the fact that MIR (`--emit mir`) and LLVM IR before optimizations (`--emit llvm-ir -C no-prepopulate-passes`) is the same for both `-C opt-level=2` and `-C opt-level=3`.

If this indeed is LLVM acting up, I guess I should follow [this guide](https://github.com/rust-lang/rustc-guide/blob/master/src/codegen/debugging.md). This seems doable but time-consuming, given that the problem only appears with multiple codegen units. So, is there a chance that I stumbled upon a known issue? I searched through the GitHub issues, but failed to find anything similar.
## [12][What rustfmt extension are you using with rust-analyser?](https://www.reddit.com/r/rust/comments/f5c8lc/what_rustfmt_extension_are_you_using_with/)
- url: https://www.reddit.com/r/rust/comments/f5c8lc/what_rustfmt_extension_are_you_using_with/
---
I'm loving rust-analyser, but using its extension doesn't seem to give me auto-format with rustfmt.

Is it likely I have something set up incorrectly, or is there a stand alone vscode extension folks are using to get that feature?
