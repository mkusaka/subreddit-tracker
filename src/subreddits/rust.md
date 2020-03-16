# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (12/2020)!](https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fjef12/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/fg0p1v/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (12/2020)?](https://www.reddit.com/r/rust/comments/fjefxj/whats_everyone_working_on_this_week_122020/)
- url: https://www.reddit.com/r/rust/comments/fjefxj/whats_everyone_working_on_this_week_122020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-12-2020/39545?u=llogiq)!
## [3][ripgrep 12 is released](https://www.reddit.com/r/rust/comments/fje2ya/ripgrep_12_is_released/)
- url: https://github.com/BurntSushi/ripgrep/releases/tag/12.0.0
---

## [4][Rust the 4th most active programming languages in GitHub](https://www.reddit.com/r/rust/comments/fj6rp9/rust_the_4th_most_active_programming_languages_in/)
- url: https://learnworthy.net/10-most-active-programming-languages-in-github/
---

## [5][Sheeit - An experimental blazing-fast spreadsheet engine that supports thousands of concurrent users and millions of cells.](https://www.reddit.com/r/rust/comments/fj5uu5/sheeit_an_experimental_blazingfast_spreadsheet/)
- url: https://github.com/wilfredwee/sheeit
---

## [6][Zenith - sort of like top or htop but with histograms, network usage, and more](https://www.reddit.com/r/rust/comments/fjdcla/zenith_sort_of_like_top_or_htop_but_with/)
- url: https://github.com/bvaisvil/zenith
---

## [7][TIL you can put closures of the same type into a collection/iterator without boxing](https://www.reddit.com/r/rust/comments/fjfqxa/til_you_can_put_closures_of_the_same_type_into_a/)
- url: https://www.reddit.com/r/rust/comments/fjfqxa/til_you_can_put_closures_of_the_same_type_into_a/
---
You can use this to implement something akin to C++ output iterator:

```
    fn output_iterator&lt;'item, 'iter, T, C&gt;(
        c: C,
    ) -&gt; impl Iterator&lt;Item = impl FnMut(T) + 'item&gt; + 'iter
    where
        C::IntoIter: 'iter,
        'item: 'iter,
        T: 'item,
        C: IntoIterator&lt;Item = &amp;'item mut T&gt;,
    {
        c.into_iter().map(|r| {
            move |t: T| {
                *r = t;
            }
        })
    }
```

You can then do this:
```
let mut vec = vec![0;10];
for (mut f, v) in output_iterator(&amp;mut vec).zip(0..) {
    f(v);
}
println!("{:?}", vec); // prints [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
```

This seems interesting, though I'm not sure how useful this pattern is in practice. It could be useful if you only want the user of your library to call a function for a certain amount of times, although I can't come up with such a situation.

[Link to Playground](https://play.rust-lang.org/?version=stable&amp;mode=debug&amp;edition=2018&amp;gist=0a6637190d5f5ce80a951ddbdcc7254f)
## [8][Flume, a 100% safe MPSC that's faster than std and gives crossbeam a run for its money](https://www.reddit.com/r/rust/comments/fj17z6/flume_a_100_safe_mpsc_thats_faster_than_std_and/)
- url: https://github.com/zesterer/flume
---

## [9][How I Start: Rust](https://www.reddit.com/r/rust/comments/fjb7fl/how_i_start_rust/)
- url: https://christine.website/blog/how-i-start-rust-2020-03-15
---

## [10][Introducing quit - Exit cleanly with an exit code](https://www.reddit.com/r/rust/comments/fj7k3b/introducing_quit_exit_cleanly_with_an_exit_code/)
- url: https://www.reddit.com/r/rust/comments/fj7k3b/introducing_quit_exit_cleanly_with_an_exit_code/
---
[Documentation](https://docs.rs/quit) | [Repository](https://github.com/dylni/quit) | [crates.io](https://crates.io/crates/quit)

This crate fixes the problems with [`std::process::exit`](https://doc.rust-lang.org/std/process/fn.exit.html). Instead of immediately terminating the program, [`quit::with_code`](https://docs.rs/quit/0.1.0/quit/fn.with_code.html) is similar to a panic. You can still set an exit code, but destructors will be called, and the stack will be unwound to the main function.

It's most useful for CLI scripts, which often exit for normal reasons, such as printing a help message or giving an error about an unknown option. [`std::process::exit`](https://doc.rust-lang.org/std/process/fn.exit.html) is usually overkill in those cases.
## [11][Rust Notebooks: Typed Arrays from String Arrays for Dataset Operation](https://www.reddit.com/r/rust/comments/fjkmxp/rust_notebooks_typed_arrays_from_string_arrays/)
- url: https://shahinrostami.com/posts/programming/rust-notebooks/typed-arrays-from-string-arrays-for-dataset-operation/
---

## [12][What's your controversal Rust opinion?](https://www.reddit.com/r/rust/comments/fj824y/whats_your_controversal_rust_opinion/)
- url: https://www.reddit.com/r/rust/comments/fj824y/whats_your_controversal_rust_opinion/
---
Do you think there's a technique more people should use that many wouldn't like? Is there a popular library you think isn't the best way of doing things (explain why). Is there any patterns you think is confusing enough that people shouldn't do? What is your controversal rust opinion?

Wow, I posted a similar question to learnprogramming and they banned me for it. What a bunch of dickwads
