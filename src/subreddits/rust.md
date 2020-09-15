# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (38/2020)!](https://www.reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ismh8n/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (38/2020)?](https://www.reddit.com/r/rust/comments/ismi9w/whats_everyone_working_on_this_week_382020/)
- url: https://www.reddit.com/r/rust/comments/ismi9w/whats_everyone_working_on_this_week_382020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-38-2020/48763?u=llogiq)!
## [3][I've been working on adding WSL support to IntelliJ Rust for the past few months. Should make Windows and Rust development a bit nicer for everyone!](https://www.reddit.com/r/rust/comments/it82wo/ive_been_working_on_adding_wsl_support_to/)
- url: https://github.com/intellij-rust/intellij-rust/pull/5014
---

## [4][Oxidizing XDG portals with zbus](https://www.reddit.com/r/rust/comments/it2hij/oxidizing_xdg_portals_with_zbus/)
- url: https://belmoussaoui.com/article/13-oxidizing-portals
---

## [5][How to make VSCode reliable for editing Rust?](https://www.reddit.com/r/rust/comments/it4exz/how_to_make_vscode_reliable_for_editing_rust/)
- url: https://www.reddit.com/r/rust/comments/it4exz/how_to_make_vscode_reliable_for_editing_rust/
---
I was using VSCode, but recent updates broke important functionality:

* vscode-rust extension + rust language server don't work for tokio and other crates with optional features:
  https://github.com/rust-lang/vscode-rust/issues/637
* vscode-rust extension + rust-analyzer are broken:
  https://github.com/rust-lang/vscode-rust/issues/852

Is there any combination of versions of VSCode, VSCode Rust Extension, RLS/rust-analyzer, and rust toolchain that just work?

For now, I've gone back to IntelliJ.  It resolves all of the tokio symbols and shows compiler errors and warnings.
* IntelliJ IDEA CE 2020.1.4
* IntelliJ Rust plugin 0.3.130.3337-201
* rust 1.47.0-beta.3
## [6][A call for contributors from the WG-prioritization team](https://www.reddit.com/r/rust/comments/isvai5/a_call_for_contributors_from_the_wgprioritization/)
- url: https://blog.rust-lang.org/2020/09/14/wg-prio-call-for-contributors.html
---

## [7][Your Language Sucks, It Doesn’t Matter](https://www.reddit.com/r/rust/comments/ismzmm/your_language_sucks_it_doesnt_matter/)
- url: https://matklad.github.io//2020/09/13/your-language-sucks.html
---

## [8][rust warns of unused variables if not using a feature](https://www.reddit.com/r/rust/comments/it201w/rust_warns_of_unused_variables_if_not_using_a/)
- url: https://www.reddit.com/r/rust/comments/it201w/rust_warns_of_unused_variables_if_not_using_a/
---
    fn main(){
        let thing = 1;
        #[cfg(feature="print_thing")]
        println!("The thing is {}", thing);
    }

The above code warns about an unused variable when the `print_thing` feature is not enabled, but the variable is only sometimes unused. Is it possible to configure the compiler to only warn about unused variables if they are unused in all configs?
## [9][Just released: C++ for Visual Studio Code v1.0. Does it improve Rust debugging experience in VS Code?](https://www.reddit.com/r/rust/comments/it3sfy/just_released_c_for_visual_studio_code_v10_does/)
- url: https://devblogs.microsoft.com/cppblog/c-in-visual-studio-code-reaches-version-1-0/
---

## [10][rust-analyzer changelog #42 – now with async blocks and more!](https://www.reddit.com/r/rust/comments/isl3cy/rustanalyzer_changelog_42_now_with_async_blocks/)
- url: https://rust-analyzer.github.io/thisweek/2020/09/14/changelog-42.html
---

## [11][Learning resources for Rust patterns](https://www.reddit.com/r/rust/comments/isxxdm/learning_resources_for_rust_patterns/)
- url: https://www.reddit.com/r/rust/comments/isxxdm/learning_resources_for_rust_patterns/
---
Hey! I've been learning Rust for some months and I've been struggling to find some good resources to learn more about the patterns used in Rust like the builder pattern. I already read the Rust Programming Language Book but now I'm looking for something to read more related to patterns, good practices, etc.

Any recommendation of books, articles, blog posts, etc.?
## [12][I2C in rust](https://www.reddit.com/r/rust/comments/it8sj6/i2c_in_rust/)
- url: https://www.reddit.com/r/rust/comments/it8sj6/i2c_in_rust/
---
does anyone here know how to create code to receive and send any keyboard typed data to and from the I2C port on the Rpi in Rust, also using the arduino.
