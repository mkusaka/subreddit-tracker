# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (39/2020)!](https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/iwxitt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 357](https://www.reddit.com/r/rust/comments/iyjzb8/this_week_in_rust_357/)
- url: https://this-week-in-rust.org/blog/2020/09/23/this-week-in-rust-357/
---

## [3][I want to donate to Rust, not Mozilla. When will it be possible?](https://www.reddit.com/r/rust/comments/iyttcn/i_want_to_donate_to_rust_not_mozilla_when_will_it/)
- url: https://www.reddit.com/r/rust/comments/iyttcn/i_want_to_donate_to_rust_not_mozilla_when_will_it/
---
I heard about Rust becoming an association. Then we could donate directly to them. I like Mozilla a lot, but how they manage their ressources is scaring me.

Is there a date when Rust is becoming an association?
## [4][Building even faster interpreters in Rust](https://www.reddit.com/r/rust/comments/iyvk73/building_even_faster_interpreters_in_rust/)
- url: https://blog.cloudflare.com/building-even-faster-interpreters-in-rust/
---

## [5][CPPCON will have a talk about bridging the gap between Rust and C++!](https://www.reddit.com/r/rust/comments/iyo6tu/cppcon_will_have_a_talk_about_bridging_the_gap/)
- url: https://www.youtube.com/watch?v=_pQGRr4P16w
---

## [6][[ANN] cargo-lock v5.0: self-contained Cargo.lock parser library with CLI (list dependencies, show dependency trees, translate lockfiles)](https://www.reddit.com/r/rust/comments/iyhme0/ann_cargolock_v50_selfcontained_cargolock_parser/)
- url: https://www.reddit.com/r/rust/comments/iyhme0/ann_cargolock_v50_selfcontained_cargolock_parser/
---
Announcing v5.0.0 of the `cargo-lock` crate, a self-contained Cargo.lock parser/serializer with full support for the V1 and V2 (merge friendly) lockfile formats based on `serde`:

- GitHub: https://github.com/RustSec/cargo-lock
- Docs.rs: https://docs.rs/cargo-lock/
- Crates.io: https://crates.io/crates/cargo-lock

The `cargo-lock` crate is entirely self-contained, relying on the `Cargo.lock` file alone as opposed to requiring an entire Cargo project/workspace. This is particularly useful for embedding `Cargo.lock` information in compiled binaries, such as in [rust-audit](https://github.com/Shnatsel/rust-audit). This release adds WASM support, in case you'd like to consume/visualize Cargo.lock in a browser!

It's primarily intended for use as a library, and is used by the following projects in that capacity:

- [`rustsec` crate](https://github.com/RustSec/rustsec-crate): client library for the [RustSec Security Advisory Database](https://rustsec.org), used by the following clients:
  - [cargo-audit](https://github.com/RustSec/cargo-audit): audit Cargo.lock files for dependencies with security vulnerabilities
  - [cargo-deny](https://github.com/EmbarkStudios/cargo-deny): dependency linter which provides `cargo-audit`-like functionality with additional features
- [built](https://crates.io/crates/built): include metadata about builds in built artifacts
- [cargo-raze](https://crates.io/crates/cargo-raze): Bazel BUILD file generator for external dependencies of monorepos.

That said, in addition to functioning as a library, it also provides an installable CLI!

# Command Line Interface

Do you find yourself doing `less Cargo.lock`, `cat Cargo.lock`, or `grep Cargo.lock` often to check your dependency information? Check out the `cli` feature of `cargo-lock`:

    $ cargo install cargo-lock --features cli

This release includes a number of improvements to the CLI. For one, you can now list your dependencies by running `cargo lock`, which outputs them in a YAML-like format:

    $ cargo lock
    - autocfg 1.0.0
    - cargo-lock 5.0.0
    - fixedbitset 0.2.0
    - gumdrop 0.8.0
    - gumdrop_derive 0.8.0
    - idna 0.2.0
    - indexmap 1.3.2
    [...]

You can also include information about transitive dependencies in the output by adding the `-d` flag:

    $ cargo lock -d
    - autocfg 1.0.0
    - cargo-lock 5.0.0
      - gumdrop 0.8.0
      - petgraph 0.5.1
      - semver 0.10.0
      - serde 1.0.116
      - toml 0.5.6
      - url 2.1.1
    - fixedbitset 0.2.0
    - gumdrop 0.8.0
      - gumdrop_derive 0.8.0
    - gumdrop_derive 0.8.0
      - proc-macro2 1.0.21
      - quote 1.0.3
      - syn 1.0.40
    [...]

Want information for a single dependency? Use the `-p` option (which is nice to combine with `-d`):

    $ cargo lock -p url -d
    - url 2.1.1
      - idna 0.2.0
      - matches 0.1.8
      - percent-encoding 2.1.0

Want full source information for each crate? Use the `-s` option:

    $ cargo lock -s
    - autocfg 1.0.0 (registry+https://github.com/rust-lang/crates.io-index)
    - cargo-lock 5.0.0
    - fixedbitset 0.2.0 (registry+https://github.com/rust-lang/crates.io-index)
    - gumdrop 0.8.0 (registry+https://github.com/rust-lang/crates.io-index)
    - gumdrop_derive 0.8.0 (registry+https://github.com/rust-lang/crates.io-index)
    [...]

For fans of `cargo tree`, there's an equivalent `cargo lock tree` which can print similar-looking dependency trees using data from Cargo.lock alone:

    $ cargo lock tree
    cargo-lock 5.0.0
    ├── url 2.1.1
    │   ├── percent-encoding 2.1.0
    │   ├── matches 0.1.8
    │   └── idna 0.2.0
    │       ├── unicode-normalization 0.1.12
    │       │   └── smallvec 1.2.0
    │       ├── unicode-bidi 0.3.4
    │       │   └── matches 0.1.8
    │       └── matches 0.1.8
    ├── toml 0.5.6
    [...]

Finally, the `cargo lock translate` command provides bidirectional translation between the V1 and V2 (a.k.a. "merge friendly") Cargo.lock formats. This also serves as a demonstration of the crate's Cargo.lock serialization support, which should now produce identical output to Cargo itself for either the V1 or V2 formats.

Enjoy!
## [7][I just finished my first personal Rust project: an interpreter for a programming language I created made of buzzwords](https://www.reddit.com/r/rust/comments/iyko2l/i_just_finished_my_first_personal_rust_project_an/)
- url: https://github.com/rotoclone/strategic-communication
---

## [8][Thursday is the last day to take the Official 2020 State of Rust Survey! Even if you're new to Rust, even if you've never used Rust, even if you've barely heard of Rust, we want to hear from you!](https://www.reddit.com/r/rust/comments/iyf101/thursday_is_the_last_day_to_take_the_official/)
- url: https://blog.rust-lang.org/2020/09/10/survey-launch.html
---

## [9][type-vec: A type-safe vector with type-level length done by TYP language](https://www.reddit.com/r/rust/comments/iyu6ey/typevec_a_typesafe_vector_with_typelevel_length/)
- url: https://www.reddit.com/r/rust/comments/iyu6ey/typevec_a_typesafe_vector_with_typelevel_length/
---
type-vec: A type-safe vector with type-level length done by TYP language

Hey rustaceans,

I'm glad to announce the first release of [TYP](https://github.com/jerry73204/typ) programming language, and [type-vec](https://github.com/jerry73204/rust-type-vec), a practical application based TYP.

TYP is a type-level programming embedded in Rust. It allows you to *compute* types in an expressive language. type-vec demonstrates the construction of the vector with a type-checked length in compile time. It unleashes the true power of Rust's zero abstraction and type safety features. It is made possible by TYP language.

Since I built [type-freak](https://github.com/jerry73204/rust-type-freak) and learned the power of the type system, I wondered an language that makes type-level programming more approachable. Thanks to willcrichton's [Tyrade](https://github.com/willcrichton/tyrade). It proves the idea can be real. TYP is a direct improvement of Tyrade, and is designed to be in used practice.

Those interested can visit the repo and read the book. Your feedback are welcome!

* TYP repo: [https://github.com/jerry73204/typ/](https://github.com/jerry73204/typ/)
* TYP book: [https://github.com/jerry73204/typ-book/](https://github.com/jerry73204/typ-book/)
* type-vec repo: [https://github.com/jerry73204/rust-type-vec/](https://github.com/jerry73204/rust-type-vec/)
## [10][Haskell's Children](https://www.reddit.com/r/rust/comments/iyawej/haskells_children/)
- url: https://owenlynch.org/posts/2020-09-16-haskells-children/
---

## [11][TiKV, a CNCF graduated project, is calling for CommunityBridge mentees!](https://www.reddit.com/r/rust/comments/iyr1gs/tikv_a_cncf_graduated_project_is_calling_for/)
- url: https://www.reddit.com/r/rust/comments/iyr1gs/tikv_a_cncf_graduated_project_is_calling_for/
---
**Hey Rustaceans!** TiKV, a CNCF graduated project, is calling for CommunityBridge mentees! (recommended skills include Rust!)  [https://twitter.com/tikvproject/status/1308706660349693952](https://twitter.com/tikvproject/status/1308706660349693952)
## [12][OMG WTF RS - Resources to help you get started with Rust](https://www.reddit.com/r/rust/comments/iy81jn/omg_wtf_rs_resources_to_help_you_get_started_with/)
- url: https://ferrous-systems.com/blog/omg-wtf-rs-resources-to-help-you-get-started-with-rust/
---

