# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (42/2020)!](https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 360](https://www.reddit.com/r/rust/comments/jbh0ci/this_week_in_rust_360/)
- url: https://this-week-in-rust.org/blog/2020/10/14/this-week-in-rust-360/
---

## [3][Valve's Proton 5.13, now uses Rust!](https://www.reddit.com/r/rust/comments/jbzorm/valves_proton_513_now_uses_rust/)
- url: https://www.reddit.com/r/rust/comments/jbzorm/valves_proton_513_now_uses_rust/
---
A few hours ago [Proton 5.13-1](https://github.com/ValveSoftware/Proton/releases/tag/proton-5.13-1b) was released to github and the steam client, and with it, among exciting new supported games, the most exciting of all: new build requirements!

&gt; The build system now depends on being able to build Rust code.

Proton is now using Rust, in production, as part of [media-converter](https://github.com/ValveSoftware/Proton/tree/proton_5.13/media-converter)!

Seeing increased Rust usage like this, in notable projects like Proton, is pretty exciting!
## [4][Announcing Tokio 0.3 and the path to 1.0](https://www.reddit.com/r/rust/comments/jbr4m5/announcing_tokio_03_and_the_path_to_10/)
- url: https://tokio.rs/blog/2020-10-tokio-0-3
---

## [5][HPACK: Huffman encoder explained](https://www.reddit.com/r/rust/comments/jc7zfr/hpack_huffman_encoder_explained/)
- url: https://kristijansedlak.medium.com/hpack-huffman-encoder-explained-61102edd6ecc
---

## [6][Question: are we pdf ready?](https://www.reddit.com/r/rust/comments/jc530c/question_are_we_pdf_ready/)
- url: https://www.reddit.com/r/rust/comments/jc530c/question_are_we_pdf_ready/
---
Are we pdf ready? What pdf crates are you using? For reading, parsing, rendering pdfs? And writing as well?
## [7][Announcing ICU4X - Modular Internationalization Library in Rust!](https://www.reddit.com/r/rust/comments/jbv3ix/announcing_icu4x_modular_internationalization/)
- url: https://github.com/unicode-org/icu4x/wiki/ICU4X-Project-Announcement
---

## [8][A pitfall of Rust’s move/copy/drop semantics and zeroing data](https://www.reddit.com/r/rust/comments/jc9fkc/a_pitfall_of_rusts_movecopydrop_semantics_and/)
- url: https://benma.github.io/2020/10/16/rust-zeroize-move.html
---

## [9][Announcing async-hsm 0.2](https://www.reddit.com/r/rust/comments/jc0c0v/announcing_asynchsm_02/)
- url: https://www.reddit.com/r/rust/comments/jc0c0v/announcing_asynchsm_02/
---
 [Async-HSM](https://crates.io/crates/async-hsm) is an asynchronous hierarchical state machine (HSM). 

States are represented by asynchronous functions and state transitions are sequences of asynchronous function-calls.
## [10][Best way to update crates that come with a derive macro?](https://www.reddit.com/r/rust/comments/jc7ex5/best_way_to_update_crates_that_come_with_a_derive/)
- url: https://www.reddit.com/r/rust/comments/jc7ex5/best_way_to_update_crates_that_come_with_a_derive/
---
The github repo for my new library contains 2 crates.   
Crate A: The main library, it has a dependency on crate B as it reexports it

Crate B: The derive macro.  


Should I always keep their version the same? Even if one crate did not have code changes? Or should I allow their version to diverge?

On one hand, I like their versions to be the same. As it makes it really obvious which versions of the crate belong with each other.  However, especially with this crate I don't expect many updates to the derive macro (It implements a very small trait, without much logic) so always updating the version seems like a waste.

Allowing their version to diverge means I don't have crate releases without code updates, however I can easily see it happen that crate A is at version 0.9 while B is still 0.1, which (if you ask me) looks stupid. I also don't know what for version bump I need to give crate B if it gets a new Derive macro that is only usable with an update to crate A. Is it a breaking change so that I know for sure that updating crate A gives the updated crate B? Or shouldn't I increase the version in that way because its not actually a breaking change?

&amp;#x200B;

Also an extra question, how do you keep dependencies up to date for libraries? For programs I normally rely on dependabot but that only works if there is a Cargo.lock file.
## [11][genpdf, a user-friendly PDF generator written in pure Rust](https://www.reddit.com/r/rust/comments/jbk94w/genpdf_a_userfriendly_pdf_generator_written_in/)
- url: https://www.reddit.com/r/rust/comments/jbk94w/genpdf_a_userfriendly_pdf_generator_written_in/
---
genpdf ([crates.io](https://crates.io/crates/genpdf), [lib.rs](https://lib.rs/crates/genpdf), [docs.rs](https://docs.rs/genpdf), [source code](https://git.sr.ht/~ireas/genpdf-rs)) is a high-level PDF generator built on top of [`printpdf`](https://lib.rs/crates/printpdf) and [`rusttype`](https://lib.rs/crates/rusttype).  It takes care of the page layout and text alignment and renders a document tree into a PDF document.  All of its dependencies are written in Rust, so you don’t need any pre-installed libraries or tools.

    // Create a document and set the default font family
    let mut doc = genpdf::Document::new("./fonts", "Liberation")
        .expect("Failed to create PDF document");
    // Change the default settings
    doc.set_margins(10);
    doc.set_title("Demo document");
    // Add one or more elements
    doc.push(genpdf::elements::Paragraph::new("This is a demo document."));
    // Render the document and write it to a file
    doc.render_to_file("output.pdf")
        .expect("Failed to write PDF file");

For a complete example with all supported elements, see the
[`examples/demo.rs`][] file that generates [this PDF document][].

[`examples/demo.rs`]: https://git.sr.ht/~ireas/genpdf-rs/tree/master/examples/demo.rs
[this PDF document]: https://git.sr.ht/~ireas/genpdf-rs/blob/master/examples/demo.pdf

For more information, see the [readme](https://sr.ht/~ireas/genpdf-rs) and the  [API documentation](https://docs.rs/genpdf).
## [12][Dota2 bindings for the official webapi in Rust](https://www.reddit.com/r/rust/comments/jbs7ad/dota2_bindings_for_the_official_webapi_in_rust/)
- url: https://github.com/sn99/dota2_webapi_bindings
---

