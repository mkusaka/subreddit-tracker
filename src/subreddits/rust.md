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
## [2][What's everyone working on this week (11/2020)?](https://www.reddit.com/r/rust/comments/fg0q9l/whats_everyone_working_on_this_week_112020/)
- url: https://www.reddit.com/r/rust/comments/fg0q9l/whats_everyone_working_on_this_week_112020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-11-2020/39269?u=llogiq)!
## [3][For your consideration: An alternative to the builder pattern, Init Struct Pattern!](https://www.reddit.com/r/rust/comments/fg6vrn/for_your_consideration_an_alternative_to_the/)
- url: https://xaeroxe.github.io/init-struct-pattern/
---

## [4][Cloudsmith releases improved Cargo support](https://www.reddit.com/r/rust/comments/fgbr71/cloudsmith_releases_improved_cargo_support/)
- url: https://blog.cloudsmith.io/2020/03/09/announcing-native-cargo-uploads-in-cloudsmith/
---

## [5][Blog post: A C# programmer examines Rust](https://www.reddit.com/r/rust/comments/fg5ble/blog_post_a_c_programmer_examines_rust/)
- url: https://treit.github.io/programming,/rust,/c%23/2020/03/06/StartingRust.html
---

## [6][Announcing rust-search-extension v0.8! The ultimate search extension for Rust.](https://www.reddit.com/r/rust/comments/fg5wo1/announcing_rustsearchextension_v08_the_ultimate/)
- url: https://www.reddit.com/r/rust/comments/fg5wo1/announcing_rustsearchextension_v08_the_ultimate/
---
Hi all!

Near a month ago, I announced **Rust Search Extension** v0.7 here and got more than 200 upvotes! Thank you all!! (In case someone missed it, here is the [original post](https://www.reddit.com/r/rust/comments/eymfxu/announcing_rustsearchextension_v07_search_std/).) 

Today, I'm so happy to announce v0.8 and along with a new slogan: **The Ultimate Search Extension for Rust**! Obviously, we're hoping the extension brings ultimate capability, flexibility, and productivity that every Rustacean could benefit from. 

Here is the basic introduction, let me know what you think! Thank you again!

# All the major features:

* Support search std docs
* Support search builtin attributes
* Support search error codes
* Support search **top 20K** crates **^(\[new\])**
* Support search official Rust book chapters **^(\[new\])**
* [Docs.rs](http://docs.rs/) enhancement (display **Feature flags**) **^(\[new\])**
* Command system **^(\[new\])**
* Support page down/up easily ^(\[new\])
* Support open in a new tab by using `Cmd`/`Alt` \+ `Enter` ^(\[new\])

# Support search top 20K crates

We bump the crates index size from top 10K to **top 20K** (more than 35k in total on [crates.io](http://crates.io/)) to help you search crates instantly. You can open the popup page to sync the latest crates index automatically.

What more, you can prefix an `!` (exclamation mark) before the keyword to search crates exclusively, prefix `!!` (double exclamation mark) to open [docs.rs](http://docs.rs/) page instead [crates.io](http://crates.io/) page!

[Prefix !! to open docs.rs page](https://preview.redd.it/m4wdifquzql41.png?width=2058&amp;format=png&amp;auto=webp&amp;s=fd2df82d5a53540d170db31099aadd63de209459)

# Support search official Rust book chapters

You can prefix `%` (percent sign) before the keyword to search all of Rust book chapters on your address bar! The result will show the title of the related page, parent chapter, and grandparent chapter.

[Search official Rust books](https://preview.redd.it/cxymbmi6yql41.png?width=2062&amp;format=png&amp;auto=webp&amp;s=b0af3d42b0e9346229b2b5375e7cda22c0222911)

# [Docs.rs](http://docs.rs/) enhancement (display Feature flags)

[Docs.rs](http://docs.rs/) don't display the crate's feature flags, which is inconvenient. So we fill the gap.

[Display Feature flags on docs.rs](https://preview.redd.it/h0ga65dnyql41.png?width=2068&amp;format=png&amp;auto=webp&amp;s=6a2dd1edaff63a64952f3064b533de54341adf23)

# Command system

The command system brings a handy set of useful and convenient commands to you. Each command starts with a `:` (colon), followed by the name, and function differently in individual. Those commands including but not limited to:

* `:help`  \- Show the help messages.
* `:yet` \- Show all Are We Yet websites.
* `:book` \- Show all Rust official books.
* `:stable` \- Show stable Rust scheduled release date in the next year.
* `:history` \- Show your local search history

[:yet command](https://preview.redd.it/mtmo5xmzyql41.png?width=2066&amp;format=png&amp;auto=webp&amp;s=0cdb78594e1862f3308d3359fad24148908d0b25)

[:stable command](https://preview.redd.it/b0hmp243zql41.png?width=2058&amp;format=png&amp;auto=webp&amp;s=449f833c43b0539d9498dbae63a38e53219a048e)

# Page down/up easily

You can press `space` after the keyword, then increase or decrease the number of `-` (hyphen) to page down or page up.

[Page down\/up](https://preview.redd.it/0ximu499zql41.png?width=2064&amp;format=png&amp;auto=webp&amp;s=59403cf995a6b0ef82bd89bcc69eb6e396c36538)

The official website [https://rust-search-extension.now.sh](https://rust-search-extension.now.sh/) gives you more details and tips about the extension, welcome to check it out.

## Github

**Rust Search Extension** is open-sourced on Github ([https://github.com/Folyd/rust-search-extension](https://github.com/Folyd/rust-search-extension)), welcome to file issue or PR!

# Installation

* [Chrome Web Store](https://chrome.google.com/webstore/detail/rust-search-extension/ennpfpdlaclocpomkiablnmbppdnlhoh)
* [Firefox](https://addons.mozilla.org/en-US/firefox/addon/rust-search-extension/)
## [7][rust-analyzer Changelog #15](https://www.reddit.com/r/rust/comments/ffvszw/rustanalyzer_changelog_15/)
- url: https://rust-analyzer.github.io/thisweek/2020/03/09/changelog-15.html
---

## [8][2020 Energy Efficiency across Programming Languages](https://www.reddit.com/r/rust/comments/ffyy31/2020_energy_efficiency_across_programming/)
- url: https://sites.google.com/view/energy-efficiency-languages/updated-functional-results-2020
---

## [9][NATS-WebUI: A web app for NATS Server written with Warp](https://www.reddit.com/r/rust/comments/fgdsyu/natswebui_a_web_app_for_nats_server_written_with/)
- url: https://github.com/sphqxe/NATS-WebUI
---

## [10][I just now wrote my first Hello World program using Rust. Excited to be a part of this community...](https://www.reddit.com/r/rust/comments/fgb4wj/i_just_now_wrote_my_first_hello_world_program/)
- url: https://www.reddit.com/r/rust/comments/fgb4wj/i_just_now_wrote_my_first_hello_world_program/
---
Suggest some beginner-friendly projects too. :)
## [11][Why are error types so rarely cloneable?](https://www.reddit.com/r/rust/comments/fg5mf7/why_are_error_types_so_rarely_cloneable/)
- url: https://www.reddit.com/r/rust/comments/fg5mf7/why_are_error_types_so_rarely_cloneable/
---
Every time I run into this, like io::Error not being cloneable, I'm confused. Why is this the case? How are people handling errors in an async/threaded environment? I need Clone on basically all of my generics, and so I end up wrapping errors with Arc.
## [12][Returning "Map" as IntoIter for IntoIterator](https://www.reddit.com/r/rust/comments/fgdp1y/returning_map_as_intoiter_for_intoiterator/)
- url: https://www.reddit.com/r/rust/comments/fgdp1y/returning_map_as_intoiter_for_intoiterator/
---
I have a struct with a `Vec` of `nalgebra::Vector3&lt;f32&gt;`. I need to implement IntoIterator for a mutable reference to this struct with the `Item` type being `&amp;mut f32` (aka, flatten into an iterator over `f32`. This can be done pretty easily using flatten and map, but I am not sure how to specify the return type for the map part of the iterator.

From my understanding, each function has a unique type. Is there any way to access this specific type (aka, the function returns a `Map` iterator using this exact function)?

I can solve this problem with boxing, but I am worried this will lead the iterator to not be inlined and it is used in performance sensitive code. Will llvm still be able to inline?

Alternatively, I would be happy to solve this problem using [the type_alias_impl_trait feature](https://github.com/rust-lang/rust/issues/63063),
but this leads to lifetimes error I think is a compiler bug.

I have all three versions of this code in a playground [here](https://play.rust-lang.org/?version=nightly&amp;mode=debug&amp;edition=2018&amp;gist=0850edff23378df08da6f4811d0c4b7d)
