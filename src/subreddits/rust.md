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
## [2][This Week in Rust 346](https://www.reddit.com/r/rust/comments/hnkws3/this_week_in_rust_346/)
- url: https://this-week-in-rust.org/blog/2020/07/08/this-week-in-rust-346/
---

## [3][Linux kernel in-tree Rust support](https://www.reddit.com/r/rust/comments/hp560p/linux_kernel_intree_rust_support/)
- url: https://lore.kernel.org/lkml/CAKwvOdmuYc8rW_H4aQG4DsJzho=F+djd68fp7mzmBp3-wY--Uw@mail.gmail.com/T/#u
---

## [4][What feature of Rust is used very often by experienced programmers, but not so much by a newbie?](https://www.reddit.com/r/rust/comments/hp6guz/what_feature_of_rust_is_used_very_often_by/)
- url: https://www.reddit.com/r/rust/comments/hp6guz/what_feature_of_rust_is_used_very_often_by/
---

## [5][Idiomatic custom number types](https://www.reddit.com/r/rust/comments/hp6dx4/idiomatic_custom_number_types/)
- url: https://www.reddit.com/r/rust/comments/hp6dx4/idiomatic_custom_number_types/
---
Hi all,

I've got a situation where it would be useful to have multiple customer number types so that it's clear that they mean different things. In my native Go, I would use separate types, which can't then interact ([playground](https://play.golang.org/p/J8nkgbDkgp5)).

As far as I can tell from the Rust book, I have two options, neither of which feels quite right. I can either use the [Newtype Pattern](https://doc.rust-lang.org/book/ch19-03-advanced-traits.html#using-the-newtype-pattern-to-implement-external-traits-on-external-types) to wrap a numerical type, but then using values of those types gets rather clumsy, as I can't do arithmetic without re-adding the operators ([playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=5da92bba6ff0da476948a61fa509518e)), or I can use [Type Aliases](https://doc.rust-lang.org/book/ch19-04-advanced-types.html#creating-type-synonyms-with-type-aliases) to define new number types, but then the compiler doesn't seem to distinguish between them, so it feels prone to type confusion bugs ([playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=fd3ab67f5c7139607d30b8ebbe0ed483)).

Is there another option I've missed? Is there a different, more idiomatic solution to the problem? I know I'm trying to recreate what I'd do in Go, which might not be the right approach in Rust.

Thanks!
## [6][cargo-todo : a simple rust tool](https://www.reddit.com/r/rust/comments/hos1um/cargotodo_a_simple_rust_tool/)
- url: https://www.reddit.com/r/rust/comments/hos1um/cargotodo_a_simple_rust_tool/
---
hi guys,

Just made [cargo-todo](https://crates.io/crates/cargo-todo), a really (i mean really) simple tool for cargo.

Just type in cargo todo inside your project directory and it will display every //todo lines inside the source code.

I'm open for feedback and requests.

I also made some other crates, feel free to check them right [here](https://crates.io/users/ProbablyClem)

edit : i'm kinda bored right know so feel free to give me any feedback or even projects to contribute to
## [7][Programming Servo: the bird’s-eyes view](https://www.reddit.com/r/rust/comments/hp4pde/programming_servo_the_birdseyes_view/)
- url: https://www.reddit.com/r/rust/comments/hp4pde/programming_servo_the_birdseyes_view/
---
A high-level overview of the architecture of the Servo Web engine, a large concurrent and multi-process system written in the Rust programming language. 

[https://medium.com/programming-servo/programming-servo-the-birds-eyes-view-201d28220b9a?source=friends\_link&amp;sk=b8610f254bf6faf8b81c81729c1b3498](https://medium.com/programming-servo/programming-servo-the-birds-eyes-view-201d28220b9a?source=friends_link&amp;sk=b8610f254bf6faf8b81c81729c1b3498)
## [8][What is Yew and Seed lacking when compared to things like React?](https://www.reddit.com/r/rust/comments/hp4p6q/what_is_yew_and_seed_lacking_when_compared_to/)
- url: https://www.reddit.com/r/rust/comments/hp4p6q/what_is_yew_and_seed_lacking_when_compared_to/
---
How is following

* Tooling, ex: Webpack, CRA
* Styling: ex: Styled components, Emotion
* Static site generation. ex: Gatsbyjs
* State management: ex: Redux
* Browser support with Polyfills
* Performance 
   *  Is DOM API problem solved?
   * Lazy loading possible?
   * How is bundle sizes and bundle splitting
## [9][mind: A productive mind](https://www.reddit.com/r/rust/comments/hp6yxx/mind_a_productive_mind/)
- url: https://www.reddit.com/r/rust/comments/hp6yxx/mind_a_productive_mind/
---
I've been working on [github.com/sayanarijit/mind](https://github.com/sayanarijit/mind), a super simple rust project to simplify keeping track of pending tasks. It's aimed at increasing productivity using the terminal any way possible.

This is not yet complete. In fact, I'm posting this to get feedback on what I should be working on next.

* Better error handling.
* More interactivity. (eliminate the need to press ENTER, fuzzy search, auto completion etc.)
* Sync (with drives, other todo/task manager apps e.g. todoist, any.do etc.)
* Reminders (so that you don't need to manually edit \`reminders.yml\`)
* Key shortcuts.
* \[mention your own suggestion\]
## [10][Target Feature vs Target CPU for Rust](https://www.reddit.com/r/rust/comments/hozs31/target_feature_vs_target_cpu_for_rust/)
- url: https://www.nickwilcox.com/blog/target_cpu_vs_target_feature/
---

## [11][Error chaining](https://www.reddit.com/r/rust/comments/hp92f6/error_chaining/)
- url: https://www.reddit.com/r/rust/comments/hp92f6/error_chaining/
---
I'm a newbie at Rust and I'm struggling to find a straight answer on how to deal with different error types in a `Result`. I'm coming from Scala where everything just gets boiled down to a single `Throwable` type for the most part.

Can I have my function return types be `Result&lt;T, impl Error&gt;` and chain them together that way?
## [12][Typeracer game you can play from Neovim live coded in Rust, the first alpha gameplay with an audience member!](https://www.reddit.com/r/rust/comments/hp7w3p/typeracer_game_you_can_play_from_neovim_live/)
- url: https://clips.twitch.tv/CrypticFastGaurLeeroyJenkins
---

