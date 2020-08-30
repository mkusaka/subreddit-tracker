# rust
## [1][Hey Rustaceans! Got an easy question? Ask here (35/2020)!](https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/)
- url: https://www.reddit.com/r/rust/comments/ifiec2/hey_rustaceans_got_an_easy_question_ask_here/
---
Mystified about strings? Borrow checker have you in a headlock? Seek help here! There are no stupid questions, only docs that haven't been written yet.

If you have a [StackOverflow](http://stackoverflow.com/) account, consider asking it there instead! StackOverflow shows up much higher in search results, so having your question there also helps future Rust users (be sure to give it [the "Rust" tag](http://stackoverflow.com/questions/tagged/rust) for maximum visibility). Note that this site is very interested in question quality. I've been asked to read a RFC I authored once. If you want your code reviewed or review other's code, there's a [codereview stackexchange](https://codereview.stackexchange.com/questions/tagged/rust), too. If you need to test your code, maybe [the Rust playground](https://play.rust-lang.org) is for you.

Here are some other venues where help may be found:

[/r/learnrust](https://www.reddit.com/r/learnrust) is a subreddit to share your questions and epiphanies learning Rust programming.

The official Rust user forums: [https://users.rust-lang.org/](https://users.rust-lang.org/).

The official Rust Programming Language Discord: [https://discord.gg/rust-lang](https://discord.gg/rust-lang)

The unofficial Rust community Discord: [https://bit.ly/rust-community](https://bit.ly/rust-community)

Also check out [last weeks' thread](https://reddit.com/r/rust/comments/i6yqng/hey_rustaceans_got_an_easy_question_ask_here/) with many good questions and answers. And if you believe your question to be either very complex or worthy of larger dissemination, feel free to create a text post.

Also if you want to be mentored by experienced Rustaceans, tell us the area of expertise that you seek.
## [2][This Week in Rust 353](https://www.reddit.com/r/rust/comments/ih8uai/this_week_in_rust_353/)
- url: https://this-week-in-rust.org/blog/2020/08/26/this-week-in-rust-353/
---

## [3][John Carmack mentioning Rust as a modern choice when talking about programming languages and the impact they have](https://www.reddit.com/r/rust/comments/ij7i8q/john_carmack_mentioning_rust_as_a_modern_choice/)
- url: https://twitter.com/ID_AA_Carmack/status/1299574198365495297?s=09
---

## [4][Writing an asynchronous MQTT Broker in Rust - Part 3](https://www.reddit.com/r/rust/comments/ijajxs/writing_an_asynchronous_mqtt_broker_in_rust_part_3/)
- url: https://hassamuddin.com/blog/rust-mqtt/channels/
---

## [5][Official /r/rust "Who's Hiring" thread for job-seekers and job-offerers [Rust 1.46]](https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/)
- url: https://www.reddit.com/r/rust/comments/iix8vw/official_rrust_whos_hiring_thread_for_jobseekers/
---
Welcome once again to the official /r/rust Who's Hiring thread!

Before we begin, job-seekers should also remember to peruse [the prior thread](https://www.reddit.com/r/rust/comments/htnv19/official_rrust_whos_hiring_thread_for_jobseekers/).

This thread will be periodically stickied to the top of /r/rust for improved visibility. The thread will be refreshed and posted anew when the next version of Rust releases in six weeks.

Please adhere to the following rules when posting:
# Rules for individuals:

 * Don't create top-level comments; those are for employers.
 * Feel free to reply to top-level comments with on-topic questions.
 * Anyone seeking work should reply to my stickied top-level comment.
 * Meta-discussion should be reserved for the distinguished comment at the very bottom.

# Rules for employers:

 * To find individuals seeking work, see the replies to the stickied top-level comment; you will need to click the "more comments" link at the bottom of the top-level comment in order to make these replies visible.
 * To make a top-level comment you must be hiring directly; no third-party recruiters.
 * One top-level comment per employer. If you have multiple job openings, please consolidate their descriptions or mention them in replies to your own top-level comment.
 * Proofread your comment after posting it and edit it if necessary to correct mistakes.
 * Please base your comment on the following template:

COMPANY: *[Company name; please link to your company's website or careers page.]*

TYPE: *[Full time, part time, internship, contract, etc.]*

DESCRIPTION: *[What does your company do, and what are you using Rust for? How much experience are you seeking and what seniority levels are you hiring for? The more details the better.]*

LOCATION: *[Where are your office or offices located? If your workplace language isn't English, please specify it.]*

REMOTE: *[Do you offer the option of working remotely? If so, do you require employees to live in certain areas or time zones?]*

VISA: *[Does your company sponsor visas?]*

CONTACT: *[How can someone get in touch with you?]*
## [6][extend 0.2.1 just released](https://www.reddit.com/r/rust/comments/ij0q8i/extend_021_just_released/)
- url: https://crates.io/crates/extend
---

## [7][Looking for a realtime audio synthesis crate/lib that can output directly to the audio device](https://www.reddit.com/r/rust/comments/ij7ehx/looking_for_a_realtime_audio_synthesis_cratelib/)
- url: https://www.reddit.com/r/rust/comments/ij7ehx/looking_for_a_realtime_audio_synthesis_cratelib/
---
Hi everyone, 

as stated in the title, I'm looking for a library or crate that can generate, mix and filter arbitrary waveforms and output them on the sound card - a realtime software syntheshizer, so to say. 

Is there something like that for Rust that for example offers a collection of waveform generators/oscillators (sine, saw, square...) that can interact with each other, for instance, one controlling the amplitude, phase or frequency of the other, a group of filters (low pass, high pass, band pass...) and stuff like that - while capable of outputting the generated waveform to the audio device in (almost) realtime? 

Thanks
## [8][Best way to create Rust AST's?](https://www.reddit.com/r/rust/comments/ij9u9t/best_way_to_create_rust_asts/)
- url: https://www.reddit.com/r/rust/comments/ij9u9t/best_way_to_create_rust_asts/
---
I'm experimenting with a project to transpile code from another language into Rust.  The overall process would go something like this:

Lang Souce -&gt; Lang AST -&gt; Rust AST -&gt; Rust Source

I could write the whole thing from scratch, but it seems as if building a Rust AST programmatically and converting this to source must already be something which exists (especially since this is possible in macro definitions).

What would be the best way to do this?  I'm really just looking to be pointed in the right direction on this topic, and if I understand the big picture, I can do my own research to figure out the details.
## [9][how to create a stream from a BytesMut](https://www.reddit.com/r/rust/comments/ijbro8/how_to_create_a_stream_from_a_bytesmut/)
- url: https://www.reddit.com/r/rust/comments/ijbro8/how_to_create_a_stream_from_a_bytesmut/
---
I am reading from `STDIN` into a big buffer (`BytesMut`) in memory, I am using something like this:

    use anyhow::Result;
    use bytes::{BufMut, BytesMut};
    use futures::stream::TryStreamExt;
    use tokio::io::stdin;
    use tokio_util::codec::{BytesCodec, FramedRead};
    
    #[tokio::main]
    async fn main() -&gt; Result&lt;()&gt; {
        let mut buffer = BytesMut::with_capacity(1024 * 1024 * 512);
        let mut input_stream = FramedRead::new(stdin(), BytesCodec::new());
        while let Some(bytes) = input_stream.try_next().await? {
            buffer.put(bytes);
        }
        println!("buf: {}", buffer.len());
        Ok(())
    }

`buffer` now contains the full input, but how can I create a stream from it

I am trying to implement the Stream trait like this:

    use std::pin::Pin;
    use std::task::{Context, Poll};
    
    #[derive(Debug)]
    struct ByteStream {
        buffer: BytesMut,
    }
    
    impl ByteStream {
        pub fn new(buffer: BytesMut) -&gt; Self {
            Self { buffer }
        }
    }
    
    impl Stream for ByteStream {
        type Item = Bytes;
    
        fn poll_next(mut self: Pin&lt;&amp;mut Self&gt;, _cx: &amp;mut Context&lt;'_&gt;) -&gt; Poll&lt;Option&lt;Self::Item&gt;&gt; {
            if self.buffer.is_empty() {
                Poll::Ready(None)
            } else {
                let buf_len = self.buffer.len();
                let out = if buf_len &gt; 8192 {
                    self.buffer.split_to(8192)
                } else {
                    self.buffer.split_to(buf_len)
                };
                Poll::Ready(Some(out.freeze()))
            }
        }
    }

I can iterate over it:

    let foo = ByteStream::new(buffer);
    
    while let Some(bytes) = foo.next().await {
        println!("{}", bytes.len());
    }

But when I want to use it with `reqwest`:

     let client = reqwest::Client::new();
     let body = reqwest::Body::wrap_stream(stream);
     client.post().body(body)

I get this error:

    error[E0271]: type mismatch resolving `&lt;ByteStream as tokio::stream::Stream&gt;::Item == std::result::Result&lt;_, _&gt;`
      --&gt; src/main.rs:70:16
       |
    70 |     let body = Body::wrap_stream(x);
       |                ^^^^^^^^^^^^^^^^^ expected struct `bytes::Bytes`, found enum `std::result::Result`
       |
       = note: expected struct `bytes::Bytes`
                    found enum `std::result::Result&lt;_, _&gt;`
       = note: required because of the requirements on the impl of `futures_core::stream::TryStream` for `ByteStream`

Any idea of how could I create the `Stream` and use it as an argument for `wrap_stram`?
## [10][Library consts config](https://www.reddit.com/r/rust/comments/ijb766/library_consts_config/)
- url: https://www.reddit.com/r/rust/comments/ijb766/library_consts_config/
---
I have a library with a number of constant integers.

These constants may vary from use case to use case.

But I can't find a way to easily configure these constants.

Currently the only way I see is to clone the project and change them. I've looked at build scripts but I can't find a way to pass an argument to them either.

Anyone know how I can fix this?
## [11][About this guide - Guide to Rustc Development](https://www.reddit.com/r/rust/comments/iizaz5/about_this_guide_guide_to_rustc_development/)
- url: https://rustc-dev-guide.rust-lang.org/
---

## [12][A new version of Ruby Firebird Extension Library using Rust instead of C](https://www.reddit.com/r/rust/comments/ijcf31/a_new_version_of_ruby_firebird_extension_library/)
- url: https://github.com/fernandobatels/rbfbclient
---

