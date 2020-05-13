# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (20/2020)!](https://www.reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ghw4v6/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This week in Rust 338](https://www.reddit.com/r/rust/comments/gikfuy/this_week_in_rust_338/)
- url: https://this-week-in-rust.org/blog/2020/05/12/this-week-in-rust-338/
---

## [3][ringbahn - a safe futures API for io-uring](https://www.reddit.com/r/rust/comments/giiy3z/ringbahn_a_safe_futures_api_for_iouring/)
- url: https://github.com/withoutboats/ringbahn
---

## [4][Why not more usage of debug_assert!?](https://www.reddit.com/r/rust/comments/girtb0/why_not_more_usage_of_debug_assert/)
- url: https://www.reddit.com/r/rust/comments/girtb0/why_not_more_usage_of_debug_assert/
---
Hi! I notice that Rust sources (e.g. libcore) do not use `debug_assert!` often. There's `// SAFETY` comments and [proofs-in-comments](https://doc.rust-lang.org/src/core/slice/mod.rs.html#1552) but these are not machine checked.

My question is why not more usage of `debug_assert!` especially in unsafe code? For example in `get_unchecked`?
## [5][Security advisories for April 2020: rustqlite, os_str_bytes, flatbuffers](https://www.reddit.com/r/rust/comments/gidtpe/security_advisories_for_april_2020_rustqlite_os/)
- url: https://www.reddit.com/r/rust/comments/gidtpe/security_advisories_for_april_2020_rustqlite_os/
---
[RustSec](https://rustsec.org) is a community database of security advisories filed against crates published to crates.io. It is maintained by the [Rust Secure Code Working Group](https://github.com/rust-secure-code/wg).

The following security issues have been identified in Rust crates in April 2020:

 * [rusqlite: Various memory safety issues](https://rustsec.org/advisories/RUSTSEC-2020-0014.html)
 * [os_str_bytes: Relies on undefined behavior of `char::from_u32_unchecked`](https://rustsec.org/advisories/RUSTSEC-2020-0012.html)
 * [flatbuffers: `read_scalar` and `read_scalar_at` allow transmuting values without `unsafe` blocks](https://rustsec.org/advisories/RUSTSEC-2020-0009.html)

You can use [cargo-audit](https://github.com/RustSec/cargo-audit) to check whether your code depends on vulnerable versions of these crates and upgrade. [A GitHub action](https://github.com/actions-rs/audit-check) that files bugs if your code depends on vulnerable crates is also available.

Additionally, we have published security advisories for two crates that intentionally violate Rust's memory safety guarantees: [fake-static](https://crates.io/crates/fake-static) and [plutonium](https://docs.rs/plutonium/0.2.2/plutonium/). This has proven to be controversial, so we have retracted the latter advisory for the time being.

So far we have abided by [the contract between safe and unsafe code laid out in the Nomicon](https://doc.rust-lang.org/nomicon/safe-unsafe-meaning.html):

&gt; **No matter what, Safe Rust can't cause Undefined Behavior.**

Thus we consider violations of that contract to be potential security issues.

Examples of code intentionally violating this contract include the [`plutonium`](https://docs.rs/plutonium/0.2.2/plutonium/) crate, or an unsound `io_uring` wrapper design descibed in [this blog post](https://boats.gitlab.io/blog/post/io-uring/).

Since the RustSec database exists to serve the Rust community and not the database maintainers, we would like to hear from *you* on how would you like intentional violations of this contract to be handled. The options are:

1. Treat them as security issues like any other and notify the public about them. If your CI/CD pipeline runs `cargo-audit`, they will be surfaced as hard failures.
2. Create a notice but surface it as a warning only, similar to how unmaintained crates are currently handled. Intentional memory safety violations would get their own distinct category.
3. Do not surface such issues in `cargo-audit` in any way, but track them in order to allow third-party tooling such as [`cargo-deny`](https://github.com/EmbarkStudios/cargo-deny) to consume this data.
4. Do nothing to inform the public about such issues.

Please let us know which option would be preferable for you and why in the comments - Reddit's comment system enables much more structured conversations than Github issues. We're also open to other suggestions on how to handle such cases.
## [6][Is there a Typescript parser for Rust?](https://www.reddit.com/r/rust/comments/giwbao/is_there_a_typescript_parser_for_rust/)
- url: https://www.reddit.com/r/rust/comments/giwbao/is_there_a_typescript_parser_for_rust/
---
I would like to parse Typescript files from an external library (e.g., [Vector3.d.ts from three.js](https://github.com/mrdoob/three.js/blob/master/src/math/Vector3.d.ts)) to extract function definitions and type information. Does anyone know of a project or crate that has this functionality? So far I have only found RESS/RESSA which only handle ECMAScript as far as I can tell.
## [7][Rust Traits: Iterator](https://www.reddit.com/r/rust/comments/gip4as/rust_traits_iterator/)
- url: https://dpbriggs.ca/blog/Rust-Traits-Iterator
---

## [8][Is it wrong of me to think that Rust crates have too many dependencies?](https://www.reddit.com/r/rust/comments/gi7v2v/is_it_wrong_of_me_to_think_that_rust_crates_have/)
- url: https://www.reddit.com/r/rust/comments/gi7v2v/is_it_wrong_of_me_to_think_that_rust_crates_have/
---
Hi, I'm sorry for the perhaps provocative title but I couldn't think of any other way I could phrase this.

I've been using Rust for a few months now. Not professionally, I've been mostly learning it and playing with it on the side. I'm in love with the language and its design and how it naturally leads you into writing correct code that does proper error handling. The borrow checker also helped me catch a few hard-to-track-down bugs in more than one occasion. However, there's been one thing I've noticed since Day 1 when I started learning it: importing even a seemingly small and focused crate *you would expect* to be pretty self-contained, more often than not requires you to download another dozen dependent crates. I'm trying my best to express this in a way that doesn't come off as offensive (which is not my intention by any means) but, from the outside, it seems as though people don't want to/don't feel like implementing a feature they need and they resort to a third-party crate that does 100 other things and in turn requires its fair amount of dependencies. I guess you could say cargo makes it very easy to import crates, almost to a fault.

This, paired with the fact that people very often don't enable LTO for release builds and the fact the Rust ABI isn't stable yet so the compiler has to statically link all your dependent crates, results in pretty big binaries being produced, at least in my view as I come from C++. Just to pick out some examples, rust-analyzer is 28M against clangd which is only 291K; ripgrep is 4.7M against GNU grep which is only 211K (binaries have already been `strip`\-ed in both cases). Again, the point of this post is not the size of Rust binaries because I understand the underlying issue and what stabilizing the ABI entails; my argument is that the ease with which cargo (which I love btw) lets you import crates perhaps contributes to this. When I compile a rust program from source it feels almost as though I'm working on a JavaScript code base, with their massive `node_modules` folders.

At the end of the day I just want an opinion from the Rust community about whether it's only me who feels this way (and I'm wrong in doing that) or if it's something other people have thought about at least once in the back of their mind.
## [9][Dynamic vs Static Dispatch](https://www.reddit.com/r/rust/comments/giqm5j/dynamic_vs_static_dispatch/)
- url: https://www.reddit.com/r/rust/comments/giqm5j/dynamic_vs_static_dispatch/
---
Hi all,

I am writing an app, and there was a place where i ended up using dynamic dispatch. I would like to see if this is avoidable, and/or if there are better ways to do it.

Basically my app can be at one of different modes. A mode is a struct with some mode-specific info, plus some global state struct. There is an event loop and at each iteration I do `mode = mode.handle_event(event)`.  The mode is a `Box&lt;dyn Mode&gt;` object, where `Mode` is a trait that abstracts what a mode of my application can do.

The `handle_event` function then handles the event, and returns the new mode that the application will be in.

Originally I had an enum for all the possible modes and matched over that enum. But this created many problems, especially when I had to change the mode that the app is in.

Are there any patterns that are commonly used in rust for analogous situtations?

thanks!
## [10][Learning Rust in 2020](https://www.reddit.com/r/rust/comments/gie64f/learning_rust_in_2020/)
- url: https://github.com/pretzelhammer/rust-blog/blob/master/posts/learning-rust-in-2020.md
---

## [11][What I learned contributing to Rust-Analyzer](https://www.reddit.com/r/rust/comments/gicwm7/what_i_learned_contributing_to_rustanalyzer/)
- url: https://dev.to/bnjjj/what-i-learned-contributing-to-rust-analyzer-4c7e
---

## [12][Filesystems: RISCV OS in Rust](https://www.reddit.com/r/rust/comments/gieoc1/filesystems_riscv_os_in_rust/)
- url: http://osblog.stephenmarz.com/ch10.html
---

