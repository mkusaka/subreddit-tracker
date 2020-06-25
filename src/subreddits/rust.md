# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (26/2020)!](https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/hdku4k/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last week's thread](https://reddit.com/r/rust/comments/h98zfz/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 344](https://www.reddit.com/r/rust/comments/hepkfq/this_week_in_rust_344/)
- url: https://this-week-in-rust.org/blog/2020/06/23/this-week-in-rust-344/
---

## [3][GNOME Builder ❤️ Rust Analyzer Part 1](https://www.reddit.com/r/rust/comments/hfiuyz/gnome_builder_rust_analyzer_part_1/)
- url: https://gunibert.de/posts/builder_and_analyzer/
---

## [4][Why is my ~500-line single-threaded Rust code running faster in an Ubuntu VM than natively on OSX?](https://www.reddit.com/r/rust/comments/hff9ox/why_is_my_500line_singlethreaded_rust_code/)
- url: https://www.reddit.com/r/rust/comments/hff9ox/why_is_my_500line_singlethreaded_rust_code/
---
Hello, I've got some code that is running *significantly* faster in a VirtualBox VM running Ubuntu 18.04 than on the host machine, a 2019 iMac running macOS Mojave. I am stumped as to why - does Rust optimize so much better for Linux that even a virtual Linux environment can run the code better than a mac? Does mac have some security checks slowing down kernel operations that it somehow lets VirtualBox bypass? Is there some feature of Rust that I am misusing really badly to get this weird behavior?

The code in question is my first Rust program of any length. You give it some letters and it finds a legal bananagrams/scrabble configuration for those letters with minimal bounding box area using a (exponential time) branch and bound approach. The repo is [here](https://github.com/DireLines/bananagrams).

For example,

    cargo run --release -- teststring -c -s 

yields the configuration

    t    
    e n g
    s   i
    t   r
    s   t


I'm trying to optimize this code as an exercise. On Linux you can use perf to profile stuff, so I made an Ubuntu VM and started to do that. I was taken aback when the program ran faster in the VM. I ran this command using [this](https://github.com/DireLines/bananagrams/commit/1caf105ee239ed5140c078a83e12bc14226eeea5) version of the code:

    cargo build --release &amp;&amp; time ./target/release/bananagrams teststring -c -s

On Mac:

    51.53 real 51.45 user 0.02 sys

On Ubuntu VM(!):

    20.10 real 20.06 user 0.01 sys

What gives?

For the record, the best time on Mac so far was 27s using the same code, but with rayon's par_iter() somewhere important instead of iter(). But weirdly, the VM beats this time single-threaded.
## [5][dabreegster/abstreet: A traffic simulation game exploring how small changes to roads affect cyclists, transit users, pedestrians, and drivers.](https://www.reddit.com/r/rust/comments/hffhxl/dabreegsterabstreet_a_traffic_simulation_game/)
- url: https://github.com/dabreegster/abstreet
---

## [6][toyDB: distributed SQL database in Rust, built from scratch to learn](https://www.reddit.com/r/rust/comments/hf6anm/toydb_distributed_sql_database_in_rust_built_from/)
- url: https://github.com/erikgrinaker/toydb
---

## [7][Tour of Rust - Let's go on an adventure! - Website redesign](https://www.reddit.com/r/rust/comments/hfeldt/tour_of_rust_lets_go_on_an_adventure_website/)
- url: http://tourofrust.com/
---

## [8][PR has been raised for `format_args_capture` feature](https://www.reddit.com/r/rust/comments/hfjv2f/pr_has_been_raised_for_format_args_capture_feature/)
- url: https://github.com/rust-lang/rust/pull/73670
---

## [9][What are the main takeaways for a Rust programmer to become familiar with C++?](https://www.reddit.com/r/rust/comments/hfhuuk/what_are_the_main_takeaways_for_a_rust_programmer/)
- url: https://www.reddit.com/r/rust/comments/hfhuuk/what_are_the_main_takeaways_for_a_rust_programmer/
---
I know Rust but am not at all familiar with C++, but whenever I see a snippet of code in C++ it seems easy enough to read which is really nice to see. I've been watching [this C++ programmer](https://www.twitch.tv/videos/659249134) pick it up over the past week and (though he's still learning) he took to it almost like it was nothing. ([Jason Turner](https://www.youtube.com/watch?v=Y5-ZgxfQvpc) picked up the basics pretty quickly too.) Meanwhile, watching them has given me a bit of familiarity with C++ and am curious what else there is to know. I'm wondering what the main things are I would need to know if I were sent to work with some C++ code next week, for example.

Here is some of what I've seen so far:

- Traits are basically like something called interfaces but differ in many respects

- Declarations are in reverse order (variable type then name) and adding const = move semantics while without const it's the same as mut

- C++ has lookup tables that it uses a lot (he wanted to make one but ended up using dynamic traits to put something similar together for practice and later on rewrote it as a match statement)

- If you take a C++ enum and add a union, you basically get a Rust enum (unions being the part that can hold values, I think)

- The auto keyword is like the i and u in isize and usize

- **-&gt;** takes values out of a struct in the same way the dot operator does, I think

- &amp; goes on the right instead of the left. Also apparently you can append it to the variable name but you can append it to the type instead if you prefer?

- Obviously no &amp;str and String, no borrow checker, no annotated lifetimes and no Result or Option

- C++ compiler messages are long and cryptic and you're supposed to scroll to the bottom to see what the actual problem is

- Various syntax differences like parentheses after if statements, no implied return at the end of functions, etc.


So that's about what I've gotten so far. Any other quick tips?
## [10][DataFrames in Rust](https://www.reddit.com/r/rust/comments/hfk83v/dataframes_in_rust/)
- url: https://www.reddit.com/r/rust/comments/hfk83v/dataframes_in_rust/
---
I started a mock up of a DataFrame library in Rust based on [Apache Arrow](https://github.com/apache/arrow). I started it just for fun, but I believe it is getting quite capable. After the first groupby benchmark it turns out we can already be more than 2x faster than pandas! 

https://preview.redd.it/a1w9chhyg1751.png?width=1400&amp;format=png&amp;auto=webp&amp;s=32817765aaf754c1afb0fee327e9336840c56f9f

[Project link](https://github.com/ritchie46/polars)
## [11][[ANN] RustCrypto's sha-1 and sha2 now support hardware acceleration on both x86 and ARM](https://www.reddit.com/r/rust/comments/hf2vcx/ann_rustcryptos_sha1_and_sha2_now_support/)
- url: https://www.reddit.com/r/rust/comments/hf2vcx/ann_rustcryptos_sha1_and_sha2_now_support/
---
x86 support was [added](https://github.com/RustCrypto/hashes/pull/167) in `sha-1 v0.9.1` and `sha2 v0.9.1`. It uses runtime detection (unlike [AES](https://github.com/RustCrypto/block-ciphers/issues/25#issuecomment-635789096)), so owners of CPUs with Intel SHA extension will benefit from it automatically. For example, on Ryzen 2700X both hash functions (SHA-1 and SHA-256) are able to process more than 2 GB/s.

Limited ARM support (Linux-only) was added earlier, but since the relevant intrinsics currently are not stable, we have to rely on assembly backend, which is disabled by default and feature-gated behind the `asm` feature.
## [12][What is the "Rust Way" of reading and parsing of JSON serial port as stream data?](https://www.reddit.com/r/rust/comments/hfg43l/what_is_the_rust_way_of_reading_and_parsing_of/)
- url: https://www.reddit.com/r/rust/comments/hfg43l/what_is_the_rust_way_of_reading_and_parsing_of/
---
Hello, a beginner here.

I have this application that reads from the serial port (using [serialport-rs](https://gitlab.com/susurrus/serialport-rs) crate). The data stream naturally coming in a sequence of 8 bytes. For example, if a message has a length of 150 bytes, we would be reading 18 messages of 8 bytes and 1 last message of 6 bytes if using the standard [std::io::read()](https://doc.rust-lang.org/std/io/trait.Read.html) function.

My data has this string format

    { node: "0x12", light: 0.0, object_temp: 25.375, ambient_temp: 25.09375, humidity: 35.155273, pressure: 97349.0 }

or if reading from with `String::from_utf8_lossy()`, we will have this format.

    "{\"node\":\"0x12\",\"light\":0.000000,\"object_temp\":25.375000,\"ambient_temp\":25.093750,\"humidity\":35.155273,\"pressure\":97349.000000}\n

I am using **serde\_json** to parse the JSON string. Due to the nature of not knowing exactly how many bytes the application will be receiving for a complete JSON string, I use `read_exact()` to fill a `Vec` of 150 bytes each. Then I split the JSON at the delimiter `"\n"`**,** and finally grab the first vector element to parse. There are a few problems with this approach that I don't know how to solve, hoping to gain some insights and knowledge from you guys.

1. reading the data into a fix byte array will leave the next data packet uncomplete, therefore losing it in the parsing process
2. Since this is a stream of data, there is no EOF to use functions such as `io::read_to_string()` or `io::read_to_end()` from `io::std` library

There is one possible approach, read each byte and insert it into a `Vec` until the "\\n" is detected, we will then parse the `Vec`, reset it, and so on. I've heard that resetting a `Vec` is quite slow and expensive. So this might be a bad approach.

Here is the snippet of the code.

&amp;#x200B;

https://preview.redd.it/6ozkg2s8iz651.png?width=1602&amp;format=png&amp;auto=webp&amp;s=d8f85e6ae63a017ccf4d55d213b983ab36b344f8

Thank you.
