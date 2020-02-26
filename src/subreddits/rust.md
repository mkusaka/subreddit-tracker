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
## [3][Securing Firefox with WebAssembly (and rust)](https://www.reddit.com/r/rust/comments/f9qk28/securing_firefox_with_webassembly_and_rust/)
- url: https://hacks.mozilla.org/2020/02/securing-firefox-with-webassembly/
---

## [4][Space Factory game in Rust + Godot "The Recall Singularity" is progressing well](https://www.reddit.com/r/rust/comments/f9m00v/space_factory_game_in_rust_godot_the_recall/)
- url: https://www.reddit.com/r/rust/comments/f9m00v/space_factory_game_in_rust_godot_the_recall/
---
When I wanted to make a Factory game set in space I surveyed the options and naturally chose Rust as the very best language to simulate the core of factories. Today [I'm publishing a retrospective on the past 3 months of development](https://medium.com/@recallsingularity/space-factory-building-in-feb-2020-752e96dfaa0a) (out of 18) which lays out features developed and challenges along the way.

Rust is a great fit for "scripting" with the Godot Engine both because it offers all the usual guarantees and benefits but also because the [godot rust bindings](https://github.com/GodotNativeTools/godot-rust) work out of the box with minimal work.

Personally I try to give back to the community by writing tutorials covering the technically difficult challenges I've solved. I've got a few ideas lined up but my [getting started with Godot + Rust tutorial](https://medium.com/@recallsingularity/gorgeous-godot-games-in-rust-1867c56045e6?source=friends_link&amp;sk=c2fd85689b4638eae4d91b743439c75f) continues to be popular and I run a [very successful discord on the combination](https://discord.gg/NRFmb8N).

Rust gamedev continues to flourish and I hope we are close to a revolution in the most popular systems language used in major games.
## [5][Intro to rustc's self profiler | Inside Rust Blog](https://www.reddit.com/r/rust/comments/f9g9ds/intro_to_rustcs_self_profiler_inside_rust_blog/)
- url: https://blog.rust-lang.org/inside-rust/2020/02/25/intro-rustc-self-profile.html
---

## [6][veriform: security-oriented Protobuf-like serialization format with no_std support](https://www.reddit.com/r/rust/comments/f9ihqm/veriform_securityoriented_protobuflike/)
- url: https://www.reddit.com/r/rust/comments/f9ihqm/veriform_securityoriented_protobuflike/
---
I'm announcing a "sneak preview" (v0.0.1) release of Veriform: a binary serialization format inspired by Protocol Buffers, but with a number of refinements which make it more useful to things like credentials, signed documents, and "blockchain" use cases:

- GitHub: https://github.com/iqlusioninc/veriform
- Docs: https://docs.rs/veriform
- Crate: https://crates.io/crates/veriform

I was hoping to get it in a slightly more polished state before making an announcement, but after seeing the [mincodec](https://www.reddit.com/r/rust/comments/f8yfuq/mincodec_extremely_spatially_efficient_true_async/) announcement yesterday (which contains a number of similar ideas, but also many differences), I thought I'd go ahead and make an announcement, both since it seems there's a lot of interest in binary encodings which support `no_std` environments, and also because it seems like there's some potential collaboration opportunities between `mincodec` and Veriform.

# Background

One of the main innovations in Protocol Buffers was a schema-driven binary serialization format with support for "unknown fields" - the ability to extend the schema while allowing clients with an older schema to remain backwards compatible. To do this every field in a message tagged with a [wire type](https://developers.google.com/protocol-buffers/docs/encoding#structure), which provides just enough information to allow the parser to calculate the length of an unknown field so it can be skipped.

Veriform extends that additional "wire type" information to make the format fully self-describing. This doesn't require significant changes from Protobufs - mostly just [more precisely tagging the contents of fields](https://github.com/protocolbuffers/protobuf/issues/2629), and marking on-the-wire which fields contain sequences of messages (`repeated` in Protobuf terminology) so a parser can always tell a sequence from a field containing a single message.

By making the format fully self-describing, it allows for computing a [structured content hash over the data](https://github.com/iqlusioninc/veriform/blob/develop/spec/draft-veriform-spec.md#structured-content-hashing-verihash), even in the presence of unknown fields. This allows for computing content hashes of (or signing) arbitrary sub-messages within a document. Computing structured content hashes in this manner is an idea borrowed from Ben Laurie's [ObjectHash](https://crates.io/crates/objecthash).

This sort of content hashing is a bit at odds with data warehousing-oriented serialization formats like Protobufs or [Cap'n Proto](https://capnproto.org/). Data warehousing formats are flexible-by-design to accommodate an evolving schema, and therefore support on-the-wire representations which can be interpreted in different ways depending on how the schema says it should be interpreted. Veriform, on the other hand, encodes all of the information needed to unambiguously recognize the entire message on-the-wire.

Perhaps somewhat surprisingly this can be done just as compactly as Protobufs (or more compactly, as Veriform supports an ASN.1-like `SEQUENCE` type which is encoded more compactly than Protobufs' `repeated` fields) with one exception: Veriform reserves one additional bit per field for flagging that field as "critical". Borrowing from X.509's notion of "critical extensions", Veriform allows for encoding critical fields which are not allowed to be ignored if unknown. This is important if fields are security-critical in nature, e.g. they contain a constraint/caveat/restriction about how a document should be interpreted which might cause security issues if ignored.

# Encoding

NOTE: There's also a [WIP specification of the encoding](https://github.com/iqlusioninc/veriform/blob/develop/spec/draft-veriform-spec.md) if you'd like to know more.

Veriform's encoding is similar to Protobufs. Here is a brief description and some of the things that differ:

## "TLV" format + sequence encoding

Protobufs and Veriform both use a "TLV" ([type-length-value](https://en.wikipedia.org/wiki/Type-length-value)) format which models messages as key/value maps. Where Protobufs model sequences (i.e. vectors/arrays) of data in messages by allowing them to be `repeated`, Veriform has a distinct, unambiguous sequence encoding (which should be more compact than Protobufs in most cases)

## 64-bit varints for framing and encoding integers

Like Protobufs, Veriform elects to use varints for field identifiers, length tags, and encoding all (unsigned and signed) integers. Message compactness is important for Veriform's use cases like credentials, as using fixed-sized integers can result in a surprisingly large amount of excess space. As an example, the [Macaroons](http://macaroons.io/) format started out using a fixed-sized integer encoding (git pkt-lines), but then moved to a more Protobuf-like V2 format after experiments showed it would save considerable space. Since Veriform fully processes all messages for security reasons (examining "critical bits" and for content hashing) it can't implement something like Cap'n Proto's "infinitely faster" parser, so between that and the space savings varints are worth it.

Veriform uses a "prefix varint" encoding called [`vint64`](https://github.com/iqlusioninc/veriform/tree/develop/rust/vint64) which stores the length of the varint itself in the least significant bits of the first byte, using CTZ to count them. This allows them to be decoded without loops - just a load, shift, and a mask, and the length can be deduced from the first byte alone, both of which are advantages over the LEB128 format used by Protobufs. It was nice to see that `mincodec` is doing something similar in their "[VLQ](https://github.com/noocene/bitbuf-vlq)" integer encoding, which looks quite similar to (possibly identical to?) `vint64` aside from some details of the actual implementation. To the `mincodec` folks if you're reading this: perhaps we can spec a common encoding and share a common implementation?

## Fully self-describing

As noted above, the full type of every field is unambiguously included on the wire. They are encoded as one of 8 "wire types", which are as follows:

| ID   | Type                    | Encoding               |
|------|-------------------------|------------------------|
| 0    | false                   | none                   |
| 1    | true                    | none                   |
| 2    | unsigned 64-bit integer | vint64                 |
| 3    | signed 64-bit integer   | vint64 (zigzag)        |
| 4    | bytes                   | length prefixed        |
| 5    | string (unicode)        | length prefixed UTF-8  |
| 6    | message (nested)        | length prefixed        |
| 7    | sequence                | length + type prefixed |

Note that unlike Protos (and borrowing a trick similar to Cap'n Proto), "false" and "true" are given their own wire types to allow boolean fields to be encoded as a single bit (Protobufs encodes each boolean value as a byte in addition to the field header).

# Implementation

[We're using Veriform](https://github.com/iqlusioninc/armistice/pull/8/files) in a project called [Armistice](https://github.com/iqlusioninc/armistice), a bare-metal cryptography / "HSM" application for the [USB armory MkII](https://github.com/iqlusioninc/usbarmory.rs). We're presently treating it as a `no_std` environment (and leveraging Realtime For the Masses), and as such Veriform has first class `no_std` support with no `alloc` requirement. This should make it interesting to anyone looking for a microcontroller-friendly encoding.

The decoder is implemented as a streaming zero-copy pull parser which is capable of incrementally decoding messages. It was quite interesting to see `mincodec` natively support async message processing - we don't do that yet, but the decoder is written in such a way which would make it easy to wrap with async support in the future.

Internally it uses a hand-written state machine which always consumes its previous state and computes a new state (in a similar manner to session types), and optionally emits events if it was able to make progress decoding the message. The implementation manages to be extremely compact despite the feature set, presently weighing in at 1.1klocs, again good for anyone looking to use it on a microcontroller.

# Future Work

Veriform is still very rough-around-the-edges. Nevertheless the current implementation, albeit low-level, is pretty much feature complete at its core. That said, you may want to wait for the following before trying it out:

- custom derive support (ala Prost and `mincodec`): using annotations on the Rust code to synthesize encoders/decoders has always been the plan. Right now we're writing the encoders/decoders by hand in Armistice as we don't want to jump to a proc macro before we've spent a lot of time handwriting the code, making sure we have a good API for the proc macro codegen to target, and looking at ways to simplify and minimize the proc macro generated code.
- schema support + other language implementations: we'd like to support a schema language ala Protobufs to eventually support implementations in other languages. This is definitely one of the big selling points of Protobufs.
## [7][Anouncing async_executors, a building block for executor agnostic libraries](https://www.reddit.com/r/rust/comments/f9t44x/anouncing_async_executors_a_building_block_for/)
- url: https://blog.wnut.pw/2020/02/25/anouncing-async_executors-a-building-block-for-executor-agnostic-libraries/
---

## [8][orion 0.15 - Argon2i, pure-Rust secretstream, extensive testing and benchmarks](https://www.reddit.com/r/rust/comments/f9o98k/orion_015_argon2i_purerust_secretstream_extensive/)
- url: https://www.reddit.com/r/rust/comments/f9o98k/orion_015_argon2i_purerust_secretstream_extensive/
---
&gt; orion is a cryptography library written in pure Rust. It aims to provide easy and usable crypto while trying to minimize the use of unsafe code.

GitHub: https://github.com/brycx/orion 
Crates.io: https://crates.io/crates/orion

After 8 months since the last major update, I'm very pleased to announce orion has finally reached `0.15`.

Overview of notable changes since the [last update (v0.14)](https://www.reddit.com/r/rust/comments/c3pn72/orion_014_released_improvements_and_call_for/):

- Focus on improving and stabilizing the API.
- Added support for Argon2i(single-threaded).
- Added support for a pure-Rust implementation of libsodiums secretstream. Thanks to [snsmac](https://github.com/snsmac) for driving this effort.
- Major improvements to the testing of the library. This includes a completely new testing framework, constant-time tests running in CI on OSX and Windows, as well as constant-time fuzzing using WASM with [sidefuzz](https://github.com/phayes/sidefuzz) in [orion-sidefuzz](https://github.com/brycx/orion-sidefuzz).
- The library is now listed as an alternative to `rust-crypto` in the RustSec advisory database.
- Performance improvements and newly [published benchmarks in the wiki](https://github.com/brycx/orion/wiki/Benchmarks) to give an overview of the performance that can be expected. Benchmarks include a daily workstation and Raspberry Pi 2. 

([See full changelog here](https://github.com/brycx/orion/blob/master/CHANGELOG.md#0150))

A huge thanks to the existing and the three new contributors that helped with this release and make the library a lot better!
## [9][I broke my rust installation](https://www.reddit.com/r/rust/comments/f9rb39/i_broke_my_rust_installation/)
- url: https://www.reddit.com/r/rust/comments/f9rb39/i_broke_my_rust_installation/
---
Hello, after updating my rust installation with `rustup update` something is wrong

    rustup default stable
    info: using existing install for 'stable-x86_64-apple-darwin'
    info: default toolchain set to 'stable-x86_64-apple-darwin'
    
      stable-x86_64-apple-darwin unchanged - rustc 1.41.0 (5e1a79984 2020-01-27)
    
    info: syncing channel updates for '1.49.0-x86_64-apple-darwin'
    info: downloading component 'rust'
    error: could not download nonexistent rust version `1.49.0-x86_64-apple-darwin`
    error: caused by: could not download file from 'https://static.rust-lang.org/dist/rust-1.49.0-x86_64-apple-darwin.tar.gz.sha256' to '/Users/jan/.rustup/tmp/fq3t2xbgx12vq1w0_file'
    error: caused by: http request returned an unsuccessful status code: 404

I removed and added the stable toolchain, but without luck. Where does the rust version `1.49.0` came from? Latest version is `1.41.0`.  Any suggestion? Thanks for help.
## [10][bottom: a cross platform TUI process/system monitor inspired by gotop and gtop](https://www.reddit.com/r/rust/comments/f99tdg/bottom_a_cross_platform_tui_processsystem_monitor/)
- url: https://github.com/ClementTsang/bottom
---

## [11][New post: Async HTTP](https://www.reddit.com/r/rust/comments/f9dg6j/new_post_async_http/)
- url: https://blog.yoshuawuyts.com/async-http/
---

## [12][keyed_priority_queue 0.1.3 released! (major performance upgrade)](https://www.reddit.com/r/rust/comments/f9rss2/keyed_priority_queue_013_released_major/)
- url: https://www.reddit.com/r/rust/comments/f9rss2/keyed_priority_queue_013_released_major/
---
At first I will talk about crate then I wish some advice about further improvements.

The crate allows changing priority of items in queue using unique key.

[https://crates.io/crates/keyed\_priority\_queue](https://crates.io/crates/keyed_priority_queue)

I had benchmarked my optimizations (all numbers against v0.1.2):

* from 20% to 40% faster construction of queue from iterator
* nearly 50% faster popping of elements from queue
* nearly 50% faster pushing in worst cases
* 10% faster random item removal
* from 10% to 30% faster random item priority change

But some things are now worse (however they was very fast and this is not a problem):

* getting the priority of existing element slower by 50% if keys and elements are usizes and 10% if strings
* random pushes sometimes slower in 70% for usizes

Benches are available in github [repository](https://github.com/AngelicosPhosphoros/keyed_priority_queue/tree/develop/benches)

**How it was optimized**

In previous versions, the main crate's type — KeyedPriorityQueue&lt;Key, Value&gt; — had very straightforward structure: binary heap with key-priority pairs and mapping from keys to heap element positions. Entries can be easily found using hashmap for **O(1)** opertations and the costs of taking or putting data into queue remain was **O(log n)**.

However, there was a big trouble with performance: to keep hashmap indices valid I needed to make O(log n) updates on hashmap which were very costly.

In the v0.1.3, I added intermediate vec: values in hashmap point to the indices of an intermediate vec, and the vec items point to the positions in heap. This significantly improve speed of many operations, since many operations require only one hash lookup (*pop, set\_priority*) or two (*remove\_item by key*). Updates in the intermediate vec require simple index lookups.

However, this intermediate Vec adds extra indirection which cause performance drawbacks on some operations with exactly one hashmap lookup like *get\_priority.*

**Futher plans (for 0.2.0)**

1. Improve API: replace panicing method by \`Result\`, add Entry API (like in HashMap).
2. And two possible enhancments but I need your advice about them

**First needed advice: use IndexMap instead of HashMap**

I can introduce Indexmap from indexmap crate instead of both hashmap and intermediate vec.

* PRO: it greatly simplifies current code
* PRO: it removes need of keeping two copies of each key so I will be able to remove \`Clone\` trait requirement and lower memory usage.
* PRO: in methods with two hash lookups, one of them can be replaced by index lookup which is faster. In *pop,* it removes need of hash lookups at all
* CON: it makes \`get\_priority\` and \`push\` operations even slower than after current modification.
* CON: it adds dependency to the crate (currently crate depends only on std)

The speed improvements are very insignificant, IMHO, because it speedups methods which already fast (e.g. popping 1000 usizes from 500'000 improved from 745 microseconds to 471 us).  Also it will add some pessimisation, e.g. pushing 1000 random usizes regressed from 5.2 milliseconds (on v0.1.3) to 6.1 ms.Are reduction memory usage and removal \`Clone\` trait requirement on keys worth such speed degradation on pushes?

**Second needed advice: Create special version for usizes**

There is an issue in my github where user Ten0 asked to add a new structure that has almost same API but uses usizes as keys instead any type and doesn't use hash maps at all. It can be achieved by using just one vec over map and keeping special \`invalid\` usize in places for evicted keys.

* PRO: it will be useful if someone already have some domain values mapped to usizes
* PRO: it obviously will be faster than any hashmap due lack of hashing and one level of indirection
* CON: user can make mistakes using this API very easily: since we need to keep remapping vec for all range of ever used keys, user can allocate a lot of memory by passing some key like 2\^50.
* CON: it increase the scope of the crate

Do I need put any effort to implement this?

Thanks in advance and sorry for my English!
