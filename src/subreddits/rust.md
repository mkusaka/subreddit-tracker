# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (37/2020)!](https://www.reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What’s everyone working on this week (37/2020)?](https://www.reddit.com/r/rust/comments/ioc6ld/whats_everyone_working_on_this_week_372020/)
- url: https://www.reddit.com/r/rust/comments/ioc6ld/whats_everyone_working_on_this_week_372020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-37-2020/48445?u=llogiq)!
## [3][The first two posts about making a language with Rust are out!](https://www.reddit.com/r/rust/comments/ip9v9g/the_first_two_posts_about_making_a_language_with/)
- url: https://arzg.github.io/lang/0
---

## [4][As a hobbyist rust developer, I want to think less about error handling](https://www.reddit.com/r/rust/comments/ip5dhv/as_a_hobbyist_rust_developer_i_want_to_think_less/)
- url: https://mbuffett.com/posts/rust-less-error-handling/
---

## [5][Map/Set traits to be generic over (Hash|BTree)(Map|Set)?](https://www.reddit.com/r/rust/comments/ipbz03/mapset_traits_to_be_generic_over_hashbtreemapset/)
- url: https://www.reddit.com/r/rust/comments/ipbz03/mapset_traits_to_be_generic_over_hashbtreemapset/
---
When writing algorithms that process a map (or set), I often write functions over `HashMap` for simplicity. Indeed, `HashMap` is now state-of-the-art optimized and suits for most cases. However, most algorithms are actually agnostic to the underlying map implementation - the only requirements are a collection supporting map operations: `insert()`, `remove()`, `entry()`, etc. So they may as well work over a `BTreeMap` if it's what the caller has.

In particular, although `HashMap` works best in most cases, there are still use cases where `BTreeMap` is preferable.

- In `no_std` environments (but with liballoc enabled), only `BTreeMap` is available (I believe it's due to the fact that no randomness source is available in bare-metal `no_std` to randomize the hash map and avoid DoS attacks?).
- When the caller already has a `BTreeMap` available (e.g. for other reasons the map has to be sorted), converting it into a temporary `HashMap` may not be beneficial.
- For key types that - for some reason - don't have a hash function available, or are expensive to hash (but cheap to order)?


I haven't seen any `Map` trait (that `HashMap` and `BTreeMap` would implement) in the standard library to encompass that, so I was curious about why.

- Did the use case I mention not arise yet?
- Is the use case too narrow to include in the stdlib - and should rather be implemented by a separate crate?
- Is such a trait already available in a crate (I must admit I haven't searched thoroughly)?
- Would it be more complex to define and use such a trait than I naively think?
## [6][Introducing inline_tweak: tweak values directly from the source](https://www.reddit.com/r/rust/comments/ip30up/introducing_inline_tweak_tweak_values_directly/)
- url: https://www.reddit.com/r/rust/comments/ip30up/introducing_inline_tweak_tweak_values_directly/
---
I recently published [inline_tweak](https://crates.io/crates/inline_tweak), a crate well received by the gamedev community, so I thought it could assist some people here.  

The idea is to easily tweak any number/boolean literal where modifying the value in the _source code_ also changes it in the program.  

A small example to get the idea:

    loop {
        // try changing the value while the program is running
        println!("{}", inline_tweak::tweak!(3.0));
    }

It also has a bunch of cool extra features which you can find [in the README](https://github.com/Uriopass/inline_tweak).

It is disabled in release mode and is extra light (only 1 dependency), so you can keep it around and it won't impact neither compile times nor performance.

How does it work? Thanks to the `file!()`, `line!()` and `column!()` macros, `inline_tweak` watches the file and parses the changes which it keeps in a `lazy_static` hashmap.
## [7][Introducing `auditable`: audit Rust binaries for known bugs or vulnerabilities in production](https://www.reddit.com/r/rust/comments/iotx5u/introducing_auditable_audit_rust_binaries_for/)
- url: https://www.reddit.com/r/rust/comments/iotx5u/introducing_auditable_audit_rust_binaries_for/
---
Rust is very promising for security-critical applications due to its memory safety guarantees. However, while vulnerabilities in Rust crates are rare, they [still exist](https://rustsec.org/advisories/), and Rust is currently missing the tooling to deal with them.

For example, Linux distros alert you if you're running a vulnerable version, and you can even opt in to automatic security updates. Cargo not only has no security update infrastructure, it doesn't even know which libraries or library versions went into compiling a certain binary, so there's no way to check if your system is vulnerable or not.

 I've embarked on a quest to fix that.

Today I'm pleased to announce the initial release of `auditable` crate. It embeds the dependency tree into the compiled executable so you can check which crates exactly were used in the build. The primary motivation is to make it possible to answer the question **"Do the Rust binaries we're actually running in production have any known vulnerabilities?"** - and even enable third parties such as cloud providers to automatically do that for you.

We provide [crates](http://docs.rs/auditable_extract/) to consume this information and easily build your own tooling, and a converter to Cargo.lock format for compatibility with existing tools. This information can already be used in conjunction with [cargo-audit](https://github.com/RustSec/cargo-audit), see example usage [here](https://github.com/Shnatsel/rust-audit#demo).

See the [repository](https://github.com/Shnatsel/rust-audit) for a demo and more info on the internals, including the [frequently asked questions](https://github.com/Shnatsel/rust-audit#faq) such as binary bloat.

The end goal is to integrate this functionality in Cargo and enable it by default on all platforms that are not tightly constrained on the size of the executable. A yet-unmerged RFC to that effect can be found [here](https://github.com/rust-lang/rfcs/pull/2801). Right now the primary blockers are:

1. [This bug in rustc](https://github.com/rust-lang/rust/issues/47384) is blocking a proper implementation that could be uplifed into Cargo.
1. We need to get some experience with the data format before we stabilize it.

If you're running production Rust workloads and would like to be able to audit them for security vulnerabilites, please get in touch. I'd be happy to assist deploying `auditable` used in a real-world setting to iron out the kinks.

And if you can hack on rustc, [you know what to do](https://github.com/rust-lang/rust/issues/47384) ;)
## [8][** Rust &amp;&amp; C++ London Joint Meetup**](https://www.reddit.com/r/rust/comments/ipf4sp/rust_c_london_joint_meetup/)
- url: https://www.reddit.com/r/rust/comments/ipf4sp/rust_c_london_joint_meetup/
---
Rust London and C++ London are proud to announce our first joint meetup The Rust &amp;&amp; C++ LDN Talks. This is the first in a series of collaborative events between our two communities which have been compared and contrasted for several years.  


The purpose of these events is to begin an exchange of ideas and experiences between our user groups.  
The Rust London User Group will open up the LDN Talks platform to the C++ London community. We will have speakers from C++ London who will share their perspectives, opinions and learning experiences with the Rust Programming language.

&amp;#x200B;

Check out the Meetup Page below

[https://www.meetup.com/Rust-London-User-Group/events/273056379/](https://www.meetup.com/Rust-London-User-Group/events/273056379/)
## [9][What makes Actix faster than other frameworks?](https://www.reddit.com/r/rust/comments/ip88wb/what_makes_actix_faster_than_other_frameworks/)
- url: https://www.reddit.com/r/rust/comments/ip88wb/what_makes_actix_faster_than_other_frameworks/
---
Why is it that Actix is so much faster than other web frameworks? Is speed a consequence of the actor model behind Actix or something else?

&amp;#x200B;

[https://www.techempower.com/benchmarks/#section=data-r18](https://www.techempower.com/benchmarks/#section=data-r18)
## [10][Is there a go-to distributed tasks queue in Rust?](https://www.reddit.com/r/rust/comments/ip438x/is_there_a_goto_distributed_tasks_queue_in_rust/)
- url: https://www.reddit.com/r/rust/comments/ip438x/is_there_a_goto_distributed_tasks_queue_in_rust/
---
I'm looking for something like Celery is for Python, maybe slightly more lightweight since I don't care about periodic tasks, or multiple possible backends -- I just need a solution that works for basic FIFO jobs. I've found https://github.com/rusty-celery/rusty-celery -- has anyone managed to use it in production? It seems that the project is well maintained but I'm afraid it's still in early stages of development.
## [11][`const_trait_impl` is pretty sweet](https://www.reddit.com/r/rust/comments/iouec5/const_trait_impl_is_pretty_sweet/)
- url: https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=19df65c14cd0f82b1f0b4a957cf7defd
---

## [12][Juniper - fast &amp; type-safe GraphQL servers in Rust](https://www.reddit.com/r/rust/comments/ioskx6/juniper_fast_typesafe_graphql_servers_in_rust/)
- url: https://blog.graphqleditor.com/graphql-rust-juniper/
---

