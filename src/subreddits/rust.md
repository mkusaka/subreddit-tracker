# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (16/2020)!](https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/g0erq1/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fw2hd8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (16/2020)?](https://www.reddit.com/r/rust/comments/g0es99/whats_everyone_working_on_this_week_162020/)
- url: https://www.reddit.com/r/rust/comments/g0es99/whats_everyone_working_on_this_week_162020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-16-2020/40880)!
## [3][Blog Post: Simple but Powerful Pratt Parsing](https://www.reddit.com/r/rust/comments/g0eusf/blog_post_simple_but_powerful_pratt_parsing/)
- url: https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html
---

## [4][Hyper Traps](https://www.reddit.com/r/rust/comments/g0gmn1/hyper_traps/)
- url: https://vorner.github.io/2020/04/13/hyper-traps.html
---

## [5][A Week of Experiencing Rust](https://www.reddit.com/r/rust/comments/g04rq9/a_week_of_experiencing_rust/)
- url: https://www.reddit.com/r/rust/comments/g04rq9/a_week_of_experiencing_rust/
---
I grew up as a C developer, and switched to Java a couple decades ago. A few years ago I decided to retool and chose Go instead of Rust. With my other hobbies on hold, I decided to poke around Rust. Here are 13 things I experienced:  


1. "`?`". This operator is great. It's just the right amount of syntax to say "I can't deal with this, give it to my caller." This is what I want, not Go's "`obj, err := doThing(..`" with pages of if statements. The syntax makes much more sense than Java's very verbose try/catch/throw exception handling. Make the syntax for error handling easy and I'll handle errors more often. So easy that:
2. `Option&lt;&gt;`, `Result&lt;&gt;`, and `?`. I discovered, like a lot of new people new to Rust, that "?" doesn't work on Option&lt;&gt;'s. I 100% understand why, but that doesn't change my desire that it did. That said, there's an issue to improve the compiler error message to direct people to using `ok_or_else`. Speaking of error messages:
3. `Result&lt;Thing,Box&lt;dyn Error&gt;&gt;`. I still don't know if this is the correct syntax, but it's what the compiler told me to do. I haven't read about "`dyn`" yet or why it's important.
4. `io::Error`. After figuring out Results and ?, io::Error was really confusing; I *think* it was this type. There's some type somewhere called Error that **is a** Result&lt;&gt;.
5. `Ok`, `Err`, `Some`, `None`. I discovered when navigating into the code of other crates that these make the flow of code quite readable and I enjoy using and reading them. Seeing "Some" says "it's possible for this function to not return the thing I'm looking for". Seeing "Err" says "this function can fail" - it doesn't matter how far away the function signature is, or how buried in generics-scoping-line noise that I don't understand yet it is.
6. Cross compiling. In Go, to compile an OSX binary on Linux, I set two environment variables: `GOOS=darwin` `GOARCH=amd64` and I'm set. I can do the same for Windows, and a handful of BSDs. The Rust tooling to cross-compile is still not there yet. If I `rustup target add foo`, then `cargo build --target foo`, I expect a binary that runs on **foo**. Right now I get errors about missing helpers (archiver/linkers). I know there are *reasons*, but I don't want reasons, I just want it to work. Even the `cross` dropin for Cargo seemed to not work in all cases.
7. Explicit allocation. I read somewhere that Rust forces the developer to be explicit about when it's allocating memory and passing it around. Of course, this is totally foreign to me. While I don't have an RFC-level detail solution to this, I do wish there was some sort of very shorthand syntax to say "it's okay to do whatever you need to make this work". If `my_foo.bar().baz().qux()` needs a few interim of allocations, then go for it. If you don't know the size of my type, then go ahead and autobox it for me. Maybe after I've been programming in Rust for a few months, I'll change my mind on this.
8. Matching variants. I'm using CLion, and I love the "autocomplete missing variants" and the compiler errors when I added new variants to enums. It felt like the tooling was making my life easier.
9. Lifetimes and ownership. Coming from garbage collected languages, I knew this was going to be tough. Two things I wish I found: 1) a code-free set of examples of different ways that lifetimes show up, maybe with diagrams / pictures and how to handle it. 2) Some sort of linter that would tell me if I'm making bad life choices. Right now I'm wrapping things in `Arc&lt;&gt;`, and cloning them as I need to give out copies. I have no idea the performance impact of doing this, nor reasonably what my alternatives would be. In fact:
10. Am I Rust Yet? In general, I'd love something that could check for common bad-practice idioms and suggest alternatives. Even if it was just a book of "if you do pattern X, because you want to do Y, here is the impact, and here is Z which has a similar outcome but isn't bad / is more idiomatic". Is my `Optional&lt;Vec&lt;Arc&lt;Thing&gt;&gt;&gt;` "right" / "fast" - no idea, I kept wrapping it in things until my code worked.
11. Strings. You probably know what I'm going to say. I've avoided putting references in my structs so far (because I'm not ready to tackle lifetimes). Converting between `String`, `&amp;String`, `&amp;str`, (and then I discovered there's also `OsString`) is a disappointing experience. I *know* I'm doing the wrong thing, but I'm not yet ready to invest in really understanding what, so instead I'm wrapping/unwrapping with `String::from` and `as_str`.
12. `pub(crate)` At one point CLion said "hey, this thing you're trying to access is private, can I make it visible for you?" and this is what it did. This syntax is really nice, it's clear what's going on with the scoping. Conversely:
13. Imports. It's 2020, and while CLion will reformat my imports, it won't remove unused ones, nor does it try to be helpful in bringing in the imports for traits I'm using. It's a bit disappointing and tedious. I feel like it's so close.

It may not sound this way, but my experience overall pretty positive. Rust's management of transitive dependencies is the perfect tradeoff (Go is still figuring itself out, and Java solved this (differently) a decade ago). Cargo being the defacto build tool is nice. I like that --test and --example are actual concepts. Obviously I wish compiling was faster, but I also think one of my dependencies is doing something clever in its `build.rs`. I'm also swallowing a two megabyte file with `RustEmbed` which I'm told stresses things out.
## [6][Anybody up for a code review?](https://www.reddit.com/r/rust/comments/g0gfoz/anybody_up_for_a_code_review/)
- url: https://www.reddit.com/r/rust/comments/g0gfoz/anybody_up_for_a_code_review/
---
I've been working on a tool called `dirsync` for synchronizing a local directory with one on a remote host.  It's built in Rust on top of `rsync` and `ssh`, and the source code is [here](https://github.com/spencerkohan/dirsync).

I'm still pretty new to Rust, and I am sure I am doing a lot of things in a non-optimal way, particularly it feels like I am unwrapping to often, and cloning too many strings.  It's only a handful of files, If anyone feels like talking a look and offering any constructive criticism I would be much obliged :)
## [7][Why does everyone recommend diesel when it's so confusing to use?](https://www.reddit.com/r/rust/comments/g07t7a/why_does_everyone_recommend_diesel_when_its_so/)
- url: https://www.reddit.com/r/rust/comments/g07t7a/why_does_everyone_recommend_diesel_when_its_so/
---
I'm not a fan of ORMs in general. I'd rather craft my own SQL queries (that way SQL stays fresh in my mind as well). But I'm a freshman when it comes to connecting my rust code to postgres so I decide to look up tutorials on how to do it.

Most of the tutorials I found use diesel to perform the schema translation. However, after about 3 hours of wrestling with diesel I just gave up and decided to switch to vanilla rust::postgres and manage my schema and migrations separately.

Why is diesel so popular amongst this community? The documentation is woefully inadequate. The `table!` macro hides way too much complexity and prevents any kind of intellisense. And ultimately, it is an ORM.

I'm not lamenting the presence of the crate by the way. I'm more lamenting so many "tutorials" and documentation using it. Even the rocket_contrib databases doco demonstrates diesel connectivity rather than vanilla dB connections.
## [8][Ruma is dead, long live Ruma!](https://www.reddit.com/r/rust/comments/g01iuh/ruma_is_dead_long_live_ruma/)
- url: https://www.ruma.io/news/ruma-is-dead-long-live-ruma-2020-04-10/
---

## [9][Valora: a graphics library and CLI focused on generative fine art for print](https://www.reddit.com/r/rust/comments/g05pnp/valora_a_graphics_library_and_cli_focused_on/)
- url: https://paytonturnage.gitbook.io/valora/
---

## [10][Split Vec&lt;T&gt; into Option&lt;(T, Vec&lt;T&gt;)&gt;](https://www.reddit.com/r/rust/comments/g0ezfp/split_vect_into_optiont_vect/)
- url: https://www.reddit.com/r/rust/comments/g0ezfp/split_vect_into_optiont_vect/
---
I'm trying to transform a `Vec&lt;T&gt;` into a tuple of `Option&lt;(T, Vec&lt;T&gt;)&gt;` where the tuple represents the `(head, tail)` of the `Vec` in the usual way.

The `[T]` (slice) type provides `[T]::split_first()` (and `split_first_mut` variant) but that of course returns `Option&lt;(&amp;T, &amp;[T])&gt;` rather than the owned items.

This is easy enough to implement using an extension trait using `Vec::swap_remove(usize)` (see [playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=ce980de8198ed313ba2de32f80063136)) but I feel like i'm missing an obviously better solution here?

If not then is there a case to be made for having `split_xxx_into()` methods on `Vec` directly or via some `SplitInto` extension trait with an impl for `Vec` similar to the playground example?
## [11][Any good GUI programming libraries for Rust?](https://www.reddit.com/r/rust/comments/g0g758/any_good_gui_programming_libraries_for_rust/)
- url: https://www.reddit.com/r/rust/comments/g0g758/any_good_gui_programming_libraries_for_rust/
---
I want to get into rust but could not any info on GUI programming/the info is really old. Currently, what are some good libraries I can use with Rust for desktop app development?
## [12][Trust-DNS 0.19.4 released, now with async-std support.](https://www.reddit.com/r/rust/comments/g04fxo/trustdns_0194_released_now_with_asyncstd_support/)
- url: https://github.com/bluejekyll/trust-dns/releases/tag/0.19.4
---

