# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (33/2020)?](https://www.reddit.com/r/rust/comments/i6yras/whats_everyone_working_on_this_week_332020/)
- url: https://www.reddit.com/r/rust/comments/i6yras/whats_everyone_working_on_this_week_332020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-33-2020/47095?u=llogiq)!
## [3][Introducing Bevy: a refreshingly simple data-driven game engine and app framework built in Rust](https://www.reddit.com/r/rust/comments/i7bcwu/introducing_bevy_a_refreshingly_simple_datadriven/)
- url: https://bevyengine.org/news/introducing-bevy/
---

## [4][[ANN] RustCrypto: `k256` and `p256` v0.2.0: pure Rust secp256k1 and NIST P-256 ECDH and ECDSA (no_std/embedded-friendly)](https://www.reddit.com/r/rust/comments/i7lca3/ann_rustcrypto_k256_and_p256_v020_pure_rust/)
- url: https://www.reddit.com/r/rust/comments/i7lca3/ann_rustcrypto_k256_and_p256_v020_pure_rust/
---
Announcing v0.4.0 releases of these RustCrypto elliptic curve crates:

- `k256`: secp256k1 (as used by Bitcoin, Ethereum, etc)
  - GitHub: https://github.com/RustCrypto/elliptic-curves/tree/master/k256
  - crates.io: https://crates.io/crates/k256
  - docs.rs: https://docs.rs/k256/
- `p256`: NIST P-256 (as used by SSL/TLS, Bluetooth, etc)
  - GitHub: https://github.com/RustCrypto/elliptic-curves/tree/master/p256
  - crates.io: https://crates.io/crates/p256
  - docs.rs: https://docs.rs/p256/

(see also [`ecdsa` v0.7](https://docs.rs/ecdsa/0.7.1/ecdsa/index.html) and [`p384` v0.3](https://docs.rs/p384/0.3.0/p384/))

The major notable new features in these releases are:

# Elliptic Curve Diffie-Hellman

Key exchange protocol which establishes a shared secret between two parties.

- [`k256` docs](https://docs.rs/k256/0.4.0/k256/ecdh/index.html) (secp256k1)
- [`p256` docs](https://docs.rs/p256/0.4.0/p256/ecdh/index.html) (NIST P-256)
- [Generic implementation](https://docs.rs/elliptic-curve/0.5.0/elliptic_curve/ecdh/index.html)

# Elliptic Curve Digital Signature Algorithm

Pervasively used public-key scheme for authenticating messages.

- [`k256` docs](https://docs.rs/k256/0.4.0/k256/ecdsa/index.html) (secp256k1)
- [`p256` docs](https://docs.rs/p256/0.4.0/p256/ecdsa/index.html) (NIST P-256)
- Generic implementation:
  - [`ecdsa::Signer`](https://docs.rs/ecdsa/0.7.1/ecdsa/index.html)
  - [`ecdsa::Verifier`](https://docs.rs/ecdsa/0.7.1/ecdsa/verifier/struct.Verifier.html)

# Notes on this release

These crates contain *experimental* pure Rust implementations of scalar/field arithmetic for the respective elliptic curves (secp256k1, NIST P-256). These implementations are new, unaudited, and haven't received much public scrutiny. We have explicitly labeled them as being at a "USE AT YOUR OWN RISK" level of maturity.

That said, these implementations utilize the best modern practices for this class of elliptic curves (complete projective formulas providing constant time scalar multiplication).

In particular:

- The `k256` crate now implements [optimized "lazy" field arithmetic based on the bitcoin-core/secp256k1 C library](https://github.com/RustCrypto/elliptic-curves/pull/59) optimized for both 32-bit and 64-bit architectures.
- The `p256` crate includes [sidechannel-resistant scalar arithmetic optimized for creating signatures on embedded platforms](https://github.com/RustCrypto/elliptic-curves/pull/83) and presently provides 64-bit scalar/field arithmetic.

This release has been a cross-functional effort, with contributions from some of the best Rust elliptic curve cryptography experts. I'd like to thank everyone who's contributed, and hope that these crates are useful, especially for embedded cryptography and cryptocurrency use cases.
## [5][Embedded Rust tooling for the 21st century](https://www.reddit.com/r/rust/comments/i7pxia/embedded_rust_tooling_for_the_21st_century/)
- url: https://ferrous-systems.com/blog/21st-century-embedded-tooling/
---

## [6][Single Page Applications using Rust](https://www.reddit.com/r/rust/comments/i7k9kw/single_page_applications_using_rust/)
- url: http://www.sheshbabu.com/posts/rust-wasm-yew-single-page-application/
---

## [7][TypeTest: An offline keyboard typing test, written with Iced](https://www.reddit.com/r/rust/comments/i7grr5/typetest_an_offline_keyboard_typing_test_written/)
- url: https://i.redd.it/3q6uzbfgj9g51.gif
---

## [8][Interview with Steve Klabnik, Rust core team contributor](https://www.reddit.com/r/rust/comments/i7q1tj/interview_with_steve_klabnik_rust_core_team/)
- url: https://evrone.com/steve-klabnik-interview
---

## [9][Input from racing wheel](https://www.reddit.com/r/rust/comments/i7o41i/input_from_racing_wheel/)
- url: https://www.reddit.com/r/rust/comments/i7o41i/input_from_racing_wheel/
---
Hello!I'm working on a racing game, and I'd like to be able to use a wheel (Like the logitech g29, and other wheels) for input and force feedback. The game engine I'm using (Godot) doesn't support wheels, but it can work with rust.Is there any way I could use rust to interface with the wheel?

I already did some searching but I can't find anything on the subject matter. I already know how to make godot work with rust.  
Edit: I'm on windows, but cross platform (Windows and linux) is preferred
## [10][Meili raises 1.5M€ for open source search in Rust](https://www.reddit.com/r/rust/comments/i71w2z/meili_raises_15m_for_open_source_search_in_rust/)
- url: https://blog.meilisearch.com/meili-fundraise/
---

## [11][Interactive Chord Diagrams with Rust](https://www.reddit.com/r/rust/comments/i7pqd0/interactive_chord_diagrams_with_rust/)
- url: https://shahinrostami.com/posts/programming/rust-notebooks/chord-diagrams/
---

## [12][IntelliJ Rust Changelog #128](https://www.reddit.com/r/rust/comments/i770v7/intellij_rust_changelog_128/)
- url: https://intellij-rust.github.io/2020/08/10/changelog-128.html
---

