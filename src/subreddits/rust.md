# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (37/2020)!](https://www.reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ioc56i/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/ijvwsk/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 355](https://www.reddit.com/r/rust/comments/ippv0q/this_week_in_rust_355/)
- url: https://this-week-in-rust.org/blog/2020/09/09/this-week-in-rust-355/
---

## [3][Clone git faster than git itself with Gitoxide 0.4](https://www.reddit.com/r/rust/comments/ir8mwq/clone_git_faster_than_git_itself_with_gitoxide_04/)
- url: https://github.com/Byron/gitoxide/releases/tag/v0.4.0
---

## [4][Announcing ultraviolet 0.6 - including support for 256-bit AVX vectors, f64 double-precision types, and performance improvements](https://www.reddit.com/r/rust/comments/ir24mp/announcing_ultraviolet_06_including_support_for/)
- url: https://fusha.moe/blog/posts/ultraviolet-0.6/
---

## [5][Apple is starting to use Rust for low-level programming](https://www.reddit.com/r/rust/comments/iqnt58/apple_is_starting_to_use_rust_for_lowlevel/)
- url: https://twitter.com/oskargroth/status/1301502690409709568?s=10
---

## [6][Announcing Actix-Web v3.0](https://www.reddit.com/r/rust/comments/iqq8k9/announcing_actixweb_v30/)
- url: https://paper.dropbox.com/published/Announcing-Actix-Web-v3.0--A7YI~P9U9aqhEOXyZJaGffjfBg-QOXXb1lXgTubzXHzUq9ONY5
---

## [7][A Few Github Action “Recipes” for Rust](https://www.reddit.com/r/rust/comments/ir16fl/a_few_github_action_recipes_for_rust/)
- url: https://shift.click/blog/github-actions-rust
---

## [8][Trait bound for "One or More"](https://www.reddit.com/r/rust/comments/ir5ue1/trait_bound_for_one_or_more/)
- url: https://www.reddit.com/r/rust/comments/ir5ue1/trait_bound_for_one_or_more/
---
How can I express a generic type that allows calling a function with a value or a slice of values? Example:

\`\`\`

fn myfunc&lt;T: OneOrMore&gt;(values: T) {

...

mycontainre.from\_iter(values.into\_iter());

....

}

myfunc(1);

myfunc(&amp;\[1, 2\]);

\`\`\`

&amp;#x200B;

Is there something like this "OneOrMore" trait? or can it be accomplished in a different way?
## [9][Conway's Game of Life on 10,000 by 10,000 Grid](https://www.reddit.com/r/rust/comments/ir2j2p/conways_game_of_life_on_10000_by_10000_grid/)
- url: https://www.reddit.com/r/rust/comments/ir2j2p/conways_game_of_life_on_10000_by_10000_grid/
---
Hi everyone,

I have been learning Rust for the past few months and for my first project, I decided to implement Conway's Game of Life on a 10,000 by 10,000 grid, as suggested by this reddit post([https://www.reddit.com/r/learnprogramming/comments/9c6yo4/a\_collection\_of\_projects\_for\_learning/](https://www.reddit.com/r/learnprogramming/comments/9c6yo4/a_collection_of_projects_for_learning/)). You can use the arrow keys to move around the grid, as this grid size will not be able to fit onto your screen. And since the computation time for the next generation is at around 1.5 seconds right now, I made the arrow key movement asynchronous so it won't feel "laggy".

&amp;#x200B;

In any case, any tips on the code quality and/or speed optimizations would be much appreciated!

[https://github.com/bbli/fast\_game\_of\_life](https://github.com/bbli/fast_game_of_life)
## [10][Any Experienced C Programmers Who Switched to Rust?](https://www.reddit.com/r/rust/comments/iqrear/any_experienced_c_programmers_who_switched_to_rust/)
- url: https://www.reddit.com/r/rust/comments/iqrear/any_experienced_c_programmers_who_switched_to_rust/
---
As the title says, I’m really curious if anyone with serious C experience, e.g. 5+ years of professional use, has switched to Rust for problems they used to solve in C. I’m interested in getting a perspective from C++ people as well, but that’s a huge can of worms because of how much C++11 improved things, and the likelihood that the C++ code experienced programmers are dealing with uses a ton of pre-C++11 patterns.
## [11][Typed Arrays from String Arrays for Dataset Operation](https://www.reddit.com/r/rust/comments/irbw0w/typed_arrays_from_string_arrays_for_dataset/)
- url: https://datacrayon.com/posts/programming/rust-notebooks/typed-arrays-from-string-arrays-for-dataset-operation/
---

## [12][i2c communication](https://www.reddit.com/r/rust/comments/irasr6/i2c_communication/)
- url: https://www.reddit.com/r/rust/comments/irasr6/i2c_communication/
---
Hey, rustaceans! I am trying to use Rust for embedded programming and I have microbit v1.5   
I am really struggling with i2c accelerometer communication, does anyone have working examples on how to read data from accelerometer with i2c?
