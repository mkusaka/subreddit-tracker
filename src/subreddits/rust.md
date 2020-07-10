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

## [3][I was wondering "what's the big deal with Rust?" ... and then I tried multithreading](https://www.reddit.com/r/rust/comments/hohjdu/i_was_wondering_whats_the_big_deal_with_rust_and/)
- url: https://www.reddit.com/r/rust/comments/hohjdu/i_was_wondering_whats_the_big_deal_with_rust_and/
---
Consider me completely sold.

Ok so the rest of this post is going to be an explanation of how I was able to speed up my program and the problems I faced, without going into too much specific detail because I'm not sure how much I can disclose.

Background: I'm new to Rust and have been learning it incrementally over the past few months. I've been making some small projects and learning new features as I need them. This is my preferred learning style. I have prior experience in C++ and Haskell which made me quite interested in Rust as a language. I've also never written a concurrent program before. So I was psyched when I was able to speed up my program 6x in a single day!

The problem: I have to read a bunch of lines (with different sizes, and an unknown number of lines) from a file, handle the first x lines separately (x is unknown), and then for the remaining lines, do some computation on each line (string splitting, regex matching, finding within a list of strings), and then write the output to another file. This is basically the worst possible case for multithreading as far as I can tell: file I/O, line by line processing but unable to do any offsets beforehand, impossible to use map or filter because I have to handle the first x lines separately. Oh, and the entire thing is in a `while` loop (I guess, a `loop` loop in Rust). The one thing I've got going for me is that the order of the output lines doesn't matter.

The solution: it turns out that multithreading this was dirt simple. I simply had to put my `loop` into a [ThreadPool.execute](https://docs.rs/threadpool/1.8.1/threadpool/struct.ThreadPool.html#method.execute), wrap my file reader and writer in an `Arc&lt;Mutex&gt;` (since they are mutable) and my list of strings in an `Arc` (since it's immutable), clone the `Arc`'s within each thread, and change some `&amp;`'s to `*`'s. Bada Bing Bada Boom. Multithreading achieved with no changes to my algorithm. Instant 6x speedup!

And then I realized that I was being stupid and was [compiling a Regex within a loop](https://docs.rs/regex/1.3.9/regex/#example-avoid-compiling-the-same-regex-in-a-loop). Once I pulled it out of the loop, I got another 10x speedup. So overall, a 60x speedup within a day. Went from taking 2 hours to 2 minutes.

Life lessons learned:

1. Rust makes multithreading extremely simple. You can do it with minimal changes to your existing code, even if the problem is strangely difficult and uses `while` loops and is impossible to map or reduce.

2. Don't compile a Regex within a loop! This simple fix was actually more impactful than the multithreading (although put together, they knocked my socks off).
## [4][Linux Developers May Discuss Allowing Rust Code Within The Kernel](https://www.reddit.com/r/rust/comments/hojg5j/linux_developers_may_discuss_allowing_rust_code/)
- url: https://www.phoronix.com/scan.php?page=news_item&amp;px=Linux-Plumbers-2020-Rust
---

## [5][Help with “Raytracing in One Weekend”](https://www.reddit.com/r/rust/comments/honu16/help_with_raytracing_in_one_weekend/)
- url: https://www.reddit.com/r/rust/comments/honu16/help_with_raytracing_in_one_weekend/
---
For those who don’t know it, [Raytracing in One Weekend](https://raytracing.github.io/) is a tutorial series where you build a raytracer. It’s written in C++, but I've working through it in Rust after seeing it suggested here. I have two problems, one general and one specific.

The specific problem is that something is clearly wrong with my code for the Cornell box example. [This](https://raytracing.github.io/images/img.cornell-box.jpg) is what it looks like in the book. [This](https://imgur.com/a/N9tQRpA) is what mine looks like. The original is somewhat noisy already, but mine is extremely grainy and dark. Is this a known problem? I've been over their and my code over and over again and just can’t figure out what the difference is. There is a small logic difference where I think their code is just overly complicated, I can expand on what I mean if need be.

The more general problem concerns architecture. Currently I have two important traits, `Geometry` and `Optics`. `Geometry` is for things that can be intersected by rays and have a bounding box. `Optics` is for things that can be hit by rays and also reflect or refract them and affect their colors. The primary implementer of `Geometry` is an enum called `Shape` (containing Spheres and Rectangles), the primary implementer of `Optics` is the `Object` struct, which contains a `Geometry` trait object and a `Material`. This approach means I have to use trait objects practically everywhere and I’m getting this feeling that people describe sometimes where you're fighting the compiler because you're going against the grain, as it were. The C++ version in the book bundles all the geometry and optics functionality together in a `hittable` class; I thought separating out those things that are purely concerned with geometry would be cleaner, but maybe I’m overcomplicating things. What are your thoughts on the most sound way to structure this?

My repository is [here](https://github.com/loewenheim/raytracing/).
## [6][TIL you can use `impl Trait` as generic parameters in return type](https://www.reddit.com/r/rust/comments/hoj6ph/til_you_can_use_impl_trait_as_generic_parameters/)
- url: https://www.reddit.com/r/rust/comments/hoj6ph/til_you_can_use_impl_trait_as_generic_parameters/
---
This means I can do this, for example:

    // The `data` field stores a list of strings, but what string type is in it
    // should not matter for the user.
    struct Storage {
        data: Vec&lt;Box&lt;str&gt;&gt;,
        // data: Vec&lt;String&gt;,
        // data: Vec&lt;Cow&lt;'static, str&gt;&gt;,
        // data: Vec&lt;beef::lean::Cow&lt;'static, str&gt;&gt;,
        // data: Vec&lt;SmallString&gt;,
    }
    impl Storage {
        // This returns a slice of something that are references to str, without
        // exposing the actual string type.
        fn data(&amp;self) -&gt; &amp;[impl AsRef&lt;str&gt;] {
            &amp;self.data
        }
    }

I was thinking of how to implement this code, it would be nice if I can use `impl Trait` in the generic parameter... and then I actually tried the code and it worked, I was like "wait, this is actually possible?" Somehow I thought you could only use `impl Trait` for the "top-level" return
type (not counting argument position). Perhaps that's because the book and the
edition guide only gave examples using it that way. Just had to get this out of my head.
## [7][[job] Work on Open Source Rust projects at Embark Studios!](https://www.reddit.com/r/rust/comments/ho3z6k/job_work_on_open_source_rust_projects_at_embark/)
- url: https://www.embark-studios.com/jobs/910166-open-source-engineer
---

## [8][Help improve the Rust wiki!](https://www.reddit.com/r/rust/comments/ho9i0h/help_improve_the_rust_wiki/)
- url: https://www.reddit.com/r/rust/comments/ho9i0h/help_improve_the_rust_wiki/
---
The inofficial [Rust Wiki](https://runrust.miraheze.org/) was announced just over a month ago, and there is already a lot of useful content! There are now **90 pages**, including redirects. I've been working on a lot of articles in the [Language Concepts](https://runrust.miraheze.org/wiki/Category:Language_Concepts) category, and there are also tutorials, ecosystem overviews, and more.

Unfortunately, progress has stalled recently, so I'd like to encourage everyone to use the Rust Wiki, fix errors you encounter, and maybe even write some content of your own!

Also see the [previous discussion](https://www.reddit.com/r/rust/comments/g7s4ss/rust_needs_a_wiki/).
## [9][How is exp(x) implemented in the standard library?](https://www.reddit.com/r/rust/comments/hom68b/how_is_expx_implemented_in_the_standard_library/)
- url: https://www.reddit.com/r/rust/comments/hom68b/how_is_expx_implemented_in_the_standard_library/
---
Hi! I just started learning rust and finished chapter three of the book. As an exercise I implemented the exponential function (base e) as such:

    fn exp(x: f64) -&gt; f64 {
        const MAX_ITER: i32 = 200; 
        let mut sum = 1.0;
        let mut term = 1.0;
    
        for n in 1..MAX_ITER {
        term *= x / n as f64;
        sum += term;
        };
    return sum
    }

I tried to figure out how this function is implemented in the standard library but could not really find anything but a reference (I'm probably misusing  reference in this context).
## [10][Niko Matsakis' Async Interviews, #8: Stjepan Glavina](https://www.reddit.com/r/rust/comments/ho8q65/niko_matsakis_async_interviews_8_stjepan_glavina/)
- url: https://smallcultfollowing.com/babysteps/blog/2020/07/09/async-interview-8-stjepan-glavina/
---

## [11][The best Rust video tutorial i have seen and its free (creator: Doug Milford)](https://www.reddit.com/r/rust/comments/hok626/the_best_rust_video_tutorial_i_have_seen_and_its/)
- url: https://www.youtube.com/playlist?list=PLLqEtX6ql2EyPAZ1M2_C0GgVd4A-_L4_5
---

## [12][prodash v7.0 now sports a new renderer to visualize progress of concurrent command-line tools](https://www.reddit.com/r/rust/comments/hojo5x/prodash_v70_now_sports_a_new_renderer_to/)
- url: https://asciinema.org/a/346619
---

