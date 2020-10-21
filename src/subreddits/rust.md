# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (43/2020)!](https://www.reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/jdwuis/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/j9l01t/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (43/2020)?](https://www.reddit.com/r/rust/comments/jdwv7h/whats_everyone_working_on_this_week_432020/)
- url: https://www.reddit.com/r/rust/comments/jdwv7h/whats_everyone_working_on_this_week_432020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-43-2020/50310?u=llogiq)!
## [3][Built with Iced / Rust: Cryptowatch Desktop](https://www.reddit.com/r/rust/comments/jey2uh/built_with_iced_rust_cryptowatch_desktop/)
- url: https://v.redd.it/dpw64usu3bu51
---

## [4][Why are there no increment (++) and decrement (--) operators in Rust?](https://www.reddit.com/r/rust/comments/jf66eu/why_are_there_no_increment_and_decrement/)
- url: https://www.reddit.com/r/rust/comments/jf66eu/why_are_there_no_increment_and_decrement/
---
I've just started learning Rust, and it struck me as a bit odd that `x++` and `x--` aren't a part of the Rust language. I did some research, and I found this vague explanation in [Rust's FAQ](https://github.com/dtolnay/rust-faq#why-doesnt-rust-have-increment-and-decrement-operators):

&gt;Preincrement and postincrement (and the decrement equivalents), while convenient, are also fairly complex. They require knowledge of evaluation order, and often lead to subtle bugs and undefined behavior in C and C++. `x = x + 1` or `x += 1` is only slightly longer, but unambiguous.

What are these "subtle bugs and undefined behavior\[s\]"? In all programming languages I know of, `x++` is exact shorthand for `x += 1`, which is in turn exact shorthand for `x = x + 1`. Likewise for `x--`. That being said, I've never used C or C++ so maybe there's something I don't know.

Thanks for the help in advance!
## [5][A federated reddit alternative written in rust - Lemmy Release v0.8.0 🥳 - Federation beta!](https://www.reddit.com/r/rust/comments/jeraka/a_federated_reddit_alternative_written_in_rust/)
- url: https://dev.lemmy.ml/post/41605
---

## [6][Ridiculously fast Unicode UTF-8 validation](https://www.reddit.com/r/rust/comments/jey87h/ridiculously_fast_unicode_utf8_validation/)
- url: https://lemire.me/blog/2020/10/20/ridiculously-fast-unicode-utf-8-validation/
---

## [7][Non-Generic Inner Functions](https://www.reddit.com/r/rust/comments/jf43xp/nongeneric_inner_functions/)
- url: https://www.possiblerust.com/pattern/non-generic-inner-functions
---

## [8][RustFest Global: a new look &amp; tickets are available!](https://www.reddit.com/r/rust/comments/jf8bc1/rustfest_global_a_new_look_tickets_are_available/)
- url: https://blog.rustfest.eu/a-new-look
---

## [9][Toy Neural Network in Rust](https://www.reddit.com/r/rust/comments/jf6t5v/toy_neural_network_in_rust/)
- url: https://www.reddit.com/r/rust/comments/jf6t5v/toy_neural_network_in_rust/
---
I ported iamtrask's [A Neural Network in 11 lines of Python (Part 1)](https://iamtrask.github.io/2015/07/12/basic-python-network/) to rust.

    // Original python code
    // https://iamtrask.github.io/2015/07/12/basic-python-network/
    // Neural Network = 3 inputs, 4 hidden, 1 output with sigmoid activation

    use ndarray::{array, Array};
    use ndarray_rand::{rand::SeedableRng, rand_distr::Uniform, RandomExt}; // random matrix generation
    use rand_isaac::isaac64::Isaac64Rng; // for setting seed

    fn main() -&gt; std::io::Result&lt;()&gt; {
        // set seeds
        let seed = 42;
        let mut rng = Isaac64Rng::seed_from_u64(seed);

        // dataset
        let x = array![[0., 0., 1.], [0., 1., 1.], [1., 0., 1.], [1., 1., 1.]];
        let y = array![[0., 1., 1., 0.]].reversed_axes();

        // initializing random weights
        let mut w0 = Array::random_using((3, 4), Uniform::new(0., 1.), &amp;mut rng); // syn0
        let mut w1 = Array::random_using((4, 1), Uniform::new(0., 1.), &amp;mut rng); // syn1

        for i in 0..1001 {
            // forward propagation
            let hidden = -x.dot(&amp;w0); // l1
            let hidden = 1. / (1. + (hidden.mapv(f64::exp)));
            let output = -hidden.dot(&amp;w1); // l2
            let output = 1. / (1. + (output.mapv(f64::exp)));
            
            // gradient calculation
            let output_delta = (&amp;y - &amp;output) * (output.clone() * (1. - &amp;output)); // l2_delta
            let hidden_delta = output_delta.dot(&amp;w1.t()) * (hidden.clone() * (1. - &amp;hidden)); // l1_delta

            // update weights
            w1 = &amp;w1 + &amp;hidden.t().dot(&amp;output_delta);
            w0 = &amp;w0 + &amp;x.t().dot(&amp;hidden_delta);

            if i % 100 == 0 {
                println!("Epoch: {} \n {:?}\n", i,output);
            }
            
        }

        Ok(())
    }

Prediction

    Epoch: 1000 
     [[0.039195556811580975],
     [0.9510824182196344],
     [0.9514737497828656],
     [0.054103145866190874]], shape=[4, 1], strides=[1, 1], layout=C (0x1), const ndim=2
## [10][Any tips for Node.js developers looking into Rust?](https://www.reddit.com/r/rust/comments/jfckz5/any_tips_for_nodejs_developers_looking_into_rust/)
- url: https://www.reddit.com/r/rust/comments/jfckz5/any_tips_for_nodejs_developers_looking_into_rust/
---
So I’m working through the ‘Rust Book’ and am really enjoying learning rust. I have virtually no exposure to C/Cpp so it’s definitely a learning curve. 

If you came from a node/js background, are there any nuggets of wisdom you can share?
## [11][Blue Pill (STM32) Rust Workshop: Chapter 1: Toolchain Setup](https://www.reddit.com/r/rust/comments/jeui3t/blue_pill_stm32_rust_workshop_chapter_1_toolchain/)
- url: https://www.youtube.com/playlist?list=PLDWmoWFf46givBRQmh5DyE27OsXMJPfag
---

## [12][Marking issues as regressions](https://www.reddit.com/r/rust/comments/jeql7u/marking_issues_as_regressions/)
- url: https://blog.rust-lang.org/2020/10/20/regression-labels.html
---

