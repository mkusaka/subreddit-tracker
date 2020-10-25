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
## [2][This Week in Rust 361](https://www.reddit.com/r/rust/comments/jg7hkt/this_week_in_rust_361/)
- url: https://this-week-in-rust.org/blog/2020/10/21/this-week-in-rust-361/
---

## [3][VTracer is a raster to vector graphics converter implemented in Rust. I am one of the authors, questions and comments welcome!](https://www.reddit.com/r/rust/comments/jhppai/vtracer_is_a_raster_to_vector_graphics_converter/)
- url: https://www.visioncortex.org/vtracer/
---

## [4][Three open source Sonos projects: efficient embedded development in Rust](https://www.reddit.com/r/rust/comments/jhhyvo/three_open_source_sonos_projects_efficient/)
- url: https://tech-blog.sonos.com/posts/three-open-source-sonos-projects-in-rust/
---

## [5][libffi-rs (Rust bindings to libffi) 1.0.0 is released](https://www.reddit.com/r/rust/comments/jhkrgl/libffirs_rust_bindings_to_libffi_100_is_released/)
- url: https://yorickpeterse.com/articles/libffi-rs-100/
---

## [6][Writing our own simple AWS Lambda Custom Runtime in Rust](https://www.reddit.com/r/rust/comments/jht8z5/writing_our_own_simple_aws_lambda_custom_runtime/)
- url: http://jamesmcm.github.io/blog/2020/10/24/lambda-runtime/#en
---

## [7][Send files to trash?](https://www.reddit.com/r/rust/comments/jhrmol/send_files_to_trash/)
- url: https://www.reddit.com/r/rust/comments/jhrmol/send_files_to_trash/
---
Does anyone know about a (cross-platform) crate that allows sending files to the trash/recycle bin?

Something similar to Python's [https://github.com/arsenetar/send2trash](https://github.com/arsenetar/send2trash), but for Rust
## [8][Debugging for dummies - how to?](https://www.reddit.com/r/rust/comments/jhpliq/debugging_for_dummies_how_to/)
- url: https://www.reddit.com/r/rust/comments/jhpliq/debugging_for_dummies_how_to/
---
Hi folks

I'm comming from Python and I have a simple question. How do you debug your app?

In Python:
User runs the app and hits a critical line - I get traceback in console/logfile. I find the line and put something like `ipdb.set_trace()` and run the app again. Once Python interpreter hits the line it kicks off interactive and nifty debugging console. I can easily debug production apps like that.

In Rust:
Since Rust is compiled when a problem occurs the app just exits and I see literally nothing. If I know where the problem is I use prints. I also tried gdb - not bad but it feels like it's debugging tool from 80's.

My understanding is to have extensive logging thru whole app (to simulate Python tracebacks). I debug with `println("{:?}", ...)` which seems lame to me. So how do you guys debug small CLI or big apps?
## [9][Callbacks On Traits.](https://www.reddit.com/r/rust/comments/jhm12j/callbacks_on_traits/)
- url: https://www.reddit.com/r/rust/comments/jhm12j/callbacks_on_traits/
---
Sorry for what seems like such a basic question, but I can't seem to find anything on this subject in Rust. I'm looking at re-writing some C++ code into Rust. One of the core functionalities is the storage of a collection of interface pointers which we can arbitrarily bind to a member function and perform a callback. (Like a C# delegate).

E.g. https://godbolt.org/z/8aqnd3

Is there some way of doing the same thing as 'invoke' in Rust with a collection of Traits?
## [10][Blog Post: Introducing Ungrammar](https://www.reddit.com/r/rust/comments/jh69jx/blog_post_introducing_ungrammar/)
- url: https://rust-analyzer.github.io/blog/2020/10/24/introducing-ungrammar.html
---

## [11][Need help with concurrency in Rust](https://www.reddit.com/r/rust/comments/jhpv5x/need_help_with_concurrency_in_rust/)
- url: https://www.reddit.com/r/rust/comments/jhpv5x/need_help_with_concurrency_in_rust/
---
I have tried learning about concurrent programming in Rust, I have read the official documentation as well as some tutorials on Youtube, but I am still unable to accomplish a very basic task.

I have a vector of some numbers, and I want to create as many threads as there are elements in this vector and do some operations on those elements (for the sake of example, lets say my program wants to square all elements of the vector). Here is what I have tried:

    use std::thread;
    
    // this function seems pointless since I could just square inside a closure, but its just for example
    fn square(s: i32) -&gt; i32 {
        s * s
    }
    
    // for vector of size N, produces N threads that together process N elements simultaneously
    fn process_parallel(mut v: &amp;Vec&lt;i32&gt;) {
        let mut handles = vec![];
        for i in 0..(v.len()) {
            let h = thread::spawn(move || {
                square(v[i])
            });
            handles.push(h);
        }
        for h in handles {
            h.join().unwrap();
        }
    }
    
    fn main() {
        let mut v = vec![1, 2, 3, 4, 5];
        process_parallel(&amp;mut v);
        // 'v' should countain [1, 4, 9, 16, 25] now
    }

This gives me an error that `v` needs to have static lifetime (which I am not sure is possible). I have also tried wrapping the vector in `std::sync::Arc` but the lifetime requirement still seems to persist. Whats the correct way to accomplish this task?

I know there are powerful external crates for concurrency such as `rayon`, which has method `par_iter_mut()` that would essentially allow me to accomplish this in a single line, but I want to learn about concurrency in Rust and how to write small tasks such as this on my own, so I don't want to move away from `std` for now. 

Any help would be appreciated.
## [12][Is there a way to traverse all `impl`s for a trait?](https://www.reddit.com/r/rust/comments/jhqnum/is_there_a_way_to_traverse_all_impls_for_a_trait/)
- url: https://www.reddit.com/r/rust/comments/jhqnum/is_there_a_way_to_traverse_all_impls_for_a_trait/
---
Hi, everyone,

I’m new to Rust (I usually use other languages), but I found it a good fit for some new project of mine. Now the problem is that said project requires applying different functions to different input, and the most elegant way would be to use traits as far as I can see. (Imagine the standard example of `Animal` and `feed()` must differ according to which animal it actually is...)

Now the number of possible implementations is really huge. Does `main.rs` need to know all of them explicitly - that would certainly increase the code size - or is there a way to loop through all available `impl`s until one matches certain criteria (e.g. `where(animal.name == args.name)`?
