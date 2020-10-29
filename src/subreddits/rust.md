# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (44/2020)!](https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (44/2020)?](https://www.reddit.com/r/rust/comments/ji8vgk/whats_everyone_working_on_this_week_442020/)
- url: https://www.reddit.com/r/rust/comments/ji8vgk/whats_everyone_working_on_this_week_442020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-44-2020/50643?u=llogiq)!
## [3][For Complex Applications, Rust is as Productive as Kotlin](https://www.reddit.com/r/rust/comments/jk8648/for_complex_applications_rust_is_as_productive_as/)
- url: https://ferrous-systems.com/blog/rust-as-productive-as-kotlin/
---

## [4][This Week in Rust 362](https://www.reddit.com/r/rust/comments/jk35ha/this_week_in_rust_362/)
- url: https://this-week-in-rust.org/blog/2020/10/28/this-week-in-rust-362/
---

## [5][rust-script - run Rust files and expressions as scripts](https://www.reddit.com/r/rust/comments/jjnyv1/rustscript_run_rust_files_and_expressions_as/)
- url: https://rust-script.org/
---

## [6][ELI5: Why do so less people use (A)GPLv3 for their crates?](https://www.reddit.com/r/rust/comments/jk5g20/eli5_why_do_so_less_people_use_agplv3_for_their/)
- url: https://www.reddit.com/r/rust/comments/jk5g20/eli5_why_do_so_less_people_use_agplv3_for_their/
---
I mainly see MIT and Apache licenses around, is (A)GPLv3 too complicated or what is the matter?

EDIT: For those who are interested, this is a detailed explanation what (A/L)GPL can and can't do: https://copyleft.org/guide/monolithic/

EDIT2: If you need help searching a license for your project you can use this site: https://choosealicense.com/
## [7][AdvancedResearch Blog: Monotonic-Solver 0.5 Released (A Rust library for Generic Automated Theorem Proving)](https://www.reddit.com/r/rust/comments/jjxye8/advancedresearch_blog_monotonicsolver_05_released/)
- url: https://advancedresearch.github.io/blog/2020-10-28-monotonic-solver-0.5-released
---

## [8][Simple remote access tool written in Rust with UDP](https://www.reddit.com/r/rust/comments/jk8n9j/simple_remote_access_tool_written_in_rust_with_udp/)
- url: https://github.com/Wafelack/remote_access_tool
---

## [9][to_trait, a crate that lets you write `thing.to::&lt;T&gt;()` instead of `Into&lt;T&gt;::into(thing)`](https://www.reddit.com/r/rust/comments/jjsjsh/to_trait_a_crate_that_lets_you_write_thingtot/)
- url: https://crates.io/crates/to_trait
---

## [10][Need help on structuring my project](https://www.reddit.com/r/rust/comments/jk69r6/need_help_on_structuring_my_project/)
- url: https://www.reddit.com/r/rust/comments/jk69r6/need_help_on_structuring_my_project/
---
Hi everyone,

I'm dabbling in Rust and have a question regarding a two-crate workspace project of mine. First, I have a library L which defines a few types and some functionality for it. Second, an application A which imports things from L, notably its structs.

**The problem**: Eventually, A wants to insert instances of L's types into a database, which I'm using diesel for. Since the original types in L do not derive diesel's macros, I cannot use them with a database.

Now, is this already some kind of code smell? I don't really want to add derives to structs in L without actually needing it. But I also don't want to have to map between L's structs and A's domain types and redefine everything. The reason for keeping L a separate library is to use it in future sub-applications within the same workspace.

What do you think is a sensible solution here? I feel like the most pragmatic would be to just derive the required diesel macros on those types. But, maybe there is something I'm not seeing. Thank you very much for your help.
## [11][IntelliJ Rust: New Functionality for Cargo Features](https://www.reddit.com/r/rust/comments/jka6bn/intellij_rust_new_functionality_for_cargo_features/)
- url: https://blog.jetbrains.com/clion/2020/10/intellij-rust-new-functionality-for-cargo-features/
---

## [12][using rustc suggestions to fix all errors](https://www.reddit.com/r/rust/comments/jk9guk/using_rustc_suggestions_to_fix_all_errors/)
- url: https://www.reddit.com/r/rust/comments/jk9guk/using_rustc_suggestions_to_fix_all_errors/
---
Hi,   
I had a "prelude" file that contained public uses of all types in my crate. However, I want to remove all \`use crate::prelude::\*\` statements in favor of the actual use statements per type.  


When I remove the prelude mod, the compiler obviously throws a lot of "cannot find type \* " errors. However it also suggests where to find them. Is there a way to fix those error similar to how \`cargo fix\` can be used to fix warnings?  


I don't want to fix those 700 errors by hand :D  


Cheers
