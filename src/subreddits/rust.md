# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (10/2020)!](https://www.reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/fc7h20/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/f8ney8/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (10/2020)?](https://www.reddit.com/r/rust/comments/fc7hn2/whats_everyone_working_on_this_week_102020/)
- url: https://www.reddit.com/r/rust/comments/fc7hn2/whats_everyone_working_on_this_week_102020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-10-2020/38939?u=llogiq)!
## [3][caniuse.rs | Rust feature search](https://www.reddit.com/r/rust/comments/fdj80k/caniusers_rust_feature_search/)
- url: https://caniuse.rs/
---

## [4][Timetill.rs: Now with updated conference lineup for 2020.](https://www.reddit.com/r/rust/comments/fdtok1/timetillrs_now_with_updated_conference_lineup_for/)
- url: https://timetill.rs/
---

## [5][WebGPU, Frame Capturing and more in the latest update for Nannou - an open source, creative coding framework for Rust!](https://www.reddit.com/r/rust/comments/fduj17/webgpu_frame_capturing_and_more_in_the_latest/)
- url: https://nannou.cc/posts/nannou_v0.13
---

## [6][Is there a way to transfer variable ownership into Iterator?](https://www.reddit.com/r/rust/comments/fdqx5o/is_there_a_way_to_transfer_variable_ownership/)
- url: https://www.reddit.com/r/rust/comments/fdqx5o/is_there_a_way_to_transfer_variable_ownership/
---
Hello,

I'd like to make Iterator to own variable and serve it's immutable reference during evaluation.

(and that immutable reference is only needed during evaluation and will not be used on the final result form)

The code below works, but I used Rc for solving this issue.

    // Thanks to https://github.com/yjh0502
    // He suggested this solution: use Rc.
    use std::rc::Rc;
    
    // Without Copy or Clone
    #[derive(Debug)]
    struct Data {
      pub s: String,
    }
    
    fn main() {
      println!("Hello, world!");
      
      get_iter().for_each(|(data, n)| {
         println!("{:?} - {}", data, n); 
      });
    }
    
    // data should be declared in the content of get_iter function.
    fn get_iter() -&gt; Box&lt;dyn Iterator&lt;Item = (Rc&lt;Data&gt;, i32)&gt;&gt;{
      let data = Rc::new(
          Data { s: "Immutable String to share".to_string() }
      );
    
      let iter = vec![1, 2, 3, 4, 5]
          .into_iter()
          .map(move |n| (Rc::clone(&amp;data), n) );
    
      Box::new(iter)
    }

What I'd like to use is something similar with below one,

    fn get_iter() -&gt; Box&lt;dyn Iterator&lt;Item = (&amp;Data, i32)&gt;&gt;{
      let data = Data { s: "Immutable String to share".to_string() };
    
      let iter = vec![1, 2, 3, 4, 5]
          .into_iter()
          .map(move |n| (&amp;data), n));
    
      Box::new(iter)
    }

Like this code below, I only need `&amp;data` until it's evaluation.

      get_iter().for_each(|(data, n)| {
         println!("{:?} - {}", data, n); 
      });
    
      // or if I use collect, it will be like this
      // no more data needed, but no idea how to make this iterator know it.
      get_iter()
          .map(|(_, n)| n)
          .collect();

after `for_each` or `collect`, I don't need `data` anymore but I have no idea on giving proper lifetime for this case.

Is there a way to

1. Transfer variable ownership into Iterator.
2. That variable doesn't implement Copy or Clone traits.
3. Iterator will serve the variable by immutable reference on iterating.
4. But it is ensured that the variable will not be used in the final form (ex. `collect`)

meet all those conditions above by setting proper lifetime param on variable `data`?

Thanks!
## [7][Cheap tricks for high-performance Rust](https://www.reddit.com/r/rust/comments/fdbszu/cheap_tricks_for_highperformance_rust/)
- url: https://deterministic.space/high-performance-rust.html
---

## [8][microfft: Embedded-friendly Fast Fourier Transforms](https://www.reddit.com/r/rust/comments/fdi1nm/microfft_embeddedfriendly_fast_fourier_transforms/)
- url: https://gitlab.com/ra_kete/microfft-rs
---

## [9][[APPEAL] Use maintenance label in your projects!](https://www.reddit.com/r/rust/comments/fdlbqb/appeal_use_maintenance_label_in_your_projects/)
- url: https://www.reddit.com/r/rust/comments/fdlbqb/appeal_use_maintenance_label_in_your_projects/
---
Hi everyone,  


Cargo.toml allow specify maintenance label. Here are some info from [documentation](https://doc.rust-lang.org/cargo/reference/manifest.html#package-metadata):

```
# Maintenance: `status` is required. Available options are:
# - `actively-developed`: New features are being added and bugs are being fixed.
# - `passively-maintained`: There are no plans for new features, but the maintainer intends to
#   respond to issues that get filed.
# - `as-is`: The crate is feature complete, the maintainer does not intend to continue working on
#   it or providing support, but it works for the purposes it was designed for.
# - `experimental`: The author wants to share it with the community but is not intending to meet
#   anyone's particular use case.
# - `looking-for-maintainer`: The current maintainer would like to transfer the crate to someone
#   else.
# - `deprecated`: The maintainer does not recommend using this crate (the description of the crate
#   can describe why, there could be a better solution available or there could be problems with
#   the crate that the author does not want to fix).
# - `none`: Displays no badge on crates.io, since the maintainer has not chosen to specify
#   their intentions, potential crate users will need to investigate on their own.
maintenance = { status = "..." }
```

This will be even more useful if lib.rs or crates.io will allow search by this labels.

I think sending PR to crates without maintenance status could be a nice contribution to make rust ecosystem better - especially for `deprecated` and `looking-for-maintainer` status.
## [10][async/await on embedded Rust](https://www.reddit.com/r/rust/comments/fdd28f/asyncawait_on_embedded_rust/)
- url: https://ferrous-systems.com/blog/async-on-embedded/
---

## [11][CHIP-8 Game Emulator in Rust for PineTime Smart Watch](https://www.reddit.com/r/rust/comments/fdo1px/chip8_game_emulator_in_rust_for_pinetime_smart/)
- url: https://lupyuen.github.io/pinetime-rust-mynewt/articles/chip8
---

## [12][Serverless + Rust: I tried it with Cloudflare Workers](https://www.reddit.com/r/rust/comments/fdmzyh/serverless_rust_i_tried_it_with_cloudflare_workers/)
- url: https://www.reddit.com/r/rust/comments/fdmzyh/serverless_rust_i_tried_it_with_cloudflare_workers/
---
I lightly [documented](https://medium.com/@psiphi75/rust-and-serverless-with-a-focus-on-cloudflare-workers-342effbc4f85) my experience with Rust serverless using Cloudflare Workers.

**TL;DR** There is a lot of promise, but the overall state of Rust on serverless is pretty immature. This is likely to change in the next 12 months.
