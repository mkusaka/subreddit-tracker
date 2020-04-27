# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (17/2020)!](https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g4nu6j/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 335](https://www.reddit.com/r/rust/comments/g6d0ac/this_week_in_rust_335/)
- url: https://this-week-in-rust.org/blog/2020/04/21/this-week-in-rust-335/
---

## [3][Teleforking a process onto a different computer!](https://www.reddit.com/r/rust/comments/g8v041/teleforking_a_process_onto_a_different_computer/)
- url: https://thume.ca/2020/04/18/telefork-forking-a-process-onto-a-different-computer/
---

## [4][First official release of rust-analyzer](https://www.reddit.com/r/rust/comments/g8zzdy/first_official_release_of_rustanalyzer/)
- url: https://rust-analyzer.github.io/blog/2020/04/20/first-release.html
---

## [5][Cargo Bloat Action: Track Rust binary sizes across builds/pull requests using Github Actions](https://www.reddit.com/r/rust/comments/g8xgoh/cargo_bloat_action_track_rust_binary_sizes_across/)
- url: https://github.com/orf/cargo-bloat-action/
---

## [6][(Almost) Lockless Stream Buffering](https://www.reddit.com/r/rust/comments/g8qgkj/almost_lockless_stream_buffering/)
- url: https://mcfelix.me/blog/shared-buffers/
---

## [7][Maximum Sum of Subarrays (Leetcode) in Rust (+ benchmarking)](https://www.reddit.com/r/rust/comments/g8y2w8/maximum_sum_of_subarrays_leetcode_in_rust/)
- url: https://www.youtube.com/watch?v=G1deF4Rehlw
---

## [8][Rust Wiki: Slack workspace for coordination](https://www.reddit.com/r/rust/comments/g8j871/rust_wiki_slack_workspace_for_coordination/)
- url: https://www.reddit.com/r/rust/comments/g8j871/rust_wiki_slack_workspace_for_coordination/
---
~~This is in response to [this](https://www.reddit.com/r/rust/comments/g7s4ss/rust_needs_a_wiki/) post where people expressed support for creating a Rust wiki.~~

~~For anyone that is interested in helping us create this wiki, please join the new [Slack workspace](https://join.slack.com/t/rustwiki/shared_invite/zt-e674kvf4-vIHNSAG3nTLbG4ef7ECnhA) so we can discuss options and coordinate efforts.~~

We've moved to the community Discord under the #wiki channel: https://discord.gg/fwSRAx

Thank you!
## [9][Do vec![0; n] and Vec::with_capacity(n) sometimes not allocate?](https://www.reddit.com/r/rust/comments/g8o113/do_vec0_n_and_vecwith_capacityn_sometimes_not/)
- url: https://www.reddit.com/r/rust/comments/g8o113/do_vec0_n_and_vecwith_capacityn_sometimes_not/
---
My understanding was that initializing a vector with `vec!` or `Vec::with_capacity` performed an initial allocation and guaranteed that no other allocations would take place so long as the Vector's length stays below that initial capacity.

However, I have been running some tests, and the actual memory use of my program isn't consistent with that understanding.  As a test, I initialized a Vector in three different ways:

1) `Vec::with_capacity(0x4_000_000)`

2) `vec![0_u8; 0x4_000_000]`

3) `vec![47_u8; 0x4_000_000]`

In each case, I complied my program in release mode, ran it, and measured its total memory use with [ps_mem](https://github.com/pixelb/ps_mem).  When I'd initialized the Vector in either of the first two ways, my program's memory use was unchanged from not creating the large Vector at all, but when I initialized it using #3, my program's memory use increased by the expected amount (~67 MiB).

Does that mean that `with_capacity` does *not* actually allocate the full capacity (and thus could still require subsequent allocations to reach that capacity)?  Or is it allocating that memory, but not in a way that ps_mem is counting as used (which seems odd – it the memory has been allocated, it's not available for other programs, which seems like use to me).

(Before anyone says anything, my *actual* use case will not involve allocating a 67 MiB Vector!  I'm just testing with large numbers because they make the effects more obvious :)  But I would like to understand how I need to initialize a vector to guarantee that I can subsequently use it without triggering any allocations)
## [10][Creating a buffer with specific indices for UDP messaging](https://www.reddit.com/r/rust/comments/g8v7xi/creating_a_buffer_with_specific_indices_for_udp/)
- url: https://www.reddit.com/r/rust/comments/g8v7xi/creating_a_buffer_with_specific_indices_for_udp/
---
Hey there! Hope we're all keeping healthy and happy.

I'm looking to implement a torrenting client just for experience. As part of that, I need to create a buffer that is looks like this [(from BEP)](http://www.bittorrent.org/beps/bep_0015.html):

| Offset | Size           | Name           | Value                           |   |
|--------|----------------|----------------|---------------------------------|---|
| 0      | 64-bit integer | protocol_id    | 0x41727101980 // magic constant |   |
| 8      | 32-bit integer | action         | 0 // connect                    |   |
| 12     | 32-bit integer | transaction_id | // Randomised ID                |   |
| 16     |                |                |                                 |   |

The thing is, I can't find a library that specifically allows me to set offsets for these buffer elements. Has anyone used something similar?

Thanks for your time with this, it's super appreciated!
Jake.
## [11][Educational Rust live coding - Building a git CLI [video]](https://www.reddit.com/r/rust/comments/g8z0xg/educational_rust_live_coding_building_a_git_cli/)
- url: https://youtu.be/YFzF1AHYjes?t=301
---

## [12][Confused with how 'any()' works on an iterator.](https://www.reddit.com/r/rust/comments/g8xp17/confused_with_how_any_works_on_an_iterator/)
- url: https://www.reddit.com/r/rust/comments/g8xp17/confused_with_how_any_works_on_an_iterator/
---
I'm screwing around with Rust now and learning iterators. All other iterator methods like filter, find and map have worked exactly as I'd expect them to, but I can't use 'any' properly. My current implementation is like this:

`for i in subscriptions.iter().any(|&amp;x| x.price &lt; 10.0) {`  
   
    `}`

Where 'subscriptions' is a vector of custom structs. This throws a `bool is not an iterator` error, which sounds odd considering that every other operation I've done on an iterator has resolved to a bool at some point. I can't tell exactly why this particular method isn't working like the others, but I have to use it for an assignment, so any help is greatly appreciated.
