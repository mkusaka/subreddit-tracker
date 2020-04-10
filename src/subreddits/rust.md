# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (15/2020)!](https://www.reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 333](https://www.reddit.com/r/rust/comments/fxqdrs/this_week_in_rust_333/)
- url: https://this-week-in-rust.org/blog/2020/04/07/this-week-in-rust-333/
---

## [3][Mental models around Ok-Wrapping](https://www.reddit.com/r/rust/comments/fyakdx/mental_models_around_okwrapping/)
- url: https://vorner.github.io/2020/04/09/wrapping-mental-models.html
---

## [4][What is wrong with Ok(match thing { ... }) ?](https://www.reddit.com/r/rust/comments/fye583/what_is_wrong_with_okmatch_thing/)
- url: https://www.reddit.com/r/rust/comments/fye583/what_is_wrong_with_okmatch_thing/
---
Sorry for yet another post on this topic. I'll keep it short.

In boats's [recent blog](https://boats.gitlab.io/blog/post/why-ok-wrapping/), he mentions:

&gt; Most of my functions with many return paths terminate with a match statement. Technically, these could be reduced to a single return path by just wrapping the whole match in an Ok, but I don’t know anyone who considers that good form, and I certainly don’t. But an experience I find quite common is that I introduce a new arm to that match as I introduce some new state to handle, and handling that new state is occassionally fallible.

I personally do not see the problem with Ok-wrapping the match. Or, if one doesn't wish to do this, introducing a let binding:
```rust
let result = match thing {
   ...
};
Ok(result)
```

As for "expressing effects", we already have syntax for that: `return Err(...);`. The only case "Ok-wrapping" would really be a boon is with multiple `return Ok(result);` paths, which I don't find to be common in practice.

I am not against Ok-Wrapping (other than recognising that the addition has a cost), but am surprised about the number of error-handling crates which have sprung up over the years and amount of discussion this topic has generated. The only error-handling facility I find lacking in `std` rust is the overhead of instantiating a new error type ([anyhow::anyhow](https://docs.rs/anyhow/1.0.28/anyhow/macro.anyhow.html) and [thiserror](https://github.com/dtolnay/thiserror) address this omission).
## [5][Brainfuck to Web Assembly compiler in &lt; 100 lines](https://www.reddit.com/r/rust/comments/fy0tqj/brainfuck_to_web_assembly_compiler_in_100_lines/)
- url: https://github.com/richardanaya/watson/blob/master/examples/bf/src/main.rs
---

## [6][Q: Rust as a hardware description language to be synthesized on an FPGA?](https://www.reddit.com/r/rust/comments/fyd57q/q_rust_as_a_hardware_description_language_to_be/)
- url: https://www.reddit.com/r/rust/comments/fyd57q/q_rust_as_a_hardware_description_language_to_be/
---
I'm just curious, are there any projects that try to adapt rust to the domain of hardware description languages (like Verilog and VHDL) and system-level modelling (SystemC)?

At a first glance this is an interesting perspective, since it seems that Rust's approach to memory safety and value access semantics maps quite nicely to RTL and Rust's powerful macro system can be used to mimic/implement more subtle aspects of hardware design such as signal definition and mapping.

Heck, maybe even async stuff could be used to synthesize state machines!

Of course all of the above is just my pure speculation. So I'd like to know the opinion of experienced people that work in this field.
## [7][GPLv2 airbag control software written in Rust](https://www.reddit.com/r/rust/comments/fya5om/gplv2_airbag_control_software_written_in_rust/)
- url: https://github.com/victoredwardocallaghan/techair
---

## [8][64K BASIC runs the classics like Oregon Trail and Super Star Trek](https://www.reddit.com/r/rust/comments/fy7s1z/64k_basic_runs_the_classics_like_oregon_trail_and/)
- url: http://basic-lang.org/
---

## [9][Deserialize an arbitrary json structure](https://www.reddit.com/r/rust/comments/fyevnx/deserialize_an_arbitrary_json_structure/)
- url: https://www.reddit.com/r/rust/comments/fyevnx/deserialize_an_arbitrary_json_structure/
---
I have a json object like following,

    {
    page1: { group1: [name: ""] },
    page2: { group1: [name: ""] }, 
    // around 500 of them...
    page500: { group22: [name: ""] }
    }

From what I know is the

* page can be 1 to 500 or even more (totally random, all have page prefix).
* each of the page can have group1 or group 2 or group 22 (totally random, all have group prefix).
* all group have a specific name string.

Been searching high and low for a way to create a struct and process the json data.

I did a reddit search before making this post, but I cannot seem to apply the knowledge gained from those posts:

* [https://www.reddit.com/r/rust/comments/fw1q36/help\_with\_constructing\_arbitrary\_json\_object/](https://www.reddit.com/r/rust/comments/fw1q36/help_with_constructing_arbitrary_json_object/)
* [https://www.reddit.com/r/learnpython/comments/c8qlyq/parsing\_a\_very\_large\_streamed\_json\_file\_and/](https://www.reddit.com/r/learnpython/comments/c8qlyq/parsing_a_very_large_streamed_json_file_and/)

I can use serde Value to read them without any speicific structure, which was described on this [github issue](https://github.com/serde-rs/json/issues/144).

It's super easy in js/python where you don't have to specify the types, but cannot figure out on rust/nim etc languages.

I cannot change the structure of the json to an array, it comes from a third party api.

Here is an example file where I'm reading data using neon and serde. The data comes from js, and I do some loops to get the insider data. [https://github.com/entrptaher/neon-serde-example/blob/master/native/src/lib.rs#L19](https://github.com/entrptaher/neon-serde-example/blob/master/native/src/lib.rs#L19)

Is there any way I can get a struct for this kind of data dynamically? Is that even healthy to do so?

Pardon me if I made any mistake. Thank you in advance.
## [10][Library-ification and analyzing Rust](https://www.reddit.com/r/rust/comments/fxs0fm/libraryification_and_analyzing_rust/)
- url: http://smallcultfollowing.com/babysteps/blog/2020/04/09/libraryification/
---

## [11][How to download all the crates on crates.io](https://www.reddit.com/r/rust/comments/fxxued/how_to_download_all_the_crates_on_cratesio/)
- url: https://www.pietroalbini.org/blog/downloading-crates-io/
---

## [12][Help with speeding up smart-pointered code](https://www.reddit.com/r/rust/comments/fyd96g/help_with_speeding_up_smartpointered_code/)
- url: https://www.reddit.com/r/rust/comments/fyd96g/help_with_speeding_up_smartpointered_code/
---
Hi. I was trying to make a randomized skiplist with rust. I needed multiple pointers to nodes and i also needed to be able to mutate the pointers which they hold, So I went with a combination of Rc and RefCell. It worked well but then i tried a completely unsafe solution with raw pointers and it is about 2x faster. Then i benchmarked my code and it seems that most of the overhead comes from borrow() and borrow\_mut() methods of RefCell (there is also a drop\_in\_place function that takes considerable time).

Do you think my datastructure is poorly designed, or is this performance drop almost natural when using refcells? (As far as i can understand the overhead of Rc is almost zero and nothing about Rc shows up in my profiling)
