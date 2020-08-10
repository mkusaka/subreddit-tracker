# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (33/2020)!](https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (33/2020)?](https://www.reddit.com/r/rust/comments/i6yras/whats_everyone_working_on_this_week_332020/)
- url: https://www.reddit.com/r/rust/comments/i6yras/whats_everyone_working_on_this_week_332020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-33-2020/47095?u=llogiq)!
## [3][Meili raises 1.5M€ for open source search in Rust](https://www.reddit.com/r/rust/comments/i71w2z/meili_raises_15m_for_open_source_search_in_rust/)
- url: https://blog.meilisearch.com/meili-fundraise/
---

## [4][Shiika is a statically-typed scripting language](https://www.reddit.com/r/rust/comments/i70f3s/shiika_is_a_staticallytyped_scripting_language/)
- url: https://github.com/yhara/shiika
---

## [5][Surviving Rust async interfaces - fasterthanli.me](https://www.reddit.com/r/rust/comments/i6p35d/surviving_rust_async_interfaces_fasterthanlime/)
- url: https://fasterthanli.me/articles/surviving-rust-async-interfaces
---

## [6][Announcing `polling-async-trait`! We can't have "real" async trait methods yet, so this is the next best thing](https://www.reddit.com/r/rust/comments/i6yuam/announcing_pollingasynctrait_we_cant_have_real/)
- url: https://www.reddit.com/r/rust/comments/i6yuam/announcing_pollingasynctrait_we_cant_have_real/
---
Inspired by the excellent `async-trait`, I wanted to see if I could create an abstraction for `async` methods in traits that was "truely" 0-cost (that is, that returns a concrete `Future` type, rather than something boxed and type-erased.) While it's currently impossible to do with an `async fn` on a trait without GATs, it turns out we can do the next best thing: asyncify `poll` methods on your trait. Thus [`polling-async-trait`](https://docs.rs/polling-async-trait/) was born.

The basic interface is very similar to `async-trait`; tag a trait with an attribute that you want to asyncify, as well as the specific polling methods:

    // This trait will have the async methods `flush`
    // and `close` generated for it.
    #[async_poll_trait]
    trait SimpleIO {
        #[async_method]
        fn poll_flush(&amp;mut self, cx: &amp;mut Context&lt;'_&gt;) -&gt; Poll&lt;io::Result&lt;()&gt;&gt;;

        // `owned` means that the future returned by `close`
        // will take irrevocable ownership of this object
        #[async_method(owned)]
        fn poll_close(&amp;mut self, cx: &amp;mut Context&lt;'_&gt;) -&gt; Poll&lt;io::Result&lt;()&gt;&gt;;
    }

This allows you to write:

    async fn flush_and_close(data: impl SimpleIO) -&gt; io::Result&lt;()&gt; {
        data.flush().await?;
        data.close().await?;
        Ok(());
    }

Try it out and let me know what you think! As always I'm open to feedback and critiques. In particular for this library I definitely want to know if anyone can produce unsoundness with the things it does with `Pin`; I'm 98% sure I covered all the bases but I'd love to have a second set of eyes on it. My hope is to move it to 1.0 ASAP once I'm confident in its soundness.
## [7][The promise of Rust async-await for embedded](https://www.reddit.com/r/rust/comments/i719p1/the_promise_of_rust_asyncawait_for_embedded/)
- url: https://tweedegolf.nl/blog/41/the-promise-of-rust-async-await-for-embedded
---

## [8][Cilk: (Toy) Compiler Infrastructure influenced by LLVM written in Rust](https://www.reddit.com/r/rust/comments/i70ftm/cilk_toy_compiler_infrastructure_influenced_by/)
- url: https://github.com/maekawatoshiki/cilk
---

## [9][ugrep: new ultrafast C++ grep claims to be faster than ripgrep](https://www.reddit.com/r/rust/comments/i6pfb2/ugrep_new_ultrafast_c_grep_claims_to_be_faster/)
- url: https://github.com/Genivia/ugrep
---

## [10][Let’s implement a Bloom Filter](https://www.reddit.com/r/rust/comments/i6qk6x/lets_implement_a_bloom_filter/)
- url: https://onatm.dev/2020/08/10/let-s-implement-a-bloom-filter/
---

## [11][Comparison of XML and HTML parsing libraries](https://www.reddit.com/r/rust/comments/i700dz/comparison_of_xml_and_html_parsing_libraries/)
- url: https://www.reddit.com/r/rust/comments/i700dz/comparison_of_xml_and_html_parsing_libraries/
---
[https://github.com/alexkreidler/rust-html-xml-comparison](https://github.com/alexkreidler/rust-html-xml-comparison)

So far I've copied in the examples from each library and written basic exemplary tests for the lower-level XML libs. It's not fully done, but provides a good overview of the landscape.

Next steps are to automatically test their capabilities on in-the-wild HTML and XML.
## [12][How To Bootstrap A Rust Web API From Scratch - Zero To Production #3](https://www.reddit.com/r/rust/comments/i6o1q4/how_to_bootstrap_a_rust_web_api_from_scratch_zero/)
- url: https://www.lpalmieri.com/posts/2020-08-09-zero-to-production-3-how-to-bootstrap-a-new-rust-web-api-from-scratch/
---

