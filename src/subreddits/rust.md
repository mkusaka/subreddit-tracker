# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (29/2020)!](https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hq8id7/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hm1pws/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 347](https://www.reddit.com/r/rust/comments/hrc4dt/this_week_in_rust_347/)
- url: https://this-week-in-rust.org/blog/2020/07/14/this-week-in-rust-347/
---

## [3][Programming languages: Now Rust project looks for a way into the Linux kernel | ZDNet](https://www.reddit.com/r/rust/comments/hstd2v/programming_languages_now_rust_project_looks_for/)
- url: https://www.zdnet.com/article/programming-languages-now-rust-project-looks-for-a-way-into-the-linux-kernel/
---

## [4][Programming Servo: just hanging around](https://www.reddit.com/r/rust/comments/hssdct/programming_servo_just_hanging_around/)
- url: https://www.reddit.com/r/rust/comments/hssdct/programming_servo_just_hanging_around/
---
An overview of how Servo, a large concurrent and multi-process code-base written in the Rust programming language, dealt with hanging threads.

[https://medium.com/programming-servo/programming-servo-just-hanging-around-d8f660c33df2?source=friends\_link&amp;sk=6efbf13743aec335bd11834c2df71783](https://medium.com/programming-servo/programming-servo-just-hanging-around-d8f660c33df2?source=friends_link&amp;sk=6efbf13743aec335bd11834c2df71783)
## [5][Announcing Rust 1.45.0 | Rust Blog](https://www.reddit.com/r/rust/comments/hsa7ca/announcing_rust_1450_rust_blog/)
- url: https://blog.rust-lang.org/2020/07/16/Rust-1.45.0.html
---

## [6][Shipping Const Generics in 2020](https://www.reddit.com/r/rust/comments/hs9hbn/shipping_const_generics_in_2020/)
- url: https://without.boats/blog/shipping-const-generics/
---

## [7][The new luminance is there](https://www.reddit.com/r/rust/comments/hsgfav/the_new_luminance_is_there/)
- url: https://phaazon.net/blog/luminance-0.40
---

## [8][What an extraordinary language](https://www.reddit.com/r/rust/comments/hso2kr/what_an_extraordinary_language/)
- url: https://www.reddit.com/r/rust/comments/hso2kr/what_an_extraordinary_language/
---
    pub fn find(&amp;self) -&gt; Option&lt;Box&lt;indextree::Node&lt;TransactionState&gt;&gt;&gt; {
            match arenaReader()
                .arena
                .lock()
                .unwrap()
                .iter()
                .find(|t2| *self == *t2.get().trans())
            {
                Some(x) =&gt; Some(Box::new(x.clone())),
                _ =&gt; None,
            }
        }

See [stackoverflow.com/questions/27791532/how-do-i-create-a-global-mutable-singleton](https://stackoverflow.com/questions/27791532/how-do-i-create-a-global-mutable-singleton)
## [9][Can anything in Rust currently match the performance of uWebSockets?](https://www.reddit.com/r/rust/comments/hsmhas/can_anything_in_rust_currently_match_the/)
- url: https://www.reddit.com/r/rust/comments/hsmhas/can_anything_in_rust_currently_match_the/
---
[https://github.com/uNetworking/uWebSockets](https://github.com/uNetworking/uWebSockets)

This is currently known as possibly the fastest pub/sub WebSocket server out there in terms of single-core performance (number of messages/bytes per second it can push out to clients/number of concurrent connected clients). I'm wondering if anyone knows if it's possible to roughly match the performance of uWebSockets in Rust.

I've tried putting together a simple single-core asynchronous tokio-based broadcast server in Rust based on Warp that broadcasts a message by just directly looping over an Arc&lt;Mutex&lt;Vec&gt;&gt; of connected clients. It's not slow, but uWebSockets can handle about 2.3x as many concurrent clients under my benchmarks (streaming about 150 broadcasts per second) before latency starts spiking.

Are there any other libraries or techniques I should be looking at that might possibly equal the performance of uWebSockets without requiring a Herculean effort? If I can get within, say, 10-30% of uWebSockets, that's probably close enough to pick Rust. However, a 2.3x difference is currently too much to ignore.
## [10][Having a hard time understanding atomic operations in Rust.](https://www.reddit.com/r/rust/comments/hskm11/having_a_hard_time_understanding_atomic/)
- url: https://www.reddit.com/r/rust/comments/hskm11/having_a_hard_time_understanding_atomic/
---
Hello everyone. Today I decided to take a break and watch the Jon Gjengset's stream of [porting Java's ConcurrentHashMap in Rust](https://www.youtube.com/watch?v=yQFWmGaFBjk). Part of the implementation, perhaps if not the majority, relies mostly on atomic operations, which I'm having a hard time understanding. This had been an unknown term for me until I checked the crossbeam crate. But it is still unclear, nonetheless, the purpose of this abstraction. May I ask; would anyone mind further explaining a little bit more about the basis of atomics and the motivation behind the crossbeam crate in which is related?

I'd deeply appreciate some guidance. Thanks in advance.
## [11][git-cm: A tool for generating conventional-commits style commits](https://www.reddit.com/r/rust/comments/hsuihh/gitcm_a_tool_for_generating_conventionalcommits/)
- url: https://www.reddit.com/r/rust/comments/hsuihh/gitcm_a_tool_for_generating_conventionalcommits/
---
Hey!

I just published a custom git command, `git-cm`. It lets you interactively generate conventional-commits styled commit messages by asking you questions.

[Here](https://github.com/SirWindfield/git-cm) is a link to the repo.

Hopefully this is as useful to others as it is to me :)
If you have ideas or found bugs, just open up a new issue :)
## [12][How do I convert &amp;str to *const c_char?](https://www.reddit.com/r/rust/comments/hsssrf/how_do_i_convert_str_to_const_c_char/)
- url: https://www.reddit.com/r/rust/comments/hsssrf/how_do_i_convert_str_to_const_c_char/
---
I know how to convert c_char to String and &amp;str. How do I do the opposite? 


Specifically, `&amp;str`  -&gt;  `*const c_char`
