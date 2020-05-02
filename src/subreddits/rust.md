# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (18/2020)!](https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g4nu6/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 336](https://www.reddit.com/r/rust/comments/gae1nt/this_week_in_rust_336/)
- url: https://www.reddit.com/r/rust/comments/gae1nt/this_week_in_rust_336/
---
[https://this-week-in-rust.org/blog/2020/04/29/this-week-in-rust-336/](https://this-week-in-rust.org/blog/2020/04/29/this-week-in-rust-336/)
## [3][Ideas from other languages that influenced Rust](https://www.reddit.com/r/rust/comments/gc4jfd/ideas_from_other_languages_that_influenced_rust/)
- url: https://doc.rust-lang.org/stable/reference/influences.html
---

## [4][Introducing, c-closures! A general purpose way to store and call Rust closures in C/C++](https://www.reddit.com/r/rust/comments/gc2gp3/introducing_cclosures_a_general_purpose_way_to/)
- url: https://crates.io/crates/c-closures
---

## [5][Should I use Rocket or Actix?](https://www.reddit.com/r/rust/comments/gc3qf7/should_i_use_rocket_or_actix/)
- url: https://www.reddit.com/r/rust/comments/gc3qf7/should_i_use_rocket_or_actix/
---
I want to build a web API in Rust. This is more of a hobby project for me and my discord server, so I do not excpect more than maybe 5 requests per second peak (I know both can handle more than that, but the point is not performance). The point is that I don't want to spend hours looking through old posts on stackoverflow or documentation to figure out how to handle a simple request.

Which of the two is easier to write in? Or is there another, even better solution for me?
## [6][Async interviews: my take thus far](https://www.reddit.com/r/rust/comments/gbpmqf/async_interviews_my_take_thus_far/)
- url: http://smallcultfollowing.com/babysteps/blog/2020/04/30/async-interviews-my-take-thus-far/
---

## [7][minisudo, a one-file privilege escalator](https://www.reddit.com/r/rust/comments/gc4hhb/minisudo_a_onefile_privilege_escalator/)
- url: https://github.com/ogham/minisudo
---

## [8][Static polymorphism, non-virtual interface, stateless types and Rust](https://www.reddit.com/r/rust/comments/gc4vfc/static_polymorphism_nonvirtual_interface/)
- url: https://robinmoussu.gitlab.io/blog/post/rust_static_polymorphism/
---

## [9][Meuse (a free private crate registry) 0.5.0 released](https://www.reddit.com/r/rust/comments/gc5ico/meuse_a_free_private_crate_registry_050_released/)
- url: https://mcorbin.fr/posts/2020-05-02-meuse-0.5.0/
---

## [10][Rust Stream: Understanding Rust Strings [video]](https://www.reddit.com/r/rust/comments/gc3z6a/rust_stream_understanding_rust_strings_video/)
- url: https://youtu.be/7I11degAElQ
---

## [11][Clone into closure syntax?](https://www.reddit.com/r/rust/comments/gbzvh4/clone_into_closure_syntax/)
- url: https://www.reddit.com/r/rust/comments/gbzvh4/clone_into_closure_syntax/
---
I don't think this exists, but maybe I'm missing something (either in the language or how I'm implementing my stuff).

I'm setting up a networked program with a GUI frontend, which means that I'm passing around a number of `Arc&lt;Mutex&lt;SomeData&gt;&gt;` clones as well as various `tokio` sync channel senders/receivers into callbacks and aysnc loops (via `tokio::spawn`).

The ugly bit I would like to avoid is needing yet another clone every time I want to use one of these bindings from within a closure. The nature of most of the UI/`tokio` callbacks requires that all bindings are moved, of course, and I'm left with (AFAICT) two options that I don't much care for:

    let network_state1 = network_state.clone();
    let network_state2 = network_state.clone();
    let network_menu = MenuTree::new()
        .leaf("Info", move |s| {
            ui::launch_info_dialog(s, network_state1.clone())
        })
        .leaf("Connect", move |s| {
            ui::launch_connect_dialog(s, network_state2.clone())
        });

and

    let mut network_menu = MenuTree::new();
    {
        let _network_state = network_state.clone();
        network_menu.add_leaf("Info", move |s| {
            // TODO: show info
        });
    }
    {
        let network_state = network_state;
        network_menu.add_leaf("Connect", move |s| {
            ui::launch_connect_dialog(s, network_state.clone())
        });
    }

Admittedly, neither of these is all that bad, but over a more complex function, that's potentially a lot of extra lines of code just to do one thing (and in the first case, I really don't like relying on "I wrote the right name" to avoid compile errors).

What I'd really like -- and maybe this exists and I just haven't found it -- is a syntax that lets me clone *into* the closure at the point of creation. Something like this:

    let clone_closure = clone network_state |_| { ... };

If I were making a proper RFC, I might consider how this works with `move` like,

    let mixed_closure = move + clone network_state |_| { ... };
    let multi_clone_closure = clone {network_state, another} |_| { ... };
    let mixed_multi_closure = move + clone {network_state, another} |_| { ... };

Am I overthinking this? Have I missed some feature of the language? Is this where I start my journey of making RFCs for the language...?
## [12][[ANN] RustCrypto: `p256` and `k256` v0.2.0: pure Rust NIST P-256 and secp256k1 curve arithmetic](https://www.reddit.com/r/rust/comments/gbjsr9/ann_rustcrypto_p256_and_k256_v020_pure_rust_nist/)
- url: https://www.reddit.com/r/rust/comments/gbjsr9/ann_rustcrypto_p256_and_k256_v020_pure_rust_nist/
---
Announcing the v0.2.0 releases of the following RustCrypto elliptic curve crates:

- `p256`: NIST P-256
  - GitHub: https://github.com/RustCrypto/elliptic-curves/tree/master/p256
  - crates.io: https://crates.io/crates/p256
  - docs.rs: https://docs.rs/p256/
- `k256`: secp256k1 (as used by Bitcoin, Ethereum, etc)
  - GitHub: https://github.com/RustCrypto/elliptic-curves/tree/master/k256
  - crates.io: https://crates.io/crates/k256
  - docs.rs: https://docs.rs/k256/

Both of these releases now implement curve/field arithmetic, namely they implement the complete Weierstrass formulas and are initially targeting correctness over performance. Because of all of that, they are suitable for environments which require small code sizes (e.g. embedded), and are designed from the ground up to work in `no_std` environments.

These are the first releases of these crates with an `arithmetic` feature. The code is brand new and has not been thoroughly reviewed, though we believe it is of high quality. Some of the field arithmetic implementations have been proptested against the ones in [fiat-rust](https://github.com/mit-plv/fiat-crypto/tree/master/fiat-rust), and we will continue to investigate ways to ensure the implementations are correct.

All of that said, USE AT YOUR OWN RISK!
