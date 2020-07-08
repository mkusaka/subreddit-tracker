# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (28/2020)!](https://www.reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hhv4z1/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (28/2020)?](https://www.reddit.com/r/rust/comments/hm1qi3/whats_everyone_working_on_this_week_282020/)
- url: https://www.reddit.com/r/rust/comments/hm1qi3/whats_everyone_working_on_this_week_282020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-28-2020/45440?u=llogiq)!
## [3][Rust is the only language that gets `await` syntax right](https://www.reddit.com/r/rust/comments/hnbz78/rust_is_the_only_language_that_gets_await_syntax/)
- url: https://www.reddit.com/r/rust/comments/hnbz78/rust_is_the_only_language_that_gets_await_syntax/
---
At first I was weirded out when the familiar `await foo` syntax [got replaced](https://boats.gitlab.io/blog/post/await-decision-ii/) by `foo.await`, but after working with other languages, I've come round and wholeheartedly agree with this decision. Chaining is just much more natural! And this is without even taking `?` into account:

C#: `(await fetchResults()).map(resultToString).join('\n')`

JavaScript: `(await fetchResults()).map(resultToString).join('\n')`

Rust: `fetchResults().await.map(resultToString).join('\n')`

It may not be apparent in this small example, but the absence of extra parentheses really helps readability if there are long argument lists or the chain is broken over multiple lines. It also plain makes sense because all actions are executed in left to right order.

I love that the Rust language designers think things through and are willing to break with established tradition if it makes things truly better. And the solid versioning/deprecation policy helps to do this with the least amount of pain for users. That's all I wanted to say!

More references:

* [Async-await status report: The syntax question](https://smallcultfollowing.com/babysteps/blog/2019/03/01/async-await-status-report/#the-syntax-question)
* [Making progress in await syntax](https://boats.gitlab.io/blog/post/await-syntax/)
* [Update on await syntax](https://boats.gitlab.io/blog/post/await-decision-ii/)
* [A final proposal for await syntax](https://boats.gitlab.io/blog/post/await-decision/)
## [4][Compiling Rust on Pine Phone](https://www.reddit.com/r/rust/comments/hn96iw/compiling_rust_on_pine_phone/)
- url: https://i.redd.it/hbb3g3a3uj951.jpg
---

## [5][Where is the rust community allowed to talk about changes in the codebase now that PR's are getting closed for discussion and posts about the changes removed on reddit?](https://www.reddit.com/r/rust/comments/hnfnti/where_is_the_rust_community_allowed_to_talk_about/)
- url: https://www.reddit.com/r/rust/comments/hnfnti/where_is_the_rust_community_allowed_to_talk_about/
---
A certain PR about sequences of elements of night and day variety got closed down to community discussion and the corresponding reddit post has also been removed. The reddit post being a discussion on both the PR and the closing down of discussion in it.

*To be clear I do not want and am not attempting to discuss the content of the PR here.*

If both a PR gets closed down and reddit posts get deleted before the PR has even been merged / closed, how are we as a community supposed  to discuss changes related to the language? Or are we simply not expected to have a voice in these matters?

I agree that politics shouldn't be discussed here, but when a change to the codebase is made off the back of a political and not technical decision (political meaning more non-technical than actually political), their needs to be a way to still discuss it. Closing down everything gives me an uneasy feeling regardless of if the PR is good or bad.

For reference: [https://www.reddit.com/r/rust/comments/hneczb/rust\_team\_is\_going\_to\_replace\_whitelist\_with/](https://www.reddit.com/r/rust/comments/hneczb/rust_team_is_going_to_replace_whitelist_with/) (which in my opinion was a mostly respectful discussion)
## [6][Approaching Rust more slowly seems to help with the learning curve](https://www.reddit.com/r/rust/comments/hneel7/approaching_rust_more_slowly_seems_to_help_with/)
- url: https://www.reddit.com/r/rust/comments/hneel7/approaching_rust_more_slowly_seems_to_help_with/
---
Just wanted to share the following beginner experience: When I started out I had a failed attempt at learning it. Looking back I think this was because I was porting old complex data structures from other software I've written. What got in the way the first time seems to be this pattern I had of learning new languages by porting graph data structures I already knew well from C#, TypeScript etc. I did this with Go and Kotlin, and that was fine, but not Rust 🤦‍♂️

This time I stared out much more [from scratch](https://youtu.be/_1TXKRYR1nA), and [it feels much better](https://clips.twitch.tv/GiantSlipperyWitchRickroll). The idea now to simply write some general programs to learn the language features. 

Rust is starting to feel really good now, and it often "clicks" for me. Perhaps it can be an advice for others just starting out too (: Taking it slow and not hitting the ground running might be wiser in this language than many others. I suspect this is the root cause of the learning curve woes we often hear about.
## [7][Concurrency Patterns in Embedded Rust](https://www.reddit.com/r/rust/comments/hmy0st/concurrency_patterns_in_embedded_rust/)
- url: https://ferrous-systems.com/blog/embedded-concurrency-patterns/
---

## [8][Getting a sparse, weighted, directed graph into Rust](https://www.reddit.com/r/rust/comments/hnd80m/getting_a_sparse_weighted_directed_graph_into_rust/)
- url: https://www.reddit.com/r/rust/comments/hnd80m/getting_a_sparse_weighted_directed_graph_into_rust/
---
I have written most of a simulation in Rust. The simulation should operate on a static, weighted, directed graph with around 17000 nodes, each with an out-degree around 6. The simulation wants to make use of shortest paths up to a threshold length and of a small number of node attributes.

I am currently generating this graph by putting data into an SQLite database using Python. But how do I get it into Rust and use it from there? There will be a gazillion queries from the simulation to the graph asking “which nodes of distance up to 1'000'000 are what distance from this focus node?” and “What is the node attribute vector for this focus node?”. Getting these queries resolved as quickly as possible is more important than a small memory footprint, but being able to run it on my laptop with its 6GB RAM for testing purposes would still be nice.

Do I want a database? Which one? With pre-computed distances up to 1'000'000 units, or is it wiser to use Dijkstra to compute them on the fly because if I need the 1'000'000 distance node, I also need the 500 distance nodes anyway? Or do I want a graph crate (which one?) that reads some standardized graph format and stuffs the whole object into memory in a nice format that can be queried by maximum distance?
## [9][io_uring-inspired async operations in Redox OS](https://www.reddit.com/r/rust/comments/hngetq/io_uringinspired_async_operations_in_redox_os/)
- url: https://www.redox-os.org/news/io_uring-1/
---

## [10][Getting in and out of trouble with Rust futures](https://www.reddit.com/r/rust/comments/hn3i45/getting_in_and_out_of_trouble_with_rust_futures/)
- url: https://fasterthanli.me/articles/getting-in-and-out-of-trouble-with-rust-futures
---

## [11][This Month in Rust GameDev #11 - June 2020](https://www.reddit.com/r/rust/comments/hn0qso/this_month_in_rust_gamedev_11_june_2020/)
- url: https://rust-gamedev.github.io/posts/newsletter-011
---

## [12][Growing Gold with Rust](https://www.reddit.com/r/rust/comments/hmqjvs/growing_gold_with_rust/)
- url: https://www.reddit.com/r/rust/comments/hmqjvs/growing_gold_with_rust/
---
Hi everyone, 

I’m a scientist working in the field of nano-optics/-technology and at work we are regularly growing gold flakes -- thin platelets out of single-crystalline gold. How it works is still a bit of magic and a better understanding would be great.

[Some real gold flakes.](https://preview.redd.it/7iil1whf8e951.png?width=800&amp;format=png&amp;auto=webp&amp;s=1b7d211abc523654aa73e1329fd2f3d78b5c1e33)

As I got interested in Rust and had a lot of time during the last couple of months, I started implementing a simulation of the growing process in Rust. And it was a pleasure!

The main challenges were to find the right data structures for handling up to billions of atoms (in the end I used ndarray with some bitshifting/masking) at an acceptable speed (I settled with BTreeSet as a store for the surface vacancies) and also to deal with the fcc lattice and its arbitrary number of stacking faults of the gold crystals. I learned a lot about proper programming and important system details e.g. that the stack is limited to only a couple of megabytes or that the OS is really lazy when allocating memory. I wasn’t able to get DTrace running on Windows, so no nice flame graphs here, but I believe the bottleneck is that there is no quick way to randomly pick an element from the BTreeSet. At least I didn’t find one. If you have an idea, please tell me!

The code can be found at [Github](https://github.com/Rene-007/flake_growth), a [Wasm version](https://www.kullock.de/flake_growth/) is also available to get a first impression and I would recommend everybody to first have a look at the [visual guide](https://www.kullock.de/flake-growth-guide/).

&amp;#x200B;

[A short animation.](https://i.redd.it/stgdksuk8e951.gif)

In the end I have come to like the language, the tooling and the values behind Rust a lot! I think it is really the way to go forward for not only systems programming. Nevertheless, I had some difficulties and got some ideas I would like to share with you:

* Initially it was quite difficult to get started because I missed the ability to play around with the data and e.g. see if my mapping from memory to the fcc lattice and back is right. It is not possible to write unit tests for that, and one must simply play around and see whether the results in the 3D scene fit or not. At work I usually use Matlab for things like that and it would be the much easier tool to figure out the mapping, but I deliberately decided not to do so. (During my PhD time I used python/numpy a lot and observed that it is (mentally) hard to switch to another language once you already invested a lot into your code. And when looking on examples such as the HipHop Virtual Machine others seems have the problem, too.) So, from my point of view it would be nice to have some “playing-around capabilities”. I think I do not mean rapid prototyping with that but rather some small loop/block you can put into your program where the compiled code is interrupted by an interpreted section. When you run the program, you will end in the interpreter loop, have direct access to all the data structure/functions and can play around with them using the Rust syntax. So, no bindings or another language needed. I am not sure if that is feasible, but it would be cool.
* The second issue circles around libraries. As a beginner/outsider it is hard to judge which library/crate is needed, which might be the best one, which one I can trust. I can sympathize with the decision to have a small and stable std library but have the feeling that there should be some additional “meta crate” which combines the most popular matured crates in a complete way, i.e. that no additional external dependencies are needed. This crate should provide a root namespace (e.g. pop::rand or pop::serde), all unsafe parts should be reviewed  (with reports/discussion openly available) and some common programming standards (documentation, api design, naming) would also be good. It should be a big honor when “your crate” gets part of this “popular crate” or when you yourself become an approved reviewer. In contrast to the standard lib the API should guarantee compatibility only within a Rust edition such that subcrates, which are not state of the art anymore, can simply be removed. But as they will still live on as a separate crate, existing users just need the remove the “pop::” prefix within their source code. Smaller incompatible API changes within a subcrate might be introduced similar to linux in a “pop.next” meta-crate to smoothen the transition to the next edition. I think this might be a good compromise for a trustworthy base which is stable but not hammered in stone forever. What do you think?
