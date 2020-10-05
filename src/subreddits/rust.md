# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (41/2020)!](https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/j57z68/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j1jgum/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (41/2020)?](https://www.reddit.com/r/rust/comments/j57zzs/whats_everyone_working_on_this_week_412020/)
- url: https://www.reddit.com/r/rust/comments/j57zzs/whats_everyone_working_on_this_week_412020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-41-2020/49684?u=llogiq)!
## [3][Making APIs more discoverable in Rustdoc](https://www.reddit.com/r/rust/comments/j5edk8/making_apis_more_discoverable_in_rustdoc/)
- url: https://www.reddit.com/r/rust/comments/j5edk8/making_apis_more_discoverable_in_rustdoc/
---
Hi everyone, I'm interested in improvements to Rustdoc that can help make an API more discoverable. Rustdoc is great if you know what you're looking for -- like "what's the return type of Vec::push?" -- but less good for less specific questions like "how do I append an element to a Vec?" More broadly, I'd like to answer two questions:

1. What kinds of discoverability questions do Rust developers ask about APIs?
2. How can Rustdoc provide better support for answering those questions?

For example, I like to be able to skim the "vocabulary" of an API, i.e. get a sense of the breadth of individual operations. But I find that particularly for large APIs (e.g. Vec which has 100+ methods including trait impls), the docs are hard to skim through even when collapsed. I've been prototyping a more dense visualization of methods, e.g. showing just the name and argument/return types grouped by mutability of self:

https://preview.redd.it/64vt9xapy6r51.png?width=1988&amp;format=png&amp;auto=webp&amp;s=e0fd28d67f95716e12314a57f8b0ea625172f79b

My question to you: what sorts of discoverability questions do you have about Rust APIs? Examples of times you struggled with documentation, or ideas for better visualizations of API information.
## [4][eventually: Event-sourcing in Rust!](https://www.reddit.com/r/rust/comments/j571fr/eventually_eventsourcing_in_rust/)
- url: https://github.com/ar3s3ru/eventually-rs
---

## [5][OS made with Rust](https://www.reddit.com/r/rust/comments/j5c5t2/os_made_with_rust/)
- url: https://www.reddit.com/r/rust/comments/j5c5t2/os_made_with_rust/
---
For the past few months I've been trying to make an operating system. My first attempt at this didn't go well and ended up being a GUI replacement for Linux Ubuntu. For the first half of my second attempt at making an OS I was just copying C code not knowing what it did or how to use it, then I heard about Rust and how it's becoming very popular. 

I've made a lot of progress since I switched to Rust for this project. Here is the link to the project on github,  [FOMOSv2-CL](https://github.com/NathanMcMillan54/FOMOSv2-CL)
## [6][Rust 2021: Foundation](https://www.reddit.com/r/rust/comments/j5fny9/rust_2021_foundation/)
- url: http://gilescope.ninja/rust/2020/10/05/Rust-2021-Foundation.html
---

## [7][Make a Language in Rust, Part Five: Binding Usages](https://www.reddit.com/r/rust/comments/j5f5tc/make_a_language_in_rust_part_five_binding_usages/)
- url: https://arzg.github.io/lang/5/
---

## [8][RSLint - an extremely fast and user friendly JavaScript linter written from scratch in rust.](https://www.reddit.com/r/rust/comments/j4z1ud/rslint_an_extremely_fast_and_user_friendly/)
- url: https://www.reddit.com/r/rust/comments/j4z1ud/rslint_an_extremely_fast_and_user_friendly/
---
Hello everyone!

I recently released v0.1.0 and v0.1.1 of a project i have been working on for a couple of months. The project is a fully fledged JavaScript linter but written from scratch in rust. I felt existing linters were too limiting and slow and decided to try and make one myself with some key points in mind:

# User facing differences
- Speed: 
  - File loading is parallelized
  - File linting is parallelized
  - Rule running is parallelized
  - (WIP) Files are incrementally reparsed for file watching
- Errors are vey rustc-like and friendly, they include labels, notes, etc
- Config uses TOML (but may eventually support json too)
- Rules are distinctly grouped (e.g. `errors`, `style`), allowing you to enable all error group rules and individually allow them using the `allowed = []` key
- Ignore commands (commands for the linter through comments) are either scoped to the entire file or only scoped to the specific statement/declaration they are on
- No need to worry about source code type (script vs module) and ECMAScript version, linter assumes script for `.js` and module for `.mjs` and assumes latest syntax. No config is needed at all.
- Error recovery, the parser and linter are 100% error tolerant, any source code can be linted no matter how wrong it is (⚠️ this is however not 100% the case currently since many recoveries still end in infinite recursion, if you find one then please submit a bug report)
- More powerful directives (commands for the linter through comments), directives are parsed using a distinct parser which reuses the js lexer (rslint_lexer) and can house multiple commands.
- CLI rule explanations with lexical syntax highlighting with rslint_lexer
- Further rule examples are autogenerated from tests

## Not yet implemented 
- Incremental reparsing (see [#16](https://github.com/RDambrosio016/RSLint/pull/16))
- TypeScript parsing and integration

# Technical differences
- rslint is a workspace and each crate has distinct jobs, rslint_cli has cli logic and a lot of dependencies for cool stuff, rslint_lexer is the js lexer and highlighter, rslint_parser is the parser, rslint_core is the core linter logic which does not know about the CLI, rslint_syntax is a simple shared crate among rslint_lexer and rslint_parser, and xtask is codegen and docgen glue.
- rslint uses a very unique syntax tree implementation (if you work on rust analyzer you will be familiar with the concept since its taken from rowan/rust analyzer), the tree is 100% lossless and error tolerant, stores tokens, and AST nodes are a thin wrapper on top of untyped nodes, this enables extremely powerful rule logic, for more reasons why this is significant you should look at my [blog post](https://rdambrosio016.github.io/rust/2020/09/18/pure-ast-based-linting-sucks.html) on the subject
- rules are grouped, there are not 200+ rules in a single directory.
- rules use a declare_lint macro which allows for easy declaration of rules, you write documentation as rust doc comments and a docgen tool can then transform it into user facing docs
- CST rules can check individual nodes without a need for a visitor, or can check individual tokens or even the root node.
- rules must all be Send + Sync.
- rule config is done directly through the rule structs using typetag

## Not yet implemented 
- Distinct types of rules, CST rules (implemented) run on individual concrete syntax trees/files, and LateRules run on all of the files being linted.
- (WIP) Rich fixer interfaces for autofixing rules efficiently using incremental reparsing. 

# Currently known big issues
- Optional chaining is not parsed correctly
- Non-simple expressions in key value patterns are not parsed correctly
- Multiline labels are sometimes spammy (this is a codespan-reporting issue which is fixed in the unreleased version of codespan)

If you would like to try the linter you can directly install the latest published version using cargo install:
```sh
cargo install rslint_cli
rslint_cli ./some/glob/pattern
```
or you can also clone the project and build it, which will include the most up to date bug fixes and changes. 

⚠️ note however that the linter is still in early development, so there will be bugs

https://github.com/RDambrosio016/RSLint

If you have any feedback i would love to talk about it!
## [9][Building a Recipe Manager - Part 1](https://www.reddit.com/r/rust/comments/j5j5jv/building_a_recipe_manager_part_1/)
- url: https://bheisler.github.io/post/recipe-manager-part-1/
---

## [10][How do I handle bad unicode in Rust?](https://www.reddit.com/r/rust/comments/j5b24l/how_do_i_handle_bad_unicode_in_rust/)
- url: https://www.reddit.com/r/rust/comments/j5b24l/how_do_i_handle_bad_unicode_in_rust/
---
For example, users will provide the text on the left, but the text on the right is actually the proper text. Also these are just examples of the problem in general collected from around the web and not a sample of the user provided data.

    uÌˆnicode -&gt; ünicode
    Ð¿ÑÐ°Ð²Ð¸Ð»ÑÐ½Ð¾ÐµÐ¿Ð¸ÑÐ°Ð½Ð¸Ðµ -&gt; правильноепитание
    â¨ââ¤Ãâ¨ââ¨ââ¨ââ¨ââ¤Ã®â¨ââ¨ââ¨â¡â¨ââ¨ââ¤Ã©â¨ââ¨ââ¨ââ¨â¡ -&gt; правильноепитание
    !Ãƒâ€šÃ‚Â¡!HONDA POW -&gt; !¡!HONDA POW

All of these are valid UTF8, but actually are mojibake from other encodings. I know of [Python's FTFY](https://ftfy.readthedocs.io/en/latest/) 
module which does all this stuff automatically, but is there anything within the Rust space that covers this use case?
## [11][Easy beginner Rust projects to contribute to as an experienced C++ engineer](https://www.reddit.com/r/rust/comments/j4wi0m/easy_beginner_rust_projects_to_contribute_to_as/)
- url: https://www.reddit.com/r/rust/comments/j4wi0m/easy_beginner_rust_projects_to_contribute_to_as/
---
Hi there,
I have been working with C++ for about 5 years and now I would like to learn Rust. What projects can you recommend me to contribute to in order to get in quickly into Rust?
Thanks
## [12][Serve only one HTTP request](https://www.reddit.com/r/rust/comments/j5iqjj/serve_only_one_http_request/)
- url: https://www.reddit.com/r/rust/comments/j5iqjj/serve_only_one_http_request/
---
For a project I need a function that listen on a specific port and serve one and only one HTTP request. I only need to extract a query parameter and make a dummy response.

I tried several libraries, like Hyper or Warp, but didn't find a nice way to do this simple task (and not end up with a full HTTP server).

What I did was to do it by myself with Tokio, I listen to a connection, accept it and handle the parsing / response manually.

Is there a nicer way to do it ?
