# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (31/2020)!](https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hynlfl/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/hurj77/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 350](https://www.reddit.com/r/rust/comments/i4a3e0/this_week_in_rust_350/)
- url: https://this-week-in-rust.org/blog/2020/08/04/this-week-in-rust-350/
---

## [3][rustc 1.44.1 is reproducible in Debian](https://www.reddit.com/r/rust/comments/i4ij47/rustc_1441_is_reproducible_in_debian/)
- url: http://reproducible.debian.net/redirect?SrcPkg=rustc
---

## [4][Rust on iOS with SDL2](https://www.reddit.com/r/rust/comments/i4o9ie/rust_on_ios_with_sdl2/)
- url: https://blog.aclysma.com/rust-on-ios-with-sdl2/
---

## [5][Crust of Rust: Channels [video]](https://www.reddit.com/r/rust/comments/i4iyku/crust_of_rust_channels_video/)
- url: https://www.youtube.com/watch?v=b4mS5UPHh20
---

## [6][Google engineers just submitted a new LLVM optimizer for consideration which gains an average of 2.33% perf.](https://www.reddit.com/r/rust/comments/i44ahg/google_engineers_just_submitted_a_new_llvm/)
- url: https://lists.llvm.org/pipermail/llvm-dev/2020-August/144012.html
---

## [7][Easy Variadic functions for Monoids](https://www.reddit.com/r/rust/comments/i4lcv2/easy_variadic_functions_for_monoids/)
- url: https://www.reddit.com/r/rust/comments/i4lcv2/easy_variadic_functions_for_monoids/
---
[I made a gist](https://gist.github.com/richardlett/9f49d2f4301868b50ea7509cbdab6581) that allows for easy creation of variadic functions from monoids.

(Besides being small, I made it as gist instead of a library so you can easily use builtin types.)

I.e. Your use case is the function being the repeated application of a binary operation, like summation,  multiplication, or function composition.

It should be no overhead (besides the identity operation), but requires nightly and the TupleList Crate.

You can create multiple variadic functions on the same type too (say you wanted to do addition and multiplication), each one is denoted by a constant generic usize. All you have to do is implement `Monoid&lt;usize label&gt;` for your type, which entails giving an identity element a binary operator. Then just name your function. Here's summation:

    #![feature(fn_traits, unboxed_closures)]
    #![feature(const_fn)]
    #![feature(const_generics)]
    
    mod veriadics;
    use veriadics::*;
    
    impl Monoid&lt;1&gt; for i32 {
        fn operator(first: Self, second: Self) -&gt; Self { first + second }
        fn identity() -&gt; Self { 0 }
    }
    
    const sum: VariFunc&lt;i32, 1&gt; = gen_function::&lt;i32, 1&gt;();
    
    fn main() {
        println!("{}", sum(1, 2, 3, 4, 5, 6, 7, 8));
    }
## [8][v0.13 - MeiliSearch, open source alternative to Algolia](https://www.reddit.com/r/rust/comments/i4qd69/v013_meilisearch_open_source_alternative_to/)
- url: https://blog.meilisearch.com/whats-new-in-0-13-0/
---

## [9][how to upload a file using reqwest but streaming the body](https://www.reddit.com/r/rust/comments/i4n9mp/how_to_upload_a_file_using_reqwest_but_streaming/)
- url: https://www.reddit.com/r/rust/comments/i4n9mp/how_to_upload_a_file_using_reqwest_but_streaming/
---
I have a big file that I would like to upload in  parts (10M each), this is the code I am using to split the file it works for writing the chunk in another file but I would like to instead stream the `buffer` and prevent having to writing to disk first:

    use std::env;
    use std::fs::File;
    use std::io::prelude::*;
    use std::io::{BufReader, Write};
    use std::process;
    
    fn main() {
        let args: Vec&lt;String&gt; = env::args().collect();
        if args.len() &lt; 2 {
            eprintln!("missing arguments: file");
            process::exit(1);
        }
        let file_path = &amp;args[1];
        let mut file = File::open(&amp;file_path).unwrap();
    
        let mut count = 1;
        loop {
            let seek = file.seek(SeekFrom::Current(0)).unwrap();
            println!("seek: {}", seek); // in case need to resume 
            let mut reader = BufReader::new(&amp;file);
            if reader.fill_buf().unwrap().is_empty() {
                break;
            }
            let mut reader = reader.take(10_485_760);
            let mut f = File::create(&amp;format!("/tmp/chunk_{}", count)).unwrap();
            loop {
                let consummed = {
                    let buffer = reader.fill_buf().unwrap();
                    if buffer.is_empty() {
                        break;
                    }
                    // how to stream the buffer
                    // client.put(url).headers(headers).body(Body::from(&amp;buffer)) ???
                    f.write_all(&amp;buffer).expect("Unable to write data");
                    buffer.len()
                };
                reader.consume(consummed);
            }
            count += 1;
        }
    }

Here is where I read in chunks for testing I am writing it to a file but how could I do the same for streaming the file.

    loop {
        let consummed = {
            let buffer = reader.fill_buf().unwrap();
            if buffer.is_empty() {
               break;
            }
            // how to stream the buffer ?
            // client.put(url).headers(headers).body(Body::from(&amp;buffer)) ?
            f.write_all(&amp;buffer).expect("Unable to write data");
            buffer.len()
        };
        reader.consume(consummed);
    }

For uploading the file in one shot I am using `FramedRead`

    let file = File::open(file_path).await?;
    let stream = FramedRead::new(file, BytesCodec::new());
    let body = Body::wrap_stream(stream);
    client.request(method, url).headers(headers).body(body)

But haven't found a way to split the file and then use `FrameRead` or how to pass only the chunk so that it can be streamed.

Any ideas?
## [10][Sirula - a simple app launcher for wayland written in rust](https://www.reddit.com/r/rust/comments/i46d1w/sirula_a_simple_app_launcher_for_wayland_written/)
- url: https://v.redd.it/l6qga4pb27f51
---

## [11][How to read non UTF-8 files in Windows?](https://www.reddit.com/r/rust/comments/i4oso4/how_to_read_non_utf8_files_in_windows/)
- url: https://www.reddit.com/r/rust/comments/i4oso4/how_to_read_non_utf8_files_in_windows/
---
Is there a mean to tell  `std::io::BufRead::lines()` to use a different "code page" than UTF-8 on Windows?

I've written a small utility to interact with command line tools and their output can contain non standard ASCII characters.  The default code page on my Windows 10 box is 850.  As such, I get errors while trying to read the output.  Change the default code page system wide to UTF-8 works but as it is beta and it interfere negatively with other programs, I cannot have it as a default.

All I can do now is to generate an error message and do a lossy conversion.

    let output = process::Command::new(SVN_CMD)
            .args(&amp;[SVN_PG, SVN_EXTERNALS, dir.to_str().unwrap()])
            .output()
            .expect(&amp;("Failed to execute ..."));
    if output.status.success() {
        let output = output.stdout;
        let lines = match String::from_utf8(output.clone()) {
            Ok(lines) =&gt; lines,
            Err(_) =&gt; String::from_utf8_lossy(&amp;output).to_string(),
        };
## [12][Why QEMU should move from C to Rust](https://www.reddit.com/r/rust/comments/i4rpyc/why_qemu_should_move_from_c_to_rust/)
- url: http://blog.vmsplice.net/2020/08/why-qemu-should-move-from-c-to-rust.html
---

