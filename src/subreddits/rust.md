# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (23/2020)!](https://www.reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/gr3r32/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 341](https://www.reddit.com/r/rust/comments/gvwvep/this_week_in_rust_341/)
- url: https://this-week-in-rust.org/blog/2020/06/02/this-week-in-rust-341/
---

## [3][System76 is hiring Rust software engineers to work on Pop!_OS](https://www.reddit.com/r/rust/comments/gx73y1/system76_is_hiring_rust_software_engineers_to/)
- url: https://www.reddit.com/r/rust/comments/gx73y1/system76_is_hiring_rust_software_engineers_to/
---
I am Jeremy Soller, the Principal Engineer at System76. We are currently looking for people in the US (for tax purposes) who have experience with Rust and would like to work on the open source Rust projects that power Pop!\_OS:

[https://pop.system76.com/](https://pop.system76.com/)

The most important skills to have are familiarity with Linux and Rust. For more information including how to apply, please see the posting:

[https://system76.com/careers](https://system76.com/careers)
## [4][I've started work on a tmuxinator successor written in Rust](https://www.reddit.com/r/rust/comments/gxidy5/ive_started_work_on_a_tmuxinator_successor/)
- url: https://www.reddit.com/r/rust/comments/gxidy5/ive_started_work_on_a_tmuxinator_successor/
---
I'm one of the maintainers of [tmuxinator](https://github.com/tmuxinator/tmuxinator), a Ruby library which allows users to manage tmux sessions using project configuration files.

tmuxinator is widely used and stable, but there are regularly issues arising around installation, distribution, etc. as a result of the project being written in Ruby (its dynamic runtime, reliance on RubyGems, etc.). Installing and using the library can be problematic for people not familiar with the Ruby ecosystem. The project is available via a few different package managers, but those can/do present other problems relating to the issues mentioned above. There are ways to "simplify" the distribution of Ruby gems by shipping a Ruby interpreter along with the library or using containers, but, IMO, both of those options are too complex.

Since I'd been experimenting with Rust for some time and wanted to dive into a project that might actually be useful, I've started working on a tmuxinator successor written in Rust: [rmuxinator](https://github.com/ethagnawl/rmuxinator). In addition to greatly simplifying distribution of the utility, the program is _much_ more robust due to Rust's nature. (For instance, project configuration files are implicitly typechecked by Serde which cleaves off a bunch of ad hoc validation that happens in tmuxinator.) This is not a performance critical project, but it should eventually be much faster, too.

rmuxinator is pre-alpha, but it is functional and I plan to keep picking at it as I find time. (It'd be great to have someone else confirm that they're able to build/run it!) It supports the most basic use case of allowing a user to define a number of project configuration options (create n windows; split window n into nn panes; scope each window/session/pane to a given directory; run a series of commands in each pane, etc.) in a TOML file and it then uses that configuration to spawn a tmux session. I've also added a number of features beyond what tmuxinator currently offers (support for all tmux hooks, support for custom pane titles, etc.). rmuxinator doesn't currently concern itself with any of the CRUD features that tmuxinator offers, but I may add those at a later time. I don't use those myself and I'm not sure how often others actually do. tmuxinator also supports scripting within a project config file using ERB, but this feature actually scares me and I don't think I'll be implementing anything similar in rmuxinator.

From a Rust perspective, I feel _okay_ about what I've done so far and I've learned quite a bit. Also, Serde is fantastic! I've really enjoyed the experience and would really like to start using Rust for $WORK (I'm available for hire!). There are undoubtedly some quirky things I'm doing because I'm still new to Rust, so I was hoping to potentially get some feedback from others about those. For instance: 

* defaulting to passing references to functions even when it's not necessary (lots of other Rust programs seem to do this, but that's not a good reason to do it)
* using `#[derive(Debug,...)]` in (almost) all of my structs
* stepping through the Serde `Config` struct one portion at a time building up vectors of strings instead of pushing that behavior down into a Serde deserializer (I looked into this and the trait implementation felt overwhelming; this would also completely invert the flow of the program and require just about every line to change.)

I'm also currently punting on some things I wasn't initially sure how to address, like: 

* shelling out multiple times using `Command::new` instead of building up and executing a single shell script to start and configure the tmux session. I don't imagine this will be too difficult (either use a vector or create a temporary file), but I haven't gotten back to it yet.
* integration testing
* publishing the built library to crates.io
## [5][Learning Rust Homework Review (Jason &amp; Jonathan Turner)](https://www.reddit.com/r/rust/comments/gxmg6n/learning_rust_homework_review_jason_jonathan/)
- url: https://www.youtube.com/watch?v=Y5-ZgxfQvpc
---

## [6][What's the recommended crypto crate?](https://www.reddit.com/r/rust/comments/gxkrop/whats_the_recommended_crypto_crate/)
- url: https://www.reddit.com/r/rust/comments/gxkrop/whats_the_recommended_crypto_crate/
---
### TL;DR

What's the goto crypto crate that everyone uses and is mostly implemented in rust.  

### The Reason
I'm working on a torrent client as a fun project, and was looking for a crypto library. I only need SHA-1 functionality and settled on this [sha-1](https://crates.io/crates/sha-1) crate currently. 

I've googled searched recommendations and looked at crates.io for top [crypto crate downloads](https://crates.io/keywords/crypto?sort=downloads). I figured i was going to be using openssl but they seem to only have rust bindings on crates.io and it's currently in a failing state.
## [7][Why the developers who use Rust love it so much - Stack Overflow Blog](https://www.reddit.com/r/rust/comments/gx85qr/why_the_developers_who_use_rust_love_it_so_much/)
- url: https://stackoverflow.blog/2020/06/05/why-the-developers-who-use-rust-love-it-so-much/
---

## [8][Tour of Rust - Chapter 6 on "Text" and more chapters in a variety of language](https://www.reddit.com/r/rust/comments/gxiqg3/tour_of_rust_chapter_6_on_text_and_more_chapters/)
- url: https://tourofrust.com/chapter_6_en.html
---

## [9][Zero To Production #1: Setup - Toolchain, IDEs, CI](https://www.reddit.com/r/rust/comments/gxqr2l/zero_to_production_1_setup_toolchain_ides_ci/)
- url: https://www.lpalmieri.com/posts/2020-06-06-zero-to-production-1-setup-toolchain-ides-ci/
---

## [10][First project, written in rust](https://www.reddit.com/r/rust/comments/gxqozw/first_project_written_in_rust/)
- url: https://i.redd.it/393ziqodfa351.gif
---

## [11][Trait bound not satisfied](https://www.reddit.com/r/rust/comments/gxn345/trait_bound_not_satisfied/)
- url: https://www.reddit.com/r/rust/comments/gxn345/trait_bound_not_satisfied/
---
I'm working on a small binary (de)serialization crate for values with a fixed serialized size at compile time.

```
pub trait Encode: Sized + EncodedSize {
    fn encode(self) -&gt; [u8; &lt;Self as EncodedSize&gt;::SIZE];
}

pub trait EncodedSize: Sized {
    const SIZE: usize;
}
```

This code produces an error telling me that the trait bound EncodedSize is not satisfied (when it in fact is). I presume this is a bug in the compiler.

While it's being fixed, is there a way to work around this that does not involve dynamic memory allocation (that is, serializing to a Vec&lt;u8&gt;)? This is for a highly realtime application so I would like to avoid dynamic allocations when possible.
## [12][I wrote a small virtual machine in rust](https://www.reddit.com/r/rust/comments/gx3emt/i_wrote_a_small_virtual_machine_in_rust/)
- url: https://www.reddit.com/r/rust/comments/gx3emt/i_wrote_a_small_virtual_machine_in_rust/
---
This is my first rust program. Would love some feedback :)

[github.com/tarekwiz/smallvm](https://github.com/tarekwiz/smallvm)
