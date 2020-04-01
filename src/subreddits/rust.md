# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (14/2020)!](https://www.reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (14/2020)?](https://www.reddit.com/r/rust/comments/frff7k/whats_everyone_working_on_this_week_142020/)
- url: https://www.reddit.com/r/rust/comments/frff7k/whats_everyone_working_on_this_week_142020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-14-2020/40159?u=llogiq)!
## [3][Writing a reverse shell in Rust?](https://www.reddit.com/r/rust/comments/fsxaaa/writing_a_reverse_shell_in_rust/)
- url: https://www.reddit.com/r/rust/comments/fsxaaa/writing_a_reverse_shell_in_rust/
---
So I'm a Rust beginner, and I thought it would be cool to write my own reverse shell in Rust. I like messing with security and occasionally do boot2root challenges like [hackthebox](https://hackthebox.eu/). Usually I just use a simple netcat shell, since meterpreter is a little bit overkill and is detected by virtually every antivirus in existance. But I thought creating my own thing would be a fun challenge. I already wrote a simple program that creates a TcpStream it can send/receive stuff from.

Now my question is, how could I go about "piping" a program (cmd or bash) through a TcpStream (like netcat's `-e` flag does)? I found a couple of existing reverse shells written in Rust, but most of them are pretty outdated and don't compile on the latest version anymore (and generally not cross-platform, which would also be nice).
## [4][Bryan Cantrill on Rust making systems programming exciting](https://www.reddit.com/r/rust/comments/fsih8g/bryan_cantrill_on_rust_making_systems_programming/)
- url: https://m.youtube.com/watch?v=vvZA9n3e5pc&amp;t=49m20s
---

## [5][Rust debugging is still really painful, VSCode particularly included](https://www.reddit.com/r/rust/comments/fspk43/rust_debugging_is_still_really_painful_vscode/)
- url: https://www.reddit.com/r/rust/comments/fspk43/rust_debugging_is_still_really_painful_vscode/
---
I'm surprised I'm writing this, but I cannot find a good debugger for Rust. I've been using VSCode as it has been the best---everything else has horribly mangled function names and whatnot. But VSCode debugging is a very broken experience. I've had issues with how it handles the embedded terminal. I'm currently having issues where code gets a segfault within the debugger which runs totally fine without it. And with *the exact same code* I've gotten "Inconsistency detected by ld.so: rtld.c: 1180: dl\_main: Assertion" which makes the launch fail. It's like once every two weeks there's some new problem.

I'm tired of realizing I have to file a VSCode bug report for Rust debugging. The lack of a decent debugger is really hampering me right now. I want to give up.
## [6][Rust in a Commercial](https://www.reddit.com/r/rust/comments/fsd5s8/rust_in_a_commercial/)
- url: https://youtu.be/h3yFOf6hIjQ
---

## [7][Oxidize 1K: March 20th, 2020 - Embedded Rust Lightning Conference - Full Video Stream](https://www.reddit.com/r/rust/comments/fsn4ux/oxidize_1k_march_20th_2020_embedded_rust/)
- url: https://www.youtube.com/watch?v=zPuELAzJyno
---

## [8][macro_rules! ($($a:expr),+) flattening](https://www.reddit.com/r/rust/comments/fsw9b1/macro_rules_aexpr_flattening/)
- url: https://www.reddit.com/r/rust/comments/fsw9b1/macro_rules_aexpr_flattening/
---
Hey guys,
I am trying to write proper DRY code by using a declarative macro. For this purpose, I want to input a string with n+2 format fillers, replace the first and the last filler with static data and the intermediate ones with my variables, and wrap the API call around that. For this, I wrote a macro_rules! macro:

    macro_rules! apicall {
        ($s:expr,$($args:expr),*) =&gt; {
            Ok(self
                .client
                .get(
                    format!(
                        $s,
                        &amp;self.apiurl,
                        $args,
                        &amp;self.apikey,
                    )
                    .as_str(),
                )
                .send()
                .await?
                .json()
                .await?)
        };
    }


The problem is that at the time I use the $args macro variable, I get a compiler error because the argument still repeats and the preprocessor doesn't know how to deal with that. So I tried to write a fix:

    macro_rules! flatten {
        ($arg:expr) =&gt; ($arg);
        ($arg1:expr, $($args:expr),+) =&gt; {
            $arg1, dewrap!($args)
        }
    }

and replaced $args with flatten!($args). This however still produces the same compiler error even though $args should match the second pattern of flatten!. How do I work around this?
## [9][Introducing TinyVec: 100% safe alternative to SmallVec and ArrayVec](https://www.reddit.com/r/rust/comments/fshuhk/introducing_tinyvec_100_safe_alternative_to/)
- url: https://www.reddit.com/r/rust/comments/fshuhk/introducing_tinyvec_100_safe_alternative_to/
---
[TinyVec](https://github.com/Lokathor/tinyvec) is a 100% safe code alternative to [SmallVec](https://github.com/servo/rust-smallvec) and [ArrayVec](https://github.com/bluss/arrayvec) crates. While SmallVec and ArrayVec create an array of unintialized memory and try to hide it from the user, TinyVec simply initializes the entire array up front. Real-world performance of this approach is surprisingly good: I have replaced SmallVec with TinyVec in `unicode-normalization` and `lewton` crates with no measurable impact on benchmarks.

The main drawback is that the type stored in TinyVec must implement `Default`, so it cannot replace SmallVec or ArrayVec in _all_ scenarios.

TinyVec is implemented as an enum of `std::Vec` and `tinyvec::ArrayVec`, which allows some optimizations that are not possible with SmallVec - for example, you can explicitly match on this enum and call `drain()` on the underlying type to avoid branching on every access.

TinyVec is designed to be a drop-in replacement for `std::Vec`, more so than SmallVec or ArrayVec that diverge from Vec behavior in some of their methods. We got a fuzzer to verify that TinyVec's behavior is identical to `std::Vec` via [arbitrary-model-tests](https://github.com/jakubadamw/arbitrary-model-tests) (which has found a few bugs!). Newly introduced methods are given deliberately long names that are unlikely to clash with future additions on Vec.

For a more detailed overview of the crate see the [docs.rs page](https://docs.rs/tinyvec/0.3.3/tinyvec/).

P.S. I'm not the author of the crate, I'm just a happy user of it.
## [10][kmon: Linux Kernel Manager and Activity Monitor written in Rust](https://www.reddit.com/r/rust/comments/fsz7ef/kmon_linux_kernel_manager_and_activity_monitor/)
- url: https://www.reddit.com/r/rust/comments/fsz7ef/kmon_linux_kernel_manager_and_activity_monitor/
---
[kmon](https://github.com/orhun/kmon) provides a text-based user interface for managing the Linux kernel modules and monitoring the kernel activities. By managing, it means loading, unloading, blacklisting and showing the information of a module. These updates in the kernel modules, logs about the hardware and other kernel messages can be tracked with the real-time activity monitor in kmon. Since the usage of different tools like dmesg and kmod are required for these tasks in Linux, kmon aims to gather them in a single terminal window and facilitate the usage as much as possible while keeping the functionality.

kmon is written in Rust and uses [tui-rs](https://github.com/fdehau/tui-rs) &amp; [termion](https://github.com/redox-os/termion) libraries for its text-based user interface.

[kmon on action](https://i.redd.it/frrmmacme7q41.gif)

**Project Homepage:** [https://github.com/orhun/kmon](https://github.com/orhun/kmon)  
**Rust Package:** [https://crates.io/crates/kmon](https://crates.io/crates/kmon)
## [11][Rust in Action update - upcoming live stream; Japanese and Korean versions will be available, probably Russian also](https://www.reddit.com/r/rust/comments/fsm1i1/rust_in_action_update_upcoming_live_stream/)
- url: https://www.reddit.com/r/rust/comments/fsm1i1/rust_in_action_update_upcoming_live_stream/
---
Hi all,

[Rust in Action](https://www.manning.com/books/rust-in-action?a_aid=rust&amp;a_bid=0367c58f&amp;chan=reddit) is an upcoming book from Manning Publications that teaches Rust by walking through systems programming examples.

Some updates

* I will be [live streaming](https://www.twitch.tv/manningpublications) on 16 April (7pm EDT)
* there is still time to provide feedback! I'm committed to addressing every comment submitted to [the liveBook](https://livebook.manning.com/book/rust-in-action/)
* the publisher informs me that Japanese and Korean translation rights have been sold and Russian is negotiation
* the book should be at the printers in May!
## [12][Specs and Legion, two very different approaches to ECS](https://www.reddit.com/r/rust/comments/fsczky/specs_and_legion_two_very_different_approaches_to/)
- url: https://csherratt.github.io/blog/posts/specs-and-legion/
---

