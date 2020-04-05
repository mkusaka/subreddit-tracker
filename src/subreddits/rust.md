# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (14/2020)!](https://www.reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/frfduy/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fnfky9/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 332](https://www.reddit.com/r/rust/comments/ftl9l3/this_week_in_rust_332/)
- url: https://this-week-in-rust.org/blog/2020/03/31/this-week-in-rust-332/
---

## [3][State machines with const-generics](https://www.reddit.com/r/rust/comments/fvciq3/state_machines_with_constgenerics/)
- url: https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=e957af6f8ba26b8f496366182178a9d2
---

## [4][To learn how base64 works I implemented it with Rust](https://www.reddit.com/r/rust/comments/fvcouv/to_learn_how_base64_works_i_implemented_it_with/)
- url: https://www.reddit.com/r/rust/comments/fvcouv/to_learn_how_base64_works_i_implemented_it_with/
---
[https://medium.com/@niklasbuechner/implementing-base64-in-rust-34ef6db1e73a](https://medium.com/@niklasbuechner/implementing-base64-in-rust-34ef6db1e73a)
## [5][Dirscan: A high performance tool for summarising large directories or drives](https://www.reddit.com/r/rust/comments/fv2czk/dirscan_a_high_performance_tool_for_summarising/)
- url: https://github.com/orf/dirscan
---

## [6][Panic message conventions?](https://www.reddit.com/r/rust/comments/fv7830/panic_message_conventions/)
- url: https://www.reddit.com/r/rust/comments/fv7830/panic_message_conventions/
---
I'm writing libraries and was wondering what the panic message convention is.

I looked around but there's no posts about this, so I looked at the rust compiler source, here are my findings:

&amp;#x200B;

* Lots of panics are lower-case without final punctuation: **("overflow when converting float to duration")**
* Panics in tests tend to have an uppercase start: **("Unexpected value {:?}")**
* Bootstrap panics include newlines: **("\\n\\nfailed to specify \`dist.sign-folder\` in \`config.toml\`\\n\\n")**
* liballoc tests contain exclamation marks at the end **("usize::MAX should trigger an overflow!")**
* Some panics include their logical domain of interest designated before a colon: **("align\_offset: align is not a power-of-two")**
* Some panics put printed data between backticks: **("invalid format arg \`{:?}\`")**
* Some panics put argument data after the final colon: **("rustc didn't succeed: {}")**

&amp;#x200B;

**My question is:** what is generally the "best" format for panic messages?
## [7][Rust Dataframe: Update 1](https://www.reddit.com/r/rust/comments/fv11ir/rust_dataframe_update_1/)
- url: https://github.com/nevi-me/rust-dataframe/blob/master/notes/update-01__04-04-2020.md
---

## [8][Help sorting BTreeMap in different order](https://www.reddit.com/r/rust/comments/fv96yi/help_sorting_btreemap_in_different_order/)
- url: https://www.reddit.com/r/rust/comments/fv96yi/help_sorting_btreemap_in_different_order/
---
I need to have two of BTreeMaps&lt;i64, Data&gt;, one sorted ascending and the other sorted descending (different data in each). Performance is very, very important. I'm already paying up in the Ordering enum (I cannot just subtract like I would in C++), so I'm trying not to pay up more with extra Options and stuff like I would by having to implement PartialOrd.

My first attempt was to create `struct IncVal(i64)` and `struct DecVal(i64)` but that was a lot of biolerplate and had to go though the Option with PartialEq. There has to be a better way.

This is a somewhat simplified example. There key is a user defined class that holds a more complex number, but just getting it to work with an i128 right now would be a small win.
## [9][How does one mutably borrow a value after conditionally returning an immutable reference to it?](https://www.reddit.com/r/rust/comments/fv9yke/how_does_one_mutably_borrow_a_value_after/)
- url: https://www.reddit.com/r/rust/comments/fv9yke/how_does_one_mutably_borrow_a_value_after/
---
I have the following pattern in my code:

```
struct Data {};

fn x(input: &amp;mut Vec&lt;Data&gt;) -&gt; Option&lt;&amp;Data&gt; {
    let (idx, data) = input.iter().enumerate().next()?;
    if false { // some condition
        Some(data)
    } else {
        input.remove(idx);
        x(input)
    }
}
```

I get the following error from the borrow checker:

```
cannot borrow `*input` as mutable because it is also borrowed as immutable
```

How does one make this work?

EDIT: I've simplified the problem to the following:

```rust
struct Data {}

fn condition(data: &amp;Data) -&gt; bool {
    true
}

fn x(input: &amp;mut Vec&lt;Data&gt;) -&gt; &amp;Data {
    let data = &amp;input[0];
    if condition(data) {
        data
    } else {
        y(input)
    }
}

fn y(input: &amp;mut Vec&lt;Data&gt;) -&gt; &amp;Data {
    todo!()
}
```
## [10][Error recovery with parser combinators (using nom)](https://www.reddit.com/r/rust/comments/fvc7hy/error_recovery_with_parser_combinators_using_nom/)
- url: https://www.eyalkalderon.com/nom-error-recovery/
---

## [11][[Question] What are some of real-world open-source Rust product?](https://www.reddit.com/r/rust/comments/fvdr7y/question_what_are_some_of_realworld_opensource/)
- url: https://www.reddit.com/r/rust/comments/fvdr7y/question_what_are_some_of_realworld_opensource/
---
Hey!

It seems that a lot of companies started adopting [Rust](https://www.rust-lang.org/production) , and more traditional companies [do as well](https://twitter.com/steveklabnik/status/1246510586416050179). However, I would like to check some products or startup's work built on top of Rust Stack and how it serves Mobile app or large-scale website as a backend.

For learning purposes, I would like to check real-world open-source Rust stack, and how teams orchestrate different services. Do you know of any repositories like that?
## [12][vim - How can I see source code for a module/trait I have imported?](https://www.reddit.com/r/rust/comments/fvb8or/vim_how_can_i_see_source_code_for_a_moduletrait_i/)
- url: https://www.reddit.com/r/rust/comments/fvb8or/vim_how_can_i_see_source_code_for_a_moduletrait_i/
---
New to rust here.

When I have \`use std::cmp::Ordering\` in my source code file, how can I see what is inside that module/trait? I tried "gf" in vim but it didn't work.
