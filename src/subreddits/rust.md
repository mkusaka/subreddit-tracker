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
## [3][Weird memory unsafety problem in safe Rust](https://www.reddit.com/r/rust/comments/f51fem/weird_memory_unsafety_problem_in_safe_rust/)
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
## [4][I audited 3 different implementation of async RwLock.](https://www.reddit.com/r/rust/comments/f4zldz/i_audited_3_different_implementation_of_async/)
- url: https://www.reddit.com/r/rust/comments/f4zldz/i_audited_3_different_implementation_of_async/
---
"Audit" is probably a strong word. Also, take this with a grain of salt. I am by no means an expert with task scheduling. I am, however, interested in using an async `RwLock` in a production environment.

What I was really interested in is answering the question: If I have a ton of readers acquiring and releasing the lock at all times, do the writers get a chance to acquire the lock, too?

To start with, I first looked at blocking implementations. If you check the docs for [`std::sync::RwLock`](https://doc.rust-lang.org/stable/std/sync/struct.RwLock.html), you'll clearly find that answering this question is entirely left up to the kernel implementation:

&gt;The priority policy of the lock is dependent on the underlying operating system's implementation, and this type does not guarantee that any particular policy will be used.

So what *actually* happens on various platforms?

* Windows: Readers and writers are fairly queued
* Linux: Readers will starve the writers
* macOS: Readers and writers are fairly queued

By "fairly queued" I mean that interleaved read/write locks will be queued in the order they attempt to acquire. With the bonus that all readers will acquire the lock in groups. This means that readers will have to wait behind any writers, and writers will have to wait behind any readers. But importantly, neither will be the dominant workload.

I first wrote a small test (on macOS), and was surprised to see the behavior that *I wanted* out of an `RwLock`. But then I ran the same code on Linux and was dismayed that it prioritizes readers. The lock on Windows behaves roughly the same as on macOS.

    #[cfg(test)]
    mod test {
        use std::sync::{Arc, RwLock};
        use std::thread::{self, JoinHandle};
        use std::time::Duration;
    
        fn create_reader(name: i32, lock: Arc&lt;RwLock&lt;u32&gt;&gt;) -&gt; JoinHandle&lt;()&gt; {
            thread::spawn(move || {
                println!("reader {}: acquiring read lock...", name);
                let guard = lock.read().unwrap();
    
                println!("reader {}: sleeping...", name);
                thread::sleep(Duration::from_secs(3));
    
                println!("reader {}: end: {}", name, *guard);
            })
        }
    
        fn create_writer(name: i32, lock: Arc&lt;RwLock&lt;u32&gt;&gt;) -&gt; JoinHandle&lt;()&gt; {
            thread::spawn(move || {
                println!("writer {}: acquiring write lock...", name);
                let guard = lock.write().unwrap();
    
                println!("writer {}: sleeping...", name);
                thread::sleep(Duration::from_secs(3));
    
                println!("writer {}: end: {}", name, *guard);
            })
        }
    
        #[test]
        fn test_rwlock() {
            let mut threads = Vec::new();
            let shared_lock = Arc::new(RwLock::new(0));
    
            // Start two reader threads
            let lock = Arc::clone(&amp;shared_lock);
            threads.push(create_reader(1, lock));
    
            let lock = Arc::clone(&amp;shared_lock);
            threads.push(create_reader(2, lock));
    
            // Wait for threads to sleep
            thread::sleep(Duration::from_millis(100));
    
            // Start a writer thread
            let lock = Arc::clone(&amp;shared_lock);
            threads.push(create_writer(1, lock));
    
            // Wait for threads to sleep
            thread::sleep(Duration::from_millis(100));
    
            // Start another two reader threads
            let lock = Arc::clone(&amp;shared_lock);
            threads.push(create_reader(3, lock));
    
            let lock = Arc::clone(&amp;shared_lock);
            threads.push(create_reader(4, lock));
    
            // Wait for threads to sleep
            thread::sleep(Duration::from_millis(100));
    
            // Start another writer thread
            let lock = Arc::clone(&amp;shared_lock);
            threads.push(create_writer(2, lock));
    
            // Wait for all threads to exit
            for t in threads {
                t.join().unwrap();
            }
        }
    }

Expected results from this test: the "sleeping..." message for readers and writers should be interleaved, with reader 1 &amp; 2 sleeping first, followed by writer 1. After this point is kind of a mixed bag, depending on the OS scheduler. Windows behavior is pretty wild! macOS seems to be very consistent. And Linux is the worst; it allows *all readers* to sleep right away, even though writer 1 wants to acquire before readers 3 &amp; 4.

The `RwLock` implementation in `parking_lot` does not have this inconsistency problem. In fact, it implements task-fair locking on all platforms: [https://docs.rs/parking\_lot/0.10.0/parking\_lot/type.RwLock.html](https://docs.rs/parking_lot/0.10.0/parking_lot/type.RwLock.html) I used this specific implementation as a benchmark for evaluating async `RwLock` implementations.

## The contestants

* [`futures_locks::RwLock`](https://docs.rs/futures-locks/0.5.0/futures_locks/struct.RwLock.html)
* [`async_std::sync::RwLock`](https://docs.rs/async-std/1.5.0/async_std/sync/struct.RwLock.html)
* [`futures_util::lock::RwLock` (currently open PR)](https://github.com/rust-lang/futures-rs/pull/2082)
* [`tokio::sync::RwLock`](https://docs.rs/tokio/0.2.11/tokio/sync/struct.RwLock.html)

edit: Tokio was added after the initial post.

In the case of `futures_locks::RwLock`, readers and writers are queued independently. There is no way for the scheduler to interleave lock acquisitions with this model. An async version of the test shows the expected readers-starve-writers behavior.

`async_std::sync::RwLock` does not queue tasks explicitly but instead relies on a bi-state counter, which can either be "locked for writing" or "locked by *n* readers". The trouble with this implementation is that the bi-state counter can remain in the read state forever. The only requirement for acquiring the read lock is that the lock is not held by a writer. This allows readers to starve writers.

`futures_util::lock::RwLock` has a very similar implementation that we see in `async-std`. Because this is an open (and active!) PR, the author is responsive to my report. And I suspect this implementation will address the starvation problem before it is reviewed.

Finally, `tokio::sync::RwLock` uses a waiter queue where readers acquire 1 permit from a semaphore, and writers acquire all permits. The waiter queue interleaves read and write lock acquisitions in FIFO order, so it does provide a task-fair locking policy! Hooray!

## The results

If you want to use `RwLock` in an async context, be aware that of the implementations I looked at, only Tokio will fairly queue the locks across tasks. A non-fair locking policy is fine in use cases where the lock is held infrequently. But with read-heavy workloads, the writer tasks will be starved.

On the flip-side, crate authors have been receptive! And I'm hopeful that raising awareness on these issues will help everyone.
## [5][WASM Vector Graphics](https://www.reddit.com/r/rust/comments/f574bg/wasm_vector_graphics/)
- url: https://crates.io/crates/wasm_svg_graphics
---

## [6][Introducing Prodash - a terminal dashboard for visualising the progress of concurrent tasks](https://www.reddit.com/r/rust/comments/f4qyvn/introducing_prodash_a_terminal_dashboard_for/)
- url: https://asciinema.org/a/301838
---

## [7][FromStr vs From&lt;&amp;str&gt; vs From&lt;String&gt; Which should I use and when ?](https://www.reddit.com/r/rust/comments/f55ur9/fromstr_vs_fromstr_vs_fromstring_which_should_i/)
- url: https://www.reddit.com/r/rust/comments/f55ur9/fromstr_vs_fromstr_vs_fromstring_which_should_i/
---
Hi,

I have a feeling that FromStr trait overlaps with the more generic From&lt;T&gt; trait. Basically I am a bit confused, when I should choose to implement FromStr and when From&lt;&amp;str&gt;. On top of that, if I choose to implement a From&lt;&amp;str&gt; trait, is there any reason why I should NOT implement the From&lt;String&gt; trait ? And if I implement a From&lt;&amp;str&gt; or a From&lt;String&gt;, is there a reason why I should NOT implement the Into&lt;&amp;str&gt; or Into&lt;String&gt; ?

Edit: For instance the Rust uuid library implements only the FromStr and if you want to convert the UUID to String, you need to use the \`to\_string\` method. Why they chose FromStr, instead of the more generic From&lt;&amp;str&gt; and why they don't provide an Into&lt;String&gt; as well instead of the fixed \`to\_string\` method? There must be a reason, but I just don't see it.
## [8][cow-utils: copy-on-write string utilities for Rust](https://www.reddit.com/r/rust/comments/f4zc5z/cowutils_copyonwrite_string_utilities_for_rust/)
- url: https://github.com/RReverser/cow-utils-rs
---

## [9][Awaiting Closures With Warp](https://www.reddit.com/r/rust/comments/f53rgu/awaiting_closures_with_warp/)
- url: https://www.reddit.com/r/rust/comments/f53rgu/awaiting_closures_with_warp/
---
Hello all.

I am making a web server using warp, and I need to await a function call inside of the .map() call for a path.

I have a minimal example of what I am trying to get working here:

    #![feature(async_closure)]
    
    #[tokio::main]
    async fn main() {
        let path = warp::post()
            .and(warp::path("path"))
            .and(warp::body::json())
            .map(async move |instance: CustomStruct| {
                let output: serde_json::Value = function_that_returns_value().await.unwrap();
                
                warp::reply::json(&amp;output)
            });
    
        warp::serve(path).run(([127, 0, 0, 1], 3030)).await
    }

This does not work because `^^^^^^^ the trait \`warp::reply::Reply\` is not implemented for \`impl core::future::future::Future\``

That is fine, because I really want to use the stable compiler.

The only thought I had to remedy the situation, would be to get the `output` variable stored somehow with a `collect()` call, and then make the `function_that_returns_value()` and `warp::reply::json()` calls outside of the map.

Any help is appreciated.

Cheers :)
## [10][Krabs: x86 bootloader that can boot vmlinux.](https://www.reddit.com/r/rust/comments/f51eao/krabs_x86_bootloader_that_can_boot_vmlinux/)
- url: https://github.com/ellbrid/krabs
---

## [11][Trying out Actix by creating a simple TODO service. Part 3](https://www.reddit.com/r/rust/comments/f50ut2/trying_out_actix_by_creating_a_simple_todo/)
- url: https://youtu.be/3vMxuM7ezEk
---

## [12][This Month in Rust GameDev #6 - January 2020](https://www.reddit.com/r/rust/comments/f4r368/this_month_in_rust_gamedev_6_january_2020/)
- url: https://rust-gamedev.github.io/posts/newsletter-006
---

