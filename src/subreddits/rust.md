# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (11/2020)!](https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 329](https://www.reddit.com/r/rust/comments/fhcuec/this_week_in_rust_329/)
- url: https://this-week-in-rust.org/blog/2020/03/10/this-week-in-rust-329/
---

## [3][beef::Cow, a more compact std::borrow::Cow](https://www.reddit.com/r/rust/comments/fif7ob/beefcow_a_more_compact_stdborrowcow/)
- url: https://github.com/maciejhirsz/beef/
---

## [4][Archetypal vs Grouped ECS architectures](https://www.reddit.com/r/rust/comments/fiemdi/archetypal_vs_grouped_ecs_architectures/)
- url: https://community.amethyst.rs/t/archetypal-vs-grouped-ecs-architectures-my-take/1344?u=erlend_sh
---

## [5][This Week in Rust is looking for a new maintainer. | Inside Rust Blog](https://www.reddit.com/r/rust/comments/fi1g1u/this_week_in_rust_is_looking_for_a_new_maintainer/)
- url: https://blog.rust-lang.org/inside-rust/2020/03/13/twir-new-lead.html
---

## [6][s3-algo: High-level, high-performance algorithms for S3 batch operations, built on top of Rusoto](https://www.reddit.com/r/rust/comments/figp3a/s3algo_highlevel_highperformance_algorithms_for/)
- url: https://www.reddit.com/r/rust/comments/figp3a/s3algo_highlevel_highperformance_algorithms_for/
---
[https://crates.io/crates/s3-algo](https://crates.io/crates/s3-algo)

(more info in repository README: [https://github.com/openanalytics/s3-algo](https://github.com/openanalytics/s3-algo))

I made this library originally to upload lots of files to S3, where previously the aws cli was used. The result is a very satisfactory performance boost, with the added feature that we can track the progress of all such multi-file uploads.

I made [s3-cli](https://github.com/Ploppz/s3-cli) (in its very infancy) in order to benchmark performance on a real S3 bucket in eu-west-1, and compare it with \`aws s3 cp ...\`. I'm sure the latter is not at all optimized for this so I suppose it's not a fair comparison. In the project I work on however, they had been using aws cli for this task up until now, so in my case it's a relevant comparison. I'm open to other benchmarking ideas! \`s3-algo\` copied 406 files totaling 1.1 GB in 6 seconds on average, and \`aws s3 cp\` on average 20 seconds. (benchmarked with \`hyperfine\`)

&amp;#x200B;

Currently supported operations:

* upload files (or in-memory data)
* list objects with a certain prefix or any other requirement
* delete or copy all objects that were listed

&amp;#x200B;

General features:

* stream files to upload requests
* detailed data about each request is gathered, and processed in programmer-given closure
* as generic as possible
* focus on parallelization
* aggressive timeouts for good performance
* retry mechanism and exponential back-off

&amp;#x200B;

Things to consider:

* I have just made a few functions, but maybe another type of abstraction would be nicer - e.g. extension trait or something like that?
* One existing project I have found is \`s4\`, or its fork \`s3-ext\`. Would it be an idea to merge efforts? So far it seems like \`s3-algo\` and \`s3-ext\` complement each other, as \`s3-ext\` focuses on doing single \[multipart\] file upload properly, while \`s3-algo\` is optimized for large number of smaller files, so far without multipart upload, etags or any more advanced S3 feature.

&amp;#x200B;

I eagerly welcome any suggestions for change (be it in the interface or internals) or PRs.
## [7][analyze-reify: a Rust/tree-sitter based CLI that analyzes clojure code for occurrences of reify](https://www.reddit.com/r/rust/comments/fig9t5/analyzereify_a_rusttreesitter_based_cli_that/)
- url: https://www.reddit.com/r/rust/comments/fig9t5/analyzereify_a_rusttreesitter_based_cli_that/
---
I think it's always fun to build useful tools while learning a new language. To me Rust is new, but I'm familiar with Clojure and I already built a few tools around Clojure code. This time I used Rust and tree-sitter to analyze Clojure code:  


[https://github.com/borkdude/analyze-reify](https://github.com/borkdude/analyze-reify)
## [8][Oxidize 1K: A Remote Embedded Rust Conference](https://www.reddit.com/r/rust/comments/fi3da0/oxidize_1k_a_remote_embedded_rust_conference/)
- url: https://oxidizeconf.com/oxidize-1k/
---

## [9][Looking for a guide (post/book fragment) on `to`, `as` and `into`.](https://www.reddit.com/r/rust/comments/fi63tl/looking_for_a_guide_postbook_fragment_on_to_as/)
- url: https://www.reddit.com/r/rust/comments/fi63tl/looking_for_a_guide_postbook_fragment_on_to_as/
---
I recall reading some document that highlighted the differences between them. I think it was a guide on designing APIs. Thanks!
## [10][Why does rust not support default arguments?](https://www.reddit.com/r/rust/comments/fi6nov/why_does_rust_not_support_default_arguments/)
- url: https://www.reddit.com/r/rust/comments/fi6nov/why_does_rust_not_support_default_arguments/
---
Just curious. Is it a philosophical stance or an implementation issue, etc?
## [11][New process-viewer release (whole new network viewer)!](https://www.reddit.com/r/rust/comments/fi5tiy/new_processviewer_release_whole_new_network_viewer/)
- url: https://blog.guillaume-gomez.fr/articles/2020-03-13+New+process-viewer+release
---

## [12][Why can't I borrow inmutably in this context?](https://www.reddit.com/r/rust/comments/fig3v6/why_cant_i_borrow_inmutably_in_this_context/)
- url: https://www.reddit.com/r/rust/comments/fig3v6/why_cant_i_borrow_inmutably_in_this_context/
---
Here the playground: https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=abe473c4500154759826602f056a55ba

If I don't touch the borrowed value after calling readonly it is ok, but if I do it complains.  
I don't understand where the problem is since by the time I modify it the second time the readonly function has already returned and it could not possibly change self since its an inmutable reference.  
EDIT:

It is also interesting that I can just reborrow mutably and keep going.  

    pub fn run(mut self) {
        println!("I consume myself");
        let text = &amp;mut self.text;
        *text = "hello".to_owned();
        self.readonly();
        let text = &amp;mut self.text;
        *text = "done".to_owned(); // This actually works
    }
