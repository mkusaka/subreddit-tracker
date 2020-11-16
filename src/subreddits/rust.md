# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (47/2020)!](https://www.reddit.com/r/rust/comments/juwjxb/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/juwjxb/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/jqrkpa/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][What's everyone working on this week (47/2020)?](https://www.reddit.com/r/rust/comments/juwkx4/whats_everyone_working_on_this_week_472020/)
- url: https://www.reddit.com/r/rust/comments/juwkx4/whats_everyone_working_on_this_week_472020/
---
New week, new Rust! What are you folks up to? Answer here or over at [rust-users](https://users.rust-lang.org/t/whats-everyone-working-on-this-week-47-2020/51526?u=llogiq)!
## [3][tiny-skia - a new, pure Rust 2D rendering library based on a Skia subset](https://www.reddit.com/r/rust/comments/juy6x7/tinyskia_a_new_pure_rust_2d_rendering_library/)
- url: https://www.reddit.com/r/rust/comments/juy6x7/tinyskia_a_new_pure_rust_2d_rendering_library/
---
https://github.com/RazrFalcon/tiny-skia

For a long time, I was looking for a great 2D rendering library to use in [resvg](https://github.com/RazrFalcon/resvg), but there was not much choice. You can either use [raqote](https://github.com/jrmuizel/raqote), which is a pure Rust library, but very slow and not actively maintained. Or use [Skia](https://skia.org/), which is probably the best 2D rendering library, but it's absurdly heavy and very hard to compile and distribute.

So the obvious solution was to write your own. But since I have (or at least had) zero knowledge about how 2D libraries work, I couldn't work on raqote myself, so I've decided to port Skia parts that I've need, mainly the CPU rendering pipeline.

After three months of work, I've ended up with tiny-skia, which does all I need (shapes, gradients, patterns, quality, performance), passes all resvg tests and not more than 2x slower than Skia (which is already blazingly fast).

So in the end, tiny-skia is the new 2D rendering library for the Rust ecosystem, which provides better quality than raqote (mainly hairline stroking, gradients and bicubic scaling) and still 5-7x faster.

The main missing feature is the text rendering, which is absurdly complex and will [take a while](https://github.com/RazrFalcon/tiny-skia/issues/1) to implement. And there is not much we can port from Skia, since it uses native libraries.

I still consider it a beta quality, so testing is welcome. There are also some small missing features and performance optimizations.
## [4][Feroxbuster - A simple, fast, recursive content discovery tool written in Rust](https://www.reddit.com/r/rust/comments/jv0xaw/feroxbuster_a_simple_fast_recursive_content/)
- url: https://i.redd.it/sfxkz1p8ajz51.gif
---

## [5][shotgun - Minimal screenshot tool written in Rust](https://www.reddit.com/r/rust/comments/jv0v32/shotgun_minimal_screenshot_tool_written_in_rust/)
- url: https://github.com/neXromancers/shotgun
---

## [6][drop-bin: Defer running expensive destructors](https://www.reddit.com/r/rust/comments/juqi8m/dropbin_defer_running_expensive_destructors/)
- url: https://www.reddit.com/r/rust/comments/juqi8m/dropbin_defer_running_expensive_destructors/
---
One of the major advantages that garbage-collected languages have is that memory freeing doesn't happen in the fast path. Smart GCs will often wait until memory usage is low or nothing needs to be processed to clear up all the garbage. To bring this same performance to Rust, I created [drop-bin](https://docs.rs/drop-bin); instead of leaving your values to be implicitly dropped, put them in the bin, and then you can clear them all out at the same time whenever you like.

To test its performance, I benchmarked dropping a `HashMap` of 1000 vectors of 500 number each. Simply dropping each value took 240us/iteration, using [defer-drop](https://docs.rs/defer-drop) took 1342ns/iteration and drop-bin took 450ns/iteration. This is cheating somewhat as the drop-bin benchmark didn't clear the bin and so the values were never actually dropped, but it is more representative of the performance you care about in many cases.
## [7][Cn v2.0.0 Released](https://www.reddit.com/r/rust/comments/jv270t/cn_v200_released/)
- url: https://gitlab.com/arijit79/cn/-/releases/v2.0.0
---

## [8][rust-analyzer changelog #51](https://www.reddit.com/r/rust/comments/jv6nof/rustanalyzer_changelog_51/)
- url: https://rust-analyzer.github.io/thisweek/2020/11/16/changelog-51.html
---

## [9][Production-Grade Logging in Rust Applications](https://www.reddit.com/r/rust/comments/jv6a5j/productiongrade_logging_in_rust_applications/)
- url: https://medium.com/better-programming/production-grade-logging-in-rust-applications-2c7fffd108a6
---

## [10][How to select between reading from a file and stdin](https://www.reddit.com/r/rust/comments/jv3q3e/how_to_select_between_reading_from_a_file_and/)
- url: https://www.reddit.com/r/rust/comments/jv3q3e/how_to_select_between_reading_from_a_file_and/
---
I'm trying to write a simple base64 encoder/decoder in Rust to learn the language, really like it so far but I have some problems when it comes to IO.

I want the option to read from either a local file or from stdin and I have been able to make some example code for both cases. The code looks almost the same and coming from a C background I try to assign a "file descriptor" using the following method:

        let mut data_src_fd;
        if data_src.eq("-") {
    	let stdin = io::stdin();
    	let handler = stdin.lock();
    	data_src_fd = handler.by_ref();
        } else {
    	let mut handler2 = match File::open(&amp;data_src) {
                Ok(file) =&gt; file,
                Err(error_description) =&gt; {
    		eprintln!(
                        "Unable to open file {} ({})",
                        data_src,
                        error_description.to_string()
    		);
    		process::exit(-1);
                }
    	};
    	data_src_fd = handler2.by_ref();
        }

This does not work whoever and gives me the following error:

    error[E0308]: mismatched types
      --&gt; src/main.rs:60:16
       |
    60 |     data_src_fd = handler2.by_ref();
       |                   ^^^^^^^^^^^^^^^^^ expected struct `std::io::StdinLock`, found struct `std::fs::File`
       |
       = note: expected mutable reference `&amp;mut std::io::StdinLock&lt;'_&gt;`
                  found mutable reference `&amp;mut std::fs::File`
    
    error: aborting due to previous error
    
    For more information about this error, try `rustc --explain E0308`.

Which make sense, but since I'm a C person I want to make something like this:

      FILE *file_fd;
      if (strncmp(file, "-", 1) == 0)
          file_fd = stdin;
      else
          file_fd = fopen(file, "r");

The Rust reading loop further down looks exactly the same in the example code that I have from reading from a file and reading from stdin:

    let mut buffer_vec = Vec::with_capacity(buffer_size);
    loop {
        match data_src_fd
            .by_ref()
            .take(buffer_size as u64)
            .read_to_end(&amp;mut buffer_vec)
        {
            Ok(chunk_size) =&gt; {
                if chunk_size == 0 {
                    break;
                }
    
                operation(&amp;buffer_vec);
    
                if chunk_size &lt; buffer_size {
                    break;
                }
    
                buffer_vec.clear();
            }
            Err(error_description) =&gt; panic!("{}", error_description),
        }
    }

It's very frustrating to have two identical loops for handling both cases so I'm wondering how I can attractive the same result as in the C code (if its possible?). I'm very new to Rust so I'm much open for suggestions for how I should be handling this and to rewrite/redesign the code if necessary.
## [11][t-rec-rs — Blazingly fast terminal recorder that generates animated gif images for the web](https://www.reddit.com/r/rust/comments/jv0r22/trecrs_blazingly_fast_terminal_recorder_that/)
- url: https://github.com/sassman/t-rec-rs
---

## [12][What are your rules of thumb around when to panic and when to return an error?](https://www.reddit.com/r/rust/comments/jv6khs/what_are_your_rules_of_thumb_around_when_to_panic/)
- url: https://www.reddit.com/r/rust/comments/jv6khs/what_are_your_rules_of_thumb_around_when_to_panic/
---
This isn't even necessarily Rust specific. So as you discuss, feel free to mention other languages or different domains, etc as you see fit.

It's easy to find hand-wavy advice along the lines of "panic is okay when the caller can't reasonably recover from it." But what does that **mean** to you? And doesn't it sound kind of presumptuous to assume what the caller may or may not want to recover from? Maybe if the file isn't found, they want to turn around and order a pizza and a margarita while submitting their resumé to the animal shelter to be a dog catcher and get the hell out of tech. How you can you predict that as library author?

On the other hand, pretty much nobody tries to do anything about memory allocation failures.

I like to refer to OCaml's docs sometimes: https://dev.realworldocaml.org/error-handling.html#scrollNav-3

Personally, I take a more "philosophical" and decidedly less pragmatic approach to errors (in all languages with which I work). I return a Result if the function can fail due to the inputs or due to the state of the world. Always. If the function does IO, I return a Result if the database is missing, the network request failed, etc. The **only** time I panic is if 100% of the reason lies within the function itself. For example, if my function is processing some data and I create a temporary array/vector/list and I *know* it's at least length 3, then I'm perfectly happy to index it directly at `list[2]` and if it crashes, it's because I wrote the function wrong.

I was thinking about this because I was reading some people criticisms of Java's checked exceptions and about IOException in particular. The argument was that IOException should be unchecked, so that Java IO functions in the standard library just bubble up IO failures transparently to the top. I'm not comfortable with that idea...
