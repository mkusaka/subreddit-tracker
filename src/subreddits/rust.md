# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (16/2020)!](https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 334](https://www.reddit.com/r/rust/comments/g1fj7p/this_week_in_rust_334/)
- url: https://this-week-in-rust.org/blog/2020/04/14/this-week-in-rust-334/
---

## [3][μfmt - a smaller, faster and panic-free alternative to core::fmt](https://www.reddit.com/r/rust/comments/g242r1/μfmt_a_smaller_faster_and_panicfree_alternative/)
- url: https://github.com/japaric/ufmt
---

## [4][regex2fat: convert regexes to FAT32, with regex-automata and a few lines of safe Rust](https://www.reddit.com/r/rust/comments/g2dxpw/regex2fat_convert_regexes_to_fat32_with/)
- url: https://github.com/8051Enthusiast/regex2fat
---

## [5][Which mqtt rust library do you recommend?](https://www.reddit.com/r/rust/comments/g2c75e/which_mqtt_rust_library_do_you_recommend/)
- url: https://www.reddit.com/r/rust/comments/g2c75e/which_mqtt_rust_library_do_you_recommend/
---
These three crates seem to be the best right now:

* [https://github.com/eclipse/paho.mqtt.rust](https://github.com/eclipse/paho.mqtt.rust)
* [https://github.com/zonyitoo/mqtt-rs](https://github.com/zonyitoo/mqtt-rs)
* [https://github.com/tekjar/rumq](https://github.com/tekjar/rumq)

Which one do you recommend? Thanks!
## [6][mini-redis: A Redis client and server implementation using Tokio. Includes extensive comments for learning.](https://www.reddit.com/r/rust/comments/g1vpo9/miniredis_a_redis_client_and_server/)
- url: https://github.com/tokio-rs/mini-redis
---

## [7][New Crate: Interactive Chord Diagrams (+ Rust Notebook Support)](https://www.reddit.com/r/rust/comments/g2cvqt/new_crate_interactive_chord_diagrams_rust/)
- url: https://shahinrostami.com/posts/programming/rust-notebooks/chord-diagrams/
---

## [8][Governance Working Group Update: Meeting 09 April 2020 | Inside Rust Blog](https://www.reddit.com/r/rust/comments/g2ch1q/governance_working_group_update_meeting_09_april/)
- url: https://blog.rust-lang.org/inside-rust/2020/04/14/Governance-WG-updated.html
---

## [9][Reliable benchmarks](https://www.reddit.com/r/rust/comments/g29kz8/reliable_benchmarks/)
- url: https://www.reddit.com/r/rust/comments/g29kz8/reliable_benchmarks/
---
I'm working to optimize some code using Criterion to benchmark, and I'm finding the benchmark results difficult to make sense of. The same comparison yields improvement one time, regression another. I'm finding this even after fixing my CPU frequency and killing most other processes. Code paths that haven't changed in years are getting "significant" performance differences. Do you all have any other tips on how to make benchmarks reliable? I'm hoping to commandeer a system dedicated to benchmarking soon, so I'll have a lot of flexibility in how things are configured.
## [10][aeron-rs - efficient reliable UDP and IPC message transport for Rust](https://www.reddit.com/r/rust/comments/g1qrpb/aeronrs_efficient_reliable_udp_and_ipc_message/)
- url: https://www.reddit.com/r/rust/comments/g1qrpb/aeronrs_efficient_reliable_udp_and_ipc_message/
---
Hi, all!

Having fast and reliable messaging is critically important for high performance distributed systems. There are number of message brokers already available for Rust. But what if you need something faster?

There is Aeron system ([https://github.com/real-logic/aeron](https://github.com/real-logic/aeron)) which can transport messages fast through IPC (shared memory) and UDP. But until recently there were no Rust interface for that. Our team from United Traders invested some efforts to obtain so called client library for Aeron written in Rust. So now it is possible to talk from Rust program (through Aeron media driver written in C) with any other Aeron enabled program (written in Rust, C++ or Java).

Our implementation of Aeron client library (called aeron-rs [https://crates.io/crates/aeron-rs](https://crates.io/crates/aeron-rs) ) is a port of C++ library to Rust. While it is completely written in Rust it also inherits design and structure of its C++ ancestor therefore some things are done not in a Rust way. I believe there are plenty improvements which can be done so it will be good if interested people will look at aeron-rs and contribute to it.

As of today aeron-rs 0.1.0 is tested (200+ unit tests, 6 integration test, number of test applications) but was not applied so far in any production environment. You are welcome to give it a try and to propose/implement improvements and fixes.

Comparative performance tests show that (on the same machine) with messages of 32 bytes length C++ Aeron client gives throughput of 6-8 millions messages per second, while aeron-rs just 1-2 millions. On 1K sized message C++ version delivers 500K msgs/sec and aeron-rs 250K msgs/sec. So from performance perspective we also need to catch up and overtake С++ :-)

Thanks!
## [11][erupt Vulkan API bindings v0.1.0 released!](https://www.reddit.com/r/rust/comments/g1tqh3/erupt_vulkan_api_bindings_v010_released/)
- url: https://www.reddit.com/r/rust/comments/g1tqh3/erupt_vulkan_api_bindings_v010_released/
---
erupt provides bindings to the Vulkan API

Links: [docs.rs](https://docs.rs/erupt/0.1.0/erupt/) [Repository](https://gitlab.com/Friz64/erupt)

While other good bindings to Vulkan exist (most prominently ash), I just released my own! My motivations were getting to know what it's like to write Vulkan bindings and trying to improve on features where ash falls behind.

**Features of erupt**

* Full Vulkan API coverage
* First-class support for all extensions
* High quality auto-generated function wrappers
* A [utility module](https://docs.rs/erupt/*/erupt/utils/index.html) aiding your use of the Vulkan API
   * [VulkanResult](https://docs.rs/erupt/*/erupt/utils/struct.VulkanResult.html): Idiomatic wrapper around a Vulkan Result
   * [surface](https://docs.rs/erupt/*/erupt/utils/surface/index.html): Create a [SurfaceKHR](https://docs.rs/erupt/*/erupt/extensions/khr_surface/struct.SurfaceKHR.html) using a [RawWindowHandle](https://docs.rs/raw-window-handle/*/raw_window_handle/enum.RawWindowHandle.html) (adapted from [ash-window](https://crates.io/crates/ash-window))
* Generated code distributed into multiple modules
* Function loading ([CoreLoader](https://docs.rs/erupt/*/erupt/struct.CoreLoader.html), [InstanceLoader](https://docs.rs/erupt/*/erupt/struct.CoreLoader.html), [DeviceLoader](https://docs.rs/erupt/*/erupt/struct.CoreLoader.html))
* Seperate Flags and FlagBits types
* A high level Builder for every struct
* Type-safe pointer chain support
* Default and Debug implementation for every type
* Complete auto-generation of everything except [utils](https://docs.rs/erupt/*/erupt/utils/index.html)

**Example: Instance Creation**

    use erupt::{vk1_0::*, CoreLoader, InstanceLoader};
    
    let mut core = CoreLoader::new()?;
    core.load_vk1_0()?;
    
    let app_info = ApplicationInfoBuilder::new().api_version(erupt::make_version(1, 0, 0));
    let instance_info = InstanceCreateInfoBuilder::new().application_info(&amp;app_info);
    let instance_handle = core
        .create_instance(&amp;instance_info, None, None)
        .expect("Failed to create instance");
    
    let mut instance = InstanceLoader::new(&amp;core, instance_handle)?;
    instance.load_vk1_0()?;
    
    // ...
    
    instance.destroy_instance(None);

**Other examples**

* [triangle](https://gitlab.com/Friz64/erupt/-/blob/master/erupt-examples/src/triangle.rs)
* [pointer-chain](https://gitlab.com/Friz64/erupt/-/blob/master/erupt-examples/src/pointer_chain.rs)
* [version](https://gitlab.com/Friz64/erupt/-/blob/master/erupt-examples/src/version.rs)

**Thank you**

* [vk-parse](https://crates.io/crates/vk-parse) for helping parse vk.xml in the [generator](https://gitlab.com/Friz64/erupt/-/tree/master/generator)
* [ash](https://crates.io/crates/ash) for helping inspiring and making this crate
* [libloading](https://crates.io/crates/libloading) for providing symbol loading
* [ash-window](https://crates.io/crates/ash-window) for providing a base for the [surface](https://docs.rs/erupt/*/erupt/utils/surface/index.html) module
* [bitflags](https://crates.io/crates/bitflags) for providing a perfect bitflag macro
* The Vulkan Community ❤️
* The Rust Community ❤️

I plan on adding more features to the [utils](https://docs.rs/erupt/*/erupt/utils/index.html) module, for example a basic Vulkan Memory Allocator to help you in write your Vulkan code.

Fun fact: This is my first crate. I released it almost exactly 2 years after starting to learn Rust :)
## [12][how to use hashmaps like this in Rust map[pos]+=elem?](https://www.reddit.com/r/rust/comments/g27vi0/how_to_use_hashmaps_like_this_in_rust_mapposelem/)
- url: https://www.reddit.com/r/rust/comments/g27vi0/how_to_use_hashmaps_like_this_in_rust_mapposelem/
---
how can I use hashmaps like this in Rust map\[pos\]+=elem? I know that I can use map.insert(pos, elem) which is equivalent to map\[pos\]=elem.
