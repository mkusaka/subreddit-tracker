# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (46/2020)!](https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 364](https://www.reddit.com/r/rust/comments/jslo80/this_week_in_rust_364/)
- url: https://this-week-in-rust.org/blog/2020/11/11/this-week-in-rust-364/
---

## [3][[Inside Rust] Source-based code coverage in nightly](https://www.reddit.com/r/rust/comments/jt45jt/inside_rust_sourcebased_code_coverage_in_nightly/)
- url: https://blog.rust-lang.org/inside-rust/2020/11/12/source-based-code-coverage.html
---

## [4][SQLite Parser Pt. 2 The Header... continues](https://www.reddit.com/r/rust/comments/jt9s2v/sqlite_parser_pt_2_the_header_continues/)
- url: https://freemasen.com/blog/sqlite-parser-pt-2/index.html
---

## [5][Learn Assembly by Writing Entirely Too Many Brainfuck Compilers in Rust](https://www.reddit.com/r/rust/comments/jsvdsy/learn_assembly_by_writing_entirely_too_many/)
- url: https://github.com/pretzelhammer/rust-blog/blob/master/posts/too-many-brainfuck-compilers.md
---

## [6][rkyv: a zero-copy deserialization framework for Rust](https://www.reddit.com/r/rust/comments/jss6h4/rkyv_a_zerocopy_deserialization_framework_for_rust/)
- url: https://www.reddit.com/r/rust/comments/jss6h4/rkyv_a_zerocopy_deserialization_framework_for_rust/
---
Hi everyone! I'm a long-time lurker and first poster so please be gentle. :)

I just released the first version of rkyv (*archive*), a zero-copy deserialization framework for Rust:

Source code: [https://github.com/djkoloski/rkyv](https://github.com/djkoloski/rkyv)  
Docs: [https://docs.rs/rkyv](https://docs.rs/rkyv), [https://docs.rs/rkyv\_dyn](https://docs.rs/rkyv_dyn), [https://docs.rs/rkyv\_typename](https://docs.rs/rkyv_typename)

rkyv is similar to other zero-copy deserialization frameworks like Cap'n Proto and FlatBuffers, but it's 100% pure rust and uses macro magic to build its serialization functions like serde does. The main feature is that it's zero-copy, meaning that all you have to do to "deserialize" your data is just cast a pointer. All of the data is serialized in a way that makes its in-memory representation the same as its archived representation.

rkyv sports a couple of neat features:

* Derive macros, even for complex and generic types
* \#\[no\_std\] support
* Hashmap support through a custom implementation based off of hashbrown
* Trait object serialization through the accompanying rkyv\_dyn crate. You can serialize out a trait object then use it with just a pointer cast!
* Plenty of examples and tests to make sure everything's working right.

rkyv was primarily made with an eye toward game development, where lots of static data needs to be read in and load times negatively impact player experience. Speaking from experience, deserialization takes up a big chunk of load times so a world without deserialization is a faster one!

I've been writing rust for a while but this is my first contribution to the community. If you're interested, take a look and leave me some feedback if you're interested. For example, I've only tested on Windows due to hardware constraints but if some tests are failing on other toolchains I'll find a way to get them fixed.

Thanks for taking a look!
## [7][Poi - a pragmatic point-free theorem prover assistant written in Rust - just multiplied matrices for the first time!](https://www.reddit.com/r/rust/comments/jt7xte/poi_a_pragmatic_pointfree_theorem_prover/)
- url: https://twitter.com/PistonDeveloper/status/1327061178951077890
---

## [8][Current state of Rust support for writing Linux kernel space code (drivers)?](https://www.reddit.com/r/rust/comments/jsxgp6/current_state_of_rust_support_for_writing_linux/)
- url: https://www.reddit.com/r/rust/comments/jsxgp6/current_state_of_rust_support_for_writing_linux/
---
I do a ton of work with the linux kernel for my dayjob, which is currently only C due to the Linux tooling supporting only C. I also do a decent bit of C++ for userspace tools, and while I would love support for a higher level language for kernel space code, I can see why them written in C++ would be an utter disaster.

Seeing Linus and the kernel mailing list giving support to rust (example [one](https://www.reddit.com/r/rust/comments/hplaz4/lkml_linus_torvalds_re_linux_kernel_intree_rust/) and [two](https://www.reddit.com/r/rust/comments/ik42b4/supporting_linux_kernel_development_in_rust/)) makes me extremely happy. But I do not see much new activity in the past two months since these announcements, so I figured I would ask, how are these efforts going? Have they stalled? Is there a place I could look to follow this closer (somewhere in the LKML?)? Closest I found is the example rust Linux kernel driver [link](https://github.com/fishinabarrel/linux-kernel-module-rust/blob/master/hello-world/src/lib.rs), but honestly, it does not look as ergonomic as a functionally equivalent C module.

Or maybe, am I just over-hyping myself?
## [9][Unexpected borrow checker behavior](https://www.reddit.com/r/rust/comments/jt1s5q/unexpected_borrow_checker_behavior/)
- url: https://www.reddit.com/r/rust/comments/jt1s5q/unexpected_borrow_checker_behavior/
---
I was hacking around on some code and came across a weird thing that rustc complains about. If you have an async function that takes a static reference and at least 2 other references, those other references must have the same lifetime.  As an example this code fails to compile with the following error:

```rust
async fn test(test: &amp;String, test2: &amp;String, test3: &amp;'static String) {}
```

```text
error[E0700]: hidden type for `impl Trait` captures lifetime that does not appear in bounds
  --&gt; src/cgi.rs:31:71
   |
31 | async fn test(test: &amp; String, test2: &amp;String, test3: &amp;'static String) {}
   |                                                                       ^
   |
note: hidden type `impl futures_lite::Future` captures lifetime smaller than the function body
  --&gt; src/cgi.rs:31:71
   |
31 | async fn test(test: &amp; String, test2: &amp;String, test3: &amp;'static String) {}
   |   
```

However adding an explicit lifetime fixes the problem:

```rust
async fn test&lt;'a&gt;(test: &amp; String, test2: &amp;'a String, test3: &amp;'static String) {}
```

This behavior is a bit unexpected and the error message it returned was not helpful.  Any ides if this a bug or just a non-helpful error message?
## [10][terender: command-line tool for rendering tera templates](https://www.reddit.com/r/rust/comments/jsxb0w/terender_commandline_tool_for_rendering_tera/)
- url: https://www.reddit.com/r/rust/comments/jsxb0w/terender_commandline_tool_for_rendering_tera/
---
I recently had several occasions where I wanted to generate source code or documentation from structured data that is available to me in JSON format. In order to facilitate that, it took me just a bit of plumbing to build a command-line tool that allows me to do that based on tera templates. I thought it might be useful to others as well, so I published it to [https://crates.io/crates/terender](https://crates.io/crates/terender). As said, I just did some plumbing, the main credits for the work go to the authors of tera, serde and serde\_json, structopt and anyhow.

I intend to add support for other file formats as well in the future, please let me know if you have a need for some specific format so I can prioritize it.

If you like the tool or if you would like me to add functionality, please let me know (either directly, or submit an issue).
## [11][How iterate a sub value in a json hashmap?](https://www.reddit.com/r/rust/comments/jsz0c5/how_iterate_a_sub_value_in_a_json_hashmap/)
- url: https://www.reddit.com/r/rust/comments/jsz0c5/how_iterate_a_sub_value_in_a_json_hashmap/
---
This is what I am using to iterate through the root level of a json:



https://pastebin.com/DwHyUDmG


it returns me with:


    keys "phones"
    keys "name"
    keys "contact"
    keys "age"


How can I iterate the content of the "contact" key
(as per commented out code)
## [12][In actix-web, how to maintain a shopping cart counter in a middleware(?) that is accessible in (near) all responses?](https://www.reddit.com/r/rust/comments/jsxlct/in_actixweb_how_to_maintain_a_shopping_cart/)
- url: https://www.reddit.com/r/rust/comments/jsxlct/in_actixweb_how_to_maintain_a_shopping_cart/
---
Building an eCommerce site, I need to display the shopping cart counter icon nearly everywhere. Now I repeat the code to get this in each response, but wonder how abstract this away, probably in a middleware?

But in a few places need to know in the response this count, so how ALSO get it from inside them?
