# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (24/2020)!](https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 342](https://www.reddit.com/r/rust/comments/h0mt9k/this_week_in_rust_342/)
- url: https://this-week-in-rust.org/blog/2020/06/10/this-week-in-rust-342/
---

## [3][Resources for learning gtk-rs](https://www.reddit.com/r/rust/comments/h0umnd/resources_for_learning_gtkrs/)
- url: https://www.reddit.com/r/rust/comments/h0umnd/resources_for_learning_gtkrs/
---
I want to learn how to make a GTK app in rust with gtk-rs. Where did you learn using gtk-rs an what tools are you using. I am reading the docs and looking at the examples, but are there any better resources like The Book for learning rust?
## [4][[ANN] New RustCrypto releases: `aead`, `block-cipher`, `crypto-mac`, `digest`, `signature`, `stream-cipher`, and more!](https://www.reddit.com/r/rust/comments/h0em5n/ann_new_rustcrypto_releases_aead_blockcipher/)
- url: https://www.reddit.com/r/rust/comments/h0em5n/ann_new_rustcrypto_releases_aead_blockcipher/
---
The [RustCrypto "2.0" Project](https://github.com/RustCrypto) (i.e. the GitHub org, not the unmaintained `rust-crypto` crate) has been hard at work doing updates and releases of practically every single one of our crates.

Almost all of our crates are written in pure Rust (sometimes with optional ASM acceleration) and nearly every single one is `no_std`/embedded-friendly and should work fine even in heapless/microcontroller environments.

Many of these crates are widely used, and for many of them this is the first major update in nearly two years, so heads up: there are new releases out, and many of them contain breaking changes.

[Follow us on Twitter at @RustCryptoOrg!](https://twitter.com/rustcryptoorg)

## Traits

The RustCrypto project is organized using a crate-per-algorithm model where each algorithm family has a crate with its own set of traits which all of the algorithm implementations depend on. We've bumped all of our traits crates. Here are the release notes:

- [`aead` v0.3.0](https://github.com/RustCrypto/traits/pull/174): authenticated encryption with associated data i.e. high-level symmetric encryption (e.g. AES-GCM)
- [`block-cipher` v0.7.0](https://github.com/RustCrypto/traits/pull/168): block cipher traits (renamed from `block-cipher-traits`)
- [`crypto-mac` v0.8.0](https://github.com/RustCrypto/traits/pull/169): message authentication codes (MACs)
- [`digest` v0.9.0](https://github.com/RustCrypto/traits/pull/186): digest algorithms a.k.a hash functions
- [`elliptic-curve` v0.4.0](https://github.com/RustCrypto/elliptic-curves/pull/39): elliptic curve cryptography (presently Weierstrass-only)
- [`signature` v1.1.0](https://github.com/RustCrypto/traits/pull/187): digital signature algorithms
- [`stream-cipher` v0.4.0](https://github.com/RustCrypto/traits/pull/170): stream cipher traits
- [`universal-hash` v0.4.0](https://github.com/RustCrypto/traits/pull/171): universal hash functions (e.g. GHASH)

Some notable highlights of cross-cutting changes:

- We've put a lot of work into documentation improvements, improving README.mds, CHANGELOG.mds, and `rustdoc`
- All crates are now MSRV 1.41 as we've upgraded to [`generic-array` v0.14](https://github.com/fizyk20/generic-array)
- We've updated to Rust 2018 edition across the board, which should make development easier
- Crates including `block-cipher`, `crypto-mac`, `stream-cipher`, and `universal-hash` now have separate traits for instantiation (`New*`) versus the core algorithm. This should be exciting news for embedded users who want to support hardware cryptographic accelerators, or anyone who wants to build abstractions around hardware-backed keys e.g. HSMs
- We've adopted "Initialize-Update-Finalize" (IUF) nomenclature across the board for all crates. This is a commonly used set of names in cryptographic API design. If you've been using any crate which depends on `crypto-mac`, `digest`, or `universal-hash` crates, you'll need to make the following changes:
    - `Input::input` is now `Update::update`
    - Methods named `*result*` to finish computing a digest or MAC have been renamed to `finalize`
    - All crates now use `Output` as the name of the algorithm's output, i.e. `digest::Output`, `crypto_mac::Output` (formerly `MacResult`)
- All crates now re-export `cratename::generic_array::typenum::consts` as `cratename::consts` for convenience
- Type aliases now exist for several places where `GenericArray` was mandatory before
- Though nearly every single one of our crates is designed from the ground up to be `no_std` friendly, we have several optional features which perform heap allocations. These features are now available in `no_std` environments with a heap under an `alloc` feature.

## AEAD Algorithms (High-level Symmetric Encryption)

We've released new versions of the following [AEAD crates](https://github.com/RustCrypto/AEADs):

- [`aes-gcm` v0.6.0](https://github.com/RustCrypto/AEADs/pull/160)
- [`aes-gcm-siv` v0.5.0](https://github.com/RustCrypto/AEADs/pull/161)
- [`aes-siv` v0.3.0](https://github.com/RustCrypto/AEADs/pull/162)
- [`chacha20poly1305` v0.5.0](https://github.com/RustCrypto/AEADs/pull/163)
- [`xsalsa20poly1305` v0.4.0](https://github.com/RustCrypto/AEADs/pull/164)

We've also released [`crypto_box` v0.2.0](https://github.com/RustCrypto/AEADs/pull/165) which provides a pure Rust implementation of NaCl's `crypto_box` public key hybrid cryptosystem.

## Block Ciphers

We've released new versions of the following [block cipher crates](https://github.com/RustCrypto/block-ciphers):

- [`aes`](https://github.com/RustCrypto/block-ciphers/tree/master/aes) v0.4.0
  - [`aesni`](https://github.com/RustCrypto/block-ciphers/tree/master/aes/aesni) v0.7.0
  - [`aes-soft`](https://github.com/RustCrypto/block-ciphers/tree/master/aes/aes-soft) v0.4.0
- [`blowfish`](https://github.com/RustCrypto/block-ciphers/tree/master/blowfish) v0.5.0
- [`cast5`](https://github.com/RustCrypto/block-ciphers/tree/master/cast5) v0.7.0
- [`des`](https://github.com/RustCrypto/block-ciphers/tree/master/des) v0.4.0
- [`idea`](https://github.com/RustCrypto/block-ciphers/tree/master/idea) v0.1.0
- [`kuznyechik`](https://github.com/RustCrypto/block-ciphers/tree/master/kuznyechik) v0.3.0
- [`magma`](https://github.com/RustCrypto/block-ciphers/tree/master/magma) v0.3.0
- [`rc2`](https://github.com/RustCrypto/block-ciphers/tree/master/rc2) v0.4.0
- [`serpent`](https://github.com/RustCrypto/block-ciphers/tree/master/serpent) v0.1.0
- [`sm4`](https://github.com/RustCrypto/block-ciphers/tree/master/sm4) v0.1.0
- [`twofish`](https://github.com/RustCrypto/block-ciphers/tree/master/twofish) v0.3.0
- [`threefish`](https://github.com/RustCrypto/block-ciphers/tree/master/threefish) v0.1.0

As well as:

- [`block-modes`](https://github.com/RustCrypto/block-ciphers/tree/master/block-modes) v0.4.0 (support for CBC, CFB, PCBC, and OFB modes)

## Digest Algorithms (a.k.a. Hash Functions)

We're still in the process of releasing crates in our [hashes repo](https://github.com/RustCrypto/hashes) and are still doing some minor updates, but in the meantime, we've cut the following releases:

- [`blake2` v0.9.0](https://github.com/RustCrypto/hashes/pull/165)
- [`k12` v0.1.0](https://github.com/RustCrypto/hashes/pull/159) (KangarooTwelve)
- [`md-5` v0.9.0](https://github.com/RustCrypto/hashes/pull/156)
- [`ripemd160` v0.9.0](https://github.com/RustCrypto/hashes/pull/162)
- [`sha-1` v0.9.0](https://github.com/RustCrypto/hashes/tree/master/sha1)
- [`sha2` v0.9.0](https://github.com/RustCrypto/hashes/tree/master/sha2)
- [`sha3` v0.9.0](https://github.com/RustCrypto/hashes/pull/163)

Expect releases of the other hash function crates in the next few days.

## Digital Signatures

We've released v0.6.0 of the [`ecdsa`](https://github.com/RustCrypto/signatures/tree/master/ecdsa) crate, which presently only provides types for ECDSA signatures. We hope to implement the ECDSA algorithm soon in a manner which is generic over the elliptic curves below. [Follow along here if you're interested](https://github.com/RustCrypto/elliptic-curves/issues/40).

## Elliptic Curves

We've released new versions of the following [elliptic curve crates](https://github.com/RustCrypto/elliptic-curves):

- [`k256` v0.3.0](https://github.com/RustCrypto/elliptic-curves/pull/41): secp256k1
- [`p256` v0.3.0](https://github.com/RustCrypto/elliptic-curves/pull/42): NIST P-256
- [`p384` v0.2.0](https://github.com/RustCrypto/elliptic-curves/pull/43): NIST P-384

We're also discussing [traits for elliptic curve arithmetic which allow algorithms to be implemented generically over the underlying elliptic curve](https://github.com/RustCrypto/elliptic-curves/issues/22). Follow along if that interests you!

## Key Derivation Functions / Password Hashing

We've released the following [key derivation function crates](https://github.com/RustCrypto/KDFs) / [password hashing function crates](https://github.com/RustCrypto/password-hashing):

- [`hkdf`](https://github.com/RustCrypto/KDFs/tree/master/hkdf) v0.9.0-alpha.0 (final release should be out soon!)
- [`pbkdf2` v0.4.0](https://github.com/RustCrypto/password-hashing/pull/36)
- [`scrypt` v0.3.0](https://github.com/RustCrypto/password-hashing/pull/37)

## Message Authentication Codes (MACs)

We've released new versions of the following [MAC crates](https://github.com/RustCrypto/MACs):

- [`cmac` v0.3.0](https://github.com/RustCrypto/MACs/pull/41)
- [`daa` v0.2.0](https://github.com/RustCrypto/MACs/pull/44)
- [`hmac` v0.8.0](https://github.com/RustCrypto/MACs/pull/46)
- [`pmac` v0.3.0](https://github.com/RustCrypto/MACs/pull/42)

## RSA

We've just released [`rsa`](https://github.com/RustCrypto/RSA) v0.3 featuring encryption using OAEP or the legacy PKCS#1v1.5 padding, signing/verifying using RSASSA-PSS, as well as updating the dependencies (e.g. `sha2` and `digest`) to the latest versions.

## Stream Ciphers

We've released new versions of the following [stream cipher crates](https://github.com/RustCrypto/stream-ciphers):

- [`aes-ctr`](https://github.com/RustCrypto/stream-ciphers/tree/master/aes-ctr) v0.4.0 (AES-CTR with optional AESNI acceleration)
- [`cfb-mode`](https://github.com/RustCrypto/stream-ciphers/tree/master/cfb-mode) v0.4.0 (generic CFB mode support)
- [`cfb8` v0.4.0](https://github.com/RustCrypto/stream-ciphers/tree/master/cfb8) (generic CFB8 mode support)
- [`chacha20` v0.4.1](https://github.com/RustCrypto/stream-ciphers/pull/139)
- [`ctr` v0.4.0](https://github.com/RustCrypto/stream-ciphers/pull/140) (generic CTR mode support)
- [`hc-256`](https://github.com/RustCrypto/stream-ciphers/tree/master/hc-256) v0.1.0
- [`ofb` v0.2.0](https://github.com/RustCrypto/stream-ciphers/tree/master/ofb) (generic OFB mode support)
- [`salsa20` v0.5.0](https://github.com/RustCrypto/stream-ciphers/pull/144)

## Universal Hashes

We've released new versions of the following [universal hash function crates](https://github.com/RustCrypto/universal-hashes):

- [`ghash` v0.3.0](https://github.com/RustCrypto/universal-hashes/pull/60)
- [`polyval` v0.4.0](https://github.com/RustCrypto/universal-hashes/pull/59)
- [`poly1305` v0.6.0](https://github.com/RustCrypto/universal-hashes/pull/58)
## [5][Announcement: Rust-PSP](https://www.reddit.com/r/rust/comments/h0izqh/announcement_rustpsp/)
- url: https://www.reddit.com/r/rust/comments/h0izqh/announcement_rustpsp/
---
There has been an immense interest in running Rust on the Sony PSP^(/s,) so I am happy to announce the rust-psp project!

Links: [GitHub](https://github.com/overdrivenpotato/rust-psp), [crates.io](https://crates.io/crates/psp), [docs.rs](https://docs.rs/psp/)

This project is a port and improvement over the unofficial C/C++ PSPSDK from 2005. It **does not** require a custom GCC toolchain to be installed. Only Rust nightly and a cargo subcommand.

## cargo-psp and the psp crate

The `psp` crate provides a `psp::sys` submodule that houses the entire Sony PSP API. This is slightly different from most platforms that have their interfaces merged into the `libc` crate. The PSP unfortunately uses non-standard dynamic linking, (and some libraries are statically linked!), so function definitions marked `extern` are not enough.

Rather than patching LLVM or rustc, we have also merged a `mipsel-sony-psp` target upstream, and published `cargo-psp`. This is a subcommand that works exactly like `cargo build`, except it also builds your crate into a PSP-specific executable format called PRX and EBOOT.PBP.

We will be reaching full parity with the original PSPSDK very soon, and the most important libraries are already implemented. Aside from library imports, we also have support for PSP-specific custom assembly instructions via our `vfpu_asm!` macro, with no need for a custom compiler toolchain.

## Next Steps

Our next major milestone is `std` support. If you are interested in helping out, please feel free to join our `#rust-psp` channel in the PSP Homebrew discord server: [https://discord.gg/WY8XhDG](https://discord.gg/WY8XhDG)
## [6][Any projects like scala cats?](https://www.reddit.com/r/rust/comments/h0yc2l/any_projects_like_scala_cats/)
- url: https://www.reddit.com/r/rust/comments/h0yc2l/any_projects_like_scala_cats/
---
I quite like this category theory library in Scala:

https://typelevel.org/cats

I was wondering if there was any similar projects in Rust. I know this is probably a bit of a longshot!
## [7][postgrest-rs: an ORM-like library for PostgREST](https://www.reddit.com/r/rust/comments/h0y2q9/postgrestrs_an_ormlike_library_for_postgrest/)
- url: https://github.com/supabase/postgrest-rs
---

## [8][rarebears: My attempt at something like pandas for rust](https://www.reddit.com/r/rust/comments/h0nwiw/rarebears_my_attempt_at_something_like_pandas_for/)
- url: https://github.com/drew-wallace/rarebears
---

## [9][Announcing rust-search-extension v0.10! Search docs, crates, error codes, and attributes, etc in your address bar instantly!](https://www.reddit.com/r/rust/comments/h0yqcl/announcing_rustsearchextension_v010_search_docs/)
- url: https://www.reddit.com/r/rust/comments/h0yqcl/announcing_rustsearchextension_v010_search_docs/
---
Hi all! I'm so happy to announce that **Rust Search Extension** has released a new version of 0.10, and I got sponsorship from [vercel.com](https://vercel.com/?utm_source=rust-search-extension).

&gt;Rust Search Extension's official website host on `now.sh`, which is one of the fascinating products belong to  [vercel.com](https://vercel.com/?utm_source=rust-search-extension). They give me a $20/mo free plan. Woww! 

https://preview.redd.it/f62866x0y9451.png?width=1382&amp;format=png&amp;auto=webp&amp;s=181e8e773e75ea6261f8dbc8a93fff96c82201bf

Version 0.10 brings some great features, and here is the detail:

# ~ (tilde) prefix to search the external crate docs

`~` is another pretty prefix to search the external crate docs. Unlike the `@` prefix, `~` gives you the power to search docs across all your favorite crates effortlessly.

For example, input `~spawn` , you can get results related to "`spawn`" from all your favorite crates (e.g. `tokio`, `async_std`, `smol`, and `futures` ).

https://preview.redd.it/6f377pfwy9451.png?width=1902&amp;format=png&amp;auto=webp&amp;s=23a9f96d1591567a3485f9b3288ca36e8fc1a661

# Searching statistics

Searching statistics page provides a set of useful charts to visualize your search history. This gives you a powerful insight view on how frequency you search, how much valuable time you saved.

These charts including:

* A calendar heatmap represents your searching history in the last year
* A percentage chart reflects each category weight you searched
* Weekly, daily, and hourly bar chart tells you searching frequency in different time dimensionality
* Top searched crates chart gives you the rank of personal crate searching

https://preview.redd.it/fojiu0t3z9451.png?width=2324&amp;format=png&amp;auto=webp&amp;s=2fec4d7eaad56c6998ad5a57aa653cbb3b600d57

https://preview.redd.it/vh8b6sy4z9451.png?width=2094&amp;format=png&amp;auto=webp&amp;s=b52f40094eb39d948479d919ff96afb102c0324f

## Open repository quickly

You can prefix `!!!` (triple !) to open crate's repository quickly. For example, input `!!!tokio`  
 then enter, the extension will obtain the repository url then redirect to the github repository page of `tokio` effortlessly.

https://preview.redd.it/7pfbpyr8z9451.png?width=1876&amp;format=png&amp;auto=webp&amp;s=359fe880be6e1acfad629ab5e4af1fd880d472c9

https://preview.redd.it/6xcv9t0az9451.png?width=1802&amp;format=png&amp;auto=webp&amp;s=bc6e2e208fd687cabe5dacfc38647445448eccee

## Customize crates searching platform

You can customize your preferred crates searching platform such as [crates.io](https://crates.io/)(default) or [lib.rs](https://lib.rs/) on the popup page.

# Miscellaneous

* Improve searching speed, approximately 10% faster than the previous version
* New commands:
   * `:tool` - search rust tools
* Bugfix:
   * Fix std doc offline mode bug
   * Fix wrong Add button state in docs.rs's src page issue

# Installation

\- [Chrome Webstore](https://chrome.google.com/webstore/detail/rust-search-extension/ennpfpdlaclocpomkiablnmbppdnlhoh)

\- [Firefox](https://addons.mozilla.org/firefox/addon/rust-search-extension/)

\- [Microsoft Edge](https://microsoftedge.microsoft.com/addons/detail/olemfibpaicdoooacpfffccidjjagmoe)

# Some links

\- Github: [https://github.com/huhu/rust-search-extension](https://github.com/huhu/rust-search-extension)

\- Website: [https://rust-search-extension.now.sh/](https://rust-search-extension.now.sh/)

&amp;#x200B;

Welcome to give it a try. If you found a bug or something missed, please don't hesitate to let me know. Of course, any issue or PR is welcome too! Thanks! :)
## [10][Smooshmonkey, new Rust js parser for Firefox](https://www.reddit.com/r/rust/comments/h0ddpi/smooshmonkey_new_rust_js_parser_for_firefox/)
- url: https://bugzilla.mozilla.org/show_bug.cgi?id=smooshmonkey
---

## [11][[Help] gtk-rs TextBuffer connect_changed and set_text feedback loop.](https://www.reddit.com/r/rust/comments/h0x0z6/help_gtkrs_textbuffer_connect_changed_and_set/)
- url: https://www.reddit.com/r/rust/comments/h0x0z6/help_gtkrs_textbuffer_connect_changed_and_set/
---
When I programatically set the text of TextBuffer of a TextView via `set_text` it would trigger the `connect_changed` event. Is there a way to programmatically set the text of a TextBuffer without triggering the `connect_change` event? Also, I would want the `connect_changed` event to only listen to user interaction to the textview. Is there a way to configure this to only listen to user interaction?
## [12][A dead simple ByteByteJump in Rust!](https://www.reddit.com/r/rust/comments/h0x2nd/a_dead_simple_bytebytejump_in_rust/)
- url: https://www.reddit.com/r/rust/comments/h0x2nd/a_dead_simple_bytebytejump_in_rust/
---
https://github.com/aachh/bytebytejump

I made this out of interest for OISCs and Rust. I'll probably try and make some more VMs with it cause this was a (5-minute) blast! :D
