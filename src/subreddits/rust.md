# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (45/2020)!](https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 363](https://www.reddit.com/r/rust/comments/joxy7n/this_week_in_rust_363/)
- url: https://this-week-in-rust.org/blog/2020/11/04/this-week-in-rust-363/
---

## [3][40 Millisencond Bug](https://www.reddit.com/r/rust/comments/jp23jq/40_millisencond_bug/)
- url: https://vorner.github.io/2020/11/06/40-ms-bug.html
---

## [4][Anu: Version control tool in Rust, written by Pijul developers](https://www.reddit.com/r/rust/comments/joqput/anu_version_control_tool_in_rust_written_by_pijul/)
- url: https://nest.anu.dev/anu/anu
---

## [5][My first contribution to the rust community!](https://www.reddit.com/r/rust/comments/jp3twx/my_first_contribution_to_the_rust_community/)
- url: https://www.reddit.com/r/rust/comments/jp3twx/my_first_contribution_to_the_rust_community/
---
After a month of hard work I am glad to finally have contributed back to the community something I think will be useful for others.

As I was building a calendar application (side project to learn rust) in rust I could not find a good crate for working with recurrence rules, so I decided to try to implement one myself!

[RRule.rs](https://RRule.rs) is a rust implementation of recurrence rules as defined in the iCalendar RFC (more info in README). It is mostly a port of the great [rrule.js](https://github.com/jakubroztocil/rrule) library, but a lot faster!

Repository can be found [here](https://github.com/fmeringdal/rust_rrule) and documentation [here](https://docs.rs/rrule).

Any feedback or contribution is highly appreciated :D
## [6][Question: is boost::units possible on Rust?](https://www.reddit.com/r/rust/comments/jp0xpd/question_is_boostunits_possible_on_rust/)
- url: https://www.reddit.com/r/rust/comments/jp0xpd/question_is_boostunits_possible_on_rust/
---
Hi Rustaceans! I have a bit theoretical question about the Rust type system. Is it possible to implement units system on Rust which will be zero-cost on run-time and check correctness of quantity operations on compile-time? I'm looking at [`Boost::Units`](https://www.boost.org/doc/libs/release/doc/html/boost_units/Units.html) as something I would like to have in Rust-world. The aim is to make it possible to write something like
```rust
Gramm(10) * Cm(3.14) / (Second(3).powi(2)) + Newton(15) * Gramm(1) / Kg(10)
```

For me it looks like at least const-generics are required, but let this question be about current (or near-future) nightly compiler.
## [7][Rust vs Go](https://www.reddit.com/r/rust/comments/jp4orc/rust_vs_go/)
- url: https://bitfieldconsulting.com/golang/rust-vs-go
---

## [8][Rust GameDev Ecosystem Survey](https://www.reddit.com/r/rust/comments/joj5e0/rust_gamedev_ecosystem_survey/)
- url: https://twitter.com/rust_gamedev/status/1324349001706229760
---

## [9][Weird question: What did I just implement?](https://www.reddit.com/r/rust/comments/joyivl/weird_question_what_did_i_just_implement/)
- url: https://www.reddit.com/r/rust/comments/joyivl/weird_question_what_did_i_just_implement/
---
Hi y'all! This is not necessarily a Rust question, but I implemented it in Rust and you are all so helpful so I'm asking here.

I wanted to generate pretty-printed WebAssembly text format code. Something like this:

```
(module
  (import
    "wasi_snapshot_preview1"
    "proc_exit"
    (func $__wasi_snapshot_preview1_proc_exit (param i32))
  )
  (func
    $_start
    (export "_start")
    (local $var i32)
    (local.set $var (i32.const 0))
    (call $__wasi_snapshot_preview1_proc_exit (i32.const 0))
  )
)
```

Basically, I want to be able to control exactly how the instructions get folded into s-expressions. To my knowledge, nothing existed that could do that. So I wrote a proof-of-concept implementation that only supports a few instructions [here](https://github.com/yagehu/wat-ast).

With this library, I can build a data structure in memory and output it to a string of `.wat` code. For an example, see this [test case](https://github.com/yagehu/wat-ast/blob/main/src/elements/module.rs).

My question is: What did I implement? Should I call it an AST? I thought normally ASTs get generated in-memory by parsers.
## [10][RustScan, a modern day port scanner in Rust. Little something I've been working on.](https://www.reddit.com/r/rust/comments/joufu2/rustscan_a_modern_day_port_scanner_in_rust_little/)
- url: http://github.com/rustScan/RustScan
---

## [11][PSA: mod lines' order matter](https://www.reddit.com/r/rust/comments/jp49g7/psa_mod_lines_order_matter/)
- url: https://www.reddit.com/r/rust/comments/jp49g7/psa_mod_lines_order_matter/
---
After creating a new macro in module B, I got a strange error. I could use it in module C,  but not in module A.

After spending way too long, I found out that in main, I first defined module A, then module B and finally module C. Since A was defined before B, I could use all symbols except the macro from B in A... Moving the mod B line to before A fixed it.

The message "help: have you added the \`#\[macro\_use\]\` on the module/import?" didn't help since I had done that.

Is there a reason it works like this?
## [12][I'm new to Rust and... Well this is the first thing that I tried, but I have no idea what to do now...](https://www.reddit.com/r/rust/comments/jopxhi/im_new_to_rust_and_well_this_is_the_first_thing/)
- url: https://www.reddit.com/r/rust/comments/jopxhi/im_new_to_rust_and_well_this_is_the_first_thing/
---
So I am learning Rust and in the meantime having a huge-but-slow side-project and I think this is the first time I actually wrote some code. But I can't get it to compile.

This is my code:

`#![allow(unused)]`

`use git2::Repository;`

`use std::fs;`

`fn main() {`

`makefolder();`

`// The repositories to clone are passed to the function here, I will have some more repo's, and they might not all be on GitHub, in future, I hope to let it read from a list and clone them automatically.`

`clonerepo_github("Marnix0810", "ClaudscapeAI");`

`}`

&amp;#x200B;

`fn makefolder() -&gt; std::io::Result&lt;()&gt; {`

`fs::create_dir("../linked-repositories")?;`

`Ok(())`

`}`

`// This is where I think it goes wrong, but I don't know what to put here instead...`

`fn clonerepo_github(u: str, r: str){`

`let url = ("`[`https://github.com/{}/{}`](https://github.com/{}/{})`", u, r);`

`let repo = match Repository::clone(url, "../linked-repositories/(r)") {`

`Ok(r) =&gt; r,`

`Err(e) =&gt; panic!("failed to clone: {}", e),`

`};`

`}`

Can anyone of you please help me?

&amp;#x200B;

PS: This is also my first time on this subreddit, I hope I'm not out of line here...

&amp;#x200B;

edit: This is the Cargo output I got:

&amp;#x200B;

   `Compiling cio-finit v0.1.0 (C:\Users\marni\Dropbox\Documenten\CIO\cio-finit)`

`error[E0308]: mismatched types`

 `--&gt; src\main.rs:7:22`

  `|`

`7 |     clonerepo_github("Marnix0810", "ClaudscapeAI");`

  `|                      ^^^^^^^^^^^^ expected \`str\`, found \`&amp;str\``

`error[E0308]: mismatched types`

 `--&gt; src\main.rs:7:36`

  `|`

`7 |     clonerepo_github("Marnix0810", "ClaudscapeAI");`

  `|                                    ^^^^^^^^^^^^^^ expected \`str\`, found \`&amp;str\``

`error[E0277]: the size for values of type \`str\` cannot be known at compilation time`

 `--&gt; src\main.rs:7:5`

  `|`

`7 |     clonerepo_github("Marnix0810", "ClaudscapeAI");`

  `|     ^^^^^^^^^^^^^^^^ doesn't have a size known at compile-time`

  `|`

  `= help: the trait \`std::marker::Sized\` is not implemented for \`str\``

  `= note: all function arguments must have a statically known size`

&amp;#x200B;

`error[E0277]: the size for values of type \`str\` cannot be known at compilation time`

  `--&gt; src\main.rs:15:21`

   `|`

`15 | fn clonerepo_github(u: str, r: str){`

   `|                     ^ doesn't have a size known at compile-time`

   `|`

   `= help: the trait \`std::marker::Sized\` is not implemented for \`str\``

   `= help: unsized locals are gated as an unstable feature`

`help: function arguments must have a statically known size, borrowed types always have a known size`

   `|`

`15 | fn clonerepo_github(u: &amp;str, r: str){`

   `|                        ^`

&amp;#x200B;

`error[E0277]: the size for values of type \`str\` cannot be known at compilation time`

  `--&gt; src\main.rs:15:29`

   `|`

`15 | fn clonerepo_github(u: str, r: str){`

   `|                             ^ doesn't have a size known at compile-time`

   `|`

   `= help: the trait \`std::marker::Sized\` is not implemented for \`str\``

   `= help: unsized locals are gated as an unstable feature`

`help: function arguments must have a statically known size, borrowed types always have a known size`

   `|`

`15 | fn clonerepo_github(u: str, r: &amp;str){`

   `|                                ^`

&amp;#x200B;

`error[E0277]: the size for values of type \`str\` cannot be known at compilation time`

  `--&gt; src\main.rs:16:15`

   `|`

`16 |     let url = ("`[`https://github.com/{}/{}`](https://github.com/{}/{})`", u, r);`

   `|               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ doesn't have a size known at compile-time`

   `|`

   `= help: within \`(&amp;'static str, str, str)\`, the trait \`std::marker::Sized\` is not implemented for \`str\``

   `= note: required because it appears within the type \`(&amp;'static str, str, str)\``

   `= note: tuples must have a statically known size to be initialized`

&amp;#x200B;

`error[E0277]: the size for values of type \`str\` cannot be known at compilation time`

  `--&gt; src\main.rs:16:9`

   `|`

`16 |     let url = ("`[`https://github.com/{}/{}`](https://github.com/{}/{})`", u, r);`

   `|         ^^^ doesn't have a size known at compile-time`

   `|`

   `= help: within \`(&amp;str, str, str)\`, the trait \`std::marker::Sized\` is not implemented for \`str\``

   `= note: required because it appears within the type \`(&amp;str, str, str)\``

   `= note: all local variables must have a statically known size`

   `= help: unsized locals are gated as an unstable feature`

&amp;#x200B;

`error[E0308]: mismatched types`

  `--&gt; src\main.rs:17:40`

   `|`

`17 |     let repo = match Repository::clone(url, "../linked-repositories/(r)") {`

   `|                                        ^^^ expected \`&amp;str\`, found tuple`

   `|`

   `= note: expected reference \`&amp;str\``

`found tuple \`(&amp;str, str, str)\``

&amp;#x200B;

`error: aborting due to 8 previous errors`

&amp;#x200B;

`Some errors have detailed explanations: E0277, E0308.`

`For more information about an error, try \`rustc --explain E0277\`.`

`error: could not compile \`cio-finit\`.`

&amp;#x200B;

`To learn more, run the command again with --verbose.`
