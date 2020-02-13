# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (7/2020)!](https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/)
- url: https://www.reddit.com/r/rust/comments/f1ucwh/hey_rustaceans_got_an_easy_question_ask_here_72020/
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

Also check out [last week's thread](https://reddit.com/r/rust/comments/ey2wte/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (7/2020)?](https://www.reddit.com/r/rust/comments/f1uej1/whats_everyone_working_on_this_week_72020/)
- url: https://www.reddit.com/r/rust/comments/f1uej1/whats_everyone_working_on_this_week_72020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-7-2020/38078?u=llogiq)!
## [3][How nnethercote optimized the heck out of LEB128 code](https://www.reddit.com/r/rust/comments/f36j05/how_nnethercote_optimized_the_heck_out_of_leb128/)
- url: https://github.com/rust-lang/rust/pull/69050#issuecomment-585508353
---

## [4][Why can't I convert Option&lt;T&gt; to Option&lt;U&gt; if U implements From&lt;T&gt;](https://www.reddit.com/r/rust/comments/f32jcv/why_cant_i_convert_optiont_to_optionu_if_u/)
- url: https://www.reddit.com/r/rust/comments/f32jcv/why_cant_i_convert_optiont_to_optionu_if_u/
---
I just came upon this because I assumed it was a thing.  

Why is it that 

    impl&lt;T,U&gt; From&lt;Option&lt;T&gt;&gt; for Option&lt;U&gt; where U: From&lt;T&gt;

is not implemented.  It seems like a fairly obvious use case to me.  Is there a particular reason they couldn't?

A similar question goes for Result
## [5][Relay Compiler being ported into Rust](https://www.reddit.com/r/rust/comments/f36rcq/relay_compiler_being_ported_into_rust/)
- url: https://www.reddit.com/r/rust/comments/f36rcq/relay_compiler_being_ported_into_rust/
---
[https://twitter.com/en\_JS/status/1227773312346075136](https://twitter.com/en_JS/status/1227773312346075136)

&gt;I’m excited to  share that we’re experimenting w porting Relay Compiler into [rustlang](https://twitter.com/hashtag/rustlang?src=hash&amp;ref_src=twsrc%5Etfw"&gt;#rustlang).  Our early results are promising: w a bit of structure it’s easy to pick  up, and we can move fast and ship high quality, efficient  code.

[https://github.com/facebook/relay/tree/master/compiler](https://github.com/facebook/relay/tree/master/compiler)
## [6][I just published crow 0.3, a pixel perfect 2D rendering engine.](https://www.reddit.com/r/rust/comments/f2wuwc/i_just_published_crow_03_a_pixel_perfect_2d/)
- url: https://www.reddit.com/r/rust/comments/f2wuwc/i_just_published_crow_03_a_pixel_perfect_2d/
---
[`crow`](https://crates.io/crates/crow) is a 2D rendering engine based on OpenGL.
I want this crate to be easy to use while still
allowing for nearly everything one might want while using pixel art.

A WIP showcase game can be found on https://github.com/lcnr/akari

Please tell me if there is anything you would want in a 2D game and don't know
how you could implement it using `crow`.
## [7][Wait on spawned child with inherited stdin](https://www.reddit.com/r/rust/comments/f38hnh/wait_on_spawned_child_with_inherited_stdin/)
- url: https://www.reddit.com/r/rust/comments/f38hnh/wait_on_spawned_child_with_inherited_stdin/
---
I've got a wrapper whose sole purpose is to run another arbitrary program, delegate the original stdio to it and wait for it to finish. All's good and easy but Child.wait() specifically closes the inherited stdin. Is there a way not to close stdin, yet wait for the child to exit? The child should exit on stdin EOF or a stdin command.
## [8][A handwired unsplitted ergo keyboard with a firmware written in Rust](https://www.reddit.com/r/rust/comments/f2o9y8/a_handwired_unsplitted_ergo_keyboard_with_a/)
- url: https://raw.githubusercontent.com/TeXitoi/keyberon-f4/master/images/keyberon-56.jpg
---

## [9][how to use reqwest without async?](https://www.reddit.com/r/rust/comments/f39ueb/how_to_use_reqwest_without_async/)
- url: https://www.reddit.com/r/rust/comments/f39ueb/how_to_use_reqwest_without_async/
---
hi,

&amp;#x200B;

is there any way to use reqwest without async, i cant simply figure out how this thing works at all.
## [10][GitHub - drbh/hotpot-db: 🌶 hottest way to store data on the web. If you like JSON, and dealing with things later.](https://www.reddit.com/r/rust/comments/f2vbji/github_drbhhotpotdb_hottest_way_to_store_data_on/)
- url: https://www.reddit.com/r/rust/comments/f2vbji/github_drbhhotpotdb_hottest_way_to_store_data_on/
---
Hey Rustaceans 👋,

I'd like to share and get feedback on a small crate called hotpot-db! It is essentially an API around SQLite's JSON extension, which allows you to store data in a NoSQL dare I say, DynamoDB like way. 

Any struct that can be Serialized to JSON can be stored and efficiently queried. 

## Flavor Palette

1. schemaless
2. reliable (uses SQLite3)
3. embeddable
4. fast (&lt;200ms to search through +500K objects)
5. JSON store
6. queryable JSON schemas 

## Full Example
```rust
use hotpot_db::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
struct Person {
    name: String,
    age: u8,
}

fn main() -&gt; Result&lt;(), hotpot_db::Error&gt; {
    let mut pot = HotPot::new();

    // lets make a new collection
    pot.create_collection("address_book")?;

    // well make a new item we want to store
    let person = Person {
        name: String::from("david holtz"),
        age: 26,
    };

    // we insert the object into the collection!
    pot.insert::&lt;Person&gt;("address_book", &amp;person)?;

    // finally we can query
    let query = QueryBuilder::new()
        .collection("address_book")
        .kind(QueryKind::Object)
        .key("name")
        .comparison("=")
        .string("david holtz")
        .finish();

    let results = pot.execute(query);
    println!("{:#?}", results);

    Ok(())
}
```

[https://github.com/drbh/hotpot-db](https://github.com/drbh/hotpot-db)

This was my first attempt at using the Builder pattern (for queries) and generics (for insert). hotpot-db has already been super helpful for me in the development process where my DB schema is in flux. hotpot-db lets me deal with schemas later and still have solid query performance.


Any feedback on the project/code and idea is greatly appreciated! 😊
## [11][Media Processing](https://www.reddit.com/r/rust/comments/f37ymd/media_processing/)
- url: https://www.reddit.com/r/rust/comments/f37ymd/media_processing/
---
I’m considering Rust for a new project, which requires processing media streams for a wide array of formats. My alternative is C or C++, using FFmpeg. 

Is there an effective/up to date Rust binding for FFmpeg or other crate that people are using for this purpose? Would be great to get a quick overview of the Rust media ecosystem for a Rust noob.
## [12][Mimalloc Rust version 0.1.14 has just been released!](https://www.reddit.com/r/rust/comments/f2vb4z/mimalloc_rust_version_0114_has_just_been_released/)
- url: https://www.reddit.com/r/rust/comments/f2vb4z/mimalloc_rust_version_0114_has_just_been_released/
---
The Rust [mimalloc wrapper](https://github.com/purpleprotocol/mimalloc_rust) crate has just had a new version released!

### Changes

* Updated [mimalloc source](https://github.com/microsoft/mimalloc) to version 1.6.0.

* Removes [libc](https://crates.io/crates/libc) from dependencies, adding support for targets targets which do not depend on it such as wasm.
