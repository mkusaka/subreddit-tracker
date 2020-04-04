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

## [3][Why I’m building a new async runtime](https://www.reddit.com/r/rust/comments/fudgku/why_im_building_a_new_async_runtime/)
- url: https://stjepang.github.io/2020/04/03/why-im-building-a-new-async-runtime.html
---

## [4][Introducing Kubie: a more powerful alternative to kubectx and kubens, written in Rust](https://www.reddit.com/r/rust/comments/fut560/introducing_kubie_a_more_powerful_alternative_to/)
- url: https://blog.sbstp.ca/introducing-kubie/
---

## [5][[OC] Fondant - macro based, configuration management library for CLI utils](https://www.reddit.com/r/rust/comments/fu8c0p/oc_fondant_macro_based_configuration_management/)
- url: https://i.redd.it/akvcc4e7hlq41.png
---

## [6][I've made a templates engine to ease writing repetitive boilerplate code [CHROBRY]](https://www.reddit.com/r/rust/comments/fuobnj/ive_made_a_templates_engine_to_ease_writing/)
- url: https://www.reddit.com/r/rust/comments/fuobnj/ive_made_a_templates_engine_to_ease_writing/
---
Hi Rustaceans!

I've made a crate (and CLI app) that generates your target language files based on template file with data types definitions and behaviours assigned to them. Loosely inspired by Rust structs and traits.

I'm using it to generate serde-like serializable data for my Unreal Engine 4 game project, i think you might find it useful too :)

Repo: [https://github.com/PsichiX/Chrobry](https://github.com/PsichiX/Chrobry)
## [7][PSA: docs.rs/std redirects to doc.rust-lang.org/std](https://www.reddit.com/r/rust/comments/fuk39x/psa_docsrsstd_redirects_to_docrustlangorgstd/)
- url: https://www.reddit.com/r/rust/comments/fuk39x/psa_docsrsstd_redirects_to_docrustlangorgstd/
---
Title. Just in case someone didn't know (just like me).
## [8][Impact of TLS dependency removal from async internals](https://www.reddit.com/r/rust/comments/fua76v/impact_of_tls_dependency_removal_from_async/)
- url: https://www.reddit.com/r/rust/comments/fua76v/impact_of_tls_dependency_removal_from_async/
---
TL;DR: The compiler is now able to optimize much better than before.

I wanted to see the impact of recent TLS dependency removal from async internal implementation (https://github.com/rust-lang/rust/pull/69033)

Here is the toy example:

```
async fn foo() -&gt; u32 {
    3
}

async fn square() -&gt; u32 {
    let a = foo().await;
    let b = foo().await;
    a * b
}

pub fn xx(cx: &amp;mut Context&lt;'_&gt;) -&gt; u32 {
    let mut f = square();
    let mut f = unsafe { Pin::new_unchecked(&amp;mut f) };
    loop {
        match f.as_mut().poll(cx) {
            Poll::Ready(x) =&gt; break x,
            Poll::Pending =&gt; {}
        }
    }
}
```

On stable 1.42.0: https://rust.godbolt.org/z/WfZMRs

On current nightly: https://rust.godbolt.org/z/jL_5sC

It would be interesting to see some real world differences.
## [9][Rust and Computational Science?](https://www.reddit.com/r/rust/comments/fugyk0/rust_and_computational_science/)
- url: https://www.reddit.com/r/rust/comments/fugyk0/rust_and_computational_science/
---
hi,

i am interested in a Master's Degree in Computational Science (Neuroscience) / Computational Neuroscience. Doing a bit of research I came across this [Wiki article](https://en.wikipedia.org/wiki/Computational_science).

It lists a few programming languages that are predominately used in that field: Fortran, Octave, Haskell, Julia, Maple, Matlab, Python, Perl, R , etc.

Rust isn't listed. Is that just because of the relative novelty of the language or are there some features missing that would make rust a bad choice language for that field?

what do those languages have that rust doesn't (yet)?
## [10][From&lt;Option&lt;T&gt;&gt; for Option&lt;U&gt; is not implemented?](https://www.reddit.com/r/rust/comments/fui7g4/fromoptiont_for_optionu_is_not_implemented/)
- url: https://www.reddit.com/r/rust/comments/fui7g4/fromoptiont_for_optionu_is_not_implemented/
---
Is there a reason that the follow blanket impl is not actually in core?  

``` rust
impl&lt;T, U: From&lt;T&gt;&gt; From&lt;Option&lt;T&gt;&gt; for Option&lt;U&gt; {
    fn from(t: Option&lt;T&gt;) -&gt; Option&lt;U&gt; {
        if let Some(inner_t) = t {
            Some(t.into())
        } else {
            None
        }
    }
}
```
## [11][Rust has to rework its error reporting](https://www.reddit.com/r/rust/comments/fubeek/rust_has_to_rework_its_error_reporting/)
- url: https://www.reddit.com/r/rust/comments/fubeek/rust_has_to_rework_its_error_reporting/
---
I really like rust and a in the beginning a great feature was the detailed error reporting where you could directly see what's wrong and how you could fix it.

I am afraid that now with `futures` rust error messages became really bad. Quite often I get error messages like these:

    error[E0599]: no method named `into_stream` found for type `futures_util::stream::stream::forward::Forward&lt;impl futures_core::stream::TryStream, futures_util::sink::map_err::SinkMapErr&lt;futures_util::sink::fanout::Fanout&lt;futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;, futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;&gt;, [closure@src/main.rs:218:47: 218:101]&gt;&gt;` in the current scope
       --&gt; src/main.rs:220:19
        |
    220 |     let foo = foo.into_stream();
        |                   ^^^^^^^^^^^ method not found in `futures_util::stream::stream::forward::Forward&lt;impl futures_core::stream::TryStream, futures_util::sink::map_err::SinkMapErr&lt;futures_util::sink::fanout::Fanout&lt;futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;, futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;&gt;, [closure@src/main.rs:218:47: 218:101]&gt;&gt;`
        |
        = note: the method `into_stream` exists but the following trait bounds were not satisfied:
                `&amp;futures_util::stream::stream::forward::Forward&lt;impl futures_core::stream::TryStream, futures_util::sink::map_err::SinkMapErr&lt;futures_util::sink::fanout::Fanout&lt;futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;, futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;&gt;, [closure@src/main.rs:218:47: 218:101]&gt;&gt; : futures_util::future::future::FutureExt`
                `&amp;futures_util::stream::stream::forward::Forward&lt;impl futures_core::stream::TryStream, futures_util::sink::map_err::SinkMapErr&lt;futures_util::sink::fanout::Fanout&lt;futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;, futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;&gt;, [closure@src/main.rs:218:47: 218:101]&gt;&gt; : futures_util::stream::try_stream::TryStreamExt`
                `&amp;mut futures_util::stream::stream::forward::Forward&lt;impl futures_core::stream::TryStream, futures_util::sink::map_err::SinkMapErr&lt;futures_util::sink::fanout::Fanout&lt;futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;, futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;&gt;, [closure@src/main.rs:218:47: 218:101]&gt;&gt; : futures_util::future::future::FutureExt`
                `&amp;mut futures_util::stream::stream::forward::Forward&lt;impl futures_core::stream::TryStream, futures_util::sink::map_err::SinkMapErr&lt;futures_util::sink::fanout::Fanout&lt;futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;, futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;&gt;, [closure@src/main.rs:218:47: 218:101]&gt;&gt; : futures_util::stream::try_stream::TryStreamExt`
                `futures_util::stream::stream::forward::Forward&lt;impl futures_core::stream::TryStream, futures_util::sink::map_err::SinkMapErr&lt;futures_util::sink::fanout::Fanout&lt;futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;, futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;&gt;, [closure@src/main.rs:218:47: 218:101]&gt;&gt; : futures_util::future::future::FutureExt`
                `futures_util::stream::stream::forward::Forward&lt;impl futures_core::stream::TryStream, futures_util::sink::map_err::SinkMapErr&lt;futures_util::sink::fanout::Fanout&lt;futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;, futures_channel::mpsc::UnboundedSender&lt;(bytes::bytes::Bytes, std::net::SocketAddr)&gt;&gt;, [closure@src/main.rs:218:47: 218:101]&gt;&gt; : futures_util::stream::try_stream::TryStreamExt`
    

I think that async and futures are really amazing tools that help you to write good and efficient code, but coding with unreadable error messages like these is no longer fun.

Do you think this situation can be resolved somehow?
## [12][Mutable Slices &amp; Borrow Checker](https://www.reddit.com/r/rust/comments/fuiwze/mutable_slices_borrow_checker/)
- url: https://www.reddit.com/r/rust/comments/fuiwze/mutable_slices_borrow_checker/
---
This post is inspired by [this tweet](https://twitter.com/plaidfinch/status/1246200064332234758) challenging everyone to get the linked piece of code ([link to playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=a7a8e3e6b08039370bcc4caa10212e4d)) working.

However, it is not quite clear to me why this piece of code fails in the first place; according to the compiler error message, the call to `slice.get_mut(index)` in line 8 mutably borrows `slice` in the previous loop iteration; but shouldn't that reference have been dropped in the previous loop iteration as well?

If I am mistaken, what is the reason this piece of code doesn't pass the borrow checker?

Looking forward to your opinions!
