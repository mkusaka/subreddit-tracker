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
## [2][What's everyone working on this week (20/2020)?](https://www.reddit.com/r/rust/comments/ghw68g/whats_everyone_working_on_this_week_202020/)
- url: https://www.reddit.com/r/rust/comments/ghw68g/whats_everyone_working_on_this_week_202020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-20-2020/42446?u=llogiq)!
## [3][Is it wrong of me to think that Rust crates have too many dependencies?](https://www.reddit.com/r/rust/comments/gi7v2v/is_it_wrong_of_me_to_think_that_rust_crates_have/)
- url: https://www.reddit.com/r/rust/comments/gi7v2v/is_it_wrong_of_me_to_think_that_rust_crates_have/
---
Hi, I'm sorry for the perhaps provocative title but I couldn't think of any other way I could phrase this.

I've been using Rust for a few months now. Not professionally, I've been mostly learning it and playing with it on the side. I'm in love with the language and its design and how it naturally leads you into writing correct code that does proper error handling. The borrow checker also helped me catch a few hard-to-track-down bugs in more than one occasion. However, there's been one thing I've noticed since Day 1 when I started learning it: importing even a seemingly small and focused crate *you would expect* to be pretty self-contained, more often than not requires you to download another dozen dependent crates. I'm trying my best to express this in a way that doesn't come off as offensive (which is not my intention by any means) but, from the outside, it seems as though people don't want to/don't feel like implementing a feature they need and they resort to a third-party crate that does 100 other things and in turn requires its fair amount of dependencies. I guess you could say cargo makes it very easy to import crates, almost to a fault.

This, paired with the fact that people very often don't enable LTO for release builds and the fact the Rust ABI isn't stable yet so the compiler has to statically link all your dependent crates, results in pretty big binaries being produced, at least in my view as I come from C++. Just to pick out some examples, rust-analyzer is 28M against clangd which is only 291K; ripgrep is 4.7M against GNU grep which is only 211K (binaries have already been `strip`\-ed in both cases). Again, the point of this post is not the size of Rust binaries because I understand the underlying issue and what stabilizing the ABI entails; my argument is that the ease with which cargo (which I love btw) lets you import crates perhaps contributes to this. When I compile a rust program from source it feels almost as though I'm working on a JavaScript code base, with their massive `node_modules` folders.

At the end of the day I just want an opinion from the Rust community about whether it's only me who feels this way (and I'm wrong in doing that) or if it's something other people have thought about at least once in the back of their mind.
## [4][Why I started learning rust and my impressions after the first week of learning](https://www.reddit.com/r/rust/comments/gi4q3d/why_i_started_learning_rust_and_my_impressions/)
- url: https://vignesh.pro/rust-learning-quest-1/
---

## [5][rust-analyzer changelog #24](https://www.reddit.com/r/rust/comments/ghn4iw/rustanalyzer_changelog_24/)
- url: https://rust-analyzer.github.io/thisweek/2020/05/11/changelog-24.html
---

## [6][A Locust-inspired Load Testing Tool In Rust](https://www.reddit.com/r/rust/comments/gi9t6c/a_locustinspired_load_testing_tool_in_rust/)
- url: https://www.tag1consulting.com/blog/goose-locust-inspired-load-testing-tool-rust
---

## [7][ASAP: As Static As Possible memory management](https://www.reddit.com/r/rust/comments/ghy61t/asap_as_static_as_possible_memory_management/)
- url: https://www.cl.cam.ac.uk/techreports/UCAM-CL-TR-908.pdf
---

## [8][Enum value from default match arm](https://www.reddit.com/r/rust/comments/giawbd/enum_value_from_default_match_arm/)
- url: https://www.reddit.com/r/rust/comments/giawbd/enum_value_from_default_match_arm/
---
Hello, I have a recurring enum pattern which tracks operations on values. I'd like to extract the next value with a very simple pattern like this:

    enum Operation {  
     Init(),  
     Sum(Box&lt;Operation&gt;),  
     Sub(Box&lt;Operation&gt;),  
     Mul(Box&lt;Operation&gt;),  
     Div(Box&lt;Operation&gt;) 
    }
    
    impl Operation {  
        pub fn next(&amp;self) -&gt; Option&lt;Operation&gt; {  
            match self {  
                Operation::Init() =&gt; None,  
                Operation::_(nxt) =&gt; Some(nxt)  
            }  
        }  
    }

Is there a way to get `nxt` without matching every single enum value?
## [9][Can I have different dependencies in debug and release builds?](https://www.reddit.com/r/rust/comments/gi909c/can_i_have_different_dependencies_in_debug_and/)
- url: https://www.reddit.com/r/rust/comments/gi909c/can_i_have_different_dependencies_in_debug_and/
---
Is it possible to include for example jemallocator only in a release build and not in debug build to decrease compilation time?

I've tried to make it optional and then in main/lib.rs use 

    #[cfg(debug_assertions)]
    extern crate jemallocator;

But it didn't work. I've also tried this in Cargo.toml:

    [target."cfg(debug_assertions)".dependencies]
    jemallocator = { version = "*"

But then rustc says

&gt; warning: Found `debug_assertions` in `target.'cfg(...)'.dependencies`. This value is not supported for selecting dependencies and will not work as expected. To learn more visit https://doc.rust-lang.org/cargo/reference/specifying-dependencies.html#platform-specific-dependencies

Is there any other way I could do this?
## [10][Wiremock 0.2 is out: HTTP mocking for Rust applications](https://www.reddit.com/r/rust/comments/ghv1wg/wiremock_02_is_out_http_mocking_for_rust/)
- url: https://www.reddit.com/r/rust/comments/ghv1wg/wiremock_02_is_out_http_mocking_for_rust/
---
**What is this?**

[Wiremock](https://crates.io/crates/wiremock) is a crate to help you test applications or routines that perform HTTP calls to third-party systems or APIs that you don't control.

You can spin up as many mock HTTP servers as you need to mock all the 3rd party APIs that your application interacts with, specifying canned responses for your requests.

Each mock HTTP server is fully isolated: tests can be run in parallel, with no interference. Each server is shut down when it goes out of scope  (e.g. end of test execution).

**What is new?**

The main missing feature of v0.1.x was **spying**: being able to assert that a `Mock` had been called (or not called) to verify that one or more side-effects had been performed (e.g. sending an email notification or checking that a cache does indeed prevent an HTTP call from happening).

Spying is now supported via the [expect](https://docs.rs/wiremock/0.2.0/wiremock/struct.Mock.html#method.expect) method: expectations are verified when the `MockServer` is dropped, triggering a panic if one of the expectations has not been satisfied.

Default behavior remains the same: if no expectation is specified, nothing will happen if a Mock is never called or called an arbitrary number of times. This is to encourage a semantic usage of spying in test suites.

There are other minor improvements, mostly around the ergonomics of using `http_types` thanks to their 1.2.0 release that ships more conversion trait implementations for several types used in our matchers.
## [11][Named Field Init in C, C++20, Zig, Rust, &amp; D + Rust's deallocation pattern](https://www.reddit.com/r/rust/comments/ghqqkk/named_field_init_in_c_c20_zig_rust_d_rusts/)
- url: https://www.youtube.com/watch?v=c-NyXKbqmQc
---

## [12][How can I set up Rust to work with Atom? Getting rust-analyzer failures.](https://www.reddit.com/r/rust/comments/gi5jlq/how_can_i_set_up_rust_to_work_with_atom_getting/)
- url: https://www.reddit.com/r/rust/comments/gi5jlq/how_can_i_set_up_rust_to_work_with_atom_getting/
---
I am very new to Rust and I am trying to write some programs using the Atom editor. I have already installed Rust (for Mac) and am currently working on the hello world program.

However, I am unable to get proper language support for Rust on atom. I downloaded the language-rust package so that the language can be supported. When I open a file with a .rs extension, I get the error: Language server command rust-analyzer failed.

How can I properly set up Atom for Rust and avoid this error?
