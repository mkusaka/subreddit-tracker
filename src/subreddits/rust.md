# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (19/2020)!](https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/gd6g9w/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/g9a5sn/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 337](https://www.reddit.com/r/rust/comments/geagy0/this_week_in_rust_337/)
- url: https://this-week-in-rust.org/blog/2020/05/05/this-week-in-rust-337/
---

## [3][async-std now using smol runtime](https://www.reddit.com/r/rust/comments/gg7gnl/asyncstd_now_using_smol_runtime/)
- url: https://github.com/async-rs/async-std/pull/757
---

## [4][A practical introduction to async programming in Rust](https://www.reddit.com/r/rust/comments/ggdw7p/a_practical_introduction_to_async_programming_in/)
- url: http://jamesmcm.github.io/blog/2020/05/06/a-practical-introduction-to-async-programming-in-rust/#en
---

## [5][Auto-currying Rust Functions: A guide to proc-macros](https://www.reddit.com/r/rust/comments/ggaetv/autocurrying_rust_functions_a_guide_to_procmacros/)
- url: https://peppe.rs/posts/auto-currying_rust_functions/
---

## [6][Async-GraphQL: A GraphQL Server Framework](https://www.reddit.com/r/rust/comments/gg71es/asyncgraphql_a_graphql_server_framework/)
- url: https://github.com/async-graphql/async-graphql
---

## [7][noob looking for a code review](https://www.reddit.com/r/rust/comments/ggde80/noob_looking_for_a_code_review/)
- url: https://www.reddit.com/r/rust/comments/ggde80/noob_looking_for_a_code_review/
---
the crux of the matter: [https://gitlab.com/jgtaylor/clusterman](https://gitlab.com/jgtaylor/clusterman)

Hello Rustaceans - I am very new with Rust. I completed the first 3/4 of the Rust Book (the online one) - before they get into advanced traits, macro stuff, etc. As a practice, I wanted to try to build a kind of "cluster manager" via ZMQ.

As I started, I realized how much "patterns" helped me learn Python/NodeJS. I also realized that I'm not able to quickly see these patterns in rust (yet). Anyway, I tried for a "Factory pattern" or "constructor pattern" - not sure what to call it in Rustean.

One of the things I was curious about - (maybe i'm an idiot because i'm thinking in terms of dicts (py) or Object (js)) - I'd wanted a structure with a vec!\[\] of cluster members, and one of them would be a "master", such that the "master" was just a reference to one of the cluster members in the Vec. I couldn't make it work because - ultimately, I kind of realized that the "master" member entry could be deleted resulting in the "master" being a dangling pointer. Was/is my conclusion correct?

This kind of leads me to architecture questions - would it be better to just leave the "master" field out of the struct, make it a method that finds a member of the Vec w/ a master flag set to true?, so something like: cluster.master() would basicaly implement cluster.members.iter().find(|m| m.master)  - where m.master would be  bool?  


actually, I have lots of questions about my code - lots of choices where I thought, "i have no idea, so I'll just try this and see." and when it compiled, I was happy. But, I'd like it to be "right", or, as I've noticed in Rust circles, "correct". Any help is much appreciated. Also, it doesn't actually do anything yet - it just prints out the data structures.
## [8][Simd Library Plans](https://www.reddit.com/r/rust/comments/gfyta3/simd_library_plans/)
- url: https://vorner.github.io/2020/05/08/simd-library-plans.html
---

## [9][serde-select: Query serde documents where you don't know the schema at compiletime](https://www.reddit.com/r/rust/comments/ggcdef/serdeselect_query_serde_documents_where_you_dont/)
- url: https://beyermatthias.de/blog/2020/05/09/serde-select/
---

## [10][Is there any asynchronous multipart/form-data implementation for Rust?](https://www.reddit.com/r/rust/comments/ggaj0k/is_there_any_asynchronous_multipartformdata/)
- url: https://www.reddit.com/r/rust/comments/ggaj0k/is_there_any_asynchronous_multipartformdata/
---
I know other web frameworks e.g. **Rocket**, **actix-web** etc have their own version of multipart/form-data implementation. But, I am looking for something which **will take a stream of bytes and allow to read each field as a stream of bytes** and It should work on any asynchronous stream and should be pluggable into any Rust asynchronous environment.

I also checked out [https://docs.rs/multipart-async/0.0.2/multipart\_async/](https://docs.rs/multipart-async/0.0.2/multipart_async/), but, its dependencies are very old and not working with the new version of **futures::stream::Stream**.
## [11][Read until EOF but not more then N bytes](https://www.reddit.com/r/rust/comments/ggdfrs/read_until_eof_but_not_more_then_n_bytes/)
- url: https://www.reddit.com/r/rust/comments/ggdfrs/read_until_eof_but_not_more_then_n_bytes/
---
Hello there,

I want to read from a source that implements std::io::Read trait until EOF but limiting to N bytes. There are helpful methods like read\_exact() and read\_to\_end(), but I has nothing found to limit the amount of bytes. In case there are more bytes than N, it must return an error.

Before I start to write a small helper function on my own, I want to ask if there is already a good solution available by the standard library.
## [12][Rust verification tools](https://www.reddit.com/r/rust/comments/gfz4gh/rust_verification_tools/)
- url: https://alastairreid.github.io/rust-verification-tools/
---

