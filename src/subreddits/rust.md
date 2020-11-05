# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (45/2020)!](https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jmijzu/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ji8ukt/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (45/2020)?](https://www.reddit.com/r/rust/comments/jmiktw/whats_everyone_working_on_this_week_452020/)
- url: https://www.reddit.com/r/rust/comments/jmiktw/whats_everyone_working_on_this_week_452020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-45-2020/50962?u=llogiq)!
## [3][Making a Snake Clone with Bevy [Rewritten and Updated for Bevy 0.3]](https://www.reddit.com/r/rust/comments/job4jo/making_a_snake_clone_with_bevy_rewritten_and/)
- url: https://mbuffett.com/posts/bevy-snake-tutorial/#0.3
---

## [4][Sauron [v0.32] - a very fast web framework with server-side rendering.](https://www.reddit.com/r/rust/comments/jodrsw/sauron_v032_a_very_fast_web_framework_with/)
- url: https://www.reddit.com/r/rust/comments/jodrsw/sauron_v032_a_very_fast_web_framework_with/
---
[Project repo](https://github.com/ivanceras/sauron)

[Complete Changelog](https://github.com/ivanceras/sauron/blob/master/Changelog.md)

[WIP Guide](https://sauron-rs.github.io/)

`0.32` is the most stable iteration so far, adding a lot of improvements without breaking any of the api.

Notable improvements:
- you can now use rust keywords such as `for`,`type` in html node! macro.
   example: &lt;input type="text"/&gt;.

- `jss`, `jss_ns`, `class_namespaced` to write css styles alongside with your component's view, while mitigating conflicting classnames with other components by prepending the component's namespace. This is not foolproof replacement for scoped style, but it comes very handy replacement for compiled stylesheets such scss or less. Example usage is in [futuristic-ui button](https://github.com/ivanceras/sauron/blob/master/examples/futuristic-ui/src/fui_button.rs)


A lot of performance improvement is done under the hood of sauron's [mt-dom](https://github.com/ivanceras/mt-dom) (read as: "empty-dom", or "mount-doom" whichever you like). In it, I've added improvements on diffing  algorithm of keyed-elements with `replace` and `skip` directive to allow developers to skip diffing on a particular nodes, for responsive/performance sensitive applications such as a code-editor(skip diffing on unchanged lines of code).

A [wiki](https://github.com/ivanceras/sauron/wiki) pages is in the works (Contributions are welcomed).
The deployed [site](https://sauron-rs.github.io/) is static site forged and generated using sauron itself to produce a searchable static site with nice rendering of source code and copy-able fake terminal for cli instructions.

A big thank you to my patreon/opencollective supporters:
- Matthew Piziak
- Freeman Latif
- YVT
- Philip Edwards
- Hugo Buddel
## [5][MetaCall: Call Python, Ruby, NodeJS, TypeScript, C#, Cobol... functions from Rust (Looking for contributors)](https://www.reddit.com/r/rust/comments/jogthj/metacall_call_python_ruby_nodejs_typescript_c/)
- url: https://www.reddit.com/r/rust/comments/jogthj/metacall_call_python_ruby_nodejs_typescript_c/
---
Hey rustaceans, I want to announce a recently published crate ( [https://crates.io/crates/metacall](https://crates.io/crates/metacall) ). This is a Rust wrapper around MetaCall library which is a multi-language runtime library (or also known as Polyglot).

The original library provides virtually support for calling functions from/to any programming language. The crate is not production ready (although we are using the C/C++ library already in a FaaS) and also there is some feedback already gathered ( [https://github.com/metacall/core/issues/56](https://github.com/metacall/core/issues/56) ) which probably I won't be able to implement because I am a noob in Rust (I'm mainly a C/C++ developer).

This Rust Proof of Concept works but it is far away from mimicking other APIs like NodeJS Port:

    const { sum } = require('sum.py');
    sum(3, 4); // 7

Where `sum.py` is this:

    def sum(a, b):
      return a + b

The Rust equivalent proposal would be something like:

    let a = metacall! {
      sum(3, 4);
      foo()
    }

&amp;#x200B;

Any feedback about the library will be highly appreciated and added to the issue for implementing it in the future. Anybody interested in contributing can contact us in our Telegram chat ( [https://t.me/joinchat/BMSVbBatp0Vi4s5l4VgUgg](https://t.me/joinchat/BMSVbBatp0Vi4s5l4VgUgg) ), or direct PR are accepted too.
## [6][ThreadScope alternative for Rust? (Tool for performance profiling of parallel programs)](https://www.reddit.com/r/rust/comments/johsyq/threadscope_alternative_for_rust_tool_for/)
- url: https://www.reddit.com/r/rust/comments/johsyq/threadscope_alternative_for_rust_tool_for/
---
For Haskell programs, there is [ThreadScope](https://wiki.haskell.org/ThreadScope)   
I am looking to fine-tune a Rust service before taking it to production and I want to introduce some parallelism/concurrency wherever possible.   
I'm going to try and use \`rayon\` and/or \`tokio\` depending on whether the functions/operations are \`async\`  or not.   


Is there any tool like \`ThreadScope\` or something similar that would help me visualize better the utilization of my processor so that I can spot issues &amp; know what to optimize?
## [7][Why Dark didn't choose Rust](https://www.reddit.com/r/rust/comments/jo2hco/why_dark_didnt_choose_rust/)
- url: https://blog.darklang.com/why-dark-didnt-choose-rust/
---

## [8][Small script to sort derives](https://www.reddit.com/r/rust/comments/joh0z0/small_script_to_sort_derives/)
- url: https://www.reddit.com/r/rust/comments/joh0z0/small_script_to_sort_derives/
---
Have you ever been annoyed that `cargo fmt` only sorts imports but not derives ?  
Fear not! Consistency has arrived in a simple form:  
A small python script that searches through all Rust files in current directory and sort derives for you according to a defined order when possible.  
If you wish to change the order or add your own derives, just change the `special` list at the top.  

[Link to gist](https://gist.github.com/Uriopass/542bfc3d489f98048c4c98e061782ac9)

/!\ Be sure to run this script in a git repository as changes are irreversible and I am not entirely sure what happens if the script is killed while running /!\  

I tested it on my 20k lines Rust project and it compiled successfully.
## [9][Ktra: Your Little Cargo Registry](https://www.reddit.com/r/rust/comments/jny8h6/ktra_your_little_cargo_registry/)
- url: https://github.com/moriturus/ktra
---

## [10][async-std v1.7.0 has been released!](https://www.reddit.com/r/rust/comments/jnueup/asyncstd_v170_has_been_released/)
- url: https://github.com/async-rs/async-std/releases/tag/v1.7.0
---

## [11][Introducing Text_Adventure an easy way to make a text adventure game.](https://www.reddit.com/r/rust/comments/jo8s4g/introducing_text_adventure_an_easy_way_to_make_a/)
- url: https://github.com/10maurycy10/text_adventure
---

## [12][With the launch of Bevy Game Engine 0.3, I'm introducing chunk base Tile Maps!](https://www.reddit.com/r/rust/comments/jntjp8/with_the_launch_of_bevy_game_engine_03_im/)
- url: https://www.reddit.com/r/rust/comments/jntjp8/with_the_launch_of_bevy_game_engine_03_im/
---
 [https://crates.io/crates/bevy\_tilemap](https://crates.io/crates/bevy_tilemap) 

This is not intended to be just another Tile Map. It is meant to be a framework and extensible by design, like Bevy. There is an emphasis placed on generic traits to accomplish that. As well as work done to keep it as close to Bevy API as possible. This allows anyone to create their own tiles, chunks and maps and still retain the speed of a handcrafted multi-threaded chunk loader and tile map. 

If you have comments or suggestions do let me know below! Thanks!
