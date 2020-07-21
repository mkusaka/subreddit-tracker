# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (30/2020)!](https://www.reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (30/2020)?](https://www.reddit.com/r/rust/comments/hurk9f/whats_everyone_working_on_this_week_302020/)
- url: https://www.reddit.com/r/rust/comments/hurk9f/whats_everyone_working_on_this_week_302020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-30-2020/46132?u=llogiq)!
## [3][Patterns to avoid borrowing mutable self more than once? (beyond RcRefcell and inner objects)](https://www.reddit.com/r/rust/comments/hv2zqo/patterns_to_avoid_borrowing_mutable_self_more/)
- url: https://www.reddit.com/r/rust/comments/hv2zqo/patterns_to_avoid_borrowing_mutable_self_more/
---
👋 from Oslo. I've been learning Rust, and making some progress. It strikes me how an important part of the learning curve is [new patterns to solve common problems](https://github.com/rust-unofficial/patterns).

One typical: I make an object with a public API that mutates its own internal state. This outer object often has other objects inside it as well, and references to results of calls methods on these needs to be kept (e.g. a lookup by ID on a sub-object containing a dataset). This pattern often yields **"cannot borrow \*self as mutable more than once at a time"** or similarly for borrowing immutable when self is mutable.

I have succeeded with Rc&lt;RefCell&gt;, and borrowing via inner smaller objects on self. I've worked through solutions to this, and I'm wondering if I'm missing obvious approaches to this API, or perhaps some more idiomatic varieties? The playground below demonstrates the problem with struct B.

E.g.: some might say "write it in a different way", but I am not sure how to create an easy API that does the same in other way? Even if you use functions and pass them an object instead of self, you have the same situation - you cannot borrow this object mutably more than once, and it gets locked behind results of method calls that use it?

Playground:  [https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=9b2018ef9f4b09770f758ee82ed511af](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=9b2018ef9f4b09770f758ee82ed511af)

Hoping someone more rusty than me have some robust answers, and that I am explaining the challenge clearly (:
## [4][A CLI Library for making CLI apps.](https://www.reddit.com/r/rust/comments/hv39p5/a_cli_library_for_making_cli_apps/)
- url: https://www.reddit.com/r/rust/comments/hv39p5/a_cli_library_for_making_cli_apps/
---
Hey everyone.

I'm planning to write a CLI app that takes the entirety of a terminal window. Example of this is "htop", the program fills the entire terminal window and can even has clickable options. So please does anyone know of a library that can achieve this or if not then maybe a way to go about this.

Thanks.
## [5][WaffleLink - A W.I.P runtime for dynamically typed languages with JIT compilation](https://www.reddit.com/r/rust/comments/hv5i8x/wafflelink_a_wip_runtime_for_dynamically_typed/)
- url: https://www.reddit.com/r/rust/comments/hv5i8x/wafflelink_a_wip_runtime_for_dynamically_typed/
---
[WaffleLink](https://github.com/jazz-lang/wafflelink) is a runtime for dynamically typed languages that supports tiering execution pipeline from interpreter to JIT compilation. Currently only Interpreter-&gt;TemplateJIT is implemented, but optimizing JIT with it's own IR will be implemented soon. Right now WaffleLink is *the first* language runtime in Rust that does On-Stack-Replacement and has tiering execution with JIT.

\- Interpreter(src/interpreter.rs): just executes code in big \`match\` statement and loop and does small profiling so JIT can generate better code

\- TemplateJIT (src/jit/\*): compiles bytecode to machine code using [https://github.com/playXE/masm-rs](https://github.com/playXE/masm-rs) library, code generated by JIT can be modified on the fly e.g if JIT is generated code for int32 binary operation based on profiling but operation fails after some executions then it will generate new code and replace jump from slow path to this new code (src/jit/mathic.rs)

**TODO List**

\- Finish implementation of JIT and interpreter, they have impl for most of opcodes right now

\- Create compiler from simple JS-like language to WaffleLink bytecode

\- Optimizing JIT

\- Inline caching for objects and arrays
## [6][How to fix `cross` builds in Rust 1.45](https://www.reddit.com/r/rust/comments/hv67j7/how_to_fix_cross_builds_in_rust_145/)
- url: https://www.reddit.com/r/rust/comments/hv67j7/how_to_fix_cross_builds_in_rust_145/
---
Rust 1.45 made some changes to LTO that have made CI using [cross](https://github.com/rust-embedded/cross) no longer work if it was set up prior to [25 Sep 2019](https://github.com/rust-embedded/cross/commit/95a7e1bb02e7f50739c29cbfcce95ef678252a23#diff-35ea51baf1fe7f0142ad5f950855dde0). This is because the scripts that `cross` use used to pass `-C lto` to rustc, which as of rust 1.45 no longer works.

The symptom to look out for is during the final linking step, when it fails:

       ....
       Compiling $pkg_name v0.7.6 (/project)
    
    error: options `-C embed-bitcode=no` and `-C lto` are incompatible
    
    error: could not compile `$pkg_name`.
    
    To learn more, run the command again with --verbose.

The quick fix is to modify `ci/before_deploy.sh` and remove `-- -C lto`:

    diff --git a/ci/before_deploy.sh b/ci/before_deploy.sh
    index 6dabfde..5618474 100644
    --- a/ci/before_deploy.sh
    +++ b/ci/before_deploy.sh
    @@ -19,7 +19,7 @@ main() {
         test -f Cargo.lock || cargo generate-lockfile
    
         # TODO Update this to build the artifacts that matter to you
    -    cross rustc --bin $pkg_name --target $TARGET --release -- -C lto
    +    cross rustc --bin $pkg_name --target $TARGET --release
    
         # TODO Update this to package the right artifacts
         cp target/$TARGET/release/$pkg_name$EXTENSION $stage/

It probably would be better to update to a newer version of `cross` that does lto the proper way via `Cargo.toml`, but this simple modification will get it building again. It also would be a good idea to add `beta` builds to your projects to catch these sorts of issues before the new compiler is released -- I'm going to go do that now.
## [7][Blog post: Three Architectures for a Responsive IDE](https://www.reddit.com/r/rust/comments/huja76/blog_post_three_architectures_for_a_responsive_ide/)
- url: https://rust-analyzer.github.io/blog/2020/07/20/three-architectures-for-responsive-ide.html
---

## [8][How to match the root path with a Warp filter?](https://www.reddit.com/r/rust/comments/hv4t0i/how_to_match_the_root_path_with_a_warp_filter/)
- url: https://www.reddit.com/r/rust/comments/hv4t0i/how_to_match_the_root_path_with_a_warp_filter/
---
By following the Warp docs I've created a simple web server that serves static files. For example to serve an index.html and script.js file at the locations `localhost:8000/index.html` and `localhost:8000/script.js` I use the following:

    let index_html = warp::path("index.html").map(...); // bunch of stuff in the map
    let script_js = warp::path("script.js").map(...);
    let routes = warp::get().and(index_html.or(script_js));
    warp::serve(routes).run(([127, 0, 0, 1], 8000)).await;

I want to serve the index file from `/` rather than `/index.html`, but I'm not sure how to specify the domain root with `warp::path`. Using `warp::path("")` or `warp::path("/")` doesn't work.

Is it possible to use a Warp filter to match the root path?
## [9][Are the Result/Option wrapper, monads?](https://www.reddit.com/r/rust/comments/hv1675/are_the_resultoption_wrapper_monads/)
- url: https://www.reddit.com/r/rust/comments/hv1675/are_the_resultoption_wrapper_monads/
---
Is just that I'm wonder if either `Result` or `Option` wrapper are monads?

As I understand the concept (naive concept) of a monad, is basically a wrapper for a value. Without going any deeper, another example of monad could be the IO monad for Haskell which lets the mutation of data, the Promise monad jn Js/Ts which wraps a value until is available or fails (similar to the Result monad) and finally the Task monad in C#, which does  similar job as the Promise in Js/Ts.
## [10][(A Few) Advanced Variable Types in Rust — Learning Rust](https://www.reddit.com/r/rust/comments/huvfjf/a_few_advanced_variable_types_in_rust_learning/)
- url: https://rust.graystorm.com/2020/07/20/a-few-advanced-variable-types-in-rust/
---

## [11][How does redis-rs scale?](https://www.reddit.com/r/rust/comments/hv1m04/how_does_redisrs_scale/)
- url: https://www.reddit.com/r/rust/comments/hv1m04/how_does_redisrs_scale/
---
I'm new to rust, and I'm investigating some pieces for a future project. I'm curious about redis-rs, I see it has async support, but doesn't currently have a connection pool.

For connections within the same datacenter (Azure or AWS), should I bother using redis-rs with r2d2, or should I just use redis-rs async directly?

How do you guys who use redis-rs in production do it, and what performance numbers are you seeing (before redis itself becomes a bottleneck)?

Sorry if I'm asking a dumb question, i've relied on connection pools too much previously.
## [12][Show r/rust: A tiny UNIX shell, from C to Rust, from 1990 to 2020](https://www.reddit.com/r/rust/comments/hv78hc/show_rrust_a_tiny_unix_shell_from_c_to_rust_from/)
- url: https://github.com/zserge/tinysh
---

