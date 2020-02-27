# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (9/2020)!](https://www.reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here_92020/)
- url: https://www.reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here_92020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/f5413m/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (9/2020)?](https://www.reddit.com/r/rust/comments/f8ng5w/whats_everyone_working_on_this_week_92020/)
- url: https://www.reddit.com/r/rust/comments/f8ng5w/whats_everyone_working_on_this_week_92020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-9-2020/38638?u=llogiq)!
## [3][Audit of the RustCrypto `aes-gcm` and `chacha20poly1305` crates by NCC group](https://www.reddit.com/r/rust/comments/fa8a96/audit_of_the_rustcrypto_aesgcm_and/)
- url: https://www.reddit.com/r/rust/comments/fa8a96/audit_of_the_rustcrypto_aesgcm_and/
---
I'm excited to announce that two of the symmetric encryption crates from the [RustCrypto/AEADs](https://github.com/RustCrypto/AEADs) project just received their first security audit 🎉

(NOTE: "RustCrypto" in this context means the Pascal Case GitHub organization, not the unmaintained kebab case `rust-crypto` crate)

In December 2019, [MobileCoin](https://www.mobilecoin.com/) engaged [NCC Group](https://www.nccgroup.trust/) to conduct a review of the following crates which provide pure Rust implementations of these authenticated encryption algorithms, along with their cryptographic dependencies:

- [`aes-gcm`](https://github.com/RustCrypto/AEADs/tree/master/aes-gcm): AEAD cipher implementation
  - [`aes`](https://crates.io/crates/aes): raw AES block cipher with [`aesni`](https://crates.io/crates/aesni) and [`aes-soft`](https://crates.io/crates/aes-soft) backends
  - [`polyval`](https://crates.io/crates/polyval): used to implement GHASH in AES-GCM
- [`chacha20poly1305`](https://github.com/RustCrypto/AEADs/tree/master/chacha20poly1305): AEAD cipher implementation
  - [`chacha20`](https://crates.io/crates/chacha20): raw ChaCha20 stream cipher
  - [`poly1305`](https://crates.io/crates/poly1305): authenticator providing message integrity

The scope of the audit was assessing cryptographic security including constant-time operation of these crates and their cryptographic dependencies, including implementations that leverage x86-64 CPU intrinsics as well as portable implementations based on integer arithmetic. Two consultants provided five person-days of efforts.

Result: there were only minor findings (mostly related to performance).

You can read the full report at:

https://research.nccgroup.com/2020/02/26/public-report-rustcrypto-aes-gcm-and-chacha20poly1305-implementation-review/

See also:

- [`aes-gcm` crate on crates.io](https://crates.io/crates/aes-gcm) | [docs](https://docs.rs/aes-gcm)
- [`chacha20poly1305` crate on crates.io](https://crates.io/crates/chacha20poly1305) | [docs](https://docs.rs/chacha20poly1305)
## [4][Example of calling a Rust library from Clojure, all combined into one executable using GraalVM](https://www.reddit.com/r/rust/comments/faaqyc/example_of_calling_a_rust_library_from_clojure/)
- url: https://www.reddit.com/r/rust/comments/faaqyc/example_of_calling_a_rust_library_from_clojure/
---
Hi! Currently learning Rust. As a fun exercise, I made an example of a Clojure program that calls a Rust library via Java JNI, which is then compiled to an executable with GraalVM.

[https://github.com/borkdude/clojure-rust-graalvm](https://github.com/borkdude/clojure-rust-graalvm-native)
## [5][Bentley's coding challenge: k most frequent words](https://www.reddit.com/r/rust/comments/fa5sbr/bentleys_coding_challenge_k_most_frequent_words/)
- url: https://codegolf.stackexchange.com/questions/188133/bentleys-coding-challenge-k-most-frequent-words
---

## [6][Better Plotting with Plotly in Rust Notebooks](https://www.reddit.com/r/rust/comments/faayaf/better_plotting_with_plotly_in_rust_notebooks/)
- url: https://shahinrostami.com/posts/programming/rust-notebooks/better-plotting-with-plotly/
---

## [7][Crosspost: Rust on RG300 consoles (and similar MIPS devices)](https://www.reddit.com/r/rust/comments/fabgxm/crosspost_rust_on_rg300_consoles_and_similar_mips/)
- url: https://www.reddit.com/r/rust_gamedev/comments/fabgof/wrote_a_rust_program_that_demonstrates_graphics/
---

## [8][This Week in Rust 327](https://www.reddit.com/r/rust/comments/fabtwg/this_week_in_rust_327/)
- url: https://this-week-in-rust.org/blog/2020/02/25/this-week-in-rust-327/
---

## [9][crates.io incident report for 2020-02-20 | Inside Rust Blog](https://www.reddit.com/r/rust/comments/f9tr55/cratesio_incident_report_for_20200220_inside_rust/)
- url: https://blog.rust-lang.org/inside-rust/2020/02/26/crates-io-incident-report.html
---

## [10][Can we write async methods? If we can,How?](https://www.reddit.com/r/rust/comments/fa9q7q/can_we_write_async_methods_if_we_canhow/)
- url: https://www.reddit.com/r/rust/comments/fa9q7q/can_we_write_async_methods_if_we_canhow/
---
I tried searching but couldn't find any place where I can get detailed info on how to write async methods.
## [11][rust-lv2: A safe, fast, and ergonomic framework to create LV2 plugins, written in Rust](https://www.reddit.com/r/rust/comments/f9wzbx/rustlv2_a_safe_fast_and_ergonomic_framework_to/)
- url: https://www.reddit.com/r/rust/comments/f9wzbx/rustlv2_a_safe_fast_and_ergonomic_framework_to/
---
About a year ago, I've released my first project with [LV2](https://lv2plug.in/), a plugin framework called [lv2rs](https://www.reddit.com/r/rust/comments/bawmbs/lv2rs_creating_lv2_plugins_in_rust/). I've received some great feedback and then joined forces with [Prokopyl](https://github.com/prokopyl) and the [RustAudio discourse](https://rust-audio.discourse.group/) to create [rust-lv2](https://crates.io/crates/lv2), the safe, fast, and ergonomic framework to create LV2 plugins, written in Rust, which now received its first release!

Using rust-lv2, you can create LV2 audio processing plugins with much less boilerplate code and worries about pointer and data handling, compared to plugins that use the raw C API directly. Thanks to Rust's incredible type system, several errors a programmer can make with the original utilities simply aren't possible.

Let's have an example. This is a basic plugin that amplifies the audio from an input port and writes it to an output port:

    // Include the prelude of the core crate. Every specification is implemented by
    // exactly one crate and their preludes always contain most of the types needed
    // to use them in a plugin.
    use lv2_core::prelude::*;

    // Most useful plugins will have ports for input and output data. In code, these
    // ports are represented by a struct implementing the `PortContainer` trait.
    #[derive(PortCollection)]
    struct Ports {
        gain: InputPort&lt;Control&gt;,
        input: InputPort&lt;Audio&gt;,
        output: OutputPort&lt;Audio&gt;,
    }

    // Every plugin defines a struct for the plugin instance. All persistent data
    // associated with a plugin instance is stored here, and is available to every
    // instance method. In this simple plugin, there is no additional instance data and
    // therefore, this struct is empty.
    struct Amp;

    // The URI is the identifier for a plugin, and how the host associates this 
    // implementation in code with its description in data. If this URI does not match
    // that used in the data files, the host will fail to load the plugin. Since many 
    // other things are also identified by URIs, there is a separate trait for them:
    // The `UriBound`. It stores the URI as a constant null-terminated byte slice and
    // provides a method to easily retrieve the URI. If the null-terminator is omitted,
    // some other parts of the system may cause undefined behaviour. Since this can not
    // checked by Rust's type system, this trait has to be unsafe. However, this is the
    // only time a user has to use the `unsafe` keyword.
    unsafe impl UriBound for Amp {
        const URI: &amp;'static [u8] = b"urn:rust-lv2-book:eg-amp-rs\0";
    }

    // Every plugin struct implements the `Plugin` trait.
    impl Plugin for Amp {
        // Set the ports type.
        type Ports = Ports;

        // This plugin does not use additional host features and therefore, we set the
        // features collection type to `()`. Other plugins may define a separate struct
        // with their required and optional features and set it here.
        type Features = ();

        // The `new` method is called by the plugin backend when it creates a new 
        // plugin instance. The host passes the plugin URI, sample rate, and bundle 
        // path for plugins that need to load additional resources (e.g. waveforms).
        // The features parameter contains host-provided features defined in LV2 
        // extensions, but this simple plugin does not use any. This method is in the
        // “instantiation” threading class, so no other methods on this instance will 
        // be called concurrently with it.
        fn new(_plugin_info: &amp;PluginInfo, _features: ()) -&gt; Option&lt;Self&gt; {
            Some(Self)
        }

        // The `run()` method is the main process function of the plugin. It processes
        // a block of audio in the audio context.
        fn run(&amp;mut self, ports: &amp;mut Ports) {
            let coef = if *(ports.gain) &gt; -90.0 {
                10.0_f32.powf(*(ports.gain) * 0.05)
            } else {
                0.0
            };

            for (in_frame, out_frame) in Iterator::zip(ports.input.iter(), ports.output.iter_mut()) {
                *out_frame = in_frame * coef;
            }
        }
    }

    // The `lv2_descriptors` macro creates the entry point to the plugin library. It
    // takes structs that implement `Plugin` and exposes them. The host will load the
    // library and call a generated function to find all the plugins defined in the
    // library.
    lv2_descriptors!(Amp);


If you want to learn more about rust-lv2, make sure to check out the [rust-LV2 book](https://janonard.github.io/rust-lv2-book/). It's an effort to translate the [original LV2 book](https://lv2plug.in/book/) and explains how to properly use the framework. Both examples were taken from their respective books.

If you have a question, have a little inconvenience or even encounter a bug, please create an issue in [the Github repository](https://github.com/RustAudio/rust-lv2/). We are striving to improve the usability of the framework and are going to implement more specifications soon.

EDIT: I've replaced every indentation with four spaces, I hope it works now!
EDIT 2: I've understood the comment now and fixed it. Sorry!
## [12][What is the motivation behind default move semantics?](https://www.reddit.com/r/rust/comments/fa9pkp/what_is_the_motivation_behind_default_move/)
- url: https://www.reddit.com/r/rust/comments/fa9pkp/what_is_the_motivation_behind_default_move/
---
I would imagine const ref could have been the default semantic, but I would like to understand why rust chose move instead. I’m sure there is an excellent reason but it’s something I can’t think of now.
