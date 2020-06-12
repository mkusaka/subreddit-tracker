# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (24/2020)!](https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gyswpo/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/guo4c1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 342](https://www.reddit.com/r/rust/comments/h0mt9k/this_week_in_rust_342/)
- url: https://this-week-in-rust.org/blog/2020/06/10/this-week-in-rust-342/
---

## [3][Rust | Fuchsia](https://www.reddit.com/r/rust/comments/h7dkv2/rust_fuchsia/)
- url: https://fuchsia.dev/fuchsia-src/development/languages/rust
---

## [4][Announcing Shredder! Garbage Collection as a Library for Rust](https://www.reddit.com/r/rust/comments/h12obf/announcing_shredder_garbage_collection_as_a/)
- url: https://blog.typingtheory.com/shredder-garbage-collection-as-a-library-for-rust/
---

## [5][Error Handling in Rust - A Beginner's Guide](https://www.reddit.com/r/rust/comments/h12itt/error_handling_in_rust_a_beginners_guide/)
- url: https://www.halcyon.hr/posts/error-handling-in-rust/
---

## [6][climake - The simple, dependency-less cli framework ✨](https://www.reddit.com/r/rust/comments/h7b19a/climake_the_simple_dependencyless_cli_framework/)
- url: https://github.com/scOwez/climake
---

## [7][Where exactly do i start?](https://www.reddit.com/r/rust/comments/h7g8b2/where_exactly_do_i_start/)
- url: https://www.reddit.com/r/rust/comments/h7g8b2/where_exactly_do_i_start/
---
I am almost Completely new to typed programming languages. all i know is javascript which doesn't have types. where can i start learning rust with all the typing stuff?
## [8][From TS to Rust](https://www.reddit.com/r/rust/comments/h7jxpf/from_ts_to_rust/)
- url: https://www.reddit.com/r/rust/comments/h7jxpf/from_ts_to_rust/
---
I started to love typed languages since I got into TS. While TS' type system offers quite a lot, mixins are clumsy and create messy code/types/generics. I am currently building a typical CRUD/graphql backend. 

Would it need three right move and time to test the waters in Rust land?
## [9][Introducing GameLisp, a scripting language for Rust game development](https://www.reddit.com/r/rust/comments/h13657/introducing_gamelisp_a_scripting_language_for/)
- url: https://gamelisp.rs/
---

## [10][Announcing rust-search-extension v0.10! Search docs, crates, error codes, and attributes, etc in your address bar instantly!](https://www.reddit.com/r/rust/comments/h0yqcl/announcing_rustsearchextension_v010_search_docs/)
- url: https://www.reddit.com/r/rust/comments/h0yqcl/announcing_rustsearchextension_v010_search_docs/
---
Hi all! I'm so happy to announce that **Rust Search Extension** has released a new version of 0.10, and I got sponsorship from [vercel.com](https://vercel.com/?utm_source=rust-search-extension).

&gt;Rust Search Extension's official website host on `now.sh`, which is one of the fascinating products belong to  [vercel.com](https://vercel.com/?utm_source=rust-search-extension). They give me a $20/mo free plan. Woww! 

https://preview.redd.it/f62866x0y9451.png?width=1382&amp;format=png&amp;auto=webp&amp;s=181e8e773e75ea6261f8dbc8a93fff96c82201bf

Version 0.10 brings some great features, and here is the detail:

# ~ (tilde) prefix to search the external crate docs

`~` is another pretty prefix to search the external crate docs. Unlike the `@` prefix, `~` gives you the power to search docs across all your favorite crates effortlessly.

For example, input `~spawn` , you can get results related to "`spawn`" from all your favorite crates (e.g. `tokio`, `async_std`, `smol`, and `futures` ).

https://preview.redd.it/6f377pfwy9451.png?width=1902&amp;format=png&amp;auto=webp&amp;s=23a9f96d1591567a3485f9b3288ca36e8fc1a661

# Searching statistics

Searching statistics page provides a set of useful charts to visualize your search history. This gives you a powerful insight view on how frequency you search, how much valuable time you saved.

These charts including:

* A calendar heatmap represents your searching history in the last year
* A percentage chart reflects each category weight you searched
* Weekly, daily, and hourly bar chart tells you searching frequency in different time dimensionality
* Top searched crates chart gives you the rank of personal crate searching

https://preview.redd.it/fojiu0t3z9451.png?width=2324&amp;format=png&amp;auto=webp&amp;s=2fec4d7eaad56c6998ad5a57aa653cbb3b600d57

https://preview.redd.it/vh8b6sy4z9451.png?width=2094&amp;format=png&amp;auto=webp&amp;s=b52f40094eb39d948479d919ff96afb102c0324f

## Open repository quickly

You can prefix `!!!` (triple !) to open crate's repository quickly. For example, input `!!!tokio`  
 then enter, the extension will obtain the repository url then redirect to the github repository page of `tokio` effortlessly.

https://preview.redd.it/7pfbpyr8z9451.png?width=1876&amp;format=png&amp;auto=webp&amp;s=359fe880be6e1acfad629ab5e4af1fd880d472c9

https://preview.redd.it/6xcv9t0az9451.png?width=1802&amp;format=png&amp;auto=webp&amp;s=bc6e2e208fd687cabe5dacfc38647445448eccee

## Customize crates searching platform

You can customize your preferred crates searching platform such as [crates.io](https://crates.io/)(default) or [lib.rs](https://lib.rs/) on the popup page.

# Miscellaneous

* Improve searching speed, approximately 10% faster than the previous version
* New commands:
   * `:tool` - search rust tools
* Bugfix:
   * Fix std doc offline mode bug
   * Fix wrong Add button state in docs.rs's src page issue

# Installation

\- [Chrome Webstore](https://chrome.google.com/webstore/detail/rust-search-extension/ennpfpdlaclocpomkiablnmbppdnlhoh)

\- [Firefox](https://addons.mozilla.org/firefox/addon/rust-search-extension/)

\- [Microsoft Edge](https://microsoftedge.microsoft.com/addons/detail/olemfibpaicdoooacpfffccidjjagmoe)

# Some links

\- Github: [https://github.com/huhu/rust-search-extension](https://github.com/huhu/rust-search-extension)

\- Website: [https://rust-search-extension.now.sh/](https://rust-search-extension.now.sh/)

&amp;#x200B;

Welcome to give it a try. If you found a bug or something missed, please don't hesitate to let me know. Of course, any issue or PR is welcome too! Thanks! :)
## [11][derive_aktor - a macro based actor library that provides typed, nominal actors](https://www.reddit.com/r/rust/comments/h7k109/derive_aktor_a_macro_based_actor_library_that/)
- url: https://www.reddit.com/r/rust/comments/h7k109/derive_aktor_a_macro_based_actor_library_that/
---
[https://github.com/insanitybit/derive\_aktor](https://github.com/insanitybit/derive_aktor)

&amp;#x200B;

Hi all. I've posted this before, years ago when I'd first started it. It's come a long way since then. This library lets you write actor oriented code, as a number of Rust libraries at this point do.

Where this library is different is:

* It's macro-based. You write sync-looking code, add a macro, it generates all of the routing, message handling, message and actor types, etc.
* It's strictly typed with, in general, static dispatch.
* The macro generates a nominal API, whereas many libraries often expose a single actor method like "send"

The end result of this is that, at least to me, it feels very much like you're just writing async rust, but actors allow you to structure your program in a number of ways that I personally like.

I think others may find this useful, and in particular I'd love to see someone contribute because I'm really not very good at writing proc macros. I've got some open issues like that you can't use a where clause on the impl block, the generics the Actor struct uses are unnecessary in some cases, tracing support should be feature-flagged, etc. This is, personally, my least favorite issue, and is very much help-wanted:

[https://github.com/insanitybit/derive\_aktor/issues/16](https://github.com/insanitybit/derive_aktor/issues/16)
## [12][Tree, lifetimes, "cannot borrow"... Can I get some help?](https://www.reddit.com/r/rust/comments/h7h005/tree_lifetimes_cannot_borrow_can_i_get_some_help/)
- url: https://www.reddit.com/r/rust/comments/h7h005/tree_lifetimes_cannot_borrow_can_i_get_some_help/
---
[Playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=d6c6bbc898879395451a0d96e501b79f) (commented, easy to follow, w/ usage example)

This piece of code gives me 2 errors I believe I understand but don't know how to fix.

SEND HELP!

Thank you
