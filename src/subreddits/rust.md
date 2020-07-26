# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (30/2020)!](https://www.reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 348](https://www.reddit.com/r/rust/comments/hvjf4i/this_week_in_rust_348/)
- url: https://this-week-in-rust.org/blog/2020/07/21/this-week-in-rust-348/
---

## [3][Per @ryan_levick: If you’re interested in making Windows applications with WinUI and XAML, let the team know by upvoting this issue!](https://www.reddit.com/r/rust/comments/hy65qk/per_ryan_levick_if_youre_interested_in_making/)
- url: https://github.com/microsoft/microsoft-ui-xaml/issues/2488
---

## [4][An introduction to Data Oriented Design with Rust](https://www.reddit.com/r/rust/comments/hxqwom/an_introduction_to_data_oriented_design_with_rust/)
- url: http://jamesmcm.github.io/blog/2020/07/25/intro-dod/#en
---

## [5][Simulate variadic generics](https://www.reddit.com/r/rust/comments/hy3v2s/simulate_variadic_generics/)
- url: https://www.reddit.com/r/rust/comments/hy3v2s/simulate_variadic_generics/
---
Hi all, I'm trying to implement a (toy) permission-based authorization system for Rocket and actix-web, and want to hear you opinions on how I could approach a specific part of the design.

My understanding is that in Rocket, [guards](https://rocket.rs/v0.4/guide/requests/#request-guards) are usually used to handle authentication and authorization. For authorization, how would you implement something that can be used like this:

    #[get("/my-protected-route")]
    fn handler(
      user: UserWithPermissions&lt;ReadPost, EditPost, ...&gt;,
      ...
    ) ...

or this:

    #[get("/my-protected-route")]
    fn handler(
      user: User,
      perm1: Perm1,
      perm2: Perm2,
      ...
    ) ...

But with the second form, `perm1` and `perm2` would depend on `user`, and I don't know how to do it.

Maybe this is the wrong approach. My high-level goal is to be able to specify a list of permissions that need to be present for a route. I'm curious how/if people have done it.
## [6][servo-unity: A browser plugin for Unity](https://www.reddit.com/r/rust/comments/hy09mm/servounity_a_browser_plugin_for_unity/)
- url: https://blog.mozvr.com/a-browser-plugin-for-unity/
---

## [7][Serious bug in Rust 1.45 stable](https://www.reddit.com/r/rust/comments/hxlj68/serious_bug_in_rust_145_stable/)
- url: https://www.reddit.com/r/rust/comments/hxlj68/serious_bug_in_rust_145_stable/
---

https://github.com/rust-lang/rust/issues/74739

It was found via a [stackoverflow question](https://stackoverflow.com/q/63087217/493729).

**Edit** tl;dr of the comments below: The bug is triggered only by very simplistic code, where all of the inputs are constant. Real-world code is therefore very unlikely to be affected. Each Rust release is tested with [crater](https://github.com/rust-lang/crater), which runs all tests for every crate on crates.io - and _**none**_ were affected. It got through because it's really not as bad as it looks.

The bug doesn't appear to be present in the most recently nightly, so it should be fixed quickly. ~~It's still a bit scary that a bug this serious could get past the tests.~~
## [8][Good reference for design patterns for multiple mutable references to data in the same container](https://www.reddit.com/r/rust/comments/hy52g1/good_reference_for_design_patterns_for_multiple/)
- url: https://www.reddit.com/r/rust/comments/hy52g1/good_reference_for_design_patterns_for_multiple/
---
I try to learn and use Rust in hobby projects and one particular problem I have noticed I always get stuck on is when I need to have two (or more) mutable references to data in the same container. Sometimes the two references are guaranteed to always be different (so the compiler should be able to check for this), but sometimes the two references may point to the same object (whether they do would only be known at runtime). It's an annoying problem, because this is trivial in other languages, but for some reason I can be stuck on trying to find an idiomatic solution to these problems in Rust for hours.

Two particular examples I have had a problem with in the past:

1) I had to modify the contents of two columns in a matrix (`ndarray`) based on each others' contents. In this case, both columns had to be mutable. My first solution was to iterate over the rows which were `vec`s, use `split_at` on the `vec` such that the two columns were now in different vectors and then I was able to get mutable references to both columns. In my next iteration I produced an iterator over the columns which was able to give me mutable references to two columns. This worked, because as the iterator moves from column to column it will consume them. However, this required me to have two branches depending on which column index (the indices were passed as arguments into the function) was lower as once you iterate past a column you can't go back. Finally, after posting an issue on GitHub the authors showed me how to do the same thing with a bit of unsafe code.

2) I had a hashmap of a certain class of objects that produce and consume a particular unit of data. The object needs to be mutable to produce or consume this data. What I wanted to do was to get a unit of this data from on object and have a second object (can also be the same object as the producer) consume it. Production/consumption are completely independent so this was a safe operation, but the compiler did not know this. I ended up simply copying the data out rather than using a reference which allowed me to drop the mutable reference to the first object after the copy.

So where can I go to find some suggestions or design patterns on how to solve this problem idiomatically in Rust?
## [9][Is gfx/portability abandonware?](https://www.reddit.com/r/rust/comments/hy30lw/is_gfxportability_abandonware/)
- url: https://www.reddit.com/r/rust/comments/hy30lw/is_gfxportability_abandonware/
---
I was considering migrating a project from MoltenVK to [gfx/portability](https://github.com/gfx-rs/portability) but it seems like there last commit was like 4 months ago and there have been Vulkan releases since then.  Does anybody know if this project is still being maintained?
## [10][Async/Await for AVR with Rust](https://www.reddit.com/r/rust/comments/hxqtl0/asyncawait_for_avr_with_rust/)
- url: https://lights0123.com/blog/2020/07/25/async-await-for-avr-with-rust/
---

## [11][async-fs: Async filesystem primitives (all runtimes, small dependencies, fast compilation)](https://www.reddit.com/r/rust/comments/hy6lmk/asyncfs_async_filesystem_primitives_all_runtimes/)
- url: https://docs.rs/async-fs
---

## [12][My biggest compliment to Rust is that it's boring, and this is an amazing compliment. – Chris Dickinson, Engineer at npm, Inc](https://www.reddit.com/r/rust/comments/hy69ms/my_biggest_compliment_to_rust_is_that_its_boring/)
- url: https://www.reddit.com/r/rust/comments/hy69ms/my_biggest_compliment_to_rust_is_that_its_boring/
---
hey guys, what do you think this means? why being boring is a compliment?   (￣～￣;)
