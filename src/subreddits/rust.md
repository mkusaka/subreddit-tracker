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
## [2][This Week in Rust 327](https://www.reddit.com/r/rust/comments/fabtwg/this_week_in_rust_327/)
- url: https://this-week-in-rust.org/blog/2020/02/25/this-week-in-rust-327/
---

## [3][Command line steganography for png images written in rust. (My First rust crate - Feedback and thoughts appreciated)](https://www.reddit.com/r/rust/comments/fbavos/command_line_steganography_for_png_images_written/)
- url: https://crates.io/crates/stegano-cli
---

## [4][I want off Mr. Golang's Wild Ride](https://www.reddit.com/r/rust/comments/fax7lj/i_want_off_mr_golangs_wild_ride/)
- url: https://fasterthanli.me/blog/2020/i-want-off-mr-golangs-wild-ride/
---

## [5][xworks - A dead-simple FAQ-style static site generator, à la anarchy.works](https://www.reddit.com/r/rust/comments/fb7jo8/xworks_a_deadsimple_faqstyle_static_site/)
- url: https://github.com/t1ra/xworks
---

## [6][How to implement a graph data structure in Rust?](https://www.reddit.com/r/rust/comments/fbb4rj/how_to_implement_a_graph_data_structure_in_rust/)
- url: https://www.reddit.com/r/rust/comments/fbb4rj/how_to_implement_a_graph_data_structure_in_rust/
---
A graph data structure usually require each node to has write access to its peer. I find implementing this hard with Rust without unsafe.

What's are standard ways for implementing a graph in Rust or a library that make this easy?
## [7][Complex Rust data structures exposed seamlessly to C++](https://www.reddit.com/r/rust/comments/faza6k/complex_rust_data_structures_exposed_seamlessly/)
- url: https://crisal.io/words/2020/02/28/C++-rust-ffi-patterns-1-complex-data-structures.html
---

## [8][pallet, a searchable datastore built on sled and tantivy](https://www.reddit.com/r/rust/comments/fb23tl/pallet_a_searchable_datastore_built_on_sled_and/)
- url: https://www.reddit.com/r/rust/comments/fb23tl/pallet_a_searchable_datastore_built_on_sled_and/
---
[docs.rs](https://docs.rs/pallet) | [github](https://github.com/kardeiz/pallet)

A searchable document datastore built on [`sled`](https://docs.rs/sled) and [`tantivy`](https://docs.rs/tantivy).

Provides a typed-tree interface to a `sled` database, with standard datastore ops (`find`, `create`, `update`, `delete`), but also Lucene/Elasticsearch style searching.

Usage:

    #[macro_use]
    extern crate serde;
    
    #[derive(Serialize, Deserialize, Debug, pallet::DocumentLike)]
    #[pallet(tree_name = "books")]
    pub struct Book {
        #[pallet(default_search_field)]
        title: String,
        #[pallet(default_search_field)]
        description: Option&lt;String&gt;,
        #[pallet(index_field_type = "u64")]
        rating: u8,
    }
    
    fn main() -&gt; Result&lt;(), Box&lt;dyn std::error::Error&gt;&gt; {
        let temp_dir = tempfile::TempDir::new_in(".")?;
    
        let db = sled::open(temp_dir.path().join("db"))?;
    
        let store = pallet::Store::builder()
            .with_db(db)
            .with_index_dir(temp_dir.path())
            .finish()?;
    
        let books = vec![
            Book {
                title: "The Old Man and the Sea".into(),
                description: Some(
                    "He was an old man who fished alone in a skiff in \
                the Gulf Stream and he had gone eighty-four days \
                now without taking a fish."
                        .into(),
                ),
                rating: 10,
            },
            Book {
                title: "The Great Gatsby".into(),
                description: Some("About a man and some other stuff".into()),
                rating: 8,
            },
        ];
    
        let _ = store.create_multi(&amp;books)?;
    
        let books = store.search("man AND rating:&gt;8")?;
    
        println!("{:?}", books);
    
        Ok(())
    }
## [9][High-level graphics?](https://www.reddit.com/r/rust/comments/fb6qo1/highlevel_graphics/)
- url: https://www.reddit.com/r/rust/comments/fb6qo1/highlevel_graphics/
---
Title. Does such a thing exist? Are there any high-level graphics libraries based on a Vulkan backend or something, that's rather easy to get into?
## [10][A tado° connected thermostatic devices Prometheus exporter written in Rust](https://www.reddit.com/r/rust/comments/fbayxd/a_tado_connected_thermostatic_devices_prometheus/)
- url: https://github.com/eko/tado-exporter
---

## [11][How should I go about fixing an incorrect comment in the code?](https://www.reddit.com/r/rust/comments/fbcaxj/how_should_i_go_about_fixing_an_incorrect_comment/)
- url: https://www.reddit.com/r/rust/comments/fbcaxj/how_should_i_go_about_fixing_an_incorrect_comment/
---
The recent anti-golang rant led me down a rabbit hole where I found myself looking at the source for Rust's monotonic time implementation in https://github.com/rust-lang/rust/blob/master/src/libstd/time.rs, where I found a backwards comment at line 286 documenting the `saturating_duration_since` function:

    /// Returns the amount of time elapsed from another instant to this one,
    /// or zero duration if that instant is earlier than this one.

As you would expect from the function name, first comment line, and actual code, the second line should actually say "or zero duration if that instant is _later_ than this one."  The code works as expected, but the comment is incorrect. I figured I'd submit a PR with the one-word fix, but wanted to be sure I followed the community guidelines. I read through https://github.com/rust-lang/rust/blob/master/CONTRIBUTING.md, and didn't see an answer to my questions, which are these:

1. For a simple doc change like this, is a bare PR enough, or does it need to reference an issue? 
2. If it needs an issue, does that issue count as a bug report? There's no actual behavioral bug. Is fixing a comment a feature request? 

Anyway, sorry for the dumb questions, just wasn't sure what the protocol is here.
## [12][Announcing arcs 0.3.0 - now with more algorithms and better docs](https://www.reddit.com/r/rust/comments/fatgr1/announcing_arcs_030_now_with_more_algorithms_and/)
- url: https://docs.rs/arcs/0.3.0/arcs/
---

